package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type SmsPlatformMgr struct {
	*_BaseMgr
}

// NewSmsPlatformMgr open func
func NewSmsPlatformMgr(db db.Repo) *SmsPlatformMgr {
	if db == nil {
		panic(fmt.Errorf("NewSmsPlatformMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &SmsPlatformMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_sms_platform"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *SmsPlatformMgr) GetTableName() string {
	return "es_sms_platform"
}

// Reset 重置gorm会话
func (obj *SmsPlatformMgr) Reset() *SmsPlatformMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *SmsPlatformMgr) Get() (result models.EsSmsPlatform, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *SmsPlatformMgr) Gets() (results []*models.EsSmsPlatform, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *SmsPlatformMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *SmsPlatformMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 平台名称
func (obj *SmsPlatformMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithOpen open获取 是否开启
func (obj *SmsPlatformMgr) WithOpen(open int) Option {
	return optionFunc(func(o *options) { o.query["open"] = open })
}

// WithConfig config获取 配置
func (obj *SmsPlatformMgr) WithConfig(config string) Option {
	return optionFunc(func(o *options) { o.query["config"] = config })
}

// WithBean bean获取 beanid
func (obj *SmsPlatformMgr) WithBean(bean string) Option {
	return optionFunc(func(o *options) { o.query["bean"] = bean })
}

// GetByOption 功能选项模式获取
func (obj *SmsPlatformMgr) GetByOption(opts ...Option) (result models.EsSmsPlatform, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *SmsPlatformMgr) GetByOptions(opts ...Option) (results []*models.EsSmsPlatform, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *SmsPlatformMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSmsPlatform, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where(options.query)
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
func (obj *SmsPlatformMgr) GetFromID(id int) (result models.EsSmsPlatform, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *SmsPlatformMgr) GetBatchFromID(ids []int) (results []*models.EsSmsPlatform, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 平台名称
func (obj *SmsPlatformMgr) GetFromName(name string) (results []*models.EsSmsPlatform, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 平台名称
func (obj *SmsPlatformMgr) GetBatchFromName(names []string) (results []*models.EsSmsPlatform, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromOpen 通过open获取内容 是否开启
func (obj *SmsPlatformMgr) GetFromOpen(open int) (results []*models.EsSmsPlatform, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`open` = ?", open).Find(&results).Error

	return
}

// GetBatchFromOpen 批量查找 是否开启
func (obj *SmsPlatformMgr) GetBatchFromOpen(opens []int) (results []*models.EsSmsPlatform, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`open` IN (?)", opens).Find(&results).Error

	return
}

// GetFromConfig 通过config获取内容 配置
func (obj *SmsPlatformMgr) GetFromConfig(config string) (results []*models.EsSmsPlatform, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`config` = ?", config).Find(&results).Error

	return
}

// GetBatchFromConfig 批量查找 配置
func (obj *SmsPlatformMgr) GetBatchFromConfig(configs []string) (results []*models.EsSmsPlatform, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`config` IN (?)", configs).Find(&results).Error

	return
}

// GetFromBean 通过bean获取内容 beanid
func (obj *SmsPlatformMgr) GetFromBean(bean string) (results []*models.EsSmsPlatform, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`bean` = ?", bean).Find(&results).Error

	return
}

// GetBatchFromBean 批量查找 beanid
func (obj *SmsPlatformMgr) GetBatchFromBean(beans []string) (results []*models.EsSmsPlatform, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`bean` IN (?)", beans).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *SmsPlatformMgr) FetchByPrimaryKey(id int) (result models.EsSmsPlatform, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSmsPlatform{}).Where("`id` = ?", id).First(&result).Error

	return
}
