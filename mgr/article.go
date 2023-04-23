package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsArticleMgr struct {
	*_BaseMgr
}

// EsArticleMgr open func
func EsArticleMgr(db *gorm.DB) *_EsArticleMgr {
	if db == nil {
		panic(fmt.Errorf("EsArticleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsArticleMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_article"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsArticleMgr) GetTableName() string {
	return "es_article"
}

// Reset 重置gorm会话
func (obj *_EsArticleMgr) Reset() *_EsArticleMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsArticleMgr) Get() (result models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsArticleMgr) Gets() (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsArticleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithArticleID article_id获取 主键
func (obj *_EsArticleMgr) WithArticleID(articleID int) Option {
	return optionFunc(func(o *options) { o.query["article_id"] = articleID })
}

// WithArticleName article_name获取 文章名称
func (obj *_EsArticleMgr) WithArticleName(articleName string) Option {
	return optionFunc(func(o *options) { o.query["article_name"] = articleName })
}

// WithCategoryID category_id获取 分类id
func (obj *_EsArticleMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithSort sort获取 文章排序
func (obj *_EsArticleMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithOutsideURL outside_url获取 外链url
func (obj *_EsArticleMgr) WithOutsideURL(outsideURL string) Option {
	return optionFunc(func(o *options) { o.query["outside_url"] = outsideURL })
}

// WithContent content获取 文章内容
func (obj *_EsArticleMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithShowPosition show_position获取 显示位置
func (obj *_EsArticleMgr) WithShowPosition(showPosition string) Option {
	return optionFunc(func(o *options) { o.query["show_position"] = showPosition })
}

// WithCreateTime create_time获取 添加时间
func (obj *_EsArticleMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithModifyTime modify_time获取 修改时间
func (obj *_EsArticleMgr) WithModifyTime(modifyTime int64) Option {
	return optionFunc(func(o *options) { o.query["modify_time"] = modifyTime })
}

// GetByOption 功能选项模式获取
func (obj *_EsArticleMgr) GetByOption(opts ...Option) (result models.EsArticle, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsArticleMgr) GetByOptions(opts ...Option) (results []*models.EsArticle, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsArticleMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsArticle, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where(options.query)
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

// GetFromArticleID 通过article_id获取内容 主键
func (obj *_EsArticleMgr) GetFromArticleID(articleID int) (result models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`article_id` = ?", articleID).First(&result).Error

	return
}

// GetBatchFromArticleID 批量查找 主键
func (obj *_EsArticleMgr) GetBatchFromArticleID(articleIDs []int) (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`article_id` IN (?)", articleIDs).Find(&results).Error

	return
}

// GetFromArticleName 通过article_name获取内容 文章名称
func (obj *_EsArticleMgr) GetFromArticleName(articleName string) (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`article_name` = ?", articleName).Find(&results).Error

	return
}

// GetBatchFromArticleName 批量查找 文章名称
func (obj *_EsArticleMgr) GetBatchFromArticleName(articleNames []string) (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`article_name` IN (?)", articleNames).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容 分类id
func (obj *_EsArticleMgr) GetFromCategoryID(categoryID int) (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 分类id
func (obj *_EsArticleMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 文章排序
func (obj *_EsArticleMgr) GetFromSort(sort int) (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 文章排序
func (obj *_EsArticleMgr) GetBatchFromSort(sorts []int) (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromOutsideURL 通过outside_url获取内容 外链url
func (obj *_EsArticleMgr) GetFromOutsideURL(outsideURL string) (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`outside_url` = ?", outsideURL).Find(&results).Error

	return
}

// GetBatchFromOutsideURL 批量查找 外链url
func (obj *_EsArticleMgr) GetBatchFromOutsideURL(outsideURLs []string) (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`outside_url` IN (?)", outsideURLs).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 文章内容
func (obj *_EsArticleMgr) GetFromContent(content string) (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 文章内容
func (obj *_EsArticleMgr) GetBatchFromContent(contents []string) (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromShowPosition 通过show_position获取内容 显示位置
func (obj *_EsArticleMgr) GetFromShowPosition(showPosition string) (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`show_position` = ?", showPosition).Find(&results).Error

	return
}

// GetBatchFromShowPosition 批量查找 显示位置
func (obj *_EsArticleMgr) GetBatchFromShowPosition(showPositions []string) (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`show_position` IN (?)", showPositions).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 添加时间
func (obj *_EsArticleMgr) GetFromCreateTime(createTime int64) (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 添加时间
func (obj *_EsArticleMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromModifyTime 通过modify_time获取内容 修改时间
func (obj *_EsArticleMgr) GetFromModifyTime(modifyTime int64) (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`modify_time` = ?", modifyTime).Find(&results).Error

	return
}

// GetBatchFromModifyTime 批量查找 修改时间
func (obj *_EsArticleMgr) GetBatchFromModifyTime(modifyTimes []int64) (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`modify_time` IN (?)", modifyTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsArticleMgr) FetchByPrimaryKey(articleID int) (result models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`article_id` = ?", articleID).First(&result).Error

	return
}

// FetchIndexByIndArticleCatid  获取多个内容
func (obj *_EsArticleMgr) FetchIndexByIndArticleCatid(categoryID int) (results []*models.EsArticle, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}
