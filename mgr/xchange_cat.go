package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsExchangeCatMgr struct {
	*_BaseMgr
}

// EsExchangeCatMgr open func
func EsExchangeCatMgr(db *gorm.DB) *_EsExchangeCatMgr {
	if db == nil {
		panic(fmt.Errorf("EsExchangeCatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsExchangeCatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_exchange_cat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsExchangeCatMgr) GetTableName() string {
	return "es_exchange_cat"
}

// Reset 重置gorm会话
func (obj *_EsExchangeCatMgr) Reset() *_EsExchangeCatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsExchangeCatMgr) Get() (result models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsExchangeCatMgr) Gets() (results []*models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsExchangeCatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCategoryID category_id获取 分类id
func (obj *_EsExchangeCatMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithName name获取 分类名称
func (obj *_EsExchangeCatMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithParentID parent_id获取 父分类
func (obj *_EsExchangeCatMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithCategoryPath category_path获取 分类id路径
func (obj *_EsExchangeCatMgr) WithCategoryPath(categoryPath string) Option {
	return optionFunc(func(o *options) { o.query["category_path"] = categoryPath })
}

// WithGoodsCount goods_count获取 商品数量
func (obj *_EsExchangeCatMgr) WithGoodsCount(goodsCount int) Option {
	return optionFunc(func(o *options) { o.query["goods_count"] = goodsCount })
}

// WithCategoryOrder category_order获取 分类排序
func (obj *_EsExchangeCatMgr) WithCategoryOrder(categoryOrder int) Option {
	return optionFunc(func(o *options) { o.query["category_order"] = categoryOrder })
}

// WithListShow list_show获取 是否在页面上显示
func (obj *_EsExchangeCatMgr) WithListShow(listShow int) Option {
	return optionFunc(func(o *options) { o.query["list_show"] = listShow })
}

// WithImage image获取 分类图片
func (obj *_EsExchangeCatMgr) WithImage(image string) Option {
	return optionFunc(func(o *options) { o.query["image"] = image })
}

// GetByOption 功能选项模式获取
func (obj *_EsExchangeCatMgr) GetByOption(opts ...Option) (result models.EsExchangeCat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsExchangeCatMgr) GetByOptions(opts ...Option) (results []*models.EsExchangeCat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsExchangeCatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsExchangeCat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where(options.query)
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

// GetFromCategoryID 通过category_id获取内容 分类id
func (obj *_EsExchangeCatMgr) GetFromCategoryID(categoryID int) (result models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where("`category_id` = ?", categoryID).First(&result).Error

	return
}

// GetBatchFromCategoryID 批量查找 分类id
func (obj *_EsExchangeCatMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 分类名称
func (obj *_EsExchangeCatMgr) GetFromName(name string) (results []*models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 分类名称
func (obj *_EsExchangeCatMgr) GetBatchFromName(names []string) (results []*models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 父分类
func (obj *_EsExchangeCatMgr) GetFromParentID(parentID int) (results []*models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 父分类
func (obj *_EsExchangeCatMgr) GetBatchFromParentID(parentIDs []int) (results []*models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromCategoryPath 通过category_path获取内容 分类id路径
func (obj *_EsExchangeCatMgr) GetFromCategoryPath(categoryPath string) (results []*models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where("`category_path` = ?", categoryPath).Find(&results).Error

	return
}

// GetBatchFromCategoryPath 批量查找 分类id路径
func (obj *_EsExchangeCatMgr) GetBatchFromCategoryPath(categoryPaths []string) (results []*models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where("`category_path` IN (?)", categoryPaths).Find(&results).Error

	return
}

// GetFromGoodsCount 通过goods_count获取内容 商品数量
func (obj *_EsExchangeCatMgr) GetFromGoodsCount(goodsCount int) (results []*models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where("`goods_count` = ?", goodsCount).Find(&results).Error

	return
}

// GetBatchFromGoodsCount 批量查找 商品数量
func (obj *_EsExchangeCatMgr) GetBatchFromGoodsCount(goodsCounts []int) (results []*models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where("`goods_count` IN (?)", goodsCounts).Find(&results).Error

	return
}

// GetFromCategoryOrder 通过category_order获取内容 分类排序
func (obj *_EsExchangeCatMgr) GetFromCategoryOrder(categoryOrder int) (results []*models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where("`category_order` = ?", categoryOrder).Find(&results).Error

	return
}

// GetBatchFromCategoryOrder 批量查找 分类排序
func (obj *_EsExchangeCatMgr) GetBatchFromCategoryOrder(categoryOrders []int) (results []*models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where("`category_order` IN (?)", categoryOrders).Find(&results).Error

	return
}

// GetFromListShow 通过list_show获取内容 是否在页面上显示
func (obj *_EsExchangeCatMgr) GetFromListShow(listShow int) (results []*models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where("`list_show` = ?", listShow).Find(&results).Error

	return
}

// GetBatchFromListShow 批量查找 是否在页面上显示
func (obj *_EsExchangeCatMgr) GetBatchFromListShow(listShows []int) (results []*models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where("`list_show` IN (?)", listShows).Find(&results).Error

	return
}

// GetFromImage 通过image获取内容 分类图片
func (obj *_EsExchangeCatMgr) GetFromImage(image string) (results []*models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where("`image` = ?", image).Find(&results).Error

	return
}

// GetBatchFromImage 批量查找 分类图片
func (obj *_EsExchangeCatMgr) GetBatchFromImage(images []string) (results []*models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where("`image` IN (?)", images).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsExchangeCatMgr) FetchByPrimaryKey(categoryID int) (result models.EsExchangeCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExchangeCat{}).Where("`category_id` = ?", categoryID).First(&result).Error

	return
}
