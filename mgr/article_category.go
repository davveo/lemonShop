package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsArticleCategoryMgr struct {
	*_BaseMgr
}

// EsArticleCategoryMgr open func
func EsArticleCategoryMgr(db *gorm.DB) *_EsArticleCategoryMgr {
	if db == nil {
		panic(fmt.Errorf("EsArticleCategoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsArticleCategoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_article_category"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsArticleCategoryMgr) GetTableName() string {
	return "es_article_category"
}

// Reset 重置gorm会话
func (obj *_EsArticleCategoryMgr) Reset() *_EsArticleCategoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsArticleCategoryMgr) Get() (result models.EsArticleCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsArticleCategoryMgr) Gets() (results []*models.EsArticleCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsArticleCategoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_EsArticleCategoryMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 分类名称
func (obj *_EsArticleCategoryMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithParentID parent_id获取 父分类id
func (obj *_EsArticleCategoryMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithPath path获取 父子路径0|10|
func (obj *_EsArticleCategoryMgr) WithPath(path string) Option {
	return optionFunc(func(o *options) { o.query["path"] = path })
}

// WithAllowDelete allow_delete获取 是否允许删除1允许 0不允许
func (obj *_EsArticleCategoryMgr) WithAllowDelete(allowDelete int16) Option {
	return optionFunc(func(o *options) { o.query["allow_delete"] = allowDelete })
}

// WithType type获取 分类类型，枚举值
func (obj *_EsArticleCategoryMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithSort sort获取 排序，正序123
func (obj *_EsArticleCategoryMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// GetByOption 功能选项模式获取
func (obj *_EsArticleCategoryMgr) GetByOption(opts ...Option) (result models.EsArticleCategory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsArticleCategoryMgr) GetByOptions(opts ...Option) (results []*models.EsArticleCategory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsArticleCategoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsArticleCategory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键
func (obj *_EsArticleCategoryMgr) GetFromID(id int) (result models.EsArticleCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_EsArticleCategoryMgr) GetBatchFromID(ids []int) (results []*models.EsArticleCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 分类名称
func (obj *_EsArticleCategoryMgr) GetFromName(name string) (results []*models.EsArticleCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 分类名称
func (obj *_EsArticleCategoryMgr) GetBatchFromName(names []string) (results []*models.EsArticleCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 父分类id
func (obj *_EsArticleCategoryMgr) GetFromParentID(parentID int) (results []*models.EsArticleCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 父分类id
func (obj *_EsArticleCategoryMgr) GetBatchFromParentID(parentIDs []int) (results []*models.EsArticleCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromPath 通过path获取内容 父子路径0|10|
func (obj *_EsArticleCategoryMgr) GetFromPath(path string) (results []*models.EsArticleCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Where("`path` = ?", path).Find(&results).Error

	return
}

// GetBatchFromPath 批量查找 父子路径0|10|
func (obj *_EsArticleCategoryMgr) GetBatchFromPath(paths []string) (results []*models.EsArticleCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Where("`path` IN (?)", paths).Find(&results).Error

	return
}

// GetFromAllowDelete 通过allow_delete获取内容 是否允许删除1允许 0不允许
func (obj *_EsArticleCategoryMgr) GetFromAllowDelete(allowDelete int16) (results []*models.EsArticleCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Where("`allow_delete` = ?", allowDelete).Find(&results).Error

	return
}

// GetBatchFromAllowDelete 批量查找 是否允许删除1允许 0不允许
func (obj *_EsArticleCategoryMgr) GetBatchFromAllowDelete(allowDeletes []int16) (results []*models.EsArticleCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Where("`allow_delete` IN (?)", allowDeletes).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 分类类型，枚举值
func (obj *_EsArticleCategoryMgr) GetFromType(_type string) (results []*models.EsArticleCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 分类类型，枚举值
func (obj *_EsArticleCategoryMgr) GetBatchFromType(_types []string) (results []*models.EsArticleCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序，正序123
func (obj *_EsArticleCategoryMgr) GetFromSort(sort int) (results []*models.EsArticleCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序，正序123
func (obj *_EsArticleCategoryMgr) GetBatchFromSort(sorts []int) (results []*models.EsArticleCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsArticleCategoryMgr) FetchByPrimaryKey(id int) (result models.EsArticleCategory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticleCategory{}).Where("`id` = ?", id).First(&result).Error

	return
}
