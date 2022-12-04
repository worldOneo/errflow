package errflow

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Transactionx is the flow wrapper of *sqlx.Tx
type Transactionx struct {
	transactionx *sqlx.Tx
	errs         errChain
}

// DriverNameFlow does the same as *sqlx.Tx.DriverName but splits the flow.
func (transactionx *Transactionx) DriverNameFlow() Splitted[string, *Transactionx] {
	return SplitOf(transactionx.DriverName(), transactionx)
}

// DriverName does the same as *sqlx.Tx.DriverName .
// This is a noop if the flow already failed.
func (transactionx *Transactionx) DriverName() string {
	if transactionx.errs.Err() != nil {
		return empty[string](transactionx.errs.Err())
	}
	result := transactionx.transactionx.DriverName()
	return result
}

// Get does the same as *sqlx.Tx but is a noop if this flow already failed.
// Returs this flow
func (transactionx *Transactionx) Get(dest interface{}, query string, args ...interface{}) *Transactionx {
	return pass(func() error { return transactionx.transactionx.Get(dest, query, args...) }, transactionx)
}

// GetContext does the same as *sqlx.Tx but is a noop if this flow already failed.
// Returs this flow
func (transactionx *Transactionx) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) *Transactionx {
	return pass(func() error { return transactionx.transactionx.GetContext(ctx, dest, query, args...) }, transactionx)
}

// NamedExecFlow does the same as *sqlx.Tx.NamedExec but splits the flow.
func (transactionx *Transactionx) NamedExecFlow(query string, arg interface{}) Splitted[*Result, *Transactionx] {
	return SplitOf(transactionx.NamedExec(query, arg), transactionx)
}

// NamedExec does the same as *sqlx.Tx but is a noop if this flow already failed.
func (transactionx *Transactionx) NamedExec(query string, arg interface{}) *Result {
	return Do(func() (*Result, error) {
		result, err := transactionx.transactionx.NamedExec(query, arg)
		return ResultOf(result, transactionx), err
	}, transactionx, EmptyResultOf)
}

// NamedExecContextFlow does the same as *sqlx.Tx.NamedExecContext but splits the flow.
func (transactionx *Transactionx) NamedExecContextFlow(ctx context.Context, query string, arg interface{}) Splitted[*Result, *Transactionx] {
	return SplitOf(transactionx.NamedExecContext(ctx, query, arg), transactionx)
}

// NamedExecContext does the same as *sqlx.Tx but is a noop if this flow already failed.
func (transactionx *Transactionx) NamedExecContext(ctx context.Context, query string, arg interface{}) *Result {
	return Do(func() (*Result, error) {
		result, err := transactionx.transactionx.NamedExecContext(ctx, query, arg)
		return ResultOf(result, transactionx), err
	}, transactionx, EmptyResultOf)
}

// NamedQueryFlow does the same as *sqlx.Tx.NamedQuery but splits the flow.
func (transactionx *Transactionx) NamedQueryFlow(query string, arg interface{}) Splitted[*sqlx.Rows, *Transactionx] {
	return SplitOf(transactionx.NamedQuery(query, arg), transactionx)
}

// NamedQuery does the same as *sqlx.Tx but is a noop if this flow already failed.
func (transactionx *Transactionx) NamedQuery(query string, arg interface{}) *sqlx.Rows {
	return Do(func() (*sqlx.Rows, error) {
		result, err := transactionx.transactionx.NamedQuery(query, arg)
		return result, err
	}, transactionx, empty[*sqlx.Rows])
}

// NamedStmtFlow does the same as *sqlx.Tx.NamedStmt but splits the flow.
func (transactionx *Transactionx) NamedStmtFlow(stmt *sqlx.NamedStmt) Splitted[*sqlx.NamedStmt, *Transactionx] {
	return SplitOf(transactionx.NamedStmt(stmt), transactionx)
}

