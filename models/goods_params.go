package models

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
