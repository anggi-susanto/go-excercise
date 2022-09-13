package transactionItem

import (
	"context"

	"github.com/anggi-susanto/go-exercise/entity"
	"github.com/anggi-susanto/go-exercise/entity/transaction"
)

type inmem struct {
	m map[*uint32]*transaction.PromoRule
}

func newInmem() *inmem {
	var m = map[*uint32]*transaction.PromoRule{}
	return &inmem{
		m: m,
	}
}

func (r *inmem) GetBySkuID(ctx context.Context, id *uint32) (res *transaction.PromoRule, err error) {
	if r.m[id] == nil {
		return nil, entity.ErrNotFound
	}
	return r.m[id], nil
}

func (r *inmem) Create(ctx context.Context, e *transaction.PromoRule) (res *transaction.PromoRule, err error) {
	if e.ID == nil {
		id := uint32(1)
		e.ID = &id
	}
	r.m[e.ID] = e
	return r.m[e.ID], nil
}
