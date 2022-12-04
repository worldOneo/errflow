package errflow

import (
	"context"
	"database/sql"
)

// Stmt is the flow wrapper of *sql.Stmt
type Stmt struct {
	stmt *sql.Stmt
	errs errChain
}

// Close does the same as *sql.Stmt but is a noop if this flow already failed.
// Returs this flow
func (stmt *Stmt) Close() *Stmt {
	return pass(func() error { return stmt.stmt.Close() }, stmt)
}

// ExecFlow does the same as *sql.Stmt.Exec but splits the flow.
func (stmt *Stmt) ExecFlow(args ...any) Splitted[*Result, *Stmt] {
	return SplitOf(stmt.Exec(args...), stmt)
}

// Exec does the same as *sql.Stmt but is a noop if this flow already failed.
func (stmt *Stmt) Exec(args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := stmt.stmt.Exec(args...)
		return ResultOf(result, stmt), err
	}, stmt, EmptyResultOf)
}

// ExecContextFlow does the same as *sql.Stmt.ExecContext but splits the flow.
func (stmt *Stmt) ExecContextFlow(ctx context.Context, args ...any) Splitted[*Result, *Stmt] {
	return SplitOf(stmt.ExecContext(ctx, args...), stmt)
}

// ExecContext does the same as *sql.Stmt but is a noop if this flow already failed.
func (stmt *Stmt) ExecContext(ctx context.Context, args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := stmt.stmt.ExecContext(ctx, args...)
		return ResultOf(result, stmt), err
	}, stmt, EmptyResultOf)
}

// QueryFlow does the same as *sql.Stmt.Query but splits the flow.
func (stmt *Stmt) QueryFlow(args ...any) Splitted[*sql.Rows, *Stmt] {
	return SplitOf(stmt.Query(args...), stmt)
}

// Query does the same as *sql.Stmt but is a noop if this flow already failed.
func (stmt *Stmt) Query(args ...any) *sql.Rows {
	return Do(func() (*sql.Rows, error) {
		result, err := stmt.stmt.Query(args...)
		return result, err
	}, stmt, empty[*sql.Rows])
}

// QueryContextFlow does the same as *sql.Stmt.QueryContext but splits the flow.
func (stmt *Stmt) QueryContextFlow(ctx context.Context, args ...any) Splitted[*sql.Rows, *Stmt] {
	return SplitOf(stmt.QueryContext(ctx, args...), stmt)
}

// QueryContext does the same as *sql.Stmt but is a noop if this flow already failed.
func (stmt *Stmt) QueryContext(ctx context.Context, args ...any) *sql.Rows {
	return Do(func() (*sql.Rows, error) {
		result, err := stmt.stmt.QueryContext(ctx, args...)
		return result, err
	}, stmt, empty[*sql.Rows])
}

// QueryRowFlow does the same as *sql.Stmt.QueryRow but splits the flow.
func (stmt *Stmt) QueryRowFlow(args ...any) Splitted[*sql.Row, *Stmt] {
	return SplitOf(stmt.QueryRow(args...), stmt)
}

// QueryRow does the same as *sql.Stmt.QueryRow .
// This is a noop if the flow already failed.
func (stmt *Stmt) QueryRow(args ...any) *sql.Row {
	if stmt.errs.Err() != nil {
		return empty[*sql.Row](stmt.errs.Err())
	}
	result := stmt.stmt.QueryRow(args...)
	return result
}

// QueryRowContextFlow does the same as *sql.Stmt.QueryRowContext but splits the flow.
func (stmt *Stmt) QueryRowContextFlow(ctx context.Context, args ...any) Splitted[*sql.Row, *Stmt] {
	return SplitOf(stmt.QueryRowContext(ctx, args...), stmt)
}

// QueryRowContext does the same as *sql.Stmt.QueryRowContext .
// This is a noop if the flow already failed.
func (stmt *Stmt) QueryRowContext(ctx context.Context, args ...any) *sql.Row {
	if stmt.errs.Err() != nil {
		return empty[*sql.Row](stmt.errs.Err())
	}
	result := stmt.stmt.QueryRowContext(ctx, args...)
	return result
}

// Err returns the error of this flow if any happend.
func (stmt *Stmt) Err() error {
	return stmt.errs.Err()
}

// Fail ends this flow with err
func (stmt *Stmt) Fail(err error) {
	stmt.errs.Fail(err)
}

// Link returns the base error of this flow.
func (stmt *Stmt) Link() *error {
	return stmt.errs.Link()
}

// LinkTo merges err as base into this flow.
func (stmt *Stmt) LinkTo(err *error) {
	stmt.errs.LinkTo(err)
}

// Unwrap returns the internal value and the produced error if any.
func (stmt *Stmt) Unwrap() (*sql.Stmt, error) {
	return stmt.stmt, stmt.Err()
}

// Raw returns the internal value.
// No Guarantees.
func (stmt *Stmt) Raw() *sql.Stmt {
	return stmt.stmt
}

// StmtOf create a new Stmt and is linked to flow.
// They are linked to gether and if one errors the error is propagated.
func StmtOf(stmt *sql.Stmt, flow Linkable) *Stmt {
	return &Stmt{stmt: stmt, errs: errChainOf(flow)}
}

// NewStmt create a new Stmt and is the root of a flow.
// It will catch any error that happens in the future.
func NewStmt(stmt *sql.Stmt) *Stmt {
	return &Stmt{stmt: stmt, errs: emptyChain()}
}

// EmptyStmtOf returns an already failed Stmt.
// Calls will have no effects on it.
func EmptyStmtOf(err error) *Stmt {
	return &Stmt{errs: errChainOfErr(err)}
}
