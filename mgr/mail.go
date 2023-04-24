package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type mailMgr struct {
	*_BaseMgr
}

// EsEmailMgr open func
func EsEmailMgr(db db.Repo) *mailMgr {
	if db == nil {
		panic(fmt.Errorf("EsEmailMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &mailMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_email"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *mailMgr) GetTableName() string {
	return "es_email"
}

// Reset 重置gorm会话
func (obj *mailMgr) Reset() *mailMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *mailMgr) Get() (result models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *mailMgr) Gets() (results []*models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *mailMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 邮件记录id
func (obj *mailMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTitle title获取 邮件标题
func (obj *mailMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithType type获取 邮件类型
func (obj *mailMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithSuccess success获取 是否成功
func (obj *mailMgr) WithSuccess(success int) Option {
	return optionFunc(func(o *options) { o.query["success"] = success })
}

// WithEmail email获取 邮件接收者
func (obj *mailMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithContent content获取 邮件内容
func (obj *mailMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithErrorNum error_num获取 错误次数
func (obj *mailMgr) WithErrorNum(errorNum int) Option {
	return optionFunc(func(o *options) { o.query["error_num"] = errorNum })
}

// WithLastSend last_send获取 最后发送时间
func (obj *mailMgr) WithLastSend(lastSend int64) Option {
	return optionFunc(func(o *options) { o.query["last_send"] = lastSend })
}

// GetByOption 功能选项模式获取
func (obj *mailMgr) GetByOption(opts ...Option) (result models.EsEmail, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *mailMgr) GetByOptions(opts ...Option) (results []*models.EsEmail, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *mailMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsEmail, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where(options.query)
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

// GetFromID 通过id获取内容 邮件记录id
func (obj *mailMgr) GetFromID(id int) (result models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 邮件记录id
func (obj *mailMgr) GetBatchFromID(ids []int) (results []*models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 邮件标题
func (obj *mailMgr) GetFromTitle(title string) (results []*models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 邮件标题
func (obj *mailMgr) GetBatchFromTitle(titles []string) (results []*models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 邮件类型
func (obj *mailMgr) GetFromType(_type string) (results []*models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 邮件类型
func (obj *mailMgr) GetBatchFromType(_types []string) (results []*models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromSuccess 通过success获取内容 是否成功
func (obj *mailMgr) GetFromSuccess(success int) (results []*models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where("`success` = ?", success).Find(&results).Error

	return
}

// GetBatchFromSuccess 批量查找 是否成功
func (obj *mailMgr) GetBatchFromSuccess(successs []int) (results []*models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where("`success` IN (?)", successs).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 邮件接收者
func (obj *mailMgr) GetFromEmail(email string) (results []*models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where("`email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找 邮件接收者
func (obj *mailMgr) GetBatchFromEmail(emails []string) (results []*models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where("`email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 邮件内容
func (obj *mailMgr) GetFromContent(content string) (results []*models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 邮件内容
func (obj *mailMgr) GetBatchFromContent(contents []string) (results []*models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromErrorNum 通过error_num获取内容 错误次数
func (obj *mailMgr) GetFromErrorNum(errorNum int) (results []*models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where("`error_num` = ?", errorNum).Find(&results).Error

	return
}

// GetBatchFromErrorNum 批量查找 错误次数
func (obj *mailMgr) GetBatchFromErrorNum(errorNums []int) (results []*models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where("`error_num` IN (?)", errorNums).Find(&results).Error

	return
}

// GetFromLastSend 通过last_send获取内容 最后发送时间
func (obj *mailMgr) GetFromLastSend(lastSend int64) (results []*models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where("`last_send` = ?", lastSend).Find(&results).Error

	return
}

// GetBatchFromLastSend 批量查找 最后发送时间
func (obj *mailMgr) GetBatchFromLastSend(lastSends []int64) (results []*models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where("`last_send` IN (?)", lastSends).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *mailMgr) FetchByPrimaryKey(id int) (result models.EsEmail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsEmail{}).Where("`id` = ?", id).First(&result).Error

	return
}
