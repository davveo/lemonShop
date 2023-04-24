package models

// EsFullDiscountGift 满优惠赠品(es_full_discount_gift)
type EsFullDiscountGift struct {
	GiftID      int     `gorm:"primaryKey;column:gift_id" json:"-"`      // 赠品id
	GiftName    string  `gorm:"column:gift_name" json:"gift_name"`       // 赠品名称
	GiftPrice   float64 `gorm:"column:gift_price" json:"gift_price"`     // 赠品金额
	GiftImg     string  `gorm:"column:gift_img" json:"gift_img"`         // 赠品图片
	GiftType    int     `gorm:"column:gift_type" json:"gift_type"`       // 赠品类型
	ActualStore int     `gorm:"column:actual_store" json:"actual_store"` // 库存
	EnableStore int     `gorm:"column:enable_store" json:"enable_store"` // 可用库存
	CreateTime  int64   `gorm:"column:create_time" json:"create_time"`   // 活动时间
	GoodsID     int     `gorm:"column:goods_id" json:"goods_id"`         // 活动商品id
	Disabled    int     `gorm:"column:disabled" json:"disabled"`         // 是否禁用
	SellerID    int     `gorm:"column:seller_id" json:"seller_id"`       // 店铺id
}

// TableName get sql table name.获取数据库表名
func (m *EsFullDiscountGift) TableName() string {
	return "es_full_discount_gift"
}

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

// EsGoodsParams 商品关联参数值(es_goods_params)
type EsGoodsParams struct {
	ID         int    `gorm:"primaryKey;column:id" json:"-"`         // 主键
	GoodsID    int    `gorm:"column:goods_id" json:"goods_id"`       // 商品id
	ParamID    int    `gorm:"column:param_id" json:"param_id"`       // 参数id
	ParamName  string `gorm:"column:param_name" json:"param_name"`   // 参数名字
	ParamValue string `gorm:"column:param_value" json:"param_value"` // 参数值
}

// TableName get sql table name.获取数据库表名
func (m *EsGoodsParams) TableName() string {
	return "es_goods_params"
}

// EsGoodsSku 商品sku(es_goods_sku)
type EsGoodsSku struct {
	SkuID          int     `gorm:"primaryKey;column:sku_id" json:"-"`             // 主键
	GoodsID        int     `gorm:"column:goods_id" json:"goods_id"`               // 商品id
	GoodsName      string  `gorm:"column:goods_name" json:"goods_name"`           // 商品名称
	Sn             string  `gorm:"column:sn" json:"sn"`                           // 商品编号
	Quantity       int     `gorm:"column:quantity" json:"quantity"`               // 库存
	EnableQuantity int     `gorm:"column:enable_quantity" json:"enable_quantity"` // 可用库存
	Price          float64 `gorm:"column:price" json:"price"`                     // 商品价格
	Specs          string  `gorm:"column:specs" json:"specs"`                     // 规格信息json
	Cost           float64 `gorm:"column:cost" json:"cost"`                       // 成本价格
	Weight         float64 `gorm:"column:weight" json:"weight"`                   // 重量
	SellerID       int     `gorm:"column:seller_id" json:"seller_id"`             // 卖家id
	SellerName     string  `gorm:"column:seller_name" json:"seller_name"`         // 卖家名称
	CategoryID     int     `gorm:"column:category_id" json:"category_id"`         // 分类id
	Thumbnail      string  `gorm:"column:thumbnail" json:"thumbnail"`             // 缩略图
	HashCode       int     `gorm:"column:hash_code" json:"hash_code"`             // 标识规格唯一性
}

// TableName get sql table name.获取数据库表名
func (m *EsGoodsSku) TableName() string {
	return "es_goods_sku"
}

// EsGoodsWords 商品分词表(es_goods_words)
type EsGoodsWords struct {
	ID       int    `gorm:"primaryKey;column:id" json:"-"`     // 主键
	Words    string `gorm:"column:words" json:"words"`         // 分词
	GoodsNum int64  `gorm:"column:goods_num" json:"goods_num"` // 数量
	Quanpin  string `gorm:"column:quanpin" json:"quanpin"`     // 全拼字母
	Szm      string `gorm:"column:szm" json:"szm"`             // 首字母
	Type     string `gorm:"column:type" json:"type"`           // 类型，系统(SYSTEM)，平台(PLATFORM)
	Sort     int    `gorm:"column:sort" json:"sort"`           // 排序
}

// TableName get sql table name.获取数据库表名
func (m *EsGoodsWords) TableName() string {
	return "es_goods_words"
}

// EsGroupbuyActive 团购活动表(es_groupbuy_active)
type EsGroupbuyActive struct {
	ActID        int    `gorm:"primaryKey;column:act_id" json:"-"`         // 活动Id
	ActName      string `gorm:"column:act_name" json:"act_name"`           // 活动名称
	StartTime    int64  `gorm:"column:start_time" json:"start_time"`       // 活动开启时间
	EndTime      int64  `gorm:"column:end_time" json:"end_time"`           // 团购结束时间
	JoinEndTime  int64  `gorm:"column:join_end_time" json:"join_end_time"` // 团购报名截止时间
	AddTime      int64  `gorm:"column:add_time" json:"add_time"`           // 团购添加时间
	ActTagID     int    `gorm:"column:act_tag_id" json:"act_tag_id"`       // 团购活动标签Id
	GoodsNum     int    `gorm:"column:goods_num" json:"goods_num"`         // 参与团购商品数量
	DeleteReason string `gorm:"column:delete_reason" json:"delete_reason"` // 删除原因
	DeleteTime   int64  `gorm:"column:delete_time" json:"delete_time"`     // 删除时间
	DeleteName   string `gorm:"column:delete_name" json:"delete_name"`     // 删除人
	DeleteStatus string `gorm:"column:delete_status" json:"delete_status"` // 是否删除 DELETED：已删除，NORMAL：正常
}

// TableName get sql table name.获取数据库表名
func (m *EsGroupbuyActive) TableName() string {
	return "es_groupbuy_active"
}

// EsGroupbuyArea 团购地区(es_groupbuy_area)
type EsGroupbuyArea struct {
	AreaID    int    `gorm:"primaryKey;column:area_id" json:"-"`  // 地区ID
	AreaName  string `gorm:"column:area_name" json:"area_name"`   // 地区名称
	AreaOrder int    `gorm:"column:area_order" json:"area_order"` // 地区排序
	AreaPath  string `gorm:"column:area_path" json:"area_path"`   // 地区路径结构
	ParentID  int    `gorm:"column:parent_id" json:"parent_id"`   // 地区父节点
}

// TableName get sql table name.获取数据库表名
func (m *EsGroupbuyArea) TableName() string {
	return "es_groupbuy_area"
}

// EsGroupbuyCat 团购分类(es_groupbuy_cat)
type EsGroupbuyCat struct {
	CatID    int    `gorm:"primaryKey;column:cat_id" json:"-"` // 分类id
	ParentID int    `gorm:"column:parent_id" json:"parent_id"` // 父类id
	CatName  string `gorm:"column:cat_name" json:"cat_name"`   // 分类名称
	CatPath  string `gorm:"column:cat_path" json:"cat_path"`   // 分类结构目录
	CatOrder int    `gorm:"column:cat_order" json:"cat_order"` // 分类排序
}

// TableName get sql table name.获取数据库表名
func (m *EsGroupbuyCat) TableName() string {
	return "es_groupbuy_cat"
}

// EsGroupbuyQuantityLog 团购商品库存日志表(es_groupbuy_quantity_log)
type EsGroupbuyQuantityLog struct {
	LogID    int    `gorm:"primaryKey;column:log_id" json:"-"` // 日志id
	OrderSn  string `gorm:"column:order_sn" json:"order_sn"`   // 订单编号
	GoodsID  int    `gorm:"column:goods_id" json:"goods_id"`   // 商品ID
	Quantity int    `gorm:"column:quantity" json:"quantity"`   // 团购售空数量
	OpTime   int64  `gorm:"column:op_time" json:"op_time"`     // 操作时间
	LogType  string `gorm:"column:log_type" json:"log_type"`   // 日志类型
	Reason   string `gorm:"column:reason" json:"reason"`       // 操作原因
	GbID     int    `gorm:"column:gb_id" json:"gb_id"`         // 团购商品id
}

// TableName get sql table name.获取数据库表名
func (m *EsGroupbuyQuantityLog) TableName() string {
	return "es_groupbuy_quantity_log"
}

// EsHalfPrice 第二件半价(es_half_price)
type EsHalfPrice struct {
	HpID        int    `gorm:"primaryKey;column:hp_id" json:"-"`      // 第二件半价活动ID
	StartTime   int64  `gorm:"column:start_time" json:"start_time"`   // 活动开始时间
	EndTime     int64  `gorm:"column:end_time" json:"end_time"`       // 活动结束时间
	Title       string `gorm:"column:title" json:"title"`             // 活动标题
	RangeType   int    `gorm:"column:range_type" json:"range_type"`   // 是否全部商品参与
	Disabled    int    `gorm:"column:disabled" json:"disabled"`       // 是否停用
	Description string `gorm:"column:description" json:"description"` // 活动说明
	SellerID    int    `gorm:"column:seller_id" json:"seller_id"`     // 商家id
}

// TableName get sql table name.获取数据库表名
func (m *EsHalfPrice) TableName() string {
	return "es_half_price"
}

