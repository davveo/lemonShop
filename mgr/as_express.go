package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type AsExpressMgr struct {
	*_BaseMgr
}

// NewAsExpressMgr open func
func NewAsExpressMgr(db db.Repo) *AsExpressMgr {
	if db == nil {
		panic(fmt.Errorf("NewAsExpressMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &AsExpressMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_as_express"),
		wdb:       db.GetDbW().Table("es_as_express"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *AsExpressMgr) GetTableName() string {
	return "es_as_express"
}

// Reset 重置gorm会话
func (obj *AsExpressMgr) Reset() *AsExpressMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *AsExpressMgr) Get() (result models.EsAsExpress, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *AsExpressMgr) Gets() (results []*models.EsAsExpress, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *AsExpressMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *AsExpressMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithServiceSn service_sn获取 售后服务单号
func (obj *AsExpressMgr) WithServiceSn(serviceSn string) Option {
	return optionFunc(func(o *options) { o.query["service_sn"] = serviceSn })
}

// WithCourierNumber courier_number获取 物流单号
func (obj *AsExpressMgr) WithCourierNumber(courierNumber string) Option {
	return optionFunc(func(o *options) { o.query["courier_number"] = courierNumber })
}

// WithCourierCompanyID courier_company_id获取 物流公司id
func (obj *AsExpressMgr) WithCourierCompanyID(courierCompanyID int) Option {
	return optionFunc(func(o *options) { o.query["courier_company_id"] = courierCompanyID })
}

// WithCourierCompany courier_company获取 物流公司名称
func (obj *AsExpressMgr) WithCourierCompany(courierCompany string) Option {
	return optionFunc(func(o *options) { o.query["courier_company"] = courierCompany })
}

// WithShipTime ship_time获取 发货时间
func (obj *AsExpressMgr) WithShipTime(shipTime int64) Option {
	return optionFunc(func(o *options) { o.query["ship_time"] = shipTime })
}

// GetByOption 功能选项模式获取
func (obj *AsExpressMgr) GetByOption(opts ...Option) (result models.EsAsExpress, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *AsExpressMgr) GetByOptions(opts ...Option) (results []*models.EsAsExpress, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *AsExpressMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsAsExpress, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Where(options.query)
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
func (obj *AsExpressMgr) GetFromID(id int) (result models.EsAsExpress, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *AsExpressMgr) GetBatchFromID(ids []int) (results []*models.EsAsExpress, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromServiceSn 通过service_sn获取内容 售后服务单号
func (obj *AsExpressMgr) GetFromServiceSn(serviceSn string) (results []*models.EsAsExpress, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Where("`service_sn` = ?", serviceSn).Find(&results).Error

	return
}

// GetBatchFromServiceSn 批量查找 售后服务单号
func (obj *AsExpressMgr) GetBatchFromServiceSn(serviceSns []string) (results []*models.EsAsExpress, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Where("`service_sn` IN (?)", serviceSns).Find(&results).Error

	return
}

// GetFromCourierNumber 通过courier_number获取内容 物流单号
func (obj *AsExpressMgr) GetFromCourierNumber(courierNumber string) (results []*models.EsAsExpress, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Where("`courier_number` = ?", courierNumber).Find(&results).Error

	return
}

// GetBatchFromCourierNumber 批量查找 物流单号
func (obj *AsExpressMgr) GetBatchFromCourierNumber(courierNumbers []string) (results []*models.EsAsExpress, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Where("`courier_number` IN (?)", courierNumbers).Find(&results).Error

	return
}

// GetFromCourierCompanyID 通过courier_company_id获取内容 物流公司id
func (obj *AsExpressMgr) GetFromCourierCompanyID(courierCompanyID int) (results []*models.EsAsExpress, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Where("`courier_company_id` = ?", courierCompanyID).Find(&results).Error

	return
}

// GetBatchFromCourierCompanyID 批量查找 物流公司id
func (obj *AsExpressMgr) GetBatchFromCourierCompanyID(courierCompanyIDs []int) (results []*models.EsAsExpress, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Where("`courier_company_id` IN (?)", courierCompanyIDs).Find(&results).Error

	return
}

// GetFromCourierCompany 通过courier_company获取内容 物流公司名称
func (obj *AsExpressMgr) GetFromCourierCompany(courierCompany string) (results []*models.EsAsExpress, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Where("`courier_company` = ?", courierCompany).Find(&results).Error

	return
}

// GetBatchFromCourierCompany 批量查找 物流公司名称
func (obj *AsExpressMgr) GetBatchFromCourierCompany(courierCompanys []string) (results []*models.EsAsExpress, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Where("`courier_company` IN (?)", courierCompanys).Find(&results).Error

	return
}

// GetFromShipTime 通过ship_time获取内容 发货时间
func (obj *AsExpressMgr) GetFromShipTime(shipTime int64) (results []*models.EsAsExpress, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Where("`ship_time` = ?", shipTime).Find(&results).Error

	return
}

// GetBatchFromShipTime 批量查找 发货时间
func (obj *AsExpressMgr) GetBatchFromShipTime(shipTimes []int64) (results []*models.EsAsExpress, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Where("`ship_time` IN (?)", shipTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *AsExpressMgr) FetchByPrimaryKey(id int) (result models.EsAsExpress, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByEsIndexID  获取多个内容
func (obj *AsExpressMgr) FetchIndexByEsIndexID(id int) (results []*models.EsAsExpress, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// FetchIndexByEsIndexServiceSn  获取多个内容
func (obj *AsExpressMgr) FetchIndexByEsIndexServiceSn(serviceSn string) (results []*models.EsAsExpress, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsExpress{}).Where("`service_sn` = ?", serviceSn).Find(&results).Error

	return
}
