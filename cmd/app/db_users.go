package main

import (
	"context"
	"cynthia/cmd/app/models"
	"cynthia/ds"
	"cynthia/store"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
)

func (db *db) UpsertUser(user *ds.User, ctx context.Context) error {
	_, err := db.pool.Exec(ctx,
		`INSERT INTO users (id, username, discord_username, avatar_hash)
        VALUES ($1, $2, $2, $3)
        ON CONFLICT (id) DO UPDATE
        SET username = EXCLUDED.username,
            discord_username = EXCLUDED.discord_username,
            avatar_hash = EXCLUDED.avatar_hash`,
		user.ID, user.Username, user.Avatar,
	)
	return err
}

func (db *db) GetBagItems(userID ds.Snowflake, ctx context.Context) ([]models.BagItem, error) {
	rows, err := db.pool.Query(ctx, `
        SELECT item_id, quantity FROM bag WHERE user_id = $1
    `, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	items := []models.BagItem{}

	for rows.Next() {
		var itemID, quantity int

		if err := rows.Scan(&itemID, &quantity); err != nil {
			return nil, err
		}

		item, ok := store.Items[itemID]

		if !ok {
			continue // or return error
		}

		items = append(items, models.BagItem{
			ItemID:      itemID,
			Name:        item.Name,
			Quantity:    quantity,
			Cost:        item.Cost,
			FlingPower:  item.FlingPower,
			FlingEffect: item.FlingEffect,
		})
	}

	return items, rows.Err()
}

func (db *db) GetUser(id ds.Snowflake, ctx context.Context) (*models.User, error, bool) {
	const query = `
        SELECT
            id,
            username,
            discord_username,
            avatar_hash,
            money,
            sprite_id,
            created_at
        FROM users
        WHERE id = $1
    `
	var user models.User
	err := db.pool.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.DiscordUsername,
		&user.AvatarHash,
		&user.Money,
		&user.SpriteID,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil, false
		}

		return nil, err, true
	}

	bag, err := db.GetBagItems(user.ID, ctx)

	if err != nil {
		return nil, err, true
	}

	user.Bag = bag

	return &user, nil, true
}

func (db *db) UpdateUsername(userID ds.Snowflake, username string, ctx context.Context) error {
	if len(username) < 3 || len(username) > 23 {
		return errors.New("username must be between 3 and 23 characters")
	}

	_, err := db.pool.Exec(ctx,
		`UPDATE users SET username = $1 WHERE id = $2`,
		username, userID,
	)

	return err
}

func (db *db) UpdateSpriteID(userID ds.Snowflake, spriteID int, ctx context.Context) error {
	_, err := db.pool.Exec(ctx,
		`UPDATE users SET sprite_id = $1 WHERE id = $2`,
		spriteID, userID,
	)

	return err
}

func (b *backend) GetBanner(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("X-Discord-ID")

	var banner []byte
	var bannerType string

	err := b.db.pool.QueryRow(r.Context(),
		`SELECT banner, banner_type FROM users WHERE id = $1`, id,
	).Scan(&banner, &bannerType)

	if err != nil || banner == nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", bannerType)
	w.Write(banner)
}

func (db *db) UpdateBanner(userID ds.Snowflake, banner []byte, contentType *string, ctx context.Context) error {
	_, err := db.pool.Exec(ctx,
		`UPDATE users SET banner = $1, banner_type = $2 WHERE id = $3`,
		banner, contentType, userID,
	)
	return err
}
