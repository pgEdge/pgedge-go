package pgedge

import (
	"context"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
)

func ConnectWithRetries(ctx context.Context, url string) (*pgx.Conn, error) {
	for {
		conn, err := pgx.Connect(ctx, url)
		if err == nil {
			return conn, nil
		}
		errMsg := err.Error()
		if strings.Contains(errMsg, "connection refused") ||
			strings.Contains(errMsg, "no such host") ||
			strings.Contains(errMsg, "password authentication failed") {
			time.Sleep(time.Second * 3)
			continue
		}
		return nil, err
	}
}
