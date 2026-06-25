package api

import (
	"camilla/service/util"
	"encoding/json"
	"net/http"
)

func (rt *_router) GetWinStats(w http.ResponseWriter, r *http.Request, ctx RequestContext) {
	id := r.PathValue("id")

	if id == "" {
		ctx.Error(w, "user not found", http.StatusNotFound, nil)
		return
	}

	stats, err := rt.db.GetWinStats(id, r.Context())

	if err != nil {
		ctx.Error(w, "database error", http.StatusInternalServerError, util.Ptr(err.Error()))
		return
	}

	if stats == nil {
		ctx.Error(w, "not found", http.StatusNotFound, nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(stats)

	if err != nil {
		ctx.Error(w, "json encoding error", http.StatusInternalServerError, util.Ptr(err.Error()))
		return
	}
}
