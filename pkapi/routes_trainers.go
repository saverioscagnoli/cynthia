package pkapi

import (
	"cynthia/assets"
	"cynthia/store"
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

// GET /sprites/trainer/sheet/map
func GetTrainerSheet(w http.ResponseWriter, r *http.Request) {
	data, err := assets.TrainerSheet.ReadFile("trainer-sheet.png")

	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(data)
}

// GET /sprites/trainer/sheet
func GetTrainerSheetMap(w http.ResponseWriter, r *http.Request) {
	data, err := assets.TrainerSheetJSON.ReadFile("trainer-sheet.json")
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// GET /sprites/trainer/count
func GetTrainerSpriteCount(w http.ResponseWriter, r *http.Request) {
	count := len(store.TrainerSprites)

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(strconv.Itoa(count)))
}
