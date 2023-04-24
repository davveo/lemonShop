package models

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