// EsHistory [...]
type EsHistory struct {
	ID         int     `gorm:"primaryKey;column:id" json:"-"`         // 主键
	GoodsID    int     `gorm:"column:goods_id" json:"goods_id"`       // 商品id
	GoodsName  string  `gorm:"column:goods_name" json:"goods_name"`   // 商品名称
	GoodsPrice float64 `gorm:"column:goods_price" json:"goods_price"` // 商品价格
	GoodsImg   string  `gorm:"column:goods_img" json:"goods_img"`     // 商品主图
	MemberID   int     `gorm:"column:member_id" json:"member_id"`     // 会员id
	MemberName string  `gorm:"column:member_name" json:"member_name"` // 会员名称
	CreateTime int64   `gorm:"column:create_time" json:"create_time"` // 创建时间，按天存
	UpdateTime int64   `gorm:"column:update_time" json:"update_time"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *EsHistory) TableName() string {
	return "es_history"
}

// EsHotKeyword 热门关键字(es_hot_keyword)
type EsHotKeyword struct {
	ID      int    `gorm:"primaryKey;column:id" json:"-"`   // 主键ID
	HotName string `gorm:"column:hot_name" json:"hot_name"` // 文字内容
	Sort    int    `gorm:"column:sort" json:"sort"`         // 排序
}

// TableName get sql table name.获取数据库表名
func (m *EsHotKeyword) TableName() string {
	return "es_hot_keyword"
}

// EsKeywordSearchHistory [...]
type EsKeywordSearchHistory struct {
	ID         int    `gorm:"primaryKey;column:id" json:"-"`
	Keyword    string `gorm:"column:keyword" json:"keyword"`         // 关键字
	Count      int    `gorm:"column:count" json:"count"`             // 搜索次数
	AddTime    int64  `gorm:"column:add_time" json:"add_time"`       // 新增时间
	ModifyTime int64  `gorm:"column:modify_time" json:"modify_time"` // 更新时间
}

// TableName get sql table name.获取数据库表名
func (m *EsKeywordSearchHistory) TableName() string {
	return "es_keyword_search_history"
}

// EsLogiCompany 物流公司(es_logi_company)
type EsLogiCompany struct {
	ID           int    `gorm:"primaryKey;column:id" json:"-"`             // 物流公司id
	Name         string `gorm:"column:name" json:"name"`                   // 物流公司名称
	Code         string `gorm:"column:code" json:"code"`                   // 物流公司code
	Kdcode       string `gorm:"column:kdcode" json:"kdcode"`               // 快递鸟物流公司code
	IsWaybill    int    `gorm:"column:is_waybill" json:"is_waybill"`       // 是否支持电子面单1：支持 0：不支持
	CustomerName string `gorm:"column:customer_name" json:"customer_name"` // 物流公司客户号
	CustomerPwd  string `gorm:"column:customer_pwd" json:"customer_pwd"`   // 物流公司电子面单密码
}

// TableName get sql table name.获取数据库表名
func (m *EsLogiCompany) TableName() string {
	return "es_logi_company"
}

// EsLogisticsCompany [...]
type EsLogisticsCompany struct {
	ID           int    `gorm:"primaryKey;column:id" json:"-"`
	Name         string `gorm:"column:name" json:"name"`
	Code         string `gorm:"column:code" json:"code"`
	IsWaybill    int    `gorm:"column:is_waybill" json:"is_waybill"`
	Kdcode       string `gorm:"column:kdcode" json:"kdcode"`
	FormItems    string `gorm:"column:form_items" json:"form_items"`
	DeleteStatus string `gorm:"column:delete_status" json:"delete_status"` // 是否删除 DELETED：已删除，NORMAL：正常
	Disabled     string `gorm:"column:disabled" json:"disabled"`           // 是否禁用 OPEN：开启，CLOSE：禁用
}

// TableName get sql table name.获取数据库表名
func (m *EsLogisticsCompany) TableName() string {
	return "es_logistics_company"
}

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

// EsMemberAsk 咨询(es_member_ask)
type EsMemberAsk struct {
	AskID       int    `gorm:"primaryKey;column:ask_id" json:"-"`       // 主键
	GoodsID     int    `gorm:"column:goods_id" json:"goods_id"`         // 商品id
	MemberID    int    `gorm:"column:member_id" json:"member_id"`       // 会员id
	Content     string `gorm:"column:content" json:"content"`           // 咨询内容
	CreateTime  int64  `gorm:"column:create_time" json:"create_time"`   // 咨询时间
	SellerID    int    `gorm:"column:seller_id" json:"seller_id"`       // 卖家id
	Reply       string `gorm:"column:reply" json:"reply"`               // 卖家回复内容
	ReplyTime   int64  `gorm:"column:reply_time" json:"reply_time"`     // 卖家回复时间
	ReplyStatus string `gorm:"column:reply_status" json:"reply_status"` // 商家回复状态 YES:已回复,NO:未回复
	Status      string `gorm:"column:status" json:"status"`             // 删除状态 DELETED:已删除,NORMAL:正常
	MemberName  string `gorm:"column:member_name" json:"member_name"`   // 会员名称
	GoodsName   string `gorm:"column:goods_name" json:"goods_name"`     // 商品名称
	MemberFace  string `gorm:"column:member_face" json:"member_face"`   // 会员头像
	AuthStatus  string `gorm:"column:auth_status" json:"auth_status"`   // 审核状态 WAIT_AUDIT:待审核,PASS_AUDIT:审核通过,REFUSE_AUDIT:审核未通过
	GoodsImg    string `gorm:"column:goods_img" json:"goods_img"`       // 商品图片
	Anonymous   string `gorm:"column:anonymous" json:"anonymous"`       // 是否匿名 YES:是，NO:否
	ReplyNum    int    `gorm:"column:reply_num" json:"reply_num"`       // 会员问题咨询回复数量
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberAsk) TableName() string {
	return "es_member_ask"
}

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

// EsMemberComment 评论(es_member_comment)
type EsMemberComment struct {
	CommentID           int    `gorm:"primaryKey;column:comment_id" json:"-"`                     // 评论主键
	GoodsID             int    `gorm:"column:goods_id" json:"goods_id"`                           // 商品id
	SkuID               int    `gorm:"column:sku_id" json:"sku_id"`                               // skuid
	MemberID            int    `gorm:"column:member_id" json:"member_id"`                         // 会员id
	SellerID            int    `gorm:"column:seller_id" json:"seller_id"`                         // 卖家id
	MemberName          string `gorm:"column:member_name" json:"member_name"`                     // 会员名称
	MemberFace          string `gorm:"column:member_face" json:"member_face"`                     // 会员头像
	GoodsName           string `gorm:"column:goods_name" json:"goods_name"`                       // 商品名称
	GoodsImg            string `gorm:"column:goods_img" json:"goods_img"`                         // 商品图片
	Content             string `gorm:"column:content" json:"content"`                             // 评论内容
	CreateTime          int64  `gorm:"column:create_time" json:"create_time"`                     // 评论时间
	HaveImage           int16  `gorm:"column:have_image" json:"have_image"`                       // 是否有图片 1 有 0 没有
	Status              int16  `gorm:"column:status" json:"status"`                               // 状态  1 正常 0 删除
	Grade               string `gorm:"column:grade" json:"grade"`                                 // 好中差评
	OrderSn             string `gorm:"column:order_sn" json:"order_sn"`                           // 订单编号
	ReplyStatus         int16  `gorm:"column:reply_status" json:"reply_status"`                   // 是否回复 1 是 0 否
	AuditStatus         string `gorm:"column:audit_status" json:"audit_status"`                   // 初评审核:待审核(WAIT_AUDIT),审核通过(PASS_AUDIT),审核拒绝(REFUSE_AUDIT)
	CommentsType        string `gorm:"column:comments_type" json:"comments_type"`                 // 评论类型：初评(INITIAL),追评(ADDITIONAL)
	ParentID            int    `gorm:"column:parent_id" json:"parent_id"`                         // 初评id，默认为0
	AdditionalStatus    int    `gorm:"column:additional_status" json:"additional_status"`         // 追加评论状态 0：未追加，1：已追加
	AdditionalContent   string `gorm:"column:additional_content" json:"additional_content"`       // 追加的评论内容
	AdditionalTime      int64  `gorm:"column:additional_time" json:"additional_time"`             // 追加评论时间
	AdditionalHaveImage int    `gorm:"column:additional_have_image" json:"additional_have_image"` // 追加的评论是否上传了图片 0：未上传，1：已上传
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberComment) TableName() string {
	return "es_member_comment"
}

// EsMemberCoupon [...]
type EsMemberCoupon struct {
	McID       int    `gorm:"primaryKey;column:mc_id" json:"-"`
	CouponID   int    `gorm:"column:coupon_id" json:"coupon_id"`
	MemberID   int    `gorm:"column:member_id" json:"member_id"`
	UsedTime   int64  `gorm:"column:used_time" json:"used_time"`
	CreateTime int64  `gorm:"column:create_time" json:"create_time"`
	OrderID    int    `gorm:"column:order_id" json:"order_id"`
	OrderSn    string `gorm:"column:order_sn" json:"order_sn"`
	MemberName string `gorm:"column:member_name" json:"member_name"`
	CouponSn   string `gorm:"column:coupon_sn" json:"coupon_sn"`
	Title      string `gorm:"column:title" json:"title"`
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberCoupon) TableName() string {
	return "es_member_coupon"
}

// EsMemberNoticeLog 会员站内消息历史(es_member_notice_log)
type EsMemberNoticeLog struct {
	ID          int    `gorm:"primaryKey;column:id" json:"-"`           // 会员历史消息id
	MemberID    int    `gorm:"column:member_id" json:"member_id"`       // 会员id
	Content     string `gorm:"column:content" json:"content"`           // 站内信内容
	SendTime    int64  `gorm:"column:send_time" json:"send_time"`       // 发送时间
	IsDel       int    `gorm:"column:is_del" json:"is_del"`             // 是否删除，0删除，1没有删除
	IsRead      int    `gorm:"column:is_read" json:"is_read"`           // 是否已读，0未读，1已读
	ReceiveTime int64  `gorm:"column:receive_time" json:"receive_time"` // 接收时间
	Title       string `gorm:"column:title" json:"title"`               // 标题
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberNoticeLog) TableName() string {
	return "es_member_notice_log"
}

// EsMemberPointHistory 会员积分表(es_member_point_history)
type EsMemberPointHistory struct {
	ID              int    `gorm:"primaryKey;column:id" json:"-"`                     // 主键ID
	MemberID        int    `gorm:"column:member_id" json:"member_id"`                 // 会员ID
	GradePoint      int64  `gorm:"column:grade_point" json:"grade_point"`             // 等级积分
	Time            int64  `gorm:"column:time" json:"time"`                           // 操作时间
	Reason          string `gorm:"column:reason" json:"reason"`                       // 操作理由
	GradePointType  int    `gorm:"column:grade_point_type" json:"grade_point_type"`   // 等级积分类型
	Operator        string `gorm:"column:operator" json:"operator"`                   // 操作者
	ConsumPoint     int64  `gorm:"column:consum_point" json:"consum_point"`           // 消费积分
	ConsumPointType int    `gorm:"column:consum_point_type" json:"consum_point_type"` // 消费积分类型
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberPointHistory) TableName() string {
	return "es_member_point_history"
}

// EsMemberReceipt 会员发票信息缓存表(es_member_receipt)
type EsMemberReceipt struct {
	ReceiptID      int    `gorm:"primaryKey;column:receipt_id" json:"-"`         // 主键ID
	MemberID       int    `gorm:"column:member_id" json:"member_id"`             // 会员ID
	ReceiptType    string `gorm:"column:receipt_type" json:"receipt_type"`       // 发票类型 ELECTRO：电子普通发票，VATORDINARY：增值税普通发票
	ReceiptTitle   string `gorm:"column:receipt_title" json:"receipt_title"`     // 发票抬头
	ReceiptContent string `gorm:"column:receipt_content" json:"receipt_content"` // 发票内容
	TaxNo          string `gorm:"column:tax_no" json:"tax_no"`                   // 纳税人识别号
	MemberMobile   string `gorm:"column:member_mobile" json:"member_mobile"`     // 收票人手机号
	MemberEmail    string `gorm:"column:member_email" json:"member_email"`       // 收票人邮箱
	IsDefault      int16  `gorm:"column:is_default" json:"is_default"`           // 是否为默认选项 0：否，1：是
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberReceipt) TableName() string {
	return "es_member_receipt"
}

// EsMemberShopScore 店铺评分(es_member_shop_score)
type EsMemberShopScore struct {
	ScoreID          int    `gorm:"primaryKey;column:score_id" json:"-"`               // 主键
	MemberID         int    `gorm:"column:member_id" json:"member_id"`                 // 会员id
	OrderSn          string `gorm:"column:order_sn" json:"order_sn"`                   // 订单编号
	DeliveryScore    int16  `gorm:"column:delivery_score" json:"delivery_score"`       // 发货速度评分
	DescriptionScore int16  `gorm:"column:description_score" json:"description_score"` // 描述相符度评分
	ServiceScore     int16  `gorm:"column:service_score" json:"service_score"`         // 服务评分
	SellerID         int    `gorm:"column:seller_id" json:"seller_id"`                 // 卖家
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberShopScore) TableName() string {
	return "es_member_shop_score"
}

// EsMemberZpzz 会员增票资质表(es_member_zpzz)
type EsMemberZpzz struct {
	ID              int    `gorm:"primaryKey;column:id" json:"-"`                   // 主键ID
	MemberID        int    `gorm:"column:member_id" json:"member_id"`               // 会员ID
	Uname           string `gorm:"column:uname" json:"uname"`                       // 会员用户名
	Status          string `gorm:"column:status" json:"status"`                     // 状态 NEW_APPLY：新申请，AUDIT_PASS：审核通过，AUDIT_REFUSE：审核未通过
	CompanyName     string `gorm:"column:company_name" json:"company_name"`         // 单位名称
	TaxpayerCode    string `gorm:"column:taxpayer_code" json:"taxpayer_code"`       // 纳税人识别码
	RegisterAddress string `gorm:"column:register_address" json:"register_address"` // 公司注册地址
	RegisterTel     string `gorm:"column:register_tel" json:"register_tel"`         // 公司注册电话
	BankName        string `gorm:"column:bank_name" json:"bank_name"`               // 开户银行
	BankAccount     string `gorm:"column:bank_account" json:"bank_account"`         // 银行账户
	AuditRemark     string `gorm:"column:audit_remark" json:"audit_remark"`         // 平台审核备注
	ApplyTime       int64  `gorm:"column:apply_time" json:"apply_time"`             // 申请时间
}

// TableName get sql table name.获取数据库表名
func (m *EsMemberZpzz) TableName() string {
	return "es_member_zpzz"
}

// EsMenu 菜单表(es_menu)
type EsMenu struct {
	ID          int    `gorm:"primaryKey;column:id" json:"-"`           // 菜单id
	ParentID    int    `gorm:"column:parent_id" json:"parent_id"`       // 父id
	Title       string `gorm:"column:title" json:"title"`               // 菜单标题
	URL         string `gorm:"column:url" json:"url"`                   // 菜单url
	IDentifier  string `gorm:"column:identifier" json:"identifier"`     // 菜单唯一标识
	AuthRegular string `gorm:"column:auth_regular" json:"auth_regular"` // 权限表达式
	DeleteFlag  int    `gorm:"column:delete_flag" json:"delete_flag"`   // 删除标记
	Path        string `gorm:"column:path" json:"path"`                 // 菜单级别标识
	Grade       int    `gorm:"column:grade" json:"grade"`               // 菜单级别
}

// TableName get sql table name.获取数据库表名
func (m *EsMenu) TableName() string {
	return "es_menu"
}

// EsMessage 站内消息(es_message)
type EsMessage struct {
	ID        int    `gorm:"primaryKey;column:id" json:"-"`       // 站内消息主键
	Title     string `gorm:"column:title" json:"title"`           // 标题
	Content   string `gorm:"column:content" json:"content"`       // 消息内容
	MemberIDs string `gorm:"column:member_ids" json:"member_ids"` // 会员id
	AdminID   int    `gorm:"column:admin_id" json:"admin_id"`     // 管理员id
	AdminName string `gorm:"column:admin_name" json:"admin_name"` // 管理员名称
	SendTime  int64  `gorm:"column:send_time" json:"send_time"`   // 发送时间
	SendType  int    `gorm:"column:send_type" json:"send_type"`   // 发送类型
	Disabled  int    `gorm:"column:disabled" json:"disabled"`     // 是否删除 0：否，1：是
}

// TableName get sql table name.获取数据库表名
func (m *EsMessage) TableName() string {
	return "es_message"
}

// EsMessageTemplate 消息模版(es_message_template)
type EsMessageTemplate struct {
	ID           int    `gorm:"primaryKey;column:id" json:"-"`             // 模版ID
	TplCode      string `gorm:"column:tpl_code" json:"tpl_code"`           // 模版编号
	TplName      string `gorm:"column:tpl_name" json:"tpl_name"`           // 模板名称
	Type         string `gorm:"column:type" json:"type"`                   // 类型(会员 ,店铺 ,其他)
	NoticeState  string `gorm:"column:notice_state" json:"notice_state"`   // 站内信提醒是否开启
	SmsState     string `gorm:"column:sms_state" json:"sms_state"`         // 短信提醒是否开启
	EmailState   string `gorm:"column:email_state" json:"email_state"`     // 邮件提醒是否开启
	Content      string `gorm:"column:content" json:"content"`             // 站内信内容
	SmsContent   string `gorm:"column:sms_content" json:"sms_content"`     // 短信内容
	EmailContent string `gorm:"column:email_content" json:"email_content"` // 邮件内容
	EmailTitle   string `gorm:"column:email_title" json:"email_title"`     // 邮件标题
}

// TableName get sql table name.获取数据库表名
func (m *EsMessageTemplate) TableName() string {
	return "es_message_template"
}

// EsMinus 单品立减(es_minus)
type EsMinus struct {
	MinusID              int     `gorm:"primaryKey;column:minus_id" json:"-"`                         // 单品立减活动id
	SingleReductionValue float64 `gorm:"column:single_reduction_value" json:"single_reduction_value"` // 单品立减金额
	StartTime            int64   `gorm:"column:start_time" json:"start_time"`                         // 起始时间
	StartTimeStr         string  `gorm:"column:start_time_str" json:"start_time_str"`                 // 起始时间字符串
	EndTime              int64   `gorm:"column:end_time" json:"end_time"`                             // 结束时间
	EndTimeStr           string  `gorm:"column:end_time_str" json:"end_time_str"`                     // 结束时间字符串
	Title                string  `gorm:"column:title" json:"title"`                                   // 单品立减活动标题
	RangeType            int     `gorm:"column:range_type" json:"range_type"`                         // 是否选择全部商品
	Disabled             int     `gorm:"column:disabled" json:"disabled"`                             // 是否停用
	Description          string  `gorm:"column:description" json:"description"`                       // 描述
	SellerID             int     `gorm:"column:seller_id" json:"seller_id"`                           // 商家id
}

// TableName get sql table name.获取数据库表名
func (m *EsMinus) TableName() string {
	return "es_minus"
}

// EsNavigation 店铺导航管理(es_navigation)
type EsNavigation struct {
	ID       int    `gorm:"primaryKey;column:id" json:"-"`   // 导航id
	Name     string `gorm:"column:name" json:"name"`         // 名称
	Disable  int    `gorm:"column:disable" json:"disable"`   // 是否显示
	Sort     int    `gorm:"column:sort" json:"sort"`         // 排序
	NavURL   string `gorm:"column:nav_url" json:"nav_url"`   // URL
	Target   int    `gorm:"column:target" json:"target"`     // 新窗口打开
	ShopID   int    `gorm:"column:shop_id" json:"shop_id"`   // 店铺id
	Contents string `gorm:"column:contents" json:"contents"` // 内容
}

// TableName get sql table name.获取数据库表名
func (m *EsNavigation) TableName() string {
	return "es_navigation"
}

// EsOrderComplainCommunication [...]
type EsOrderComplainCommunication struct {
	CommunicationID int    `gorm:"primaryKey;column:communication_id" json:"-"` // 主键
	ComplainID      int    `gorm:"column:complain_id" json:"complain_id"`       // 投诉id
	Content         string `gorm:"column:content" json:"content"`               // 对话内容
	CreateTime      int64  `gorm:"column:create_time" json:"create_time"`       // 对话时间
	Owner           string `gorm:"column:owner" json:"owner"`                   // 所属，买家/卖家
	OwnerName       string `gorm:"column:owner_name" json:"owner_name"`         // 对话所属名称
	OwnerID         int    `gorm:"column:owner_id" json:"owner_id"`             // 对话所属id,卖家id/买家id
}

// TableName get sql table name.获取数据库表名
func (m *EsOrderComplainCommunication) TableName() string {
	return "es_order_complain_communication"
}

// EsOrderLog 订单日志表(es_order_log)
type EsOrderLog struct {
	LogID   int    `gorm:"primaryKey;column:log_id" json:"-"` // 主键ID
	OrderSn string `gorm:"column:order_sn" json:"order_sn"`   // 订单编号
	OpID    int    `gorm:"column:op_id" json:"op_id"`         // 操作者id
	OpName  string `gorm:"column:op_name" json:"op_name"`     // 操作者名称
	Message string `gorm:"column:message" json:"message"`     // 日志信息
	OpTime  int64  `gorm:"column:op_time" json:"op_time"`     // 操作时间
}

// TableName get sql table name.获取数据库表名
func (m *EsOrderLog) TableName() string {
	return "es_order_log"
}

// EsOrderMeta 订单扩展信息表(es_order_meta)
type EsOrderMeta struct {
	MetaID    int    `gorm:"primaryKey;column:meta_id" json:"-"`  // 主键ID
	OrderSn   string `gorm:"column:order_sn" json:"order_sn"`     // 订单编号
	MetaKey   string `gorm:"column:meta_key" json:"meta_key"`     // 扩展-键
	MetaValue string `gorm:"column:meta_value" json:"meta_value"` // 扩展-值
	Status    string `gorm:"column:status" json:"status"`         // 售后状态
}

// TableName get sql table name.获取数据库表名
func (m *EsOrderMeta) TableName() string {
	return "es_order_meta"
}

// EsOrderOutStatus 订单出库状态(es_order_out_status)
type EsOrderOutStatus struct {
	OutID     int    `gorm:"primaryKey;column:out_id" json:"-"`   // 主键
	OrderSn   string `gorm:"column:order_sn" json:"order_sn"`     // 订单编号
	OutType   string `gorm:"column:out_type" json:"out_type"`     // 出库类型
	OutStatus string `gorm:"column:out_status" json:"out_status"` // 出库状态
}

// TableName get sql table name.获取数据库表名
func (m *EsOrderOutStatus) TableName() string {
	return "es_order_out_status"
}

// EsPage 楼层(es_page)
type EsPage struct {
	PageID     int    `gorm:"primaryKey;column:page_id" json:"-"`    // 主键id
	PageName   string `gorm:"column:page_name" json:"page_name"`     // 楼层名称
	PageData   string `gorm:"column:page_data" json:"page_data"`     // 楼层数据
	PageType   string `gorm:"column:page_type" json:"page_type"`     // 页面类型
	ClientType string `gorm:"column:client_type" json:"client_type"` // 客户端类型
}

// TableName get sql table name.获取数据库表名
func (m *EsPage) TableName() string {
	return "es_page"
}

// EsParameterGroup 参数组(es_parameter_group)
type EsParameterGroup struct {
	GroupID    int    `gorm:"primaryKey;column:group_id" json:"-"`   // 主键
	GroupName  string `gorm:"column:group_name" json:"group_name"`   // 参数组名称
	CategoryID int    `gorm:"column:category_id" json:"category_id"` // 关联分类id
	Sort       int    `gorm:"column:sort" json:"sort"`               // 排序
}

// TableName get sql table name.获取数据库表名
func (m *EsParameterGroup) TableName() string {
	return "es_parameter_group"
}

// EsParameters 参数(es_parameters)
type EsParameters struct {
	ParamID    int    `gorm:"primaryKey;column:param_id" json:"-"`   // 主键
	ParamName  string `gorm:"column:param_name" json:"param_name"`   // 参数名称
	ParamType  int    `gorm:"column:param_type" json:"param_type"`   // 参数类型1 输入项   2 选择项
	Options    string `gorm:"column:options" json:"options"`         // 选择值，当参数类型是选择项2时，必填，逗号分隔
	IsIndex    int    `gorm:"column:is_index" json:"is_index"`       // 是否可索引，0 不显示 1 显示
	Required   int    `gorm:"column:required" json:"required"`       // 是否必填是  1    否   0
	GroupID    int    `gorm:"column:group_id" json:"group_id"`       // 参数分组id
	CategoryID int    `gorm:"column:category_id" json:"category_id"` // 分类id
	Sort       int    `gorm:"column:sort" json:"sort"`               // 排序
}

// TableName get sql table name.获取数据库表名
func (m *EsParameters) TableName() string {
	return "es_parameters"
}

// EsPayLog 收款单(es_pay_log)
type EsPayLog struct {
	PayLogID      int     `gorm:"primaryKey;column:pay_log_id" json:"-"`         // 收款单ID
	OrderSn       string  `gorm:"column:order_sn" json:"order_sn"`               // 订单编号
	PayWay        string  `gorm:"column:pay_way" json:"pay_way"`                 // 支付方式
	PayType       string  `gorm:"column:pay_type" json:"pay_type"`               // 支付类型
	PayTime       int64   `gorm:"column:pay_time" json:"pay_time"`               // 付款时间
	PayMoney      float64 `gorm:"column:pay_money" json:"pay_money"`             // 付款金额
	PayMemberName string  `gorm:"column:pay_member_name" json:"pay_member_name"` // 付款会员名
	PayStatus     string  `gorm:"column:pay_status" json:"pay_status"`           // 付款状态
	PayLogSn      string  `gorm:"column:pay_log_sn" json:"pay_log_sn"`           // 流水号
	PayOrderNo    string  `gorm:"column:pay_order_no" json:"pay_order_no"`       // 支付方式返回流水号
}

// TableName get sql table name.获取数据库表名
func (m *EsPayLog) TableName() string {
	return "es_pay_log"
}

// EsPaymentBill 支付帐单(es_payment_bill)
type EsPaymentBill struct {
	BillID          int     `gorm:"primaryKey;column:bill_id" json:"-"`                // 主键id
	Sn              string  `gorm:"column:sn" json:"sn"`                               // 单号
	OutTradeNo      string  `gorm:"column:out_trade_no" json:"out_trade_no"`           // 提交给第三方平台单号
	ReturnTradeNo   string  `gorm:"column:return_trade_no" json:"return_trade_no"`     // 第三方平台返回交易号
	IsPay           int     `gorm:"column:is_pay" json:"is_pay"`                       // 是否已支付
	TradeType       string  `gorm:"column:trade_type" json:"trade_type"`               // 交易类型
	PaymentName     string  `gorm:"column:payment_name" json:"payment_name"`           // 支付方式名称
	PayConfig       string  `gorm:"column:pay_config" json:"pay_config"`               // 支付参数
	TradePrice      float64 `gorm:"column:trade_price" json:"trade_price"`             // 交易金额
	PaymentPluginID string  `gorm:"column:payment_plugin_id" json:"payment_plugin_id"` // 支付插件
}

// TableName get sql table name.获取数据库表名
func (m *EsPaymentBill) TableName() string {
	return "es_payment_bill"
}

// EsPaymentMethod 支付方式(es_payment_method)
type EsPaymentMethod struct {
	MethodID        int    `gorm:"primaryKey;column:method_id" json:"-"`              // 支付方式id
	MethodName      string `gorm:"column:method_name" json:"method_name"`             // 支付方式名称
	PluginID        string `gorm:"column:plugin_id" json:"plugin_id"`                 // 支付插件名称
	PcConfig        string `gorm:"column:pc_config" json:"pc_config"`                 // pc是否可用
	WapConfig       string `gorm:"column:wap_config" json:"wap_config"`               // wap是否可用
	AppNativeConfig string `gorm:"column:app_native_config" json:"app_native_config"` // app 原生是否可用
	Image           string `gorm:"column:image" json:"image"`                         // 支付方式图片
	IsRetrace       int    `gorm:"column:is_retrace" json:"is_retrace"`               // 是否支持原路退回
	AppReactConfig  string `gorm:"column:app_react_config" json:"app_react_config"`   // app RN是否可用
	MiniConfig      string `gorm:"column:mini_config" json:"mini_config"`             // 小程序配置
}

// TableName get sql table name.获取数据库表名
func (m *EsPaymentMethod) TableName() string {
	return "es_payment_method"
}

// EsPintuan [...]
type EsPintuan struct {
	PromotionID    int    `gorm:"primaryKey;column:promotion_id" json:"-"`       // id
	PromotionName  string `gorm:"column:promotion_name" json:"promotion_name"`   // 活动名称
	PromotionTitle string `gorm:"column:promotion_title" json:"promotion_title"` // 活动标题
	StartTime      int64  `gorm:"column:start_time" json:"start_time"`           // 活动开始时间
	EndTime        int64  `gorm:"column:end_time" json:"end_time"`               // 活动结束时间
	RequiredNum    int    `gorm:"column:required_num" json:"required_num"`       // 成团人数
	LimitNum       int    `gorm:"column:limit_num" json:"limit_num"`             // 限购数量
	EnableMocker   int16  `gorm:"column:enable_mocker" json:"enable_mocker"`     // 虚拟成团
	CreateTime     int64  `gorm:"column:create_time" json:"create_time"`         // 创建时间
	Status         string `gorm:"column:status" json:"status"`                   // 活动状态
	OptionStatus   string `gorm:"column:option_status" json:"option_status"`     // api请求操作状态
	SellerName     string `gorm:"column:seller_name" json:"seller_name"`         // 商家名称
	SellerID       int64  `gorm:"column:seller_id" json:"seller_id"`             // 商家ID
}

// TableName get sql table name.获取数据库表名
func (m *EsPintuan) TableName() string {
	return "es_pintuan"
}

// EsPintuanChildOrder [...]
type EsPintuanChildOrder struct {
	ChildOrderID int     `gorm:"primaryKey;column:child_order_id" json:"-"` // 子订单id
	OrderSn      string  `gorm:"column:order_sn" json:"order_sn"`           // 订单编号
	MemberID     int     `gorm:"column:member_id" json:"member_id"`         // 会员id
	SkuID        int     `gorm:"column:sku_id" json:"sku_id"`               // 会员id
	PintuanID    int     `gorm:"column:pintuan_id" json:"pintuan_id"`       // 拼团活动id
	OrderStatus  string  `gorm:"column:order_status" json:"order_status"`   // 订单状态
	OrderID      int     `gorm:"column:order_id" json:"order_id"`           // 主订单id
	MemberName   string  `gorm:"column:member_name" json:"member_name"`     // 买家名称
	OriginPrice  float64 `gorm:"column:origin_price" json:"origin_price"`   // 原价
	SalesPrice   float64 `gorm:"column:sales_price" json:"sales_price"`     // 拼团价
}

// TableName get sql table name.获取数据库表名
func (m *EsPintuanChildOrder) TableName() string {
	return "es_pintuan_child_order"
}

// EsPintuanGoods [...]
type EsPintuanGoods struct {
	ID             int     `gorm:"primaryKey;column:id" json:"-"`                 // id
	SkuID          int     `gorm:"column:sku_id" json:"sku_id"`                   // sku_id
	GoodsName      string  `gorm:"column:goods_name" json:"goods_name"`           // 商品名称
	OriginPrice    float64 `gorm:"column:origin_price" json:"origin_price"`       // 原价
	SalesPrice     float64 `gorm:"column:sales_price" json:"sales_price"`         // 活动价
	Sn             string  `gorm:"column:sn" json:"sn"`                           // sn
	SoldQuantity   int     `gorm:"column:sold_quantity" json:"sold_quantity"`     // 已售数量
	LockedQuantity int     `gorm:"column:locked_quantity" json:"locked_quantity"` // 待发货数量
	PintuanID      int     `gorm:"column:pintuan_id" json:"pintuan_id"`           // 拼团活动id
	GoodsID        int     `gorm:"column:goods_id" json:"goods_id"`               // goods_id
	Specs          string  `gorm:"column:specs" json:"specs"`                     // 规格
	SellerID       int     `gorm:"column:seller_id" json:"seller_id"`             // 卖家id
	SellerName     string  `gorm:"column:seller_name" json:"seller_name"`         // 卖家名称
	Thumbnail      string  `gorm:"column:thumbnail" json:"thumbnail"`             // 商品缩略图
}

// TableName get sql table name.获取数据库表名
func (m *EsPintuanGoods) TableName() string {
	return "es_pintuan_goods"
}

// EsPintuanOrder [...]
type EsPintuanOrder struct {
	OrderID        int    `gorm:"primaryKey;column:order_id" json:"-"`           // 订单id
	PintuanID      int    `gorm:"column:pintuan_id" json:"pintuan_id"`           // 拼团id
	EndTime        int64  `gorm:"column:end_time" json:"end_time"`               // 结束时间
	SkuID          int    `gorm:"column:sku_id" json:"sku_id"`                   // sku_id
	RequiredNum    int    `gorm:"column:required_num" json:"required_num"`       // 成团人数
	OfferedNum     int    `gorm:"column:offered_num" json:"offered_num"`         // 已参团人数
	OfferedPersons string `gorm:"column:offered_persons" json:"offered_persons"` // 参团人
	OrderStatus    string `gorm:"column:order_status" json:"order_status"`       // 订单状态
	GoodsID        int    `gorm:"column:goods_id" json:"goods_id"`               // 商品id
	Thumbnail      string `gorm:"column:thumbnail" json:"thumbnail"`             // 商品缩略图
	GoodsName      string `gorm:"column:goods_name" json:"goods_name"`           // 商品名称
}

// TableName get sql table name.获取数据库表名
func (m *EsPintuanOrder) TableName() string {
	return "es_pintuan_order"
}

// EsPromotionGoods 有效活动商品对照表(es_promotion_goods)
type EsPromotionGoods struct {
	PgID          int     `gorm:"primaryKey;column:pg_id" json:"-"`            // 商品对照ID
	GoodsID       int     `gorm:"column:goods_id" json:"goods_id"`             // 商品id
	ProductID     int     `gorm:"column:product_id" json:"product_id"`         // 货品id
	StartTime     int64   `gorm:"column:start_time" json:"start_time"`         // 活动开始时间
	EndTime       int64   `gorm:"column:end_time" json:"end_time"`             // 活动结束时间
	ActivityID    int     `gorm:"column:activity_id" json:"activity_id"`       // 活动id
	PromotionType string  `gorm:"column:promotion_type" json:"promotion_type"` // 促销工具类型
	Title         string  `gorm:"column:title" json:"title"`                   // 活动标题
	Num           int     `gorm:"column:num" json:"num"`                       // 参与活动的商品数量
	Price         float64 `gorm:"column:price" json:"price"`                   // 活动时商品的价格
	SellerID      int     `gorm:"column:seller_id" json:"seller_id"`           // 商家ID
}

// TableName get sql table name.获取数据库表名
func (m *EsPromotionGoods) TableName() string {
	return "es_promotion_goods"
}

// EsQuantityLog 库存日志表(es_quantity_log)
type EsQuantityLog struct {
	LogID          int    `gorm:"primaryKey;column:log_id" json:"-"`             // 日志id
	OrderSn        string `gorm:"column:order_sn" json:"order_sn"`               // 订单编号
	GoodsID        int    `gorm:"column:goods_id" json:"goods_id"`               // 商品id
	SkuID          int    `gorm:"column:sku_id" json:"sku_id"`                   // sku id
	Quantity       int    `gorm:"column:quantity" json:"quantity"`               // 库存数
	EnableQuantity int    `gorm:"column:enable_quantity" json:"enable_quantity"` // 可用库存
	OpTime         int    `gorm:"column:op_time" json:"op_time"`                 // 操作时间
	LogType        string `gorm:"column:log_type" json:"log_type"`               // 日志类型
	Reason         string `gorm:"column:reason" json:"reason"`                   // 原因
}

// TableName get sql table name.获取数据库表名
func (m *EsQuantityLog) TableName() string {
	return "es_quantity_log"
}

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

// EsReceiptContent 发票内容(es_receipt_content)
type EsReceiptContent struct {
	ID      int    `gorm:"primaryKey;column:id" json:"-"` // 发票内容id
	Content string `gorm:"column:content" json:"content"` // 发票内容
}

// TableName get sql table name.获取数据库表名
func (m *EsReceiptContent) TableName() string {
	return "es_receipt_content"
}

// EsReceiptFile 会员电子发票附件表(es_receipt_file)
type EsReceiptFile struct {
	ID        int    `gorm:"primaryKey;column:id" json:"-"`
	HistoryID int    `gorm:"column:history_id" json:"history_id"`
	ElecFile  string `gorm:"column:elec_file" json:"elec_file"`
}

// TableName get sql table name.获取数据库表名
func (m *EsReceiptFile) TableName() string {
	return "es_receipt_file"
}

// EsRegions 地区(es_regions)
type EsRegions struct {
	ID          int    `gorm:"primaryKey;column:id" json:"-"`           // 地区id
	ParentID    int    `gorm:"column:parent_id" json:"parent_id"`       // 父地区id
	RegionPath  string `gorm:"column:region_path" json:"region_path"`   // 路径
	RegionGrade int    `gorm:"column:region_grade" json:"region_grade"` // 级别
	LocalName   string `gorm:"column:local_name" json:"local_name"`     // 名称
	Zipcode     string `gorm:"column:zipcode" json:"zipcode"`           // 邮编
	Cod         string `gorm:"column:cod" json:"cod"`                   // 是否支持货到付款
}

// TableName get sql table name.获取数据库表名
func (m *EsRegions) TableName() string {
	return "es_regions"
}

// EsRole 角色表(es_role)
type EsRole struct {
	RoleID       int    `gorm:"primaryKey;column:role_id" json:"-"`        // 主键ID
	RoleName     string `gorm:"column:role_name" json:"role_name"`         // 角色名称
	AuthIDs      string `gorm:"column:auth_ids" json:"auth_ids"`           // 角色介绍
	RoleDescribe string `gorm:"column:role_describe" json:"role_describe"` // 角色描述
}

// TableName get sql table name.获取数据库表名
func (m *EsRole) TableName() string {
	return "es_role"
}

// EsSeckill 限时抢购入库(es_seckill)
type EsSeckill struct {
	SeckillID     int    `gorm:"primaryKey;column:seckill_id" json:"-"`       // 主键id
	SeckillName   string `gorm:"column:seckill_name" json:"seckill_name"`     // 活动名称
	StartDay      int64  `gorm:"column:start_day" json:"start_day"`           // 活动日期
	ApplyEndTime  int64  `gorm:"column:apply_end_time" json:"apply_end_time"` // 报名截至时间
	SeckillRule   string `gorm:"column:seckill_rule" json:"seckill_rule"`     // 申请规则
	SellerIDs     string `gorm:"column:seller_ids" json:"seller_ids"`         // 商家id
	SeckillStatus string `gorm:"column:seckill_status" json:"seckill_status"` // 状态
	DeleteStatus  string `gorm:"column:delete_status" json:"delete_status"`   // 是否删除 DELETED：已删除，NORMAL：正常
}

// TableName get sql table name.获取数据库表名
func (m *EsSeckill) TableName() string {
	return "es_seckill"
}

// EsSeckillApply 限时抢购申请(es_seckill_apply)
type EsSeckillApply struct {
	ApplyID       int     `gorm:"primaryKey;column:apply_id" json:"-"`         // 主键ID
	SeckillID     int     `gorm:"column:seckill_id" json:"seckill_id"`         // 活动id
	TimeLine      int     `gorm:"column:time_line" json:"time_line"`           // 时刻
	StartDay      int64   `gorm:"column:start_day" json:"start_day"`           // 活动开始日期
	GoodsID       int     `gorm:"column:goods_id" json:"goods_id"`             // 商品ID
	GoodsName     string  `gorm:"column:goods_name" json:"goods_name"`         // 商品名称
	SkuID         int     `gorm:"column:sku_id" json:"sku_id"`                 // 商品规格ID
	SellerID      int     `gorm:"column:seller_id" json:"seller_id"`           // 商家ID
	ShopName      string  `gorm:"column:shop_name" json:"shop_name"`           // 商家名称
	Price         float64 `gorm:"column:price" json:"price"`                   // 价格
	SoldQuantity  int     `gorm:"column:sold_quantity" json:"sold_quantity"`   // 售空数量
	Status        string  `gorm:"column:status" json:"status"`                 // 申请状态
	FailReason    string  `gorm:"column:fail_reason" json:"fail_reason"`       // 驳回原因
	SalesNum      int     `gorm:"column:sales_num" json:"sales_num"`           // 已售数量
	OriginalPrice float64 `gorm:"column:original_price" json:"original_price"` // 商品原始价格
}

// TableName get sql table name.获取数据库表名
func (m *EsSeckillApply) TableName() string {
	return "es_seckill_apply"
}

// EsSeckillRange 限时抢购时刻(es_seckill_range)
type EsSeckillRange struct {
	RangeID   int `gorm:"primaryKey;column:range_id" json:"-"` // 主键
	SeckillID int `gorm:"column:seckill_id" json:"seckill_id"` // 限时抢购活动id
	RangeTime int `gorm:"column:range_time" json:"range_time"` // 整点时刻
}

// TableName get sql table name.获取数据库表名
func (m *EsSeckillRange) TableName() string {
	return "es_seckill_range"
}

// EsSellerBill 店铺结算(es_seller_bill)
type EsSellerBill struct {
	ID                int     `gorm:"primaryKey;column:id" json:"-"`                       // id
	SellerID          int     `gorm:"column:seller_id" json:"seller_id"`                   // 商家id
	CreateTime        int64   `gorm:"column:create_time" json:"create_time"`               // 创建时间
	OrderSn           string  `gorm:"column:order_sn" json:"order_sn"`                     // 订单编号
	Expenditure       float64 `gorm:"column:expenditure" json:"expenditure"`               // 商品反现支出
	ReturnExpenditure float64 `gorm:"column:return_expenditure" json:"return_expenditure"` // 返还商品反现支出
}

// TableName get sql table name.获取数据库表名
func (m *EsSellerBill) TableName() string {
	return "es_seller_bill"
}

// EsSensitiveWords 敏感词(es_sensitive_words)
type EsSensitiveWords struct {
	ID         int    `gorm:"primaryKey;column:id" json:"-"`         // 主键
	WordName   string `gorm:"column:word_name" json:"word_name"`     // 敏感词名称
	CreateTime int64  `gorm:"column:create_time" json:"create_time"` // 创建时间
	IsDelete   int16  `gorm:"column:is_delete" json:"is_delete"`     // 删除状态  1正常 0 删除
}

// TableName get sql table name.获取数据库表名
func (m *EsSensitiveWords) TableName() string {
	return "es_sensitive_words"
}

// EsSettings 系统设置(es_settings)
type EsSettings struct {
	ID       int    `gorm:"primaryKey;column:id" json:"-"`     // 系统设置id
	CfgValue string `gorm:"column:cfg_value" json:"cfg_value"` // 系统配置信息
	CfgGroup string `gorm:"column:cfg_group" json:"cfg_group"` // 业务设置标识
}

// TableName get sql table name.获取数据库表名
func (m *EsSettings) TableName() string {
	return "es_settings"
}

// EsShipTemplate 运费模版(es_ship_template)
type EsShipTemplate struct {
	ID       int    `gorm:"primaryKey;column:id" json:"-"`     // 主键
	SellerID int    `gorm:"column:seller_id" json:"seller_id"` // 卖家id
	Name     string `gorm:"column:name" json:"name"`           // 模板名称
	Type     int16  `gorm:"column:type" json:"type"`           // 1 重量算运费 2 计件算运费
}

// TableName get sql table name.获取数据库表名
func (m *EsShipTemplate) TableName() string {
	return "es_ship_template"
}

// EsShipTemplateChild 运费模板子模板(es_ship_template_child)
type EsShipTemplateChild struct {
	ID               int     `gorm:"primaryKey;column:id" json:"-"`                     // 主键
	TemplateID       int     `gorm:"column:template_id" json:"template_id"`             // 外键  模板id
	FirstCompany     int     `gorm:"column:first_company" json:"first_company"`         // 首个（件）
	FirstPrice       float64 `gorm:"column:first_price" json:"first_price"`             // 首个（件）价格
	ContinuedCompany int     `gorm:"column:continued_company" json:"continued_company"` // 续个（件）
	ContinuedPrice   float64 `gorm:"column:continued_price" json:"continued_price"`     // 续个（件）价格
	Area             string  `gorm:"column:area" json:"area"`                           // 地区json
	AreaID           string  `gorm:"column:area_id" json:"area_id"`                     // 地区id
}

// TableName get sql table name.获取数据库表名
func (m *EsShipTemplateChild) TableName() string {
	return "es_ship_template_child"
}

// EsShop 店铺(es_shop)
type EsShop struct {
	ShopID         int    `gorm:"primaryKey;column:shop_id" json:"-"`            // 店铺Id
	MemberID       int    `gorm:"column:member_id" json:"member_id"`             // 会员Id
	MemberName     string `gorm:"column:member_name" json:"member_name"`         // 会员名称
	ShopName       string `gorm:"column:shop_name" json:"shop_name"`             // 店铺名称
	ShopDisable    string `gorm:"column:shop_disable" json:"shop_disable"`       // 店铺状态
	ShopCreatetime int64  `gorm:"column:shop_createtime" json:"shop_createtime"` // 店铺创建时间
	ShopEndtime    int64  `gorm:"column:shop_endtime" json:"shop_endtime"`       // 店铺关闭时间
}

// TableName get sql table name.获取数据库表名
func (m *EsShop) TableName() string {
	return "es_shop"
}

// EsShopCat 店铺分组(es_shop_cat)
type EsShopCat struct {
	ShopCatID   int    `gorm:"primaryKey;column:shop_cat_id" json:"-"`    // 店铺分组id
	ShopCatPid  int    `gorm:"column:shop_cat_pid" json:"shop_cat_pid"`   // 店铺分组父ID
	ShopID      int    `gorm:"column:shop_id" json:"shop_id"`             // 店铺id
	ShopCatName string `gorm:"column:shop_cat_name" json:"shop_cat_name"` // 店铺分组名称
	Disable     int    `gorm:"column:disable" json:"disable"`             // 店铺分组显示状态:1显示 0不显示
	Sort        int    `gorm:"column:sort" json:"sort"`                   // 排序
	CatPath     string `gorm:"column:cat_path" json:"cat_path"`           // 分组路径
}

// TableName get sql table name.获取数据库表名
func (m *EsShopCat) TableName() string {
	return "es_shop_cat"
}

// EsShopLogiRel 店铺物流公司对照表(es_shop_logi_rel)
type EsShopLogiRel struct {
	LogiID int `gorm:"column:logi_id" json:"logi_id"` // 物流公司ID
	ShopID int `gorm:"column:shop_id" json:"shop_id"` // 店铺ID
}

// TableName get sql table name.获取数据库表名
func (m *EsShopLogiRel) TableName() string {
	return "es_shop_logi_rel"
}

// EsShopLogisticsSetting [...]
type EsShopLogisticsSetting struct {
	ID          int    `gorm:"primaryKey;column:id" json:"-"`
	ShopID      int    `gorm:"column:shop_id" json:"shop_id"`
	LogisticsID int    `gorm:"column:logistics_id" json:"logistics_id"`
	Config      string `gorm:"column:config" json:"config"`
}

// TableName get sql table name.获取数据库表名
func (m *EsShopLogisticsSetting) TableName() string {
	return "es_shop_logistics_setting"
}

// EsShopMenu 菜单管理店铺(es_shop_menu)
type EsShopMenu struct {
	ID          int    `gorm:"primaryKey;column:id" json:"-"`           // 店铺菜单id
	ParentID    int    `gorm:"column:parent_id" json:"parent_id"`       // 菜单父id
	Title       string `gorm:"column:title" json:"title"`               // 菜单标题
	IDentifier  string `gorm:"column:identifier" json:"identifier"`     // 菜单唯一标识
	AuthRegular string `gorm:"column:auth_regular" json:"auth_regular"` // 权限表达式
	DeleteFlag  int    `gorm:"column:delete_flag" json:"delete_flag"`   // 删除标记
	Path        string `gorm:"column:path" json:"path"`                 // 菜单级别标识
	Grade       int    `gorm:"column:grade" json:"grade"`               // 菜单级别
}

// TableName get sql table name.获取数据库表名
func (m *EsShopMenu) TableName() string {
	return "es_shop_menu"
}

// EsShopNoticeLog 店铺站内消息(es_shop_notice_log)
type EsShopNoticeLog struct {
	ID            int    `gorm:"primaryKey;column:id" json:"-"`               // id
	ShopID        int    `gorm:"column:shop_id" json:"shop_id"`               // 店铺ID
	NoticeContent string `gorm:"column:notice_content" json:"notice_content"` // 站内信内容
	SendTime      int64  `gorm:"column:send_time" json:"send_time"`           // 发送时间
	IsDelete      int    `gorm:"column:is_delete" json:"is_delete"`           // 是否删除 ：1 删除   0  未删除
	IsRead        int    `gorm:"column:is_read" json:"is_read"`               // 是否已读 ：1已读   0 未读
	Type          string `gorm:"column:type" json:"type"`                     // 消息类型
}

// TableName get sql table name.获取数据库表名
func (m *EsShopNoticeLog) TableName() string {
	return "es_shop_notice_log"
}

// EsShopRole 店铺角色(es_shop_role)
type EsShopRole struct {
	RoleID       int    `gorm:"primaryKey;column:role_id" json:"-"`        // 角色主键
	RoleName     string `gorm:"column:role_name" json:"role_name"`         // 角色名称
	AuthIDs      string `gorm:"column:auth_ids" json:"auth_ids"`           // 角色
	RoleDescribe string `gorm:"column:role_describe" json:"role_describe"` // 角色描述
	ShopID       int    `gorm:"column:shop_id" json:"shop_id"`             // 店铺id
}

// TableName get sql table name.获取数据库表名
func (m *EsShopRole) TableName() string {
	return "es_shop_role"
}

// EsShopSilde 店铺幻灯片(es_shop_silde)
type EsShopSilde struct {
	SildeID  int    `gorm:"primaryKey;column:silde_id" json:"-"` // 幻灯片Id
	ShopID   int    `gorm:"column:shop_id" json:"shop_id"`       // 店铺Id
	SildeURL string `gorm:"column:silde_url" json:"silde_url"`   // 幻灯片URL
	Img      string `gorm:"column:img" json:"img"`               // 图片
}

// TableName get sql table name.获取数据库表名
func (m *EsShopSilde) TableName() string {
	return "es_shop_silde"
}

// EsShopThemes 店铺模版(es_shop_themes)
type EsShopThemes struct {
	ID        int    `gorm:"primaryKey;column:id" json:"-"`       // 模版id
	Name      string `gorm:"column:name" json:"name"`             // 模版名称
	ImagePath string `gorm:"column:image_path" json:"image_path"` // 模版图片路径
	IsDefault int    `gorm:"column:is_default" json:"is_default"` // 是否为默认
	Type      string `gorm:"column:type" json:"type"`             // 模版类型
	Mark      string `gorm:"column:mark" json:"mark"`             // 模版标识
}

// TableName get sql table name.获取数据库表名
func (m *EsShopThemes) TableName() string {
	return "es_shop_themes"
}

// EsShortURL 短链接(es_short_url)
type EsShortURL struct {
	ID  int64  `gorm:"primaryKey;column:id" json:"-"` // id
	URL string `gorm:"column:url" json:"url"`         // 跳转地址
	Su  string `gorm:"column:su" json:"su"`           // 短链接key
}

// TableName get sql table name.获取数据库表名
func (m *EsShortURL) TableName() string {
	return "es_short_url"
}

// EsSiteNavigation 导航栏(es_site_navigation)
type EsSiteNavigation struct {
	NavigationID   int    `gorm:"primaryKey;column:navigation_id" json:"-"`      // 主键
	NavigationName string `gorm:"column:navigation_name" json:"navigation_name"` // 导航名称
	URL            string `gorm:"column:url" json:"url"`                         // 导航地址
	ClientType     string `gorm:"column:client_type" json:"client_type"`         // 客户端类型
	Image          string `gorm:"column:image" json:"image"`                     // 图片
	Sort           int    `gorm:"column:sort" json:"sort"`                       // 排序
}

// TableName get sql table name.获取数据库表名
func (m *EsSiteNavigation) TableName() string {
	return "es_site_navigation"
}

// EsSmsPlatform 短信网关(es_sms_platform)
type EsSmsPlatform struct {
	ID     int    `gorm:"primaryKey;column:id" json:"-"` // 主键ID
	Name   string `gorm:"column:name" json:"name"`       // 平台名称
	Open   int    `gorm:"column:open" json:"open"`       // 是否开启
	Config string `gorm:"column:config" json:"config"`   // 配置
	Bean   string `gorm:"column:bean" json:"bean"`       // beanid
}

// TableName get sql table name.获取数据库表名
func (m *EsSmsPlatform) TableName() string {
	return "es_sms_platform"
}

// EsSMTP 邮件(es_smtp)
type EsSMTP struct {
	ID           int    `gorm:"primaryKey;column:id" json:"-"`               // 主键ID
	Host         string `gorm:"column:host" json:"host"`                     // 主机
	Username     string `gorm:"column:username" json:"username"`             // 用户名
	Password     string `gorm:"column:password" json:"password"`             // 密码
	LastSendTime int64  `gorm:"column:last_send_time" json:"last_send_time"` // 最后发信时间
	SendCount    int    `gorm:"column:send_count" json:"send_count"`         // 已发数
	MaxCount     int    `gorm:"column:max_count" json:"max_count"`           // 最大发信数
	MailFrom     string `gorm:"column:mail_from" json:"mail_from"`           // 发信邮箱
	Port         int    `gorm:"column:port" json:"port"`                     // 端口
	OpenSsl      int    `gorm:"column:open_ssl" json:"open_ssl"`             // ssl是否开启
}

// TableName get sql table name.获取数据库表名
func (m *EsSMTP) TableName() string {
	return "es_smtp"
}

// EsSpecValues 规格值(es_spec_values)
type EsSpecValues struct {
	SpecValueID int    `gorm:"primaryKey;column:spec_value_id" json:"-"` // 主键
	SpecID      int    `gorm:"column:spec_id" json:"spec_id"`            // 规格项id
	SpecValue   string `gorm:"column:spec_value" json:"spec_value"`      // 规格值名字
	SellerID    int    `gorm:"column:seller_id" json:"seller_id"`        // 所属卖家
	SpecName    string `gorm:"column:spec_name" json:"spec_name"`        // 规格名字
}

// TableName get sql table name.获取数据库表名
func (m *EsSpecValues) TableName() string {
	return "es_spec_values"
}

// EsSssGoodsData 统计库商品数据(es_sss_goods_data)
type EsSssGoodsData struct {
	ID           int     `gorm:"primaryKey;column:id" json:"-"`             // id
	GoodsID      int     `gorm:"column:goods_id" json:"goods_id"`           // 商品id
	GoodsName    string  `gorm:"column:goods_name" json:"goods_name"`       // 商品名称
	BrandID      int     `gorm:"column:brand_id" json:"brand_id"`           // 品牌id
	CategoryID   int     `gorm:"column:category_id" json:"category_id"`     // 分类id
	CategoryPath string  `gorm:"column:category_path" json:"category_path"` // 分类路径
	SellerID     int     `gorm:"column:seller_id" json:"seller_id"`         // 商家id
	ShopCatID    int     `gorm:"column:shop_cat_id" json:"shop_cat_id"`     // 商家分类id
	Price        float64 `gorm:"column:price" json:"price"`                 // 商品价格
	FavoriteNum  int     `gorm:"column:favorite_num" json:"favorite_num"`   // 收藏数量
	MarketEnable int16   `gorm:"column:market_enable" json:"market_enable"` // 是否上架
}

// TableName get sql table name.获取数据库表名
func (m *EsSssGoodsData) TableName() string {
	return "es_sss_goods_data"
}

// EsSssGoodsPv 商品访问量统计表(es_sss_goods_pv)
type EsSssGoodsPv struct {
	ID        int    `gorm:"primaryKey;column:id" json:"-"`       // 主键id
	SellerID  int    `gorm:"column:seller_id" json:"seller_id"`   // 店铺id
	GoodsID   int    `gorm:"column:goods_id" json:"goods_id"`     // 商品id
	GoodsName string `gorm:"column:goods_name" json:"goods_name"` // 商品名称
	VsYear    int    `gorm:"column:vs_year" json:"vs_year"`       // 年份
	VsNum     int    `gorm:"column:vs_num" json:"vs_num"`         // 访问量
	VsMonth   int    `gorm:"column:vs_month" json:"vs_month"`     // 月份
	VsDay     int    `gorm:"column:vs_day" json:"vs_day"`         // 天
}

// TableName get sql table name.获取数据库表名
func (m *EsSssGoodsPv) TableName() string {
	return "es_sss_goods_pv"
}

// EsSssGoodsPv2015 [...]
type EsSssGoodsPv2015 struct {
	ID        int    `gorm:"column:id" json:"id"`
	SellerID  int    `gorm:"column:seller_id" json:"seller_id"`
	GoodsID   int    `gorm:"column:goods_id" json:"goods_id"`
	GoodsName string `gorm:"column:goods_name" json:"goods_name"`
	VsYear    int    `gorm:"column:vs_year" json:"vs_year"`
	VsMonth   int    `gorm:"column:vs_month" json:"vs_month"`
	VsDay     int    `gorm:"column:vs_day" json:"vs_day"`
	VsNum     int    `gorm:"column:vs_num" json:"vs_num"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssGoodsPv2015) TableName() string {
	return "es_sss_goods_pv_2015"
}

