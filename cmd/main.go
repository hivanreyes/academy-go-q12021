package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hivanreyes/academy-go-q12021/pkg/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()
	routes.RegisterPokemonRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))
}
