-- +goose Up
-- +goose StatementBegin
INSERT INTO `users` (id, name, type) VALUES (1, 'Christy Schumm', 'coach');
INSERT INTO `users` (id, name, type) VALUES (2, 'Natalia Stanton Jr.', 'coach');
INSERT INTO `users` (id, name, type) VALUES (3, 'Nola Murazik V', 'coach');
INSERT INTO `users` (id, name, type) VALUES (4, 'Elyssa O\'Kon', 'coach');
INSERT INTO `users` (id, name, type) VALUES (5, 'Dr. Geovany Keebler', 'coach');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
