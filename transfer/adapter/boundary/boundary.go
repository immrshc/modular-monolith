package boundary

import (
	"context"
	"github.com/immrshc/modular-monolith/transfer/usecase/port"
)

type Transfer interface {
	ReserveTransfer(ctx context.Context, amount int) error
}

type transfer struct {
	atu port.AccountTransferUsecase
}

func NewTransfer(atu port.AccountTransferUsecase) Transfer{
	return &transfer{atu: atu}
}

func (t *transfer) ReserveTransfer(ctx context.Context, amount int) error {
	return t.atu.ReserveTransfer(ctx, amount)
}
