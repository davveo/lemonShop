package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsSettingsMgr struct {
	*_BaseMgr
}

// EsSettingsMgr open func
func EsSettingsMgr(db *gorm.DB) *_EsSettingsMgr {
	if db == nil {
		panic(fmt.Errorf("EsSettingsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsSettingsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_settings"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsSettingsMgr) GetTableName() string {
	return "es_settings"
}

// Reset 重置gorm会话
func (obj *_EsSettingsMgr) Reset() *_EsSettingsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsSettingsMgr) Get() (result models.EsSettings, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSettings{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsSettingsMgr) Gets() (results []*models.EsSettings, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSettings{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsSettingsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsSettings{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 系统设置id
func (obj *_EsSettingsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithCfgValue cfg_value获取 系统配置信息
func (obj *_EsSettingsMgr) WithCfgValue(cfgValue string) Option {
	return optionFunc(func(o *options) { o.query["cfg_value"] = cfgValue })
}

// WithCfgGroup cfg_group获取 业务设置标识
func (obj *_EsSettingsMgr) WithCfgGroup(cfgGroup string) Option {
	return optionFunc(func(o *options) { o.query["cfg_group"] = cfgGroup })
}

// GetByOption 功能选项模式获取
func (obj *_EsSettingsMgr) GetByOption(opts ...Option) (result models.EsSettings, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSettings{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsSettingsMgr) GetByOptions(opts ...Option) (results []*models.EsSettings, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSettings{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsSettingsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSettings, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsSettings{}).Where(options.query)
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

// GetFromID 通过id获取内容 系统设置id
func (obj *_EsSettingsMgr) GetFromID(id int) (result models.EsSettings, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSettings{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 系统设置id
func (obj *_EsSettingsMgr) GetBatchFromID(ids []int) (results []*models.EsSettings, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSettings{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromCfgValue 通过cfg_value获取内容 系统配置信息
func (obj *_EsSettingsMgr) GetFromCfgValue(cfgValue string) (results []*models.EsSettings, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSettings{}).Where("`cfg_value` = ?", cfgValue).Find(&results).Error

	return
}

// GetBatchFromCfgValue 批量查找 系统配置信息
func (obj *_EsSettingsMgr) GetBatchFromCfgValue(cfgValues []string) (results []*models.EsSettings, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSettings{}).Where("`cfg_value` IN (?)", cfgValues).Find(&results).Error

	return
}

// GetFromCfgGroup 通过cfg_group获取内容 业务设置标识
func (obj *_EsSettingsMgr) GetFromCfgGroup(cfgGroup string) (results []*models.EsSettings, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSettings{}).Where("`cfg_group` = ?", cfgGroup).Find(&results).Error

	return
}

// GetBatchFromCfgGroup 批量查找 业务设置标识
func (obj *_EsSettingsMgr) GetBatchFromCfgGroup(cfgGroups []string) (results []*models.EsSettings, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSettings{}).Where("`cfg_group` IN (?)", cfgGroups).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsSettingsMgr) FetchByPrimaryKey(id int) (result models.EsSettings, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSettings{}).Where("`id` = ?", id).First(&result).Error

	return
}
