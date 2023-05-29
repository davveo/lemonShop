package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type ConnectSettingMgr struct {
	*_BaseMgr
}

// NewConnectSettingMgr open func
func NewConnectSettingMgr(db db.Repo) *ConnectSettingMgr {
	if db == nil {
		panic(fmt.Errorf("NewConnectSettingMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &ConnectSettingMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_connect_setting"),
		wdb:       db.GetDbW().Table("es_connect_setting"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *ConnectSettingMgr) GetTableName() string {
	return "es_connect_setting"
}

// Reset 重置gorm会话
func (obj *ConnectSettingMgr) Reset() *ConnectSettingMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *ConnectSettingMgr) Get() (result models.EsConnectSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnectSetting{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *ConnectSettingMgr) Gets() (results []*models.EsConnectSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *ConnectSettingMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *ConnectSettingMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithType type获取 授权类型
func (obj *ConnectSettingMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithConfig config获取 参数组
func (obj *ConnectSettingMgr) WithConfig(config string) Option {
	return optionFunc(func(o *options) { o.query["config"] = config })
}

// WithName name获取 标题
func (obj *ConnectSettingMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// GetByOption 功能选项模式获取
func (obj *ConnectSettingMgr) GetByOption(opts ...Option) (result models.EsConnectSetting, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *ConnectSettingMgr) GetByOptions(opts ...Option) (results []*models.EsConnectSetting, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *ConnectSettingMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsConnectSetting, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where(options.query)
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
func (obj *ConnectSettingMgr) GetFromID(id int) (result models.EsConnectSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *ConnectSettingMgr) GetBatchFromID(ids []int) (results []*models.EsConnectSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 授权类型
func (obj *ConnectSettingMgr) GetFromType(_type string) (results []*models.EsConnectSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 授权类型
func (obj *ConnectSettingMgr) GetBatchFromType(_types []string) (results []*models.EsConnectSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromConfig 通过config获取内容 参数组
func (obj *ConnectSettingMgr) GetFromConfig(config string) (results []*models.EsConnectSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where("`config` = ?", config).Find(&results).Error

	return
}

// GetBatchFromConfig 批量查找 参数组
func (obj *ConnectSettingMgr) GetBatchFromConfig(configs []string) (results []*models.EsConnectSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where("`config` IN (?)", configs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 标题
func (obj *ConnectSettingMgr) GetFromName(name string) (results []*models.EsConnectSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 标题
func (obj *ConnectSettingMgr) GetBatchFromName(names []string) (results []*models.EsConnectSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *ConnectSettingMgr) FetchByPrimaryKey(id int) (result models.EsConnectSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnectSetting{}).Where("`id` = ?", id).First(&result).Error

	return
}
