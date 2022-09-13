package transaction

import (
	"context"

	"github.com/anggi-susanto/go-exercise/entity/transaction"
)

type Reader interface {
	Get(ctx context.Context, id *uint32) (res *transaction.Transaction, err error)
}

type Writer interface {
	Create(ctx context.Context, e *transaction.Transaction) (res *transaction.Transaction, err error)
	Update(ctx context.Context, e *transaction.Transaction) (err error)
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	CreateTransaction(
		ctx context.Context, userID uint32, subTotal float64, discount float64, total float64, status string,
	) (res *transaction.Transaction, err error)
	UpdateTransaction(ctx context.Context, e *transaction.Transaction) error
	GetTransaction(ctx context.Context, transactionId *uint32) (res transaction.Transaction, err error)
}
