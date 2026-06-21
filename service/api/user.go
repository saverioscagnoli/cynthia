package api

import (
	"cynthia/constants"
	"cynthia/ds"
	"encoding/json"
	"io"
	"net/http"
)

func (rt *_router) GetLoggedUser(w http.ResponseWriter, r *http.Request, ctx RequestContext) {
	id := r.Header.Get("X-Discord-ID")
	user, err := rt.db.GetUser(id, r.Context())

	if user == nil && err == nil {
		ctx.Logger.Error("User not found in database", "id", id)
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	if err != nil {
		ctx.Logger.Error("Error retrieving user from database", "err", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		ctx.Logger.Error("Json decoding error in user retrieval", "err", err)
		http.Error(w, "json encoding error", http.StatusInternalServerError)
	}
}

func (rt *_router) UpdateUsername(w http.ResponseWriter, r *http.Request, ctx RequestContext) {
	id := r.Header.Get("X-Discord-ID")

	var body struct{ Username string }

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		ctx.Logger.Error("Failed to update username", "err", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = rt.db.UpdateUsername(id, body.Username, r.Context())

	if err != nil {
		ctx.Logger.Error("Failed to update username", "err", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rt *_router) UpdateTrainerSprite(w http.ResponseWriter, r *http.Request, ctx RequestContext) {
	id := r.Header.Get("X-Discord-ID")

	var body struct {
		ID int `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ctx.Logger.Error("Failed to update trainer sprite id", "err", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err := rt.db.UpdateTrainerSprite(id, &body.ID, r.Context()); err != nil {
		ctx.Logger.Error("Failed to update trainer sprite id", "err", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rt *_router) GetUserBanner(w http.ResponseWriter, r *http.Request, ctx RequestContext) {
	id := r.Header.Get("X-Discord-ID")

	banner, err := rt.db.GetUserBanner(id, r.Context())

	if err != nil {
		ctx.Logger.Error("Failed to get user banner", "err", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if banner == nil {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Write(*banner)
}

func (rt *_router) UpdateUserBanner(w http.ResponseWriter, r *http.Request, ctx RequestContext) {
	id := r.Header.Get("X-Discord-ID")

	r.Body = http.MaxBytesReader(w, r.Body, constants.MaxUserBannerSize)

	if err := r.ParseMultipartForm(constants.MaxUserBannerSize); err != nil {
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

	if err := rt.db.UpdateUserBanner(ds.Snowflake(id), &data, &contentType, r.Context()); err != nil {
		ctx.Logger.Error("Failed to update banner", "err", err)
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) DeleteUserBanner(w http.ResponseWriter, r *http.Request, ctx RequestContext) {
	id := r.Header.Get("X-Discord-ID")

	if err := rt.db.UpdateUserBanner(ds.Snowflake(id), nil, nil, r.Context()); err != nil {
		ctx.Logger.Error("Failed to delete banner", "err", err)
		http.Error(w, "db error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
