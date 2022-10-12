package transfer

import (
	"context"

	"github.com/immrshch/modular-monolith/transfer/adapter/boundary"
)

type Client struct {
	transfer boundary.Transfer
}

func NewClient(transfer boundary.Transfer) *Client {
	return &Client{transfer: transfer}
}

func (c *Client) ReserveTransfer(ctx context.Context, amount int) error {
	return c.transfer.ReserveTransfer(ctx, amount)
}
