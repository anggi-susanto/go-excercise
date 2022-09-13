package transaction

import (
	"context"
	"testing"
	"time"

	"github.com/anggi-susanto/go-exercise/entity"
	"github.com/anggi-susanto/go-exercise/entity/transaction"
	"github.com/stretchr/testify/assert"
)

func newFixtureTransaction() *transaction.Transaction {
	id := uint32(1)
	discount := float64(1)
	return &transaction.Transaction{
		ID:          &id,
		UserID:      1,
		SubTotal:    2.0,
		Discount:    &discount,
		Total:       1.0,
		Status:      "cart",
		CreatedDate: time.Now(),
	}
}

func TestCreateTransaction(t *testing.T) {
	t.Parallel()
	repo := newInmem()
	m := NewService(repo)
	u := newFixtureTransaction()
	ctx := context.Background()
	_, err := m.CreateTransaction(ctx, u.UserID, u.SubTotal, *u.Discount, u.Total, u.Status)
	assert.Nil(t, err)
	assert.False(t, u.CreatedDate.IsZero())
}

func TestUpdateTransaction(t *testing.T) {
	t.Parallel()
	repo := newInmem()
	m := NewService(repo)
	u := newFixtureTransaction()
	ctx := context.Background()
	tr, err := m.CreateTransaction(ctx, u.UserID, u.SubTotal, *u.Discount, u.Total, u.Status)
	assert.Nil(t, err)
	noDataID := uint32(2)
	_, err = m.GetTransaction(ctx, &noDataID)
	assert.Equal(t, err, entity.ErrNotFound)

	saved, err := m.GetTransaction(ctx, tr.ID)
	assert.Nil(t, err)
	saved.UserID = 2
	assert.Nil(t, m.UpdateTransaction(ctx, saved))
	updated, err := m.GetTransaction(ctx, saved.ID)
	assert.Nil(t, err)
	assert.Equal(t, updated.UserID, uint32(2))
	updated.Status = ""
	err = m.UpdateTransaction(ctx, updated)
	assert.Equal(t, err, entity.ErrInvalidEntity)

}
