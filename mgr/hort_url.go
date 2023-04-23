package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsShortURLMgr struct {
	*_BaseMgr
}

// EsShortURLMgr open func
func EsShortURLMgr(db *gorm.DB) *_EsShortURLMgr {
	if db == nil {
		panic(fmt.Errorf("EsShortURLMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsShortURLMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_short_url"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsShortURLMgr) GetTableName() string {
	return "es_short_url"
}

// Reset 重置gorm会话
func (obj *_EsShortURLMgr) Reset() *_EsShortURLMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsShortURLMgr) Get() (result models.EsShortURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShortURL{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsShortURLMgr) Gets() (results []*models.EsShortURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShortURL{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsShortURLMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsShortURL{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_EsShortURLMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithURL url获取 跳转地址
func (obj *_EsShortURLMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithSu su获取 短链接key
func (obj *_EsShortURLMgr) WithSu(su string) Option {
	return optionFunc(func(o *options) { o.query["su"] = su })
}

// GetByOption 功能选项模式获取
func (obj *_EsShortURLMgr) GetByOption(opts ...Option) (result models.EsShortURL, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsShortURL{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsShortURLMgr) GetByOptions(opts ...Option) (results []*models.EsShortURL, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsShortURL{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsShortURLMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsShortURL, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsShortURL{}).Where(options.query)
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

// GetFromID 通过id获取内容 id
func (obj *_EsShortURLMgr) GetFromID(id int64) (result models.EsShortURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShortURL{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_EsShortURLMgr) GetBatchFromID(ids []int64) (results []*models.EsShortURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShortURL{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容 跳转地址
func (obj *_EsShortURLMgr) GetFromURL(url string) (results []*models.EsShortURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShortURL{}).Where("`url` = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量查找 跳转地址
func (obj *_EsShortURLMgr) GetBatchFromURL(urls []string) (results []*models.EsShortURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShortURL{}).Where("`url` IN (?)", urls).Find(&results).Error

	return
}

// GetFromSu 通过su获取内容 短链接key
func (obj *_EsShortURLMgr) GetFromSu(su string) (results []*models.EsShortURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShortURL{}).Where("`su` = ?", su).Find(&results).Error

	return
}

// GetBatchFromSu 批量查找 短链接key
func (obj *_EsShortURLMgr) GetBatchFromSu(sus []string) (results []*models.EsShortURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShortURL{}).Where("`su` IN (?)", sus).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsShortURLMgr) FetchByPrimaryKey(id int64) (result models.EsShortURL, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShortURL{}).Where("`id` = ?", id).First(&result).Error

	return
}
