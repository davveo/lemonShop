package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsConnectSettingMgr struct {
	*_BaseMgr
}

// EsConnectSettingMgr open func
func EsConnectSettingMgr(db *gorm.DB) *_EsConnectSettingMgr {
	if db == nil {
		panic(fmt.Errorf("EsConnectSettingMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsConnectSettingMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_connect_setting"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsConnectSettingMgr) GetTableName() string {
	return "es_connect_setting"
}

// Reset 重置gorm会话
func (obj *_EsConnectSettingMgr) Reset() *_EsConnectSettingMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsConnectSettingMgr) Get() (result models.EsConnectSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnectSetting{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsConnectSettingMgr) Gets() (results []*models.EsConnectSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsConnectSettingMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_EsConnectSettingMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithType type获取 授权类型
func (obj *_EsConnectSettingMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithConfig config获取 参数组
func (obj *_EsConnectSettingMgr) WithConfig(config string) Option {
	return optionFunc(func(o *options) { o.query["config"] = config })
}

// WithName name获取 标题
func (obj *_EsConnectSettingMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *_EsConnectSettingMgr) GetByOption(opts ...Option) (result models.EsConnectSetting, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsConnectSettingMgr) GetByOptions(opts ...Option) (results []*models.EsConnectSetting, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsConnectSettingMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsConnectSetting, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where(options.query)
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

// GetFromID 通过id获取内容 id
func (obj *_EsConnectSettingMgr) GetFromID(id int) (result models.EsConnectSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_EsConnectSettingMgr) GetBatchFromID(ids []int) (results []*models.EsConnectSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 授权类型
func (obj *_EsConnectSettingMgr) GetFromType(_type string) (results []*models.EsConnectSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 授权类型
func (obj *_EsConnectSettingMgr) GetBatchFromType(_types []string) (results []*models.EsConnectSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromConfig 通过config获取内容 参数组
func (obj *_EsConnectSettingMgr) GetFromConfig(config string) (results []*models.EsConnectSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where("`config` = ?", config).Find(&results).Error

	return
}

// GetBatchFromConfig 批量查找 参数组
func (obj *_EsConnectSettingMgr) GetBatchFromConfig(configs []string) (results []*models.EsConnectSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where("`config` IN (?)", configs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 标题
func (obj *_EsConnectSettingMgr) GetFromName(name string) (results []*models.EsConnectSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 标题
func (obj *_EsConnectSettingMgr) GetBatchFromName(names []string) (results []*models.EsConnectSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsConnectSettingMgr) FetchByPrimaryKey(id int) (result models.EsConnectSetting, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where("`id` = ?", id).First(&result).Error

	return
}
