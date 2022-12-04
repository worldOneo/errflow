package errflow

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type VirtualTxx struct {
	tx   *sqlx.Tx
	errs errChain
}

func emptyTransactionOf(err error) *VirtualTxx {
	return &VirtualTxx{errs: errChainOfErr(err)}
}

func BeginTxx(ctx context.Context, opts *sql.TxOptions, db *sqlx.DB) *VirtualTxx {
	tx, err := db.BeginTxx(ctx, opts)
	return &VirtualTxx{tx: tx, errs: errChainOfErr(err)}
}

func Beginx(db *sqlx.DB) *VirtualTxx {
	tx, err := db.Beginx()
	return &VirtualTxx{tx: tx, errs: errChainOfErr(err)}
}

func (tx *VirtualTxx) Commit() *VirtualTxx {
	return pass(func() error { return tx.tx.Commit() }, tx)
}

func (tx *VirtualTxx) Get(dest interface{}, query string, args ...interface{}) *VirtualTxx {
	return tx.GetContext(context.Background(), dest, query, args...)
}

func (tx *VirtualTxx) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) *VirtualTxx {
	return pass(func() error { return tx.tx.GetContext(ctx, dest, query, args...) }, tx)
}

func (tx *VirtualTxx) ExecFlow(query string, args ...any) Splitted[*VirtualResult, *VirtualTxx] {
	return tx.ExecContextFlow(context.Background(), query, args...)
}

func (tx *VirtualTxx) ExecContextFlow(ctx context.Context, query string, args ...any) Splitted[*VirtualResult, *VirtualTxx] {
	return SplitOf(tx.ExecContext(ctx, query, args...).(*VirtualResult), tx)
}

func (tx *VirtualTxx) Exec(query string, args ...any) sql.Result {
	return tx.ExecContext(context.Background(), query, args...)
}

func (tx *VirtualTxx) ExecContext(ctx context.Context, query string, args ...any) sql.Result {
	return Do(func() (sql.Result, error) {
		result, err := tx.tx.ExecContext(ctx, query, args...)
		return result, err
	}, tx, emptyResultOf)
}

func (tx *VirtualTxx) Prepare(query string) *VirtualStmt {
	return tx.PrepareContext(context.TODO(), query)
}

func (tx *VirtualTxx) PrepareContext(ctx context.Context, query string) *VirtualStmt {
	return Do(func() (*VirtualStmt, error) {
		stmt, err := tx.tx.PrepareContext(ctx, query)
		return stmtOf(stmt), err
	}, tx, emptyStmtOf)
}

func (tx *VirtualTxx) Rollback() error {
	return tx.Rollback()
}

func (tx *VirtualTxx) Stmt(stmt *sql.Stmt) *VirtualStmt {
	return tx.StmtContext(context.Background(), stmt)
}

func (tx *VirtualTxx) StmtContext(ctx context.Context, stmt *sql.Stmt) *VirtualStmt {
	return stmtOf(tx.tx.StmtContext(ctx, stmt))
}

func (tx *VirtualTxx) Err() error {
	return tx.errs.Err()
}

func (tx *VirtualTxx) Fail(err error) {
	tx.errs.Fail(err)
}

func (tx *VirtualTxx) Link() *error {
	return tx.errs.Link()
}
