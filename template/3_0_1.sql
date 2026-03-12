-- 群聊设置
-- 创建时间：2026-03-12
-- 更新机器人实例库
alter table skills
    add env_vars JSON null comment '环境变量列表' after source;