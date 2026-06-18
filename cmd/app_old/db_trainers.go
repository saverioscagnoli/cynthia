package main

import (
	"cynthia/cmd/app/models"
	"cynthia/ds"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (db *DB) GetUser(discordID ds.Snowflake) (*models.User, error) {
	const query = `
		SELECT
			discord_id,
			username,
			avatar,
			trainer_id,
			created_at
		FROM users
		WHERE discord_id = $1
	`

	var user models.User

	err := db.Pool.QueryRow(db.Context, query, discordID).Scan(
		&user.DiscordID,
		&user.Username,
		&user.Avatar,
		&user.TrainerID,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (db *DB) GetTrainer(discordID ds.Snowflake) (*models.Trainer, error) {
	const query = `
			SELECT
				id,
				money,
				created_at,
				trainer_sprite_id
			FROM trainers
			WHERE id = $1
		`

	var trainer models.Trainer

	err := db.Pool.QueryRow(db.Context, query, discordID).Scan(
		&trainer.ID,
		&trainer.Money,
		&trainer.CreatedAt,
		&trainer.SpriteID,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &trainer, nil
}
