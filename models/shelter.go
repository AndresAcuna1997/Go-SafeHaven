package models

import (
	"context"
	"errors"
	"fmt"
	"time"

	"safehaven.com/m/db"
)

type Shelter struct {
	ID             int64
	Name           string `binding:"required"`
	Description    string `binding:"required"`
	Address        string `binding:"required"`
	RefugeeCount   int    `json:"refugee_count" binding:"required"`
	ContactPhone   string `json:"contact_phone" binding:"required"`
	ContactEmail   string `json:"contact_email" binding:"required"`
	CreatedAt      time.Time
	OrganizationId int64 `json:"organization_id"`
	City           int64 `binding:"required"`
}

func (s Shelter) Save() (Shelter, error) {
	query := `INSERT INTO shelter (name, description, address, refugeescount, contactphone, contactemail, organization_id, city)
						VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id, createdAt`

	row := db.DB.QueryRow(context.Background(), query, s.Name, s.Description, s.Address, s.RefugeeCount, s.ContactPhone, s.ContactEmail, s.OrganizationId, s.City)
	err := row.Scan(&s.ID, &s.CreatedAt)

	if err != nil {
		return Shelter{}, err
	}

	return s, nil
}

func (s Shelter) Update() (Shelter, error) {
	query := `UPDATE shelter
						SET name = $1, 
						description = $2,
						address =$3,
						refugeescount=$4, 
						contactphone=$5, 
						contactemail=$6, 
						city=$7
						WHERE id = $8`

	commandTag, err := db.DB.Exec(context.Background(), query,
		s.Name,
		s.Description,
		s.Address,
		s.RefugeeCount,
		s.ContactPhone,
		s.ContactEmail,
		s.City,
		s.ID)

	if err != nil {
		return Shelter{}, err
	}

	if commandTag.RowsAffected() == 0 {
		return Shelter{}, errors.New("no filas fueron afectadas")
	}

	return s, nil
}

func (s Shelter) Delete() error {
	query := `DELETE FROM shelter WHERE id = $1`

	cmdTag, err := db.DB.Exec(context.Background(), query, s.ID)

	if err != nil {
		fmt.Println("1", err)
		return errors.New("no se pudo borrar el refugio")
	}

	if cmdTag.RowsAffected() == 0 {
		fmt.Println("1", err)
		return errors.New("no filas fueron afectadas")
	}

	return nil
}

func GetAllShelters() ([]Shelter, error) {
	query := `SELECT * FROM shelter`

	rows, err := db.DB.Query(context.Background(), query)

	if err != nil {
		return []Shelter{}, err
	}

	var shelters []Shelter

	for rows.Next() {
		var shelter Shelter
		err := rows.Scan(&shelter.ID, &shelter.Name, &shelter.Description, &shelter.Address, &shelter.RefugeeCount, &shelter.ContactPhone, &shelter.ContactEmail, &shelter.CreatedAt, &shelter.OrganizationId, &shelter.City)

		if err != nil {
			return nil, err
		}

		shelters = append(shelters, shelter)
	}

	if err = rows.Err(); err != nil {
		return []Shelter{}, err
	}

	return shelters, nil
}

func GetSingleShelter(shelterId int64) (Shelter, error) {

	var shelter Shelter

	query := `SELECT * FROM shelter WHERE id = $1`

	row := db.DB.QueryRow(context.Background(), query, shelterId)

	err := row.Scan(&shelter.ID, &shelter.Name, &shelter.Description, &shelter.Address, &shelter.RefugeeCount, &shelter.ContactPhone, &shelter.ContactEmail, &shelter.CreatedAt, &shelter.OrganizationId, &shelter.City)

	if err != nil {
		fmt.Println(err)
		return Shelter{}, err
	}

	return shelter, nil
}
