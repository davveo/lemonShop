package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type ShipTemplateChildMgr struct {
	*_BaseMgr
}

// NewShipTemplateChildMgr open func
func NewShipTemplateChildMgr(db db.Repo) *ShipTemplateChildMgr {
	if db == nil {
		panic(fmt.Errorf("NewShipTemplateChildMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &ShipTemplateChildMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_ship_template_child"), wdb: db.GetDbW().Table("es_ship_template_child"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *ShipTemplateChildMgr) GetTableName() string {
	return "es_ship_template_child"
}

// Reset 重置gorm会话
func (obj *ShipTemplateChildMgr) Reset() *ShipTemplateChildMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *ShipTemplateChildMgr) Get() (result models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *ShipTemplateChildMgr) Gets() (results []*models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *ShipTemplateChildMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *ShipTemplateChildMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTemplateID template_id获取 外键  模板id
func (obj *ShipTemplateChildMgr) WithTemplateID(templateID int) Option {
	return optionFunc(func(o *options) { o.query["template_id"] = templateID })
}

// WithFirstCompany first_company获取 首个（件）
func (obj *ShipTemplateChildMgr) WithFirstCompany(firstCompany int) Option {
	return optionFunc(func(o *options) { o.query["first_company"] = firstCompany })
}

// WithFirstPrice first_price获取 首个（件）价格
func (obj *ShipTemplateChildMgr) WithFirstPrice(firstPrice float64) Option {
	return optionFunc(func(o *options) { o.query["first_price"] = firstPrice })
}

// WithContinuedCompany continued_company获取 续个（件）
func (obj *ShipTemplateChildMgr) WithContinuedCompany(continuedCompany int) Option {
	return optionFunc(func(o *options) { o.query["continued_company"] = continuedCompany })
}

// WithContinuedPrice continued_price获取 续个（件）价格
func (obj *ShipTemplateChildMgr) WithContinuedPrice(continuedPrice float64) Option {
	return optionFunc(func(o *options) { o.query["continued_price"] = continuedPrice })
}

// WithArea area获取 地区json
func (obj *ShipTemplateChildMgr) WithArea(area string) Option {
	return optionFunc(func(o *options) { o.query["area"] = area })
}

// WithAreaID area_id获取 地区id
func (obj *ShipTemplateChildMgr) WithAreaID(areaID string) Option {
	return optionFunc(func(o *options) { o.query["area_id"] = areaID })
}

// GetByOption 功能选项模式获取
func (obj *ShipTemplateChildMgr) GetByOption(opts ...Option) (result models.EsShipTemplateChild, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *ShipTemplateChildMgr) GetByOptions(opts ...Option) (results []*models.EsShipTemplateChild, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *ShipTemplateChildMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsShipTemplateChild, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where(options.query)
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
func (obj *ShipTemplateChildMgr) GetFromID(id int) (result models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *ShipTemplateChildMgr) GetBatchFromID(ids []int) (results []*models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTemplateID 通过template_id获取内容 外键  模板id
func (obj *ShipTemplateChildMgr) GetFromTemplateID(templateID int) (results []*models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where("`template_id` = ?", templateID).Find(&results).Error

	return
}

// GetBatchFromTemplateID 批量查找 外键  模板id
func (obj *ShipTemplateChildMgr) GetBatchFromTemplateID(templateIDs []int) (results []*models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where("`template_id` IN (?)", templateIDs).Find(&results).Error

	return
}

// GetFromFirstCompany 通过first_company获取内容 首个（件）
func (obj *ShipTemplateChildMgr) GetFromFirstCompany(firstCompany int) (results []*models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where("`first_company` = ?", firstCompany).Find(&results).Error

	return
}

// GetBatchFromFirstCompany 批量查找 首个（件）
func (obj *ShipTemplateChildMgr) GetBatchFromFirstCompany(firstCompanys []int) (results []*models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where("`first_company` IN (?)", firstCompanys).Find(&results).Error

	return
}

// GetFromFirstPrice 通过first_price获取内容 首个（件）价格
func (obj *ShipTemplateChildMgr) GetFromFirstPrice(firstPrice float64) (results []*models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where("`first_price` = ?", firstPrice).Find(&results).Error

	return
}

// GetBatchFromFirstPrice 批量查找 首个（件）价格
func (obj *ShipTemplateChildMgr) GetBatchFromFirstPrice(firstPrices []float64) (results []*models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where("`first_price` IN (?)", firstPrices).Find(&results).Error

	return
}

// GetFromContinuedCompany 通过continued_company获取内容 续个（件）
func (obj *ShipTemplateChildMgr) GetFromContinuedCompany(continuedCompany int) (results []*models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where("`continued_company` = ?", continuedCompany).Find(&results).Error

	return
}

// GetBatchFromContinuedCompany 批量查找 续个（件）
func (obj *ShipTemplateChildMgr) GetBatchFromContinuedCompany(continuedCompanys []int) (results []*models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where("`continued_company` IN (?)", continuedCompanys).Find(&results).Error

	return
}

// GetFromContinuedPrice 通过continued_price获取内容 续个（件）价格
func (obj *ShipTemplateChildMgr) GetFromContinuedPrice(continuedPrice float64) (results []*models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where("`continued_price` = ?", continuedPrice).Find(&results).Error

	return
}

// GetBatchFromContinuedPrice 批量查找 续个（件）价格
func (obj *ShipTemplateChildMgr) GetBatchFromContinuedPrice(continuedPrices []float64) (results []*models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where("`continued_price` IN (?)", continuedPrices).Find(&results).Error

	return
}

// GetFromArea 通过area获取内容 地区json
func (obj *ShipTemplateChildMgr) GetFromArea(area string) (results []*models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where("`area` = ?", area).Find(&results).Error

	return
}

// GetBatchFromArea 批量查找 地区json
func (obj *ShipTemplateChildMgr) GetBatchFromArea(areas []string) (results []*models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where("`area` IN (?)", areas).Find(&results).Error

	return
}

// GetFromAreaID 通过area_id获取内容 地区id
func (obj *ShipTemplateChildMgr) GetFromAreaID(areaID string) (results []*models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where("`area_id` = ?", areaID).Find(&results).Error

	return
}

// GetBatchFromAreaID 批量查找 地区id
func (obj *ShipTemplateChildMgr) GetBatchFromAreaID(areaIDs []string) (results []*models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where("`area_id` IN (?)", areaIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *ShipTemplateChildMgr) FetchByPrimaryKey(id int) (result models.EsShipTemplateChild, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShipTemplateChild{}).Where("`id` = ?", id).First(&result).Error

	return
}
