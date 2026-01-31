-- 群聊设置
-- 创建时间：2026-01-27
-- 更新机器人实例库
alter table global_settings
    drop column image_model;

alter table friend_settings
    drop column image_model;

alter table chat_room_settings
    drop column image_model;