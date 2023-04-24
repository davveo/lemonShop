package models

// EsPaymentMethod 支付方式(es_payment_method)
type EsPaymentMethod struct {
	MethodID        int    `gorm:"primaryKey;column:method_id" json:"-"`              // 支付方式id
	MethodName      string `gorm:"column:method_name" json:"method_name"`             // 支付方式名称
	PluginID        string `gorm:"column:plugin_id" json:"plugin_id"`                 // 支付插件名称
	PcConfig        string `gorm:"column:pc_config" json:"pc_config"`                 // pc是否可用
	WapConfig       string `gorm:"column:wap_config" json:"wap_config"`               // wap是否可用
	AppNativeConfig string `gorm:"column:app_native_config" json:"app_native_config"` // app 原生是否可用
	Image           string `gorm:"column:image" json:"image"`                         // 支付方式图片
	IsRetrace       int    `gorm:"column:is_retrace" json:"is_retrace"`               // 是否支持原路退回
	AppReactConfig  string `gorm:"column:app_react_config" json:"app_react_config"`   // app RN是否可用
	MiniConfig      string `gorm:"column:mini_config" json:"mini_config"`             // 小程序配置
}

// TableName get sql table name.获取数据库表名
func (m *EsPaymentMethod) TableName() string {
	return "es_payment_method"
}
