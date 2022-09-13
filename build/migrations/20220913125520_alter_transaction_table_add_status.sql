-- +goose Up
-- +goose StatementBegin
ALTER TABLE `transactions` 
ADD COLUMN `status` enum('cart','checkout') NOT NULL DEFAULT 'cart' AFTER `total`;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `transactions` 
DROP COLUMN `status`;
-- +goose StatementEnd
