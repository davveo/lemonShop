package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsKeywordSearchHistoryMgr struct {
	*_BaseMgr
}

// EsKeywordSearchHistoryMgr open func
func EsKeywordSearchHistoryMgr(db *gorm.DB) *_EsKeywordSearchHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("EsKeywordSearchHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsKeywordSearchHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_keyword_search_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsKeywordSearchHistoryMgr) GetTableName() string {
	return "es_keyword_search_history"
}

// Reset 重置gorm会话
func (obj *_EsKeywordSearchHistoryMgr) Reset() *_EsKeywordSearchHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsKeywordSearchHistoryMgr) Get() (result models.EsKeywordSearchHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsKeywordSearchHistoryMgr) Gets() (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsKeywordSearchHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_EsKeywordSearchHistoryMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithKeyword keyword获取 关键字
func (obj *_EsKeywordSearchHistoryMgr) WithKeyword(keyword string) Option {
	return optionFunc(func(o *options) { o.query["keyword"] = keyword })
}

// WithCount count获取 搜索次数
func (obj *_EsKeywordSearchHistoryMgr) WithCount(count int) Option {
	return optionFunc(func(o *options) { o.query["count"] = count })
}

// WithAddTime add_time获取 新增时间
func (obj *_EsKeywordSearchHistoryMgr) WithAddTime(addTime int64) Option {
	return optionFunc(func(o *options) { o.query["add_time"] = addTime })
}

// WithModifyTime modify_time获取 更新时间
func (obj *_EsKeywordSearchHistoryMgr) WithModifyTime(modifyTime int64) Option {
	return optionFunc(func(o *options) { o.query["modify_time"] = modifyTime })
}

// GetByOption 功能选项模式获取
func (obj *_EsKeywordSearchHistoryMgr) GetByOption(opts ...Option) (result models.EsKeywordSearchHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsKeywordSearchHistoryMgr) GetByOptions(opts ...Option) (results []*models.EsKeywordSearchHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsKeywordSearchHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsKeywordSearchHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *_EsKeywordSearchHistoryMgr) GetFromID(id int) (result models.EsKeywordSearchHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_EsKeywordSearchHistoryMgr) GetBatchFromID(ids []int) (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromKeyword 通过keyword获取内容 关键字
func (obj *_EsKeywordSearchHistoryMgr) GetFromKeyword(keyword string) (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`keyword` = ?", keyword).Find(&results).Error

	return
}

// GetBatchFromKeyword 批量查找 关键字
func (obj *_EsKeywordSearchHistoryMgr) GetBatchFromKeyword(keywords []string) (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`keyword` IN (?)", keywords).Find(&results).Error

	return
}

// GetFromCount 通过count获取内容 搜索次数
func (obj *_EsKeywordSearchHistoryMgr) GetFromCount(count int) (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`count` = ?", count).Find(&results).Error

	return
}

// GetBatchFromCount 批量查找 搜索次数
func (obj *_EsKeywordSearchHistoryMgr) GetBatchFromCount(counts []int) (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`count` IN (?)", counts).Find(&results).Error

	return
}

// GetFromAddTime 通过add_time获取内容 新增时间
func (obj *_EsKeywordSearchHistoryMgr) GetFromAddTime(addTime int64) (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`add_time` = ?", addTime).Find(&results).Error

	return
}

// GetBatchFromAddTime 批量查找 新增时间
func (obj *_EsKeywordSearchHistoryMgr) GetBatchFromAddTime(addTimes []int64) (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`add_time` IN (?)", addTimes).Find(&results).Error

	return
}

// GetFromModifyTime 通过modify_time获取内容 更新时间
func (obj *_EsKeywordSearchHistoryMgr) GetFromModifyTime(modifyTime int64) (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`modify_time` = ?", modifyTime).Find(&results).Error

	return
}

// GetBatchFromModifyTime 批量查找 更新时间
func (obj *_EsKeywordSearchHistoryMgr) GetBatchFromModifyTime(modifyTimes []int64) (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`modify_time` IN (?)", modifyTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsKeywordSearchHistoryMgr) FetchByPrimaryKey(id int) (result models.EsKeywordSearchHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`id` = ?", id).First(&result).Error

	return
}
