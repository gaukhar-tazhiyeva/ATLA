/*!!!!!!!!!!!!!!!!!! DID NOT ADD CRUD YET !!!!!!!!!!!!!!!!!!*/

package model

import (
	//"context"
	"database/sql"
	"log"
	//"time"
)

type Episode struct {
	ID           int    `json:"id"`
	Season_ID    int    `json:"season_id"`
	Title        string `json:"title"`
	Character_ID int    `json: "character_id"`
	CreatedAt    string `json:"createdAt"`
	ApdatedAt    string `json:"updatedAt"`
}

type EpisodeModel struct {
	DB       *sql.DB
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}
