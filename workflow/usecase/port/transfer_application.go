package port

import (
	"context"

	"github.com/immrshc/modular-monolith/workflow/entity"
)

type ApplicationUsecase interface {
	GetApplication(ctx context.Context, id int64) error
	ApproveApplication(ctx context.Context, amount int) error
}

type ApplicationRepository interface {
	UpdateApplication(ctx context.Context, id int64) error
}

type ApplicationRegistrationClient interface {
	GetCompany(ctx context.Context, id int64) *entity.Company
}

type ApplicationTransferClient interface {
	ReserveTransfer(ctx context.Context, amount int) error
}

type ApplicationTx interface {
	ExecApplicationTx(
		context.Context,
		func(ctx context.Context, repo ApplicationRepository, tc ApplicationTransferClient, rc ApplicationRegistrationClient) error,
	) error
}
