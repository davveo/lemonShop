package vo

type SmsPlatformVo struct {
	Id          int          `json:"id"`
	Name        string       `json:"name"`
	Open        int          `json:"open"`
	Bean        string       `json:"bean"`
	Config      string       `json:"config"`
	ConfigItems []ConfigItem `json:"configItems"`
}
