package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/justverena/ATLA/pkg/atla/model"
	"github.com/justverena/ATLA/pkg/atla/validator"
)

func (app *application) createEpisodeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Air_Date string `json:"air_date"`
		// CreatedAt string `json:"createdAt"`
		// UpdatedAt string `json:"updatedAt"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		log.Println(err)
		app.errorResponse(w, r, http.StatusBadRequest, "Invalid request payload")
		return
	}

	episode := &model.Episode{
		ID:       input.ID,
		Title:    input.Title,
		Air_Date: input.Air_Date,
		// CreatedAt: input.CreatedAt,
		// UpdatedAt: input.UpdatedAt,
	}

	err = app.models.Episodes.Insert(episode)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, envelope{"episodes": episode}, nil)
}

func (app *application) getEpisodeList(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title string
		model.Filters
	}
	v := validator.New()
	qs := r.URL.Query()

	input.Title = app.readStrings(qs, "title", "")

	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readStrings(qs, "sort", "id")

	input.Filters.SortSafeList = []string{
		"id", "title",
		"-id", "-title",
	}

	if model.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	episodes, metadata, err := app.models.Episodes.GetAll(input.Title, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"episodes": episodes, "metadata": metadata}, nil)
}

func (app *application) getEpisodeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	episode, err := app.models.Episodes.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"episodes": episode}, nil)
}

func (app *application) updateEpisodeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	episode, err := app.models.Episodes.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		ID       *int    `json:"id"`
		Title    *string `json:"title"`
		Air_Date *string `json:"air_date"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Title != nil {
		episode.Title = *input.Title
	}

	if input.Air_Date != nil {
		episode.Air_Date = *input.Air_Date
	}
	v := validator.New()

	if model.ValidateEpisode(v, episode); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.Episodes.Update(episode)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"episodes": episode}, nil)
}

func (app *application) deleteEpisodeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Episodes.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, model.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"message": "success"}, nil)
}

func (app *application) getCharacterEpisode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	characterID, err := strconv.Atoi(vars["id"])
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, "Invalid character ID")
		return
	}

	character, err := app.models.Characters.Get(characterID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	episode, err := app.models.Episodes.GetByCharacter(characterID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"character": character}, nil)

	app.writeJSON(w, http.StatusOK, envelope{"episode": episode}, nil)
}
