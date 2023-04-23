package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsMessageTemplateMgr struct {
	*_BaseMgr
}

// EsMessageTemplateMgr open func
func EsMessageTemplateMgr(db *gorm.DB) *_EsMessageTemplateMgr {
	if db == nil {
		panic(fmt.Errorf("EsMessageTemplateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsMessageTemplateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_message_template"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsMessageTemplateMgr) GetTableName() string {
	return "es_message_template"
}

// Reset 重置gorm会话
func (obj *_EsMessageTemplateMgr) Reset() *_EsMessageTemplateMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsMessageTemplateMgr) Get() (result models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsMessageTemplateMgr) Gets() (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsMessageTemplateMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 模版ID
func (obj *_EsMessageTemplateMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTplCode tpl_code获取 模版编号
func (obj *_EsMessageTemplateMgr) WithTplCode(tplCode string) Option {
	return optionFunc(func(o *options) { o.query["tpl_code"] = tplCode })
}

// WithTplName tpl_name获取 模板名称
func (obj *_EsMessageTemplateMgr) WithTplName(tplName string) Option {
	return optionFunc(func(o *options) { o.query["tpl_name"] = tplName })
}

// WithType type获取 类型(会员 ,店铺 ,其他)
func (obj *_EsMessageTemplateMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithNoticeState notice_state获取 站内信提醒是否开启
func (obj *_EsMessageTemplateMgr) WithNoticeState(noticeState string) Option {
	return optionFunc(func(o *options) { o.query["notice_state"] = noticeState })
}

// WithSmsState sms_state获取 短信提醒是否开启
func (obj *_EsMessageTemplateMgr) WithSmsState(smsState string) Option {
	return optionFunc(func(o *options) { o.query["sms_state"] = smsState })
}

// WithEmailState email_state获取 邮件提醒是否开启
func (obj *_EsMessageTemplateMgr) WithEmailState(emailState string) Option {
	return optionFunc(func(o *options) { o.query["email_state"] = emailState })
}

// WithContent content获取 站内信内容
func (obj *_EsMessageTemplateMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithSmsContent sms_content获取 短信内容
func (obj *_EsMessageTemplateMgr) WithSmsContent(smsContent string) Option {
	return optionFunc(func(o *options) { o.query["sms_content"] = smsContent })
}

// WithEmailContent email_content获取 邮件内容
func (obj *_EsMessageTemplateMgr) WithEmailContent(emailContent string) Option {
	return optionFunc(func(o *options) { o.query["email_content"] = emailContent })
}

// WithEmailTitle email_title获取 邮件标题
func (obj *_EsMessageTemplateMgr) WithEmailTitle(emailTitle string) Option {
	return optionFunc(func(o *options) { o.query["email_title"] = emailTitle })
}

// GetByOption 功能选项模式获取
func (obj *_EsMessageTemplateMgr) GetByOption(opts ...Option) (result models.EsMessageTemplate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsMessageTemplateMgr) GetByOptions(opts ...Option) (results []*models.EsMessageTemplate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsMessageTemplateMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsMessageTemplate, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where(options.query)
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

// GetFromID 通过id获取内容 模版ID
func (obj *_EsMessageTemplateMgr) GetFromID(id int) (result models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 模版ID
func (obj *_EsMessageTemplateMgr) GetBatchFromID(ids []int) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTplCode 通过tpl_code获取内容 模版编号
func (obj *_EsMessageTemplateMgr) GetFromTplCode(tplCode string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`tpl_code` = ?", tplCode).Find(&results).Error

	return
}

// GetBatchFromTplCode 批量查找 模版编号
func (obj *_EsMessageTemplateMgr) GetBatchFromTplCode(tplCodes []string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`tpl_code` IN (?)", tplCodes).Find(&results).Error

	return
}

// GetFromTplName 通过tpl_name获取内容 模板名称
func (obj *_EsMessageTemplateMgr) GetFromTplName(tplName string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`tpl_name` = ?", tplName).Find(&results).Error

	return
}

// GetBatchFromTplName 批量查找 模板名称
func (obj *_EsMessageTemplateMgr) GetBatchFromTplName(tplNames []string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`tpl_name` IN (?)", tplNames).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型(会员 ,店铺 ,其他)
func (obj *_EsMessageTemplateMgr) GetFromType(_type string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 类型(会员 ,店铺 ,其他)
func (obj *_EsMessageTemplateMgr) GetBatchFromType(_types []string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromNoticeState 通过notice_state获取内容 站内信提醒是否开启
func (obj *_EsMessageTemplateMgr) GetFromNoticeState(noticeState string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`notice_state` = ?", noticeState).Find(&results).Error

	return
}

// GetBatchFromNoticeState 批量查找 站内信提醒是否开启
func (obj *_EsMessageTemplateMgr) GetBatchFromNoticeState(noticeStates []string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`notice_state` IN (?)", noticeStates).Find(&results).Error

	return
}

// GetFromSmsState 通过sms_state获取内容 短信提醒是否开启
func (obj *_EsMessageTemplateMgr) GetFromSmsState(smsState string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`sms_state` = ?", smsState).Find(&results).Error

	return
}

// GetBatchFromSmsState 批量查找 短信提醒是否开启
func (obj *_EsMessageTemplateMgr) GetBatchFromSmsState(smsStates []string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`sms_state` IN (?)", smsStates).Find(&results).Error

	return
}

// GetFromEmailState 通过email_state获取内容 邮件提醒是否开启
func (obj *_EsMessageTemplateMgr) GetFromEmailState(emailState string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`email_state` = ?", emailState).Find(&results).Error

	return
}

// GetBatchFromEmailState 批量查找 邮件提醒是否开启
func (obj *_EsMessageTemplateMgr) GetBatchFromEmailState(emailStates []string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`email_state` IN (?)", emailStates).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 站内信内容
func (obj *_EsMessageTemplateMgr) GetFromContent(content string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 站内信内容
func (obj *_EsMessageTemplateMgr) GetBatchFromContent(contents []string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromSmsContent 通过sms_content获取内容 短信内容
func (obj *_EsMessageTemplateMgr) GetFromSmsContent(smsContent string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`sms_content` = ?", smsContent).Find(&results).Error

	return
}

// GetBatchFromSmsContent 批量查找 短信内容
func (obj *_EsMessageTemplateMgr) GetBatchFromSmsContent(smsContents []string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`sms_content` IN (?)", smsContents).Find(&results).Error

	return
}

// GetFromEmailContent 通过email_content获取内容 邮件内容
func (obj *_EsMessageTemplateMgr) GetFromEmailContent(emailContent string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`email_content` = ?", emailContent).Find(&results).Error

	return
}

// GetBatchFromEmailContent 批量查找 邮件内容
func (obj *_EsMessageTemplateMgr) GetBatchFromEmailContent(emailContents []string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`email_content` IN (?)", emailContents).Find(&results).Error

	return
}

// GetFromEmailTitle 通过email_title获取内容 邮件标题
func (obj *_EsMessageTemplateMgr) GetFromEmailTitle(emailTitle string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`email_title` = ?", emailTitle).Find(&results).Error

	return
}

// GetBatchFromEmailTitle 批量查找 邮件标题
func (obj *_EsMessageTemplateMgr) GetBatchFromEmailTitle(emailTitles []string) (results []*models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`email_title` IN (?)", emailTitles).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsMessageTemplateMgr) FetchByPrimaryKey(id int) (result models.EsMessageTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMessageTemplate{}).Where("`id` = ?", id).First(&result).Error

	return
}
