package transactionItem

import (
	"context"

	"github.com/anggi-susanto/go-exercise/entity/transaction"
)

type Reader interface {
	GetBySkuID(ctx context.Context, id *uint32) (res *transaction.PromoRule, err error)
}

type Repository interface {
	Reader
}

type UseCase interface {
	GetPromoRule(ctx context.Context, skuID *uint32) (res *transaction.PromoRule, err error)
}
