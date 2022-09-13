package transactionItem

import (
	"context"

	"github.com/anggi-susanto/go-exercise/entity/transaction"
)

type Reader interface {
	Get(ctx context.Context, id *uint32) (res *transaction.Sku, err error)
}

type Writer interface {
	Create(ctx context.Context, e *transaction.Sku) (res *transaction.Sku, err error)
	Update(ctx context.Context, e *transaction.Sku) (err error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateSku(ctx context.Context, sku string, name string, price float64, inventoryQty int) (res *transaction.Sku, err error)
	UpdateSku(ctx context.Context, e *transaction.Sku) error
	GetSku(ctx context.Context, skuID *uint32) (res *transaction.Sku, err error)
}
