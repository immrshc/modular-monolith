package port

import "github.com/immrshc/modular-monolith/registration/entity"

type CompanyUsecase interface {
	GetCompany() *entity.Company
}
