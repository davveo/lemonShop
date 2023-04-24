package models

// EsMemberCollectionShop 会员收藏店铺表(es_member_collection_shop)
type EsMemberCollectionShop struct {
	ID                    int     `gorm:"primaryKey;column:id" json:"-"`                                 // 收藏id
	MemberID              int     `gorm:"column:member_id" json:"member_id"`                             // 会员id
	ShopID                int     `gorm:"column:shop_id" json:"shop_id"`                                 // 店铺id
	ShopName              string  `gorm:"column:shop_name" json:"shop_name"`                             // 店铺名称
	CreateTime            int64   `gorm:"column:create_time" json:"create_time"`                         // 收藏时间
	Logo                  string  `gorm:"column:logo" json:"logo"`                                       // 店铺logo
	ShopProvince          string  `gorm:"column:shop_province" json:"shop_province"`                     // 店铺所在省
	ShopCity              string  `gorm:"column:shop_city" json:"shop_city"`                             // 店铺所在市
	ShopRegion            string  `gorm:"column:shop_region" json:"shop_region"`                         // 店铺所在县
	ShopTown              string  `gorm:"column:shop_town" json:"shop_town"`                             // 店铺所在镇
	ShopPraiseRate        float64 `gorm:"column:shop_praise_rate" json:"shop_praise_rate"`               // 店铺好评率
	ShopDescriptionCredit float64 `gorm:"column:shop_description_credit" json:"shop_description_credit"` // 店铺描述相符度
	ShopServiceCredit     float64 `gorm:"column:shop_service_credit" json:"shop_service_credit"`         // 服务态度分数
	ShopDeliveryCredit    float64 `gorm:"column:shop_delivery_credit" json:"shop_delivery_credit"`       // 发货速度分数
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberCollectionShop) TableName() string {
	return "es_member_collection_shop"
}
