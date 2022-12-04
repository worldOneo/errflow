package errflow

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"

	"github.com/jmoiron/sqlx"
)

// DBx is the flow wrapper of *sqlx.DB
type DBx struct {
	dBx  *sqlx.DB
	errs errChain
}

// BeginTxxFlow does the same as *sqlx.DB.BeginTxx but splits the flow.
func (dBx *DBx) BeginTxxFlow(ctx context.Context, opts *sql.TxOptions) Splitted[*Transactionx, *DBx] {
	return SplitOf(dBx.BeginTxx(ctx, opts), dBx)
}

// BeginTxx does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) BeginTxx(ctx context.Context, opts *sql.TxOptions) *Transactionx {
	return Do(func() (*Transactionx, error) {
		result, err := dBx.dBx.BeginTxx(ctx, opts)
		return TransactionxOf(result, dBx), err
	}, dBx, EmptyTransactionxOf)
}

// BeginxFlow does the same as *sqlx.DB.Beginx but splits the flow.
func (dBx *DBx) BeginxFlow() Splitted[*Transactionx, *DBx] {
	return SplitOf(dBx.Beginx(), dBx)
}

// Beginx does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) Beginx() *Transactionx {
	return Do(func() (*Transactionx, error) {
		result, err := dBx.dBx.Beginx()
		return TransactionxOf(result, dBx), err
	}, dBx, EmptyTransactionxOf)
}

// ConnxFlow does the same as *sqlx.DB.Connx but splits the flow.
func (dBx *DBx) ConnxFlow(ctx context.Context) Splitted[*sqlx.Conn, *DBx] {
	return SplitOf(dBx.Connx(ctx), dBx)
}

// Connx does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) Connx(ctx context.Context) *sqlx.Conn {
	return Do(func() (*sqlx.Conn, error) {
		result, err := dBx.dBx.Connx(ctx)
		return result, err
	}, dBx, empty[*sqlx.Conn])
}

// DriverNameFlow does the same as *sqlx.DB.DriverName but splits the flow.
func (dBx *DBx) DriverNameFlow() Splitted[string, *DBx] {
	return SplitOf(dBx.DriverName(), dBx)
}

// DriverName does the same as *sqlx.DB.DriverName .
// This is a noop if the flow already failed.
func (dBx *DBx) DriverName() string {
	if dBx.errs.Err() != nil {
		return empty[string](dBx.errs.Err())
	}
	result := dBx.dBx.DriverName()
	return result
}

// Get does the same as *sqlx.DB but is a noop if this flow already failed.
// Returs this flow
func (dBx *DBx) Get(dest interface{}, query string, args ...interface{}) *DBx {
	return pass(func() error { return dBx.dBx.Get(dest, query, args...) }, dBx)
}

// GetContext does the same as *sqlx.DB but is a noop if this flow already failed.
// Returs this flow
func (dBx *DBx) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) *DBx {
	return pass(func() error { return dBx.dBx.GetContext(ctx, dest, query, args...) }, dBx)
}

// NamedExecFlow does the same as *sqlx.DB.NamedExec but splits the flow.
func (dBx *DBx) NamedExecFlow(query string, arg interface{}) Splitted[*Result, *DBx] {
	return SplitOf(dBx.NamedExec(query, arg), dBx)
}

// NamedExec does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) NamedExec(query string, arg interface{}) *Result {
	return Do(func() (*Result, error) {
		result, err := dBx.dBx.NamedExec(query, arg)
		return ResultOf(result, dBx), err
	}, dBx, EmptyResultOf)
}

// NamedExecContextFlow does the same as *sqlx.DB.NamedExecContext but splits the flow.
func (dBx *DBx) NamedExecContextFlow(ctx context.Context, query string, arg interface{}) Splitted[*Result, *DBx] {
	return SplitOf(dBx.NamedExecContext(ctx, query, arg), dBx)
}

// NamedExecContext does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) NamedExecContext(ctx context.Context, query string, arg interface{}) *Result {
	return Do(func() (*Result, error) {
		result, err := dBx.dBx.NamedExecContext(ctx, query, arg)
		return ResultOf(result, dBx), err
	}, dBx, EmptyResultOf)
}

// NamedQueryFlow does the same as *sqlx.DB.NamedQuery but splits the flow.
func (dBx *DBx) NamedQueryFlow(query string, arg interface{}) Splitted[*sqlx.Rows, *DBx] {
	return SplitOf(dBx.NamedQuery(query, arg), dBx)
}

