package models

// EsMemberAddress 会员收货地址表(es_member_address)
type EsMemberAddress struct {
	AddrID          int    `gorm:"primaryKey;column:addr_id" json:"-"`                // 主键ID
	MemberID        int    `gorm:"column:member_id" json:"member_id"`                 // 会员ID
	Name            string `gorm:"column:name" json:"name"`                           // 收货人姓名
	Country         string `gorm:"column:country" json:"country"`                     // 收货人国籍
	ProvinceID      int    `gorm:"column:province_id" json:"province_id"`             // 所属省份ID
	CityID          int    `gorm:"column:city_id" json:"city_id"`                     // 所属城市ID
	CountyID        int    `gorm:"column:county_id" json:"county_id"`                 // 所属县(区)ID
	TownID          int    `gorm:"column:town_id" json:"town_id"`                     // 所属城镇ID
	County          string `gorm:"column:county" json:"county"`                       // 所属县(区)名称
	City            string `gorm:"column:city" json:"city"`                           // 所属城市名称
	Province        string `gorm:"column:province" json:"province"`                   // 所属省份名称
	Town            string `gorm:"column:town" json:"town"`                           // 所属城镇名称
	Addr            string `gorm:"column:addr" json:"addr"`                           // 详细地址
	Tel             string `gorm:"column:tel" json:"tel"`                             // 联系电话(一般指座机)
	Mobile          string `gorm:"column:mobile" json:"mobile"`                       // 手机号码
	DefAddr         int    `gorm:"column:def_addr" json:"def_addr"`                   // 是否为默认收货地址
	ShipAddressName string `gorm:"column:ship_address_name" json:"ship_address_name"` // 地址别名
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberAddress) TableName() string {
	return "es_member_address"
}
