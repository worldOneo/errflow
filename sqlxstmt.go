package errflow

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Stmtx is the flow wrapper of *sqlx.Stmt
type Stmtx struct {
	stmtx *sqlx.Stmt
	errs  errChain
}

// Get does the same as *sqlx.Stmt but is a noop if this flow already failed.
// Returs this flow
func (stmtx *Stmtx) Get(dest interface{}, args ...interface{}) *Stmtx {
	return pass(func() error { return stmtx.stmtx.Get(dest, args...) }, stmtx)
}

// GetContext does the same as *sqlx.Stmt but is a noop if this flow already failed.
// Returs this flow
func (stmtx *Stmtx) GetContext(ctx context.Context, dest interface{}, args ...interface{}) *Stmtx {
	return pass(func() error { return stmtx.stmtx.GetContext(ctx, dest, args...) }, stmtx)
}

// QueryRowxFlow does the same as *sqlx.Stmt.QueryRowx but splits the flow.
func (stmtx *Stmtx) QueryRowxFlow(args ...interface{}) Splitted[*sqlx.Row, *Stmtx] {
	return SplitOf(stmtx.QueryRowx(args...), stmtx)
}

// QueryRowx does the same as *sqlx.Stmt.QueryRowx .
// This is a noop if the flow already failed.
func (stmtx *Stmtx) QueryRowx(args ...interface{}) *sqlx.Row {
	if stmtx.errs.Err() != nil {
		return empty[*sqlx.Row](stmtx.errs.Err())
	}
	result := stmtx.stmtx.QueryRowx(args...)
	return result
}

// QueryRowxContextFlow does the same as *sqlx.Stmt.QueryRowxContext but splits the flow.
func (stmtx *Stmtx) QueryRowxContextFlow(ctx context.Context, args ...interface{}) Splitted[*sqlx.Row, *Stmtx] {
	return SplitOf(stmtx.QueryRowxContext(ctx, args...), stmtx)
}

// QueryRowxContext does the same as *sqlx.Stmt.QueryRowxContext .
// This is a noop if the flow already failed.
func (stmtx *Stmtx) QueryRowxContext(ctx context.Context, args ...interface{}) *sqlx.Row {
	if stmtx.errs.Err() != nil {
		return empty[*sqlx.Row](stmtx.errs.Err())
	}
	result := stmtx.stmtx.QueryRowxContext(ctx, args...)
	return result
}

// QueryxFlow does the same as *sqlx.Stmt.Queryx but splits the flow.
func (stmtx *Stmtx) QueryxFlow(args ...interface{}) Splitted[*sqlx.Rows, *Stmtx] {
	return SplitOf(stmtx.Queryx(args...), stmtx)
}

// Queryx does the same as *sqlx.Stmt but is a noop if this flow already failed.
func (stmtx *Stmtx) Queryx(args ...interface{}) *sqlx.Rows {
	return Do(func() (*sqlx.Rows, error) {
		result, err := stmtx.stmtx.Queryx(args...)
		return result, err
	}, stmtx, empty[*sqlx.Rows])
}

// QueryxContextFlow does the same as *sqlx.Stmt.QueryxContext but splits the flow.
func (stmtx *Stmtx) QueryxContextFlow(ctx context.Context, args ...interface{}) Splitted[*sqlx.Rows, *Stmtx] {
	return SplitOf(stmtx.QueryxContext(ctx, args...), stmtx)
}

// QueryxContext does the same as *sqlx.Stmt but is a noop if this flow already failed.
func (stmtx *Stmtx) QueryxContext(ctx context.Context, args ...interface{}) *sqlx.Rows {
	return Do(func() (*sqlx.Rows, error) {
		result, err := stmtx.stmtx.QueryxContext(ctx, args...)
		return result, err
	}, stmtx, empty[*sqlx.Rows])
}

// Select does the same as *sqlx.Stmt but is a noop if this flow already failed.
// Returs this flow
func (stmtx *Stmtx) Select(dest interface{}, args ...interface{}) *Stmtx {
	return pass(func() error { return stmtx.stmtx.Select(dest, args...) }, stmtx)
}

// SelectContext does the same as *sqlx.Stmt but is a noop if this flow already failed.
// Returs this flow
func (stmtx *Stmtx) SelectContext(ctx context.Context, dest interface{}, args ...interface{}) *Stmtx {
	return pass(func() error { return stmtx.stmtx.SelectContext(ctx, dest, args...) }, stmtx)
}

// UnsafeFlow does the same as *sqlx.Stmt.Unsafe but splits the flow.
func (stmtx *Stmtx) UnsafeFlow() Splitted[*Stmtx, *Stmtx] {
	return SplitOf(stmtx.Unsafe(), stmtx)
}

// Unsafe does the same as *sqlx.Stmt.Unsafe  but wraps the result into another Flow.
// This is a noop if the flow already failed.
func (stmtx *Stmtx) Unsafe() *Stmtx {
	if stmtx.errs.Err() != nil {
		return EmptyStmtxOf(stmtx.errs.Err())
	}
	result := stmtx.stmtx.Unsafe()
	return StmtxOf(result, stmtx)
}

// Close does the same as *sqlx.Stmt but is a noop if this flow already failed.
// Returs this flow
func (stmtx *Stmtx) Close() *Stmtx {
	return pass(func() error { return stmtx.stmtx.Close() }, stmtx)
}

