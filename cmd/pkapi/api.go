//go:build ignore

package main

import (
	"cynthia/cmd/pkapi/routes"
	"cynthia/cmd/pkapi/store"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
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

		slog.Info(fmt.Sprintf("/%s %s %s", r.Method, r.URL.Path, time.Since(start)))
	})
}

func main() {
	godotenv.Load()

	path := flag.String("path", "./pokemon.db", "Database path")
	flag.Parse()

	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      slog.LevelInfo,
		TimeFormat: time.TimeOnly,
	})))

	slog.Info("Extracting data in-memory...")
	store.Extract(*path)
	slog.Info("Extracted data successfully.")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /pokemon/{id}", routes.GetPokemon)
	mux.HandleFunc("GET /item/{id}", routes.GetItem)
	mux.HandleFunc("GET /type/{id}", routes.GetType)

	mux.HandleFunc("GET /sprites/pokemon/{id}", routes.GetPokemonSprites)
	mux.HandleFunc("GET /sprites/pokemon/{id}/{type}", routes.GetPokemonSprite)
	mux.HandleFunc("GET /sprites/item/{id}", routes.GetItemSprite)
	mux.HandleFunc("GET /sprites/type/{id}", routes.GetTypeSprite)

	mux.HandleFunc("GET /sprites/trainer/sheet", routes.GetTrainerSheet)
	mux.HandleFunc("GET /sprites/trainer/sheet/map", routes.GetTrainerSheetMap)
	mux.HandleFunc("GET /sprites/trainer/count", routes.GetTrainerSpriteCount)
	mux.HandleFunc("GET /sprites/trainer/{id}", routes.GetTrainerSprite)

	mux.HandleFunc("GET /stat/{id}", routes.GetStat)

	port := os.Getenv("PKAPI_PORT")
	slog.Info("Listening on port " + port)

	c := cors.Default()
	handler := c.Handler(logRequests(mux))

	http.ListenAndServe("0.0.0.0:"+port, handler)
}
