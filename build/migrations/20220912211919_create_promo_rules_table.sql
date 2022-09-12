-- +goose Up
-- +goose StatementBegin
CREATE TABLE `promo_rules` (
  `id` int(11) NOT NULL,
  `sku_id` int(11) NOT NULL,
  `requirement_type` enum('min_quantity','min_subtotal') NOT NULL DEFAULT 'min_quantity',
  `requirement_value` decimal(10,2) NOT NULL DEFAULT '1.00',
  `reward_type` enum('product','discount_percent','discount_product') DEFAULT 'discount_percent',
  `reward_value` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_sku_promo_rules` (`sku_id`),
  CONSTRAINT `fk_sku_promo_rules` FOREIGN KEY (`sku_id`) REFERENCES `skus` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `promo_rules`;
-- +goose StatementEnd
