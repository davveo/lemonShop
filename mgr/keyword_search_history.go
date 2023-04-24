package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type KeywordSearchHistoryMgr struct {
	*_BaseMgr
}

// NewKeywordSearchHistoryMgr open func
func NewKeywordSearchHistoryMgr(db db.Repo) *KeywordSearchHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("NewKeywordSearchHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &KeywordSearchHistoryMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_keyword_search_history"), wdb: db.GetDbW().Table("es_keyword_search_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *KeywordSearchHistoryMgr) GetTableName() string {
	return "es_keyword_search_history"
}

// Reset 重置gorm会话
func (obj *KeywordSearchHistoryMgr) Reset() *KeywordSearchHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *KeywordSearchHistoryMgr) Get() (result models.EsKeywordSearchHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *KeywordSearchHistoryMgr) Gets() (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *KeywordSearchHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *KeywordSearchHistoryMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithKeyword keyword获取 关键字
func (obj *KeywordSearchHistoryMgr) WithKeyword(keyword string) Option {
	return optionFunc(func(o *options) { o.query["keyword"] = keyword })
}

// WithCount count获取 搜索次数
func (obj *KeywordSearchHistoryMgr) WithCount(count int) Option {
	return optionFunc(func(o *options) { o.query["count"] = count })
}

// WithAddTime add_time获取 新增时间
func (obj *KeywordSearchHistoryMgr) WithAddTime(addTime int64) Option {
	return optionFunc(func(o *options) { o.query["add_time"] = addTime })
}

// WithModifyTime modify_time获取 更新时间
func (obj *KeywordSearchHistoryMgr) WithModifyTime(modifyTime int64) Option {
	return optionFunc(func(o *options) { o.query["modify_time"] = modifyTime })
}

// GetByOption 功能选项模式获取
func (obj *KeywordSearchHistoryMgr) GetByOption(opts ...Option) (result models.EsKeywordSearchHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *KeywordSearchHistoryMgr) GetByOptions(opts ...Option) (results []*models.EsKeywordSearchHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *KeywordSearchHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsKeywordSearchHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where(options.query)
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
func (obj *KeywordSearchHistoryMgr) GetFromID(id int) (result models.EsKeywordSearchHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *KeywordSearchHistoryMgr) GetBatchFromID(ids []int) (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromKeyword 通过keyword获取内容 关键字
func (obj *KeywordSearchHistoryMgr) GetFromKeyword(keyword string) (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`keyword` = ?", keyword).Find(&results).Error

	return
}

// GetBatchFromKeyword 批量查找 关键字
func (obj *KeywordSearchHistoryMgr) GetBatchFromKeyword(keywords []string) (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`keyword` IN (?)", keywords).Find(&results).Error

	return
}

// GetFromCount 通过count获取内容 搜索次数
func (obj *KeywordSearchHistoryMgr) GetFromCount(count int) (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`count` = ?", count).Find(&results).Error

	return
}

// GetBatchFromCount 批量查找 搜索次数
func (obj *KeywordSearchHistoryMgr) GetBatchFromCount(counts []int) (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`count` IN (?)", counts).Find(&results).Error

	return
}

// GetFromAddTime 通过add_time获取内容 新增时间
func (obj *KeywordSearchHistoryMgr) GetFromAddTime(addTime int64) (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`add_time` = ?", addTime).Find(&results).Error

	return
}

// GetBatchFromAddTime 批量查找 新增时间
func (obj *KeywordSearchHistoryMgr) GetBatchFromAddTime(addTimes []int64) (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`add_time` IN (?)", addTimes).Find(&results).Error

	return
}

// GetFromModifyTime 通过modify_time获取内容 更新时间
func (obj *KeywordSearchHistoryMgr) GetFromModifyTime(modifyTime int64) (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`modify_time` = ?", modifyTime).Find(&results).Error

	return
}

// GetBatchFromModifyTime 批量查找 更新时间
func (obj *KeywordSearchHistoryMgr) GetBatchFromModifyTime(modifyTimes []int64) (results []*models.EsKeywordSearchHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`modify_time` IN (?)", modifyTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *KeywordSearchHistoryMgr) FetchByPrimaryKey(id int) (result models.EsKeywordSearchHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsKeywordSearchHistory{}).Where("`id` = ?", id).First(&result).Error

	return
}
