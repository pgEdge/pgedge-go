package pgedge

import (
	"context"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

func RetryConnect(ctx context.Context, url string) (*pgx.Conn, error) {
	for {
		conn, err := pgx.Connect(ctx, url)
		if err == nil {
			return conn, nil
		}
		if strings.Contains(err.Error(), "connection refused") ||
			strings.Contains(err.Error(), "no such host") ||
			strings.Contains(err.Error(), "password authentication failed") {
			time.Sleep(time.Second * 3)
			continue
		}
		return nil, err
	}
}
