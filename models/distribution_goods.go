package models

// EsDistributionGoods 分销商品返现(es_distribution_goods)
type EsDistributionGoods struct {
	ID           int64   `gorm:"primaryKey;column:id" json:"-"`             // id
	GoodsID      int64   `gorm:"column:goods_id" json:"goods_id"`           // 商品id
	Grade1Rebate float64 `gorm:"column:grade1_rebate" json:"grade1_rebate"` // 一级返现
	Grade2Rebate float64 `gorm:"column:grade2_rebate" json:"grade2_rebate"` // 二级返现
}

// TableName get sql table name.获取数据库表名
func (m *EsDistributionGoods) TableName() string {
	return "es_distribution_goods"
}
