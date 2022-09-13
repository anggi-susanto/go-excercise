package transactionItem

import (
	"context"
	"testing"

	"github.com/anggi-susanto/go-exercise/entity"
	"github.com/anggi-susanto/go-exercise/entity/transaction"
	"github.com/stretchr/testify/assert"
)

func newFixturePromoRule() *transaction.PromoRule {
	id := uint32(1)
	return &transaction.PromoRule{
		ID:               &id,
		SkuID:            1,
		RequirementType:  "a",
		RequirementValue: "a",
		RewardType:       "",
		RewardValue:      1,
	}
}

func TestGetPromoRule(t *testing.T) {
	t.Parallel()
	repo := newInmem()
	m := NewService(repo)
	u := newFixturePromoRule()
	ctx := context.Background()
	p, err := repo.Create(ctx, u)
	assert.Nil(t, err)
	noDataID := uint32(2)
	_, err = m.GetPromoRule(ctx, &noDataID)
	assert.Equal(t, err, entity.ErrNotFound)
	saved, err := m.GetPromoRule(ctx, p.ID)
	assert.Nil(t, err)
	assert.Equal(t, saved.ID, p.ID)
}
