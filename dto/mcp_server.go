package dto

// MCPTransportType MCP服务器传输类型
type MCPTransportType string

const (
	MCPTransportTypeStdio MCPTransportType = "stdio" // 命令行模式（标准输入输出）
	MCPTransportTypeSSE   MCPTransportType = "sse"   // Server-Sent Events模式
	MCPTransportTypeHTTP  MCPTransportType = "http"  // HTTP模式
	MCPTransportTypeWS    MCPTransportType = "ws"    // WebSocket模式
)

// MCPAuthType MCP认证类型
type MCPAuthType string

const (
	MCPAuthTypeNone   MCPAuthType = "none"   // 无认证
	MCPAuthTypeBearer MCPAuthType = "bearer" // Bearer Token认证
	MCPAuthTypeBasic  MCPAuthType = "basic"  // Basic认证
	MCPAuthTypeAPIKey MCPAuthType = "apikey" // API Key认证
)

type MCPServer struct {
	ID          uint64           `json:"id"`
	Name        string           `json:"name"`
	IsBuiltIn   *bool            `json:"is_built_in"`
	Description string           `json:"description"`
	Transport   MCPTransportType `json:"transport"`
	Enabled     *bool            `json:"enabled"`
	Priority    int              `json:"priority"`

	// Stdio模式配置（命令行模式）
	Command    string            `json:"command"`
	Args       []string          `json:"args"` // []string
	WorkingDir string            `json:"working_dir"`
	Env        map[string]string `json:"env"` // map[string]string

	// 网络模式配置（SSE/HTTP/WS共用）
	URL           string            `json:"url"`
	AuthType      MCPAuthType       `json:"auth_type"`
	AuthToken     string            `json:"auth_token"`
	AuthUsername  string            `json:"auth_username"`
	AuthPassword  string            `json:"auth_password"`
	Headers       map[string]string `json:"headers"` // map[string]string
	TLSSkipVerify *bool             `json:"tls_skip_verify"`

	// 超时和重连配置
	ConnectTimeout    int   `json:"connect_timeout"`
	ReadTimeout       int   `json:"read_timeout"`
	WriteTimeout      int   `json:"write_timeout"`
	MaxRetries        int   `json:"max_retries"`
	RetryInterval     int   `json:"retry_interval"`
	HeartbeatEnable   *bool `json:"heartbeat_enable"`
	HeartbeatInterval int   `json:"heartbeat_interval"`

	// 高级配置
	Capabilities map[string]string `json:"capabilities"`  // 服务器支持的能力
	CustomConfig map[string]string `json:"custom_config"` // 其他自定义配置
	Tags         []string          `json:"tags"`          // []string，用于分类和过滤

	// 状态追踪
	LastConnectedAt *string `json:"last_connected_at"`
	LastError       string  `json:"last_error"`
	ConnectionCount int64   `json:"connection_count"`
	ErrorCount      int64   `json:"error_count"`

	// 时间戳
	CreatedAt *string `json:"created_at"`
	UpdatedAt *string `json:"updated_at"`
	DeletedAt *string `json:"deleted_at,omitempty"`
}

type MCPServerTool struct {
	Name        string `json:"name"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}
