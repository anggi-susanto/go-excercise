-- +goose Up
-- +goose StatementBegin
CREATE TABLE `appointment_requests` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` enum('request','reschedule') NOT NULL DEFAULT 'request',
  `requester_id` int(11) NOT NULL,
  `approver_id` int(11) NOT NULL,
  `date` datetime NOT NULL,
  `status` enum('approved','rescheduled','rejected','requested') NOT NULL DEFAULT 'requested',
  `reschedule_request_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `in_appointment_req_date` (`date`) USING BTREE,
  KEY `fk_app_req_requester_user` (`requester_id`),
  KEY `fk_app_req_approver_user` (`approver_id`),
  CONSTRAINT `fk_app_req_approver_user` FOREIGN KEY (`approver_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE,
  CONSTRAINT `fk_app_req_requester_user` FOREIGN KEY (`requester_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `appointment_requests`;
-- +goose StatementEnd
