package transaction_test

import (
	"testing"

	"github.com/anggi-susanto/go-exercise/entity"
	"github.com/anggi-susanto/go-exercise/entity/transaction"
	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	trans, err := transaction.NewTransaction(1, 2.0, 1.0, 1.0)
	assert.Nil(t, err)
	assert.Equal(t, trans.UserID, uint32(1))
	assert.Equal(t, trans.SubTotal, 2.0)
	assert.Equal(t, trans.Total, 1.0)
}

func TestTransactionValidate(t *testing.T) {
	type test struct {
		userID   uint32
		subtotal float64
		discount float64
		total    float64
		want     error
	}

	tests := []test{
		{
			userID:   1,
			subtotal: 2.0,
			discount: 1.0,
			total:    1.0,
			want:     nil,
		},
		{
			userID:   0,
			subtotal: 2.0,
			discount: 1.0,
			total:    1.0,
			want:     entity.ErrInvalidEntity,
		},
		{
			userID:   1,
			subtotal: 0,
			discount: 1.0,
			total:    1.0,
			want:     entity.ErrInvalidEntity,
		},
		{
			userID:   0,
			subtotal: 2.0,
			discount: 1.0,
			total:    0,
			want:     entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := transaction.NewTransaction(tc.userID, tc.subtotal, tc.discount, tc.total)
		assert.Equal(t, err, tc.want)
	}
}
