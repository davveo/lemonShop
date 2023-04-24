package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type ArticleMgr struct {
	*_BaseMgr
}

// EsArticleMgr open func
func EsArticleMgr(db db.Repo) *ArticleMgr {
	if db == nil {
		panic(fmt.Errorf("EsArticleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &ArticleMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_article"),
		wdb:       db.GetDbW().Table("es_article"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *ArticleMgr) GetTableName() string {
	return "es_article"
}

func (obj *ArticleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithArticleID article_id获取 主键
func (obj *ArticleMgr) WithArticleID(articleID int) Option {
	return optionFunc(func(o *options) { o.query["article_id"] = articleID })
}

// WithArticleName article_name获取 文章名称
func (obj *ArticleMgr) WithArticleName(articleName string) Option {
	return optionFunc(func(o *options) { o.query["article_name"] = articleName })
}

// WithCategoryID category_id获取 分类id
func (obj *ArticleMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithSort sort获取 文章排序
func (obj *ArticleMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithOutsideURL outside_url获取 外链url
func (obj *ArticleMgr) WithOutsideURL(outsideURL string) Option {
	return optionFunc(func(o *options) { o.query["outside_url"] = outsideURL })
}

// WithContent content获取 文章内容
func (obj *ArticleMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithShowPosition show_position获取 显示位置
func (obj *ArticleMgr) WithShowPosition(showPosition string) Option {
	return optionFunc(func(o *options) { o.query["show_position"] = showPosition })
}

// WithCreateTime create_time获取 添加时间
func (obj *ArticleMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithModifyTime modify_time获取 修改时间
func (obj *ArticleMgr) WithModifyTime(modifyTime int64) Option {
	return optionFunc(func(o *options) { o.query["modify_time"] = modifyTime })
}

// GetByOption 功能选项模式获取
func (obj *ArticleMgr) GetByOption(opts ...Option) (result models.EsArticle, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *ArticleMgr) GetByOptions(opts ...Option) (results []*models.EsArticle, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *ArticleMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsArticle, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where(options.query)
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
func (obj *ArticleMgr) GetFromArticleID(articleID int) (result models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`article_id` = ?", articleID).First(&result).Error

	return
}

// GetBatchFromArticleID 批量查找 主键
func (obj *ArticleMgr) GetBatchFromArticleID(articleIDs []int) (results []*models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`article_id` IN (?)", articleIDs).Find(&results).Error

	return
}

// GetFromArticleName 通过article_name获取内容 文章名称
func (obj *ArticleMgr) GetFromArticleName(articleName string) (results []*models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`article_name` = ?", articleName).Find(&results).Error

	return
}

// GetBatchFromArticleName 批量查找 文章名称
func (obj *ArticleMgr) GetBatchFromArticleName(articleNames []string) (results []*models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`article_name` IN (?)", articleNames).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容 分类id
func (obj *ArticleMgr) GetFromCategoryID(categoryID int) (results []*models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 分类id
func (obj *ArticleMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 文章排序
func (obj *ArticleMgr) GetFromSort(sort int) (results []*models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 文章排序
func (obj *ArticleMgr) GetBatchFromSort(sorts []int) (results []*models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromOutsideURL 通过outside_url获取内容 外链url
func (obj *ArticleMgr) GetFromOutsideURL(outsideURL string) (results []*models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`outside_url` = ?", outsideURL).Find(&results).Error

	return
}

// GetBatchFromOutsideURL 批量查找 外链url
func (obj *ArticleMgr) GetBatchFromOutsideURL(outsideURLs []string) (results []*models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`outside_url` IN (?)", outsideURLs).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 文章内容
func (obj *ArticleMgr) GetFromContent(content string) (results []*models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 文章内容
func (obj *ArticleMgr) GetBatchFromContent(contents []string) (results []*models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromShowPosition 通过show_position获取内容 显示位置
func (obj *ArticleMgr) GetFromShowPosition(showPosition string) (results []*models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`show_position` = ?", showPosition).Find(&results).Error

	return
}

// GetBatchFromShowPosition 批量查找 显示位置
func (obj *ArticleMgr) GetBatchFromShowPosition(showPositions []string) (results []*models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`show_position` IN (?)", showPositions).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 添加时间
func (obj *ArticleMgr) GetFromCreateTime(createTime int64) (results []*models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 添加时间
func (obj *ArticleMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromModifyTime 通过modify_time获取内容 修改时间
func (obj *ArticleMgr) GetFromModifyTime(modifyTime int64) (results []*models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`modify_time` = ?", modifyTime).Find(&results).Error

	return
}

// GetBatchFromModifyTime 批量查找 修改时间
func (obj *ArticleMgr) GetBatchFromModifyTime(modifyTimes []int64) (results []*models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`modify_time` IN (?)", modifyTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *ArticleMgr) FetchByPrimaryKey(articleID int) (result models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`article_id` = ?", articleID).First(&result).Error

	return
}

// FetchIndexByIndArticleCatid  获取多个内容
func (obj *ArticleMgr) FetchIndexByIndArticleCatid(categoryID int) (results []*models.EsArticle, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsArticle{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}
