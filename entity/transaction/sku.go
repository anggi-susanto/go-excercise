package transaction

import "github.com/anggi-susanto/go-exercise/entity"

type Sku struct {
	ID           *uint32
	SKU          string
	Name         string
	Price        float64
	InventoryQty int
}

func NewSku(sku string, name string, price float64, inventoryQty int) (*Sku, error) {
	s := &Sku{
		SKU:          sku,
		Name:         name,
		Price:        price,
		InventoryQty: inventoryQty,
	}

	err := s.Validate()

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}

	return s, nil
}
func (s *Sku) Validate() error {
	if s.SKU == "" || s.Name == "" || s.Price < float64(1) {
		return entity.ErrInvalidEntity
	}

	return nil
}
