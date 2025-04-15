CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `wechat_id` varchar(64) DEFAULT NULL,
  `display_name` varchar(64) DEFAULT NULL,
  `role` bigint DEFAULT '1',
  `status` bigint DEFAULT '1',
  `avatar_url` varchar(500) DEFAULT '',
  `last_login_at` bigint DEFAULT '0',
  `created_at` bigint DEFAULT NULL,
  `deleted_at` DATETIME DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_we_chat_id` (`wechat_id`),
  KEY `idx_user_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `robot` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `robot_code` VARCHAR(64) NOT NULL,
  `owner` VARCHAR(64) NOT NULL,
  `device_id` VARCHAR(255) DEFAULT NULL,
  `device_name` VARCHAR(255) DEFAULT NULL,
  `wechat_id` VARCHAR(64) DEFAULT NULL,
  `nickname` VARCHAR(255) DEFAULT NULL,
  `avatar` VARCHAR(255) DEFAULT NULL,
  `status` ENUM('online', 'offline', 'error') NOT NULL DEFAULT 'offline',
  `redis_db` BIGINT UNSIGNED NOT NULL DEFAULT '1',
  `error_message` TEXT DEFAULT NULL,
  `last_login_at` BIGINT DEFAULT NULL,
  `created_at` BIGINT DEFAULT NULL,
  `updated_at` BIGINT DEFAULT NULL,
  `deleted_at` DATETIME DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_robot_code` (`robot_code`),
  UNIQUE KEY `idx_redis_db` (`redis_db`),
  KEY `idx_owner` (`owner`),
  KEY `idx_wechat_id` (`wechat_id`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;