package repository

import (
	"context"

	"github.com/anggi-susanto/go-exercise/entity"
	"github.com/anggi-susanto/go-exercise/entity/transaction"
	"github.com/jmoiron/sqlx"
	"go.opencensus.io/trace"
)

type PromoRulesMySQL struct {
	db *sqlx.DB
}

func NewPromoRulesMysql(db *sqlx.DB) *PromoRulesMySQL {
	return &PromoRulesMySQL{
		db: db,
	}
}

func (p *PromoRulesMySQL) GetBySkuID(ctx context.Context, skuID uint32) (res transaction.PromoRule, err error) {
	ctx, span := trace.StartSpan(ctx, "promoRule: PromoRulesMySQL.GetBySkuID")
	defer span.End()

	query := `select id, requirement_type, requirement_value, reward_type, reward_value from promo_rules where sku_id=?`
	row := p.db.QueryRowContext(
		ctx,
		query,
		skuID,
	)

	err = row.Scan(
		&res.ID,
		&res.RequirementType,
		&res.RequirementValue,
		&res.RewardType,
		&res.RewardValue,
	)

	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			err = entity.ErrPromoRulesNotFound
		}
		return
	}

	return
}
