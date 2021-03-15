package models

import (
	"fmt"

	"github.com/hivanreyes/academy-go-q12021/pkg/utils"
)

type Pokemon struct {
	Id   string
	Name string
}

const fileName = "data/example.csv"

func GetAllPokemon() []Pokemon {
	lines, err := utils.ReadCsv(fileName)
	if err != nil {
		fmt.Println("Error loading csv")
		panic(err)
	}

	var data []Pokemon

	// Loop through lines and append to data
	for _, line := range lines {
		item := Pokemon{
			Id:   line[0],
			Name: line[1],
		}

		data = append(data, item)
	}

	return data
}
