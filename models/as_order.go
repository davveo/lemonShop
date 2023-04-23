package models

// EsAsOrder 售后服务单表(es_as_order)
type EsAsOrder struct {
	ID            int    `gorm:"primaryKey;column:id" json:"-"`               // 主键ID
	Sn            string `gorm:"column:sn" json:"sn"`                         // 售后单号
	OrderSn       string `gorm:"column:order_sn" json:"order_sn"`             // 订单编号
	MemberID      int    `gorm:"column:member_id" json:"member_id"`           // 用户ID
	MemberName    string `gorm:"column:member_name" json:"member_name"`       // 用户名
	SellerID      int    `gorm:"column:seller_id" json:"seller_id"`           // 商家ID
	SellerName    string `gorm:"column:seller_name" json:"seller_name"`       // 商家名称
	CreateTime    int64  `gorm:"column:create_time" json:"create_time"`       // 创建时间
	Mobile        string `gorm:"column:mobile" json:"mobile"`                 // 手机号
	ServiceType   string `gorm:"column:service_type" json:"service_type"`     // 售后类型 RETURN_GOODS：退货，CHANGE_GOODS：换货，SUPPLY_AGAIN_GOODS：补发货品，ORDER_CANCEL：取消订单（订单已付款且未收货之前）
	ServiceStatus string `gorm:"column:service_status" json:"service_status"` // 售后单状态 APPLY：待审核，PASS：审核通过，REFUSE：审核拒绝，WAIT_FOR_MANUAL：待人工处理，STOCK_IN：退货入库，REFUNDING：退款中，REFUNDFAIL：退款失败，COMPLETE：完成
	Reason        string `gorm:"column:reason" json:"reason"`                 // 申请原因
	ApplyVouchers string `gorm:"column:apply_vouchers" json:"apply_vouchers"` // 申请凭证
	ProblemDesc   string `gorm:"column:problem_desc" json:"problem_desc"`     // 问题描述
	GoodsJSON     string `gorm:"column:goods_json" json:"goods_json"`         // 售后商品信息json
	Disabled      string `gorm:"column:disabled" json:"disabled"`             // 删除状态 DELETED：已删除 NORMAL：正常
	AuditRemark   string `gorm:"column:audit_remark" json:"audit_remark"`     // 商家审核备注
	StockRemark   string `gorm:"column:stock_remark" json:"stock_remark"`     // 商家入库备注
	RefundRemark  string `gorm:"column:refund_remark" json:"refund_remark"`   // 平台退款备注
	CloseReason   string `gorm:"column:close_reason" json:"close_reason"`     // 取消原因
	ReturnAddr    string `gorm:"column:return_addr" json:"return_addr"`       // 退货地址信息
	NewOrderSn    string `gorm:"column:new_order_sn" json:"new_order_sn"`     // 新订单编号
	CreateChannel string `gorm:"column:create_channel" json:"create_channel"` // 售后服务单创建渠道 NORMAL：正常渠道创建，PINTUAN：拼团失败自动创建
}

// TableName get sql table name.获取数据库表名
func (m *EsAsOrder) TableName() string {
	return "es_as_order"
}
