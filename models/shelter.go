package models

import (
	"context"
	"time"

	"safehaven.com/m/db"
)

type Shelter struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Address        string    `json:"address"`
	RefugeeCount   int       `json:"refugee_count"`
	ContactPhone   string    `json:"contact_phone"`
	ContactEmail   string    `json:"contact_email"`
	CreatedAt      time.Time `json:"created_at"`
	OrganizationId int       `json:"organization_id"`
	City           int       `json:"city"`
}

func (s Shelter) Save() (Shelter, error) {
	query := `INSERT INTO shelter (name, description, address, refugeescount, contactphone, contactemail, organization_id)
						VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id, createdAt`

	row := db.DB.QueryRow(context.Background(), query, s.Name, s.Description, s.Address, s.RefugeeCount, s.ContactPhone, s.ContactEmail, s.OrganizationId)
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
