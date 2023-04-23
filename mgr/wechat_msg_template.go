package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsWechatMsgTemplateMgr struct {
	*_BaseMgr
}

// EsWechatMsgTemplateMgr open func
func EsWechatMsgTemplateMgr(db *gorm.DB) *_EsWechatMsgTemplateMgr {
	if db == nil {
		panic(fmt.Errorf("EsWechatMsgTemplateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsWechatMsgTemplateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_wechat_msg_template"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsWechatMsgTemplateMgr) GetTableName() string {
	return "es_wechat_msg_template"
}

// Reset 重置gorm会话
func (obj *_EsWechatMsgTemplateMgr) Reset() *_EsWechatMsgTemplateMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsWechatMsgTemplateMgr) Get() (result models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsWechatMsgTemplateMgr) Gets() (results []*models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsWechatMsgTemplateMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_EsWechatMsgTemplateMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMsgTmpName msg_tmp_name获取 模板名称
func (obj *_EsWechatMsgTemplateMgr) WithMsgTmpName(msgTmpName string) Option {
	return optionFunc(func(o *options) { o.query["msg_tmp_name"] = msgTmpName })
}

// WithMsgTmpSn msg_tmp_sn获取 消息编号
func (obj *_EsWechatMsgTemplateMgr) WithMsgTmpSn(msgTmpSn string) Option {
	return optionFunc(func(o *options) { o.query["msg_tmp_sn"] = msgTmpSn })
}

// WithTemplateID template_id获取 模板ID
func (obj *_EsWechatMsgTemplateMgr) WithTemplateID(templateID string) Option {
	return optionFunc(func(o *options) { o.query["template_id"] = templateID })
}

// WithMsgFirst msg_first获取 消息开头文字
func (obj *_EsWechatMsgTemplateMgr) WithMsgFirst(msgFirst string) Option {
	return optionFunc(func(o *options) { o.query["msg_first"] = msgFirst })
}

// WithMsgRemark msg_remark获取 消息结尾备注文字
func (obj *_EsWechatMsgTemplateMgr) WithMsgRemark(msgRemark string) Option {
	return optionFunc(func(o *options) { o.query["msg_remark"] = msgRemark })
}

// WithIsOpen is_open获取 是否开启
func (obj *_EsWechatMsgTemplateMgr) WithIsOpen(isOpen int16) Option {
	return optionFunc(func(o *options) { o.query["is_open"] = isOpen })
}

// WithTmpType tmp_type获取 模板类型，枚举
func (obj *_EsWechatMsgTemplateMgr) WithTmpType(tmpType string) Option {
	return optionFunc(func(o *options) { o.query["tmp_type"] = tmpType })
}

// GetByOption 功能选项模式获取
func (obj *_EsWechatMsgTemplateMgr) GetByOption(opts ...Option) (result models.EsWechatMsgTemplate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsWechatMsgTemplateMgr) GetByOptions(opts ...Option) (results []*models.EsWechatMsgTemplate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsWechatMsgTemplateMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsWechatMsgTemplate, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键
func (obj *_EsWechatMsgTemplateMgr) GetFromID(id int) (result models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_EsWechatMsgTemplateMgr) GetBatchFromID(ids []int) (results []*models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMsgTmpName 通过msg_tmp_name获取内容 模板名称
func (obj *_EsWechatMsgTemplateMgr) GetFromMsgTmpName(msgTmpName string) (results []*models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where("`msg_tmp_name` = ?", msgTmpName).Find(&results).Error

	return
}

// GetBatchFromMsgTmpName 批量查找 模板名称
func (obj *_EsWechatMsgTemplateMgr) GetBatchFromMsgTmpName(msgTmpNames []string) (results []*models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where("`msg_tmp_name` IN (?)", msgTmpNames).Find(&results).Error

	return
}

// GetFromMsgTmpSn 通过msg_tmp_sn获取内容 消息编号
func (obj *_EsWechatMsgTemplateMgr) GetFromMsgTmpSn(msgTmpSn string) (results []*models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where("`msg_tmp_sn` = ?", msgTmpSn).Find(&results).Error

	return
}

// GetBatchFromMsgTmpSn 批量查找 消息编号
func (obj *_EsWechatMsgTemplateMgr) GetBatchFromMsgTmpSn(msgTmpSns []string) (results []*models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where("`msg_tmp_sn` IN (?)", msgTmpSns).Find(&results).Error

	return
}

// GetFromTemplateID 通过template_id获取内容 模板ID
func (obj *_EsWechatMsgTemplateMgr) GetFromTemplateID(templateID string) (results []*models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where("`template_id` = ?", templateID).Find(&results).Error

	return
}

// GetBatchFromTemplateID 批量查找 模板ID
func (obj *_EsWechatMsgTemplateMgr) GetBatchFromTemplateID(templateIDs []string) (results []*models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where("`template_id` IN (?)", templateIDs).Find(&results).Error

	return
}

// GetFromMsgFirst 通过msg_first获取内容 消息开头文字
func (obj *_EsWechatMsgTemplateMgr) GetFromMsgFirst(msgFirst string) (results []*models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where("`msg_first` = ?", msgFirst).Find(&results).Error

	return
}

// GetBatchFromMsgFirst 批量查找 消息开头文字
func (obj *_EsWechatMsgTemplateMgr) GetBatchFromMsgFirst(msgFirsts []string) (results []*models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where("`msg_first` IN (?)", msgFirsts).Find(&results).Error

	return
}

// GetFromMsgRemark 通过msg_remark获取内容 消息结尾备注文字
func (obj *_EsWechatMsgTemplateMgr) GetFromMsgRemark(msgRemark string) (results []*models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where("`msg_remark` = ?", msgRemark).Find(&results).Error

	return
}

// GetBatchFromMsgRemark 批量查找 消息结尾备注文字
func (obj *_EsWechatMsgTemplateMgr) GetBatchFromMsgRemark(msgRemarks []string) (results []*models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where("`msg_remark` IN (?)", msgRemarks).Find(&results).Error

	return
}

// GetFromIsOpen 通过is_open获取内容 是否开启
func (obj *_EsWechatMsgTemplateMgr) GetFromIsOpen(isOpen int16) (results []*models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where("`is_open` = ?", isOpen).Find(&results).Error

	return
}

// GetBatchFromIsOpen 批量查找 是否开启
func (obj *_EsWechatMsgTemplateMgr) GetBatchFromIsOpen(isOpens []int16) (results []*models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where("`is_open` IN (?)", isOpens).Find(&results).Error

	return
}

// GetFromTmpType 通过tmp_type获取内容 模板类型，枚举
func (obj *_EsWechatMsgTemplateMgr) GetFromTmpType(tmpType string) (results []*models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where("`tmp_type` = ?", tmpType).Find(&results).Error

	return
}

// GetBatchFromTmpType 批量查找 模板类型，枚举
func (obj *_EsWechatMsgTemplateMgr) GetBatchFromTmpType(tmpTypes []string) (results []*models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where("`tmp_type` IN (?)", tmpTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsWechatMsgTemplateMgr) FetchByPrimaryKey(id int) (result models.EsWechatMsgTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWechatMsgTemplate{}).Where("`id` = ?", id).First(&result).Error

	return
}
