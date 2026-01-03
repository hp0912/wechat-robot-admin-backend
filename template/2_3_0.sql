-- 群聊设置
-- 创建时间：2026-01-02
-- 更新机器人实例库
alter table chat_room_settings
    add short_video_parsing_enabled tinyint(1) default 1 not null comment '是否启用短视频解析功能' after ltts_settings;