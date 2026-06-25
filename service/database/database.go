package database

import (
	"camilla/ds"
	"camilla/service/database/models"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AppDatabase interface {
	Ping(ctx context.Context) error
	GetUser(id ds.Snowflake, ctx context.Context) (*models.User, error)
	GetOrInsertUser(user *ds.User, ctx context.Context) (*models.User, error)
	GetBagItems(userID ds.Snowflake, ctx context.Context) ([]models.BagItem, error)
	GetUserBanner(userID ds.Snowflake, ctx context.Context) (*[]byte, error)

	GetWinStats(userID ds.Snowflake, ctx context.Context) (*models.WinStats, error)

	UpsertUser(user *ds.User, ctx context.Context) error
	UpdateUsername(userID ds.Snowflake, username string, ctx context.Context) error
	UpdateTrainerSprite(userID ds.Snowflake, spriteID *int, ctx context.Context) error
	UpdateUserBanner(userID ds.Snowflake, banner *[]byte, contentType *string, ctx context.Context) error
}

type dbimpl struct {
	*pgxpool.Pool
}

func New(pool *pgxpool.Pool) (AppDatabase, error) {
	if err := createSchema(context.Background(), pool); err != nil {
		return nil, err
	}

	return &dbimpl{Pool: pool}, nil
}

func (db *dbimpl) Ping(ctx context.Context) error {
	return db.Pool.Ping(ctx)
}
