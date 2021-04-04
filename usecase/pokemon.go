package usecase

import "github.com/hivanreyes/academy-go-q12021/model"

type UseCase struct {
	service Service
}

// Service interface
type Service interface {
	ReadPokemon() ([]model.Pokemon, error)
	SavePokemon() ([]model.Pokemon, error)
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
