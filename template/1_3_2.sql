-- 1.3.2版本更新脚本
CREATE TABLE IF NOT EXISTS `oss_settings` (
  `id` BIGINT AUTO_INCREMENT PRIMARY KEY COMMENT '表主键ID',
  `auto_upload_image` BOOLEAN DEFAULT FALSE COMMENT '启用自动上传图片',
  `auto_upload_image_mode` ENUM('all', 'ai_only') NOT NULL DEFAULT 'ai_only' COMMENT '自动上传图片模式',
  `auto_upload_video` BOOLEAN DEFAULT FALSE COMMENT '启用自动上传视频',
  `auto_upload_video_mode` ENUM('all', 'ai_only') NOT NULL DEFAULT 'ai_only' COMMENT '自动上传视频模式',
  `auto_upload_file` BOOLEAN DEFAULT FALSE COMMENT '启用自动上传文件',
  `auto_upload_file_mode` ENUM('all', 'ai_only') NOT NULL DEFAULT 'ai_only' COMMENT '自动上传文件模式',
  `oss_provider` ENUM('aliyun', 'tencent_cloud', 'cloudflare') NOT NULL DEFAULT 'aliyun' COMMENT '对象存储服务商',
  `aliyun_oss_settings` JSON COMMENT '阿里云OSS配置项',
  `tencent_cloud_oss_settings` JSON COMMENT '腾讯云OSS配置项',
  `cloudflare_r2_settings` JSON COMMENT 'Cloudflare R2配置项',
  `created_at` BIGINT NOT NULL,
  `updated_at` BIGINT NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;