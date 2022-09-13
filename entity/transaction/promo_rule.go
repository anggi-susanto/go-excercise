package transaction

import "github.com/anggi-susanto/go-exercise/entity"

type PromoRule struct {
	ID               *uint32
	SkuID            uint32
	RequirementType  string
	RequirementValue string
	RewardType       string
	RewardValue      int
}

func NewPromoRule(skuID uint32, requirementType string, requirementValue string, rewardType string, rewardValue int) (*PromoRule, error) {
	p := &PromoRule{
		SkuID:            skuID,
		RequirementType:  requirementType,
		RequirementValue: requirementValue,
		RewardType:       rewardType,
		RewardValue:      rewardValue,
	}

	err := p.Validate()

	if err != nil {
		return nil, entity.ErrInvalidEntity
	}

	return p, nil
}

func (p *PromoRule) Validate() error {
	if p.SkuID < 1 || p.RequirementType == "" || p.RequirementValue == "" || p.RewardType == "" || p.RewardValue < 1 {
		return entity.ErrInvalidEntity
	}
	return nil
}
