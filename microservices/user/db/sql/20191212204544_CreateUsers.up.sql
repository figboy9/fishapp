CREATE TABLE `users`(
  `id` INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `email` VARCHAR(255) UNIQUE NOT NULL,
  `encrypted_password` VARCHAR(255) NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  `introduction` TEXT NOT NULL,
  `sex` INT(11) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL
);
