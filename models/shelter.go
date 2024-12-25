package models

import (
	"context"
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
