package repository

import (
	"context"

	"github.com/anggi-susanto/go-exercise/entity/transaction"
	"github.com/jmoiron/sqlx"
	"go.opencensus.io/trace"
)

type TransactionItemMySQL struct {
	db *sqlx.DB
}

func NewTransactionItemMysql(db *sqlx.DB) *TransactionItemMySQL {
	return &TransactionItemMySQL{
		db: db,
	}
}

func (t *TransactionItemMySQL) Create(ctx context.Context, e *transaction.TransactionItem) (res transaction.TransactionItem, err error) {
	ctx, span := trace.StartSpan(ctx, "transaction: TransactionItemMySQL.Create")
	defer span.End()

	query := `INSERT INTO transaction_items 
			(transaction_id,base_price,discount,created_date,updated_date) 
		VALUES (?,?,? NOW(), NOW());
		SELECT id, transaction_id, base_price, discount, created_date, updated_date from transaction_items WHERE id=(SELECT LAST_INSERT_ID();`
	row := t.db.QueryRowContext(
		ctx,
		query,
		e.TransactionID,
		e.BasePrice,
		e.Discount,
	)
	err = row.Scan(&res)
	if err != nil {
		return
	}

	return
}

func (t *TransactionItemMySQL) Get(ctx context.Context, id *uint32) (res transaction.TransactionItem, err error) {
	ctx, span := trace.StartSpan(ctx, "transaction: TransactionItemMySQL.Get")
	defer span.End()
	query := `SELECT id, transaction_id, base_price, discount, created_date, updated_date from transaction_items WHERE id=?;`
	row := t.db.QueryRowContext(
		ctx,
		query,
		id,
	)
	err = row.Scan(&res)
	if err != nil {
		return
	}

	return
}

func (t *TransactionItemMySQL) Update(ctx context.Context, e *transaction.TransactionItem) (err error) {
	ctx, span := trace.StartSpan(ctx, "transaction: TransactionMySQL.Update")
	defer span.End()

	query := `UPDATE transaction_items SET BasePrice=?, discount=?, qty=?, aupdated_date=NOW() WHERE id=?`
	_, err = t.db.ExecContext(
		ctx,
		query,
		e.BasePrice,
		e.Discount,
		e.Qty,
		e.ID,
	)

	if err != nil {
		return
	}
	return
}
