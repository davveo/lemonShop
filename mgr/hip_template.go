package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsShipTemplateMgr struct {
	*_BaseMgr
}

// EsShipTemplateMgr open func
func EsShipTemplateMgr(db *gorm.DB) *_EsShipTemplateMgr {
	if db == nil {
		panic(fmt.Errorf("EsShipTemplateMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsShipTemplateMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_ship_template"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsShipTemplateMgr) GetTableName() string {
	return "es_ship_template"
}

// Reset 重置gorm会话
func (obj *_EsShipTemplateMgr) Reset() *_EsShipTemplateMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsShipTemplateMgr) Get() (result models.EsShipTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShipTemplate{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsShipTemplateMgr) Gets() (results []*models.EsShipTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShipTemplate{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsShipTemplateMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsShipTemplate{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_EsShipTemplateMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSellerID seller_id获取 卖家id
func (obj *_EsShipTemplateMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithName name获取 模板名称
func (obj *_EsShipTemplateMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithType type获取 1 重量算运费 2 计件算运费
func (obj *_EsShipTemplateMgr) WithType(_type int16) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// GetByOption 功能选项模式获取
func (obj *_EsShipTemplateMgr) GetByOption(opts ...Option) (result models.EsShipTemplate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsShipTemplate{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsShipTemplateMgr) GetByOptions(opts ...Option) (results []*models.EsShipTemplate, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsShipTemplate{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsShipTemplateMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsShipTemplate, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsShipTemplate{}).Where(options.query)
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
func (obj *_EsShipTemplateMgr) GetFromID(id int) (result models.EsShipTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShipTemplate{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_EsShipTemplateMgr) GetBatchFromID(ids []int) (results []*models.EsShipTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShipTemplate{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 卖家id
func (obj *_EsShipTemplateMgr) GetFromSellerID(sellerID int) (results []*models.EsShipTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShipTemplate{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 卖家id
func (obj *_EsShipTemplateMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsShipTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShipTemplate{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 模板名称
func (obj *_EsShipTemplateMgr) GetFromName(name string) (results []*models.EsShipTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShipTemplate{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 模板名称
func (obj *_EsShipTemplateMgr) GetBatchFromName(names []string) (results []*models.EsShipTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShipTemplate{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 1 重量算运费 2 计件算运费
func (obj *_EsShipTemplateMgr) GetFromType(_type int16) (results []*models.EsShipTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShipTemplate{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 1 重量算运费 2 计件算运费
func (obj *_EsShipTemplateMgr) GetBatchFromType(_types []int16) (results []*models.EsShipTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShipTemplate{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsShipTemplateMgr) FetchByPrimaryKey(id int) (result models.EsShipTemplate, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShipTemplate{}).Where("`id` = ?", id).First(&result).Error

	return
}
