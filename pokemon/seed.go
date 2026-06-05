//go:build ignore

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

type PokemonList struct {
	Count   int    `json:"count"`
	Next    string `json:"next"`
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Sprites struct {
	FrontDefault     string `json:"front_default"`
	FrontShiny       string `json:"front_shiny"`
	FrontFemale      string `json:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female"`
	BackDefault      string `json:"back_default"`
	BackShiny        string `json:"back_shiny"`
	BackFemale       string `json:"back_female"`
	BackShinyFemale  string `json:"back_shiny_female"`
}

type Pokemon struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Sprites Sprites `json:"sprites"`
}

func downloadSprite(url, path string) {
	if url == "" {
		return
	}
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		fmt.Printf("  failed to download %s\n", path)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		fmt.Printf("  failed to create %s\n", path)
		return
	}
	defer file.Close()
	io.Copy(file, resp.Body)
}

func main() {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	spritesDir := filepath.Join(dir, "db", "sprites")

	os.MkdirAll(spritesDir, 0755)

	url := "http://localhost:8000/api/v2/pokemon?limit=100&offset=0"

	for url != "" {
		resp, _ := http.Get(url)
		var list PokemonList
		json.NewDecoder(resp.Body).Decode(&list)
		resp.Body.Close()

		for _, p := range list.Results {
			r, _ := http.Get(p.URL)
			var pokemon Pokemon
			json.NewDecoder(r.Body).Decode(&pokemon)
			r.Body.Close()

			fmt.Printf("downloading %s (#%d)...\n", pokemon.Name, pokemon.ID)

			pokemonDir := filepath.Join(spritesDir, fmt.Sprintf("%d", pokemon.ID))
			os.MkdirAll(pokemonDir, 0755)

			s := pokemon.Sprites
			downloadSprite(s.FrontDefault, filepath.Join(pokemonDir, "front.png"))
			downloadSprite(s.FrontShiny, filepath.Join(pokemonDir, "front-shiny.png"))
			downloadSprite(s.FrontFemale, filepath.Join(pokemonDir, "front-female.png"))
			downloadSprite(s.FrontShinyFemale, filepath.Join(pokemonDir, "front-shiny-female.png"))
			downloadSprite(s.BackDefault, filepath.Join(pokemonDir, "back.png"))
			downloadSprite(s.BackShiny, filepath.Join(pokemonDir, "back-shiny.png"))
			downloadSprite(s.BackFemale, filepath.Join(pokemonDir, "back-female.png"))
			downloadSprite(s.BackShinyFemale, filepath.Join(pokemonDir, "back-shiny-female.png"))
		}

		url = list.Next
	}

	fmt.Println("done!")
}
