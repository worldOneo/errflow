package errflow

import "database/sql"

type VirtualResult struct {
	res          sql.Result
	errs         errChain
	rowsAffected Box[int64]
	lastInserted Box[int64]
}

func emptyResultOf(err error) sql.Result {
	return &VirtualResult{errs: errChainOfErr(err)}
}

func resultOf(res sql.Result, parent Linkable) sql.Result {
	return &VirtualResult{res: res, errs: errChainOf(parent)}
}

func (res *VirtualResult) LastInsertId() (int64, error) {
	v, ok := res.lastInserted.Get()
	if !ok {
		v, _ = res.lastInserted.Fill(res.res.LastInsertId).Get()
	}
	return v, res.lastInserted.Err()
}

func (res *VirtualResult) RowsAffected() (int64, error) {
	v, ok := res.rowsAffected.Get()
	if !ok {
		v, _ = res.rowsAffected.Fill(res.res.RowsAffected).Get()
	}
	return v, res.rowsAffected.Err()
}

func (res *VirtualResult) Err() error {
	return res.errs.Err()
}

func (res *VirtualResult) Fail(err error) {
	res.errs.Fail(err)
}

func (res *VirtualResult) Link() *error {
	return res.errs.Link()
}

func (res *VirtualResult) LinkTo(err *error)  {
	res.errs.LinkTo(err)
}