// NamedStmt does the same as *sqlx.Tx.NamedStmt .
// This is a noop if the flow already failed.
func (transactionx *Transactionx) NamedStmt(stmt *sqlx.NamedStmt) *sqlx.NamedStmt {
	if transactionx.errs.Err() != nil {
		return empty[*sqlx.NamedStmt](transactionx.errs.Err())
	}
	result := transactionx.transactionx.NamedStmt(stmt)
	return result
}

// NamedStmtContextFlow does the same as *sqlx.Tx.NamedStmtContext but splits the flow.
func (transactionx *Transactionx) NamedStmtContextFlow(ctx context.Context, stmt *sqlx.NamedStmt) Splitted[*sqlx.NamedStmt, *Transactionx] {
	return SplitOf(transactionx.NamedStmtContext(ctx, stmt), transactionx)
}

// NamedStmtContext does the same as *sqlx.Tx.NamedStmtContext .
// This is a noop if the flow already failed.
func (transactionx *Transactionx) NamedStmtContext(ctx context.Context, stmt *sqlx.NamedStmt) *sqlx.NamedStmt {
	if transactionx.errs.Err() != nil {
		return empty[*sqlx.NamedStmt](transactionx.errs.Err())
	}
	result := transactionx.transactionx.NamedStmtContext(ctx, stmt)
	return result
}

// PrepareNamedFlow does the same as *sqlx.Tx.PrepareNamed but splits the flow.
func (transactionx *Transactionx) PrepareNamedFlow(query string) Splitted[*sqlx.NamedStmt, *Transactionx] {
	return SplitOf(transactionx.PrepareNamed(query), transactionx)
}

// PrepareNamed does the same as *sqlx.Tx but is a noop if this flow already failed.
func (transactionx *Transactionx) PrepareNamed(query string) *sqlx.NamedStmt {
	return Do(func() (*sqlx.NamedStmt, error) {
		result, err := transactionx.transactionx.PrepareNamed(query)
		return result, err
	}, transactionx, empty[*sqlx.NamedStmt])
}

// PrepareNamedContextFlow does the same as *sqlx.Tx.PrepareNamedContext but splits the flow.
func (transactionx *Transactionx) PrepareNamedContextFlow(ctx context.Context, query string) Splitted[*sqlx.NamedStmt, *Transactionx] {
	return SplitOf(transactionx.PrepareNamedContext(ctx, query), transactionx)
}

// PrepareNamedContext does the same as *sqlx.Tx but is a noop if this flow already failed.
func (transactionx *Transactionx) PrepareNamedContext(ctx context.Context, query string) *sqlx.NamedStmt {
	return Do(func() (*sqlx.NamedStmt, error) {
		result, err := transactionx.transactionx.PrepareNamedContext(ctx, query)
		return result, err
	}, transactionx, empty[*sqlx.NamedStmt])
}

// PreparexFlow does the same as *sqlx.Tx.Preparex but splits the flow.
func (transactionx *Transactionx) PreparexFlow(query string) Splitted[*Stmtx, *Transactionx] {
	return SplitOf(transactionx.Preparex(query), transactionx)
}

// Preparex does the same as *sqlx.Tx but is a noop if this flow already failed.
func (transactionx *Transactionx) Preparex(query string) *Stmtx {
	return Do(func() (*Stmtx, error) {
		result, err := transactionx.transactionx.Preparex(query)
		return StmtxOf(result, transactionx), err
	}, transactionx, EmptyStmtxOf)
}

// PreparexContextFlow does the same as *sqlx.Tx.PreparexContext but splits the flow.
func (transactionx *Transactionx) PreparexContextFlow(ctx context.Context, query string) Splitted[*Stmtx, *Transactionx] {
	return SplitOf(transactionx.PreparexContext(ctx, query), transactionx)
}

