package routes

import (
	"github.com/gorilla/mux"
	"github.com/hivanreyes/academy-go-q12021/pkg/controllers"
)

var RegisterPokemonRoutes = func(router *mux.Router) {
	router.HandleFunc("/getAllPokemon", controllers.GetPokemon).Methods("GET")
}
