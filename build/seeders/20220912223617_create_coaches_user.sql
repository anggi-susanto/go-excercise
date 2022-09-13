-- +goose Up
-- +goose StatementBegin
INSERT INTO `users`
(id, name, type)
VALUES (1, 'Christy Schumm', 'coach'),(2, 'Natalia Stanton Jr.', 'coach'),(3, 'Nola Murazik V', 'coach'),(4, 'Elyssa O\'Kon', 'coach'),(5, 'Dr. Geovany Keebler', 'coach');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