// PreparexContext does the same as *sqlx.Tx but is a noop if this flow already failed.
func (transactionx *Transactionx) PreparexContext(ctx context.Context, query string) *Stmtx {
	return Do(func() (*Stmtx, error) {
		result, err := transactionx.transactionx.PreparexContext(ctx, query)
		return StmtxOf(result, transactionx), err
	}, transactionx, EmptyStmtxOf)
}

// QueryRowxFlow does the same as *sqlx.Tx.QueryRowx but splits the flow.
func (transactionx *Transactionx) QueryRowxFlow(query string, args ...interface{}) Splitted[*sqlx.Row, *Transactionx] {
	return SplitOf(transactionx.QueryRowx(query, args...), transactionx)
}

// QueryRowx does the same as *sqlx.Tx.QueryRowx .
// This is a noop if the flow already failed.
func (transactionx *Transactionx) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	if transactionx.errs.Err() != nil {
		return empty[*sqlx.Row](transactionx.errs.Err())
	}
	result := transactionx.transactionx.QueryRowx(query, args...)
	return result
}

// QueryRowxContextFlow does the same as *sqlx.Tx.QueryRowxContext but splits the flow.
func (transactionx *Transactionx) QueryRowxContextFlow(ctx context.Context, query string, args ...interface{}) Splitted[*sqlx.Row, *Transactionx] {
	return SplitOf(transactionx.QueryRowxContext(ctx, query, args...), transactionx)
}

// QueryRowxContext does the same as *sqlx.Tx.QueryRowxContext .
// This is a noop if the flow already failed.
func (transactionx *Transactionx) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	if transactionx.errs.Err() != nil {
		return empty[*sqlx.Row](transactionx.errs.Err())
	}
	result := transactionx.transactionx.QueryRowxContext(ctx, query, args...)
	return result
}

// QueryxFlow does the same as *sqlx.Tx.Queryx but splits the flow.
func (transactionx *Transactionx) QueryxFlow(query string, args ...interface{}) Splitted[*sqlx.Rows, *Transactionx] {
	return SplitOf(transactionx.Queryx(query, args...), transactionx)
}

// Queryx does the same as *sqlx.Tx but is a noop if this flow already failed.
func (transactionx *Transactionx) Queryx(query string, args ...interface{}) *sqlx.Rows {
	return Do(func() (*sqlx.Rows, error) {
		result, err := transactionx.transactionx.Queryx(query, args...)
		return result, err
	}, transactionx, empty[*sqlx.Rows])
}

// QueryxContextFlow does the same as *sqlx.Tx.QueryxContext but splits the flow.
func (transactionx *Transactionx) QueryxContextFlow(ctx context.Context, query string, args ...interface{}) Splitted[*sqlx.Rows, *Transactionx] {
	return SplitOf(transactionx.QueryxContext(ctx, query, args...), transactionx)
}

// QueryxContext does the same as *sqlx.Tx but is a noop if this flow already failed.
func (transactionx *Transactionx) QueryxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Rows {
	return Do(func() (*sqlx.Rows, error) {
		result, err := transactionx.transactionx.QueryxContext(ctx, query, args...)
		return result, err
	}, transactionx, empty[*sqlx.Rows])
}

// RebindFlow does the same as *sqlx.Tx.Rebind but splits the flow.
func (transactionx *Transactionx) RebindFlow(query string) Splitted[string, *Transactionx] {
	return SplitOf(transactionx.Rebind(query), transactionx)
}

// Rebind does the same as *sqlx.Tx.Rebind .
// This is a noop if the flow already failed.
func (transactionx *Transactionx) Rebind(query string) string {
	if transactionx.errs.Err() != nil {
		return empty[string](transactionx.errs.Err())
	}
	result := transactionx.transactionx.Rebind(query)
	return result
}

// Select does the same as *sqlx.Tx but is a noop if this flow already failed.
// Returs this flow
func (transactionx *Transactionx) Select(dest interface{}, query string, args ...interface{}) *Transactionx {
	return pass(func() error { return transactionx.transactionx.Select(dest, query, args...) }, transactionx)
}

