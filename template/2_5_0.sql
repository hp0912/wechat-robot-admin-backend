-- 群聊设置
-- 创建时间：2026-02-18
-- 更新机器人实例库
alter table chat_room_settings
    add wxhb_notify_enabled tinyint(1) default 0 null comment '是否启用微信红包通知功能' after short_video_parsing_enabled;

alter table chat_room_settings
    add wxhb_notify_member_list text not null comment '微信红包通知的成员列表，逗号分隔的微信ID' after wxhb_notify_enabled;

alter table chat_room_settings
    add podcast_enabled tinyint(1) default 0 null comment '是否启用AI播客功能' after wxhb_notify_member_list;

alter table chat_room_settings
    add podcast_config json null comment 'AI播客配置' after podcast_enabled;