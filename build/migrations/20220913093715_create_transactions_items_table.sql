-- +goose Up
-- +goose StatementBegin
CREATE TABLE `transaction_items`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `transaction_id` int(11) NOT NULL,
  `sku_id` int(11) NOT NULL,
  `base_price` decimal(10, 2) NOT NULL,
  `discount` decimal(10, 2) NOT NULL DEFAULT 0,
  `created_date` datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updated_date` datetime(0) NULL ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`),
  INDEX `in_transaction_item_created`(`created_date`) USING BTREE,
  INDEX `in_transaction_item_updated`(`updated_date`) USING BTREE,
  CONSTRAINT `fk_transaction_item_trans` FOREIGN KEY (`transaction_id`) REFERENCES `exercise`.`transactions` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE,
  CONSTRAINT `fk_transaction_item_sku` FOREIGN KEY (`sku_id`) REFERENCES `exercise`.`skus` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `transaction_items`;
-- +goose StatementEnd