// SelectContext does the same as *sqlx.Tx but is a noop if this flow already failed.
// Returs this flow
func (transactionx *Transactionx) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) *Transactionx {
	return pass(func() error { return transactionx.transactionx.SelectContext(ctx, dest, query, args...) }, transactionx)
}

// StmtxFlow does the same as *sqlx.Tx.Stmtx but splits the flow.
func (transactionx *Transactionx) StmtxFlow(stmt interface{}) Splitted[*Stmtx, *Transactionx] {
	return SplitOf(transactionx.Stmtx(stmt), transactionx)
}

// Stmtx does the same as *sqlx.Tx.Stmtx  but wraps the result into another Flow.
// This is a noop if the flow already failed.
func (transactionx *Transactionx) Stmtx(stmt interface{}) *Stmtx {
	if transactionx.errs.Err() != nil {
		return EmptyStmtxOf(transactionx.errs.Err())
	}
	result := transactionx.transactionx.Stmtx(stmt)
	return StmtxOf(result, transactionx)
}

// StmtxContextFlow does the same as *sqlx.Tx.StmtxContext but splits the flow.
func (transactionx *Transactionx) StmtxContextFlow(ctx context.Context, stmt interface{}) Splitted[*Stmtx, *Transactionx] {
	return SplitOf(transactionx.StmtxContext(ctx, stmt), transactionx)
}

// StmtxContext does the same as *sqlx.Tx.StmtxContext  but wraps the result into another Flow.
// This is a noop if the flow already failed.
func (transactionx *Transactionx) StmtxContext(ctx context.Context, stmt interface{}) *Stmtx {
	if transactionx.errs.Err() != nil {
		return EmptyStmtxOf(transactionx.errs.Err())
	}
	result := transactionx.transactionx.StmtxContext(ctx, stmt)
	return StmtxOf(result, transactionx)
}

// UnsafeFlow does the same as *sqlx.Tx.Unsafe but splits the flow.
func (transactionx *Transactionx) UnsafeFlow() Splitted[*Transactionx, *Transactionx] {
	return SplitOf(transactionx.Unsafe(), transactionx)
}

// Unsafe does the same as *sqlx.Tx.Unsafe  but wraps the result into another Flow.
// This is a noop if the flow already failed.
func (transactionx *Transactionx) Unsafe() *Transactionx {
	if transactionx.errs.Err() != nil {
		return EmptyTransactionxOf(transactionx.errs.Err())
	}
	result := transactionx.transactionx.Unsafe()
	return TransactionxOf(result, transactionx)
}

// Commit does the same as *sqlx.Tx but is a noop if this flow already failed.
// Returs this flow
func (transactionx *Transactionx) Commit() *Transactionx {
	return pass(func() error { return transactionx.transactionx.Commit() }, transactionx)
}

// ExecFlow does the same as *sqlx.Tx.Exec but splits the flow.
func (transactionx *Transactionx) ExecFlow(query string, args ...any) Splitted[*Result, *Transactionx] {
	return SplitOf(transactionx.Exec(query, args...), transactionx)
}

// Exec does the same as *sqlx.Tx but is a noop if this flow already failed.
func (transactionx *Transactionx) Exec(query string, args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := transactionx.transactionx.Exec(query, args...)
		return ResultOf(result, transactionx), err
	}, transactionx, EmptyResultOf)
}

// ExecContextFlow does the same as *sqlx.Tx.ExecContext but splits the flow.
func (transactionx *Transactionx) ExecContextFlow(ctx context.Context, query string, args ...any) Splitted[*Result, *Transactionx] {
	return SplitOf(transactionx.ExecContext(ctx, query, args...), transactionx)
}

