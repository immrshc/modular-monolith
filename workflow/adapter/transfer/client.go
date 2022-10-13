package transfer

import (
	"context"

	"github.com/immrshc/modular-monolith/transfer/adapter/boundary"
	"github.com/immrshc/modular-monolith/transfer/usecase"
)

type ClientBuilder struct {
	// boundary.Transferの組み立てに必要なsql.DB以外の値を受け取る
	A int
	B int
}

type Client struct {
	transfer *boundary.Transfer
}

func NewClientBuilder(a, b int) *ClientBuilder {
	return &ClientBuilder{A: a, B: b}
}

func (c *ClientBuilder) NewClient(db interface{}) *Client {
	// boundary.Transferの組み立て方法の知識を持つ
	//transfer := boundary.NewTransfer(usecase.NewAccountTransfer(c.A, c.B, db))
	transfer := boundary.NewTransfer(usecase.NewAccountTransfer())
	return &Client{transfer: transfer}
}

func (c *Client) ReserveTransfer(ctx context.Context, amount int) error {
	return c.transfer.ReserveTransfer(ctx, amount)
}
