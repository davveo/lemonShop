package models

// EsDraftGoods 草稿商品(es_draft_goods)
type EsDraftGoods struct {
	DraftGoodsID        int     `gorm:"primaryKey;column:draft_goods_id" json:"-"`                 // 草稿商品id
	Quantity            int     `gorm:"column:quantity" json:"quantity"`                           // 商品总库存
	Original            string  `gorm:"column:original" json:"original"`                           // 商品原始图片
	SellerID            int     `gorm:"column:seller_id" json:"seller_id"`                         // 商品所属卖家ID
	ShopCatID           int     `gorm:"column:shop_cat_id" json:"shop_cat_id"`                     // 商品所属店铺类目ID
	TemplateID          int     `gorm:"column:template_id" json:"template_id"`                     // 商品运费模板ID
	GoodsTransfeeCharge int     `gorm:"column:goods_transfee_charge" json:"goods_transfee_charge"` // 是否为买家承担运费
	SellerName          string  `gorm:"column:seller_name" json:"seller_name"`                     // 商品所属店铺名称
	PageTitle           string  `gorm:"column:page_title" json:"page_title"`                       // seo 标题
	MetaKeywords        string  `gorm:"column:meta_keywords" json:"meta_keywords"`                 // seo关键字
	MetaDescription     string  `gorm:"column:meta_description" json:"meta_description"`           // seo描述
	CreateTime          int64   `gorm:"column:create_time" json:"create_time"`                     // 商品添加时间
	HaveSpec            int     `gorm:"column:have_spec" json:"have_spec"`                         // 是否开启规格
	Sn                  string  `gorm:"column:sn" json:"sn"`                                       // 商品编号
	GoodsName           string  `gorm:"column:goods_name" json:"goods_name"`                       // 商品名称
	BrandID             int     `gorm:"column:brand_id" json:"brand_id"`                           // 商品品牌ID
	CategoryID          int     `gorm:"column:category_id" json:"category_id"`                     // 商品分类ID
	Weight              float64 `gorm:"column:weight" json:"weight"`                               // 商品重量
	Intro               string  `gorm:"column:intro" json:"intro"`                                 // 商品详情
	Price               float64 `gorm:"column:price" json:"price"`                                 // 商品价格
	Cost                float64 `gorm:"column:cost" json:"cost"`                                   // 商品成本价
	Mktprice            float64 `gorm:"column:mktprice" json:"mktprice"`                           // 商品市场价
	GoodsType           string  `gorm:"column:goods_type" json:"goods_type"`                       // 商品类型
	ExchangeMoney       float64 `gorm:"column:exchange_money" json:"exchange_money"`               // 积分商品需要的金额
	ExchangePoint       int     `gorm:"column:exchange_point" json:"exchange_point"`               // 积分商品需要的积分
	ExchangeCategoryID  int     `gorm:"column:exchange_category_id" json:"exchange_category_id"`   // 积分商品的分类id
}

// TableName get sql table name.获取数据库表名
func (m *EsDraftGoods) TableName() string {
	return "es_draft_goods"
}
