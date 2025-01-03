package models

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"safehaven.com/m/db"
)

type Refugee struct {
	ID             int64
	Name           string          `binding:"required"`
	RefugeeType    string          `json:"refugee_type" binding:"required"`
	Size           string          `binding:"required"`
	Age            int64           `binding:"required"`
	AdditionalInfo json.RawMessage `json:"additional_info" binding:"required"`
	Pictures       json.RawMessage `binding:"required"`
	CreatedAt      time.Time
	ShelterId      int64 `json:"shelter_id"`
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

func (r Refugee) Update() (Refugee, error) {
	query := `UPDATE refugee SET name = $1,
					type = $2, size = $3, age = $4,
					additionalinfo =$5, pictures =$6 	WHERE id = $7`

	cmndTag, err := db.DB.Exec(context.Background(), query, r.Name, r.RefugeeType, r.Size, r.Age, r.AdditionalInfo, r.Pictures, r.ID)

	if err != nil {
		return Refugee{}, errors.New("no se pudo actualizar el animalito")
	}

	if cmndTag.RowsAffected() == 0 {
		return Refugee{}, errors.New("no se afecto ninguna fila")
	}

	return r, nil
}

func (r Refugee) Delete() error {
	query := `DELETE FROM refugee WHERE id = $1`

	cmdTag, err := db.DB.Exec(context.Background(), query, r.ID)

	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return errors.New("no se afecto ninguna fila")
	}

	return nil
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

func FindSingleRefugee(id int64) (Refugee, error) {
	//Crear Struct
	var refugee Refugee
	//Crear Query
	query := `SELECT * from refugee WHERE id = $1`

	row := db.DB.QueryRow(context.Background(), query, id)

	err := row.Scan(&refugee.ID, &refugee.Name, &refugee.RefugeeType, &refugee.Size, &refugee.Age, &refugee.AdditionalInfo, &refugee.Pictures, &refugee.CreatedAt, &refugee.ShelterId)

	if err != nil {
		fmt.Println(err)
		return Refugee{}, errors.New("error obteniendo el animalito")
	}

	return refugee, nil
}
