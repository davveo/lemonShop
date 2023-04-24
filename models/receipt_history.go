package models

// EsReceiptHistory 会员开票历史记录表(es_receipt_history)
type EsReceiptHistory struct {
	HistoryID      int     `gorm:"primaryKey;column:history_id" json:"-"`         // 主键ID
	OrderSn        string  `gorm:"column:order_sn" json:"order_sn"`               // 订单编号
	OrderPrice     float64 `gorm:"column:order_price" json:"order_price"`         // 订单金额
	SellerID       int     `gorm:"column:seller_id" json:"seller_id"`             // 开票商家ID
	SellerName     string  `gorm:"column:seller_name" json:"seller_name"`         // 开票商家
	MemberID       int     `gorm:"column:member_id" json:"member_id"`             // 会员ID
	Status         int16   `gorm:"column:status" json:"status"`                   // 开票状态 0：未开，1：已开
	ReceiptMethod  string  `gorm:"column:receipt_method" json:"receipt_method"`   // 开票方式(针对增值税专用发票)
	ReceiptType    string  `gorm:"column:receipt_type" json:"receipt_type"`       // 发票类型 ELECTRO：电子普通发票，VATORDINARY：增值税普通发票，VATOSPECIAL：增值税专用发票
	LogiID         int     `gorm:"column:logi_id" json:"logi_id"`                 // 物流公司ID
	LogiName       string  `gorm:"column:logi_name" json:"logi_name"`             // 物流公司名称
	LogiCode       string  `gorm:"column:logi_code" json:"logi_code"`             // 快递单号
	ReceiptTitle   string  `gorm:"column:receipt_title" json:"receipt_title"`     // 发票抬头
	ReceiptContent string  `gorm:"column:receipt_content" json:"receipt_content"` // 发票内容
	TaxNo          string  `gorm:"column:tax_no" json:"tax_no"`                   // 纳税人识别号
	RegAddr        string  `gorm:"column:reg_addr" json:"reg_addr"`               // 注册地址
	RegTel         string  `gorm:"column:reg_tel" json:"reg_tel"`                 // 注册电话
	BankName       string  `gorm:"column:bank_name" json:"bank_name"`             // 开户银行
	BankAccount    string  `gorm:"column:bank_account" json:"bank_account"`       // 银行账户
	MemberName     string  `gorm:"column:member_name" json:"member_name"`         // 收票人名称
	MemberMobile   string  `gorm:"column:member_mobile" json:"member_mobile"`     // 收票人手机号
	MemberEmail    string  `gorm:"column:member_email" json:"member_email"`       // 收票人邮箱
	ProvinceID     int     `gorm:"column:province_id" json:"province_id"`         // 收票地址-所属省份ID
	CityID         int     `gorm:"column:city_id" json:"city_id"`                 // 收票地址-所属城市ID
	CountyID       int     `gorm:"column:county_id" json:"county_id"`             // 收票地址-所属区县ID
	TownID         int     `gorm:"column:town_id" json:"town_id"`                 // 收票地址-所属乡镇ID
	Province       string  `gorm:"column:province" json:"province"`               // 收票地址-所属省份
	City           string  `gorm:"column:city" json:"city"`                       // 收票地址-所属城市
	County         string  `gorm:"column:county" json:"county"`                   // 收票地址-所属区县
	Town           string  `gorm:"column:town" json:"town"`                       // 收票地址-所属乡镇
	DetailAddr     string  `gorm:"column:detail_addr" json:"detail_addr"`         // 收票地址-详细地址
	AddTime        int64   `gorm:"column:add_time" json:"add_time"`               // 申请开票日期
	GoodsJSON      string  `gorm:"column:goods_json" json:"goods_json"`           // 订单商品信息
	OrderStatus    string  `gorm:"column:order_status" json:"order_status"`       // 订单出库状态，NEW/CONFIRM
	Uname          string  `gorm:"column:uname" json:"uname"`                     // 会员名称
}

// TableName get sql table name.获取数据库表名
func (m *EsReceiptHistory) TableName() string {
	return "es_receipt_history"
}
