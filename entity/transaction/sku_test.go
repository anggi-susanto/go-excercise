package transaction_test

import (
	"testing"

	"github.com/anggi-susanto/go-exercise/entity"
	"github.com/anggi-susanto/go-exercise/entity/transaction"
	"github.com/stretchr/testify/assert"
)

func TestNewSku(t *testing.T) {
	s, err := transaction.NewSku("a", "a", 1.0, 1)
	assert.Nil(t, err)
	assert.Equal(t, s.SKU, "a")
	assert.Equal(t, s.Name, "a")
	assert.Equal(t, s.Price, 1.0)
	assert.Equal(t, s.InventoryQty, 1)
}

func TestSkuValidate(t *testing.T) {
	type test struct {
		sku   string
		name  string
		price float64
		inv   int
		want  error
	}

	tests := []test{
		{
			sku:   "a",
			name:  "a",
			price: 1.0,
			inv:   1,
			want:  nil,
		},
		{
			sku:   "",
			name:  "a",
			price: 1.0,
			inv:   1,
			want:  entity.ErrInvalidEntity,
		},
		{
			sku:   "a",
			name:  "",
			price: 1.0,
			inv:   1,
			want:  entity.ErrInvalidEntity,
		},
		{
			sku:   "a",
			name:  "a",
			price: 0,
			inv:   1,
			want:  entity.ErrInvalidEntity,
		},
		{
			sku:   "a",
			name:  "a",
			price: 1.0,
			inv:   0,
			want:  entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := transaction.NewSku(tc.sku, tc.name, tc.price, tc.inv)
		assert.Equal(t, err, tc.want)
	}
}
