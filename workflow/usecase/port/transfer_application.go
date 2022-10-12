package port

import "context"

type ApplicationUsecase interface {
	ApproveApplication(ctx context.Context, amount int) error
}

type TransferClient interface {
	ReserveTransfer(ctx context.Context, amount int) error
}