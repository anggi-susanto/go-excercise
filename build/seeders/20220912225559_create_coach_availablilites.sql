-- +goose Up
-- +goose StatementBegin
INSERT INTO `coach_availabilities` VALUES (1, 1, 'America/North_Dakota/New_Salem', 'Monday', '09:00:00', '17:30:00'),(2, 1, 'America/North_Dakota/New_Salem', 'Tuesday', '08:00:00', '16:00:00'),(3, 1, 'America/North_Dakota/New_Salem', 'Thursday', '09:00:00', '16:00:00'),(4, 1, 'America/North_Dakota/New_Salem', 'Friday', '07:00:00', '14:00:00'),(5, 2, 'Central Time (US & Canada)', 'Tuesday', '08:00:00', '10:00:00'),(6, 2, 'Central Time (US & Canada)', 'Wednesday', '11:00:00', '18:00:00'),(7, 2, 'Central Time (US & Canada)', 'Saturday', '09:00:00', '15:00:00'),(8, 2, 'Central Time (US & Canada)', 'Sunday', '08:00:00', '15:00:00'),(9, 3, 'America/Yakutat', 'Monday', '08:00:00', '10:00:00'),(10, 3, 'America/Yakutat', 'Tuesday', '11:00:00', '13:00:00'),(11, 3, 'America/Yakutat', 'Wednesday', '08:00:00', '10:00:00'),(12, 3, 'America/Yakutat', 'Saturday', '08:00:00', '11:00:00'),(13, 3, 'America/Yakutat', 'Sunday', '07:00:00', '09:00:00'),(14, 4, 'Central Time (US & Canada)', 'Monday', '09:00:00', '15:00:00'),(15, 4, 'Central Time (US & Canada)', 'Tuesday', '06:00:00', '13:00:00'),(16, 4, 'Central Time (US & Canada)', 'Wednesday', '06:00:00', '11:00:00'),(17, 4, 'Central Time (US & Canada)', 'Friday', '08:00:00', '12:00:00'),(18, 4, 'Central Time (US & Canada)', 'Saturday', '09:00:00', '14:00:00'),(19, 4, 'Central Time (US & Canada)', 'Sunday', '08:00:00', '10:00:00'),(20, 5, 'Central Time (US & Canada)', 'Thursday', '07:00:00', '14:00:00'),(21, 5, 'Central Time (US & Canada)', 'Thursday', '15:00:00', '17:00:00');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
