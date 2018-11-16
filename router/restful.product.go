package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"

	"github.com/joedha8/AnimeStore/model"
)

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	products, err := model.GetAllProduct(r.Context(), DbPool)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string][]model.ProductModel{"produtc": products})
}

func GetOneProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	product, err := model.GetOneProduct(r.Context(), DbPool, id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]model.ProductModel{"product": product})
}

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	var product model.ProductModel
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	id, err := product.Insert(r.Context(), DbPool)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	product.ID = id

	respondWithJSON(w, http.StatusCreated, map[string]model.ProductModel{"product": product})
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var product model.ProductModel
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	product.ID = id

	if err := product.Update(r.Context(), DbPool); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"product": "success"})
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := model.DeleteProduct(r.Context(), DbPool, id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"product": "success"})
}
