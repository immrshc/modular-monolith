package repository

import (
	"context"

	"github.com/immrshc/modular-monolith/workflow/adapter/registration"
	"github.com/immrshc/modular-monolith/workflow/adapter/transfer"
	"github.com/immrshc/modular-monolith/workflow/usecase/port"
)

type TxExecutor struct {
	//tx *sql.Tx
	tx interface{}
	tcb *transfer.ClientBuilder
	rcb *registration.ClientBuilder
}

func NewTxExecutor(tx interface{}, tcb *transfer.ClientBuilder, rcb *registration.ClientBuilder) *TxExecutor {
	return &TxExecutor{tx: tx, tcb: tcb, rcb: rcb}
}

func (te *TxExecutor) ExecApplicationTx(
	ctx context.Context,
	op func(ctx context.Context, repo port.ApplicationRepository, tc port.ApplicationTransferClient, rc port.ApplicationRegistrationClient) error,
) error {
	// BeginTx
	// defer te.tx.Commit()
	return op(ctx, NewDB(te.tx), te.tcb.NewClient(te.tx), te.rcb.NewClient(te.tx))
}