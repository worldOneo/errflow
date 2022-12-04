package errflow

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type DBx struct {
	dBx  *sqlx.DB
	errs errChain
}

func (dBx *DBx) BeginTxxFlow(ctx context.Context, opts *sql.TxOptions) Splitted[*Transactionx, *DBx] {
	return SplitOf(dBx.BeginTxx(ctx, opts), dBx)
}

func (dBx *DBx) BeginTxx(ctx context.Context, opts *sql.TxOptions) *Transactionx {
	return Do(func() (*Transactionx, error) {
		result, err := dBx.dBx.BeginTxx(ctx, opts)
		return TransactionxOf(result, dBx), err
	}, dBx, EmptyTransactionxOf)
}

func (dBx *DBx) BeginxFlow() Splitted[*Transactionx, *DBx] {
	return SplitOf(dBx.Beginx(), dBx)
}

func (dBx *DBx) Beginx() *Transactionx {
	return Do(func() (*Transactionx, error) {
		result, err := dBx.dBx.Beginx()
		return TransactionxOf(result, dBx), err
	}, dBx, EmptyTransactionxOf)
}

func (dBx *DBx) ConnxFlow(ctx context.Context) Splitted[*sqlx.Conn, *DBx] {
	return SplitOf(dBx.Connx(ctx), dBx)
}

func (dBx *DBx) Connx(ctx context.Context) *sqlx.Conn {
	return Do(func() (*sqlx.Conn, error) {
		result, err := dBx.dBx.Connx(ctx)
		return result, err
	}, dBx, empty[*sqlx.Conn])
}

func (dBx *DBx) DriverNameFlow() Splitted[string, *DBx] {
	return SplitOf(dBx.DriverName(), dBx)
}

func (dBx *DBx) DriverName() string {
	return dBx.dBx.DriverName()
}

func (dBx *DBx) Get(dest interface{}, query string, args ...interface{}) *DBx {
	return pass(func() error { return dBx.dBx.Get(dest, query, args...) }, dBx)
}

func (dBx *DBx) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) *DBx {
	return pass(func() error { return dBx.dBx.GetContext(ctx, dest, query, args...) }, dBx)
}

func (dBx *DBx) NamedExecFlow(query string, arg interface{}) Splitted[sql.Result, *DBx] {
	return SplitOf(dBx.NamedExec(query, arg), dBx)
}

func (dBx *DBx) NamedExec(query string, arg interface{}) sql.Result {
	return Do(func() (sql.Result, error) {
		result, err := dBx.dBx.NamedExec(query, arg)
		return result, err
	}, dBx, empty[sql.Result])
}

func (dBx *DBx) NamedExecContextFlow(ctx context.Context, query string, arg interface{}) Splitted[sql.Result, *DBx] {
	return SplitOf(dBx.NamedExecContext(ctx, query, arg), dBx)
}

func (dBx *DBx) NamedExecContext(ctx context.Context, query string, arg interface{}) sql.Result {
	return Do(func() (sql.Result, error) {
		result, err := dBx.dBx.NamedExecContext(ctx, query, arg)
		return result, err
	}, dBx, empty[sql.Result])
}

func (dBx *DBx) NamedQueryFlow(query string, arg interface{}) Splitted[*sqlx.Rows, *DBx] {
	return SplitOf(dBx.NamedQuery(query, arg), dBx)
}

func (dBx *DBx) NamedQuery(query string, arg interface{}) *sqlx.Rows {
	return Do(func() (*sqlx.Rows, error) {
		result, err := dBx.dBx.NamedQuery(query, arg)
		return result, err
	}, dBx, empty[*sqlx.Rows])
}

func (dBx *DBx) NamedQueryContextFlow(ctx context.Context, query string, arg interface{}) Splitted[*sqlx.Rows, *DBx] {
	return SplitOf(dBx.NamedQueryContext(ctx, query, arg), dBx)
}

func (dBx *DBx) NamedQueryContext(ctx context.Context, query string, arg interface{}) *sqlx.Rows {
	return Do(func() (*sqlx.Rows, error) {
		result, err := dBx.dBx.NamedQueryContext(ctx, query, arg)
		return result, err
	}, dBx, empty[*sqlx.Rows])
}

func (dBx *DBx) PrepareNamedFlow(query string) Splitted[*sqlx.NamedStmt, *DBx] {
	return SplitOf(dBx.PrepareNamed(query), dBx)
}

func (dBx *DBx) PrepareNamed(query string) *sqlx.NamedStmt {
	return Do(func() (*sqlx.NamedStmt, error) {
		result, err := dBx.dBx.PrepareNamed(query)
		return result, err
	}, dBx, empty[*sqlx.NamedStmt])
}

func (dBx *DBx) PrepareNamedContextFlow(ctx context.Context, query string) Splitted[*sqlx.NamedStmt, *DBx] {
	return SplitOf(dBx.PrepareNamedContext(ctx, query), dBx)
}

