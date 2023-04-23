package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsSmsPlatformMgr struct {
	*_BaseMgr
}

// EsSmsPlatformMgr open func
func EsSmsPlatformMgr(db *gorm.DB) *_EsSmsPlatformMgr {
	if db == nil {
		panic(fmt.Errorf("EsSmsPlatformMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsSmsPlatformMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_sms_platform"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsSmsPlatformMgr) GetTableName() string {
	return "es_sms_platform"
}

// Reset 重置gorm会话
func (obj *_EsSmsPlatformMgr) Reset() *_EsSmsPlatformMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsSmsPlatformMgr) Get() (result models.EsSmsPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsSmsPlatformMgr) Gets() (results []*models.EsSmsPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsSmsPlatformMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *_EsSmsPlatformMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 平台名称
func (obj *_EsSmsPlatformMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithOpen open获取 是否开启
func (obj *_EsSmsPlatformMgr) WithOpen(open int) Option {
	return optionFunc(func(o *options) { o.query["open"] = open })
}

// WithConfig config获取 配置
func (obj *_EsSmsPlatformMgr) WithConfig(config string) Option {
	return optionFunc(func(o *options) { o.query["config"] = config })
}

// WithBean bean获取 beanid
func (obj *_EsSmsPlatformMgr) WithBean(bean string) Option {
	return optionFunc(func(o *options) { o.query["bean"] = bean })
}

// GetByOption 功能选项模式获取
func (obj *_EsSmsPlatformMgr) GetByOption(opts ...Option) (result models.EsSmsPlatform, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsSmsPlatformMgr) GetByOptions(opts ...Option) (results []*models.EsSmsPlatform, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsSmsPlatformMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSmsPlatform, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where(options.query)
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
func (obj *_EsSmsPlatformMgr) GetFromID(id int) (result models.EsSmsPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *_EsSmsPlatformMgr) GetBatchFromID(ids []int) (results []*models.EsSmsPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 平台名称
func (obj *_EsSmsPlatformMgr) GetFromName(name string) (results []*models.EsSmsPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 平台名称
func (obj *_EsSmsPlatformMgr) GetBatchFromName(names []string) (results []*models.EsSmsPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromOpen 通过open获取内容 是否开启
func (obj *_EsSmsPlatformMgr) GetFromOpen(open int) (results []*models.EsSmsPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`open` = ?", open).Find(&results).Error

	return
}

// GetBatchFromOpen 批量查找 是否开启
func (obj *_EsSmsPlatformMgr) GetBatchFromOpen(opens []int) (results []*models.EsSmsPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`open` IN (?)", opens).Find(&results).Error

	return
}

// GetFromConfig 通过config获取内容 配置
func (obj *_EsSmsPlatformMgr) GetFromConfig(config string) (results []*models.EsSmsPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`config` = ?", config).Find(&results).Error

	return
}

// GetBatchFromConfig 批量查找 配置
func (obj *_EsSmsPlatformMgr) GetBatchFromConfig(configs []string) (results []*models.EsSmsPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`config` IN (?)", configs).Find(&results).Error

	return
}

// GetFromBean 通过bean获取内容 beanid
func (obj *_EsSmsPlatformMgr) GetFromBean(bean string) (results []*models.EsSmsPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`bean` = ?", bean).Find(&results).Error

	return
}

// GetBatchFromBean 批量查找 beanid
func (obj *_EsSmsPlatformMgr) GetBatchFromBean(beans []string) (results []*models.EsSmsPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`bean` IN (?)", beans).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsSmsPlatformMgr) FetchByPrimaryKey(id int) (result models.EsSmsPlatform, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`id` = ?", id).First(&result).Error

	return
}
