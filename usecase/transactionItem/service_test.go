package transactionItem

import (
	"context"
	"testing"
	"time"

	"github.com/anggi-susanto/go-exercise/entity"
	"github.com/anggi-susanto/go-exercise/entity/transaction"
	"github.com/stretchr/testify/assert"
)

func newFixtureTransactionItem() *transaction.TransactionItem {
	id := uint32(1)
	discount := float64(1)
	return &transaction.TransactionItem{
		ID:            &id,
		TransactionID: 1,
		SkuID:         1,
		BasePrice:     2.0,
		Discount:      &discount,
		Qty:           1.0,
		CreatedDate:   time.Now(),
	}
}

func TestCreateTransactionItem(t *testing.T) {
	t.Parallel()
	repo := newInmem()
	m := NewService(repo)
	u := newFixtureTransactionItem()
	ctx := context.Background()
	_, err := m.CreateTransactionItem(ctx, u.TransactionID, u.SkuID, u.BasePrice, *u.Discount, u.Qty)
	assert.Nil(t, err)
	assert.False(t, u.CreatedDate.IsZero())
}

func TestUpdateTransactionItem(t *testing.T) {
	t.Parallel()
	repo := newInmem()
	m := NewService(repo)
	u := newFixtureTransactionItem()
	ctx := context.Background()

	ti, err := m.CreateTransactionItem(ctx, u.TransactionID, u.SkuID, u.BasePrice, *u.Discount, u.Qty)
	assert.Nil(t, err)
	assert.False(t, u.CreatedDate.IsZero())

	noDataID := uint32(2)
	_, err = m.GetTransactionItem(ctx, &noDataID)
	assert.Equal(t, err, entity.ErrNotFound)

	saved, err := m.GetTransactionItem(ctx, ti.ID)
	assert.Nil(t, err)

	saved.SkuID = 2
	assert.Nil(t, m.UpdateTransactionItem(ctx, saved))
	updated, err := m.GetTransactionItem(ctx, saved.ID)
	assert.Nil(t, err)
	assert.Equal(t, updated.SkuID, uint32(2))

	updated.Qty = 0
	err = m.UpdateTransactionItem(ctx, updated)
	assert.Equal(t, err, entity.ErrInvalidEntity)
}
