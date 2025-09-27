package dto

type TrustSoftData struct {
	SoftConfig *string `json:"softConfig,omitempty"`
	SoftData   string  `json:"softData,omitempty"`
}

type TrustResponseData struct {
	SoftData    *TrustSoftData `json:"softData,omitempty"`
	DeviceToken *string        `json:"deviceToken,omitempty"`
	TimeStamp   *uint32        `json:"timeStamp,omitempty"`
}

type TrustResponse struct {
	BaseResponse      *BaseResponse      `json:"BaseResponse,omitempty"`
	TrustResponseData *TrustResponseData `json:"TrustResponseData,omitempty"`
}

type LoginDataInfo struct {
	Type     byte
	UserName string
	PassWord string
	//伪密码
	NewPassWord string
	//登录数据 62/A16
	LoginData string
	Ticket    string
	NewType   int
	Language  string
}

type DeviceInfo struct {
	UUIDOne            string `json:"uuidone"`
	UUIDTwo            string `json:"uuidtwo"`
	Imei               string `json:"imei"`
	DeviceID           string `json:"deviceid"`
	DeviceName         string `json:"devicename"`
	DeviceMac          string `json:"Devicemac"`
	TimeZone           string `json:"timezone"`
	Language           string `json:"language"`
	DeviceBrand        string `json:"devicebrand"`
	RealCountry        string `json:"realcountry"`
	IphoneVer          string `json:"iphonever"`
	BundleID           string `json:"boudleid"`
	OsType             string `json:"ostype"`
	AdSource           string `json:"adsource"`
	OsTypeNumber       string `json:"ostypenumber"`
	CoreCount          uint32 `json:"corecount"`
	CarrierName        string `json:"carriername"`
	SoftTypeXML        string `json:"softtypexml"`
	ClientCheckDataXML string `json:"clientcheckdataxml"`
	// extInfo
	GUID1 string `json:"guid1"` //data   path  uuid
	GUID2 string `json:"guid2"` //bundle path  uuid
	Sdi   string `json:"sdi"`   //md5(uuid)

	InstallTime  uint64 `json:"installtime"`  //random
	KernBootTime uint64 `json:"kernboottime"` //random

	Sysverplist *Stat `json:"sysverplist"` ////System/Library/CoreServices/SystemVersion.plist
	Dyldcache   *Stat `json:"dyldcache"`   ///System/Library/Caches/com.apple.dyld/dyld_shared_cache_arm64
	Var         *Stat `json:"var"`         ///private/var
	Etcgroup    *Stat `json:"etcgroup"`    ///private/etc/group
	Etchosts    *Stat `json:"etchosts"`    ///private/etc/hosts

	Apfs *Statfs `json:"apfs"` //apfs

	DeviceToken TrustResponse
}

type AndroidDeviceInfo struct {
	Imei                string
	AndriodId           string
	PhoneSerial         string
	WidevineDeviceID    string
	WidevineProvisionID string
	AndriodFsId         string
	AndriodBssId        string
	AndriodSsId         string
	WLanAddress         string
	PackageSign         string
	Androidversion      string
	RadioVersion        string
	Manufacturer        string
	BuildID             string
	BuildFP             string
	BuildBoard          string
	PhoneModel          string
	Hardware            string
	Features            string
	WifiName            string
	WifiFullName        string
	KernelReleaseNumber string
	Arch                string
	SfMD5               string
	SfArmMD5            string
	SfArm64MD5          string
	SbMD5               string
}

