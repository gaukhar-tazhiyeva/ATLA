package model

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Character struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Gender    string `json:"gender"`
	Status    string `json:"status"`
	Nation    string `json:"nation"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CharacterModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func (m CharacterModel) Insert(character *Character) error {
	query := `
		INSERT INTO characters (name, age, gender, status, nation) 
		VALUES ($1, $2, $3) 
		RETURNING id, created_at, updated_at
		`
	args := []interface{}{character.Name, character.Age, character.Gender, character.Status, character.Nation}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&character.ID, &character.CreatedAt, &character.UpdatedAt)
}

func (m CharacterModel) Get(id int) (*Character, error) {
	query := `
		SELECT id, name, age, gender, status, nation, created_at, updated_at 
		FROM characters
		WHERE id = $1
		`
	var character Character
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&character.ID, &character.Name, &character.Age, &character.Gender, &character.Status, &character.Nation, &character.CreatedAt, &character.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &character, nil
}

func (m CharacterModel) Update(character *Character) error {
	query := `
		UPDATE characters
		SET name = $1, age = $2, gender = $3, status = $4, nation = $5
		WHERE id = $6
		RETURNING updated_at
		`
	args := []interface{}{character.Name, character.Age, character.Gender, character.Status, character.Nation, character.ID}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&character.UpdatedAt)
}

func (m CharacterModel) Delete(id int) error {
	query := `
		DELETE FROM characters
		WHERE id = $1
		`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}