// NamedQuery does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) NamedQuery(query string, arg interface{}) *sqlx.Rows {
	return Do(func() (*sqlx.Rows, error) {
		result, err := dBx.dBx.NamedQuery(query, arg)
		return result, err
	}, dBx, empty[*sqlx.Rows])
}

// NamedQueryContextFlow does the same as *sqlx.DB.NamedQueryContext but splits the flow.
func (dBx *DBx) NamedQueryContextFlow(ctx context.Context, query string, arg interface{}) Splitted[*sqlx.Rows, *DBx] {
	return SplitOf(dBx.NamedQueryContext(ctx, query, arg), dBx)
}

// NamedQueryContext does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) NamedQueryContext(ctx context.Context, query string, arg interface{}) *sqlx.Rows {
	return Do(func() (*sqlx.Rows, error) {
		result, err := dBx.dBx.NamedQueryContext(ctx, query, arg)
		return result, err
	}, dBx, empty[*sqlx.Rows])
}

// PrepareNamedFlow does the same as *sqlx.DB.PrepareNamed but splits the flow.
func (dBx *DBx) PrepareNamedFlow(query string) Splitted[*sqlx.NamedStmt, *DBx] {
	return SplitOf(dBx.PrepareNamed(query), dBx)
}

// PrepareNamed does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) PrepareNamed(query string) *sqlx.NamedStmt {
	return Do(func() (*sqlx.NamedStmt, error) {
		result, err := dBx.dBx.PrepareNamed(query)
		return result, err
	}, dBx, empty[*sqlx.NamedStmt])
}

// PrepareNamedContextFlow does the same as *sqlx.DB.PrepareNamedContext but splits the flow.
func (dBx *DBx) PrepareNamedContextFlow(ctx context.Context, query string) Splitted[*sqlx.NamedStmt, *DBx] {
	return SplitOf(dBx.PrepareNamedContext(ctx, query), dBx)
}

// PrepareNamedContext does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) PrepareNamedContext(ctx context.Context, query string) *sqlx.NamedStmt {
	return Do(func() (*sqlx.NamedStmt, error) {
		result, err := dBx.dBx.PrepareNamedContext(ctx, query)
		return result, err
	}, dBx, empty[*sqlx.NamedStmt])
}

// PreparexFlow does the same as *sqlx.DB.Preparex but splits the flow.
func (dBx *DBx) PreparexFlow(query string) Splitted[*Stmtx, *DBx] {
	return SplitOf(dBx.Preparex(query), dBx)
}

// Preparex does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) Preparex(query string) *Stmtx {
	return Do(func() (*Stmtx, error) {
		result, err := dBx.dBx.Preparex(query)
		return StmtxOf(result, dBx), err
	}, dBx, EmptyStmtxOf)
}

// PreparexContextFlow does the same as *sqlx.DB.PreparexContext but splits the flow.
func (dBx *DBx) PreparexContextFlow(ctx context.Context, query string) Splitted[*Stmtx, *DBx] {
	return SplitOf(dBx.PreparexContext(ctx, query), dBx)
}

// PreparexContext does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) PreparexContext(ctx context.Context, query string) *Stmtx {
	return Do(func() (*Stmtx, error) {
		result, err := dBx.dBx.PreparexContext(ctx, query)
		return StmtxOf(result, dBx), err
	}, dBx, EmptyStmtxOf)
}

// QueryRowxFlow does the same as *sqlx.DB.QueryRowx but splits the flow.
func (dBx *DBx) QueryRowxFlow(query string, args ...interface{}) Splitted[*sqlx.Row, *DBx] {
	return SplitOf(dBx.QueryRowx(query, args...), dBx)
}

// QueryRowx does the same as *sqlx.DB.QueryRowx .
// This is a noop if the flow already failed.
func (dBx *DBx) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	if dBx.errs.Err() != nil {
		return empty[*sqlx.Row](dBx.errs.Err())
	}
	result := dBx.dBx.QueryRowx(query, args...)
	return result
}

// QueryRowxContextFlow does the same as *sqlx.DB.QueryRowxContext but splits the flow.
func (dBx *DBx) QueryRowxContextFlow(ctx context.Context, query string, args ...interface{}) Splitted[*sqlx.Row, *DBx] {
	return SplitOf(dBx.QueryRowxContext(ctx, query, args...), dBx)
}

