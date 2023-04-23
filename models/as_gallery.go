package models

// EsAsGallery 售后服务图片信息表(es_as_gallery)
type EsAsGallery struct {
	ID        int    `gorm:"primaryKey;column:id" json:"-"`       // 主键ID
	ServiceSn string `gorm:"column:service_sn" json:"service_sn"` // 售后单号
	Img       string `gorm:"column:img" json:"img"`               // 图片链接
}

// TableName get sql table name.获取数据库表名
func (m *EsAsGallery) TableName() string {
	return "es_as_gallery"
}
