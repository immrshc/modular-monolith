package registration

import (
	"context"

	"github.com/immrshc/modular-monolith/registration/adapter/boundary"
	"github.com/immrshc/modular-monolith/registration/usecase"
	"github.com/immrshc/modular-monolith/workflow/entity"
)

type ClientBuilder struct {
	// boundary.Registrationの組み立てに必要なsql.DB以外の値を受け取る
	A int
	B int
}

type Client struct {
	registration *boundary.Registration
}

func NewClientBuilder(a, b int) *ClientBuilder {
	return &ClientBuilder{A: a, B: b}
}

func (c *ClientBuilder) NewClient(db interface{}) *Client {
	// boundary.Registrationの組み立て方法の知識を持つ
	//registration := boundary.NewRegistration(usecase.NewCompany(c.A, c.B, db))
	registration := boundary.NewRegistration(usecase.NewCompany())
	return &Client{registration: registration}
}

func (c *Client) GetCompany(ctx context.Context, id int64) *entity.Company {
	// 事業所・申込モジュールの関数呼び出し
	com := c.registration.GetCompany(ctx, id)
	// registrationのモデルをworkflowのモデルに変換する
	return &entity.Company{ID: com.ID}
}
