package utility

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DatabaseTxStub struct{}

func (tx *DatabaseTxStub) Begin(ctx context.Context) (pgx.Tx, error) {
	return nil, nil
}
func (tx *DatabaseTxStub) Commit(ctx context.Context) error {
	return nil
}
func (tx *DatabaseTxStub) Rollback(ctx context.Context) error {
	return nil
}
func (tx *DatabaseTxStub) Conn() *pgx.Conn {
	return nil
}
func (tx *DatabaseTxStub) CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, reader pgx.CopyFromSource) (int64, error) {
	return 0, nil
}
func (tx *DatabaseTxStub) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return pgconn.CommandTag{}, nil
}
func (tx *DatabaseTxStub) LargeObjects() pgx.LargeObjects {
	return pgx.LargeObjects{}
}
func (tx *DatabaseTxStub) Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error) {
	return nil, nil
}
func (tx *DatabaseTxStub) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return nil, nil
}
func (tx *DatabaseTxStub) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return nil
}
func (tx *DatabaseTxStub) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return nil
}

type DatabasePoolStub struct{}

func (pool *DatabasePoolStub) Begin(ctx context.Context) (pgx.Tx, error) {
	return pgx.Tx(&DatabaseTxStub{}), nil
}
func (pool *DatabasePoolStub) Close() {}
func (pool *DatabasePoolStub) Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error) {
	return pgconn.CommandTag{}, nil
}
func (pool *DatabasePoolStub) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	return nil, nil
}
func (pool *DatabasePoolStub) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return nil
}
func (pool *DatabasePoolStub) Ping(ctx context.Context) error {
	return nil
}
func (pool *DatabasePoolStub) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	return nil
}
