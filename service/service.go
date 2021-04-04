package service

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/go-resty/resty/v2"

	model "github.com/hivanreyes/academy-go-q12021/model"
	"github.com/hivanreyes/academy-go-q12021/utils"
)

const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon"
const newFileName = "data/pokemon1.csv"
const filename = "data/pokemon.csv"

// Service
type Service struct{}

// New Service
func New() *Service {
	return &Service{}
}

// Get all pokemons
func (s *Service) ReadPokemon() ([][]string, error) {
	// Open CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Skip first row (line)
	row1, err := bufio.NewReader(file).ReadSlice('\n')
	if err != nil {
		return nil, err
	}

	_, err = file.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		return nil, err
	}

	// Read File into a Variable
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, nil
}

func getDataFromApi() (*model.Response, error) {
	data := &model.Response{Result: &[]model.PokemonApi{}}

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"limit":  "100",
			"offset": "200",
		}).
		SetHeader("Accept", "application/json").
		Get(pokeApiUrl)

	if err != nil {
		fmt.Println("Error getting pokemons")
		panic(err)
	}

	err = json.Unmarshal([]byte(resp.Body()), data)

	if err != nil {
		fmt.Println("Error parsing pokemons")
		panic(err)
	}

	return data, nil
}

// Get data from api and save data into CSV
func (s *Service) SavePokemon() error {
	data, errData := getDataFromApi()

	if errData != nil {
		return errData
	}

	utils.RemoveFile(newFileName)

	f, err := os.Create(newFileName)

	if err != nil {
		return err
	}

	defer f.Close()

	if err != nil {
		fmt.Println("Error creating csv")
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	// Add headers to the csv
	if err := w.Write([]string{"id", "name"}); err != nil {
		fmt.Println("Error adding titles to csv")
	}

	for _, pokemon := range *data.Result {
		// Getting pokemon ID
		re, _ := regexp.Compile("/pokemon/(.*)/")
		values := re.FindStringSubmatch(pokemon.Url)
		var id = values[1]
		var name = pokemon.Name
		if err := w.Write([]string{id, name}); err != nil {
			fmt.Println("Error adding pokemon to csv")
		}
	}

	return nil
}
