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

func (s *Service) GetPromoRule(ctx context.Context, skuID *uint32) (res *transaction.PromoRule, err error) {
	t, err := s.repo.GetBySkuID(ctx, skuID)

	if t == nil {
		err = entity.ErrNotFound
		return
	}
	if err != nil {
		return
	}
	return t, nil
}
