package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"regexp"

	"github.com/go-resty/resty/v2"
)

type Response struct {
	Count    int        `json:"count,omitempty"`
	Next     string     `json:"next,omitempty"`
	Previous string     `json:"previous,omitempty"`
	Result   *[]Pokemon `json:"results"`
}

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

const pokeApiUrl = "https://pokeapi.co/api/v2/pokemon"
const newFileName = "data/pokemon.csv"

func getDataFromApi(limit string, offset string) *Response {
	data := &Response{Result: &[]Pokemon{}}

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"limit":  limit,
			"offset": offset,
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

	return data
}

func removeFile(filename string) {
	err := os.Remove(filename)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func PopulatePokemon(limit string, offset string) (string, error) {
	data := getDataFromApi(limit, offset)
	removeFile(newFileName)

	f, err := os.Create(newFileName)
	defer f.Close()

	if err != nil {
		fmt.Println("Error creating csv")
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	if err := w.Write([]string{"id", "name"}); err != nil {
		fmt.Println("Error addint titles to csv")
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

	return "File successfully created", nil
}
