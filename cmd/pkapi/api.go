//go:build ignore

package main

import (
	"cynthia/cmd/pkapi/routes"
	"cynthia/cmd/pkapi/store"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"
	_ "github.com/mattn/go-sqlite3"
)

type responseWriter struct {
	http.ResponseWriter
	bytesWritten int
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	n, err := rw.ResponseWriter.Write(b)
	rw.bytesWritten += n
	return n, err
}

func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{ResponseWriter: w}
		next.ServeHTTP(rw, r)
		slog.Info("request",
			"method", r.Method,
			"path", r.URL.Path,
			"remote", r.RemoteAddr,
			"duration", time.Since(start),
			"size_kb", float64(rw.bytesWritten)/1024,
		)
	})
}

func main() {
	godotenv.Load()

	path := flag.String("path", "./pokemon.db", "Database path")
	flag.Parse()

	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      slog.LevelInfo,
		TimeFormat: time.Kitchen,
	})))

	slog.Info("Extracting data in-memory...")
	store.Extract(*path)
	slog.Info("Extracted data successfully.")

	mux := http.NewServeMux()

	mux.HandleFunc("/pokemon/{id}", routes.GetPokemon)
	mux.HandleFunc("/item/{id}", routes.GetItem)
	mux.HandleFunc("/type/{id}", routes.GetType)

	mux.HandleFunc("/sprites/pokemon/{id}", routes.GetPokemonSprites)
	mux.HandleFunc("/sprites/pokemon/{id}/{type}", routes.GetPokemonSprite)
	mux.HandleFunc("/sprites/item/{id}", routes.GetItemSprite)
	mux.HandleFunc("/sprites/type/{id}", routes.GetTypeSprite)

	port := os.Getenv("PKAPI_PORT")
	slog.Info("Listening on port " + port)

	http.ListenAndServe("0.0.0.0:"+port, logRequests(mux))
}
