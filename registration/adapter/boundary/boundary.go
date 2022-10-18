package boundary

import (
	"context"

	"github.com/immrshc/modular-monolith/registration/entity"
	"github.com/immrshc/modular-monolith/registration/usecase/port"
)

type Registration struct {
	cu port.CompanyUsecase
}

func NewRegistration(cu port.CompanyUsecase) *Registration {
	return &Registration{cu: cu}
}

func (r *Registration) GetCompany(ctx context.Context, id int64) *entity.Company {
	// 自身のモジュール内のユースケースを呼び出す
	return r.cu.GetCompany()
}
