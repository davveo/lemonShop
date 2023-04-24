package models

// EsMemberCollectionGoods 会员商品收藏表(es_member_collection_goods)
type EsMemberCollectionGoods struct {
	ID         int     `gorm:"primaryKey;column:id" json:"-"`         // 主键ID
	MemberID   int     `gorm:"column:member_id" json:"member_id"`     // 会员ID
	GoodsID    int     `gorm:"column:goods_id" json:"goods_id"`       // 收藏商品ID
	CreateTime int64   `gorm:"column:create_time" json:"create_time"` // 收藏商品时间
	GoodsName  string  `gorm:"column:goods_name" json:"goods_name"`   // 商品名称
	GoodsPrice float64 `gorm:"column:goods_price" json:"goods_price"` // 商品价格
	GoodsSn    string  `gorm:"column:goods_sn" json:"goods_sn"`       // 商品编号
	GoodsImg   string  `gorm:"column:goods_img" json:"goods_img"`     // 商品图片
	ShopID     int     `gorm:"column:shop_id" json:"shop_id"`         // 店铺id
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberCollectionGoods) TableName() string {
	return "es_member_collection_goods"
}
