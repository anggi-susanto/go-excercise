package transaction

import (
	"context"

	"github.com/anggi-susanto/go-exercise/entity"
	"github.com/anggi-susanto/go-exercise/entity/transaction"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) CreateTransaction(
	ctx context.Context, userID uint32, subTotal float64, discount float64, total float64, status string,
) (res *transaction.Transaction, err error) {
	t, err := transaction.NewTransaction(userID, subTotal, discount, total, status)
	if err != nil {
		return
	}

	return s.repo.Create(ctx, t)
}
func (s *Service) UpdateTransaction(ctx context.Context, e *transaction.Transaction) error {
	err := e.Validate()
	if err != nil {
		return err
	}

	return s.repo.Update(ctx, e)
}
func (s *Service) GetTransaction(ctx context.Context, transactionId *uint32) (res *transaction.Transaction, err error) {
	t, err := s.repo.Get(ctx, transactionId)

	if t == nil {
		err = entity.ErrNotFound
		return
	}
	if err != nil {
		return
	}
	return t, nil
}
