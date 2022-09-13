package transactionItem

import (
	"context"

	"github.com/anggi-susanto/go-exercise/entity/transaction"
)

type Reader interface {
	Get(ctx context.Context, id *uint32) (res *transaction.TransactionItem, err error)
}

type Writer interface {
	Create(ctx context.Context, e *transaction.TransactionItem) (res *transaction.TransactionItem, err error)
	Update(ctx context.Context, e *transaction.TransactionItem) (err error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateTransactionItem(ctx context.Context, transactionID uint32, skuID uint32, basePrice float64, discount float64, qty int) (res *transaction.TransactionItem, err error)
	UpdateTransactionItem(ctx context.Context, e *transaction.TransactionItem) error
	GetTransactionItem(ctx context.Context, transactionItemId uint32) (res *transaction.TransactionItem, err error)
}
