package transactionItem

import (
	"context"

	"github.com/anggi-susanto/go-exercise/entity"
	"github.com/anggi-susanto/go-exercise/entity/transaction"
)

type inmem struct {
	m map[*uint32]*transaction.Sku
}

func newInmem() *inmem {
	var m = map[*uint32]*transaction.Sku{}
	return &inmem{
		m: m,
	}
}

func (r *inmem) Get(ctx context.Context, id *uint32) (res *transaction.Sku, err error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

func (r *inmem) Create(ctx context.Context, e *transaction.Sku) (res *transaction.Sku, err error) {
	if e.ID == nil {
		id := uint32(1)
		e.ID = &id
	}
	r.m[e.ID] = e
	return r.m[e.ID], nil
}
func (r *inmem) Update(ctx context.Context, e *transaction.Sku) (err error) {
	_, err = r.Get(ctx, e.ID)
	if err != nil {
		return err
	}
	r.m[e.ID] = e
	return nil
}
