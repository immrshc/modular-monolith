package usecase

import (
	"context"

	"github.com/immrshc/modular-monolith/workflow/usecase/port"
)

type TransferApplication struct {
	repo port.ApplicationRepository
	tc   port.ApplicationTransferClient
	rc   port.ApplicationRegistrationClient
	tx   port.ApplicationTx
}

func NewTransferApplication(
	repo port.ApplicationRepository,
	tc port.ApplicationTransferClient,
	rc port.ApplicationRegistrationClient,
	tx port.ApplicationTx,
) *TransferApplication {
	return &TransferApplication{repo: repo, tc: tc, rc: rc, tx: tx}
}

func (ta *TransferApplication) GetApplication(ctx context.Context, id int64) error {
	ta.rc.GetCompany(ctx, id)
	return nil
}

func (ta *TransferApplication) ApproveApplication(ctx context.Context, amount int) error {
	return ta.tx.ExecApplicationTx(
		ctx,
		func(
			ctx context.Context,
			repo port.ApplicationRepository,
			tc port.ApplicationTransferClient,
			rc port.ApplicationRegistrationClient,
		) error {
			if err := repo.UpdateApplication(ctx, 1); err != nil {
				return err
			}
			return tc.ReserveTransfer(ctx, amount)
		},
	)
}