// ExecFlow does the same as *sqlx.Stmt.Exec but splits the flow.
func (stmtx *Stmtx) ExecFlow(args ...any) Splitted[*Result, *Stmtx] {
	return SplitOf(stmtx.Exec(args...), stmtx)
}

// Exec does the same as *sqlx.Stmt but is a noop if this flow already failed.
func (stmtx *Stmtx) Exec(args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := stmtx.stmtx.Exec(args...)
		return ResultOf(result, stmtx), err
	}, stmtx, EmptyResultOf)
}

// ExecContextFlow does the same as *sqlx.Stmt.ExecContext but splits the flow.
func (stmtx *Stmtx) ExecContextFlow(ctx context.Context, args ...any) Splitted[*Result, *Stmtx] {
	return SplitOf(stmtx.ExecContext(ctx, args...), stmtx)
}

// ExecContext does the same as *sqlx.Stmt but is a noop if this flow already failed.
func (stmtx *Stmtx) ExecContext(ctx context.Context, args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := stmtx.stmtx.ExecContext(ctx, args...)
		return ResultOf(result, stmtx), err
	}, stmtx, EmptyResultOf)
}

// QueryFlow does the same as *sqlx.Stmt.Query but splits the flow.
func (stmtx *Stmtx) QueryFlow(args ...any) Splitted[*sql.Rows, *Stmtx] {
	return SplitOf(stmtx.Query(args...), stmtx)
}

// Query does the same as *sqlx.Stmt but is a noop if this flow already failed.
func (stmtx *Stmtx) Query(args ...any) *sql.Rows {
	return Do(func() (*sql.Rows, error) {
		result, err := stmtx.stmtx.Query(args...)
		return result, err
	}, stmtx, empty[*sql.Rows])
}

// QueryContextFlow does the same as *sqlx.Stmt.QueryContext but splits the flow.
func (stmtx *Stmtx) QueryContextFlow(ctx context.Context, args ...any) Splitted[*sql.Rows, *Stmtx] {
	return SplitOf(stmtx.QueryContext(ctx, args...), stmtx)
}

// QueryContext does the same as *sqlx.Stmt but is a noop if this flow already failed.
func (stmtx *Stmtx) QueryContext(ctx context.Context, args ...any) *sql.Rows {
	return Do(func() (*sql.Rows, error) {
		result, err := stmtx.stmtx.QueryContext(ctx, args...)
		return result, err
	}, stmtx, empty[*sql.Rows])
}

// QueryRowFlow does the same as *sqlx.Stmt.QueryRow but splits the flow.
func (stmtx *Stmtx) QueryRowFlow(args ...any) Splitted[*sql.Row, *Stmtx] {
	return SplitOf(stmtx.QueryRow(args...), stmtx)
}

// QueryRow does the same as *sqlx.Stmt.QueryRow .
// This is a noop if the flow already failed.
func (stmtx *Stmtx) QueryRow(args ...any) *sql.Row {
	if stmtx.errs.Err() != nil {
		return empty[*sql.Row](stmtx.errs.Err())
	}
	result := stmtx.stmtx.QueryRow(args...)
	return result
}

// QueryRowContextFlow does the same as *sqlx.Stmt.QueryRowContext but splits the flow.
func (stmtx *Stmtx) QueryRowContextFlow(ctx context.Context, args ...any) Splitted[*sql.Row, *Stmtx] {
	return SplitOf(stmtx.QueryRowContext(ctx, args...), stmtx)
}

// QueryRowContext does the same as *sqlx.Stmt.QueryRowContext .
// This is a noop if the flow already failed.
func (stmtx *Stmtx) QueryRowContext(ctx context.Context, args ...any) *sql.Row {
	if stmtx.errs.Err() != nil {
		return empty[*sql.Row](stmtx.errs.Err())
	}
	result := stmtx.stmtx.QueryRowContext(ctx, args...)
	return result
}

// Err returns the error of this flow if any happend.
func (stmtx *Stmtx) Err() error {
	return stmtx.errs.Err()
}

// Fail ends this flow with err
func (stmtx *Stmtx) Fail(err error) {
	stmtx.errs.Fail(err)
}

// Link returns the base error of this flow.
func (stmtx *Stmtx) Link() *error {
	return stmtx.errs.Link()
}

// LinkTo merges err as base into this flow.
func (stmtx *Stmtx) LinkTo(err *error) {
	stmtx.errs.LinkTo(err)
}

// Unwrap returns the internal value and the produced error if any.
func (stmtx *Stmtx) Unwrap() (*sqlx.Stmt, error) {
	return stmtx.stmtx, stmtx.Err()
}

// Raw returns the internal value.
// No Guarantees.
func (stmtx *Stmtx) Raw() *sqlx.Stmt {
	return stmtx.stmtx
}

// StmtxOf create a new Stmtx and is linked to flow.
// They are linked to gether and if one errors the error is propagated.
func StmtxOf(stmtx *sqlx.Stmt, flow Linkable) *Stmtx {
	return &Stmtx{stmtx: stmtx, errs: errChainOf(flow)}
}

// NewStmtx create a new Stmtx and is the root of a flow.
// It will catch any error that happens in the future.
func NewStmtx(stmtx *sqlx.Stmt) *Stmtx {
	return &Stmtx{stmtx: stmtx, errs: emptyChain()}
}

// EmptyStmtxOf returns an already failed Stmtx.
// Calls will have no effects on it.
func EmptyStmtxOf(err error) *Stmtx {
	return &Stmtx{errs: errChainOfErr(err)}
}
