package requests

type Cart struct {
	UserID uint32 `json:"user_id"`
	SkuID  uint32 `json:"sku_id"`
	Qty    int    `json:"qty"`
}
