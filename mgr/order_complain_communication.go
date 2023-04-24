package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type OrderComplainCommunicationMgr struct {
	*_BaseMgr
}

// NewOrderComplainCommunicationMgr open func
func NewOrderComplainCommunicationMgr(db db.Repo) *OrderComplainCommunicationMgr {
	if db == nil {
		panic(fmt.Errorf("NewOrderComplainCommunicationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &OrderComplainCommunicationMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_order_complain_communication"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *OrderComplainCommunicationMgr) GetTableName() string {
	return "es_order_complain_communication"
}

// Reset 重置gorm会话
func (obj *OrderComplainCommunicationMgr) Reset() *OrderComplainCommunicationMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *OrderComplainCommunicationMgr) Get() (result models.EsOrderComplainCommunication, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *OrderComplainCommunicationMgr) Gets() (results []*models.EsOrderComplainCommunication, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *OrderComplainCommunicationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCommunicationID communication_id获取 主键
func (obj *OrderComplainCommunicationMgr) WithCommunicationID(communicationID int) Option {
	return optionFunc(func(o *options) { o.query["communication_id"] = communicationID })
}

// WithComplainID complain_id获取 投诉id
func (obj *OrderComplainCommunicationMgr) WithComplainID(complainID int) Option {
	return optionFunc(func(o *options) { o.query["complain_id"] = complainID })
}

// WithContent content获取 对话内容
func (obj *OrderComplainCommunicationMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithCreateTime create_time获取 对话时间
func (obj *OrderComplainCommunicationMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithOwner owner获取 所属，买家/卖家
func (obj *OrderComplainCommunicationMgr) WithOwner(owner string) Option {
	return optionFunc(func(o *options) { o.query["owner"] = owner })
}

// WithOwnerName owner_name获取 对话所属名称
func (obj *OrderComplainCommunicationMgr) WithOwnerName(ownerName string) Option {
	return optionFunc(func(o *options) { o.query["owner_name"] = ownerName })
}

// WithOwnerID owner_id获取 对话所属id,卖家id/买家id
func (obj *OrderComplainCommunicationMgr) WithOwnerID(ownerID int) Option {
	return optionFunc(func(o *options) { o.query["owner_id"] = ownerID })
}

// GetByOption 功能选项模式获取
func (obj *OrderComplainCommunicationMgr) GetByOption(opts ...Option) (result models.EsOrderComplainCommunication, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *OrderComplainCommunicationMgr) GetByOptions(opts ...Option) (results []*models.EsOrderComplainCommunication, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *OrderComplainCommunicationMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsOrderComplainCommunication, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromCommunicationID 通过communication_id获取内容 主键
func (obj *OrderComplainCommunicationMgr) GetFromCommunicationID(communicationID int) (result models.EsOrderComplainCommunication, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Where("`communication_id` = ?", communicationID).First(&result).Error

	return
}

// GetBatchFromCommunicationID 批量查找 主键
func (obj *OrderComplainCommunicationMgr) GetBatchFromCommunicationID(communicationIDs []int) (results []*models.EsOrderComplainCommunication, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Where("`communication_id` IN (?)", communicationIDs).Find(&results).Error

	return
}

// GetFromComplainID 通过complain_id获取内容 投诉id
func (obj *OrderComplainCommunicationMgr) GetFromComplainID(complainID int) (results []*models.EsOrderComplainCommunication, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Where("`complain_id` = ?", complainID).Find(&results).Error

	return
}

// GetBatchFromComplainID 批量查找 投诉id
func (obj *OrderComplainCommunicationMgr) GetBatchFromComplainID(complainIDs []int) (results []*models.EsOrderComplainCommunication, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Where("`complain_id` IN (?)", complainIDs).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 对话内容
func (obj *OrderComplainCommunicationMgr) GetFromContent(content string) (results []*models.EsOrderComplainCommunication, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 对话内容
func (obj *OrderComplainCommunicationMgr) GetBatchFromContent(contents []string) (results []*models.EsOrderComplainCommunication, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 对话时间
func (obj *OrderComplainCommunicationMgr) GetFromCreateTime(createTime int64) (results []*models.EsOrderComplainCommunication, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 对话时间
func (obj *OrderComplainCommunicationMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsOrderComplainCommunication, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromOwner 通过owner获取内容 所属，买家/卖家
func (obj *OrderComplainCommunicationMgr) GetFromOwner(owner string) (results []*models.EsOrderComplainCommunication, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Where("`owner` = ?", owner).Find(&results).Error

	return
}

// GetBatchFromOwner 批量查找 所属，买家/卖家
func (obj *OrderComplainCommunicationMgr) GetBatchFromOwner(owners []string) (results []*models.EsOrderComplainCommunication, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Where("`owner` IN (?)", owners).Find(&results).Error

	return
}

// GetFromOwnerName 通过owner_name获取内容 对话所属名称
func (obj *OrderComplainCommunicationMgr) GetFromOwnerName(ownerName string) (results []*models.EsOrderComplainCommunication, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Where("`owner_name` = ?", ownerName).Find(&results).Error

	return
}

// GetBatchFromOwnerName 批量查找 对话所属名称
func (obj *OrderComplainCommunicationMgr) GetBatchFromOwnerName(ownerNames []string) (results []*models.EsOrderComplainCommunication, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Where("`owner_name` IN (?)", ownerNames).Find(&results).Error

	return
}

// GetFromOwnerID 通过owner_id获取内容 对话所属id,卖家id/买家id
func (obj *OrderComplainCommunicationMgr) GetFromOwnerID(ownerID int) (results []*models.EsOrderComplainCommunication, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Where("`owner_id` = ?", ownerID).Find(&results).Error

	return
}

// GetBatchFromOwnerID 批量查找 对话所属id,卖家id/买家id
func (obj *OrderComplainCommunicationMgr) GetBatchFromOwnerID(ownerIDs []int) (results []*models.EsOrderComplainCommunication, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Where("`owner_id` IN (?)", ownerIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *OrderComplainCommunicationMgr) FetchByPrimaryKey(communicationID int) (result models.EsOrderComplainCommunication, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsOrderComplainCommunication{}).Where("`communication_id` = ?", communicationID).First(&result).Error

	return
}
