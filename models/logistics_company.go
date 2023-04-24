package models

// EsLogisticsCompany [...]
type EsLogisticsCompany struct {
	ID           int    `gorm:"primaryKey;column:id" json:"-"`
	Name         string `gorm:"column:name" json:"name"`
	Code         string `gorm:"column:code" json:"code"`
	IsWaybill    int    `gorm:"column:is_waybill" json:"is_waybill"`
	Kdcode       string `gorm:"column:kdcode" json:"kdcode"`
	FormItems    string `gorm:"column:form_items" json:"form_items"`
	DeleteStatus string `gorm:"column:delete_status" json:"delete_status"` // 是否删除 DELETED：已删除，NORMAL：正常
	Disabled     string `gorm:"column:disabled" json:"disabled"`           // 是否禁用 OPEN：开启，CLOSE：禁用
}

// TableName get sql table name.获取数据库表名
func (m *EsLogisticsCompany) TableName() string {
	return "es_logistics_company"
}
