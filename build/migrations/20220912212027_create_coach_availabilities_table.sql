-- +goose Up
-- +goose StatementBegin
CREATE TABLE `coach_availabilities` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `timezone` varchar(255) NOT NULL DEFAULT 'Asia/Jakarta',
  `day_of_week` enum('Sunday','Monday','Tuesday','Wednesday','Thursday','Friday','Saturday') NOT NULL,
  `time_start` time DEFAULT NULL,
  `time_end` time DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `in_time_start_end` (`time_start`,`time_end`) USING BTREE,
  KEY `fk_coach_availability` (`user_id`),
  CONSTRAINT `fk_coach_availability` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `coach_availabilities`;
-- +goose StatementEnd
