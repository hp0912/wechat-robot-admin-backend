CREATE TABLE IF NOT EXISTS `global_settings` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '公共配置表主键ID',
  -- 聊天模型AI设置
  `chat_ai_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用AI聊天功能',
  `chat_ai_trigger` VARCHAR(20) DEFAULT '' COMMENT '触发聊天AI的关键词',
  `chat_base_url` VARCHAR(255) DEFAULT '' COMMENT '聊天AI的基础URL地址',
  `chat_api_key` VARCHAR(255) DEFAULT '' COMMENT '聊天AI的API密钥',
  `chat_model` VARCHAR(100) DEFAULT '' COMMENT '聊天AI使用的模型名称',
  `chat_prompt` TEXT COMMENT '聊天AI系统提示词',
  `max_completion_tokens` INT DEFAULT 0 COMMENT '最大回复',
  -- 绘图模型AI设置
  `image_ai_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用AI绘图功能',
  `image_model` VARCHAR(255) DEFAULT '' COMMENT '绘图AI模型',
  `image_ai_settings` JSON COMMENT '绘图AI配置项',
  -- 文本转语音AI设置
  `tts_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用AI文本转语音功能',
  `tts_settings` JSON COMMENT '文本转语音AI配置项',
  `ltts_settings` JSON COMMENT '长文本转语音AI配置项',
  -- 拍一拍
  `pat_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用AI拍一拍功能',
  `pat_type` ENUM('text', 'voice') NOT NULL DEFAULT 'text' COMMENT '拍一拍方式：text-文本，voice-语音',
  `pat_text` BOOLEAN DEFAULT FALSE COMMENT '拍一拍文本',
  -- 欢迎新人
  `welcome_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用新成员加群欢迎功能',
  `welcome_type` ENUM('text','emoji', 'image', 'url') NOT NULL DEFAULT 'text' COMMENT '欢迎方式：text-文本，emoji-表情，image-图片，url-链接',
  `welcome_text` VARCHAR(255) DEFAULT '' COMMENT '欢迎新成员的文本',
  `welcome_emoji_md5` VARCHAR(64) DEFAULT '' COMMENT '欢迎新成员的表情MD5',
  `welcome_emoji_len` BIGINT DEFAULT 0 COMMENT '欢迎新成员的表情MD5长度',
  `welcome_image_url` VARCHAR(255) DEFAULT '' COMMENT '欢迎新成员的图片URL',
  `welcome_url` VARCHAR(255) DEFAULT '' COMMENT '欢迎新成员的URL',
  -- 退出群聊提醒
  `leave_chat_room_alert_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用群聊退出提醒功能',
  `leave_chat_room_alert_text` VARCHAR(255) DEFAULT '' COMMENT '群聊退出提醒文本',
  -- 群聊排行榜
  `chat_room_ranking_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用群聊排行榜功能',
  `chat_room_ranking_daily_cron` VARCHAR(255) DEFAULT '' COMMENT '每日定时任务表达式',
  `chat_room_ranking_weekly_cron` VARCHAR(255) DEFAULT '' COMMENT '每周定时任务表达式',
  `chat_room_ranking_month_cron` VARCHAR(255) DEFAULT '' COMMENT '每月定时任务表达式',
  -- 群聊总结
  `chat_room_summary_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用聊天记录总结功能',
  `chat_room_summary_model` VARCHAR(100) DEFAULT '' COMMENT '聊天总结使用的AI模型名称',
  `chat_room_summary_cron` VARCHAR(100) DEFAULT '' COMMENT '群聊总结的定时任务表达式',
  -- 每日早报
  `news_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用每日早报功能',
  `news_type` ENUM('text', 'image') NOT NULL DEFAULT 'text' COMMENT '是否启用每日早报功能',
  `news_cron` VARCHAR(100) DEFAULT '' COMMENT '每日早报的定时任务表达式',
  -- 每日早安
  `morning_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用早安问候功能',
  `morning_cron` VARCHAR(100) DEFAULT '' COMMENT '早安问候的定时任务表达式',
  -- 同步联系人
  `friend_sync_cron` VARCHAR(100) DEFAULT '' COMMENT '好友同步的定时任务表达式'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `friend_settings` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '群聊配置表主键ID',
  `wechat_id` VARCHAR(64) DEFAULT '' COMMENT '好友微信ID',
  -- 聊天模型AI设置
  `chat_ai_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用AI聊天功能',
  `chat_base_url` VARCHAR(255) DEFAULT '' COMMENT '聊天AI的基础URL地址',
  `chat_api_key` VARCHAR(255) DEFAULT '' COMMENT '聊天AI的API密钥',
  `chat_model` VARCHAR(100) DEFAULT '' COMMENT '聊天AI使用的模型名称',
  `chat_prompt` TEXT COMMENT '聊天AI系统提示词',
  `max_completion_tokens` INT DEFAULT 0 COMMENT '最大回复',
  -- 绘图模型AI设置
  `image_ai_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用AI绘图功能',
  `image_model` VARCHAR(255) DEFAULT '' COMMENT '绘图AI模型',
  `image_ai_settings` JSON COMMENT '绘图AI配置项',
   -- 文本转语音AI设置
  `tts_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用AI文本转语音功能',
  `tts_settings` JSON COMMENT '文本转语音AI配置项',
  `ltts_settings` JSON COMMENT '长文本转语音AI配置项',
  KEY `idx_wechat_id` (`wechat_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `chat_room_settings` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '群聊配置表主键ID',
  `chat_room_id` VARCHAR(64) DEFAULT '' COMMENT '群聊微信ID',
  -- 聊天模型AI设置
  `chat_ai_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用AI聊天功能',
  `chat_ai_trigger` VARCHAR(20) DEFAULT '' COMMENT '触发聊天AI的关键词',
  `chat_base_url` VARCHAR(255) DEFAULT '' COMMENT '聊天AI的基础URL地址',
  `chat_api_key` VARCHAR(255) DEFAULT '' COMMENT '聊天AI的API密钥',
  `chat_model` VARCHAR(100) DEFAULT '' COMMENT '聊天AI使用的模型名称',
  `chat_prompt` TEXT COMMENT '聊天AI系统提示词',
  `max_completion_tokens` INT DEFAULT 0 COMMENT '最大回复',
  -- 绘图模型AI设置
  `image_ai_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用AI绘图功能',
  `image_model` VARCHAR(255) DEFAULT '' COMMENT '绘图AI模型',
  `image_ai_settings` JSON COMMENT '绘图AI配置项',
   -- 文本转语音AI设置
  `tts_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用AI文本转语音功能',
  `tts_settings` JSON COMMENT '文本转语音AI配置项',
  `ltts_settings` JSON COMMENT '长文本转语音AI配置项',
    -- 拍一拍
  `pat_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用AI拍一拍功能',
  `pat_type` ENUM('text', 'voice') NOT NULL DEFAULT 'text' COMMENT '拍一拍方式：text-文本，voice-语音',
  `pat_text` VARCHAR(255) DEFAULT '' COMMENT '拍一拍文本',
  `pat_voice_timbre` VARCHAR(100) DEFAULT '' COMMENT '拍一拍音色',
  -- 欢迎新人
  `welcome_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用新成员加群欢迎功能',
  `welcome_type` ENUM('text','emoji', 'image', 'url') NOT NULL DEFAULT 'text' COMMENT '欢迎方式：text-文本，emoji-表情，image-图片，url-链接',
  `welcome_text` VARCHAR(255) DEFAULT '' COMMENT '欢迎新成员的文本',
  `welcome_emoji_md5` VARCHAR(64) DEFAULT '' COMMENT '欢迎新成员的表情MD5',
  `welcome_emoji_len` BIGINT DEFAULT 0 COMMENT '欢迎新成员的表情MD5长度',
  `welcome_image_url` VARCHAR(255) DEFAULT '' COMMENT '欢迎新成员的图片URL',
  `welcome_url` VARCHAR(255) DEFAULT '' COMMENT '欢迎新成员的URL',
    -- 退出群聊提醒
  `leave_chat_room_alert_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用群聊退出提醒功能',
  `leave_chat_room_alert_text` VARCHAR(255) DEFAULT '' COMMENT '群聊退出提醒文本',
  -- 群聊排行榜
  `chat_room_ranking_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用群聊排行榜功能',
  -- 群聊总结
  `chat_room_summary_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用聊天记录总结功能',
  `chat_room_summary_model` VARCHAR(100) DEFAULT '' COMMENT '聊天总结使用的AI模型名称',
  -- 每日早报
  `news_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用每日早报功能',
  `news_type` ENUM('', 'text', 'image') DEFAULT NULL COMMENT '每日早报类型：text-文本，image-图片',
  -- 每日早安
  `morning_enabled` BOOLEAN DEFAULT FALSE COMMENT '是否启用早安问候功能',
  KEY `idx_chat_room_id` (`chat_room_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `messages` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `msg_id` BIGINT NOT NULL,
  `client_msg_id` BIGINT NOT NULL,
  `type` INT NOT NULL,
  `app_msg_type` INT DEFAULT NULL,
  `is_chat_room` BOOLEAN DEFAULT FALSE COMMENT '是否为群聊消息',
  `is_at_me` BOOLEAN DEFAULT FALSE COMMENT '消息是否@我',
  `is_ai_context` BOOLEAN DEFAULT FALSE COMMENT '消息是否是AI的上下文',
  `is_recalled` BOOLEAN DEFAULT FALSE COMMENT '消息是否已经撤回',
  `content` TEXT,
  `display_full_content` TEXT,
  `message_source` TEXT,
  `from_wxid` VARCHAR(64),
  `sender_wxid` VARCHAR(64),
  `reply_wxid` VARCHAR(64),
  `to_wxid` VARCHAR(64),
  `attachment_url` VARCHAR(512),
  `created_at` BIGINT NOT NULL,
  `updated_at` BIGINT NOT NULL,
  UNIQUE KEY `uniq_msg_id` (`msg_id`),
  KEY `idx_from_wxid` (`from_wxid`),
  KEY `idx_type` (`type`),
  KEY `idx_sender_wxid` (`sender_wxid`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `system_messages` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
  `msg_id` BIGINT NOT NULL,
  `client_msg_id` BIGINT NOT NULL,
  `type` INT NOT NULL,
  `image_url` VARCHAR(512) DEFAULT NULL COMMENT '图片URL',
  `description` VARCHAR(255) DEFAULT NULL COMMENT '备注',
  `content` TEXT,
  `from_wxid` VARCHAR(64),
  `to_wxid` VARCHAR(64),
  `status` INT NOT NULL DEFAULT 0 COMMENT '消息状态：0-未处理，1-已处理',
  `is_read` BOOLEAN DEFAULT FALSE COMMENT '是否已读',
  `created_at` BIGINT NOT NULL,
  `updated_at` BIGINT NOT NULL,
  UNIQUE KEY `uniq_msg_id` (`msg_id`),
  KEY `idx_from_wxid` (`from_wxid`),
  KEY `idx_type` (`type`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `contacts` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `wechat_id` VARCHAR(64) NOT NULL COMMENT '微信ID',
  `alias` VARCHAR(64) DEFAULT NULL COMMENT '微信号',
  `nickname` VARCHAR(64) DEFAULT NULL COMMENT '昵称',
  `avatar` VARCHAR(255) DEFAULT NULL COMMENT '头像',
  `type` ENUM('friend','chat_room', 'official_account') NOT NULL COMMENT '联系人类型：friend-好友，chat_room-群组，official_account-公众号',
  `remark` VARCHAR(255) DEFAULT NULL COMMENT '备注',
  `pyinitial` VARCHAR(64) DEFAULT NULL COMMENT '昵称拼音首字母大写',
  `quan_pin` VARCHAR(255) DEFAULT NULL COMMENT '昵称拼音全拼小写',
  `sex` TINYINT DEFAULT 0 COMMENT '性别 0：未知 1：男 2：女',
  `country` VARCHAR(64) DEFAULT NULL COMMENT '国家',
  `province` VARCHAR(64) DEFAULT NULL COMMENT '省份',
  `city` VARCHAR(64) DEFAULT NULL COMMENT '城市',
  `signature` VARCHAR(255) DEFAULT NULL COMMENT '个性签名',
  `sns_background` VARCHAR(255) DEFAULT NULL COMMENT '朋友圈背景图',
  `chat_room_owner` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '群聊所有者微信ID',
  `created_at` BIGINT NOT NULL COMMENT '创建时间',
  `last_active_at` BIGINT NOT NULL COMMENT '最近活跃时间',
  `updated_at` BIGINT NOT NULL COMMENT '更新时间',
  `deleted_at` DATETIME DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_wechat_id` (`wechat_id`),
  KEY `idx_last_active_at` (`last_active_at`),
  KEY `idx_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `chat_room_members` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `chat_room_id` VARCHAR(64) NOT NULL COMMENT '群ID',
  `wechat_id` VARCHAR(64) NOT NULL COMMENT '微信ID',
  `alias` VARCHAR(64) DEFAULT NULL COMMENT '微信号',
  `nickname` VARCHAR(64) DEFAULT NULL COMMENT '昵称',
  `avatar` VARCHAR(255) DEFAULT NULL COMMENT '头像',
  `inviter_wechat_id` VARCHAR(64) NOT NULL COMMENT '邀请人微信ID',
  `is_admin` BOOLEAN DEFAULT FALSE COMMENT '是否群管理员',
  `is_leaved` BOOLEAN DEFAULT FALSE COMMENT '是否已经离开群聊',
  `score` BIGINT DEFAULT NULL COMMENT '积分',
  `remark` VARCHAR(255) DEFAULT NULL COMMENT '备注',
  `joined_at` BIGINT NOT NULL COMMENT '加入时间',
  `last_active_at` BIGINT NOT NULL COMMENT '最近活跃时间',
  `leaved_at` BIGINT DEFAULT NULL COMMENT '离开时间',
  PRIMARY KEY (`id`),
  KEY `idx_chat_room_id` (`chat_room_id`),
  KEY `idx_wechat_id` (`wechat_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `ai_task` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `contact_id` VARCHAR(64) NOT NULL COMMENT '联系人ID',
  `message_id` BIGINT NOT NULL COMMENT '消息ID，关联messages表的msg_id',
  `ai_provider_task_id` VARCHAR(64) DEFAULT NULL COMMENT 'AI服务商任务ID',
  `ai_task_type` ENUM('tts', 'ltts') NOT NULL COMMENT 'ltts-长文本转语音',
  `ai_task_status` ENUM('pending', 'processing', 'completed', 'failed') NOT NULL COMMENT '任务状态：pending-待处理，processing-处理中，completed-已完成，failed-已失败',
  `extra` JSON COMMENT '额外信息',
  `created_at` BIGINT NOT NULL COMMENT '创建时间',
  `updated_at` BIGINT NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_contact_id` (`contact_id`),
  UNIQUE KEY  `uk_message_id` (`message_id`),
  UNIQUE KEY  `uk_ai_provider_task_id` (`ai_provider_task_id`),
  KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;