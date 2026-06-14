package routes

import (
	"cynthia/cmd/pkapi/store"
	"encoding/json"
	"net/http"
	"strconv"
)

// GET /stat/{id}
func GetStat(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	item, ok := store.Stats[id]

	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(item)

	if err != nil {
		http.Error(w, "failed to encode json", http.StatusInternalServerError)
	}
}
