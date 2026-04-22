-- 系统设置
-- 创建时间：2026-04-21
-- 更新机器人实例库
alter table system_settings
    modify column notification_type enum('push_plus', 'email', 'wechat_work_app') not null default 'push_plus' comment '通知方式：push_plus-推送加，email-邮件，wechat_work_app-企业微信应用';

alter table system_settings
    add wechat_work_corp_id varchar(255) default '' null comment '企业微信企业ID' after push_plus_token;

alter table system_settings
    add wechat_work_agent_id varchar(255) default '' null comment '企业微信应用AgentId' after wechat_work_corp_id;

alter table system_settings
    add wechat_work_secret varchar(255) default '' null comment '企业微信应用Secret' after wechat_work_agent_id;

alter table system_settings
    add wechat_work_proxy_url varchar(255) default '' null comment '企业微信代理地址' after wechat_work_secret;

alter table system_settings
    add wechat_work_to_user varchar(512) default '' null comment '企业微信推送用户ID，留空默认ALL' after wechat_work_proxy_url;