// ExecContext does the same as *sqlx.Tx but is a noop if this flow already failed.
func (transactionx *Transactionx) ExecContext(ctx context.Context, query string, args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := transactionx.transactionx.ExecContext(ctx, query, args...)
		return ResultOf(result, transactionx), err
	}, transactionx, EmptyResultOf)
}

// PrepareFlow does the same as *sqlx.Tx.Prepare but splits the flow.
func (transactionx *Transactionx) PrepareFlow(query string) Splitted[*Stmt, *Transactionx] {
	return SplitOf(transactionx.Prepare(query), transactionx)
}

// Prepare does the same as *sqlx.Tx but is a noop if this flow already failed.
func (transactionx *Transactionx) Prepare(query string) *Stmt {
	return Do(func() (*Stmt, error) {
		result, err := transactionx.transactionx.Prepare(query)
		return StmtOf(result, transactionx), err
	}, transactionx, EmptyStmtOf)
}

// PrepareContextFlow does the same as *sqlx.Tx.PrepareContext but splits the flow.
func (transactionx *Transactionx) PrepareContextFlow(ctx context.Context, query string) Splitted[*Stmt, *Transactionx] {
	return SplitOf(transactionx.PrepareContext(ctx, query), transactionx)
}

// PrepareContext does the same as *sqlx.Tx but is a noop if this flow already failed.
func (transactionx *Transactionx) PrepareContext(ctx context.Context, query string) *Stmt {
	return Do(func() (*Stmt, error) {
		result, err := transactionx.transactionx.PrepareContext(ctx, query)
		return StmtOf(result, transactionx), err
	}, transactionx, EmptyStmtOf)
}

// QueryFlow does the same as *sqlx.Tx.Query but splits the flow.
func (transactionx *Transactionx) QueryFlow(query string, args ...any) Splitted[*sql.Rows, *Transactionx] {
	return SplitOf(transactionx.Query(query, args...), transactionx)
}

// Query does the same as *sqlx.Tx but is a noop if this flow already failed.
func (transactionx *Transactionx) Query(query string, args ...any) *sql.Rows {
	return Do(func() (*sql.Rows, error) {
		result, err := transactionx.transactionx.Query(query, args...)
		return result, err
	}, transactionx, empty[*sql.Rows])
}

// QueryContextFlow does the same as *sqlx.Tx.QueryContext but splits the flow.
func (transactionx *Transactionx) QueryContextFlow(ctx context.Context, query string, args ...any) Splitted[*sql.Rows, *Transactionx] {
	return SplitOf(transactionx.QueryContext(ctx, query, args...), transactionx)
}

// QueryContext does the same as *sqlx.Tx but is a noop if this flow already failed.
func (transactionx *Transactionx) QueryContext(ctx context.Context, query string, args ...any) *sql.Rows {
	return Do(func() (*sql.Rows, error) {
		result, err := transactionx.transactionx.QueryContext(ctx, query, args...)
		return result, err
	}, transactionx, empty[*sql.Rows])
}

// QueryRowFlow does the same as *sqlx.Tx.QueryRow but splits the flow.
func (transactionx *Transactionx) QueryRowFlow(query string, args ...any) Splitted[*sql.Row, *Transactionx] {
	return SplitOf(transactionx.QueryRow(query, args...), transactionx)
}

// QueryRow does the same as *sqlx.Tx.QueryRow .
// This is a noop if the flow already failed.
func (transactionx *Transactionx) QueryRow(query string, args ...any) *sql.Row {
	if transactionx.errs.Err() != nil {
		return empty[*sql.Row](transactionx.errs.Err())
	}
	result := transactionx.transactionx.QueryRow(query, args...)
	return result
}

// QueryRowContextFlow does the same as *sqlx.Tx.QueryRowContext but splits the flow.
func (transactionx *Transactionx) QueryRowContextFlow(ctx context.Context, query string, args ...any) Splitted[*sql.Row, *Transactionx] {
	return SplitOf(transactionx.QueryRowContext(ctx, query, args...), transactionx)
}

