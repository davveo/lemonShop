package models

// EsReceiptAddress 会员收票地址表(es_receipt_address)
type EsReceiptAddress struct {
	ID           int    `gorm:"primaryKey;column:id" json:"-"`             // 主键ID
	MemberID     int    `gorm:"column:member_id" json:"member_id"`         // 会员ID
	MemberName   string `gorm:"column:member_name" json:"member_name"`     // 收票人姓名
	MemberMobile string `gorm:"column:member_mobile" json:"member_mobile"` // 收票人手机号
	ProvinceID   int    `gorm:"column:province_id" json:"province_id"`     // 收票地址-所属省份ID
	CityID       int    `gorm:"column:city_id" json:"city_id"`             // 收票地址-所属城市ID
	CountyID     int    `gorm:"column:county_id" json:"county_id"`         // 收票地址-所属区县ID
	TownID       int    `gorm:"column:town_id" json:"town_id"`             // 收票地址-所属乡镇ID
	Province     string `gorm:"column:province" json:"province"`           // 收票地址-所属省份
	City         string `gorm:"column:city" json:"city"`                   // 收票地址-所属城市
	County       string `gorm:"column:county" json:"county"`               // 收票地址-所属区县
	Town         string `gorm:"column:town" json:"town"`                   // 收票地址-所属乡镇
	DetailAddr   string `gorm:"column:detail_addr" json:"detail_addr"`     // 收票地址-详细地址
}

// TableName get sql table name.获取数据库表名
func (m *EsReceiptAddress) TableName() string {
	return "es_receipt_address"
}