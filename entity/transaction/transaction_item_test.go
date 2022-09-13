package transaction_test

import (
	"testing"

	"github.com/anggi-susanto/go-exercise/entity"
	"github.com/anggi-susanto/go-exercise/entity/transaction"
	"github.com/stretchr/testify/assert"
)

func TestNewTransactionItem(t *testing.T) {
	tr, err := transaction.NewTransactionItem(1, 1, 1.0, 0.5, 1)

	assert.Nil(t, err)
	assert.Equal(t, tr.TransactionID, uint32(1))
	assert.Equal(t, tr.SkuID, uint32(1))
	assert.Equal(t, tr.BasePrice, 1.0)
	assert.Equal(t, tr.Qty, 1)
}

func TestTransactionItemValidate(t *testing.T) {
	type test struct {
		transactionID uint32
		skuID         uint32
		basePrice     float64
		discount      float64
		qty           int
		want          error
	}

	tests := []test{
		{
			transactionID: 1,
			skuID:         1,
			basePrice:     1.0,
			discount:      0,
			qty:           1,
			want:          nil,
		},
		{
			transactionID: 0,
			skuID:         1,
			basePrice:     1.0,
			discount:      0,
			qty:           1,
			want:          entity.ErrInvalidEntity,
		},
		{
			transactionID: 1,
			skuID:         0,
			basePrice:     1.0,
			discount:      0,
			qty:           1,
			want:          entity.ErrInvalidEntity,
		},
		{
			transactionID: 1,
			skuID:         1,
			basePrice:     0,
			discount:      0,
			qty:           1,
			want:          entity.ErrInvalidEntity,
		},
		{
			transactionID: 1,
			skuID:         1,
			basePrice:     1,
			discount:      0,
			qty:           0,
			want:          entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := transaction.NewTransactionItem(tc.transactionID, tc.skuID, tc.basePrice, tc.discount, tc.qty)
		assert.Equal(t, err, tc.want)
	}
}
