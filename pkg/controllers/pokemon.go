package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/hivanreyes/academy-go-q12021/pkg/models"
)

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	pokemon := models.GetAllPokemon()
	res, _ := json.Marshal(pokemon)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func PopulatePokemon(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	pokemon := models.PopulateAllPokemon(limit, offset)
	res, _ := json.Marshal(pokemon)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
