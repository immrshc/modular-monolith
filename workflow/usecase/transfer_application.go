package usecase

import (
	"context"
	"github.com/immrshc/modular-monolith/workflow/usecase/port"
)

type TransferApplication struct {
	transferClient port.TransferClient
}

func NewTransferApplication(tc port.TransferClient) *TransferApplication {
	return &TransferApplication{transferClient: tc}
}

func (ta *TransferApplication) ApproveApplication(ctx context.Context, amount int) error {
	return ta.transferClient.ReserveTransfer(ctx, amount)
}