// EsSssGoodsPv2017 [...]
type EsSssGoodsPv2017 struct {
	ID        int    `gorm:"column:id" json:"id"`
	SellerID  int    `gorm:"column:seller_id" json:"seller_id"`
	GoodsID   int    `gorm:"column:goods_id" json:"goods_id"`
	GoodsName string `gorm:"column:goods_name" json:"goods_name"`
	VsYear    int    `gorm:"column:vs_year" json:"vs_year"`
	VsMonth   int    `gorm:"column:vs_month" json:"vs_month"`
	VsDay     int    `gorm:"column:vs_day" json:"vs_day"`
	VsNum     int    `gorm:"column:vs_num" json:"vs_num"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssGoodsPv2017) TableName() string {
	return "es_sss_goods_pv_2017"
}

// EsSssGoodsPv2018 [...]
type EsSssGoodsPv2018 struct {
	ID        int    `gorm:"column:id" json:"id"`
	SellerID  int    `gorm:"column:seller_id" json:"seller_id"`
	GoodsID   int    `gorm:"column:goods_id" json:"goods_id"`
	GoodsName string `gorm:"column:goods_name" json:"goods_name"`
	VsYear    int    `gorm:"column:vs_year" json:"vs_year"`
	VsMonth   int    `gorm:"column:vs_month" json:"vs_month"`
	VsDay     int    `gorm:"column:vs_day" json:"vs_day"`
	VsNum     int    `gorm:"column:vs_num" json:"vs_num"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssGoodsPv2018) TableName() string {
	return "es_sss_goods_pv_2018"
}

