package models

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