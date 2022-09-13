package transactionItem

import (
	"context"

	"github.com/anggi-susanto/go-exercise/entity"
	"github.com/anggi-susanto/go-exercise/entity/transaction"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
func (s *Service) CreateTransactionItem(ctx context.Context, transactionID uint32, skuID uint32, basePrice float64, discount float64, qty int) (res *transaction.TransactionItem, err error) {
	t, err := transaction.NewTransactionItem(transactionID, skuID, basePrice, discount, qty)
	if err != nil {
		return
	}

	return s.repo.Create(ctx, t)
}
func (s *Service) UpdateTransactionItem(ctx context.Context, e *transaction.TransactionItem) error {
	err := e.Validate()
	if err != nil {
		return err
	}

	return s.repo.Update(ctx, e)
}
func (s *Service) GetTransactionItem(ctx context.Context, transactionItemId *uint32) (res *transaction.TransactionItem, err error) {
	t, err := s.repo.Get(ctx, transactionItemId)

	if t == nil {
		err = entity.ErrNotFound
		return
	}
	if err != nil {
		return
	}
	return t, nil
}
