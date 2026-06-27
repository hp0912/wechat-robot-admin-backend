package dto

type TimelineObject struct {
	ID                     uint64        `json:"ID"`
	Username               string        `json:"Username"`
	CreateTime             uint32        `json:"CreateTime"`
	ContentDesc            string        `json:"ContentDesc"`
	ContentDescShowType    uint32        `json:"ContentDescShowType"`
	ContentDescScene       uint32        `json:"ContentDescScene"`
	Private                uint32        `json:"Private"`
	SightFolded            uint32        `json:"SightFolded,omitempty"`
	ShowFlag               uint32        `json:"ShowFlag,omitempty"`
	ContentAttr            string        `json:"ContentAttr,omitempty"`
	SourceUserName         string        `json:"SourceUserName"`
	SourceNickName         string        `json:"SourceNickName"`
	PublicUserName         string        `json:"PublicUserName"`
	PublicBrandContactType uint32        `json:"PublicBrandContactType,omitempty"`
	StatisticsData         string        `json:"StatisticsData"`
	StatExtStr             string        `json:"StatExtStr,omitempty"`
	CanvasInfoXML          string        `json:"CanvasInfoXML,omitempty"`
	AppInfo                AppInfo       `json:"AppInfo"`
	WeappInfo              WeappInfo     `json:"WeappInfo,omitempty"`
	ContentObject          ContentObject `json:"ContentObject"`
	ActionInfo             ActionInfo    `json:"ActionInfo"`
	Location               Location      `json:"Location"`
	StreamVideo            StreamVideo   `json:"StreamVideo"`
}

type AppInfo struct {
	ID            string `json:"ID"`
	Version       string `json:"Version,omitempty"`
	AppName       string `json:"AppName,omitempty"`
	InstallUrl    string `json:"InstallUrl,omitempty"`
	FromUrl       string `json:"FromUrl,omitempty"`
	IsForceUpdate uint32 `json:"IsForceUpdate,omitempty"`
	IsHidden      uint32 `json:"IsHidden,omitempty"`
}

type WeappInfo struct {
	AppUserName      string `json:"AppUserName"`
	PagePath         string `json:"PagePath"`
	Version          string `json:"Version"`
	IsHidden         uint32 `json:"IsHidden"`
	DebugMode        string `json:"DebugMode"`
	ShareActionId    string `json:"ShareActionId"`
	IsGame           string `json:"IsGame"`
	MessageExtraData string `json:"MessageExtraData"`
	SubType          string `json:"SubType"`
	PreloadResources string `json:"PreloadResources"`
}

type ContentObject struct {
	ContentStyle    uint32    `json:"ContentStyle"`
	ContentSubStyle uint32    `json:"ContentSubStyle,omitempty"`
	Title           string    `json:"Title"`
	Description     string    `json:"Description"`
	ContentUrl      string    `json:"ContentUrl"`
	MediaList       MediaList `json:"MediaList"`
}

type MediaList struct {
	Media []Media `json:"Media"`
}

type Media struct {
	ID               uint64          `json:"ID"`
	IDStr            string          `json:"IDStr,omitempty"`
	Type             uint32          `json:"Type"`
	Title            string          `json:"Title"`
	Description      string          `json:"Description"`
	Private          uint32          `json:"Private"`
	UserData         string          `json:"UserData,omitempty"`
	SubType          uint32          `json:"SubType,omitempty"`
	VideoSize        *VideoSize      `json:"VideoSize,omitempty"`
	HD               *URL            `json:"HD,omitempty"`
	UHD              *URL            `json:"UHD,omitempty"`
	URL              URL             `json:"URL"`
	Thumb            Thumb           `json:"Thumb"`
	Size             Size            `json:"Size"`
	VideoDuration    float64         `json:"VideoDuration,omitempty"`
	VideoDurationStr string          `json:"VideoDurationStr,omitempty"`
	VideoColdDLRule  VideoColdDLRule `json:"VideoColdDLRule,omitempty"`
}

type VideoSize struct {
	Width  string `json:"Width"`
	Height string `json:"Height"`
}

type URL struct {
	Type     string `json:"Type"`
	MD5      string `json:"MD5"`
	VideoMD5 string `json:"VideoMD5"`
	Value    string `json:"Value"`
}

type Thumb struct {
	Type  string `json:"Type"`
	Value string `json:"Value"`
}

type Size struct {
	Width     string `json:"Width"`
	Height    string `json:"Height"`
	TotalSize string `json:"TotalSize"`
}

type VideoColdDLRule struct {
	All string `json:"All"`
}

type ActionInfo struct {
	AppMsg AppMsg `json:"AppMsg"`
}

