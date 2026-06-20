package pkapi

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

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

func Start(port string) {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /pokemon/{id}", GetPokemon)
	mux.HandleFunc("GET /item/{id}", GetItem)
	mux.HandleFunc("GET /type/{id}", GetType)

	mux.HandleFunc("GET /sprites/pokemon/{id}", GetPokemonSprites)
	mux.HandleFunc("GET /sprites/pokemon/{id}/{type}", GetPokemonSprite)
	mux.HandleFunc("GET /sprites/item/{id}", GetItemSprite)
	mux.HandleFunc("GET /sprites/type/{id}", GetTypeSprite)

	mux.HandleFunc("GET /sprites/trainer/sheet", GetTrainerSheet)
	mux.HandleFunc("GET /sprites/trainer/sheet/map", GetTrainerSheetMap)
	mux.HandleFunc("GET /sprites/trainer/count", GetTrainerSpriteCount)
	mux.HandleFunc("GET /sprites/trainer/{id}", GetTrainerSprite)

	mux.HandleFunc("GET /stat/{id}", GetStat)

	slog.Info("Listening on port " + port)

	c := cors.Default()
	handler := c.Handler(logRequests(mux))

	http.ListenAndServe(":"+port, handler)
}