type RobotLoginData struct {
	Uin                        uint32
	Wxid                       string
	Pwd                        string
	Uuid                       string
	Aeskey                     string
	NotifyKey                  string
	Deviceid_str               string
	Deviceid_byte              string
	DeviceType                 string
	ClientVersion              int32
	DeviceName                 string
	NickName                   string
	HeadUrl                    string
	Email                      string
	Alais                      string
	Mobile                     string
	Mmtlsip                    string
	ShortHost                  string
	LongHost                   string
	Sessionkey                 string
	Sessionkey_2               string
	Autoauthkey                string
	Autoauthkeylen             int32
	Clientsessionkey           string
	Serversessionkey           string
	HybridEcdhPrivkey          string
	HybridEcdhPubkey           string
	HybridEcdhInitServerPubKey string
	Loginecdhkey               string
	Cooike                     string
	LoginMode                  string
	Proxy                      ProxyInfo
	MmtlsKey                   MmtlsClient
	DeviceToken                TrustResponse
	SyncKey                    string
	Data62                     string
	RomModel                   string
	Imei                       string
	SoftType                   string
	OsVersion                  string
	RsaPublicKey               string
	RsaPrivateKey              string
	Dns                        []Dns
	// 登录的Rsa 密钥版本
	LoginRsaVer uint32
	// 是否开启服务
	EnableService bool
	EcPublicKey   string `json:"ecpukey"`
	EcPrivateKey  string `json:"ecprkey"`
	Ticket        string
	LoginDataInfo LoginDataInfo
	// 设备信息62
	DeviceInfo *DeviceInfo
	//A16信息
	DeviceInfoA16 *AndroidDeviceInfo
	// 登录时间
	LoginDate int64
	// 刷新 tonken 时间
	RefreshTokenDate int64
}

type ProxyInfo struct {
	ProxyIp       string
	ProxyUser     string
	ProxyPassword string
}

type BaseResponse struct {
	Ret    *int32            `json:"ret,omitempty"`
	ErrMsg *SKBuiltinStringT `json:"errMsg,omitempty"`
}

type Dns struct {
	Ip   string
	Host string
}

type Timespec struct {
	Tvsec  uint64 `json:"tv_sec"`
	Tvnsec uint64 `json:"tv_nsec"`
}

type Statfs struct {
	Type        uint64 `json:"type"`        // f_type = 26
	Fstypename  string `json:"fstypename"`  //apfs
	Flags       uint64 `json:"flags"`       //statfs f_flags =  1417728009
	Mntonname   string `json:"mntonname"`   // /
	Mntfromname string `json:"mntfromname"` // com.apple.os.update-{%{96}s}@/dev/disk0s1s1
	Fsid        uint64 `json:"fsid"`        // f_fsid[0]
}

type Stat struct {
	Inode   uint64   `json:"inode"`
	Statime Timespec `json:"st_atime"`
	Stmtime Timespec `json:"st_mtime"`
	Stctime Timespec `json:"st_ctime"`
	Stbtime Timespec `json:"st_btime"`
}

type MmtlsClient struct {
	Shakehandpubkey    string
	Shakehandpubkeylen int32
	Shakehandprikey    string
	Shakehandprikeylen int32

	Shakehandpubkey_2   string
	Shakehandpubkeylen2 int32
	Shakehandprikey_2   string
	Shakehandprikeylen2 int32

	Mserverpubhashs     string
	ServerSeq           int
	ClientSeq           int
	ShakehandECDHkey    string
	ShakehandECDHkeyLen int32

	Encrptmmtlskey  string
	Decryptmmtlskey string
	EncrptmmtlsIv   string
	DecryptmmtlsIv  string

	CurDecryptSeqIv string
	CurEncryptSeqIv string

	Decrypt_part2_hash256            string
	Decrypt_part3_hash256            string
	ShakehandECDHkeyhash             string
	Hkdfexpand_pskaccess_key         string
	Hkdfexpand_pskrefresh_key        string
	HkdfExpand_info_serverfinish_key string
	Hkdfexpand_clientfinish_key      string
	Hkdfexpand_secret_key            string

	Hkdfexpand_application_key string
	Encrptmmtlsapplicationkey  string
	Decryptmmtlsapplicationkey string
	EncrptmmtlsapplicationIv   string
	DecryptmmtlsapplicationIv  string

	Earlydatapart       string
	Newsendbufferhashs  string
	Encrptshortmmtlskey string
	Encrptshortmmtlsiv  string
	Decrptshortmmtlskey string
	Decrptshortmmtlsiv  string

	//http才需要
	Pskkey    string
	Pskiv     string
	MmtlsMode uint
}
