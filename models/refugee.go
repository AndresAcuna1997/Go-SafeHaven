package models

import (
	"context"
	"encoding/json"
	"time"

	"safehaven.com/m/db"
)

type Refugee struct {
	ID             int             `json:"id"`
	Name           string          `json:"name"`
	RefugeeType    string          `json:"refugeeType"`
	Size           string          `json:"size"`
	Age            int             `json:"age"`
	AdditionalInfo json.RawMessage `json:"additionalInfo"`
	Pictures       json.RawMessage `json:"pictures"`
	CreatedAt      time.Time       `json:"createdAt"`
	ShelterId      int             `json:"shelterId"`
}

func (r Refugee) Save() (Refugee, error) {
	query := `INSERT INTO refugee (name,type,size,age,additionalInfo, pictures, shelter_id)
						VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id, createdAt`

	row := db.DB.QueryRow(context.Background(), query, r.Name, r.RefugeeType, r.Size, r.Age, r.AdditionalInfo, r.Pictures, r.ShelterId)
	err := row.Scan(&r.ID, &r.CreatedAt)

	if err != nil {
		return Refugee{}, err
	}

	return r, nil
}

func GetAllRefugees() ([]Refugee, error) {
	query := `SELECT * FROM refugee`

	rows, err := db.DB.Query(context.Background(), query)

	if err != nil {
		return []Refugee{}, nil
	}

	var refugees []Refugee

	for rows.Next() {
		var refugee Refugee

		err := rows.Scan(&refugee.ID, &refugee.Name, &refugee.RefugeeType, &refugee.Size, &refugee.Age, &refugee.AdditionalInfo, &refugee.Pictures, &refugee.CreatedAt, &refugee.ShelterId)

		if err != nil {
			return nil, err
		}

		refugees = append(refugees, refugee)
	}

	if err = rows.Err(); err != nil {
		return []Refugee{}, err
	}
	return refugees, nil
}
