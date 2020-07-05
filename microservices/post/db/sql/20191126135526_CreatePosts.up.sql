CREATE TABLE `posts`(
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(255) NOT NULL,
  `content` VARCHAR(5000) NOT NULL,
  `fishing_spot_type_id` INT(11) NOT NULL,
  `prefecture_id` INT(11),
  `meeting_place_id` VARCHAR(255) NOT NULL,
  `meeting_at` DATETIME NOT NULL,
  `max_apply` INT(11) NOT NULL,
  `user_id` INT(11) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
);