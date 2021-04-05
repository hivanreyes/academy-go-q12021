package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	ReadPokemon(w http.ResponseWriter, r *http.Request)
	SavePokemon(w http.ResponseWriter, r *http.Request)
	ReadConcurrentPokemon(w http.ResponseWriter, r *http.Request)
}

// New initialize Router
func New(c Controller) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/getAllPokemon", c.ReadPokemon).Methods("GET")
	r.HandleFunc("/populateAllPokemon", c.SavePokemon).Methods("GET")
	r.HandleFunc("/getConcurrentPokemon", c.ReadConcurrentPokemon).Methods("GET")

	return r
}
