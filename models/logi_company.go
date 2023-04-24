package models

// EsLogiCompany 物流公司(es_logi_company)
type EsLogiCompany struct {
	ID           int    `gorm:"primaryKey;column:id" json:"-"`             // 物流公司id
	Name         string `gorm:"column:name" json:"name"`                   // 物流公司名称
	Code         string `gorm:"column:code" json:"code"`                   // 物流公司code
	Kdcode       string `gorm:"column:kdcode" json:"kdcode"`               // 快递鸟物流公司code
	IsWaybill    int    `gorm:"column:is_waybill" json:"is_waybill"`       // 是否支持电子面单1：支持 0：不支持
	CustomerName string `gorm:"column:customer_name" json:"customer_name"` // 物流公司客户号
	CustomerPwd  string `gorm:"column:customer_pwd" json:"customer_pwd"`   // 物流公司电子面单密码
}

// TableName get sql table name.获取数据库表名
func (m *EsLogiCompany) TableName() string {
	return "es_logi_company"
}
