package pkapi

import (
	"cynthia/store"
	"encoding/json"
	"net/http"
	"strconv"
)

// GET /type/{id}
func GetType(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	t, ok := store.Types[id]

	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(t)

	if err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
		return
	}
}

// GET /sprites/type/{id}
func GetTypeSprite(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	sprite, ok := store.TypeSprites[id]

	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if sprite == nil {
		http.Error(w, "sprite not available", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(*sprite)
}
