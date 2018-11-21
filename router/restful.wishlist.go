package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"

	"github.com/joedha8/AnimeStore/model"
)

func GetAllWishlist(w http.ResponseWriter, r *http.Request) {
	wishlists, err := model.GetAllWishlist(r.Context(), DbPool)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string][]model.WishlistModel{"wishlist": wishlists})
}

func GetOneWishlist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	wishlist, err := model.GetOneWishlist(r.Context(), DbPool, id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]model.WishlistModel{"wishlist": wishlist})
}

func InsertWishlist(w http.ResponseWriter, r *http.Request) {
	var wishlist model.WishlistModel
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&wishlist); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	id, err := wishlist.Insert(r.Context(), DbPool)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
	}
	wishlist.ID = id

	respondWithJSON(w, http.StatusCreated, map[string]model.WishlistModel{"wishlist": wishlist})
}

func UpdateWishlist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var wishlist model.WishlistModel
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&wishlist); err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	wishlist.ID = id

	if err := wishlist.Update(r.Context(), DbPool); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"wishlist": "success"})
}

func DeleteWishlist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.FromString(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := model.DeleteWishlist(r.Context(), DbPool, id); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"wishlist": "success"})
}
