package models

// EsDistributionOrder 分销订单(es_distribution_order)
type EsDistributionOrder struct {
	ID                  int64   `gorm:"primaryKey;column:id" json:"-"`                             // id
	OrderID             int64   `gorm:"column:order_id" json:"order_id"`                           // 订单id
	OrderSn             string  `gorm:"column:order_sn" json:"order_sn"`                           // 订单sn
	BuyerMemberID       int64   `gorm:"column:buyer_member_id" json:"buyer_member_id"`             // 购买会员id
	BuyerMemberName     string  `gorm:"column:buyer_member_name" json:"buyer_member_name"`         // 购买会员名称
	MemberIDLv1         int64   `gorm:"column:member_id_lv1" json:"member_id_lv1"`                 // 一级分销商id
	MemberIDLv2         int64   `gorm:"column:member_id_lv2" json:"member_id_lv2"`                 // 二级分销商id
	BillID              int64   `gorm:"column:bill_id" json:"bill_id"`                             // 结算单id
	SettleCycle         int64   `gorm:"column:settle_cycle" json:"settle_cycle"`                   // 解冻日期
	CreateTime          int64   `gorm:"column:create_time" json:"create_time"`                     // 订单创建时间
	OrderPrice          float64 `gorm:"column:order_price" json:"order_price"`                     // 订单金额
	Grade1Rebate        float64 `gorm:"column:grade1_rebate" json:"grade1_rebate"`                 // 1级提成金额
	Grade2Rebate        float64 `gorm:"column:grade2_rebate" json:"grade2_rebate"`                 // 2级提成金额
	Grade1SellbackPrice float64 `gorm:"column:grade1_sellback_price" json:"grade1_sellback_price"` // 1级退款金额
	Grade2SellbackPrice float64 `gorm:"column:grade2_sellback_price" json:"grade2_sellback_price"` // 2级退款金额
	IsReturn            int16   `gorm:"column:is_return" json:"is_return"`                         // 是否退货
	ReturnMoney         float64 `gorm:"column:return_money" json:"return_money"`                   // 退款金额
	IsWithdraw          int16   `gorm:"column:is_withdraw" json:"is_withdraw"`                     // 是否已经提现
	SellerID            int64   `gorm:"column:seller_id" json:"seller_id"`                         // 店铺id
	SellerName          string  `gorm:"column:seller_name" json:"seller_name"`                     // 店铺名称
	Lv1Point            float64 `gorm:"column:lv1_point" json:"lv1_point"`                         // 一级饭点比例
	Lv2Point            float64 `gorm:"column:lv2_point" json:"lv2_point"`                         // 二级返点比例
	GoodsRebate         string  `gorm:"column:goods_rebate" json:"goods_rebate"`                   // 商品返现描述详细
}

// TableName get sql table name.获取数据库表名
func (m *EsDistributionOrder) TableName() string {
	return "es_distribution_order"
}
