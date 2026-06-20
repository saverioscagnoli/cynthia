package main

import "context"

func (db *db) Ping(ctx context.Context) error {
	return db.pool.Ping(ctx)
}
