package errflow

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Stmtx struct {
	stmt *sqlx.Stmt
	errs errChain
}

func (stmt *Stmtx) ExecFlow(args ...any) Splitted[*Result, *Stmtx] {
	return SplitOf(stmt.Exec(args...), stmt)
}

func (stmt *Stmtx) Exec(args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := stmt.stmt.Exec(args...)
		return ResultOf(result, stmt), err
	}, stmt, EmptyResultOf)
}

func (stmt *Stmtx) ExecContextFlow(ctx context.Context, args ...any) Splitted[*Result, *Stmtx] {
	return SplitOf(stmt.ExecContext(ctx, args...), stmt)
}

func (stmt *Stmtx) ExecContext(ctx context.Context, args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := stmt.stmt.ExecContext(ctx, args...)
		return ResultOf(result, stmt), err
	}, stmt, EmptyResultOf)
}
func (stmt *Stmtx) Get(dest any, args ...any) *Stmtx {
	return pass(func() error { return stmt.stmt.Get(dest, args...) }, stmt)
}

func (stmt *Stmtx) GetContext(ctx context.Context, dest any, args ...any) *Stmtx {
	return pass(func() error { return stmt.stmt.GetContext(ctx, dest, args...) }, stmt)
}

func (stmt *Stmtx) Select(dest any, args ...any) *Stmtx {
	return pass(func() error { return stmt.stmt.Select(dest, args...) }, stmt)
}

func (stmt *Stmtx) SelectContext(ctx context.Context, dest any, args ...any) *Stmtx {
	return pass(func() error { return stmt.stmt.GetContext(ctx, dest, args...) }, stmt)
}

func (stmt *Stmtx) Err() error {
	return stmt.errs.Err()
}

func (stmt *Stmtx) Fail(err error) {
	stmt.errs.Fail(err)
}

func (stmt *Stmtx) Link() *error {
	return stmt.errs.Link()
}

func (stmt *Stmtx) LinkTo(err *error) {
	stmt.errs.LinkTo(err)
}

func (stmt *Stmtx) Unwrap() (*sqlx.Stmt, error) {
	return stmt.stmt, stmt.Err()
}

func (stmt *Stmtx) Raw() *sqlx.Stmt {
	return stmt.stmt
}

func StmtxOf(stmt *sqlx.Stmt, flow Linkable) *Stmtx {
	return &Stmtx{stmt: stmt, errs: errChainOf(flow)}
}

func NewStmtxOf(stmt *sqlx.Stmt) *Stmtx {
	return &Stmtx{stmt: stmt, errs: emptyChain()}
}

func EmptyStmtxOf(err error) *Stmtx {
	return &Stmtx{errs: errChainOfErr(err)}
}
