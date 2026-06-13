package routes

import (
	"cynthia/cmd/pkapi/store"
	"encoding/json"
	"net/http"
	"strconv"
)

// GET /item/{id}
func GetItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	item, ok := store.Items[id]

	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(item)

	if err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
	}
}

// GET /sprites/item/{id}
func GetItemSprite(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	sprite, ok := store.ItemSprites[id]

	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if sprite == nil {
		http.Error(w, "not available", http.StatusNotFound)
		return
	}

	w.Header().Set("ContentType", "image/png")
	w.Write(*sprite)
}
