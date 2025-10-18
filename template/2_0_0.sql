-- MCP服务器配置表
-- 支持四种传输模式：stdio(命令行)、sse、http、ws(WebSocket)
-- 创建时间：2025-10-02
-- 更新机器人实例库
CREATE TABLE IF NOT EXISTS `mcp_servers` (
  -- 基础字段
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'MCP服务器配置表主键ID',
  `name` varchar(100) NOT NULL COMMENT 'MCP服务器名称',
  `is_built_in` tinyint(1) DEFAULT 0 COMMENT '是否内置MCP服务器',
  `description` varchar(500) DEFAULT '' COMMENT 'MCP服务器描述',
  `transport` enum('stdio','sse','http','ws') NOT NULL COMMENT '传输类型：stdio-命令行，sse-SSE，http-HTTP，ws-WebSocket',
  `enabled` tinyint(1) DEFAULT 0 COMMENT '是否启用该MCP服务器',
  `priority` int DEFAULT 0 COMMENT '优先级，数字越大优先级越高',
  
  -- Stdio模式专用字段（命令行模式）
  `command` varchar(255) DEFAULT '' COMMENT '命令行模式的可执行命令',
  `args` json DEFAULT NULL COMMENT '命令行参数数组 []string',
  `working_dir` varchar(500) DEFAULT '' COMMENT '工作目录',
  `env` json DEFAULT NULL COMMENT '环境变量键值对 map[string]string',
  
  -- 网络模式专用字段（SSE/HTTP/WS共用）
  `url` varchar(500) DEFAULT '' COMMENT '服务器URL地址（SSE/HTTP/WS模式）',
  `auth_type` enum('none','bearer','basic','apikey') DEFAULT 'none' COMMENT '认证类型：none-无认证，bearer-Bearer Token，basic-Basic认证，apikey-API Key',
  `auth_token` varchar(500) DEFAULT '' COMMENT '认证令牌（Bearer Token或API Key）',
  `auth_username` varchar(100) DEFAULT '' COMMENT 'Basic认证用户名',
  `auth_password` varchar(255) DEFAULT '' COMMENT 'Basic认证密码',
  `headers` json DEFAULT NULL COMMENT '自定义HTTP请求头 map[string]string',
  `tls_skip_verify` tinyint(1) DEFAULT 0 COMMENT '是否跳过TLS证书验证',
  
  -- 超时和重连配置
  `connect_timeout` int DEFAULT 30 COMMENT '连接超时时间（秒）',
  `read_timeout` int DEFAULT 60 COMMENT '读取超时时间（秒）',
  `write_timeout` int DEFAULT 60 COMMENT '写入超时时间（秒）',
  `max_retries` int DEFAULT 3 COMMENT '最大重试次数',
  `retry_interval` int DEFAULT 5 COMMENT '重试间隔时间（秒）',
  `heartbeat_enable` tinyint(1) DEFAULT 1 COMMENT '是否启用心跳检测',
  `heartbeat_interval` int DEFAULT 30 COMMENT '心跳间隔时间（秒）',
  
  -- 高级配置
  `capabilities` json DEFAULT NULL COMMENT 'MCP服务器能力配置',
  `custom_config` json DEFAULT NULL COMMENT '自定义配置项',
  `tags` json DEFAULT NULL COMMENT '标签列表 []string',
  
  -- 状态追踪
  `last_connected_at` datetime DEFAULT NULL COMMENT '最后连接成功时间',
  `last_error` text COMMENT '最后一次错误信息',
  `connection_count` bigint DEFAULT 0 COMMENT '累计连接次数',
  `error_count` bigint DEFAULT 0 COMMENT '累计错误次数',
  
  -- 时间戳
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '软删除时间',
  
  PRIMARY KEY (`id`),
  KEY `idx_mcp_servers_deleted_at` (`deleted_at`),
  KEY `idx_enabled_priority` (`enabled`, `priority`),
  KEY `idx_transport` (`transport`),
  KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='MCP服务器配置表';