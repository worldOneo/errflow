package errflow

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"
)

// DB is the flow wrapper of *sql.DB
type DB struct {
	dB   *sql.DB
	errs errChain
}

// BeginFlow does the same as *sql.DB.Begin but splits the flow.
func (dB *DB) BeginFlow() Splitted[*sql.Tx, *DB] {
	return SplitOf(dB.Begin(), dB)
}

// Begin does the same as *sql.DB but is a noop if this flow already failed.
func (dB *DB) Begin() *sql.Tx {
	return Do(func() (*sql.Tx, error) {
		result, err := dB.dB.Begin()
		return result, err
	}, dB, empty[*sql.Tx])
}

// BeginTxFlow does the same as *sql.DB.BeginTx but splits the flow.
func (dB *DB) BeginTxFlow(ctx context.Context, opts *sql.TxOptions) Splitted[*sql.Tx, *DB] {
	return SplitOf(dB.BeginTx(ctx, opts), dB)
}

// BeginTx does the same as *sql.DB but is a noop if this flow already failed.
func (dB *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) *sql.Tx {
	return Do(func() (*sql.Tx, error) {
		result, err := dB.dB.BeginTx(ctx, opts)
		return result, err
	}, dB, empty[*sql.Tx])
}

// Close does the same as *sql.DB but is a noop if this flow already failed.
// Returs this flow
func (dB *DB) Close() *DB {
	return pass(func() error { return dB.dB.Close() }, dB)
}

// ConnFlow does the same as *sql.DB.Conn but splits the flow.
func (dB *DB) ConnFlow(ctx context.Context) Splitted[*sql.Conn, *DB] {
	return SplitOf(dB.Conn(ctx), dB)
}

// Conn does the same as *sql.DB but is a noop if this flow already failed.
func (dB *DB) Conn(ctx context.Context) *sql.Conn {
	return Do(func() (*sql.Conn, error) {
		result, err := dB.dB.Conn(ctx)
		return result, err
	}, dB, empty[*sql.Conn])
}

// DriverFlow does the same as *sql.DB.Driver but splits the flow.
func (dB *DB) DriverFlow() Splitted[driver.Driver, *DB] {
	return SplitOf(dB.Driver(), dB)
}

// Driver does the same as *sql.DB.Driver .
// This is a noop if the flow already failed.
func (dB *DB) Driver() driver.Driver {
	if dB.errs.Err() != nil {
		return empty[driver.Driver](dB.errs.Err())
	}
	result := dB.dB.Driver()
	return result
}

// ExecFlow does the same as *sql.DB.Exec but splits the flow.
func (dB *DB) ExecFlow(query string, args ...any) Splitted[*Result, *DB] {
	return SplitOf(dB.Exec(query, args...), dB)
}

// Exec does the same as *sql.DB but is a noop if this flow already failed.
func (dB *DB) Exec(query string, args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := dB.dB.Exec(query, args...)
		return ResultOf(result, dB), err
	}, dB, EmptyResultOf)
}

// ExecContextFlow does the same as *sql.DB.ExecContext but splits the flow.
func (dB *DB) ExecContextFlow(ctx context.Context, query string, args ...any) Splitted[*Result, *DB] {
	return SplitOf(dB.ExecContext(ctx, query, args...), dB)
}

// ExecContext does the same as *sql.DB but is a noop if this flow already failed.
func (dB *DB) ExecContext(ctx context.Context, query string, args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := dB.dB.ExecContext(ctx, query, args...)
		return ResultOf(result, dB), err
	}, dB, EmptyResultOf)
}

// Ping does the same as *sql.DB but is a noop if this flow already failed.
// Returs this flow
func (dB *DB) Ping() *DB {
	return pass(func() error { return dB.dB.Ping() }, dB)
}

// PingContext does the same as *sql.DB but is a noop if this flow already failed.
// Returs this flow
func (dB *DB) PingContext(ctx context.Context) *DB {
	return pass(func() error { return dB.dB.PingContext(ctx) }, dB)
}

// PrepareFlow does the same as *sql.DB.Prepare but splits the flow.
func (dB *DB) PrepareFlow(query string) Splitted[*Stmt, *DB] {
	return SplitOf(dB.Prepare(query), dB)
}

// Prepare does the same as *sql.DB but is a noop if this flow already failed.
func (dB *DB) Prepare(query string) *Stmt {
	return Do(func() (*Stmt, error) {
		result, err := dB.dB.Prepare(query)
		return StmtOf(result, dB), err
	}, dB, EmptyStmtOf)
}

// PrepareContextFlow does the same as *sql.DB.PrepareContext but splits the flow.
func (dB *DB) PrepareContextFlow(ctx context.Context, query string) Splitted[*Stmt, *DB] {
	return SplitOf(dB.PrepareContext(ctx, query), dB)
}

// PrepareContext does the same as *sql.DB but is a noop if this flow already failed.
func (dB *DB) PrepareContext(ctx context.Context, query string) *Stmt {
	return Do(func() (*Stmt, error) {
		result, err := dB.dB.PrepareContext(ctx, query)
		return StmtOf(result, dB), err
	}, dB, EmptyStmtOf)
}

// QueryFlow does the same as *sql.DB.Query but splits the flow.
func (dB *DB) QueryFlow(query string, args ...any) Splitted[*sql.Rows, *DB] {
	return SplitOf(dB.Query(query, args...), dB)
}

