package boundary

import (
	"context"

	"github.com/immrshc/modular-monolith/transfer/usecase/port"
)

type Transfer struct {
	atu port.AccountTransferUsecase
}

func NewTransfer(atu port.AccountTransferUsecase) *Transfer{
	return &Transfer{atu: atu}
}

func (t *Transfer) ReserveTransfer(ctx context.Context, amount int) error {
	return t.atu.ReserveTransfer(ctx, amount)
}
