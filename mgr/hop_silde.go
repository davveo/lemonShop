package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsShopSildeMgr struct {
	*_BaseMgr
}

// EsShopSildeMgr open func
func EsShopSildeMgr(db *gorm.DB) *_EsShopSildeMgr {
	if db == nil {
		panic(fmt.Errorf("EsShopSildeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsShopSildeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_shop_silde"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsShopSildeMgr) GetTableName() string {
	return "es_shop_silde"
}

// Reset 重置gorm会话
func (obj *_EsShopSildeMgr) Reset() *_EsShopSildeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsShopSildeMgr) Get() (result models.EsShopSilde, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopSilde{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsShopSildeMgr) Gets() (results []*models.EsShopSilde, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopSilde{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsShopSildeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsShopSilde{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSildeID silde_id获取 幻灯片Id
func (obj *_EsShopSildeMgr) WithSildeID(sildeID int) Option {
	return optionFunc(func(o *options) { o.query["silde_id"] = sildeID })
}

// WithShopID shop_id获取 店铺Id
func (obj *_EsShopSildeMgr) WithShopID(shopID int) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithSildeURL silde_url获取 幻灯片URL
func (obj *_EsShopSildeMgr) WithSildeURL(sildeURL string) Option {
	return optionFunc(func(o *options) { o.query["silde_url"] = sildeURL })
}

// WithImg img获取 图片
func (obj *_EsShopSildeMgr) WithImg(img string) Option {
	return optionFunc(func(o *options) { o.query["img"] = img })
}

// GetByOption 功能选项模式获取
func (obj *_EsShopSildeMgr) GetByOption(opts ...Option) (result models.EsShopSilde, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsShopSildeMgr) GetByOptions(opts ...Option) (results []*models.EsShopSilde, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsShopSildeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsShopSilde, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where(options.query)
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

// GetFromSildeID 通过silde_id获取内容 幻灯片Id
func (obj *_EsShopSildeMgr) GetFromSildeID(sildeID int) (result models.EsShopSilde, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`silde_id` = ?", sildeID).First(&result).Error

	return
}

// GetBatchFromSildeID 批量查找 幻灯片Id
func (obj *_EsShopSildeMgr) GetBatchFromSildeID(sildeIDs []int) (results []*models.EsShopSilde, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`silde_id` IN (?)", sildeIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺Id
func (obj *_EsShopSildeMgr) GetFromShopID(shopID int) (results []*models.EsShopSilde, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺Id
func (obj *_EsShopSildeMgr) GetBatchFromShopID(shopIDs []int) (results []*models.EsShopSilde, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromSildeURL 通过silde_url获取内容 幻灯片URL
func (obj *_EsShopSildeMgr) GetFromSildeURL(sildeURL string) (results []*models.EsShopSilde, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`silde_url` = ?", sildeURL).Find(&results).Error

	return
}

// GetBatchFromSildeURL 批量查找 幻灯片URL
func (obj *_EsShopSildeMgr) GetBatchFromSildeURL(sildeURLs []string) (results []*models.EsShopSilde, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`silde_url` IN (?)", sildeURLs).Find(&results).Error

	return
}

// GetFromImg 通过img获取内容 图片
func (obj *_EsShopSildeMgr) GetFromImg(img string) (results []*models.EsShopSilde, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`img` = ?", img).Find(&results).Error

	return
}

// GetBatchFromImg 批量查找 图片
func (obj *_EsShopSildeMgr) GetBatchFromImg(imgs []string) (results []*models.EsShopSilde, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`img` IN (?)", imgs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsShopSildeMgr) FetchByPrimaryKey(sildeID int) (result models.EsShopSilde, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`silde_id` = ?", sildeID).First(&result).Error

	return
}

// FetchIndexByIndShopSilde  获取多个内容
func (obj *_EsShopSildeMgr) FetchIndexByIndShopSilde(shopID int) (results []*models.EsShopSilde, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}
