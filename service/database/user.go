package database

import (
	"context"
	"cynthia/cmd/app/models"
	"cynthia/ds"
	"cynthia/store"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (db *dbimpl) GetUser(id ds.Snowflake, ctx context.Context) (*models.User, error) {
	const query = `
        SELECT
            id,
            username,
            discord_username,
            avatar_hash,
            money,
            sprite_id,
            banner,
            created_at
        FROM users
        WHERE id = $1
    `
	var user models.User

	err := db.Pool.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.DiscordUsername,
		&user.AvatarHash,
		&user.Money,
		&user.SpriteID,
		&user.Banner,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	bag, err := db.GetBagItems(user.ID, ctx)

	if err != nil {
		return nil, err
	}

	user.Bag = bag

	return &user, nil
}

func (db *dbimpl) GetOrInsertUser(user *ds.User, ctx context.Context) (*models.User, error) {
	_, err := db.Pool.Exec(ctx,
		`INSERT INTO users (id, username, discord_username, avatar_hash)
        VALUES ($1, $2, $2, $3)
        ON CONFLICT (id) DO NOTHING`,
		user.ID, user.Username, user.Avatar,
	)

	if err != nil {
		return nil, err
	}

	return db.GetUser(user.ID, ctx)
}

func (db *dbimpl) GetBagItems(userID ds.Snowflake, ctx context.Context) ([]models.BagItem, error) {
	rows, err := db.Pool.Query(ctx, `
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
			return nil, nil
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

func (db *dbimpl) GetUserBanner(userID ds.Snowflake, ctx context.Context) (*[]byte, error) {
	var banner []byte

	err := db.Pool.QueryRow(ctx, `
        SELECT banner FROM users WHERE id = $1
    `, userID).Scan(&banner)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &banner, nil
}

func (db *dbimpl) UpsertUser(user *ds.User, ctx context.Context) error {
	_, err := db.Pool.Exec(ctx,
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

func (db *dbimpl) UpdateUsername(userID ds.Snowflake, username string, ctx context.Context) error {
	_, err := db.Pool.Exec(ctx,
		`UPDATE users SET username = $1 WHERE id = $2`,
		username, userID,
	)

	return err
}

func (db *dbimpl) UpdateTrainerSprite(userID ds.Snowflake, spriteID *int, ctx context.Context) error {
	_, err := db.Pool.Exec(ctx,
		`UPDATE users SET sprite_id = $1 WHERE id = $2`,
		spriteID, userID,
	)

	return err
}

func (db *dbimpl) UpdateUserBanner(userID ds.Snowflake, banner *[]byte, contentType *string, ctx context.Context) error {
	_, err := db.Pool.Exec(ctx,
		`UPDATE users SET banner = $1, banner_type = $2 WHERE id = $3`,
		banner, contentType, userID,
	)

	return err
}
