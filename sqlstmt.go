package errflow

import (
	"context"
	"database/sql"
)

type VirtualStmt struct {
	stmt *sql.Stmt
	errs errChain
}

func stmtOf(stmt *sql.Stmt) *VirtualStmt {
	return &VirtualStmt{stmt: stmt, errs: emptyChain()}
}

func emptyStmtOf(err error) *VirtualStmt {
	return &VirtualStmt{errs: errChainOfErr(err)}
}

func (s *VirtualStmt) Exec(args ...any) sql.Result {
	return s.ExecContext(context.Background(), args...)
}

func (s *VirtualStmt) ExecContext(ctx context.Context, args ...any) sql.Result {
	return Do(func() (sql.Result, error) {
		result, err := s.stmt.ExecContext(ctx, args...)
		return result, err
	}, s, emptyResultOf)
}

func (s *VirtualStmt) Err() error {
	return s.errs.Err()
}

func (s *VirtualStmt) Fail(err error) {
	s.errs.Fail(err)
}

func (s *VirtualStmt) Link() *error {
	return s.errs.Link()
}

func (s *VirtualStmt) Destruct() (*sql.Stmt, error) {
	return s.stmt, s.errs.Err()
}

func (s *VirtualStmt) Raw() *sql.Stmt {
	return s.stmt
}