// Query does the same as *sql.DB but is a noop if this flow already failed.
func (dB *DB) Query(query string, args ...any) *sql.Rows {
	return Do(func() (*sql.Rows, error) {
		result, err := dB.dB.Query(query, args...)
		return result, err
	}, dB, empty[*sql.Rows])
}

// QueryContextFlow does the same as *sql.DB.QueryContext but splits the flow.
func (dB *DB) QueryContextFlow(ctx context.Context, query string, args ...any) Splitted[*sql.Rows, *DB] {
	return SplitOf(dB.QueryContext(ctx, query, args...), dB)
}

// QueryContext does the same as *sql.DB but is a noop if this flow already failed.
func (dB *DB) QueryContext(ctx context.Context, query string, args ...any) *sql.Rows {
	return Do(func() (*sql.Rows, error) {
		result, err := dB.dB.QueryContext(ctx, query, args...)
		return result, err
	}, dB, empty[*sql.Rows])
}

// QueryRowFlow does the same as *sql.DB.QueryRow but splits the flow.
func (dB *DB) QueryRowFlow(query string, args ...any) Splitted[*sql.Row, *DB] {
	return SplitOf(dB.QueryRow(query, args...), dB)
}

// QueryRow does the same as *sql.DB.QueryRow .
// This is a noop if the flow already failed.
func (dB *DB) QueryRow(query string, args ...any) *sql.Row {
	if dB.errs.Err() != nil {
		return empty[*sql.Row](dB.errs.Err())
	}
	result := dB.dB.QueryRow(query, args...)
	return result
}

// QueryRowContextFlow does the same as *sql.DB.QueryRowContext but splits the flow.
func (dB *DB) QueryRowContextFlow(ctx context.Context, query string, args ...any) Splitted[*sql.Row, *DB] {
	return SplitOf(dB.QueryRowContext(ctx, query, args...), dB)
}

// QueryRowContext does the same as *sql.DB.QueryRowContext .
// This is a noop if the flow already failed.
func (dB *DB) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	if dB.errs.Err() != nil {
		return empty[*sql.Row](dB.errs.Err())
	}
	result := dB.dB.QueryRowContext(ctx, query, args...)
	return result
}

// SetConnMaxIdleTime does the same as *sql.DB.SetConnMaxIdleTime but is a noop if this flow
// failed.
// Returns this flow.
func (dB *DB) SetConnMaxIdleTime(d time.Duration) *DB {
	if dB.errs.Err() != nil {
		return dB
	}
	dB.dB.SetConnMaxIdleTime(d)
	return dB
}

// SetConnMaxLifetime does the same as *sql.DB.SetConnMaxLifetime but is a noop if this flow
// failed.
// Returns this flow.
func (dB *DB) SetConnMaxLifetime(d time.Duration) *DB {
	if dB.errs.Err() != nil {
		return dB
	}
	dB.dB.SetConnMaxLifetime(d)
	return dB
}

// SetMaxIdleConns does the same as *sql.DB.SetMaxIdleConns but is a noop if this flow
// failed.
// Returns this flow.
func (dB *DB) SetMaxIdleConns(n int) *DB {
	if dB.errs.Err() != nil {
		return dB
	}
	dB.dB.SetMaxIdleConns(n)
	return dB
}

// SetMaxOpenConns does the same as *sql.DB.SetMaxOpenConns but is a noop if this flow
// failed.
// Returns this flow.
func (dB *DB) SetMaxOpenConns(n int) *DB {
	if dB.errs.Err() != nil {
		return dB
	}
	dB.dB.SetMaxOpenConns(n)
	return dB
}

// StatsFlow does the same as *sql.DB.Stats but splits the flow.
func (dB *DB) StatsFlow() Splitted[sql.DBStats, *DB] {
	return SplitOf(dB.Stats(), dB)
}

// Stats does the same as *sql.DB.Stats .
// This is a noop if the flow already failed.
func (dB *DB) Stats() sql.DBStats {
	if dB.errs.Err() != nil {
		return empty[sql.DBStats](dB.errs.Err())
	}
	result := dB.dB.Stats()
	return result
}

// Err returns the error of this flow if any happend.
func (dB *DB) Err() error {
	return dB.errs.Err()
}

// Fail ends this flow with err
func (dB *DB) Fail(err error) {
	dB.errs.Fail(err)
}

// Link returns the base error of this flow.
func (dB *DB) Link() *error {
	return dB.errs.Link()
}

// LinkTo merges err as base into this flow.
func (dB *DB) LinkTo(err *error) {
	dB.errs.LinkTo(err)
}

// Unwrap returns the internal value and the produced error if any.
func (dB *DB) Unwrap() (*sql.DB, error) {
	return dB.dB, dB.Err()
}

// Raw returns the internal value.
// No Guarantees.
func (dB *DB) Raw() *sql.DB {
	return dB.dB
}

// DBOf create a new DB and is linked to flow.
// They are linked to gether and if one errors the error is propagated.
func DBOf(dB *sql.DB, flow Linkable) *DB {
	return &DB{dB: dB, errs: errChainOf(flow)}
}

// NewDB create a new DB and is the root of a flow.
// It will catch any error that happens in the future.
func NewDB(dB *sql.DB) *DB {
	return &DB{dB: dB, errs: emptyChain()}
}

// EmptyDBOf returns an already failed DB.
// Calls will have no effects on it.
func EmptyDBOf(err error) *DB {
	return &DB{errs: errChainOfErr(err)}
}
