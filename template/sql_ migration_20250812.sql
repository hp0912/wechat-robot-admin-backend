-- admin 数据库，如果没改过环境变量的话，是 robot_admin
-- 新增 user.api_token 字段及索引
-- use `robot_admin`;
ALTER TABLE `user`
    ADD COLUMN `api_token` VARCHAR(128) DEFAULT NULL AFTER `display_name`;
CREATE INDEX `idx_user_api_token` ON `user` (`api_token`);