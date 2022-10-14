package repository

import (
	"context"
)

type DBRepository struct {
	//h *sql.Tx
	h interface{}
}

func NewDB(h interface{}) *DBRepository {
	return &DBRepository{h: h}
}

func (d *DBRepository) UpdateApplication(ctx context.Context, id int64) error {
	return nil
}
