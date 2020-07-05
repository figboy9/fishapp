CREATE TABLE `apply_posts`(
  `id` INT(11) NOT NULL AUTO_INCREMENT,
  `post_id` INT(11) NOT NULL,
  `user_id` INT(11) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE (`post_id`, `user_id`),
  FOREIGN KEY (`post_id`) 
    REFERENCES posts(`id`)
    ON DELETE CASCADE
);