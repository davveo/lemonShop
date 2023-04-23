package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsExpressPlatformMgr struct {
	*_BaseMgr
}

// EsExpressPlatformMgr open func
func EsExpressPlatformMgr(db *gorm.DB) *_EsExpressPlatformMgr {
	if db == nil {
		panic(fmt.Errorf("EsExpressPlatformMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsExpressPlatformMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_express_platform"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsExpressPlatformMgr) GetTableName() string {
	return "es_express_platform"
}

// Reset 重置gorm会话
func (obj *_EsExpressPlatformMgr) Reset() *_EsExpressPlatformMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsExpressPlatformMgr) Get() (result models.EsExpressPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExpressPlatform{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsExpressPlatformMgr) Gets() (results []*models.EsExpressPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExpressPlatform{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsExpressPlatformMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsExpressPlatform{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 快递平台id
func (obj *_EsExpressPlatformMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 快递平台名称
func (obj *_EsExpressPlatformMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithOpen open获取 是否开启快递平台,1开启，0未开启
func (obj *_EsExpressPlatformMgr) WithOpen(open int) Option {
	return optionFunc(func(o *options) { o.query["open"] = open })
}

// WithConfig config获取 快递平台配置
func (obj *_EsExpressPlatformMgr) WithConfig(config string) Option {
	return optionFunc(func(o *options) { o.query["config"] = config })
}

// WithBean bean获取 快递平台beanid
func (obj *_EsExpressPlatformMgr) WithBean(bean string) Option {
	return optionFunc(func(o *options) { o.query["bean"] = bean })
}

// GetByOption 功能选项模式获取
func (obj *_EsExpressPlatformMgr) GetByOption(opts ...Option) (result models.EsExpressPlatform, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsExpressPlatform{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsExpressPlatformMgr) GetByOptions(opts ...Option) (results []*models.EsExpressPlatform, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsExpressPlatform{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsExpressPlatformMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsExpressPlatform, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsExpressPlatform{}).Where(options.query)
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

// GetFromID 通过id获取内容 快递平台id
func (obj *_EsExpressPlatformMgr) GetFromID(id int) (result models.EsExpressPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExpressPlatform{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 快递平台id
func (obj *_EsExpressPlatformMgr) GetBatchFromID(ids []int) (results []*models.EsExpressPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExpressPlatform{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 快递平台名称
func (obj *_EsExpressPlatformMgr) GetFromName(name string) (results []*models.EsExpressPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExpressPlatform{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 快递平台名称
func (obj *_EsExpressPlatformMgr) GetBatchFromName(names []string) (results []*models.EsExpressPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExpressPlatform{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromOpen 通过open获取内容 是否开启快递平台,1开启，0未开启
func (obj *_EsExpressPlatformMgr) GetFromOpen(open int) (results []*models.EsExpressPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExpressPlatform{}).Where("`open` = ?", open).Find(&results).Error

	return
}

// GetBatchFromOpen 批量查找 是否开启快递平台,1开启，0未开启
func (obj *_EsExpressPlatformMgr) GetBatchFromOpen(opens []int) (results []*models.EsExpressPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExpressPlatform{}).Where("`open` IN (?)", opens).Find(&results).Error

	return
}

// GetFromConfig 通过config获取内容 快递平台配置
func (obj *_EsExpressPlatformMgr) GetFromConfig(config string) (results []*models.EsExpressPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExpressPlatform{}).Where("`config` = ?", config).Find(&results).Error

	return
}

// GetBatchFromConfig 批量查找 快递平台配置
func (obj *_EsExpressPlatformMgr) GetBatchFromConfig(configs []string) (results []*models.EsExpressPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExpressPlatform{}).Where("`config` IN (?)", configs).Find(&results).Error

	return
}

// GetFromBean 通过bean获取内容 快递平台beanid
func (obj *_EsExpressPlatformMgr) GetFromBean(bean string) (results []*models.EsExpressPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExpressPlatform{}).Where("`bean` = ?", bean).Find(&results).Error

	return
}

// GetBatchFromBean 批量查找 快递平台beanid
func (obj *_EsExpressPlatformMgr) GetBatchFromBean(beans []string) (results []*models.EsExpressPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExpressPlatform{}).Where("`bean` IN (?)", beans).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsExpressPlatformMgr) FetchByPrimaryKey(id int) (result models.EsExpressPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsExpressPlatform{}).Where("`id` = ?", id).First(&result).Error

	return
}
