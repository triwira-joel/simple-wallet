CREATE TABLE `users` (
  `id` integer NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(255),
  `email` varchar(255),
  `password` varchar(255),
  `status` varchar(255),
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

CREATE TABLE `accounts` (
  `id` integer NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `user_id` integer,
  `balance` bigint,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

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

ALTER TABLE `accounts` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `transaction_history` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `transaction_history` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);
