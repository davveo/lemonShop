package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type ShopCatMgr struct {
	*_BaseMgr
}

// NewShopCatMgr open func
func NewShopCatMgr(db db.Repo) *ShopCatMgr {
	if db == nil {
		panic(fmt.Errorf("NewShopCatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &ShopCatMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_shop_cat"), wdb: db.GetDbW().Table("es_shop_cat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *ShopCatMgr) GetTableName() string {
	return "es_shop_cat"
}

// Reset 重置gorm会话
func (obj *ShopCatMgr) Reset() *ShopCatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *ShopCatMgr) Get() (result models.EsShopCat, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *ShopCatMgr) Gets() (results []*models.EsShopCat, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *ShopCatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithShopCatID shop_cat_id获取 店铺分组id
func (obj *ShopCatMgr) WithShopCatID(shopCatID int) Option {
	return optionFunc(func(o *options) { o.query["shop_cat_id"] = shopCatID })
}

// WithShopCatPid shop_cat_pid获取 店铺分组父ID
func (obj *ShopCatMgr) WithShopCatPid(shopCatPid int) Option {
	return optionFunc(func(o *options) { o.query["shop_cat_pid"] = shopCatPid })
}

// WithShopID shop_id获取 店铺id
func (obj *ShopCatMgr) WithShopID(shopID int) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithShopCatName shop_cat_name获取 店铺分组名称
func (obj *ShopCatMgr) WithShopCatName(shopCatName string) Option {
	return optionFunc(func(o *options) { o.query["shop_cat_name"] = shopCatName })
}

// WithDisable disable获取 店铺分组显示状态:1显示 0不显示
func (obj *ShopCatMgr) WithDisable(disable int) Option {
	return optionFunc(func(o *options) { o.query["disable"] = disable })
}

// WithSort sort获取 排序
func (obj *ShopCatMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithCatPath cat_path获取 分组路径
func (obj *ShopCatMgr) WithCatPath(catPath string) Option {
	return optionFunc(func(o *options) { o.query["cat_path"] = catPath })
}

// GetByOption 功能选项模式获取
func (obj *ShopCatMgr) GetByOption(opts ...Option) (result models.EsShopCat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *ShopCatMgr) GetByOptions(opts ...Option) (results []*models.EsShopCat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *ShopCatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsShopCat, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where(options.query)
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

// GetFromShopCatID 通过shop_cat_id获取内容 店铺分组id
func (obj *ShopCatMgr) GetFromShopCatID(shopCatID int) (result models.EsShopCat, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_cat_id` = ?", shopCatID).First(&result).Error

	return
}

// GetBatchFromShopCatID 批量查找 店铺分组id
func (obj *ShopCatMgr) GetBatchFromShopCatID(shopCatIDs []int) (results []*models.EsShopCat, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_cat_id` IN (?)", shopCatIDs).Find(&results).Error

	return
}

// GetFromShopCatPid 通过shop_cat_pid获取内容 店铺分组父ID
func (obj *ShopCatMgr) GetFromShopCatPid(shopCatPid int) (results []*models.EsShopCat, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_cat_pid` = ?", shopCatPid).Find(&results).Error

	return
}

// GetBatchFromShopCatPid 批量查找 店铺分组父ID
func (obj *ShopCatMgr) GetBatchFromShopCatPid(shopCatPids []int) (results []*models.EsShopCat, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_cat_pid` IN (?)", shopCatPids).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *ShopCatMgr) GetFromShopID(shopID int) (results []*models.EsShopCat, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *ShopCatMgr) GetBatchFromShopID(shopIDs []int) (results []*models.EsShopCat, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromShopCatName 通过shop_cat_name获取内容 店铺分组名称
func (obj *ShopCatMgr) GetFromShopCatName(shopCatName string) (results []*models.EsShopCat, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_cat_name` = ?", shopCatName).Find(&results).Error

	return
}

// GetBatchFromShopCatName 批量查找 店铺分组名称
func (obj *ShopCatMgr) GetBatchFromShopCatName(shopCatNames []string) (results []*models.EsShopCat, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_cat_name` IN (?)", shopCatNames).Find(&results).Error

	return
}

// GetFromDisable 通过disable获取内容 店铺分组显示状态:1显示 0不显示
func (obj *ShopCatMgr) GetFromDisable(disable int) (results []*models.EsShopCat, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`disable` = ?", disable).Find(&results).Error

	return
}

// GetBatchFromDisable 批量查找 店铺分组显示状态:1显示 0不显示
func (obj *ShopCatMgr) GetBatchFromDisable(disables []int) (results []*models.EsShopCat, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`disable` IN (?)", disables).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *ShopCatMgr) GetFromSort(sort int) (results []*models.EsShopCat, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *ShopCatMgr) GetBatchFromSort(sorts []int) (results []*models.EsShopCat, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromCatPath 通过cat_path获取内容 分组路径
func (obj *ShopCatMgr) GetFromCatPath(catPath string) (results []*models.EsShopCat, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`cat_path` = ?", catPath).Find(&results).Error

	return
}

// GetBatchFromCatPath 批量查找 分组路径
func (obj *ShopCatMgr) GetBatchFromCatPath(catPaths []string) (results []*models.EsShopCat, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`cat_path` IN (?)", catPaths).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *ShopCatMgr) FetchByPrimaryKey(shopCatID int) (result models.EsShopCat, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_cat_id` = ?", shopCatID).First(&result).Error

	return
}

// FetchIndexByEsShopCatShopid  获取多个内容
func (obj *ShopCatMgr) FetchIndexByEsShopCatShopid(shopID int) (results []*models.EsShopCat, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}
