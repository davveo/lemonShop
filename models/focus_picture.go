package models

// EsFocusPicture 焦点图(es_focus_picture)
type EsFocusPicture struct {
	ID             int    `gorm:"primaryKey;column:id" json:"-"`                 // 主键id
	PicURL         string `gorm:"column:pic_url" json:"pic_url"`                 // 图片地址
	OperationType  string `gorm:"column:operation_type" json:"operation_type"`   // 操作类型
	OperationParam string `gorm:"column:operation_param" json:"operation_param"` // 操作参数
	OperationURL   string `gorm:"column:operation_url" json:"operation_url"`     // 操作地址
	ClientType     string `gorm:"column:client_type" json:"client_type"`         // 客户端类型 pc:pc楼层mobile:移动端楼层
}

// TableName get sql table name.获取数据库表名
func (m *EsFocusPicture) TableName() string {
	return "es_focus_picture"
}
