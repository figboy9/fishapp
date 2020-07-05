CREATE TABLE `messages`(
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `body` VARCHAR(255) NOT NULL,
  `room_id` INT(11) NOT NULL,
  `user_id` INT(11) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`room_id`) 
    REFERENCES rooms(`id`)
    ON DELETE CASCADE
)