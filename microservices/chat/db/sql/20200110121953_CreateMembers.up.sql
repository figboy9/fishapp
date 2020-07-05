CREATE TABLE `members`(
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `room_id` INT(11) NOT NULL,
  `user_id` INT(11) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE (`room_id`, `user_id`),
  FOREIGN KEY (`room_id`) 
    REFERENCES rooms(`id`)
    ON DELETE CASCADE
)