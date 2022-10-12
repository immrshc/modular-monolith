package port

import "context"

type AccountTransferUsecase interface {
	ReserveTransfer(ctx context.Context, amount int) error
}
