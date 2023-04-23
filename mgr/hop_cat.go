package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsShopCatMgr struct {
	*_BaseMgr
}

// EsShopCatMgr open func
func EsShopCatMgr(db *gorm.DB) *_EsShopCatMgr {
	if db == nil {
		panic(fmt.Errorf("EsShopCatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsShopCatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_shop_cat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsShopCatMgr) GetTableName() string {
	return "es_shop_cat"
}

// Reset 重置gorm会话
func (obj *_EsShopCatMgr) Reset() *_EsShopCatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsShopCatMgr) Get() (result models.EsShopCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsShopCatMgr) Gets() (results []*models.EsShopCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsShopCatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithShopCatID shop_cat_id获取 店铺分组id
func (obj *_EsShopCatMgr) WithShopCatID(shopCatID int) Option {
	return optionFunc(func(o *options) { o.query["shop_cat_id"] = shopCatID })
}

// WithShopCatPid shop_cat_pid获取 店铺分组父ID
func (obj *_EsShopCatMgr) WithShopCatPid(shopCatPid int) Option {
	return optionFunc(func(o *options) { o.query["shop_cat_pid"] = shopCatPid })
}

// WithShopID shop_id获取 店铺id
func (obj *_EsShopCatMgr) WithShopID(shopID int) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithShopCatName shop_cat_name获取 店铺分组名称
func (obj *_EsShopCatMgr) WithShopCatName(shopCatName string) Option {
	return optionFunc(func(o *options) { o.query["shop_cat_name"] = shopCatName })
}

// WithDisable disable获取 店铺分组显示状态:1显示 0不显示
func (obj *_EsShopCatMgr) WithDisable(disable int) Option {
	return optionFunc(func(o *options) { o.query["disable"] = disable })
}

// WithSort sort获取 排序
func (obj *_EsShopCatMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithCatPath cat_path获取 分组路径
func (obj *_EsShopCatMgr) WithCatPath(catPath string) Option {
	return optionFunc(func(o *options) { o.query["cat_path"] = catPath })
}

// GetByOption 功能选项模式获取
func (obj *_EsShopCatMgr) GetByOption(opts ...Option) (result models.EsShopCat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsShopCatMgr) GetByOptions(opts ...Option) (results []*models.EsShopCat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsShopCatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsShopCat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where(options.query)
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
func (obj *_EsShopCatMgr) GetFromShopCatID(shopCatID int) (result models.EsShopCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_cat_id` = ?", shopCatID).First(&result).Error

	return
}

// GetBatchFromShopCatID 批量查找 店铺分组id
func (obj *_EsShopCatMgr) GetBatchFromShopCatID(shopCatIDs []int) (results []*models.EsShopCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_cat_id` IN (?)", shopCatIDs).Find(&results).Error

	return
}

// GetFromShopCatPid 通过shop_cat_pid获取内容 店铺分组父ID
func (obj *_EsShopCatMgr) GetFromShopCatPid(shopCatPid int) (results []*models.EsShopCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_cat_pid` = ?", shopCatPid).Find(&results).Error

	return
}

// GetBatchFromShopCatPid 批量查找 店铺分组父ID
func (obj *_EsShopCatMgr) GetBatchFromShopCatPid(shopCatPids []int) (results []*models.EsShopCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_cat_pid` IN (?)", shopCatPids).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_EsShopCatMgr) GetFromShopID(shopID int) (results []*models.EsShopCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_EsShopCatMgr) GetBatchFromShopID(shopIDs []int) (results []*models.EsShopCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromShopCatName 通过shop_cat_name获取内容 店铺分组名称
func (obj *_EsShopCatMgr) GetFromShopCatName(shopCatName string) (results []*models.EsShopCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_cat_name` = ?", shopCatName).Find(&results).Error

	return
}

// GetBatchFromShopCatName 批量查找 店铺分组名称
func (obj *_EsShopCatMgr) GetBatchFromShopCatName(shopCatNames []string) (results []*models.EsShopCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_cat_name` IN (?)", shopCatNames).Find(&results).Error

	return
}

// GetFromDisable 通过disable获取内容 店铺分组显示状态:1显示 0不显示
func (obj *_EsShopCatMgr) GetFromDisable(disable int) (results []*models.EsShopCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`disable` = ?", disable).Find(&results).Error

	return
}

// GetBatchFromDisable 批量查找 店铺分组显示状态:1显示 0不显示
func (obj *_EsShopCatMgr) GetBatchFromDisable(disables []int) (results []*models.EsShopCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`disable` IN (?)", disables).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_EsShopCatMgr) GetFromSort(sort int) (results []*models.EsShopCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_EsShopCatMgr) GetBatchFromSort(sorts []int) (results []*models.EsShopCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromCatPath 通过cat_path获取内容 分组路径
func (obj *_EsShopCatMgr) GetFromCatPath(catPath string) (results []*models.EsShopCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`cat_path` = ?", catPath).Find(&results).Error

	return
}

// GetBatchFromCatPath 批量查找 分组路径
func (obj *_EsShopCatMgr) GetBatchFromCatPath(catPaths []string) (results []*models.EsShopCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`cat_path` IN (?)", catPaths).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsShopCatMgr) FetchByPrimaryKey(shopCatID int) (result models.EsShopCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_cat_id` = ?", shopCatID).First(&result).Error

	return
}

// FetchIndexByEsShopCatShopid  获取多个内容
func (obj *_EsShopCatMgr) FetchIndexByEsShopCatShopid(shopID int) (results []*models.EsShopCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopCat{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}
