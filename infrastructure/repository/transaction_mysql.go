package repository

import (
	"context"

	"github.com/anggi-susanto/go-exercise/entity/transaction"
	"github.com/jmoiron/sqlx"
	"go.opencensus.io/trace"
)

type TransactionMySQL struct {
	db *sqlx.DB
}

func NewTransactionMySQL(db *sqlx.DB) *TransactionMySQL {
	return &TransactionMySQL{
		db: db,
	}
}

func (t *TransactionMySQL) Create(ctx context.Context, e *transaction.Transaction) (res *transaction.Transaction, err error) {
	ctx, span := trace.StartSpan(ctx, "transaction: TransactionMySQL.Create")
	defer span.End()

	query := `INSERT INTO transactions 
			(user_id,sub_total,discount,total,created_date,updated_date) 
		VALUES (?,?,?,? NOW(), NOW());
		SELECT id, user_id, sub_total, discount, total, status, created_date, updated_date from transactions WHERE id=(SELECT LAST_INSERT_ID();`
	row := t.db.QueryRowContext(
		ctx,
		query,
		e.UserID,
		e.SubTotal,
		e.Discount,
		e.Total,
	)
	err = row.Scan(&res)
	if err != nil {
		return
	}

	return
}

func (t *TransactionMySQL) Get(ctx context.Context, id uint32) (res *transaction.Transaction, err error) {
	ctx, span := trace.StartSpan(ctx, "transaction: TransactionMySQL.Get")
	defer span.End()
	query := `SELECT id, user_id, sub_total, discount, total, created_date, updated_date from transactions WHERE id=?;`
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

func (s *TransactionMySQL) Update(ctx context.Context, e *transaction.Transaction) (err error) {
	ctx, span := trace.StartSpan(ctx, "transaction: TransactionMySQL.Update")
	defer span.End()

	query := `UPDATE transactions SET sub_total=?, discount=?, total=?, status=?, updated_date=NOW() WHERE id=?`
	_, err = s.db.ExecContext(
		ctx,
		query,
		e.SubTotal,
		e.Discount,
		e.Total,
		e.Status,
		e.ID,
	)

	if err != nil {
		return
	}

	return
}