// EsSssGoodsPv2019 [...]
type EsSssGoodsPv2019 struct {
	ID        int    `gorm:"primaryKey;column:id" json:"-"`
	SellerID  int    `gorm:"column:seller_id" json:"seller_id"`
	GoodsID   int    `gorm:"column:goods_id" json:"goods_id"`
	GoodsName string `gorm:"column:goods_name" json:"goods_name"`
	VsYear    int    `gorm:"column:vs_year" json:"vs_year"`
	VsMonth   int    `gorm:"column:vs_month" json:"vs_month"`
	VsDay     int    `gorm:"column:vs_day" json:"vs_day"`
	VsNum     int    `gorm:"column:vs_num" json:"vs_num"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssGoodsPv2019) TableName() string {
	return "es_sss_goods_pv_2019"
}

// EsSssMemberRegisterData 注册会员信息(es_sss_member_register_data)
type EsSssMemberRegisterData struct {
	ID         int    `gorm:"primaryKey;column:id" json:"-"`         // id
	MemberID   int    `gorm:"column:member_id" json:"member_id"`     // 会员id
	MemberName string `gorm:"column:member_name" json:"member_name"` // 会员名字
	CreateTime int    `gorm:"column:create_time" json:"create_time"` // 注册日期
}

// TableName get sql table name.获取数据库表名
func (m *EsSssMemberRegisterData) TableName() string {
	return "es_sss_member_register_data"
}

// EsSssOrderData 订单数据(es_sss_order_data)
type EsSssOrderData struct {
	ID             int     `gorm:"primaryKey;column:id" json:"-"`                   // id
	Sn             string  `gorm:"column:sn" json:"sn"`                             // 订单编号
	BuyerID        int     `gorm:"column:buyer_id" json:"buyer_id"`                 // 会员id
	BuyerName      string  `gorm:"column:buyer_name" json:"buyer_name"`             // 商家名称
	SellerID       int     `gorm:"column:seller_id" json:"seller_id"`               // 商家id
	SellerName     string  `gorm:"column:seller_name" json:"seller_name"`           // 商家名称
	OrderStatus    string  `gorm:"column:order_status" json:"order_status"`         // 订单状态
	PayStatus      string  `gorm:"column:pay_status" json:"pay_status"`             // 付款状态
	OrderPrice     float64 `gorm:"column:order_price" json:"order_price"`           // 订单金额
	GoodsNum       int     `gorm:"column:goods_num" json:"goods_num"`               // 订单商品数量
	ShipProvinceID int     `gorm:"column:ship_province_id" json:"ship_province_id"` // 省id
	ShipCityID     int     `gorm:"column:ship_city_id" json:"ship_city_id"`         // 市id
	CreateTime     int     `gorm:"column:create_time" json:"create_time"`           // 订单创建时间
}

// TableName get sql table name.获取数据库表名
func (m *EsSssOrderData) TableName() string {
	return "es_sss_order_data"
}

// EsSssOrderData2015 [...]
type EsSssOrderData2015 struct {
	ID             int     `gorm:"primaryKey;column:id" json:"-"`                   // id
	Sn             string  `gorm:"column:sn" json:"sn"`                             // 订单编号
	BuyerID        int     `gorm:"column:buyer_id" json:"buyer_id"`                 // 会员id
	BuyerName      string  `gorm:"column:buyer_name" json:"buyer_name"`             // 商家名称
	SellerID       int     `gorm:"column:seller_id" json:"seller_id"`               // 商家id
	SellerName     string  `gorm:"column:seller_name" json:"seller_name"`           // 商家名称
	OrderStatus    string  `gorm:"column:order_status" json:"order_status"`         // 订单状态
	PayStatus      string  `gorm:"column:pay_status" json:"pay_status"`             // 付款状态
	OrderPrice     float64 `gorm:"column:order_price" json:"order_price"`           // 订单金额
	GoodsNum       int     `gorm:"column:goods_num" json:"goods_num"`               // 订单商品数量
	ShipProvinceID int     `gorm:"column:ship_province_id" json:"ship_province_id"` // 省id
	ShipCityID     int     `gorm:"column:ship_city_id" json:"ship_city_id"`         // 市id
	CreateTime     int     `gorm:"column:create_time" json:"create_time"`           // 订单创建时间
}

// TableName get sql table name.获取数据库表名
func (m *EsSssOrderData2015) TableName() string {
	return "es_sss_order_data_2015"
}

// EsSssOrderData2017 [...]
type EsSssOrderData2017 struct {
	ID             int     `gorm:"primaryKey;column:id" json:"-"`                   // id
	Sn             string  `gorm:"column:sn" json:"sn"`                             // 订单编号
	BuyerID        int     `gorm:"column:buyer_id" json:"buyer_id"`                 // 会员id
	BuyerName      string  `gorm:"column:buyer_name" json:"buyer_name"`             // 商家名称
	SellerID       int     `gorm:"column:seller_id" json:"seller_id"`               // 商家id
	SellerName     string  `gorm:"column:seller_name" json:"seller_name"`           // 商家名称
	OrderStatus    string  `gorm:"column:order_status" json:"order_status"`         // 订单状态
	PayStatus      string  `gorm:"column:pay_status" json:"pay_status"`             // 付款状态
	OrderPrice     float64 `gorm:"column:order_price" json:"order_price"`           // 订单金额
	GoodsNum       int     `gorm:"column:goods_num" json:"goods_num"`               // 订单商品数量
	ShipProvinceID int     `gorm:"column:ship_province_id" json:"ship_province_id"` // 省id
	ShipCityID     int     `gorm:"column:ship_city_id" json:"ship_city_id"`         // 市id
	CreateTime     int     `gorm:"column:create_time" json:"create_time"`           // 订单创建时间
}

// TableName get sql table name.获取数据库表名
func (m *EsSssOrderData2017) TableName() string {
	return "es_sss_order_data_2017"
}

// EsSssOrderData2018 [...]
type EsSssOrderData2018 struct {
	ID             int     `gorm:"primaryKey;column:id" json:"-"`                   // id
	Sn             string  `gorm:"column:sn" json:"sn"`                             // 订单编号
	BuyerID        int     `gorm:"column:buyer_id" json:"buyer_id"`                 // 会员id
	BuyerName      string  `gorm:"column:buyer_name" json:"buyer_name"`             // 商家名称
	SellerID       int     `gorm:"column:seller_id" json:"seller_id"`               // 商家id
	SellerName     string  `gorm:"column:seller_name" json:"seller_name"`           // 商家名称
	OrderStatus    string  `gorm:"column:order_status" json:"order_status"`         // 订单状态
	PayStatus      string  `gorm:"column:pay_status" json:"pay_status"`             // 付款状态
	OrderPrice     float64 `gorm:"column:order_price" json:"order_price"`           // 订单金额
	GoodsNum       int     `gorm:"column:goods_num" json:"goods_num"`               // 订单商品数量
	ShipProvinceID int     `gorm:"column:ship_province_id" json:"ship_province_id"` // 省id
	ShipCityID     int     `gorm:"column:ship_city_id" json:"ship_city_id"`         // 市id
	CreateTime     int     `gorm:"column:create_time" json:"create_time"`           // 订单创建时间
}

// TableName get sql table name.获取数据库表名
func (m *EsSssOrderData2018) TableName() string {
	return "es_sss_order_data_2018"
}

// EsSssOrderData2019 [...]
type EsSssOrderData2019 struct {
	ID             int     `gorm:"primaryKey;column:id" json:"-"`
	Sn             string  `gorm:"column:sn" json:"sn"`
	BuyerID        int     `gorm:"column:buyer_id" json:"buyer_id"`
	BuyerName      string  `gorm:"column:buyer_name" json:"buyer_name"`
	SellerID       int     `gorm:"column:seller_id" json:"seller_id"`
	SellerName     string  `gorm:"column:seller_name" json:"seller_name"`
	OrderStatus    string  `gorm:"column:order_status" json:"order_status"`
	PayStatus      string  `gorm:"column:pay_status" json:"pay_status"`
	OrderPrice     float64 `gorm:"column:order_price" json:"order_price"`
	GoodsNum       int     `gorm:"column:goods_num" json:"goods_num"`
	ShipProvinceID int     `gorm:"column:ship_province_id" json:"ship_province_id"`
	ShipCityID     int     `gorm:"column:ship_city_id" json:"ship_city_id"`
	CreateTime     int     `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssOrderData2019) TableName() string {
	return "es_sss_order_data_2019"
}

