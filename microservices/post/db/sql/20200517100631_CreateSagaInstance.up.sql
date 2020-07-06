CREATE TABLE `saga_instance`(
  `id` VARCHAR(255) PRIMARY KEY,
  `saga_type` VARCHAR(255) NOT NULL,
  `saga_data` JSON NOT NULL,
  `current_state` VARCHAR(255) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL
);