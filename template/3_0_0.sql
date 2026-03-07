-- Agent Skills 表
-- 创建时间：2026-03-06
-- 更新机器人实例库
CREATE TABLE IF NOT EXISTS `skills` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'Skill主键ID',
  `name` VARCHAR(128) NOT NULL COMMENT 'Skill名称',
  `path` VARCHAR(512) NOT NULL COMMENT 'Skill在磁盘上的绝对路径',
  `enabled` TINYINT(1) DEFAULT 1 COMMENT '是否启用',
  `source_type` VARCHAR(20) DEFAULT 'local' COMMENT '来源类型：local/git',
  `source` JSON COMMENT '来源详情(JSON)',
  `installed_at` DATETIME DEFAULT NULL COMMENT '安装时间',
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_skills_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='Agent Skills 表';