// EsSssOrderGoodsData 订单商品表(es_sss_order_goods_data)
type EsSssOrderGoodsData struct {
	ID           int     `gorm:"primaryKey;column:id" json:"-"`             // id
	OrderSn      string  `gorm:"column:order_sn" json:"order_sn"`           // 订单编号
	GoodsID      int     `gorm:"column:goods_id" json:"goods_id"`           // 商品id
	GoodsName    string  `gorm:"column:goods_name" json:"goods_name"`       // 商品名称
	GoodsNum     int     `gorm:"column:goods_num" json:"goods_num"`         // 数量
	Price        float64 `gorm:"column:price" json:"price"`                 // 商品单价
	SubTotal     float64 `gorm:"column:sub_total" json:"sub_total"`         // 小计
	CategoryPath string  `gorm:"column:category_path" json:"category_path"` // 分类path
	CategoryID   int     `gorm:"column:category_id" json:"category_id"`     // 分类id
	CreateTime   int     `gorm:"column:create_time" json:"create_time"`     // 创建时间
	IndustryID   int64   `gorm:"column:industry_id" json:"industry_id"`     // 行业id
}

// TableName get sql table name.获取数据库表名
func (m *EsSssOrderGoodsData) TableName() string {
	return "es_sss_order_goods_data"
}

