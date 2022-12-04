package errflow

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Transactionx struct {
	transactionx *sqlx.Tx
	errs         errChain
}

func (transactionx *Transactionx) DriverNameFlow() Splitted[string, *Transactionx] {
	return SplitOf(transactionx.DriverName(), transactionx)
}

func (transactionx *Transactionx) DriverName() string {
	result := transactionx.transactionx.DriverName()
	return result
}

func (transactionx *Transactionx) Get(dest interface{}, query string, args ...interface{}) *Transactionx {
	return pass(func() error { return transactionx.transactionx.Get(dest, query, args...) }, transactionx)
}

func (transactionx *Transactionx) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) *Transactionx {
	return pass(func() error { return transactionx.transactionx.GetContext(ctx, dest, query, args...) }, transactionx)
}

func (transactionx *Transactionx) NamedExecFlow(query string, arg interface{}) Splitted[*Result, *Transactionx] {
	return SplitOf(transactionx.NamedExec(query, arg), transactionx)
}

func (transactionx *Transactionx) NamedExec(query string, arg interface{}) *Result {
	return Do(func() (*Result, error) {
		result, err := transactionx.transactionx.NamedExec(query, arg)
		return ResultOf(result, transactionx), err
	}, transactionx, EmptyResultOf)
}

func (transactionx *Transactionx) NamedExecContextFlow(ctx context.Context, query string, arg interface{}) Splitted[*Result, *Transactionx] {
	return SplitOf(transactionx.NamedExecContext(ctx, query, arg), transactionx)
}

func (transactionx *Transactionx) NamedExecContext(ctx context.Context, query string, arg interface{}) *Result {
	return Do(func() (*Result, error) {
		result, err := transactionx.transactionx.NamedExecContext(ctx, query, arg)
		return ResultOf(result, transactionx), err
	}, transactionx, EmptyResultOf)
}

func (transactionx *Transactionx) NamedQueryFlow(query string, arg interface{}) Splitted[*sqlx.Rows, *Transactionx] {
	return SplitOf(transactionx.NamedQuery(query, arg), transactionx)
}

func (transactionx *Transactionx) NamedQuery(query string, arg interface{}) *sqlx.Rows {
	return Do(func() (*sqlx.Rows, error) {
		result, err := transactionx.transactionx.NamedQuery(query, arg)
		return result, err
	}, transactionx, empty[*sqlx.Rows])
}

func (transactionx *Transactionx) NamedStmtFlow(stmt *sqlx.NamedStmt) Splitted[*sqlx.NamedStmt, *Transactionx] {
	return SplitOf(transactionx.NamedStmt(stmt), transactionx)
}

func (transactionx *Transactionx) NamedStmt(stmt *sqlx.NamedStmt) *sqlx.NamedStmt {
	result := transactionx.transactionx.NamedStmt(stmt)
	return result
}

func (transactionx *Transactionx) NamedStmtContextFlow(ctx context.Context, stmt *sqlx.NamedStmt) Splitted[*sqlx.NamedStmt, *Transactionx] {
	return SplitOf(transactionx.NamedStmtContext(ctx, stmt), transactionx)
}

func (transactionx *Transactionx) NamedStmtContext(ctx context.Context, stmt *sqlx.NamedStmt) *sqlx.NamedStmt {
	result := transactionx.transactionx.NamedStmtContext(ctx, stmt)
	return result
}

func (transactionx *Transactionx) PrepareNamedFlow(query string) Splitted[*sqlx.NamedStmt, *Transactionx] {
	return SplitOf(transactionx.PrepareNamed(query), transactionx)
}

func (transactionx *Transactionx) PrepareNamed(query string) *sqlx.NamedStmt {
	return Do(func() (*sqlx.NamedStmt, error) {
		result, err := transactionx.transactionx.PrepareNamed(query)
		return result, err
	}, transactionx, empty[*sqlx.NamedStmt])
}

func (transactionx *Transactionx) PrepareNamedContextFlow(ctx context.Context, query string) Splitted[*sqlx.NamedStmt, *Transactionx] {
	return SplitOf(transactionx.PrepareNamedContext(ctx, query), transactionx)
}

func (transactionx *Transactionx) PrepareNamedContext(ctx context.Context, query string) *sqlx.NamedStmt {
	return Do(func() (*sqlx.NamedStmt, error) {
		result, err := transactionx.transactionx.PrepareNamedContext(ctx, query)
		return result, err
	}, transactionx, empty[*sqlx.NamedStmt])
}

func (transactionx *Transactionx) PreparexFlow(query string) Splitted[*Stmtx, *Transactionx] {
	return SplitOf(transactionx.Preparex(query), transactionx)
}

func (transactionx *Transactionx) Preparex(query string) *Stmtx {
	return Do(func() (*Stmtx, error) {
		result, err := transactionx.transactionx.Preparex(query)
		return StmtxOf(result, transactionx), err
	}, transactionx, EmptyStmtxOf)
}

