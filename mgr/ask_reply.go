package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type AskReplyMgr struct {
	*_BaseMgr
}

// NewAskReplyMgr open func
func NewAskReplyMgr(db db.Repo) *AskReplyMgr {
	if db == nil {
		panic(fmt.Errorf("NewAskReplyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &AskReplyMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_ask_reply"),
		wdb:       db.GetDbW().Table("es_ask_reply"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *AskReplyMgr) GetTableName() string {
	return "es_ask_reply"
}

// Reset 重置gorm会话
func (obj *AskReplyMgr) Reset() *AskReplyMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *AskReplyMgr) Get() (result models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *AskReplyMgr) Gets() (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *AskReplyMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *AskReplyMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAskID ask_id获取 会员咨询id
func (obj *AskReplyMgr) WithAskID(askID int) Option {
	return optionFunc(func(o *options) { o.query["ask_id"] = askID })
}

// WithMemberID member_id获取 会员id
func (obj *AskReplyMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMemberName member_name获取 会员名称
func (obj *AskReplyMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithContent content获取 回复内容
func (obj *AskReplyMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithReplyTime reply_time获取 回复时间
func (obj *AskReplyMgr) WithReplyTime(replyTime int64) Option {
	return optionFunc(func(o *options) { o.query["reply_time"] = replyTime })
}

// WithAnonymous anonymous获取 是否匿名 YES:是，NO:否
func (obj *AskReplyMgr) WithAnonymous(anonymous string) Option {
	return optionFunc(func(o *options) { o.query["anonymous"] = anonymous })
}

// WithAuthStatus auth_status获取 审核状态 WAIT_AUDIT:待审核,PASS_AUDIT:审核通过,REFUSE_AUDIT:审核未通过
func (obj *AskReplyMgr) WithAuthStatus(authStatus string) Option {
	return optionFunc(func(o *options) { o.query["auth_status"] = authStatus })
}

// WithIsDel is_del获取 删除状态 DELETED：已删除 NORMAL：正常
func (obj *AskReplyMgr) WithIsDel(isDel string) Option {
	return optionFunc(func(o *options) { o.query["is_del"] = isDel })
}

// WithReplyStatus reply_status获取 是否已回复 YES:是，NO:否
func (obj *AskReplyMgr) WithReplyStatus(replyStatus string) Option {
	return optionFunc(func(o *options) { o.query["reply_status"] = replyStatus })
}

// WithCreateTime create_time获取 创建时间
func (obj *AskReplyMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *AskReplyMgr) GetByOption(opts ...Option) (result models.EsAskReply, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *AskReplyMgr) GetByOptions(opts ...Option) (results []*models.EsAskReply, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *AskReplyMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsAskReply, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键ID
func (obj *AskReplyMgr) GetFromID(id int) (result models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *AskReplyMgr) GetBatchFromID(ids []int) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromAskID 通过ask_id获取内容 会员咨询id
func (obj *AskReplyMgr) GetFromAskID(askID int) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`ask_id` = ?", askID).Find(&results).Error

	return
}

// GetBatchFromAskID 批量查找 会员咨询id
func (obj *AskReplyMgr) GetBatchFromAskID(askIDs []int) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`ask_id` IN (?)", askIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *AskReplyMgr) GetFromMemberID(memberID int) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *AskReplyMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 会员名称
func (obj *AskReplyMgr) GetFromMemberName(memberName string) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 会员名称
func (obj *AskReplyMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 回复内容
func (obj *AskReplyMgr) GetFromContent(content string) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 回复内容
func (obj *AskReplyMgr) GetBatchFromContent(contents []string) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromReplyTime 通过reply_time获取内容 回复时间
func (obj *AskReplyMgr) GetFromReplyTime(replyTime int64) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`reply_time` = ?", replyTime).Find(&results).Error

	return
}

// GetBatchFromReplyTime 批量查找 回复时间
func (obj *AskReplyMgr) GetBatchFromReplyTime(replyTimes []int64) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`reply_time` IN (?)", replyTimes).Find(&results).Error

	return
}

// GetFromAnonymous 通过anonymous获取内容 是否匿名 YES:是，NO:否
func (obj *AskReplyMgr) GetFromAnonymous(anonymous string) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`anonymous` = ?", anonymous).Find(&results).Error

	return
}

// GetBatchFromAnonymous 批量查找 是否匿名 YES:是，NO:否
func (obj *AskReplyMgr) GetBatchFromAnonymous(anonymouss []string) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`anonymous` IN (?)", anonymouss).Find(&results).Error

	return
}

// GetFromAuthStatus 通过auth_status获取内容 审核状态 WAIT_AUDIT:待审核,PASS_AUDIT:审核通过,REFUSE_AUDIT:审核未通过
func (obj *AskReplyMgr) GetFromAuthStatus(authStatus string) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`auth_status` = ?", authStatus).Find(&results).Error

	return
}

// GetBatchFromAuthStatus 批量查找 审核状态 WAIT_AUDIT:待审核,PASS_AUDIT:审核通过,REFUSE_AUDIT:审核未通过
func (obj *AskReplyMgr) GetBatchFromAuthStatus(authStatuss []string) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`auth_status` IN (?)", authStatuss).Find(&results).Error

	return
}

// GetFromIsDel 通过is_del获取内容 删除状态 DELETED：已删除 NORMAL：正常
func (obj *AskReplyMgr) GetFromIsDel(isDel string) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`is_del` = ?", isDel).Find(&results).Error

	return
}

// GetBatchFromIsDel 批量查找 删除状态 DELETED：已删除 NORMAL：正常
func (obj *AskReplyMgr) GetBatchFromIsDel(isDels []string) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`is_del` IN (?)", isDels).Find(&results).Error

	return
}

// GetFromReplyStatus 通过reply_status获取内容 是否已回复 YES:是，NO:否
func (obj *AskReplyMgr) GetFromReplyStatus(replyStatus string) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`reply_status` = ?", replyStatus).Find(&results).Error

	return
}

// GetBatchFromReplyStatus 批量查找 是否已回复 YES:是，NO:否
func (obj *AskReplyMgr) GetBatchFromReplyStatus(replyStatuss []string) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`reply_status` IN (?)", replyStatuss).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *AskReplyMgr) GetFromCreateTime(createTime int64) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *AskReplyMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *AskReplyMgr) FetchByPrimaryKey(id int) (result models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByEsIndexAskID  获取多个内容
func (obj *AskReplyMgr) FetchIndexByEsIndexAskID(askID int) (results []*models.EsAskReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAskReply{}).Where("`ask_id` = ?", askID).Find(&results).Error

	return
}
