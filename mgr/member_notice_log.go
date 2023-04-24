package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type MemberNoticeLogMgr struct {
	*_BaseMgr
}

// NewMemberNoticeLogMgr open func
func NewMemberNoticeLogMgr(db db.Repo) *MemberNoticeLogMgr {
	if db == nil {
		panic(fmt.Errorf("NewMemberNoticeLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &MemberNoticeLogMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_member_notice_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *MemberNoticeLogMgr) GetTableName() string {
	return "es_member_notice_log"
}

// Reset 重置gorm会话
func (obj *MemberNoticeLogMgr) Reset() *MemberNoticeLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *MemberNoticeLogMgr) Get() (result models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *MemberNoticeLogMgr) Gets() (results []*models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *MemberNoticeLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 会员历史消息id
func (obj *MemberNoticeLogMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员id
func (obj *MemberNoticeLogMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithContent content获取 站内信内容
func (obj *MemberNoticeLogMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithSendTime send_time获取 发送时间
func (obj *MemberNoticeLogMgr) WithSendTime(sendTime int64) Option {
	return optionFunc(func(o *options) { o.query["send_time"] = sendTime })
}

// WithIsDel is_del获取 是否删除，0删除，1没有删除
func (obj *MemberNoticeLogMgr) WithIsDel(isDel int) Option {
	return optionFunc(func(o *options) { o.query["is_del"] = isDel })
}

// WithIsRead is_read获取 是否已读，0未读，1已读
func (obj *MemberNoticeLogMgr) WithIsRead(isRead int) Option {
	return optionFunc(func(o *options) { o.query["is_read"] = isRead })
}

// WithReceiveTime receive_time获取 接收时间
func (obj *MemberNoticeLogMgr) WithReceiveTime(receiveTime int64) Option {
	return optionFunc(func(o *options) { o.query["receive_time"] = receiveTime })
}

// WithTitle title获取 标题
func (obj *MemberNoticeLogMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// GetByOption 功能选项模式获取
func (obj *MemberNoticeLogMgr) GetByOption(opts ...Option) (result models.EsMemberNoticeLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *MemberNoticeLogMgr) GetByOptions(opts ...Option) (results []*models.EsMemberNoticeLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *MemberNoticeLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsMemberNoticeLog, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where(options.query)
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

// GetFromID 通过id获取内容 会员历史消息id
func (obj *MemberNoticeLogMgr) GetFromID(id int) (result models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 会员历史消息id
func (obj *MemberNoticeLogMgr) GetBatchFromID(ids []int) (results []*models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *MemberNoticeLogMgr) GetFromMemberID(memberID int) (results []*models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *MemberNoticeLogMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 站内信内容
func (obj *MemberNoticeLogMgr) GetFromContent(content string) (results []*models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 站内信内容
func (obj *MemberNoticeLogMgr) GetBatchFromContent(contents []string) (results []*models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromSendTime 通过send_time获取内容 发送时间
func (obj *MemberNoticeLogMgr) GetFromSendTime(sendTime int64) (results []*models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where("`send_time` = ?", sendTime).Find(&results).Error

	return
}

// GetBatchFromSendTime 批量查找 发送时间
func (obj *MemberNoticeLogMgr) GetBatchFromSendTime(sendTimes []int64) (results []*models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where("`send_time` IN (?)", sendTimes).Find(&results).Error

	return
}

// GetFromIsDel 通过is_del获取内容 是否删除，0删除，1没有删除
func (obj *MemberNoticeLogMgr) GetFromIsDel(isDel int) (results []*models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where("`is_del` = ?", isDel).Find(&results).Error

	return
}

// GetBatchFromIsDel 批量查找 是否删除，0删除，1没有删除
func (obj *MemberNoticeLogMgr) GetBatchFromIsDel(isDels []int) (results []*models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where("`is_del` IN (?)", isDels).Find(&results).Error

	return
}

// GetFromIsRead 通过is_read获取内容 是否已读，0未读，1已读
func (obj *MemberNoticeLogMgr) GetFromIsRead(isRead int) (results []*models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where("`is_read` = ?", isRead).Find(&results).Error

	return
}

// GetBatchFromIsRead 批量查找 是否已读，0未读，1已读
func (obj *MemberNoticeLogMgr) GetBatchFromIsRead(isReads []int) (results []*models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where("`is_read` IN (?)", isReads).Find(&results).Error

	return
}

// GetFromReceiveTime 通过receive_time获取内容 接收时间
func (obj *MemberNoticeLogMgr) GetFromReceiveTime(receiveTime int64) (results []*models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where("`receive_time` = ?", receiveTime).Find(&results).Error

	return
}

// GetBatchFromReceiveTime 批量查找 接收时间
func (obj *MemberNoticeLogMgr) GetBatchFromReceiveTime(receiveTimes []int64) (results []*models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where("`receive_time` IN (?)", receiveTimes).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 标题
func (obj *MemberNoticeLogMgr) GetFromTitle(title string) (results []*models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 标题
func (obj *MemberNoticeLogMgr) GetBatchFromTitle(titles []string) (results []*models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *MemberNoticeLogMgr) FetchByPrimaryKey(id int) (result models.EsMemberNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberNoticeLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
