package models

// EsAsGoods 售后服务商品信息表(es_as_goods)
type EsAsGoods struct {
	ID         int     `gorm:"primaryKey;column:id" json:"-"`         // 主键ID
	ServiceSn  string  `gorm:"column:service_sn" json:"service_sn"`   // 售后服务单号
	GoodsID    int     `gorm:"column:goods_id" json:"goods_id"`       // 商品ID
	SkuID      int     `gorm:"column:sku_id" json:"sku_id"`           // 商品SKUID
	ShipNum    int     `gorm:"column:ship_num" json:"ship_num"`       // 发货数量
	Price      float64 `gorm:"column:price" json:"price"`             // 商品成交价
	ReturnNum  int     `gorm:"column:return_num" json:"return_num"`   // 退还数量
	StorageNum int     `gorm:"column:storage_num" json:"storage_num"` // 入库数量
	GoodsName  string  `gorm:"column:goods_name" json:"goods_name"`   // 商品名称
	GoodsSn    string  `gorm:"column:goods_sn" json:"goods_sn"`       // 商品编号
	GoodsImage string  `gorm:"column:goods_image" json:"goods_image"` // 商品缩略图
	SpecJSON   string  `gorm:"column:spec_json" json:"spec_json"`     // 商品规格信息
}

// TableName get sql table name.获取数据库表名
func (m *EsAsGoods) TableName() string {
	return "es_as_goods"
}
