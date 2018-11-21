package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"

	"github.com/joedha8/AnimeStore/model"
)

func GetAllCategory(w http.ResponseWriter, r *http.Request) {
	categories, err := model.GetAllCategory(r.Context(), DbPool)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string][]model.CategoryModel{"category": categories})
}

func GetOneCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	category, err := model.GetOneCategory(r.Context(), DbPool, id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]model.CategoryModel{"category": category})
}

func InsertCategory(w http.ResponseWriter, r *http.Request) {
	var category model.CategoryModel
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	id, err := category.Insert(r.Context(), DbPool)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	category.ID = id

	respondWithJSON(w, http.StatusCreated, map[string]model.CategoryModel{"category": category})
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var category model.CategoryModel
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&category); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	category.ID = id

	if err := category.Update(r.Context(), DbPool); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"category": "success"})
}

func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := model.DeleteCategory(r.Context(), DbPool, id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"category": "success"})
}
