package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/lmittmann/tint"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	path := flag.String("path", "./pokemon.db", "Database path")
	flag.Parse()

	slog.SetDefault(slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		Level:      slog.LevelInfo,
		TimeFormat: time.Kitchen,
	})))

	slog.Info("Extracting data in-memory...")

	Init(*path)

	slog.Info("Extracted all data successfully.", "pokemons", len(Pokemons))

	mux := http.NewServeMux()

	mux.HandleFunc("GET /pokemon/sprite/{id}/{type}", func(w http.ResponseWriter, req *http.Request) {
		idStr := req.PathValue("id")
		spriteType := req.PathValue("type")

		id, err := strconv.Atoi(idStr)

		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		p, ok := Pokemons[id]

		if !ok {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		sprites := map[string]*[]byte{
			"front":              p.Sprites.Front,
			"front-shiny":        p.Sprites.FrontShiny,
			"back":               p.Sprites.Back,
			"back-shiny":         p.Sprites.BackShiny,
			"front-female":       p.Sprites.FrontFemale,
			"front-shiny-female": p.Sprites.FrontShinyFemale,
			"back-female":        p.Sprites.BackFemale,
			"back-shiny-female":  p.Sprites.BackShinyFemale,
		}

		sprite, ok := sprites[spriteType]
		if !ok || sprite == nil {
			http.Error(w, "sprite not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "image/png")
		w.Write(*sprite)
	})

	slog.Info("Listening on port 9000")
	http.ListenAndServe(":9000", mux)
}
