package models

import (
	"context"
	"time"

	"safehaven.com/m/db"
)

type Organization struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (o Organization) Save() (Organization, error) {
	query := `INSERT INTO organization (name, description)
            VALUES ($1,$2) RETURNING id, createdAt`

	row := db.DB.QueryRow(context.Background(), query, o.Name, o.Description)
	err := row.Scan(&o.ID, &o.CreatedAt)

	if err != nil {
		return Organization{}, err
	}

	return o, nil
}

func GetOrgs() ([]Organization, error) {
	query := `SELECT * FROM organization`

	rows, err := db.DB.Query(context.Background(), query)

	if err != nil {
		return []Organization{}, err
	}

	defer rows.Close()

	var orgs []Organization

	for rows.Next() {
		var org Organization

		err := rows.Scan(&org.ID, &org.Name, &org.Description, &org.CreatedAt)

		if err != nil {
			return []Organization{}, err
		}

		orgs = append(orgs, org)
	}

	if err := rows.Err(); err != nil {
		return []Organization{}, err
	}

	return orgs, nil
}