// QueryRowxContext does the same as *sqlx.DB.QueryRowxContext .
// This is a noop if the flow already failed.
func (dBx *DBx) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	if dBx.errs.Err() != nil {
		return empty[*sqlx.Row](dBx.errs.Err())
	}
	result := dBx.dBx.QueryRowxContext(ctx, query, args...)
	return result
}

// QueryxFlow does the same as *sqlx.DB.Queryx but splits the flow.
func (dBx *DBx) QueryxFlow(query string, args ...interface{}) Splitted[*sqlx.Rows, *DBx] {
	return SplitOf(dBx.Queryx(query, args...), dBx)
}

// Queryx does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) Queryx(query string, args ...interface{}) *sqlx.Rows {
	return Do(func() (*sqlx.Rows, error) {
		result, err := dBx.dBx.Queryx(query, args...)
		return result, err
	}, dBx, empty[*sqlx.Rows])
}

// QueryxContextFlow does the same as *sqlx.DB.QueryxContext but splits the flow.
func (dBx *DBx) QueryxContextFlow(ctx context.Context, query string, args ...interface{}) Splitted[*sqlx.Rows, *DBx] {
	return SplitOf(dBx.QueryxContext(ctx, query, args...), dBx)
}

// QueryxContext does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) QueryxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Rows {
	return Do(func() (*sqlx.Rows, error) {
		result, err := dBx.dBx.QueryxContext(ctx, query, args...)
		return result, err
	}, dBx, empty[*sqlx.Rows])
}

// RebindFlow does the same as *sqlx.DB.Rebind but splits the flow.
func (dBx *DBx) RebindFlow(query string) Splitted[string, *DBx] {
	return SplitOf(dBx.Rebind(query), dBx)
}

// Rebind does the same as *sqlx.DB.Rebind .
// This is a noop if the flow already failed.
func (dBx *DBx) Rebind(query string) string {
	if dBx.errs.Err() != nil {
		return empty[string](dBx.errs.Err())
	}
	result := dBx.dBx.Rebind(query)
	return result
}

// Select does the same as *sqlx.DB but is a noop if this flow already failed.
// Returs this flow
func (dBx *DBx) Select(dest interface{}, query string, args ...interface{}) *DBx {
	return pass(func() error { return dBx.dBx.Select(dest, query, args...) }, dBx)
}

// SelectContext does the same as *sqlx.DB but is a noop if this flow already failed.
// Returs this flow
func (dBx *DBx) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) *DBx {
	return pass(func() error { return dBx.dBx.SelectContext(ctx, dest, query, args...) }, dBx)
}

// BeginFlow does the same as *sqlx.DB.Begin but splits the flow.
func (dBx *DBx) BeginFlow() Splitted[*sql.Tx, *DBx] {
	return SplitOf(dBx.Begin(), dBx)
}

// Begin does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) Begin() *sql.Tx {
	return Do(func() (*sql.Tx, error) {
		result, err := dBx.dBx.Begin()
		return result, err
	}, dBx, empty[*sql.Tx])
}

// BeginTxFlow does the same as *sqlx.DB.BeginTx but splits the flow.
func (dBx *DBx) BeginTxFlow(ctx context.Context, opts *sql.TxOptions) Splitted[*sql.Tx, *DBx] {
	return SplitOf(dBx.BeginTx(ctx, opts), dBx)
}

// BeginTx does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) BeginTx(ctx context.Context, opts *sql.TxOptions) *sql.Tx {
	return Do(func() (*sql.Tx, error) {
		result, err := dBx.dBx.BeginTx(ctx, opts)
		return result, err
	}, dBx, empty[*sql.Tx])
}

// Close does the same as *sqlx.DB but is a noop if this flow already failed.
// Returs this flow
func (dBx *DBx) Close() *DBx {
	return pass(func() error { return dBx.dBx.Close() }, dBx)
}

// ConnFlow does the same as *sqlx.DB.Conn but splits the flow.
func (dBx *DBx) ConnFlow(ctx context.Context) Splitted[*sql.Conn, *DBx] {
	return SplitOf(dBx.Conn(ctx), dBx)
}

// Conn does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) Conn(ctx context.Context) *sql.Conn {
	return Do(func() (*sql.Conn, error) {
		result, err := dBx.dBx.Conn(ctx)
		return result, err
	}, dBx, empty[*sql.Conn])
}

// DriverFlow does the same as *sqlx.DB.Driver but splits the flow.
func (dBx *DBx) DriverFlow() Splitted[driver.Driver, *DBx] {
	return SplitOf(dBx.Driver(), dBx)
}

