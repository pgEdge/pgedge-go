package pgedge

import (
	"fmt"
	"strings"
)

// Credentials used to connect to a database
type Credentials struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	SSLMode  string `json:"sslmode"`
}

// GetPort returns the port, with a default of 5432 if unset
func (c *Credentials) GetPort() int {
	if c.Port == 0 {
		return 5432
	}
	return c.Port
}

// GetHost returns the host, with a default of "localhost" if unset
func (c *Credentials) GetHost() string {
	if c.Host == "" {
		return "localhost"
	}
	return c.Host
}

// GetSSLMode returns the SSL mode, with a default of "verify-full" if unset
func (c *Credentials) GetSSLMode() string {
	if c.SSLMode == "" {
		return "verify-full"
	}
	return c.SSLMode
}

// URL for the database
func (c *Credentials) URL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.Username, c.Password, c.GetHost(), c.GetPort(), c.Database, c.GetSSLMode())
}

// DSN returns the Data Source Name for the database
func (c *Credentials) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		c.GetHost(), c.Username, c.Password, c.Database, c.GetPort(), c.GetSSLMode())
}

// Options returns command line options for the database that may be provided
// to a command like psql, pg_dump, or pg_restore. The password and sslmode are
// not included since they can't be provided as options.
func (c *Credentials) Options() []string {
	return []string{
		fmt.Sprintf("--host=%s", c.GetHost()),
		fmt.Sprintf("--port=%d", c.GetPort()),
		fmt.Sprintf("--dbname=%s", c.Database),
		fmt.Sprintf("--username=%s", c.Username),
	}
}

// Environment returns libpq compatible environment variables for the database.
func (c *Credentials) Environment() []string {
	// https://www.postgresql.org/docs/current/libpq-envars.html
	return []string{
		fmt.Sprintf("PGHOST=%s", c.GetHost()),
		fmt.Sprintf("PGPORT=%d", c.GetPort()),
		fmt.Sprintf("PGDATABASE=%s", c.Database),
		fmt.Sprintf("PGUSER=%s", c.Username),
		fmt.Sprintf("PGPASSWORD=%s", c.Password),
		fmt.Sprintf("PGSSLMODE=%s", c.GetSSLMode()),
	}
}

// EnvironmentString returns libpq compatible environment variables for the
// database as a single string.
func (c *Credentials) EnvironmentString() string {
	return strings.Join(c.Environment(), " ")
}
