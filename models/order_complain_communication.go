package models

// EsOrderComplainCommunication [...]
type EsOrderComplainCommunication struct {
	CommunicationID int    `gorm:"primaryKey;column:communication_id" json:"-"` // 主键
	ComplainID      int    `gorm:"column:complain_id" json:"complain_id"`       // 投诉id
	Content         string `gorm:"column:content" json:"content"`               // 对话内容
	CreateTime      int64  `gorm:"column:create_time" json:"create_time"`       // 对话时间
	Owner           string `gorm:"column:owner" json:"owner"`                   // 所属，买家/卖家
	OwnerName       string `gorm:"column:owner_name" json:"owner_name"`         // 对话所属名称
	OwnerID         int    `gorm:"column:owner_id" json:"owner_id"`             // 对话所属id,卖家id/买家id
}

// TableName get sql table name.获取数据库表名
func (m *EsOrderComplainCommunication) TableName() string {
	return "es_order_complain_communication"
}
