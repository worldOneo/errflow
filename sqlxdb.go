package errflow

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type VirtualDB struct {
	db   *sqlx.DB
	errs errChain
}

func DB(db *sqlx.DB) *VirtualDB {
	return &VirtualDB{db: db, errs: emptyChain()}
}

func (db *VirtualDB) BeginTxx(ctx context.Context, opts *sql.TxOptions) *VirtualTxx {
	return BeginTxx(ctx, opts, db.db)
}

func (db *VirtualDB) Beginx() *VirtualTxx {
	return Do(func() (*VirtualTxx, error) {
		tx, err := db.db.Beginx()
		return &VirtualTxx{tx: tx, errs: errChainOf(db)}, err
	}, db, emptyTransactionOf)
}

func (db *VirtualDB) Get(dest interface{}, query string, args ...interface{}) *VirtualDB {
	return db.GetContext(context.Background(), dest, query, args...)
}

func (db *VirtualDB) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) *VirtualDB {
	return pass(func() error { return db.db.GetContext(ctx, dest, query, args...) }, db)
}

func (db *VirtualDB) Select(dest interface{}, query string, args ...interface{}) *VirtualDB {
	return db.SelectContext(context.Background(), dest, query, args...)
}

func (db *VirtualDB) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) *VirtualDB {
	return pass(func() error { return db.db.SelectContext(ctx, dest, query, args...) }, db)
}

func (db *VirtualDB) NamedExec(query string, arg interface{}) sql.Result {
	return db.NamedExecContext(context.Background(), query, arg)
}

func (db *VirtualDB) NamedExecContext(ctx context.Context, query string, arg interface{}) sql.Result {
	return Do(func() (sql.Result, error) {
		result, err := db.db.NamedExec(query, arg)
		return result, err
	}, db, emptyResultOf)
}

func (db *VirtualDB) Exec(query string, args ...interface{}) sql.Result {
	return db.ExecContext(context.Background(), query, args...)
}

func (db *VirtualDB) ExecContext(ctx context.Context, query string, args ...interface{}) sql.Result {
	return Do(func() (sql.Result, error) {
		result, err := db.db.ExecContext(ctx, query, args...)
		return result, err
	}, db, emptyResultOf)
}

func (db *VirtualDB) Err() error {
	return db.errs.Err()
}

func (db *VirtualDB) Fail(err error) {
	db.errs.Fail(err)
}

func (db *VirtualDB) Link() *error {
	return db.errs.Link()
}

func (res *VirtualDB) LinkTo(err *error)  {
	res.errs.LinkTo(err)
}
