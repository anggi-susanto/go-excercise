-- +goose Up
-- +goose StatementBegin
INSERT INTO `skus` VALUES (1, '120P90', 'Google Home', 49.99, 10), (2, '43N23P', 'MacBook Pro', 5399.99, 5), (3, 'A304SD', 'Alexa Speaker', 109.50, 10), (4, '234234', 'Raspberry Pi B', 30.00, 2);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
