package pgedge

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCredentials(t *testing.T) {
	c := &Credentials{
		Username: "app",
		Password: "abcd1234",
		Database: "mydb",
		SSLMode:  "require",
	}
	require.Equal(t, "localhost", c.GetHost())
	require.Equal(t, 5432, c.GetPort())
	require.Equal(t, "postgres://app:abcd1234@localhost:5432/mydb?sslmode=require", c.URL())
	require.Equal(t, "host=localhost user=app password=abcd1234 dbname=mydb port=5432 sslmode=require", c.DSN())
	require.Equal(t, []string{
		"--host=localhost",
		"--port=5432",
		"--dbname=mydb",
		"--username=app",
	}, c.Options())
	require.Equal(t, []string{
		"PGHOST=localhost",
		"PGPORT=5432",
		"PGDATABASE=mydb",
		"PGUSER=app",
		"PGPASSWORD=abcd1234",
		"PGSSLMODE=require",
	}, c.Environment())
	require.Equal(t, "PGHOST=localhost PGPORT=5432 PGDATABASE=mydb PGUSER=app PGPASSWORD=abcd1234 PGSSLMODE=require", c.EnvironmentString())
}
