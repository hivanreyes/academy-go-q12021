package usecase

type UseCase struct {
	service Service
}

// Service interface
type Service interface {
	ReadPokemon() ([][]string, error)
	SavePokemon() error
}

// New UseCase
func New(service Service) *UseCase {
	return &UseCase{service}
}

// ReadPokemon Usecase
func (u *UseCase) ReadPokemon() ([][]string, error) {
	return u.service.ReadPokemon()
}

// SavePokemon Usecase
func (u *UseCase) SavePokemon() error {
	return u.service.SavePokemon()
}