func (transactionx *Transactionx) PreparexContextFlow(ctx context.Context, query string) Splitted[*Stmtx, *Transactionx] {
	return SplitOf(transactionx.PreparexContext(ctx, query), transactionx)
}

func (transactionx *Transactionx) PreparexContext(ctx context.Context, query string) *Stmtx {
	return Do(func() (*Stmtx, error) {
		result, err := transactionx.transactionx.PreparexContext(ctx, query)
		return StmtxOf(result, transactionx), err
	}, transactionx, EmptyStmtxOf)
}

func (transactionx *Transactionx) QueryRowxFlow(query string, args ...interface{}) Splitted[*sqlx.Row, *Transactionx] {
	return SplitOf(transactionx.QueryRowx(query, args...), transactionx)
}

func (transactionx *Transactionx) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	result := transactionx.transactionx.QueryRowx(query, args...)
	return result
}

func (transactionx *Transactionx) QueryRowxContextFlow(ctx context.Context, query string, args ...interface{}) Splitted[*sqlx.Row, *Transactionx] {
	return SplitOf(transactionx.QueryRowxContext(ctx, query, args...), transactionx)
}

func (transactionx *Transactionx) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	result := transactionx.transactionx.QueryRowxContext(ctx, query, args...)
	return result
}

func (transactionx *Transactionx) QueryxFlow(query string, args ...interface{}) Splitted[*sqlx.Rows, *Transactionx] {
	return SplitOf(transactionx.Queryx(query, args...), transactionx)
}

func (transactionx *Transactionx) Queryx(query string, args ...interface{}) *sqlx.Rows {
	return Do(func() (*sqlx.Rows, error) {
		result, err := transactionx.transactionx.Queryx(query, args...)
		return result, err
	}, transactionx, empty[*sqlx.Rows])
}

func (transactionx *Transactionx) QueryxContextFlow(ctx context.Context, query string, args ...interface{}) Splitted[*sqlx.Rows, *Transactionx] {
	return SplitOf(transactionx.QueryxContext(ctx, query, args...), transactionx)
}

func (transactionx *Transactionx) QueryxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Rows {
	return Do(func() (*sqlx.Rows, error) {
		result, err := transactionx.transactionx.QueryxContext(ctx, query, args...)
		return result, err
	}, transactionx, empty[*sqlx.Rows])
}

func (transactionx *Transactionx) RebindFlow(query string) Splitted[string, *Transactionx] {
	return SplitOf(transactionx.Rebind(query), transactionx)
}

func (transactionx *Transactionx) Rebind(query string) string {
	result := transactionx.transactionx.Rebind(query)
	return result
}

func (transactionx *Transactionx) Select(dest interface{}, query string, args ...interface{}) *Transactionx {
	return pass(func() error { return transactionx.transactionx.Select(dest, query, args...) }, transactionx)
}

func (transactionx *Transactionx) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) *Transactionx {
	return pass(func() error { return transactionx.transactionx.SelectContext(ctx, dest, query, args...) }, transactionx)
}

func (transactionx *Transactionx) StmtxFlow(stmt interface{}) Splitted[*Stmtx, *Transactionx] {
	return SplitOf(transactionx.Stmtx(stmt), transactionx)
}

func (transactionx *Transactionx) Stmtx(stmt interface{}) *Stmtx {
	result := transactionx.transactionx.Stmtx(stmt)
	return StmtxOf(result, transactionx)
}

func (transactionx *Transactionx) StmtxContextFlow(ctx context.Context, stmt interface{}) Splitted[*Stmtx, *Transactionx] {
	return SplitOf(transactionx.StmtxContext(ctx, stmt), transactionx)
}

func (transactionx *Transactionx) StmtxContext(ctx context.Context, stmt interface{}) *Stmtx {
	result := transactionx.transactionx.StmtxContext(ctx, stmt)
	return StmtxOf(result, transactionx)
}

func (transactionx *Transactionx) UnsafeFlow() Splitted[*Transactionx, *Transactionx] {
	return SplitOf(transactionx.Unsafe(), transactionx)
}

func (transactionx *Transactionx) Unsafe() *Transactionx {
	result := transactionx.transactionx.Unsafe()
	return TransactionxOf(result, transactionx)
}

func (transactionx *Transactionx) Commit() *Transactionx {
	return pass(func() error { return transactionx.transactionx.Commit() }, transactionx)
}

func (transactionx *Transactionx) ExecFlow(query string, args ...any) Splitted[*Result, *Transactionx] {
	return SplitOf(transactionx.Exec(query, args...), transactionx)
}

func (transactionx *Transactionx) Exec(query string, args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := transactionx.transactionx.Exec(query, args...)
		return ResultOf(result, transactionx), err
	}, transactionx, EmptyResultOf)
}

func (transactionx *Transactionx) ExecContextFlow(ctx context.Context, query string, args ...any) Splitted[*Result, *Transactionx] {
	return SplitOf(transactionx.ExecContext(ctx, query, args...), transactionx)
}

