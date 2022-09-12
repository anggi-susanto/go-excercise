-- +goose Up
-- +goose StatementBegin
CREATE TABLE `appointments` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `coach_id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `date` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `in_date` (`date`) USING BTREE,
  KEY `fk_appointment_user` (`user_id`),
  KEY `fk_appointment_coach` (`coach_id`),
  CONSTRAINT `fk_appointment_coach` FOREIGN KEY (`coach_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE,
  CONSTRAINT `fk_appointment_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `appointments`;
-- +goose StatementEnd
