package schema

import (
	"context"
	"strings"

	"github.com/jackc/pgx/v5"
)

type RepsetAddAllTablesInput struct {
	RepsetName string
	Schemas    []string
}

type RepsetAddTableInput struct {
	RepsetName string
	TableName  string
	SyncData   bool
}

type Spock interface {
	RepsetAddAllTables(context.Context, RepsetAddAllTablesInput) error
	RepsetAddTable(context.Context, RepsetAddTableInput) error
}

type SpockImpl struct {
	conn *pgx.Conn
}

func NewSpock(conn *pgx.Conn) *SpockImpl {
	return &SpockImpl{conn: conn}
}

func (s *SpockImpl) RepsetAddAllTables(ctx context.Context, input RepsetAddAllTablesInput) error {
	_, err := s.conn.Exec(ctx, `SELECT spock.repset_add_all_tables($1, $2);`, input.RepsetName, input.Schemas)
	return err
}

func (s *SpockImpl) RepsetAddTable(ctx context.Context, input RepsetAddTableInput) error {
	_, err := s.conn.Exec(ctx, `SELECT spock.repset_add_table($1, $2, $3, NULL, NULL);`,
		input.RepsetName, input.TableName, input.SyncData)
	if err != nil {
		if IsUniqueConstraintError(err) {
			return nil // table is already in the repset
		}
		return err
	}
	return nil
}

func IsUniqueConstraintError(err error) bool {
	return strings.Contains(err.Error(), "duplicate key value violates unique constraint")
}