// EsSssOrderGoodsData2015 [...]
type EsSssOrderGoodsData2015 struct {
	ID           int     `gorm:"column:id" json:"id"`
	OrderSn      string  `gorm:"column:order_sn" json:"order_sn"`
	GoodsID      int     `gorm:"column:goods_id" json:"goods_id"`
	GoodsName    string  `gorm:"column:goods_name" json:"goods_name"`
	GoodsNum     int     `gorm:"column:goods_num" json:"goods_num"`
	Price        float64 `gorm:"column:price" json:"price"`
	SubTotal     float64 `gorm:"column:sub_total" json:"sub_total"`
	CategoryPath string  `gorm:"column:category_path" json:"category_path"`
	CategoryID   int     `gorm:"column:category_id" json:"category_id"`
	CreateTime   int     `gorm:"column:create_time" json:"create_time"`
	IndustryID   int     `gorm:"column:industry_id" json:"industry_id"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssOrderGoodsData2015) TableName() string {
	return "es_sss_order_goods_data_2015"
}

// EsSssOrderGoodsData2017 [...]
type EsSssOrderGoodsData2017 struct {
	ID           int     `gorm:"column:id" json:"id"`
	OrderSn      string  `gorm:"column:order_sn" json:"order_sn"`
	GoodsID      int     `gorm:"column:goods_id" json:"goods_id"`
	GoodsName    string  `gorm:"column:goods_name" json:"goods_name"`
	GoodsNum     int     `gorm:"column:goods_num" json:"goods_num"`
	Price        float64 `gorm:"column:price" json:"price"`
	SubTotal     float64 `gorm:"column:sub_total" json:"sub_total"`
	CategoryPath string  `gorm:"column:category_path" json:"category_path"`
	CategoryID   int     `gorm:"column:category_id" json:"category_id"`
	CreateTime   int     `gorm:"column:create_time" json:"create_time"`
	IndustryID   int     `gorm:"column:industry_id" json:"industry_id"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssOrderGoodsData2017) TableName() string {
	return "es_sss_order_goods_data_2017"
}