func (dBx *DBx) PrepareNamedContext(ctx context.Context, query string) *sqlx.NamedStmt {
	return Do(func() (*sqlx.NamedStmt, error) {
		result, err := dBx.dBx.PrepareNamedContext(ctx, query)
		return result, err
	}, dBx, empty[*sqlx.NamedStmt])
}

func (dBx *DBx) PreparexFlow(query string) Splitted[*sqlx.Stmt, *DBx] {
	return SplitOf(dBx.Preparex(query), dBx)
}

func (dBx *DBx) Preparex(query string) *sqlx.Stmt {
	return Do(func() (*sqlx.Stmt, error) {
		result, err := dBx.dBx.Preparex(query)
		return result, err
	}, dBx, empty[*sqlx.Stmt])
}

func (dBx *DBx) PreparexContextFlow(ctx context.Context, query string) Splitted[*sqlx.Stmt, *DBx] {
	return SplitOf(dBx.PreparexContext(ctx, query), dBx)
}

func (dBx *DBx) PreparexContext(ctx context.Context, query string) *sqlx.Stmt {
	return Do(func() (*sqlx.Stmt, error) {
		result, err := dBx.dBx.PreparexContext(ctx, query)
		return result, err
	}, dBx, empty[*sqlx.Stmt])
}

func (dBx *DBx) QueryRowxFlow(query string, args ...interface{}) Splitted[*sqlx.Row, *DBx] {
	return SplitOf(dBx.QueryRowx(query, args...), dBx)
}

func (dBx *DBx) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	return dBx.dBx.QueryRowx(query, args...)
}

func (dBx *DBx) QueryRowxContextFlow(ctx context.Context, query string, args ...interface{}) Splitted[*sqlx.Row, *DBx] {
	return SplitOf(dBx.QueryRowxContext(ctx, query, args...), dBx)
}

func (dBx *DBx) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return dBx.dBx.QueryRowxContext(ctx, query, args...)
}

func (dBx *DBx) QueryxFlow(query string, args ...interface{}) Splitted[*sqlx.Rows, *DBx] {
	return SplitOf(dBx.Queryx(query, args...), dBx)
}

func (dBx *DBx) Queryx(query string, args ...interface{}) *sqlx.Rows {
	return Do(func() (*sqlx.Rows, error) {
		result, err := dBx.dBx.Queryx(query, args...)
		return result, err
	}, dBx, empty[*sqlx.Rows])
}

func (dBx *DBx) QueryxContextFlow(ctx context.Context, query string, args ...interface{}) Splitted[*sqlx.Rows, *DBx] {
	return SplitOf(dBx.QueryxContext(ctx, query, args...), dBx)
}

func (dBx *DBx) QueryxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Rows {
	return Do(func() (*sqlx.Rows, error) {
		result, err := dBx.dBx.QueryxContext(ctx, query, args...)
		return result, err
	}, dBx, empty[*sqlx.Rows])
}

func (dBx *DBx) RebindFlow(query string) Splitted[string, *DBx] {
	return SplitOf(dBx.Rebind(query), dBx)
}

func (dBx *DBx) Rebind(query string) string {
	return dBx.dBx.Rebind(query)
}

func (dBx *DBx) Select(dest interface{}, query string, args ...interface{}) *DBx {
	return pass(func() error { return dBx.dBx.Select(dest, query, args...) }, dBx)
}

func (dBx *DBx) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) *DBx {
	return pass(func() error { return dBx.dBx.SelectContext(ctx, dest, query, args...) }, dBx)
}

func (dBx *DBx) UnsafeFlow() Splitted[*sqlx.DB, *DBx] {
	return SplitOf(dBx.Unsafe(), dBx)
}

func (dBx *DBx) Unsafe() *sqlx.DB {
	return dBx.dBx.Unsafe()
}

func (dBx *DBx) Err() error {
	return dBx.errs.Err()
}

func (dBx *DBx) Fail(err error) {
	dBx.errs.Fail(err)
}

func (dBx *DBx) Link() *error {
	return dBx.errs.Link()
}

func (dBx *DBx) LinkTo(err *error) {
	dBx.errs.LinkTo(err)
}

func (dBx *DBx) Unwrap() (*sqlx.DB, error) {
	return dBx.dBx, dBx.Err()
}

func (dBx *DBx) Raw() *sqlx.DB {
	return dBx.dBx
}

func DBxOf(dBx *sqlx.DB, flow Linkable) *DBx {
	return &DBx{dBx: dBx, errs: errChainOf(flow)}
}

func NewDBx(dBx *sqlx.DB) *DBx {
	return &DBx{dBx: dBx, errs: emptyChain()}
}

func EmptyDBxOf(err error) *DBx {
	return &DBx{errs: errChainOfErr(err)}
}
