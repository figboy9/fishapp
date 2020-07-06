CREATE TABLE `images`(
  `id` INT(11) PRIMARY KEY AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `owner_id` INT(11) NOT NULL,
  `owner_type` INT(11) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL
);
