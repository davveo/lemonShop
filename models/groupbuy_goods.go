package models

// EsGroupbuyGoods 团购商品(es_groupbuy_goods)
type EsGroupbuyGoods struct {
	GbID          int     `gorm:"primaryKey;column:gb_id" json:"-"`            // 团购商品Id
	SkuID         int     `gorm:"column:sku_id" json:"sku_id"`                 // 货品Id
	ActID         int     `gorm:"column:act_id" json:"act_id"`                 // 活动Id
	CatID         int     `gorm:"column:cat_id" json:"cat_id"`                 // cat_id
	AreaID        int     `gorm:"column:area_id" json:"area_id"`               // 地区Id
	GbName        string  `gorm:"column:gb_name" json:"gb_name"`               // 团购名称
	GbTitle       string  `gorm:"column:gb_title" json:"gb_title"`             // 副标题
	GoodsName     string  `gorm:"column:goods_name" json:"goods_name"`         // 商品名称
	GoodsID       int     `gorm:"column:goods_id" json:"goods_id"`             // 商品Id
	OriginalPrice float64 `gorm:"column:original_price" json:"original_price"` // 原始价格
	Price         float64 `gorm:"column:price" json:"price"`                   // 团购价格
	ImgURL        string  `gorm:"column:img_url" json:"img_url"`               // 图片地址
	GoodsNum      int     `gorm:"column:goods_num" json:"goods_num"`           // 商品总数
	VisualNum     int     `gorm:"column:visual_num" json:"visual_num"`         // 虚拟数量
	LimitNum      int     `gorm:"column:limit_num" json:"limit_num"`           // 限购数量
	BuyNum        int     `gorm:"column:buy_num" json:"buy_num"`               // 已团购数量
	ViewNum       int     `gorm:"column:view_num" json:"view_num"`             // 浏览数量
	Remark        string  `gorm:"column:remark" json:"remark"`                 // 介绍
	GbStatus      int     `gorm:"column:gb_status" json:"gb_status"`           // 状态
	AddTime       int64   `gorm:"column:add_time" json:"add_time"`             // 添加时间
	WapThumbnail  string  `gorm:"column:wap_thumbnail" json:"wap_thumbnail"`   // wap缩略图
	Thumbnail     string  `gorm:"column:thumbnail" json:"thumbnail"`           // pc缩略图
	SellerID      int     `gorm:"column:seller_id" json:"seller_id"`           // 商家ID
	SellerName    string  `gorm:"column:seller_name" json:"seller_name"`       // 店铺名称
}

// TableName get sql table name.获取数据库表名
func (m *EsGroupbuyGoods) TableName() string {
	return "es_groupbuy_goods"
}
