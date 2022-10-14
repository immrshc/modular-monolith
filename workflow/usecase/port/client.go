package port

import (
	"context"

	"github.com/immrshc/modular-monolith/workflow/entity"
)

type RegistrationClient interface {
	GetCompany(ctx context.Context, id int64) *entity.Company
}

type TransferClient interface {
	ReserveTransfer(ctx context.Context, amount int) error
}
