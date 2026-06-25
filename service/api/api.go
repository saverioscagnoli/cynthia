package api

import (
	"camilla/ds"
	"camilla/service/database"
	"errors"
	"log/slog"
	"net/http"
	"os"
)

type Config struct {
	Logger             *slog.Logger
	Database           database.AppDatabase
	SnowflakeGenerator *ds.SnowflakeGenerator
}

type Router interface {
	Handler() http.Handler
	Close() error
}

func New(conf Config) (Router, error) {
	if conf.Logger == nil {
		return nil, errors.New("A logger is required")
	}

	if conf.Database == nil {
		return nil, errors.New("A database is required")
	}

	if conf.SnowflakeGenerator == nil {
		return nil, errors.New("A snowflake generator is required")
	}

	mux := http.NewServeMux()

	frontendURL := os.Getenv("FRONTEND_URL")
	jwtSecret := os.Getenv("JWT_SECRET")
	oauth2ClientID := os.Getenv("OAUTH2_CLIENT_ID")
	oauth2Secret := os.Getenv("OAUTH2_SECRET")
	oauth2RedirectURI := os.Getenv("OAUTH2_REDIRECT_URI")

	return &_router{
		mux:               mux,
		logger:            conf.Logger,
		db:                conf.Database,
		snowflake:         conf.SnowflakeGenerator,
		frontendURL:       frontendURL,
		jwtSecret:         jwtSecret,
		oauth2ClientID:    oauth2ClientID,
		oauth2Secret:      oauth2Secret,
		oauth2RedirectURI: oauth2RedirectURI,
		isProd:            os.Getenv("APP_ENV") == "prod",
	}, nil

}

type _router struct {
	mux *http.ServeMux
	// This is the logger for everything that is not
	// involved with requests
	logger            *slog.Logger
	db                database.AppDatabase
	snowflake         *ds.SnowflakeGenerator
	frontendURL       string
	jwtSecret         string
	oauth2ClientID    string
	oauth2Secret      string
	oauth2RedirectURI string
	isProd            bool
}
