package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type WaybillMgr struct {
	*_BaseMgr
}

// NewWaybillMgr open func
func NewWaybillMgr(db db.Repo) *WaybillMgr {
	if db == nil {
		panic(fmt.Errorf("NewWaybillMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &WaybillMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_waybill"), wdb: db.GetDbW().Table("es_waybill"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *WaybillMgr) GetTableName() string {
	return "es_waybill"
}

// Reset 重置gorm会话
func (obj *WaybillMgr) Reset() *WaybillMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *WaybillMgr) Get() (result models.EsWaybill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWaybill{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *WaybillMgr) Gets() (results []*models.EsWaybill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWaybill{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *WaybillMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsWaybill{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 电子面单id
func (obj *WaybillMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *WaybillMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithOpen open获取 是否开启
func (obj *WaybillMgr) WithOpen(open int) Option {
	return optionFunc(func(o *options) { o.query["open"] = open })
}

// WithConfig config获取 电子面单配置
func (obj *WaybillMgr) WithConfig(config string) Option {
	return optionFunc(func(o *options) { o.query["config"] = config })
}

// WithBean bean获取 beanid
func (obj *WaybillMgr) WithBean(bean string) Option {
	return optionFunc(func(o *options) { o.query["bean"] = bean })
}

// GetByOption 功能选项模式获取
func (obj *WaybillMgr) GetByOption(opts ...Option) (result models.EsWaybill, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWaybill{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *WaybillMgr) GetByOptions(opts ...Option) (results []*models.EsWaybill, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWaybill{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *WaybillMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsWaybill, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsWaybill{}).Where(options.query)
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

// GetFromID 通过id获取内容 电子面单id
func (obj *WaybillMgr) GetFromID(id int) (result models.EsWaybill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWaybill{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 电子面单id
func (obj *WaybillMgr) GetBatchFromID(ids []int) (results []*models.EsWaybill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWaybill{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *WaybillMgr) GetFromName(name string) (results []*models.EsWaybill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWaybill{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *WaybillMgr) GetBatchFromName(names []string) (results []*models.EsWaybill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWaybill{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromOpen 通过open获取内容 是否开启
func (obj *WaybillMgr) GetFromOpen(open int) (results []*models.EsWaybill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWaybill{}).Where("`open` = ?", open).Find(&results).Error

	return
}

// GetBatchFromOpen 批量查找 是否开启
func (obj *WaybillMgr) GetBatchFromOpen(opens []int) (results []*models.EsWaybill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWaybill{}).Where("`open` IN (?)", opens).Find(&results).Error

	return
}

// GetFromConfig 通过config获取内容 电子面单配置
func (obj *WaybillMgr) GetFromConfig(config string) (results []*models.EsWaybill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWaybill{}).Where("`config` = ?", config).Find(&results).Error

	return
}

// GetBatchFromConfig 批量查找 电子面单配置
func (obj *WaybillMgr) GetBatchFromConfig(configs []string) (results []*models.EsWaybill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWaybill{}).Where("`config` IN (?)", configs).Find(&results).Error

	return
}

// GetFromBean 通过bean获取内容 beanid
func (obj *WaybillMgr) GetFromBean(bean string) (results []*models.EsWaybill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWaybill{}).Where("`bean` = ?", bean).Find(&results).Error

	return
}

// GetBatchFromBean 批量查找 beanid
func (obj *WaybillMgr) GetBatchFromBean(beans []string) (results []*models.EsWaybill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWaybill{}).Where("`bean` IN (?)", beans).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *WaybillMgr) FetchByPrimaryKey(id int) (result models.EsWaybill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWaybill{}).Where("`id` = ?", id).First(&result).Error

	return
}
