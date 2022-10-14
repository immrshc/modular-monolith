package usecase

import "github.com/immrshc/modular-monolith/registration/entity"

type Company struct{}

func NewCompany() *Company {
	return &Company{}
}

func (c *Company) GetCompany() *entity.Company {
	return &entity.Company{
		ID:      1,
		Name:    "hoge",
		Address: "fuga",
	}
}
