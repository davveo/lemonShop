package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type CustomWordsMgr struct {
	*_BaseMgr
}

// NewCustomWordsMgr open func
func NewCustomWordsMgr(db db.Repo) *CustomWordsMgr {
	if db == nil {
		panic(fmt.Errorf("NewCustomWordsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &CustomWordsMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_custom_words"),
		wdb:       db.GetDbW().Table("es_custom_words"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *CustomWordsMgr) GetTableName() string {
	return "es_custom_words"
}

// Reset 重置gorm会话
func (obj *CustomWordsMgr) Reset() *CustomWordsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *CustomWordsMgr) Get() (result models.EsCustomWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCustomWords{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *CustomWordsMgr) Gets() (results []*models.EsCustomWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCustomWords{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *CustomWordsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsCustomWords{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *CustomWordsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *CustomWordsMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithAddTime add_time获取 添加时间
func (obj *CustomWordsMgr) WithAddTime(addTime int64) Option {
	return optionFunc(func(o *options) { o.query["add_time"] = addTime })
}

// WithDisabled disabled获取 显示 1  隐藏 0
func (obj *CustomWordsMgr) WithDisabled(disabled int16) Option {
	return optionFunc(func(o *options) { o.query["disabled"] = disabled })
}

// WithModifyTime modify_time获取 修改时间
func (obj *CustomWordsMgr) WithModifyTime(modifyTime int64) Option {
	return optionFunc(func(o *options) { o.query["modify_time"] = modifyTime })
}

// GetByOption 功能选项模式获取
func (obj *CustomWordsMgr) GetByOption(opts ...Option) (result models.EsCustomWords, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCustomWords{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *CustomWordsMgr) GetByOptions(opts ...Option) (results []*models.EsCustomWords, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCustomWords{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *CustomWordsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsCustomWords, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsCustomWords{}).Where(options.query)
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
func (obj *CustomWordsMgr) GetFromID(id int) (result models.EsCustomWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCustomWords{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *CustomWordsMgr) GetBatchFromID(ids []int) (results []*models.EsCustomWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCustomWords{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *CustomWordsMgr) GetFromName(name string) (results []*models.EsCustomWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCustomWords{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *CustomWordsMgr) GetBatchFromName(names []string) (results []*models.EsCustomWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCustomWords{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromAddTime 通过add_time获取内容 添加时间
func (obj *CustomWordsMgr) GetFromAddTime(addTime int64) (results []*models.EsCustomWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCustomWords{}).Where("`add_time` = ?", addTime).Find(&results).Error

	return
}

// GetBatchFromAddTime 批量查找 添加时间
func (obj *CustomWordsMgr) GetBatchFromAddTime(addTimes []int64) (results []*models.EsCustomWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCustomWords{}).Where("`add_time` IN (?)", addTimes).Find(&results).Error

	return
}

// GetFromDisabled 通过disabled获取内容 显示 1  隐藏 0
func (obj *CustomWordsMgr) GetFromDisabled(disabled int16) (results []*models.EsCustomWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCustomWords{}).Where("`disabled` = ?", disabled).Find(&results).Error

	return
}

// GetBatchFromDisabled 批量查找 显示 1  隐藏 0
func (obj *CustomWordsMgr) GetBatchFromDisabled(disableds []int16) (results []*models.EsCustomWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCustomWords{}).Where("`disabled` IN (?)", disableds).Find(&results).Error

	return
}

// GetFromModifyTime 通过modify_time获取内容 修改时间
func (obj *CustomWordsMgr) GetFromModifyTime(modifyTime int64) (results []*models.EsCustomWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCustomWords{}).Where("`modify_time` = ?", modifyTime).Find(&results).Error

	return
}

// GetBatchFromModifyTime 批量查找 修改时间
func (obj *CustomWordsMgr) GetBatchFromModifyTime(modifyTimes []int64) (results []*models.EsCustomWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCustomWords{}).Where("`modify_time` IN (?)", modifyTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *CustomWordsMgr) FetchByPrimaryKey(id int) (result models.EsCustomWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCustomWords{}).Where("`id` = ?", id).First(&result).Error

	return
}
