package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"

	"github.com/joedha8/AnimeStore/model"
)

func GetAllDetailOrder(w http.ResponseWriter, r *http.Request) {
	detailOrders, err := model.GetAllDetailOrder(r.Context(), DbPool)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string][]model.DetailOrderModel{"detail_order": detailOrders})
}

func GetOneDetailOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	detailOrder, err := model.GetOneDetailOrder(r.Context(), DbPool, id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]model.DetailOrderModel{"detail_order": detailOrder})
}

func InsertDetailOrder(w http.ResponseWriter, r *http.Request) {
	var detailOrder model.DetailOrderModel
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&detailOrder); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	id, err := detailOrder.Insert(r.Context(), DbPool)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	detailOrder.ID = id

	respondWithJSON(w, http.StatusCreated, map[string]model.DetailOrderModel{"detail_order": detailOrder})
}

func UpdateDetailOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var detailOrder model.DetailOrderModel
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&detailOrder); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	detailOrder.ID = id

	if err := detailOrder.Update(r.Context(), DbPool); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"detail_order": "success"})
}

func DeleteDetailOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := model.DeleteDetailOrder(r.Context(), DbPool, id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"product": "success"})
}
