package port

import "context"

type TxExecutor interface {
	ExecTx(
		context.Context,
		func(ctx context.Context, repo Repository, tc TransferClient, rc RegistrationClient) error,
	) error
}

type Repository interface {
	ApplicationRepository
}

type ApplicationRepository interface {
	UpdateApplication(ctx context.Context, id int64) error
}
