package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/lmittmann/tint"
)

func main() {
	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		AddSource:  true,
		Level:      slog.LevelInfo,
		TimeFormat: time.Kitchen,
	})))

	mux := http.NewServeMux()

	mux.HandleFunc("GET /pokemon", func(w http.ResponseWriter, req *http.Request) {
		json.NewEncoder(w).Encode("Fah!")
	})

	slog.Info("Listening on port 9000")
	http.ListenAndServe(":9000", mux)
}
