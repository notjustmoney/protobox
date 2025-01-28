package outboxsql

import (
	"context"
	"database/sql"

	"github.com/Masterminds/squirrel"
)

type dbCtx struct {
	db *sql.DB
}

func (d *dbCtx) executor(ctx context.Context) squirrel.BaseRunner {
	tx, ok := txFromContext(ctx)
	if !ok {
		return d.db
	}
	return tx
}

type txContextKey struct{}

func withTx(ctx context.Context, tx *sql.Tx) context.Context {
	return context.WithValue(ctx, txContextKey{}, tx)
}

func txFromContext(ctx context.Context) (*sql.Tx, bool) {
	tx, ok := ctx.Value(txContextKey{}).(*sql.Tx)
	return tx, ok
}
