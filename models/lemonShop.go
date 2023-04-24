package models

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
