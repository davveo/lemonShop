package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsReceiptContentMgr struct {
	*_BaseMgr
}

// EsReceiptContentMgr open func
func EsReceiptContentMgr(db *gorm.DB) *_EsReceiptContentMgr {
	if db == nil {
		panic(fmt.Errorf("EsReceiptContentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsReceiptContentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_receipt_content"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsReceiptContentMgr) GetTableName() string {
	return "es_receipt_content"
}

// Reset 重置gorm会话
func (obj *_EsReceiptContentMgr) Reset() *_EsReceiptContentMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsReceiptContentMgr) Get() (result models.EsReceiptContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptContent{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsReceiptContentMgr) Gets() (results []*models.EsReceiptContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsReceiptContentMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 发票内容id
func (obj *_EsReceiptContentMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithContent content获取 发票内容
func (obj *_EsReceiptContentMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// GetByOption 功能选项模式获取
func (obj *_EsReceiptContentMgr) GetByOption(opts ...Option) (result models.EsReceiptContent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsReceiptContentMgr) GetByOptions(opts ...Option) (results []*models.EsReceiptContent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsReceiptContentMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsReceiptContent, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Where(options.query)
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

// GetFromID 通过id获取内容 发票内容id
func (obj *_EsReceiptContentMgr) GetFromID(id int) (result models.EsReceiptContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 发票内容id
func (obj *_EsReceiptContentMgr) GetBatchFromID(ids []int) (results []*models.EsReceiptContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 发票内容
func (obj *_EsReceiptContentMgr) GetFromContent(content string) (results []*models.EsReceiptContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 发票内容
func (obj *_EsReceiptContentMgr) GetBatchFromContent(contents []string) (results []*models.EsReceiptContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsReceiptContentMgr) FetchByPrimaryKey(id int) (result models.EsReceiptContent, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Where("`id` = ?", id).First(&result).Error

	return
}
