package pkapi

import (
	"cynthia/store"
	"encoding/json"
	"net/http"
	"strconv"
)

// GET /pokemon/{id}
func GetPokemon(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	p, ok := store.Pokemons[id]

	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(p)

	if err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}

// GET /sprites/pokemon/{id}
func GetPokemonSprites(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	sprites, ok := store.PokemonSpriteMap[id]

	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(sprites)

	if err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
	}
}

// GET /sprites/pokemon/{id}/{typ}
func GetPokemonSprite(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.PathValue("id"))
	spriteType, ok := store.ParseSprite(r.PathValue("type"))

	if !ok {
		http.Error(w, "invalid sprite type", http.StatusBadRequest)
		return
	}

	sprites, ok := store.PokemonSpriteMap[id]

	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	sprite := sprites.Get(spriteType)

	if sprite == nil {
		http.Error(w, "sprite not available", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(*sprite)
}