// Driver does the same as *sqlx.DB.Driver .
// This is a noop if the flow already failed.
func (dBx *DBx) Driver() driver.Driver {
	if dBx.errs.Err() != nil {
		return empty[driver.Driver](dBx.errs.Err())
	}
	result := dBx.dBx.Driver()
	return result
}

// ExecFlow does the same as *sqlx.DB.Exec but splits the flow.
func (dBx *DBx) ExecFlow(query string, args ...any) Splitted[*Result, *DBx] {
	return SplitOf(dBx.Exec(query, args...), dBx)
}

// Exec does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) Exec(query string, args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := dBx.dBx.Exec(query, args...)
		return ResultOf(result, dBx), err
	}, dBx, EmptyResultOf)
}

// ExecContextFlow does the same as *sqlx.DB.ExecContext but splits the flow.
func (dBx *DBx) ExecContextFlow(ctx context.Context, query string, args ...any) Splitted[*Result, *DBx] {
	return SplitOf(dBx.ExecContext(ctx, query, args...), dBx)
}

// ExecContext does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) ExecContext(ctx context.Context, query string, args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := dBx.dBx.ExecContext(ctx, query, args...)
		return ResultOf(result, dBx), err
	}, dBx, EmptyResultOf)
}

// Ping does the same as *sqlx.DB but is a noop if this flow already failed.
// Returs this flow
func (dBx *DBx) Ping() *DBx {
	return pass(func() error { return dBx.dBx.Ping() }, dBx)
}

// PingContext does the same as *sqlx.DB but is a noop if this flow already failed.
// Returs this flow
func (dBx *DBx) PingContext(ctx context.Context) *DBx {
	return pass(func() error { return dBx.dBx.PingContext(ctx) }, dBx)
}

// PrepareFlow does the same as *sqlx.DB.Prepare but splits the flow.
func (dBx *DBx) PrepareFlow(query string) Splitted[*Stmt, *DBx] {
	return SplitOf(dBx.Prepare(query), dBx)
}

// Prepare does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) Prepare(query string) *Stmt {
	return Do(func() (*Stmt, error) {
		result, err := dBx.dBx.Prepare(query)
		return StmtOf(result, dBx), err
	}, dBx, EmptyStmtOf)
}

// PrepareContextFlow does the same as *sqlx.DB.PrepareContext but splits the flow.
func (dBx *DBx) PrepareContextFlow(ctx context.Context, query string) Splitted[*Stmt, *DBx] {
	return SplitOf(dBx.PrepareContext(ctx, query), dBx)
}

// PrepareContext does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) PrepareContext(ctx context.Context, query string) *Stmt {
	return Do(func() (*Stmt, error) {
		result, err := dBx.dBx.PrepareContext(ctx, query)
		return StmtOf(result, dBx), err
	}, dBx, EmptyStmtOf)
}

// QueryFlow does the same as *sqlx.DB.Query but splits the flow.
func (dBx *DBx) QueryFlow(query string, args ...any) Splitted[*sql.Rows, *DBx] {
	return SplitOf(dBx.Query(query, args...), dBx)
}

// Query does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) Query(query string, args ...any) *sql.Rows {
	return Do(func() (*sql.Rows, error) {
		result, err := dBx.dBx.Query(query, args...)
		return result, err
	}, dBx, empty[*sql.Rows])
}

// QueryContextFlow does the same as *sqlx.DB.QueryContext but splits the flow.
func (dBx *DBx) QueryContextFlow(ctx context.Context, query string, args ...any) Splitted[*sql.Rows, *DBx] {
	return SplitOf(dBx.QueryContext(ctx, query, args...), dBx)
}

// QueryContext does the same as *sqlx.DB but is a noop if this flow already failed.
func (dBx *DBx) QueryContext(ctx context.Context, query string, args ...any) *sql.Rows {
	return Do(func() (*sql.Rows, error) {
		result, err := dBx.dBx.QueryContext(ctx, query, args...)
		return result, err
	}, dBx, empty[*sql.Rows])
}

// QueryRowFlow does the same as *sqlx.DB.QueryRow but splits the flow.
func (dBx *DBx) QueryRowFlow(query string, args ...any) Splitted[*sql.Row, *DBx] {
	return SplitOf(dBx.QueryRow(query, args...), dBx)
}

