CREATE TABLE `accounts` (
  `id` integer NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `user_id` integer,
  `balance` bigint,
  `created_at` timestamp DEFAULT (now()),
  `updated_at` timestamp DEFAULT (now())
);

ALTER TABLE `accounts` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);