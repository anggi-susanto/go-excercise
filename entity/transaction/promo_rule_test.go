package transaction_test

import (
	"testing"

	"github.com/anggi-susanto/go-exercise/entity"
	"github.com/anggi-susanto/go-exercise/entity/transaction"
	"github.com/stretchr/testify/assert"
)

func TestNewPromoRule(t *testing.T) {
	p, err := transaction.NewPromoRule(1, "a", "a", "a", 1)
	assert.Nil(t, err)
	assert.Equal(t, p.SkuID, uint32(1))
	assert.Equal(t, p.RequirementType, "a")
	assert.Equal(t, p.RequirementValue, "a")
	assert.Equal(t, p.RewardType, "a")
	assert.Equal(t, p.RewardValue, 1)
}

func TestPromoRuleValidate(t *testing.T) {
	type test struct {
		skuID uint32
		reqT  string
		reqV  string
		rewT  string
		rewV  int
		want  error
	}

	tests := []test{
		{
			skuID: 1,
			reqT:  "a",
			reqV:  "a",
			rewT:  "a",
			rewV:  1,
			want:  nil,
		},
		{
			skuID: 0,
			reqT:  "a",
			reqV:  "a",
			rewT:  "a",
			rewV:  1,
			want:  entity.ErrInvalidEntity,
		},
		{
			skuID: 1,
			reqT:  "",
			reqV:  "a",
			rewT:  "a",
			rewV:  1,
			want:  entity.ErrInvalidEntity,
		},
		{
			skuID: 1,
			reqT:  "a",
			reqV:  "",
			rewT:  "a",
			rewV:  1,
			want:  entity.ErrInvalidEntity,
		},
		{
			skuID: 1,
			reqT:  "a",
			reqV:  "a",
			rewT:  "",
			rewV:  1,
			want:  entity.ErrInvalidEntity,
		},
		{
			skuID: 1,
			reqT:  "a",
			reqV:  "a",
			rewT:  "a",
			rewV:  0,
			want:  entity.ErrInvalidEntity,
		},
	}

	for _, tc := range tests {
		_, err := transaction.NewPromoRule(tc.skuID, tc.reqT, tc.reqV, tc.rewT, tc.rewV)
		assert.Equal(t, err, tc.want)
	}
}