// QueryRow does the same as *sqlx.DB.QueryRow .
// This is a noop if the flow already failed.
func (dBx *DBx) QueryRow(query string, args ...any) *sql.Row {
	if dBx.errs.Err() != nil {
		return empty[*sql.Row](dBx.errs.Err())
	}
	result := dBx.dBx.QueryRow(query, args...)
	return result
}

// QueryRowContextFlow does the same as *sqlx.DB.QueryRowContext but splits the flow.
func (dBx *DBx) QueryRowContextFlow(ctx context.Context, query string, args ...any) Splitted[*sql.Row, *DBx] {
	return SplitOf(dBx.QueryRowContext(ctx, query, args...), dBx)
}

// QueryRowContext does the same as *sqlx.DB.QueryRowContext .
// This is a noop if the flow already failed.
func (dBx *DBx) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	if dBx.errs.Err() != nil {
		return empty[*sql.Row](dBx.errs.Err())
	}
	result := dBx.dBx.QueryRowContext(ctx, query, args...)
	return result
}

// SetConnMaxIdleTime does the same as *sqlx.DB.SetConnMaxIdleTime but is a noop if this flow
// failed.
// Returns this flow.
func (dBx *DBx) SetConnMaxIdleTime(d time.Duration) *DBx {
	if dBx.errs.Err() != nil {
		return dBx
	}
	dBx.dBx.SetConnMaxIdleTime(d)
	return dBx
}

// SetConnMaxLifetime does the same as *sqlx.DB.SetConnMaxLifetime but is a noop if this flow
// failed.
// Returns this flow.
func (dBx *DBx) SetConnMaxLifetime(d time.Duration) *DBx {
	if dBx.errs.Err() != nil {
		return dBx
	}
	dBx.dBx.SetConnMaxLifetime(d)
	return dBx
}

// SetMaxIdleConns does the same as *sqlx.DB.SetMaxIdleConns but is a noop if this flow
// failed.
// Returns this flow.
func (dBx *DBx) SetMaxIdleConns(n int) *DBx {
	if dBx.errs.Err() != nil {
		return dBx
	}
	dBx.dBx.SetMaxIdleConns(n)
	return dBx
}

// SetMaxOpenConns does the same as *sqlx.DB.SetMaxOpenConns but is a noop if this flow
// failed.
// Returns this flow.
func (dBx *DBx) SetMaxOpenConns(n int) *DBx {
	if dBx.errs.Err() != nil {
		return dBx
	}
	dBx.dBx.SetMaxOpenConns(n)
	return dBx
}

// StatsFlow does the same as *sqlx.DB.Stats but splits the flow.
func (dBx *DBx) StatsFlow() Splitted[sql.DBStats, *DBx] {
	return SplitOf(dBx.Stats(), dBx)
}

// Stats does the same as *sqlx.DB.Stats .
// This is a noop if the flow already failed.
func (dBx *DBx) Stats() sql.DBStats {
	if dBx.errs.Err() != nil {
		return empty[sql.DBStats](dBx.errs.Err())
	}
	result := dBx.dBx.Stats()
	return result
}

// Err returns the error of this flow if any happend.
func (dBx *DBx) Err() error {
	return dBx.errs.Err()
}

// Fail ends this flow with err
func (dBx *DBx) Fail(err error) {
	dBx.errs.Fail(err)
}

// Link returns the base error of this flow.
func (dBx *DBx) Link() *error {
	return dBx.errs.Link()
}

// LinkTo merges err as base into this flow.
func (dBx *DBx) LinkTo(err *error) {
	dBx.errs.LinkTo(err)
}

// Unwrap returns the internal value and the produced error if any.
func (dBx *DBx) Unwrap() (*sqlx.DB, error) {
	return dBx.dBx, dBx.Err()
}

// Raw returns the internal value.
// No Guarantees.
func (dBx *DBx) Raw() *sqlx.DB {
	return dBx.dBx
}

// DBxOf create a new DBx and is linked to flow.
// They are linked to gether and if one errors the error is propagated.
func DBxOf(dBx *sqlx.DB, flow Linkable) *DBx {
	return &DBx{dBx: dBx, errs: errChainOf(flow)}
}

// NewDBx create a new DBx and is the root of a flow.
// It will catch any error that happens in the future.
func NewDBx(dBx *sqlx.DB) *DBx {
	return &DBx{dBx: dBx, errs: emptyChain()}
}

// EmptyDBxOf returns an already failed DBx.
// Calls will have no effects on it.
func EmptyDBxOf(err error) *DBx {
	return &DBx{errs: errChainOfErr(err)}
}
