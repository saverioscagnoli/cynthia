package api

import (
	"encoding/json"
	"net/http"
)

func (rt *_router) GetWinStats(w http.ResponseWriter, r *http.Request, ctx RequestContext) {
	id := r.PathValue("id")

	if id == "" {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	stats, err := rt.db.GetWinStats(id, r.Context())

	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	if stats == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(stats)

	if err != nil {
		http.Error(w, "encoding error", http.StatusInternalServerError)
		return
	}
}
