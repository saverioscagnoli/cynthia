package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func (b *backend) Healthcheck(w http.ResponseWriter, r *http.Request) {
	if err := b.db.Ping(r.Context()); err != nil {
		http.Error(w, fmt.Sprintf("healthcheck: %s", err.Error()), http.StatusInternalServerError)
		slog.Error("Healthcheck", "err", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
