package controller

import (
	"encoding/json"
	"net/http"

	"github.com/hivanreyes/academy-go-q12021/model"
)

type UseCase interface {
	ReadPokemon() ([]model.Pokemon, error)
	SavePokemon() ([]model.Pokemon, error)
	ReadConcurrentPokemon() ([]model.Pokemon, error)
}

// Usecase struct
type PokeUsecase struct {
	useCase UseCase
}

// Initialize controller
func New(u UseCase) *PokeUsecase {
	return &PokeUsecase{u}
}

// Read Pokemon controller
func (u *PokeUsecase) ReadPokemon(w http.ResponseWriter, r *http.Request) {

	pokemon, err := u.useCase.ReadPokemon()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(pokemon)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Save Pokemon controller
func (u *PokeUsecase) SavePokemon(w http.ResponseWriter, r *http.Request) {
	pokemon, err := u.useCase.SavePokemon()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(pokemon)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// Read Pokemon Concurrently controller
func (u *PokeUsecase) ReadConcurrentPokemon(w http.ResponseWriter, r *http.Request) {
	typeItem := r.URL.Query().Get("type")
	items := r.URL.Query().Get("items")
	itemPerWorker := r.URL.Query().Get("items_per_workers")

	if typeItem == "" || items == "" || itemPerWorker == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pokemon, err := u.useCase.ReadConcurrentPokemon(typeItem, items, itemPerWorker)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(pokemon)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
