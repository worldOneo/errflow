package errflow

import "database/sql"

// Result is the flow wrapper of sql.Result
type Result struct {
	result sql.Result
	errs   errChain
}

// LastInsertIdFlow does the same as sql.Result.LastInsertId but splits the flow.
func (result *Result) LastInsertIdFlow() Splitted[int64, *Result] {
	return SplitOf(result.LastInsertId(), result)
}

// LastInsertId does the same as sql.Result but is a noop if this flow already failed.
func (result *Result) LastInsertId() int64 {
	return Do(func() (int64, error) {
		result, err := result.result.LastInsertId()
		return result, err
	}, result, empty[int64])
}

// RowsAffectedFlow does the same as sql.Result.RowsAffected but splits the flow.
func (result *Result) RowsAffectedFlow() Splitted[int64, *Result] {
	return SplitOf(result.RowsAffected(), result)
}

// RowsAffected does the same as sql.Result but is a noop if this flow already failed.
func (result *Result) RowsAffected() int64 {
	return Do(func() (int64, error) {
		result, err := result.result.RowsAffected()
		return result, err
	}, result, empty[int64])
}

// Err returns the error of this flow if any happend.
func (result *Result) Err() error {
	return result.errs.Err()
}

// Fail ends this flow with err
func (result *Result) Fail(err error) {
	result.errs.Fail(err)
}

// Link returns the base error of this flow.
func (result *Result) Link() *error {
	return result.errs.Link()
}

// LinkTo merges err as base into this flow.
func (result *Result) LinkTo(err *error) {
	result.errs.LinkTo(err)
}

// Unwrap returns the internal value and the produced error if any.
func (result *Result) Unwrap() (sql.Result, error) {
	return result.result, result.Err()
}

// Raw returns the internal value.
// No Guarantees.
func (result *Result) Raw() sql.Result {
	return result.result
}

// ResultOf create a new Result and is linked to flow.
// They are linked to gether and if one errors the error is propagated.
func ResultOf(result sql.Result, flow Linkable) *Result {
	return &Result{result: result, errs: errChainOf(flow)}
}

// NewResult create a new Result and is the root of a flow.
// It will catch any error that happens in the future.
func NewResult(result sql.Result) *Result {
	return &Result{result: result, errs: emptyChain()}
}

// EmptyResultOf returns an already failed Result.
// Calls will have no effects on it.
func EmptyResultOf(err error) *Result {
	return &Result{errs: errChainOfErr(err)}
}
