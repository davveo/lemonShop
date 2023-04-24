package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type ReceiptContentMgr struct {
	*_BaseMgr
}

// NewReceiptContentMgr open func
func NewReceiptContentMgr(db db.Repo) *ReceiptContentMgr {
	if db == nil {
		panic(fmt.Errorf("NewReceiptContentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &ReceiptContentMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_receipt_content"), wdb: db.GetDbW().Table("es_receipt_content"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *ReceiptContentMgr) GetTableName() string {
	return "es_receipt_content"
}

// Reset 重置gorm会话
func (obj *ReceiptContentMgr) Reset() *ReceiptContentMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *ReceiptContentMgr) Get() (result models.EsReceiptContent, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptContent{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *ReceiptContentMgr) Gets() (results []*models.EsReceiptContent, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *ReceiptContentMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 发票内容id
func (obj *ReceiptContentMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithContent content获取 发票内容
func (obj *ReceiptContentMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// GetByOption 功能选项模式获取
func (obj *ReceiptContentMgr) GetByOption(opts ...Option) (result models.EsReceiptContent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *ReceiptContentMgr) GetByOptions(opts ...Option) (results []*models.EsReceiptContent, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *ReceiptContentMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsReceiptContent, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Where(options.query)
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
func (obj *ReceiptContentMgr) GetFromID(id int) (result models.EsReceiptContent, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 发票内容id
func (obj *ReceiptContentMgr) GetBatchFromID(ids []int) (results []*models.EsReceiptContent, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 发票内容
func (obj *ReceiptContentMgr) GetFromContent(content string) (results []*models.EsReceiptContent, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 发票内容
func (obj *ReceiptContentMgr) GetBatchFromContent(contents []string) (results []*models.EsReceiptContent, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *ReceiptContentMgr) FetchByPrimaryKey(id int) (result models.EsReceiptContent, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptContent{}).Where("`id` = ?", id).First(&result).Error

	return
}
