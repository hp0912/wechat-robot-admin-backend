CREATE TABLE IF NOT EXISTS `system_settings` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '表主键ID',
  `offline_notification_enabled` BOOLEAN DEFAULT FALSE COMMENT '启用离线通知',
  `notification_type` ENUM('push_plus', 'email') NOT NULL DEFAULT 'push_plus' COMMENT '通知方式：push_plus-推送加，email-邮件',
  `push_plus_url` VARCHAR(255) DEFAULT '' COMMENT 'Push Plus的URL',
  `push_plus_token` VARCHAR(255) DEFAULT '' COMMENT 'Push Plus的Token',
  `auto_verify_user` BOOLEAN DEFAULT FALSE COMMENT '自动通过好友验证',
  `verify_user_delay` INT DEFAULT 60 COMMENT '自动通过好友验证延迟时间(秒)',
  `auto_chatroom_invite` BOOLEAN DEFAULT FALSE COMMENT '自动邀请进群'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;