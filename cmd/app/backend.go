package main

import (
	"log/slog"
	"net/http"
)

func (a *App) SetupMuxHandler() {
	a.mux.HandleFunc("GET /healthcheck", func(w http.ResponseWriter, r *http.Request) {
		if err := a.DB.Ping(a.DB.Context); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			slog.Error("Healthcheck", "err", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
}