func (transactionx *Transactionx) ExecContext(ctx context.Context, query string, args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := transactionx.transactionx.ExecContext(ctx, query, args...)
		return ResultOf(result, transactionx), err
	}, transactionx, EmptyResultOf)
}

func (transactionx *Transactionx) PrepareFlow(query string) Splitted[*Stmt, *Transactionx] {
	return SplitOf(transactionx.Prepare(query), transactionx)
}

func (transactionx *Transactionx) Prepare(query string) *Stmt {
	return Do(func() (*Stmt, error) {
		result, err := transactionx.transactionx.Prepare(query)
		return StmtOf(result, transactionx), err
	}, transactionx, EmptyStmtOf)
}

func (transactionx *Transactionx) PrepareContextFlow(ctx context.Context, query string) Splitted[*Stmt, *Transactionx] {
	return SplitOf(transactionx.PrepareContext(ctx, query), transactionx)
}

func (transactionx *Transactionx) PrepareContext(ctx context.Context, query string) *Stmt {
	return Do(func() (*Stmt, error) {
		result, err := transactionx.transactionx.PrepareContext(ctx, query)
		return StmtOf(result, transactionx), err
	}, transactionx, EmptyStmtOf)
}

func (transactionx *Transactionx) QueryFlow(query string, args ...any) Splitted[*sql.Rows, *Transactionx] {
	return SplitOf(transactionx.Query(query, args...), transactionx)
}

func (transactionx *Transactionx) Query(query string, args ...any) *sql.Rows {
	return Do(func() (*sql.Rows, error) {
		result, err := transactionx.transactionx.Query(query, args...)
		return result, err
	}, transactionx, empty[*sql.Rows])
}

func (transactionx *Transactionx) QueryContextFlow(ctx context.Context, query string, args ...any) Splitted[*sql.Rows, *Transactionx] {
	return SplitOf(transactionx.QueryContext(ctx, query, args...), transactionx)
}

func (transactionx *Transactionx) QueryContext(ctx context.Context, query string, args ...any) *sql.Rows {
	return Do(func() (*sql.Rows, error) {
		result, err := transactionx.transactionx.QueryContext(ctx, query, args...)
		return result, err
	}, transactionx, empty[*sql.Rows])
}

func (transactionx *Transactionx) QueryRowFlow(query string, args ...any) Splitted[*sql.Row, *Transactionx] {
	return SplitOf(transactionx.QueryRow(query, args...), transactionx)
}

func (transactionx *Transactionx) QueryRow(query string, args ...any) *sql.Row {
	result := transactionx.transactionx.QueryRow(query, args...)
	return result
}

func (transactionx *Transactionx) QueryRowContextFlow(ctx context.Context, query string, args ...any) Splitted[*sql.Row, *Transactionx] {
	return SplitOf(transactionx.QueryRowContext(ctx, query, args...), transactionx)
}

func (transactionx *Transactionx) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	result := transactionx.transactionx.QueryRowContext(ctx, query, args...)
	return result
}

func (transactionx *Transactionx) Rollback() *Transactionx {
	return pass(func() error { return transactionx.transactionx.Rollback() }, transactionx)
}

func (transactionx *Transactionx) StmtFlow(stmt *sql.Stmt) Splitted[*Stmt, *Transactionx] {
	return SplitOf(transactionx.Stmt(stmt), transactionx)
}

func (transactionx *Transactionx) Stmt(stmt *sql.Stmt) *Stmt {
	result := transactionx.transactionx.Stmt(stmt)
	return StmtOf(result, transactionx)
}

func (transactionx *Transactionx) StmtContextFlow(ctx context.Context, stmt *sql.Stmt) Splitted[*Stmt, *Transactionx] {
	return SplitOf(transactionx.StmtContext(ctx, stmt), transactionx)
}

func (transactionx *Transactionx) StmtContext(ctx context.Context, stmt *sql.Stmt) *Stmt {
	result := transactionx.transactionx.StmtContext(ctx, stmt)
	return StmtOf(result, transactionx)
}

func (transactionx *Transactionx) Err() error {
	return transactionx.errs.Err()
}

func (transactionx *Transactionx) Fail(err error) {
	transactionx.errs.Fail(err)
}

func (transactionx *Transactionx) Link() *error {
	return transactionx.errs.Link()
}

func (transactionx *Transactionx) LinkTo(err *error) {
	transactionx.errs.LinkTo(err)
}

func (transactionx *Transactionx) Unwrap() (*sqlx.Tx, error) {
	return transactionx.transactionx, transactionx.Err()
}

func (transactionx *Transactionx) Raw() *sqlx.Tx {
	return transactionx.transactionx
}

func TransactionxOf(transactionx *sqlx.Tx, flow Linkable) *Transactionx {
	return &Transactionx{transactionx: transactionx, errs: errChainOf(flow)}
}

func NewTransactionx(transactionx *sqlx.Tx) *Transactionx {
	return &Transactionx{transactionx: transactionx, errs: emptyChain()}
}

func EmptyTransactionxOf(err error) *Transactionx {
	return &Transactionx{errs: errChainOfErr(err)}
}
