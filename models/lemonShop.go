package models

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
