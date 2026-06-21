package api

import "net/http"

func (rt *_router) Healthcheck(w http.ResponseWriter, r *http.Request, ctx RequestContext) {
	if err := rt.db.Ping(r.Context()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
