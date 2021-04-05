package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hivanreyes/academy-go-q12021/model"
)

type UseCase interface {
	ReadPokemon() ([]model.Pokemon, error)
	SavePokemon() ([]model.Pokemon, error)
	ReadConcurrentPokemon(typeItem string, items int, itemPerWorker int) ([]model.Pokemon, error)
}

// PokeUsecase usecase struct
type PokeUsecase struct {
	useCase UseCase
}

// New initialize controller
func New(u UseCase) *PokeUsecase {
	return &PokeUsecase{u}
}

// ReadPokemon Read Pokemon controller
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

// SavePokemon Save Pokemon controller
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

// ReadConcurrentPokemon Read Pokemons concurrently controller
func (u *PokeUsecase) ReadConcurrentPokemon(w http.ResponseWriter, r *http.Request) {
	typeItem := r.URL.Query().Get("type")
	items := r.URL.Query().Get("items")
	itemPerWorker := r.URL.Query().Get("items_per_workers")

	itemsNumber, err := strconv.Atoi(items)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	itemsWorker, err := strconv.Atoi(itemPerWorker)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if typeItem == "" || items == "" || itemPerWorker == "" || itemsWorker > itemsNumber {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	pokemon, err := u.useCase.ReadConcurrentPokemon(typeItem, itemsNumber, itemsWorker)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, _ := json.Marshal(pokemon)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
