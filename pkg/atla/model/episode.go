package model

import (
	//"context"
	"database/sql"
	"log"
	//"time"
)

type Episode struct {
	ID        int    `json:"id"`
	Title     string `json:"name"`
	Air_Date  string `json:"air_date"`
	CreatedAt string `json:"createdAt"`
	ApdatedAt string `json:"updatedAt"`
}
type EpisodeModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
