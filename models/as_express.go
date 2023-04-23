package models

// EsAsExpress 售后服务用户退还商品物流信息表(es_as_express)
type EsAsExpress struct {
	ID               int    `gorm:"primaryKey;column:id" json:"-"`                       // 主键ID
	ServiceSn        string `gorm:"column:service_sn" json:"service_sn"`                 // 售后服务单号
	CourierNumber    string `gorm:"column:courier_number" json:"courier_number"`         // 物流单号
	CourierCompanyID int    `gorm:"column:courier_company_id" json:"courier_company_id"` // 物流公司id
	CourierCompany   string `gorm:"column:courier_company" json:"courier_company"`       // 物流公司名称
	ShipTime         int64  `gorm:"column:ship_time" json:"ship_time"`                   // 发货时间
}

// TableName get sql table name.获取数据库表名
func (m *EsAsExpress) TableName() string {
	return "es_as_express"
}