// EsSssOrderGoodsData2018 [...]
type EsSssOrderGoodsData2018 struct {
	ID           int     `gorm:"primaryKey;column:id" json:"-"`             // id
	OrderSn      string  `gorm:"column:order_sn" json:"order_sn"`           // 订单编号
	GoodsID      int     `gorm:"column:goods_id" json:"goods_id"`           // 商品id
	GoodsName    string  `gorm:"column:goods_name" json:"goods_name"`       // 商品名称
	GoodsNum     int     `gorm:"column:goods_num" json:"goods_num"`         // 数量
	Price        float64 `gorm:"column:price" json:"price"`                 // 商品单价
	SubTotal     float64 `gorm:"column:sub_total" json:"sub_total"`         // 小计
	CategoryPath string  `gorm:"column:category_path" json:"category_path"` // 分类path
	CategoryID   int     `gorm:"column:category_id" json:"category_id"`     // 分类id
	CreateTime   int     `gorm:"column:create_time" json:"create_time"`     // 创建时间
	IndustryID   int64   `gorm:"column:industry_id" json:"industry_id"`     // 行业id
}

// TableName get sql table name.获取数据库表名
func (m *EsSssOrderGoodsData2018) TableName() string {
	return "es_sss_order_goods_data_2018"
}

// EsSssOrderGoodsData2019 [...]
type EsSssOrderGoodsData2019 struct {
	ID           int     `gorm:"primaryKey;column:id" json:"-"`
	OrderSn      string  `gorm:"column:order_sn" json:"order_sn"`
	GoodsID      int     `gorm:"column:goods_id" json:"goods_id"`
	GoodsName    string  `gorm:"column:goods_name" json:"goods_name"`
	GoodsNum     int     `gorm:"column:goods_num" json:"goods_num"`
	Price        float64 `gorm:"column:price" json:"price"`
	SubTotal     float64 `gorm:"column:sub_total" json:"sub_total"`
	CategoryPath string  `gorm:"column:category_path" json:"category_path"`
	CategoryID   int     `gorm:"column:category_id" json:"category_id"`
	CreateTime   int     `gorm:"column:create_time" json:"create_time"`
	IndustryID   int     `gorm:"column:industry_id" json:"industry_id"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssOrderGoodsData2019) TableName() string {
	return "es_sss_order_goods_data_2019"
}

// EsSssRefundData 退货数据(es_sss_refund_data)
type EsSssRefundData struct {
	ID          int     `gorm:"primaryKey;column:id" json:"-"`           // id
	SellerID    int     `gorm:"column:seller_id" json:"seller_id"`       // 商家id
	OrderSn     string  `gorm:"column:order_sn" json:"order_sn"`         // 订单sn
	RefundSn    string  `gorm:"column:refund_sn" json:"refund_sn"`       // 售后订单sn
	RefundPrice float64 `gorm:"column:refund_price" json:"refund_price"` // 退还金额
	CreateTime  int     `gorm:"column:create_time" json:"create_time"`   // 创建日期
}

// TableName get sql table name.获取数据库表名
func (m *EsSssRefundData) TableName() string {
	return "es_sss_refund_data"
}

// EsSssRefundData2015 [...]
type EsSssRefundData2015 struct {
	ID          int     `gorm:"column:id" json:"id"`
	SellerID    int     `gorm:"column:seller_id" json:"seller_id"`
	RefundSn    string  `gorm:"column:refund_sn" json:"refund_sn"`
	OrderSn     string  `gorm:"column:order_sn" json:"order_sn"`
	RefundPrice float64 `gorm:"column:refund_price" json:"refund_price"`
	CreateTime  int     `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssRefundData2015) TableName() string {
	return "es_sss_refund_data_2015"
}

