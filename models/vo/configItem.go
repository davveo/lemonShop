package vo

type ConfigItem struct {
	Name       string        `json:"name"`
	Text       string        `json:"text"`
	ConfigType string        `json:"configType"`
	Value      interface{}   `json:"value"`
	Options    []RadioOption `json:"options"`
}