type AppMsg struct {
	MediaTagName  string `json:"MediaTagName,omitempty"`
	MessageExt    string `json:"MessageExt,omitempty"`
	MessageAction string `json:"MessageAction"`
}

type Location struct {
	PoiClassifyId   string `json:"PoiClassifyId"`
	PoiName         string `json:"PoiName"`
	PoiAddress      string `json:"PoiAddress"`
	PoiClassifyType uint32 `json:"PoiClassifyType"`
	City            string `json:"City"`
}

type StreamVideo struct {
	StreamVideoUrl      string `json:"StreamVideoUrl"`
	StreamVideoThumbUrl string `json:"StreamVideoThumbUrl"`
	StreamVideoWebUrl   string `json:"StreamVideoWebUrl"`
}

type MomentsGetListRequest struct {
	ID           int64  `form:"id" json:"id"  binding:"required"`
	FristPageMd5 string `form:"frist_page_md5" json:"frist_page_md5"`
	MaxID        string `form:"max_id" json:"max_id" binding:"required"`
}

type MomentsGetListResponse struct {
	FirstPageMd5          *string            `json:"FirstPageMd5,omitempty"`
	ObjectCount           *uint32            `json:"ObjectCount,omitempty"`
	ObjectList            []*SnsObject       `json:"ObjectList,omitempty"`
	NewRequestTime        *uint32            `json:"NewRequestTime,omitempty"`
	ObjectCountForSameMd5 *uint32            `json:"ObjectCountForSameMd5,omitempty"`
	ControlFlag           *uint32            `json:"ControlFlag,omitempty"`
	ServerConfig          *SnsServerConfig   `json:"ServerConfig,omitempty"`
	AdvertiseCount        *uint32            `json:"AdvertiseCount,omitempty"`
	AdvertiseList         *string            `json:"AdvertiseList,omitempty"`
	Session               *SKBuiltinString_S `json:"Session,omitempty"`
	RecCount              *uint32            `json:"RecCount,omitempty"`
	RecList               *uint32            `json:"RecList,omitempty"`
}

type SnsObject struct {
	Id                   *uint64              `json:"Id,omitempty"`
	IdStr                string               `json:"IdStr,omitempty"`
	Username             *string              `json:"Username,omitempty"`
	Nickname             *string              `json:"Nickname,omitempty"`
	Avatar               *string              `json:"Avatar,omitempty"`
	CreateTime           *uint32              `json:"CreateTime,omitempty"`
	ObjectDesc           *SKBuiltinString_S   `json:"ObjectDesc,omitempty"`
	TimelineObject       *TimelineObject      `json:"TimelineObject,omitempty"`
	LikeFlag             *uint32              `json:"LikeFlag,omitempty"`
	LikeCount            *uint32              `json:"LikeCount,omitempty"`
	LikeUserListCount    *uint32              `json:"LikeUserListCount,omitempty"`
	LikeUserList         []*SnsCommentInfo    `json:"LikeUserList,omitempty"`
	CommentCount         *uint32              `json:"CommentCount,omitempty"`
	CommentUserListCount *uint32              `json:"CommentUserListCount,omitempty"`
	CommentUserList      []*SnsCommentInfo    `json:"CommentUserList,omitempty"`
	WithUserCount        *uint32              `json:"WithUserCount,omitempty"`
	WithUserListCount    *uint32              `json:"WithUserListCount,omitempty"`
	WithUserList         []*SnsCommentInfo    `json:"WithUserList,omitempty"`
	ExtFlag              *uint32              `json:"ExtFlag,omitempty"`
	NoChange             *uint32              `json:"NoChange,omitempty"`
	GroupCount           *uint32              `json:"GroupCount,omitempty"`
	GroupList            []*SnsGroup          `json:"GroupList,omitempty"`
	IsNotRichText        *uint32              `json:"IsNotRichText,omitempty"`
	ReferUsername        *string              `json:"ReferUsername,omitempty"`
	ReferId              *uint64              `json:"ReferId,omitempty"`
	BlackListCount       *uint32              `json:"BlackListCount,omitempty"`
	BlackList            []*SKBuiltinString_S `json:"BlackList,omitempty"`
	DeleteFlag           *uint32              `json:"DeleteFlag,omitempty"`
	GroupUserCount       *uint32              `json:"GroupUserCount,omitempty"`
	GroupUser            []*SKBuiltinString_S `json:"GroupUser,omitempty"`
	ObjectOperations     []*SKBuiltinString_S `json:"ObjectOperations,omitempty"`
	SnsRedEnvelops       *SnsRedEnvelops      `json:"SnsRedEnvelops,omitempty"`
	PreDownloadInfo      *PreDownloadInfo     `json:"PreDownloadInfo,omitempty"`
	WeAppInfo            *SnsWeAppInfo        `json:"WeAppInfo,omitempty"`
}

