package routes

import (
	"cynthia/cmd/pkapi/store"
	"net/http"
	"strconv"
)

// GET /sprites/trainer/{id}
func GetTrainerSprite(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	sprite, ok := store.TrainerSprites[id]

	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(*sprite)
}

// GET /sprites/trainer/count
func GetTrainerSpriteCount(w http.ResponseWriter, r *http.Request) {
	count := len(store.TrainerSprites)

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.Itoa(count)))
}
