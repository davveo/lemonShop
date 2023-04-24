package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type HotKeywordMgr struct {
	*_BaseMgr
}

// NewHotKeywordMgr open func
func NewHotKeywordMgr(db db.Repo) *HotKeywordMgr {
	if db == nil {
		panic(fmt.Errorf("NewHotKeywordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &HotKeywordMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_hot_keyword"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *HotKeywordMgr) GetTableName() string {
	return "es_hot_keyword"
}

// Reset 重置gorm会话
func (obj *HotKeywordMgr) Reset() *HotKeywordMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *HotKeywordMgr) Get() (result models.EsHotKeyword, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsHotKeyword{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *HotKeywordMgr) Gets() (results []*models.EsHotKeyword, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsHotKeyword{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *HotKeywordMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsHotKeyword{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *HotKeywordMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithHotName hot_name获取 文字内容
func (obj *HotKeywordMgr) WithHotName(hotName string) Option {
	return optionFunc(func(o *options) { o.query["hot_name"] = hotName })
}

// WithSort sort获取 排序
func (obj *HotKeywordMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// GetByOption 功能选项模式获取
func (obj *HotKeywordMgr) GetByOption(opts ...Option) (result models.EsHotKeyword, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsHotKeyword{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *HotKeywordMgr) GetByOptions(opts ...Option) (results []*models.EsHotKeyword, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsHotKeyword{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *HotKeywordMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsHotKeyword, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsHotKeyword{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键ID
func (obj *HotKeywordMgr) GetFromID(id int) (result models.EsHotKeyword, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsHotKeyword{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *HotKeywordMgr) GetBatchFromID(ids []int) (results []*models.EsHotKeyword, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsHotKeyword{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromHotName 通过hot_name获取内容 文字内容
func (obj *HotKeywordMgr) GetFromHotName(hotName string) (results []*models.EsHotKeyword, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsHotKeyword{}).Where("`hot_name` = ?", hotName).Find(&results).Error

	return
}

// GetBatchFromHotName 批量查找 文字内容
func (obj *HotKeywordMgr) GetBatchFromHotName(hotNames []string) (results []*models.EsHotKeyword, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsHotKeyword{}).Where("`hot_name` IN (?)", hotNames).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *HotKeywordMgr) GetFromSort(sort int) (results []*models.EsHotKeyword, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsHotKeyword{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *HotKeywordMgr) GetBatchFromSort(sorts []int) (results []*models.EsHotKeyword, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsHotKeyword{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *HotKeywordMgr) FetchByPrimaryKey(id int) (result models.EsHotKeyword, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsHotKeyword{}).Where("`id` = ?", id).First(&result).Error

	return
}
