/*!!!!!!!!!!!!!!!!!! HAS CRUD !!!!!!!!!!!!!!!!!!*/

package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/justverena/ATLA/pkg/atla/model"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func (app *application) createCharacterHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		Age       int    `json:"age"`
		Gender    string `json:"gender"`
		Status    string `json:"status"`
		Nation    string `json:"nation"`
		CreatedAt string `json:"createdAt"`
		UpdatedAt string `json:"updatedAt"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	character := &model.Character{
		ID:        input.ID,
		Name:      input.Name,
		Age:       input.Age,
		Gender:    input.Gender,
		Status:    input.Status,
		Nation:    input.Nation,
		CreatedAt: input.CreatedAt,
		UpdatedAt: input.UpdatedAt,
	}

	err = app.models.Characters.Insert(character)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusCreated, character)
}

func (app *application) getCharacterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["characterID"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid character ID")
		return
	}

	character, err := app.models.Characters.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	app.respondWithJSON(w, http.StatusOK, character)
}

func (app *application) updateCharacterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["characterID"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid character ID")
		return
	}

	character, err := app.models.Characters.Get(id)
	if err != nil {
		app.respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	var input struct {
		CharacterName *string `json:"characterName"`
		Age           *int    `json:"age"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if input.CharacterName != nil {
		character.Name = *input.CharacterName
	}

	if input.Age != nil {
		character.Age = *input.Age
	}

	err = app.models.Characters.Update(character)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, character)
}

func (app *application) deleteCharacterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["characterID"]

	id, err := strconv.Atoi(param)
	if err != nil || id < 1 {
		app.respondWithError(w, http.StatusBadRequest, "Invalid character ID")
		return
	}

	err = app.models.Characters.Delete(id)
	if err != nil {
		app.respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	app.respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (app *application) readJSON(_ http.ResponseWriter, r *http.Request, dst interface{}) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(dst)
	if err != nil {
		return err
	}

	return nil
}
