-- 群成员管理
-- 创建时间：2025-12-31
-- 更新机器人实例库
UPDATE chat_room_members SET `score` = 0 WHERE `score` is NULL;

alter table chat_room_members
    add is_blacklisted tinyint(1) default 0 not null comment '是否在黑名单' after is_admin;

alter table chat_room_members
    modify score bigint default 0 not null comment '积分';

alter table chat_room_members
    add temporary_score bigint default 0 not null comment '临时积分' after score;

alter table chat_room_members
    add temporary_score_expiry bigint default 0 not null comment '临时积分有效期' after temporary_score;

alter table chat_room_members
    add frozen_score bigint default 0 not null comment '冻结积分' after temporary_score_expiry;

alter table chat_room_members
    add frozen_temporary_score bigint default 0 not null comment '冻结临时积分' after frozen_score;

