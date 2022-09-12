-- +goose Up
-- +goose StatementBegin
CREATE TABLE `skus` (
  `id` int(11) NOT NULL,
  `SKU` varchar(50) NOT NULL,
  `name` varchar(255) NOT NULL,
  `price` decimal(10,2) NOT NULL DEFAULT '0.00',
  `inventory_qty` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `skus`;
-- +goose StatementEnd
