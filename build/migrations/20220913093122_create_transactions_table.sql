-- +goose Up
-- +goose StatementBegin
CREATE TABLE `transactions`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `sub_total` decimal(10, 2) NULL,
  `discount` decimal(10, 2) NULL,
  `total` decimal(10, 2) NULL,
  `created_date` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  `updated_date` datetime(0) NULL,
  PRIMARY KEY (`id`),
  INDEX `in_transaction_create`(`created_date`) USING BTREE,
  INDEX `in_transaction_updated`(`updated_date`) USING BTREE,
  CONSTRAINT `fk_transaction_user_id` FOREIGN KEY (`user_id`) REFERENCES `exercise`.`users` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `transactions`;
-- +goose StatementEnd
