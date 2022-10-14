package port

import (
	"context"
)

type ApplicationUsecase interface {
	GetApplication(ctx context.Context, id int64) error
	ApproveApplication(ctx context.Context, amount int) error
}
