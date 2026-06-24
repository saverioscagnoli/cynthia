package database

import (
	"context"
	"cynthia/ds"
	"cynthia/service/database/models"
)

func (db *dbimpl) GetWinStats(userID ds.Snowflake, ctx context.Context) (*models.WinStats, error) {
	const query = `
		SELECT
			COUNT(*) FILTER (WHERE winner_id = $1)                                                AS wins,
			COUNT(*) FILTER (WHERE (player1_id = $1 OR player2_id = $1)
				AND winner_id IS NOT NULL AND winner_id != $1
				AND status = 'finished')                                                          AS losses,
			COUNT(*) FILTER (WHERE (player1_id = $1 OR player2_id = $1)
				AND winner_id IS NULL
				AND status = 'finished')                                                          AS draws,
			COUNT(*) FILTER (WHERE winner_id = $1 AND type = 'single')                           AS single_wins,
			COUNT(*) FILTER (WHERE winner_id = $1 AND type = 'double')                           AS double_wins
		FROM matches
		WHERE (player1_id = $1 OR player2_id = $1)
		  AND status = 'finished'
	`

	var stats models.WinStats

	err := db.Pool.QueryRow(ctx, query, userID).Scan(
		&stats.Wins,
		&stats.Losses,
		&stats.Draws,
		&stats.SingleWins,
		&stats.DoubleWins,
	)

	if err != nil {
		return nil, err
	}

	total := stats.Wins + stats.Losses + stats.Draws

	if total > 0 {
		stats.Winrate = float32(stats.Wins) / float32(total)
	}

	// streak: walk matches ordered by date, newest first
	streakQuery := `
		SELECT winner_id
		FROM matches
		WHERE (player1_id = $1 OR player2_id = $1)
		  AND status = 'finished'
		ORDER BY finished_at DESC
	`

	rows, err := db.Pool.Query(ctx, streakQuery, userID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	current := 0
	best := 0

	for rows.Next() {
		var winnerID *string

		if err := rows.Scan(&winnerID); err != nil {
			return nil, err
		}

		if winnerID != nil && *winnerID == string(userID) {
			current++
			if current > best {
				best = current
			}
		} else {
			if current > best {
				best = current
			}
			current = 0
		}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	stats.CurrentStreak = current
	stats.BestStreak = best

	return &stats, nil
}
