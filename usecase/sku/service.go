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
func (s *Service) CreateSku(ctx context.Context, sku string, name string, price float64, inventoryQty int) (res *transaction.Sku, err error) {
	t, err := transaction.NewSku(sku, name, price, inventoryQty)
	if err != nil {
		return
	}

	return s.repo.Create(ctx, t)
}
func (s *Service) UpdateSku(ctx context.Context, e *transaction.Sku) error {
	err := e.Validate()
	if err != nil {
		return err
	}

	return s.repo.Update(ctx, e)
}
func (s *Service) GetSku(ctx context.Context, skuID *uint32) (res *transaction.Sku, err error) {
	t, err := s.repo.Get(ctx, skuID)

	if t == nil {
		err = entity.ErrNotFound
		return
	}
	if err != nil {
		return
	}
	return t, nil
}
