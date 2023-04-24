package models

// EsSssMemberRegisterData 注册会员信息(es_sss_member_register_data)
type EsSssMemberRegisterData struct {
	ID         int    `gorm:"primaryKey;column:id" json:"-"`         // id
	MemberID   int    `gorm:"column:member_id" json:"member_id"`     // 会员id
	MemberName string `gorm:"column:member_name" json:"member_name"` // 会员名字
	CreateTime int    `gorm:"column:create_time" json:"create_time"` // 注册日期
}

// TableName get sql table name.获取数据库表名
func (m *EsSssMemberRegisterData) TableName() string {
	return "es_sss_member_register_data"
}
