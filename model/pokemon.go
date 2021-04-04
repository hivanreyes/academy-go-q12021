package model

// Pokemon struct
type Pokemon struct {
	Id   string
	Name string
}

// PokemonApi struct
type PokemonApi struct {
	Name string
	Url  string
}

// Response struct
type Response struct {
	Count    int           `json:"count,omitempty"`
	Next     string        `json:"next,omitempty"`
	Previous string        `json:"previous,omitempty"`
	Result   *[]PokemonApi `json:"results"`
}
