package api

import (
	"camilla/constants"
	"camilla/service/util"
	"encoding/json"
	"io"
	"net/http"
)

func (rt *_router) GetLoggedUser(w http.ResponseWriter, r *http.Request, ctx RequestContext) {
	claims := ctx.Claims(r)
	user, err := rt.db.GetUser(claims.UserID, r.Context())

	if user == nil && err == nil {
		ctx.Error(w, "user not found", http.StatusNotFound, nil)
		return
	}

	if err != nil {
		ctx.Logger.Error("Error retrieving user from database", "err", err)
		ctx.Error(w, "error retrieving user from database", http.StatusInternalServerError, util.Ptr(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		ctx.Error(w, "json encoding error", http.StatusInternalServerError, util.Ptr(err.Error()))
	}
}

func (rt *_router) GetUser(w http.ResponseWriter, r *http.Request, ctx RequestContext) {
	id := r.PathValue("id")

	user, err := rt.db.GetUser(id, r.Context())

	if err != nil {
		ctx.Error(w, "database error", http.StatusInternalServerError, util.Ptr(err.Error()))
		return
	}

	if user == nil {
		ctx.Error(w, "user not found", http.StatusNotFound, nil)
		return
	}

	if user.Banner != nil {
		user.Banner = &[]byte{byte(1)}
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)

	if err != nil {
		ctx.Error(w, "json encoding error", http.StatusInternalServerError, util.Ptr(err.Error()))
	}
}

func (rt *_router) UpdateUsername(w http.ResponseWriter, r *http.Request, ctx RequestContext) {
	claims := ctx.Claims(r)

	var body struct{ Username string }

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		ctx.Error(w, "bad request", http.StatusBadRequest, nil)
		return
	}

	err = rt.db.UpdateUsername(claims.UserID, body.Username, r.Context())

	if err != nil {
		ctx.Error(w, "internal server error", http.StatusInternalServerError, nil)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rt *_router) UpdateTrainerSprite(w http.ResponseWriter, r *http.Request, ctx RequestContext) {
	claims := ctx.Claims(r)

	var body struct {
		ID int `json:"id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		ctx.Error(w, "bad request", http.StatusBadRequest, nil)
		return
	}

	if err := rt.db.UpdateTrainerSprite(claims.UserID, &body.ID, r.Context()); err != nil {
		ctx.Error(w, "database error", http.StatusInternalServerError, util.Ptr(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rt *_router) GetUserBanner(w http.ResponseWriter, r *http.Request, ctx RequestContext) {
	id := r.PathValue("id")

	banner, err := rt.db.GetUserBanner(id, r.Context())

	if err != nil {
		ctx.Error(w, "database error", http.StatusInternalServerError, util.Ptr(err.Error()))
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
	claims := ctx.Claims(r)

	r.Body = http.MaxBytesReader(w, r.Body, constants.MaxUserBannerSize)

	if err := r.ParseMultipartForm(constants.MaxUserBannerSize); err != nil {
		ctx.Error(w, "banner is over 5MB", http.StatusRequestEntityTooLarge, nil)
		return
	}

	file, _, err := r.FormFile("banner")

	if err != nil {
		ctx.Error(w, "missing banner", http.StatusBadRequest, nil)
		return
	}

	defer file.Close()

	data, err := io.ReadAll(file)

	if err != nil {
		ctx.Error(w, "internal error", http.StatusInternalServerError, util.Ptr(err.Error()))
		return
	}

	contentType := http.DetectContentType(data)

	if contentType != "image/jpeg" && contentType != "image/png" && contentType != "image/gif" && contentType != "image/webp" {
		ctx.Error(w, "banner must be an image", http.StatusBadRequest, nil)
		return
	}

	if err := rt.db.UpdateUserBanner(claims.UserID, &data, &contentType, r.Context()); err != nil {
		ctx.Error(w, "database error", http.StatusInternalServerError, util.Ptr(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rt *_router) DeleteUserBanner(w http.ResponseWriter, r *http.Request, ctx RequestContext) {
	claims := ctx.Claims(r)

	if err := rt.db.UpdateUserBanner(claims.UserID, nil, nil, r.Context()); err != nil {
		ctx.Error(w, "database error", http.StatusInternalServerError, util.Ptr(err.Error()))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
