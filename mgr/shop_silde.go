package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type ShopSildeMgr struct {
	*_BaseMgr
}

// NewShopSildeMgr open func
func NewShopSildeMgr(db db.Repo) *ShopSildeMgr {
	if db == nil {
		panic(fmt.Errorf("NewShopSildeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &ShopSildeMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_shop_silde"), wdb: db.GetDbW().Table("es_shop_silde"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *ShopSildeMgr) GetTableName() string {
	return "es_shop_silde"
}

// Reset 重置gorm会话
func (obj *ShopSildeMgr) Reset() *ShopSildeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *ShopSildeMgr) Get() (result models.EsShopSilde, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopSilde{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *ShopSildeMgr) Gets() (results []*models.EsShopSilde, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopSilde{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *ShopSildeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsShopSilde{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSildeID silde_id获取 幻灯片Id
func (obj *ShopSildeMgr) WithSildeID(sildeID int) Option {
	return optionFunc(func(o *options) { o.query["silde_id"] = sildeID })
}

// WithShopID shop_id获取 店铺Id
func (obj *ShopSildeMgr) WithShopID(shopID int) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithSildeURL silde_url获取 幻灯片URL
func (obj *ShopSildeMgr) WithSildeURL(sildeURL string) Option {
	return optionFunc(func(o *options) { o.query["silde_url"] = sildeURL })
}

// WithImg img获取 图片
func (obj *ShopSildeMgr) WithImg(img string) Option {
	return optionFunc(func(o *options) { o.query["img"] = img })
}

// GetByOption 功能选项模式获取
func (obj *ShopSildeMgr) GetByOption(opts ...Option) (result models.EsShopSilde, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *ShopSildeMgr) GetByOptions(opts ...Option) (results []*models.EsShopSilde, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *ShopSildeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsShopSilde, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where(options.query)
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
func (obj *ShopSildeMgr) GetFromSildeID(sildeID int) (result models.EsShopSilde, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`silde_id` = ?", sildeID).First(&result).Error

	return
}

// GetBatchFromSildeID 批量查找 幻灯片Id
func (obj *ShopSildeMgr) GetBatchFromSildeID(sildeIDs []int) (results []*models.EsShopSilde, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`silde_id` IN (?)", sildeIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺Id
func (obj *ShopSildeMgr) GetFromShopID(shopID int) (results []*models.EsShopSilde, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺Id
func (obj *ShopSildeMgr) GetBatchFromShopID(shopIDs []int) (results []*models.EsShopSilde, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromSildeURL 通过silde_url获取内容 幻灯片URL
func (obj *ShopSildeMgr) GetFromSildeURL(sildeURL string) (results []*models.EsShopSilde, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`silde_url` = ?", sildeURL).Find(&results).Error

	return
}

// GetBatchFromSildeURL 批量查找 幻灯片URL
func (obj *ShopSildeMgr) GetBatchFromSildeURL(sildeURLs []string) (results []*models.EsShopSilde, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`silde_url` IN (?)", sildeURLs).Find(&results).Error

	return
}

// GetFromImg 通过img获取内容 图片
func (obj *ShopSildeMgr) GetFromImg(img string) (results []*models.EsShopSilde, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`img` = ?", img).Find(&results).Error

	return
}

// GetBatchFromImg 批量查找 图片
func (obj *ShopSildeMgr) GetBatchFromImg(imgs []string) (results []*models.EsShopSilde, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`img` IN (?)", imgs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *ShopSildeMgr) FetchByPrimaryKey(sildeID int) (result models.EsShopSilde, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`silde_id` = ?", sildeID).First(&result).Error

	return
}

// FetchIndexByIndShopSilde  获取多个内容
func (obj *ShopSildeMgr) FetchIndexByIndShopSilde(shopID int) (results []*models.EsShopSilde, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopSilde{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}
