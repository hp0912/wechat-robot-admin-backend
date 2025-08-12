-- admin 数据库，如果没改过环境变量的话，是 robot_admin
-- 新增 user.api_token 字段及索引
-- use `robot_admin`;
DROP INDEX `idx_user_we_chat_id` ON `user`;
CREATE UNIQUE INDEX `uk_user_we_chat_id` ON `user` (`wechat_id`);
ALTER TABLE `user`
    ADD COLUMN `api_token` VARCHAR(128) DEFAULT NULL AFTER `display_name`;
CREATE UNIQUE INDEX `uk_user_api_token` ON `user` (`api_token`);

-- 机器人实例库，跟机器人 code 有关
-- use `机器人 code`;
alter table system_settings
    add api_token_enabled tinyint(1) default 0 not null comment '启用API Token' after id;