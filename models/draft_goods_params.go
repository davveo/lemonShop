package models

// EsDraftGoodsParams 草稿商品参数表(es_draft_goods_params)
type EsDraftGoodsParams struct {
	ID           int    `gorm:"primaryKey;column:id" json:"-"`               // ID
	DraftGoodsID int    `gorm:"column:draft_goods_id" json:"draft_goods_id"` // 草稿ID
	ParamID      int    `gorm:"column:param_id" json:"param_id"`             // 参数ID
	ParamName    string `gorm:"column:param_name" json:"param_name"`         // 参数名
	ParamValue   string `gorm:"column:param_value" json:"param_value"`       // 参数值
}

// TableName get sql table name.获取数据库表名
func (m *EsDraftGoodsParams) TableName() string {
	return "es_draft_goods_params"
}
