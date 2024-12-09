package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect(uri string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), uri)
	if err != nil {
		return nil, err
	}
	return conn, nil

}
