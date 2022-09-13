package transaction

import (
	"time"

	"github.com/anggi-susanto/go-exercise/entity"
)

type TransactionItem struct {
	ID            *uint32
	TransactionID uint32
	SkuID         uint32
	BasePrice     float64
	Discount      *float64
	Qty           int
	CreatedDate   time.Time
	UpdatedDate   *time.Time
}

func NewTransactionItem(transactionID uint32, skuID uint32, basePrice float64, discount float64, qty int) (*TransactionItem, error) {
	t := &TransactionItem{
		TransactionID: transactionID,
		SkuID:         skuID,
		BasePrice:     basePrice,
		Discount:      &discount,
		Qty:           qty,
		CreatedDate:   time.Now(),
	}

	err := t.Validate()
	if err != nil {
		return nil, entity.ErrInvalidEntity
	}

	return t, nil

}

func (t *TransactionItem) Validate() error {
	if t.TransactionID < 1 || t.SkuID < 1 || t.BasePrice < 1 || t.Qty < 1 {
		return entity.ErrInvalidEntity
	}

	return nil
}
