package vo

type SiteSetting struct {
	SiteName      string `json:"site_name"`       //站点名称
	Title         string `json:"title"`           // 网站标题
	Keywords      string `json:"keywords"`        // 网站关键字
	Description   string `json:"descript"`        // 网站描述
	SiteOn        int    `json:"siteon"`          // 网站是否开启，0开启，-1关闭
	CloseReason   string `json:"close_reson"`     // 关闭原因
	Logo          string `json:"logo"`            // 网站logo
	GlobalAuthKey string `json:"global_auth_key"` // 加密秘钥
	DefaultImg    string `json:"default_img"`     // 默认图片
	TestMode      int    `json:"test_mode"`       // 测试模式
}
