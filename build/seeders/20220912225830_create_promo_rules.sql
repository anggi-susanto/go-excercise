-- +goose Up
-- +goose StatementBegin
INSERT INTO `promo_rules` VALUES (1, 2, 'min_quantity', 1.00, 'product', 4), (2, 1, 'min_quantity', 3.00, 'discount_product', 1), (3, 3, 'min_quantity', 1.00, 'discount_percent', 10);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
