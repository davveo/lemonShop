package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type CategoryMgr struct {
	*_BaseMgr
}

// NewCategoryMgr open func
func NewCategoryMgr(db db.Repo) *CategoryMgr {
	if db == nil {
		panic(fmt.Errorf("NewCategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &CategoryMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_category_spec"),
		wdb:       db.GetDbW().Table("es_category_spec"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *CategoryMgr) GetTableName() string {
	return "es_category"
}

// Reset 重置gorm会话
func (obj *CategoryMgr) Reset() *CategoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *CategoryMgr) Get() (result models.EsCategory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *CategoryMgr) Gets() (results []*models.EsCategory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *CategoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCategoryID category_id获取 主键
func (obj *CategoryMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithName name获取 分类名称
func (obj *CategoryMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithParentID parent_id获取 分类父id
func (obj *CategoryMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithCategoryPath category_path获取 分类父子路径
func (obj *CategoryMgr) WithCategoryPath(categoryPath string) Option {
	return optionFunc(func(o *options) { o.query["category_path"] = categoryPath })
}

// WithGoodsCount goods_count获取 该分类下商品数量
func (obj *CategoryMgr) WithGoodsCount(goodsCount int) Option {
	return optionFunc(func(o *options) { o.query["goods_count"] = goodsCount })
}

// WithCategoryOrder category_order获取 分类排序
func (obj *CategoryMgr) WithCategoryOrder(categoryOrder int) Option {
	return optionFunc(func(o *options) { o.query["category_order"] = categoryOrder })
}

// WithImage image获取 分类图标
func (obj *CategoryMgr) WithImage(image string) Option {
	return optionFunc(func(o *options) { o.query["image"] = image })
}

// GetByOption 功能选项模式获取
func (obj *CategoryMgr) GetByOption(opts ...Option) (result models.EsCategory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *CategoryMgr) GetByOptions(opts ...Option) (results []*models.EsCategory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *CategoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsCategory, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where(options.query)
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

// GetFromCategoryID 通过category_id获取内容 主键
func (obj *CategoryMgr) GetFromCategoryID(categoryID int) (result models.EsCategory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`category_id` = ?", categoryID).First(&result).Error

	return
}

// GetBatchFromCategoryID 批量查找 主键
func (obj *CategoryMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsCategory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 分类名称
func (obj *CategoryMgr) GetFromName(name string) (results []*models.EsCategory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 分类名称
func (obj *CategoryMgr) GetBatchFromName(names []string) (results []*models.EsCategory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 分类父id
func (obj *CategoryMgr) GetFromParentID(parentID int) (results []*models.EsCategory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 分类父id
func (obj *CategoryMgr) GetBatchFromParentID(parentIDs []int) (results []*models.EsCategory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromCategoryPath 通过category_path获取内容 分类父子路径
func (obj *CategoryMgr) GetFromCategoryPath(categoryPath string) (results []*models.EsCategory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`category_path` = ?", categoryPath).Find(&results).Error

	return
}

// GetBatchFromCategoryPath 批量查找 分类父子路径
func (obj *CategoryMgr) GetBatchFromCategoryPath(categoryPaths []string) (results []*models.EsCategory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`category_path` IN (?)", categoryPaths).Find(&results).Error

	return
}

// GetFromGoodsCount 通过goods_count获取内容 该分类下商品数量
func (obj *CategoryMgr) GetFromGoodsCount(goodsCount int) (results []*models.EsCategory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`goods_count` = ?", goodsCount).Find(&results).Error

	return
}

// GetBatchFromGoodsCount 批量查找 该分类下商品数量
func (obj *CategoryMgr) GetBatchFromGoodsCount(goodsCounts []int) (results []*models.EsCategory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`goods_count` IN (?)", goodsCounts).Find(&results).Error

	return
}

// GetFromCategoryOrder 通过category_order获取内容 分类排序
func (obj *CategoryMgr) GetFromCategoryOrder(categoryOrder int) (results []*models.EsCategory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`category_order` = ?", categoryOrder).Find(&results).Error

	return
}

// GetBatchFromCategoryOrder 批量查找 分类排序
func (obj *CategoryMgr) GetBatchFromCategoryOrder(categoryOrders []int) (results []*models.EsCategory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`category_order` IN (?)", categoryOrders).Find(&results).Error

	return
}

// GetFromImage 通过image获取内容 分类图标
func (obj *CategoryMgr) GetFromImage(image string) (results []*models.EsCategory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`image` = ?", image).Find(&results).Error

	return
}

// GetBatchFromImage 批量查找 分类图标
func (obj *CategoryMgr) GetBatchFromImage(images []string) (results []*models.EsCategory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`image` IN (?)", images).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *CategoryMgr) FetchByPrimaryKey(categoryID int) (result *models.EsCategory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`category_id` = ?", categoryID).First(&result).Error

	return
}

// FetchIndexByIndGoodsCatParentid  获取多个内容
func (obj *CategoryMgr) FetchIndexByIndGoodsCatParentid(parentID int) (results []*models.EsCategory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCategory{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

func (obj *CategoryMgr) Create(specs *models.EsCategory) (err error) {
	err = obj.wdb.WithContext(obj.ctx).Create(specs).Error

	return
}

func (obj *CategoryMgr) Update(updates map[string]interface{}, opts ...Option) (err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.wdb.WithContext(obj.ctx).Model(models.EsCategory{}).
		Where(options.query).Updates(updates).Error
	return
}

func (obj *CategoryMgr) UpdateByModel(updates *models.EsCategory, opts ...Option) (err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.wdb.WithContext(obj.ctx).Model(models.EsCategory{}).
		Where(options.query).Updates(updates).Error
	return
}