type SnsServerConfig struct {
	PostMentionLimit      *int32 `json:"PostMentionLimit,omitempty"`
	CopyAndPasteWordLimit *int32 `json:"CopyAndPasteWordLimit,omitempty"`
}

type SnsCommentInfo struct {
	Username        *string `json:"Username,omitempty"`
	Nickname        *string `json:"Nickname,omitempty"`
	Source          *uint32 `json:"Source,omitempty"`
	Type            *uint32 `json:"Type,omitempty"`
	Content         *string `json:"Content,omitempty"`
	CreateTime      *uint32 `json:"CreateTime,omitempty"`
	CommentId       *int32  `json:"CommentId,omitempty"`
	ReplyCommentId  *int32  `json:"ReplyCommentId,omitempty"`
	ReplyUsername   *string `json:"ReplyUsername,omitempty"`
	IsNotRichText   *uint32 `json:"IsNotRichText,omitempty"`
	ReplyCommentId2 *uint64 `json:"ReplyCommentId2,omitempty"`
	CommentId2      *uint64 `json:"CommentId2,omitempty"`
	DeleteFlag      *uint32 `json:"DeleteFlag,omitempty"`
	CommentFlag     *uint32 `json:"CommentFlag,omitempty"`
}

type SnsGroup struct {
	GroupId *uint64 `json:"GroupId,omitempty"`
}

type SnsRedEnvelops struct {
	RewardCount    *uint32 `json:"RewardCount,omitempty"`
	RewardUserList *string `json:"RewardUserList,omitempty"`
	ReportId       *uint32 `json:"ReportId,omitempty"`
	ReportKey      *uint32 `json:"ReportKey,omitempty"`
	ResourceId     *uint32 `json:"ResourceId,omitempty"`
}

type PreDownloadInfo struct {
	PreDownloadPercent *uint32 `json:"PreDownloadPercent,omitempty"`
	PreDownloadNetType *uint32 `json:"PreDownloadNetType,omitempty"`
	NoPreDownloadRange *string `json:"NoPreDownloadRange,omitempty"`
}

type SnsWeAppInfo struct {
	MapPoiId    *string `json:"MapPoiId,omitempty"`
	AppId       *uint32 `json:"AppId,omitempty"`
	UserName    *string `json:"UserName,omitempty"`
	RedirectUrl *string `json:"RedirectUrl,omitempty"`
	ShowType    *uint32 `json:"ShowType,omitempty"`
	RScore      *uint32 `json:"RScore,omitempty"`
}

type MomentsDownFriendCircleMediaRequest struct {
	ID  int64  `form:"id" json:"id"  binding:"required"`
	Url string `form:"url" json:"url" binding:"required"`
	Key string `form:"key" json:"key"`
}

type FriendCircleMedia struct {
	ClientId         *string         `json:"ClientId,omitempty"`
	BufferUrl        *SnsBufferUrl   `json:"BufferUrl,omitempty"`
	ThumbUrlCount    *uint32         `json:"ThumbUrlCount,omitempty"`
	ThumbUrls        []*SnsBufferUrl `json:"ThumbUrls,omitempty"`
	Id               *uint64         `json:"Id,omitempty"`
	IdStr            *string         `json:"IdStr,omitempty"`
	Type             *uint32         `json:"Type,omitempty"`
	Size             Size            `json:"Size"`
	VideoDuration    float64         `json:"VideoDuration,omitempty"`
	VideoDurationStr string          `json:"VideoDurationStr,omitempty"`
}

type SnsBufferUrl struct {
	Url  *string `json:"Url,omitempty"`
	Type *uint32 `json:"Type,omitempty"`
}

type MomentPostRequest struct {
	ID           int64               `form:"id" json:"id" binding:"required"`
	Content      string              `form:"content" json:"content"`
	MediaList    []FriendCircleMedia `form:"media_list" json:"media_list"`
	WithUserList []string            `form:"with_user_list" json:"with_user_list"`
	ShareType    string              `form:"share_type" json:"share_type" binding:"required"`
	ShareWith    []string            `form:"share_with" json:"share_with"`
	DoNotShare   []string            `form:"donot_share" json:"donot_share"`
}

type MomentPostResponse struct {
	SnsObject *SnsObject `json:"SnsObject,omitempty"`
	SpamTips  *string    `json:"SpamTips,omitempty"`
}

type MomentOpRequest struct {
	ID        int64  `form:"id" json:"id" binding:"required"`
	MomentID  string `form:"MomentID" json:"MomentID" binding:"required"`
	Type      uint32 `form:"Type" json:"Type" binding:"required"`
	CommentId uint32 `form:"CommentId" json:"CommentId"`
}

