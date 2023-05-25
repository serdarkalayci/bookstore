package redis

import (
	"context"
)

type databaseCollection interface {
	GetIntValue(ctx context.Context, id string) (int, error)
}

type dbHelper interface {
	GetIntValue(ctx context.Context, id string) (int, error)
	Ping(ctx context.Context) (bool)
}