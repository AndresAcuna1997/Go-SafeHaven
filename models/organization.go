package models

import (
	"context"
	"errors"
	"time"

	"safehaven.com/m/db"
	"safehaven.com/m/utils"
)

type Organization struct {
	ID          int64
	Name        string
	Email       string `binding:"required"`
	Password    string `binding:"required"`
	Description string
	CreatedAt   time.Time
}

func (o Organization) Save() (Organization, error) {
	query := `INSERT INTO organization (name, description, email, password)
            VALUES ($1,$2,$3,$4) RETURNING id, createdAt`

	hashedPass, err := utils.HashPassword(o.Password)

	if err != nil {
		return Organization{}, err
	}

	o.Password = hashedPass

	row := db.DB.QueryRow(context.Background(), query, o.Name, o.Description, o.Email, o.Password)
	err = row.Scan(&o.ID, &o.CreatedAt)

	if err != nil {
		return Organization{}, err
	}

	return o, nil
}

func GetOrgs() ([]Organization, error) {
	query := `SELECT id, name, description, email, createdAt FROM organization`

	rows, err := db.DB.Query(context.Background(), query)

	if err != nil {
		return []Organization{}, err
	}

	defer rows.Close()

	var orgs []Organization

	for rows.Next() {
		var org Organization

		err := rows.Scan(&org.ID, &org.Name, &org.Description, &org.Email, &org.CreatedAt)

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

func (o *Organization) ValidateCredential() error {
	query := `SELECT id, password FROM organization WHERE email = $1`

	row := db.DB.QueryRow(context.Background(), query, o.Email)

	var dbPassword string

	err := row.Scan(&o.ID, &dbPassword)

	if err != nil {
		return errors.New("no existe una cuenta con ese email")
	}

	if !utils.CheckPassword(o.Password, dbPassword) {
		return errors.New("credenciales invalidas")
	}

	return nil
}
