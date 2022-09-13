-- +goose Up
-- +goose StatementBegin
ALTER TABLE `transaction_items` 
ADD COLUMN `qty` int(11) NULL AFTER `discount`;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `transaction_items` 
DROP COLUMN `qty`;
-- +goose StatementEnd
