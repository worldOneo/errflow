package errflow

import "database/sql"

type Result struct {
	result sql.Result
	errs   errChain
}

func (result *Result) LastInsertIdFlow() Splitted[int64, *Result] {
	return SplitOf(result.LastInsertId(), result)
}

func (result *Result) LastInsertId() int64 {
	return Do(func() (int64, error) {
		result, err := result.result.LastInsertId()
		return result, err
	}, result, empty[int64])
}

func (result *Result) RowsAffectedFlow() Splitted[int64, *Result] {
	return SplitOf(result.RowsAffected(), result)
}

func (result *Result) RowsAffected() int64 {
	return Do(func() (int64, error) {
		result, err := result.result.RowsAffected()
		return result, err
	}, result, empty[int64])
}

func (result *Result) Err() error {
	return result.errs.Err()
}

func (result *Result) Fail(err error) {
	result.errs.Fail(err)
}

func (result *Result) Link() *error {
	return result.errs.Link()
}

func (result *Result) LinkTo(err *error) {
	result.errs.LinkTo(err)
}

func (result *Result) Unwrap() (sql.Result, error) {
	return result.result, result.Err()
}

func (result *Result) Raw() sql.Result {
	return result.result
}

func ResultOf(result sql.Result, flow Linkable) *Result {
	return &Result{result: result, errs: errChainOf(flow)}
}

func NewResult(result sql.Result) *Result {
	return &Result{result: result, errs: emptyChain()}
}

func EmptyResultOf(err error) *Result {
	return &Result{errs: errChainOfErr(err)}
}
