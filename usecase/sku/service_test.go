package transactionItem

import (
	"context"
	"testing"

	"github.com/anggi-susanto/go-exercise/entity"
	"github.com/anggi-susanto/go-exercise/entity/transaction"
	"github.com/stretchr/testify/assert"
)

func newFixtureSku() *transaction.Sku {
	id := uint32(1)
	return &transaction.Sku{
		ID:           &id,
		SKU:          "a",
		Name:         "a",
		Price:        2.0,
		InventoryQty: 1,
	}
}

func TestCreateSku(t *testing.T) {
	t.Parallel()
	repo := newInmem()
	m := NewService(repo)
	u := newFixtureSku()
	ctx := context.Background()
	_, err := m.CreateSku(ctx, u.SKU, u.Name, u.Price, u.InventoryQty)
	assert.Nil(t, err)
}

func TestUpdateSku(t *testing.T) {
	t.Parallel()
	repo := newInmem()
	m := NewService(repo)
	u := newFixtureSku()
	ctx := context.Background()

	s, err := m.CreateSku(ctx, u.SKU, u.Name, u.Price, u.InventoryQty)
	assert.Nil(t, err)

	noDataID := uint32(2)
	_, err = m.GetSku(ctx, &noDataID)
	assert.Equal(t, err, entity.ErrNotFound)

	saved, err := m.GetSku(ctx, s.ID)
	assert.Nil(t, err)

	saved.SKU = "b"
	assert.Nil(t, m.UpdateSku(ctx, saved))
	updated, err := m.GetSku(ctx, saved.ID)
	assert.Nil(t, err)
	assert.Equal(t, updated.SKU, "b")
}
