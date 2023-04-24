package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type MessageMgr struct {
	*_BaseMgr
}

// NewMessageMgr open func
func NewMessageMgr(db db.Repo) *MessageMgr {
	if db == nil {
		panic(fmt.Errorf("NewMessageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &MessageMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_message"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *MessageMgr) GetTableName() string {
	return "es_message"
}

// Reset 重置gorm会话
func (obj *MessageMgr) Reset() *MessageMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *MessageMgr) Get() (result models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *MessageMgr) Gets() (results []*models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *MessageMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 站内消息主键
func (obj *MessageMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitle title获取 标题
func (obj *MessageMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithContent content获取 消息内容
func (obj *MessageMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithMemberIDs member_ids获取 会员id
func (obj *MessageMgr) WithMemberIDs(memberIDs string) Option {
	return optionFunc(func(o *options) { o.query["member_ids"] = memberIDs })
}

// WithAdminID admin_id获取 管理员id
func (obj *MessageMgr) WithAdminID(adminID int) Option {
	return optionFunc(func(o *options) { o.query["admin_id"] = adminID })
}

// WithAdminName admin_name获取 管理员名称
func (obj *MessageMgr) WithAdminName(adminName string) Option {
	return optionFunc(func(o *options) { o.query["admin_name"] = adminName })
}

// WithSendTime send_time获取 发送时间
func (obj *MessageMgr) WithSendTime(sendTime int64) Option {
	return optionFunc(func(o *options) { o.query["send_time"] = sendTime })
}

// WithSendType send_type获取 发送类型
func (obj *MessageMgr) WithSendType(sendType int) Option {
	return optionFunc(func(o *options) { o.query["send_type"] = sendType })
}

// WithDisabled disabled获取 是否删除 0：否，1：是
func (obj *MessageMgr) WithDisabled(disabled int) Option {
	return optionFunc(func(o *options) { o.query["disabled"] = disabled })
}

// GetByOption 功能选项模式获取
func (obj *MessageMgr) GetByOption(opts ...Option) (result models.EsMessage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *MessageMgr) GetByOptions(opts ...Option) (results []*models.EsMessage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *MessageMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsMessage, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where(options.query)
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

// GetFromID 通过id获取内容 站内消息主键
func (obj *MessageMgr) GetFromID(id int) (result models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 站内消息主键
func (obj *MessageMgr) GetBatchFromID(ids []int) (results []*models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标题
func (obj *MessageMgr) GetFromTitle(title string) (results []*models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 标题
func (obj *MessageMgr) GetBatchFromTitle(titles []string) (results []*models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 消息内容
func (obj *MessageMgr) GetFromContent(content string) (results []*models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 消息内容
func (obj *MessageMgr) GetBatchFromContent(contents []string) (results []*models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromMemberIDs 通过member_ids获取内容 会员id
func (obj *MessageMgr) GetFromMemberIDs(memberIDs string) (results []*models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`member_ids` = ?", memberIDs).Find(&results).Error

	return
}

// GetBatchFromMemberIDs 批量查找 会员id
func (obj *MessageMgr) GetBatchFromMemberIDs(memberIDss []string) (results []*models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`member_ids` IN (?)", memberIDss).Find(&results).Error

	return
}

// GetFromAdminID 通过admin_id获取内容 管理员id
func (obj *MessageMgr) GetFromAdminID(adminID int) (results []*models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`admin_id` = ?", adminID).Find(&results).Error

	return
}

// GetBatchFromAdminID 批量查找 管理员id
func (obj *MessageMgr) GetBatchFromAdminID(adminIDs []int) (results []*models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`admin_id` IN (?)", adminIDs).Find(&results).Error

	return
}

// GetFromAdminName 通过admin_name获取内容 管理员名称
func (obj *MessageMgr) GetFromAdminName(adminName string) (results []*models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`admin_name` = ?", adminName).Find(&results).Error

	return
}

// GetBatchFromAdminName 批量查找 管理员名称
func (obj *MessageMgr) GetBatchFromAdminName(adminNames []string) (results []*models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`admin_name` IN (?)", adminNames).Find(&results).Error

	return
}

// GetFromSendTime 通过send_time获取内容 发送时间
func (obj *MessageMgr) GetFromSendTime(sendTime int64) (results []*models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`send_time` = ?", sendTime).Find(&results).Error

	return
}

// GetBatchFromSendTime 批量查找 发送时间
func (obj *MessageMgr) GetBatchFromSendTime(sendTimes []int64) (results []*models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`send_time` IN (?)", sendTimes).Find(&results).Error

	return
}

// GetFromSendType 通过send_type获取内容 发送类型
func (obj *MessageMgr) GetFromSendType(sendType int) (results []*models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`send_type` = ?", sendType).Find(&results).Error

	return
}

// GetBatchFromSendType 批量查找 发送类型
func (obj *MessageMgr) GetBatchFromSendType(sendTypes []int) (results []*models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`send_type` IN (?)", sendTypes).Find(&results).Error

	return
}

// GetFromDisabled 通过disabled获取内容 是否删除 0：否，1：是
func (obj *MessageMgr) GetFromDisabled(disabled int) (results []*models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`disabled` = ?", disabled).Find(&results).Error

	return
}

// GetBatchFromDisabled 批量查找 是否删除 0：否，1：是
func (obj *MessageMgr) GetBatchFromDisabled(disableds []int) (results []*models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`disabled` IN (?)", disableds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *MessageMgr) FetchByPrimaryKey(id int) (result models.EsMessage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMessage{}).Where("`id` = ?", id).First(&result).Error

	return
}
