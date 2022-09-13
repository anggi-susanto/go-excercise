package repository

import (
	"context"

	"github.com/anggi-susanto/go-exercise/entity"
	"github.com/anggi-susanto/go-exercise/entity/transaction"
	"github.com/jmoiron/sqlx"
	"go.opencensus.io/trace"
)

type SkuMySQL struct {
	db *sqlx.DB
}

func NewSkuMySQL(db *sqlx.DB) *SkuMySQL {
	return &SkuMySQL{
		db: db,
	}
}

func (s *SkuMySQL) Create(ctx context.Context, e *transaction.Sku) (res *transaction.Sku, err error) {
	ctx, span := trace.StartSpan(ctx, "sku: SkuMySQL.Create")
	defer span.End()

	query := `INSERT INTO skus
			(SKU,name,price,invetory_qty) 
		VALUES (?,?,?,?);
		SELECT id, SKU,name,price,invetory_qtye from skus WHERE id=(SELECT LAST_INSERT_ID();`
	row := s.db.QueryRowContext(
		ctx,
		query,
		e.SKU,
		e.Name,
		e.Price,
		e.InventoryQty,
	)
	err = row.Scan(&res)
	if err != nil {
		return
	}

	return
}

func (s *SkuMySQL) Get(ctx context.Context, id transaction.Sku) (res transaction.Sku, err error) {
	ctx, span := trace.StartSpan(ctx, "sku: SkuMySQL.Get")
	defer span.End()

	query := `select id, SKU, name, price, inventory_qty from skus where id=?`
	row := s.db.QueryRowContext(
		ctx,
		query,
		id,
	)

	err = row.Scan(
		&res.ID,
		&res.Name,
		&res.Price,
		&res.InventoryQty,
	)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			err = entity.ErrSkuNotFound
		}
		return
	}

	return
}

func (s *SkuMySQL) Update(ctx context.Context, skuID uint32, inventoryQty int) (err error) {
	ctx, span := trace.StartSpan(ctx, "sku: SkuMySQL.UpdateSku")
	defer span.End()

	query := `UPDATE sku SET inventory_qty=? WHERE sku_id=?`
	_, err = s.db.ExecContext(
		ctx,
		query,
		inventoryQty,
		skuID,
	)

	if err != nil {
		return
	}

	return
}
