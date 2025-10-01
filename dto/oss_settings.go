package dto

type AutoUploadMode string

const (
	AutoUploadModeAll    AutoUploadMode = "all" // 全部上传
	AutoUploadModeAIOnly AutoUploadMode = "ai_only"
)

type OSSProvider string

const (
	OSSProviderAliyun       OSSProvider = "aliyun"        // 阿里云 OSS
	OSSProviderTencentCloud OSSProvider = "tencent_cloud" // 腾讯云 COS
	OSSProviderCloudflare   OSSProvider = "cloudflare"    // Cloudflare R2
)

type OSSSettingsRequest struct {
	ID                      uint64         `json:"id"`
	AutoUploadImage         *bool          `json:"auto_upload_image"`
	AutoUploadImageMode     AutoUploadMode `json:"auto_upload_image_mode"`
	AutoUploadVideo         *bool          `json:"auto_upload_video"`
	AutoUploadVideoMode     AutoUploadMode `json:"auto_upload_video_mode"`
	AutoUploadFile          *bool          `json:"auto_upload_file"`
	AutoUploadFileMode      AutoUploadMode `json:"auto_upload_file_mode"`
	OSSProvider             OSSProvider    `json:"oss_provider"`
	AliyunOSSSettings       map[string]any `json:"aliyun_oss_settings"`
	TencentCloudOSSSettings map[string]any `json:"tencent_cloud_oss_settings"`
	CloudflareR2Settings    map[string]any `json:"cloudflare_r2_settings"`
}

type OSSSettings struct {
	ID                      uint64         `json:"id"`
	AutoUploadImage         *bool          `json:"auto_upload_image"`
	AutoUploadImageMode     AutoUploadMode `json:"auto_upload_image_mode"`
	AutoUploadVideo         *bool          `json:"auto_upload_video"`
	AutoUploadVideoMode     AutoUploadMode `json:"auto_upload_video_mode"`
	AutoUploadFile          *bool          `json:"auto_upload_file"`
	AutoUploadFileMode      AutoUploadMode `json:"auto_upload_file_mode"`
	OSSProvider             OSSProvider    `json:"oss_provider"`
	AliyunOSSSettings       map[string]any `json:"aliyun_oss_settings"`
	TencentCloudOSSSettings map[string]any `json:"tencent_cloud_oss_settings"`
	CloudflareR2Settings    map[string]any `json:"cloudflare_r2_settings"`
	CreatedAt               int64          `json:"created_at"`
	UpdatedAt               int64          `json:"updated_at"`
}
