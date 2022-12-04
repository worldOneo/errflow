package errflow

import (
	"context"
	"database/sql"
)

type Stmt struct {
	stmt *sql.Stmt
	errs errChain
}

func (stmt *Stmt) ExecFlow(args ...any) Splitted[*Result, *Stmt] {
	return SplitOf(stmt.Exec(args...), stmt)
}

func (stmt *Stmt) Exec(args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := stmt.stmt.Exec(args...)
		return ResultOf(result, stmt), err
	}, stmt, EmptyResultOf)
}

func (stmt *Stmt) ExecContextFlow(ctx context.Context, args ...any) Splitted[*Result, *Stmt] {
	return SplitOf(stmt.ExecContext(ctx, args...), stmt)
}

func (stmt *Stmt) ExecContext(ctx context.Context, args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := stmt.stmt.ExecContext(ctx, args...)
		return ResultOf(result, stmt), err
	}, stmt, EmptyResultOf)
}

func (stmt *Stmt) Err() error {
	return stmt.errs.Err()
}

func (stmt *Stmt) Fail(err error) {
	stmt.errs.Fail(err)
}

func (stmt *Stmt) Link() *error {
	return stmt.errs.Link()
}

func (stmt *Stmt) LinkTo(err *error) {
	stmt.errs.LinkTo(err)
}

func (stmt *Stmt) Unwrap() (*sql.Stmt, error) {
	return stmt.stmt, stmt.Err()
}

func (stmt *Stmt) Raw() *sql.Stmt {
	return stmt.stmt
}

func StmtOf(stmt *sql.Stmt, flow Linkable) *Stmt {
	return &Stmt{stmt: stmt, errs: errChainOf(flow)}
}

func NewStmt(stmt *sql.Stmt) *Stmt {
	return &Stmt{stmt: stmt, errs: emptyChain()}
}

func EmptyStmtOf(err error) *Stmt {
	return &Stmt{errs: errChainOfErr(err)}
}