// EsSssRefundData2017 [...]
type EsSssRefundData2017 struct {
	ID          int     `gorm:"column:id" json:"id"`
	SellerID    int     `gorm:"column:seller_id" json:"seller_id"`
	RefundSn    string  `gorm:"column:refund_sn" json:"refund_sn"`
	OrderSn     string  `gorm:"column:order_sn" json:"order_sn"`
	RefundPrice float64 `gorm:"column:refund_price" json:"refund_price"`
	CreateTime  int     `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssRefundData2017) TableName() string {
	return "es_sss_refund_data_2017"
}

// EsSssRefundData2018 [...]
type EsSssRefundData2018 struct {
	ID          int     `gorm:"column:id" json:"id"`
	SellerID    int     `gorm:"column:seller_id" json:"seller_id"`
	RefundSn    string  `gorm:"column:refund_sn" json:"refund_sn"`
	OrderSn     string  `gorm:"column:order_sn" json:"order_sn"`
	RefundPrice float64 `gorm:"column:refund_price" json:"refund_price"`
	CreateTime  int     `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssRefundData2018) TableName() string {
	return "es_sss_refund_data_2018"
}

// EsSssRefundData2019 [...]
type EsSssRefundData2019 struct {
	ID          int     `gorm:"primaryKey;column:id" json:"-"`
	SellerID    int     `gorm:"column:seller_id" json:"seller_id"`
	RefundSn    string  `gorm:"column:refund_sn" json:"refund_sn"`
	OrderSn     string  `gorm:"column:order_sn" json:"order_sn"`
	RefundPrice float64 `gorm:"column:refund_price" json:"refund_price"`
	CreateTime  int     `gorm:"column:create_time" json:"create_time"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssRefundData2019) TableName() string {
	return "es_sss_refund_data_2019"
}

// EsSssShopData 商品数据表(es_sss_shop_data)
type EsSssShopData struct {
	ID          int    `gorm:"primaryKey;column:id" json:"-"`           // 主键id
	SellerID    int    `gorm:"column:seller_id" json:"seller_id"`       // 店铺id
	SellerName  string `gorm:"column:seller_name" json:"seller_name"`   // 店铺名称
	FavoriteNum int    `gorm:"column:favorite_num" json:"favorite_num"` // 店铺被收藏数量
	ShopDisable string `gorm:"column:shop_disable" json:"shop_disable"` // 店铺是否开启
}

// TableName get sql table name.获取数据库表名
func (m *EsSssShopData) TableName() string {
	return "es_sss_shop_data"
}

// EsSssShopPv 店铺访问量统计表(es_sss_shop_pv)
type EsSssShopPv struct {
	ID       int `gorm:"primaryKey;column:id" json:"-"`     // 主键id
	SellerID int `gorm:"column:seller_id" json:"seller_id"` // 店铺id
	VsYear   int `gorm:"column:vs_year" json:"vs_year"`     // 年份
	VsMonth  int `gorm:"column:vs_month" json:"vs_month"`   // 月份
	VsDay    int `gorm:"column:vs_day" json:"vs_day"`       // 日期
	VsNum    int `gorm:"column:vs_num" json:"vs_num"`       // 访问量
}

// TableName get sql table name.获取数据库表名
func (m *EsSssShopPv) TableName() string {
	return "es_sss_shop_pv"
}

// EsSssShopPv2015 [...]
type EsSssShopPv2015 struct {
	ID       int `gorm:"column:id" json:"id"`
	SellerID int `gorm:"column:seller_id" json:"seller_id"`
	VsYear   int `gorm:"column:vs_year" json:"vs_year"`
	VsMonth  int `gorm:"column:vs_month" json:"vs_month"`
	VsDay    int `gorm:"column:vs_day" json:"vs_day"`
	VsNum    int `gorm:"column:vs_num" json:"vs_num"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssShopPv2015) TableName() string {
	return "es_sss_shop_pv_2015"
}

// EsSssShopPv2017 [...]
type EsSssShopPv2017 struct {
	ID       int `gorm:"column:id" json:"id"`
	SellerID int `gorm:"column:seller_id" json:"seller_id"`
	VsYear   int `gorm:"column:vs_year" json:"vs_year"`
	VsMonth  int `gorm:"column:vs_month" json:"vs_month"`
	VsDay    int `gorm:"column:vs_day" json:"vs_day"`
	VsNum    int `gorm:"column:vs_num" json:"vs_num"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssShopPv2017) TableName() string {
	return "es_sss_shop_pv_2017"
}

// EsSssShopPv2018 [...]
type EsSssShopPv2018 struct {
	ID       int `gorm:"column:id" json:"id"`
	SellerID int `gorm:"column:seller_id" json:"seller_id"`
	VsYear   int `gorm:"column:vs_year" json:"vs_year"`
	VsMonth  int `gorm:"column:vs_month" json:"vs_month"`
	VsDay    int `gorm:"column:vs_day" json:"vs_day"`
	VsNum    int `gorm:"column:vs_num" json:"vs_num"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssShopPv2018) TableName() string {
	return "es_sss_shop_pv_2018"
}

// EsSssShopPv2019 [...]
type EsSssShopPv2019 struct {
	ID       int `gorm:"primaryKey;column:id" json:"-"`
	SellerID int `gorm:"column:seller_id" json:"seller_id"`
	VsYear   int `gorm:"column:vs_year" json:"vs_year"`
	VsMonth  int `gorm:"column:vs_month" json:"vs_month"`
	VsDay    int `gorm:"column:vs_day" json:"vs_day"`
	VsNum    int `gorm:"column:vs_num" json:"vs_num"`
}

// TableName get sql table name.获取数据库表名
func (m *EsSssShopPv2019) TableName() string {
	return "es_sss_shop_pv_2019"
}

// EsTagGoods 标签商品关联(es_tag_goods)
type EsTagGoods struct {
	TagID   int `gorm:"column:tag_id" json:"tag_id"`     // 标签id
	GoodsID int `gorm:"column:goods_id" json:"goods_id"` // 商品id
}

// TableName get sql table name.获取数据库表名
func (m *EsTagGoods) TableName() string {
	return "es_tag_goods"
}

// EsTags 商品标签(es_tags)
type EsTags struct {
	TagID    int    `gorm:"primaryKey;column:tag_id" json:"-"` // 主键
	TagName  string `gorm:"column:tag_name" json:"tag_name"`   // 标签名字
	SellerID int    `gorm:"column:seller_id" json:"seller_id"` // 所属卖家
	Mark     string `gorm:"column:mark" json:"mark"`           // 关键字
}

// TableName get sql table name.获取数据库表名
func (m *EsTags) TableName() string {
	return "es_tags"
}

// EsTransactionRecord 交易记录表(es_transaction_record)
type EsTransactionRecord struct {
	RecordID int     `gorm:"primaryKey;column:record_id" json:"-"` // 主键ID
	OrderSn  string  `gorm:"column:order_sn" json:"order_sn"`      // 订单编号
	GoodsID  int     `gorm:"column:goods_id" json:"goods_id"`      // 商品ID
	GoodsNum int     `gorm:"column:goods_num" json:"goods_num"`    // 商品数量
	RogTime  int64   `gorm:"column:rog_time" json:"rog_time"`      // 确认收货时间
	Uname    string  `gorm:"column:uname" json:"uname"`            // 用户名
	Price    float64 `gorm:"column:price" json:"price"`            // 交易价格
	MemberID int     `gorm:"column:member_id" json:"member_id"`    // 会员ID
}

// TableName get sql table name.获取数据库表名
func (m *EsTransactionRecord) TableName() string {
	return "es_transaction_record"
}

// EsUpgradeLog 升级日志(es_upgrade_log)
type EsUpgradeLog struct {
	ID         int64  `gorm:"primaryKey;column:id" json:"-"`           // id
	MemberID   int64  `gorm:"column:member_id" json:"member_id"`       // 会员id
	MemberName string `gorm:"column:member_name" json:"member_name"`   // 会员名称
	Type       string `gorm:"column:type" json:"type"`                 // 切换类型
	OldTplID   int64  `gorm:"column:old_tpl_id" json:"old_tpl_id"`     // 旧的模板id
	NewTplID   int64  `gorm:"column:new_tpl_id" json:"new_tpl_id"`     // 新的模板id
	NewTplName string `gorm:"column:new_tpl_name" json:"new_tpl_name"` // 新的模板名称
	OldTplName string `gorm:"column:old_tpl_name" json:"old_tpl_name"` // 旧的模板名称
	CreateTime int    `gorm:"column:create_time" json:"create_time"`   // 创建日期
}

// TableName get sql table name.获取数据库表名
func (m *EsUpgradeLog) TableName() string {
	return "es_upgrade_log"
}

// EsUploader 存储方案(es_uploader)
type EsUploader struct {
	ID     int    `gorm:"primaryKey;column:id" json:"-"` // 储存id
	Name   string `gorm:"column:name" json:"name"`       // 存储名称
	Open   int    `gorm:"column:open" json:"open"`       // 是否开启
	Config string `gorm:"column:config" json:"config"`   // 存储配置
	Bean   string `gorm:"column:bean" json:"bean"`       // 存储插件id
}

// TableName get sql table name.获取数据库表名
func (m *EsUploader) TableName() string {
	return "es_uploader"
}

// EsWaybill 电子面单(es_waybill)
type EsWaybill struct {
	ID     int    `gorm:"primaryKey;column:id" json:"-"` // 电子面单id
	Name   string `gorm:"column:name" json:"name"`       // 名称
	Open   int    `gorm:"column:open" json:"open"`       // 是否开启
	Config string `gorm:"column:config" json:"config"`   // 电子面单配置
	Bean   string `gorm:"column:bean" json:"bean"`       // beanid
}

// TableName get sql table name.获取数据库表名
func (m *EsWaybill) TableName() string {
	return "es_waybill"
}

// EsWechatMsgTemplate [...]
type EsWechatMsgTemplate struct {
	ID         int    `gorm:"primaryKey;column:id" json:"-"`           // 主键
	MsgTmpName string `gorm:"column:msg_tmp_name" json:"msg_tmp_name"` // 模板名称
	MsgTmpSn   string `gorm:"column:msg_tmp_sn" json:"msg_tmp_sn"`     // 消息编号
	TemplateID string `gorm:"column:template_id" json:"template_id"`   // 模板ID
	MsgFirst   string `gorm:"column:msg_first" json:"msg_first"`       // 消息开头文字
	MsgRemark  string `gorm:"column:msg_remark" json:"msg_remark"`     // 消息结尾备注文字
	IsOpen     int16  `gorm:"column:is_open" json:"is_open"`           // 是否开启
	TmpType    string `gorm:"column:tmp_type" json:"tmp_type"`         // 模板类型，枚举
}

// TableName get sql table name.获取数据库表名
func (m *EsWechatMsgTemplate) TableName() string {
	return "es_wechat_msg_template"
}

// EsWithdrawApply 提现申请(es_withdraw_apply)
type EsWithdrawApply struct {
	ID             int     `gorm:"primaryKey;column:id" json:"-"`                 // id
	ApplyMoney     float64 `gorm:"column:apply_money" json:"apply_money"`         // 提现金额
	MemberID       int     `gorm:"column:member_id" json:"member_id"`             // 会员id
	MemberName     string  `gorm:"column:member_name" json:"member_name"`         // 会员名称
	ApplyRemark    string  `gorm:"column:apply_remark" json:"apply_remark"`       // 申请备注
	InspectRemark  string  `gorm:"column:inspect_remark" json:"inspect_remark"`   // 审核备注
	TransferRemark string  `gorm:"column:transfer_remark" json:"transfer_remark"` // 转账备注
	ApplyTime      int     `gorm:"column:apply_time" json:"apply_time"`           // 申请时间
	InspectTime    int     `gorm:"column:inspect_time" json:"inspect_time"`       // 审核时间
	TransferTime   int     `gorm:"column:transfer_time" json:"transfer_time"`     // 转账时间
	Status         string  `gorm:"column:status" json:"status"`                   // 状态
	Sn             string  `gorm:"column:sn" json:"sn"`                           // 编号
	IP             string  `gorm:"column:ip" json:"ip"`                           // ip地址
}

// TableName get sql table name.获取数据库表名
func (m *EsWithdrawApply) TableName() string {
	return "es_withdraw_apply"
}

// EsWithdrawSetting 分销提现设置(es_withdraw_setting)
type EsWithdrawSetting struct {
	ID       int    `gorm:"primaryKey;column:id" json:"-"`     // id
	MemberID int    `gorm:"column:member_id" json:"member_id"` // 用户id
	Param    string `gorm:"column:param" json:"param"`         // 参数
}

// TableName get sql table name.获取数据库表名
func (m *EsWithdrawSetting) TableName() string {
	return "es_withdraw_setting"
}