type MomentOpResponse struct {
	OpCount   *uint32 `json:"opCount,omitempty"`
	OpRetList []int32 `json:"opRetList,omitempty"`
}

type MomentPrivacySettingsRequest struct {
	ID       int64  `form:"id" json:"id" binding:"required"`
	Function uint32 `form:"Function" json:"Function"`
	Value    uint32 `form:"Value" json:"Value"`
}

type MomentPrivacySettingsResponse struct {
	Ret      *int32    `json:"ret,omitempty"`
	OplogRet *OplogRet `json:"oplogRet,omitempty"`
}

type OplogRet struct {
	Count  *uint32 `json:"count,omitempty"`
	Ret    []byte  `json:"ret,omitempty"`
	ErrMsg []byte  `json:"errMsg,omitempty"`
}

type FriendCircleCommentRequest struct {
	ID             int64  `form:"id" json:"id" binding:"required"`
	Type           uint32 `form:"Type" json:"Type" binding:"required"`
	MomentId       string `form:"MomentId" json:"MomentId" binding:"required"`
	ReplyCommnetId uint32 `form:"ReplyCommnetId" json:"ReplyCommnetId"`
	Content        string `form:"Content" json:"Content"`
}

type FriendCircleGetDetailRequest struct {
	ID           int64  `form:"id" json:"id" binding:"required"`
	Towxid       string `form:"Towxid" json:"Towxid" binding:"required"`
	Fristpagemd5 string `form:"Fristpagemd5" json:"Fristpagemd5"`
	Maxid        string `form:"Maxid" json:"Maxid"`
}

type FriendCircleGetIdDetailRequest struct {
	ID       int64  `form:"id" json:"id" binding:"required"`
	Towxid   string `form:"Towxid" json:"Towxid" binding:"required"`
	MomentId string `form:"MomentId" json:"MomentId" binding:"required"`
}

type SnsCommentResponse struct {
	SnsObject *SnsObject `json:"snsObject,omitempty"`
}

type SnsUserInfo struct {
	SnsFlag       *uint32 `protobuf:"varint,1,opt,name=SnsFlag" json:"SnsFlag,omitempty"`
	SnsBgimgId    *string `protobuf:"bytes,2,opt,name=SnsBgimgId" json:"SnsBgimgId,omitempty"`
	SnsBgobjectId *uint64 `protobuf:"varint,3,opt,name=SnsBgobjectId" json:"SnsBgobjectId,omitempty"`
	SnsFlagEx     *uint32 `protobuf:"varint,4,opt,name=SnsFlagEx" json:"SnsFlagEx,omitempty"`
}

type SnsUserPageResponse struct {
	FristPageMd5          *string            `json:"FristPageMd5,omitempty"`
	ObjectCount           *uint32            `json:"ObjectCount,omitempty"`
	ObjectList            []*SnsObject       `json:"ObjectList,omitempty"`
	ObjectTotalCount      []uint32           `json:"ObjectTotalCount,omitempty"`
	SnsUserInfo           []*SnsUserInfo     `json:"SnsUserInfo,omitempty"`
	NewRequestTime        []uint32           `json:"NewRequestTime,omitempty"`
	ObjectCountForSameMd5 []uint32           `json:"ObjectCountForSameMd5,omitempty"`
	ServerConfig          []*SnsServerConfig `json:"ServerConfig,omitempty"`
	LimitedId             []uint64           `json:"LimitedId,omitempty"`
	ContinueId            []uint64           `json:"ContinueId,omitempty"`
	RetTips               []string           `json:"RetTips,omitempty"`
}

type SnsObjectDetailResponse struct {
	Object *SnsObject `json:"object,omitempty"`
}

type MomentSettings struct {
	ID                      int64  `json:"id"`
	AutoLike                bool   `json:"auto_like"`
	AutoComment             bool   `json:"auto_comment"`
	Whitelist               string `json:"whitelist"`
	Blacklist               string `json:"blacklist"`
	AIBaseURL               string `json:"ai_base_url"`
	AIAPIKey                string `json:"ai_api_key"`
	WorkflowModel           string `json:"workflow_model"`
	CommentModel            string `json:"comment_model"`
	ImageUnderstandingModel string `json:"image_understanding_model"`
	VideoUnderstandingModel string `json:"video_understanding_model"`
	CommentPrompt           string `json:"comment_prompt"`
	MaxCompletionTokens     int    `json:"max_completion_tokens"`
	SyncInterval            int    `json:"sync_interval" binding:"min=10,max=86400"`
}