// QueryRowContext does the same as *sqlx.Tx.QueryRowContext .
// This is a noop if the flow already failed.
func (transactionx *Transactionx) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	if transactionx.errs.Err() != nil {
		return empty[*sql.Row](transactionx.errs.Err())
	}
	result := transactionx.transactionx.QueryRowContext(ctx, query, args...)
	return result
}

// Rollback does the same as *sqlx.Tx but is a noop if this flow already failed.
// Returs this flow
func (transactionx *Transactionx) Rollback() *Transactionx {
	return pass(func() error { return transactionx.transactionx.Rollback() }, transactionx)
}

// StmtFlow does the same as *sqlx.Tx.Stmt but splits the flow.
func (transactionx *Transactionx) StmtFlow(stmt *sql.Stmt) Splitted[*Stmt, *Transactionx] {
	return SplitOf(transactionx.Stmt(stmt), transactionx)
}

// Stmt does the same as *sqlx.Tx.Stmt  but wraps the result into another Flow.
// This is a noop if the flow already failed.
func (transactionx *Transactionx) Stmt(stmt *sql.Stmt) *Stmt {
	if transactionx.errs.Err() != nil {
		return EmptyStmtOf(transactionx.errs.Err())
	}
	result := transactionx.transactionx.Stmt(stmt)
	return StmtOf(result, transactionx)
}

// StmtContextFlow does the same as *sqlx.Tx.StmtContext but splits the flow.
func (transactionx *Transactionx) StmtContextFlow(ctx context.Context, stmt *sql.Stmt) Splitted[*Stmt, *Transactionx] {
	return SplitOf(transactionx.StmtContext(ctx, stmt), transactionx)
}

// StmtContext does the same as *sqlx.Tx.StmtContext  but wraps the result into another Flow.
// This is a noop if the flow already failed.
func (transactionx *Transactionx) StmtContext(ctx context.Context, stmt *sql.Stmt) *Stmt {
	if transactionx.errs.Err() != nil {
		return EmptyStmtOf(transactionx.errs.Err())
	}
	result := transactionx.transactionx.StmtContext(ctx, stmt)
	return StmtOf(result, transactionx)
}

// Err returns the error of this flow if any happend.
func (transactionx *Transactionx) Err() error {
	return transactionx.errs.Err()
}

// Fail ends this flow with err
func (transactionx *Transactionx) Fail(err error) {
	transactionx.errs.Fail(err)
}

// Link returns the base error of this flow.
func (transactionx *Transactionx) Link() *error {
	return transactionx.errs.Link()
}

// LinkTo merges err as base into this flow.
func (transactionx *Transactionx) LinkTo(err *error) {
	transactionx.errs.LinkTo(err)
}

// Unwrap returns the internal value and the produced error if any.
func (transactionx *Transactionx) Unwrap() (*sqlx.Tx, error) {
	return transactionx.transactionx, transactionx.Err()
}

// Raw returns the internal value.
// No Guarantees.
func (transactionx *Transactionx) Raw() *sqlx.Tx {
	return transactionx.transactionx
}

// TransactionxOf create a new Transactionx and is linked to flow.
// They are linked to gether and if one errors the error is propagated.
func TransactionxOf(transactionx *sqlx.Tx, flow Linkable) *Transactionx {
	return &Transactionx{transactionx: transactionx, errs: errChainOf(flow)}
}

// NewTransactionx create a new Transactionx and is the root of a flow.
// It will catch any error that happens in the future.
func NewTransactionx(transactionx *sqlx.Tx) *Transactionx {
	return &Transactionx{transactionx: transactionx, errs: emptyChain()}
}

// EmptyTransactionxOf returns an already failed Transactionx.
// Calls will have no effects on it.
func EmptyTransactionxOf(err error) *Transactionx {
	return &Transactionx{errs: errChainOfErr(err)}
}
