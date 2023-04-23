package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsGroupbuyCatMgr struct {
	*_BaseMgr
}

// EsGroupbuyCatMgr open func
func EsGroupbuyCatMgr(db *gorm.DB) *_EsGroupbuyCatMgr {
	if db == nil {
		panic(fmt.Errorf("EsGroupbuyCatMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsGroupbuyCatMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_groupbuy_cat"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsGroupbuyCatMgr) GetTableName() string {
	return "es_groupbuy_cat"
}

// Reset 重置gorm会话
func (obj *_EsGroupbuyCatMgr) Reset() *_EsGroupbuyCatMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsGroupbuyCatMgr) Get() (result models.EsGroupbuyCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyCat{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsGroupbuyCatMgr) Gets() (results []*models.EsGroupbuyCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyCat{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsGroupbuyCatMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyCat{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCatID cat_id获取 分类id
func (obj *_EsGroupbuyCatMgr) WithCatID(catID int) Option {
	return optionFunc(func(o *options) { o.query["cat_id"] = catID })
}

// WithParentID parent_id获取 父类id
func (obj *_EsGroupbuyCatMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithCatName cat_name获取 分类名称
func (obj *_EsGroupbuyCatMgr) WithCatName(catName string) Option {
	return optionFunc(func(o *options) { o.query["cat_name"] = catName })
}

// WithCatPath cat_path获取 分类结构目录
func (obj *_EsGroupbuyCatMgr) WithCatPath(catPath string) Option {
	return optionFunc(func(o *options) { o.query["cat_path"] = catPath })
}

// WithCatOrder cat_order获取 分类排序
func (obj *_EsGroupbuyCatMgr) WithCatOrder(catOrder int) Option {
	return optionFunc(func(o *options) { o.query["cat_order"] = catOrder })
}

// GetByOption 功能选项模式获取
func (obj *_EsGroupbuyCatMgr) GetByOption(opts ...Option) (result models.EsGroupbuyCat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyCat{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsGroupbuyCatMgr) GetByOptions(opts ...Option) (results []*models.EsGroupbuyCat, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyCat{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsGroupbuyCatMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsGroupbuyCat, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyCat{}).Where(options.query)
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

// GetFromCatID 通过cat_id获取内容 分类id
func (obj *_EsGroupbuyCatMgr) GetFromCatID(catID int) (result models.EsGroupbuyCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyCat{}).Where("`cat_id` = ?", catID).First(&result).Error

	return
}

// GetBatchFromCatID 批量查找 分类id
func (obj *_EsGroupbuyCatMgr) GetBatchFromCatID(catIDs []int) (results []*models.EsGroupbuyCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyCat{}).Where("`cat_id` IN (?)", catIDs).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 父类id
func (obj *_EsGroupbuyCatMgr) GetFromParentID(parentID int) (results []*models.EsGroupbuyCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyCat{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 父类id
func (obj *_EsGroupbuyCatMgr) GetBatchFromParentID(parentIDs []int) (results []*models.EsGroupbuyCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyCat{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromCatName 通过cat_name获取内容 分类名称
func (obj *_EsGroupbuyCatMgr) GetFromCatName(catName string) (results []*models.EsGroupbuyCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyCat{}).Where("`cat_name` = ?", catName).Find(&results).Error

	return
}

// GetBatchFromCatName 批量查找 分类名称
func (obj *_EsGroupbuyCatMgr) GetBatchFromCatName(catNames []string) (results []*models.EsGroupbuyCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyCat{}).Where("`cat_name` IN (?)", catNames).Find(&results).Error

	return
}

// GetFromCatPath 通过cat_path获取内容 分类结构目录
func (obj *_EsGroupbuyCatMgr) GetFromCatPath(catPath string) (results []*models.EsGroupbuyCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyCat{}).Where("`cat_path` = ?", catPath).Find(&results).Error

	return
}

// GetBatchFromCatPath 批量查找 分类结构目录
func (obj *_EsGroupbuyCatMgr) GetBatchFromCatPath(catPaths []string) (results []*models.EsGroupbuyCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyCat{}).Where("`cat_path` IN (?)", catPaths).Find(&results).Error

	return
}

// GetFromCatOrder 通过cat_order获取内容 分类排序
func (obj *_EsGroupbuyCatMgr) GetFromCatOrder(catOrder int) (results []*models.EsGroupbuyCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyCat{}).Where("`cat_order` = ?", catOrder).Find(&results).Error

	return
}

// GetBatchFromCatOrder 批量查找 分类排序
func (obj *_EsGroupbuyCatMgr) GetBatchFromCatOrder(catOrders []int) (results []*models.EsGroupbuyCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyCat{}).Where("`cat_order` IN (?)", catOrders).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsGroupbuyCatMgr) FetchByPrimaryKey(catID int) (result models.EsGroupbuyCat, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyCat{}).Where("`cat_id` = ?", catID).First(&result).Error

	return
}
