-- 新增 user.api_token 字段及索引
ALTER TABLE `user`
    ADD COLUMN `api_token` VARCHAR(128) DEFAULT NULL AFTER `display_name`;
CREATE INDEX `idx_user_api_token` ON `user` (`api_token`);