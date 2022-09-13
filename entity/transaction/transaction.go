package transaction

import (
	"time"

	"github.com/anggi-susanto/go-exercise/entity"
)

type Transaction struct {
	ID          *uint32
	UserID      uint32
	SubTotal    float64
	Discount    *float64
	Total       float64
	CreatedDate time.Time
	UpdatedDate *time.Time
}

func NewTransaction(userID uint32, subTotal float64, discount float64, total float64) (*Transaction, error) {
	t := &Transaction{
		UserID:      userID,
		SubTotal:    subTotal,
		Discount:    &discount,
		Total:       total,
		CreatedDate: time.Now(),
	}

	err := t.Validate()
	if err != nil {
		return nil, entity.ErrInvalidEntity
	}

	return t, nil
}

func (t *Transaction) Validate() error {
	if t.UserID < 1 || t.SubTotal < 1 || t.Total < 1 {
		return entity.ErrInvalidEntity
	}

	return nil
}
