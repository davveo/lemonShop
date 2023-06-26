package models

// AsChange 售后服务用户收货地址信息表(es_as_change)
type AsChange struct {
	ID         int    `gorm:"primaryKey;column:id" json:"-"`         // 主键ID
	ServiceSn  string `gorm:"column:service_sn" json:"service_sn"`   // 售后服务单号
	ShipName   string `gorm:"column:ship_name" json:"ship_name"`     // 收货人姓名
	ProvinceID int    `gorm:"column:province_id" json:"province_id"` // 收货地址省份id
	CityID     int    `gorm:"column:city_id" json:"city_id"`         // 收货地址城市id
	CountyID   int    `gorm:"column:county_id" json:"county_id"`     // 收货地址区县id
	TownID     int    `gorm:"column:town_id" json:"town_id"`         // 收货地址城镇id
	Province   string `gorm:"column:province" json:"province"`       // 收货地址省份名称
	City       string `gorm:"column:city" json:"city"`               // 收货地址城市名称
	County     string `gorm:"column:county" json:"county"`           // 收货地址县(区)名称
	Town       string `gorm:"column:town" json:"town"`               // 收货地址城镇名称
	ShipAddr   string `gorm:"column:ship_addr" json:"ship_addr"`     // 收货地址
	ShipMobile string `gorm:"column:ship_mobile" json:"ship_mobile"` // 收货人手机号
}

// TableName get sql table name.获取数据库表名
func (m *AsChange) TableName() string {
	return "es_as_change"
}
