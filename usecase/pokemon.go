package usecase

import (
	"github.com/hivanreyes/academy-go-q12021/model"
)

type UseCase struct {
	service Service
}

// Service interface
type Service interface {
	ReadPokemon() ([]model.Pokemon, error)
	SavePokemon() ([]model.Pokemon, error)
	ReadConcurrentPokemon(typeItem string, items string, itemPerWorker string, pokemons []model.Pokemon) ([]model.Pokemon, error)
}

// New UseCase
func New(service Service) *UseCase {
	return &UseCase{service}
}

// ReadPokemon Usecase
func (u *UseCase) ReadPokemon() ([]model.Pokemon, error) {
	return u.service.ReadPokemon()
}

// SavePokemon Usecase
func (u *UseCase) SavePokemon() ([]model.Pokemon, error) {
	return u.service.SavePokemon()
}

// ReadPokemon Usecase
func (u *UseCase) ReadConcurrentPokemon(typeItem string, items string, itemPerWorker string) ([]model.Pokemon, error) {
	pokemons, err := u.service.ReadPokemon()

	if err != nil {
		return nil, err
	}

	return u.service.ReadConcurrentPokemon(typeItem, items, itemPerWorker, pokemons)
}
