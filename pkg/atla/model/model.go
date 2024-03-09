/*!!!!!!!!!!!!!!!!!! CRUD IN HANDLERS !!!!!!!!!!!!!!!!!!*/

package model

import (
	"database/sql"
	"log"
	"os"
)

type Models struct {
	Characters CharacterModel
	Episodes   EpisodeModel
}

func NewModels(db *sql.DB) Models {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	return Models{
		Characters: CharacterModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
		Episodes: EpisodeModel{
			DB:       db,
			InfoLog:  infoLog,
			ErrorLog: errorLog,
		},
	}
}
