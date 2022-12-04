package errflow

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DBx struct {
	db   *sqlx.DB
	errs errChain
}

func (db *DBx) ExecFlow(query string, args ...any) Splitted[*Result, *DBx] {
	return SplitOf(db.Exec(query, args...), db)
}

func (db *DBx) Exec(query string, args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := db.db.Exec(query, args...)
		return ResultOf(result, db), err
	}, db, EmptyResultOf)
}

func (db *DBx) ExecContextFlow(ctx context.Context, query string, args ...any) Splitted[*Result, *DBx] {
	return SplitOf(db.ExecContext(ctx, query, args...), db)
}

func (db *DBx) ExecContext(ctx context.Context, query string, args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := db.db.ExecContext(ctx, query, args...)
		return ResultOf(result, db), err
	}, db, EmptyResultOf)
}
func (db *DBx) Get(dest any, query string, args ...any) *DBx {
	return pass(func() error { return db.db.Get(dest, query, args...) }, db)
}

func (db *DBx) GetContext(ctx context.Context, dest any, query string, args ...any) *DBx {
	return pass(func() error { return db.db.GetContext(ctx, dest, query, args...) }, db)
}

func (db *DBx) Select(dest any, query string, args ...any) *DBx {
	return pass(func() error { return db.db.Select(dest, query, args...) }, db)
}

func (db *DBx) SelectContext(ctx context.Context, dest any, query string, args ...any) *DBx {
	return pass(func() error { return db.db.GetContext(ctx, dest, query, args...) }, db)
}

func (db *DBx) Beginx() *Transactionx {
	return Do(func() (*Transactionx, error) {
		result, err := db.db.Beginx()
		return TransactionxOf(result, db), err
	}, db, EmptyTransactionxOf)
}

func (db *DBx) BeginTxx(ctx context.Context, opts *sql.TxOptions) *Transactionx {
	return Do(func() (*Transactionx, error) {
		result, err := db.db.BeginTxx(ctx, opts)
		return TransactionxOf(result, db), err
	}, db, EmptyTransactionxOf)
}

func (db *DBx) Err() error {
	return db.errs.Err()
}

func (db *DBx) Fail(err error) {
	db.errs.Fail(err)
}

func (db *DBx) Link() *error {
	return db.errs.Link()
}

func (db *DBx) LinkTo(err *error) {
	db.errs.LinkTo(err)
}

func (db *DBx) Unwrap() (*sqlx.DB, error) {
	return db.db, db.Err()
}

func (db *DBx) Raw() *sqlx.DB {
	return db.db
}

func DBxOf(db *sqlx.DB, flow Linkable) *DBx {
	return &DBx{db: db, errs: errChainOf(flow)}
}

func NewDBxOf(db *sqlx.DB) *DBx {
	return &DBx{db: db, errs: emptyChain()}
}

func EmptyDBxOf(err error) *DBx {
	return &DBx{errs: errChainOfErr(err)}
}
