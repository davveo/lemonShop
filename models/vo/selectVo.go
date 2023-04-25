package vo

type SelectVo struct {
	Id       int64  `json:"id"`
	Text     string `json:"text"`
	Selected bool   `json:"selected"`
}
