package errflow

func (tx *Transactionx) Stmt(stmt any) *Stmtx {
	return StmtxOf(tx.tx.Stmtx(stmt), tx)
}
