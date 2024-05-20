CREATE TABLE `transaction_history` (
  `id` integer NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `transaction_type` varchar(255),
  `amount` bigint,
  `user_id` integer,
  `account_id` integer,
  `timestamp` timestamp,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

ALTER TABLE `transaction_history` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `transaction_history` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);
