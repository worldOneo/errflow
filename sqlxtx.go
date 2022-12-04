package errflow

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type Transactionx struct {
	tx   *sqlx.Tx
	errs errChain
}

func (tx *Transactionx) ExecFlow(query string, args ...any) Splitted[*Result, *Transactionx] {
	return SplitOf(tx.Exec(query, args...), tx)
}

func (tx *Transactionx) Exec(query string, args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := tx.tx.Exec(query, args...)
		return ResultOf(result, tx), err
	}, tx, EmptyResultOf)
}

func (tx *Transactionx) ExecContextFlow(ctx context.Context, query string, args ...any) Splitted[*Result, *Transactionx] {
	return SplitOf(tx.ExecContext(ctx, query, args...), tx)
}

func (tx *Transactionx) ExecContext(ctx context.Context, query string, args ...any) *Result {
	return Do(func() (*Result, error) {
		result, err := tx.tx.ExecContext(ctx, query, args...)
		return ResultOf(result, tx), err
	}, tx, EmptyResultOf)
}
func (tx *Transactionx) Get(dest any, query string, args ...any) *Transactionx {
	return pass(func() error { return tx.tx.Get(dest, query, args...) }, tx)
}

func (tx *Transactionx) GetContext(ctx context.Context, dest any, query string, args ...any) *Transactionx {
	return pass(func() error { return tx.tx.GetContext(ctx, dest, query, args...) }, tx)
}

func (tx *Transactionx) Commit() *Transactionx {
	return pass(func() error { return tx.tx.Commit() }, tx)
}

func (tx *Transactionx) Rollback() *Transactionx {
	return pass(func() error { return tx.tx.Rollback() }, tx)
}

func (tx *Transactionx) Err() error {
	return tx.errs.Err()
}

func (tx *Transactionx) Fail(err error) {
	tx.errs.Fail(err)
}

func (tx *Transactionx) Link() *error {
	return tx.errs.Link()
}

func (tx *Transactionx) LinkTo(err *error) {
	tx.errs.LinkTo(err)
}

func (tx *Transactionx) Unwrap() (*sqlx.Tx, error) {
	return tx.tx, tx.Err()
}

func (tx *Transactionx) Raw() *sqlx.Tx {
	return tx.tx
}

func TransactionxOf(tx *sqlx.Tx, flow Linkable) *Transactionx {
	return &Transactionx{tx: tx, errs: errChainOf(flow)}
}

func NewTransactionx(tx *sqlx.Tx) *Transactionx {
	return &Transactionx{tx: tx, errs: emptyChain()}
}

func EmptyTransactionxOf(err error) *Transactionx {
	return &Transactionx{errs: errChainOfErr(err)}
}
