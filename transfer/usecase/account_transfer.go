package usecase

import (
	"context"
	"fmt"
)

type AccountTransfer struct {
}

func NewAccountTransfer() *AccountTransfer {
	return &AccountTransfer{}
}

func (at *AccountTransfer) ReserveTransfer(ctx context.Context, amount int) error {
	fmt.Printf("=============%d円=============\n", amount)
	return nil
}
