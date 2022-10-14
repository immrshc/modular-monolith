package usecase

import (
	"context"

	"github.com/immrshc/modular-monolith/workflow/usecase/port"
)

type TransferApplication struct {
	repo port.Repository
	tc   port.TransferClient
	rc   port.RegistrationClient
	tx   port.TxExecutor
}

func NewTransferApplication(repo port.Repository, tc port.TransferClient,
	rc port.RegistrationClient, tx port.TxExecutor) *TransferApplication {
	return &TransferApplication{repo: repo, tc: tc, rc: rc, tx: tx}
}

func (ta *TransferApplication) GetApplication(ctx context.Context, id int64) error {
	ta.rc.GetCompany(ctx, id)
	return nil
}

func (ta *TransferApplication) ApproveApplication(ctx context.Context, amount int) error {
	return ta.tx.ExecTx(
		ctx,
		func(ctx context.Context, repo port.Repository, tc port.TransferClient, rc port.RegistrationClient) error {
			if err := repo.UpdateApplication(ctx, 1); err != nil {
				return err
			}
			return tc.ReserveTransfer(ctx, amount)
		},
	)
}
