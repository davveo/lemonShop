package models

// EsGoodsGallery 商品相册(es_goods_gallery)
type EsGoodsGallery struct {
	ImgID     int    `gorm:"primaryKey;column:img_id" json:"-"` // 主键
	GoodsID   int    `gorm:"column:goods_id" json:"goods_id"`   // 商品主键
	Thumbnail string `gorm:"column:thumbnail" json:"thumbnail"` // 缩略图路径
	Small     string `gorm:"column:small" json:"small"`         // 小图路径
	Big       string `gorm:"column:big" json:"big"`             // 大图路径
	Original  string `gorm:"column:original" json:"original"`   // 原图路径
	Tiny      string `gorm:"column:tiny" json:"tiny"`           // 极小图路径
	Isdefault int    `gorm:"column:isdefault" json:"isdefault"` // 是否是默认图片1   0没有默认
	Sort      int    `gorm:"column:sort" json:"sort"`           // 排序
}

// TableName get sql table name.获取数据库表名
func (m *EsGoodsGallery) TableName() string {
	return "es_goods_gallery"
}
