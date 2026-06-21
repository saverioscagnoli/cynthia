package main

import (
	"cynthia/ds"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

func (b *backend) GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("X-Discord-ID")
	user, err, found := b.db.GetUser(id, r.Context())

	if !found {
		slog.Error("User not found in database", "id", id)
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if err != nil {
		slog.Error("Error retrieving user from database", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		slog.Error("Json decoding error in user retrieval", "err", err)
		http.Error(w, "json encoding error", http.StatusInternalServerError)
	}
}

func (b *backend) UpdateUsername(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("X-Discord-ID")

	var body struct{ Username string }

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		slog.Error("Failed to update username", "err", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = b.db.UpdateUsername(id, body.Username, r.Context())

	if err != nil {
		slog.Error("Failed to update username", "err", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (b *backend) UpdateTrainerSpriteID(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("X-Discord-ID")

	var body struct {
		ID int `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		slog.Error("Failed to update trainer sprite id", "err", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err := b.db.UpdateSpriteID(id, body.ID, r.Context()); err != nil {
		slog.Error("Failed to update trainer sprite id", "err", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

const maxBannerSize = 5 * 1024 * 1024

func (b *backend) UpdateBanner(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("X-Discord-ID")

	r.Body = http.MaxBytesReader(w, r.Body, maxBannerSize)

	if err := r.ParseMultipartForm(maxBannerSize); err != nil {
		http.Error(w, "banner must be under 5MB", http.StatusRequestEntityTooLarge)
		return
	}

	file, _, err := r.FormFile("banner")

	if err != nil {
		http.Error(w, "missing banner", http.StatusBadRequest)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		http.Error(w, "failed to read file", http.StatusInternalServerError)
		return
	}

	contentType := http.DetectContentType(data)

	if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/gif" {
		http.Error(w, "banner must be an image", http.StatusBadRequest)
		return
	}

	if err := b.db.UpdateBanner(ds.Snowflake(id), data, &contentType, r.Context()); err != nil {
		slog.Error("Failed to update banner", "err", err)
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (b *backend) DeleteBanner(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("X-Discord-ID")

	if err := b.db.UpdateBanner(ds.Snowflake(id), nil, nil, r.Context()); err != nil {
		slog.Error("Failed to delete banner", "err", err)
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
