package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsAsChangeMgr struct {
	*_BaseMgr
}

// EsAsChangeMgr open func
func EsAsChangeMgr(db *gorm.DB) *_EsAsChangeMgr {
	if db == nil {
		panic(fmt.Errorf("EsAsChangeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsAsChangeMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_as_change"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsAsChangeMgr) GetTableName() string {
	return "es_as_change"
}

// Get 获取
func (obj *_EsAsChangeMgr) Get() (result models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsAsChangeMgr) Gets() (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsAsChangeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *_EsAsChangeMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithServiceSn service_sn获取 售后服务单号
func (obj *_EsAsChangeMgr) WithServiceSn(serviceSn string) Option {
	return optionFunc(func(o *options) { o.query["service_sn"] = serviceSn })
}

// WithShipName ship_name获取 收货人姓名
func (obj *_EsAsChangeMgr) WithShipName(shipName string) Option {
	return optionFunc(func(o *options) { o.query["ship_name"] = shipName })
}

// WithProvinceID province_id获取 收货地址省份id
func (obj *_EsAsChangeMgr) WithProvinceID(provinceID int) Option {
	return optionFunc(func(o *options) { o.query["province_id"] = provinceID })
}

// WithCityID city_id获取 收货地址城市id
func (obj *_EsAsChangeMgr) WithCityID(cityID int) Option {
	return optionFunc(func(o *options) { o.query["city_id"] = cityID })
}

// WithCountyID county_id获取 收货地址区县id
func (obj *_EsAsChangeMgr) WithCountyID(countyID int) Option {
	return optionFunc(func(o *options) { o.query["county_id"] = countyID })
}

// WithTownID town_id获取 收货地址城镇id
func (obj *_EsAsChangeMgr) WithTownID(townID int) Option {
	return optionFunc(func(o *options) { o.query["town_id"] = townID })
}

// WithProvince province获取 收货地址省份名称
func (obj *_EsAsChangeMgr) WithProvince(province string) Option {
	return optionFunc(func(o *options) { o.query["province"] = province })
}

// WithCity city获取 收货地址城市名称
func (obj *_EsAsChangeMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithCounty county获取 收货地址县(区)名称
func (obj *_EsAsChangeMgr) WithCounty(county string) Option {
	return optionFunc(func(o *options) { o.query["county"] = county })
}

// WithTown town获取 收货地址城镇名称
func (obj *_EsAsChangeMgr) WithTown(town string) Option {
	return optionFunc(func(o *options) { o.query["town"] = town })
}

// WithShipAddr ship_addr获取 收货地址
func (obj *_EsAsChangeMgr) WithShipAddr(shipAddr string) Option {
	return optionFunc(func(o *options) { o.query["ship_addr"] = shipAddr })
}

// WithShipMobile ship_mobile获取 收货人手机号
func (obj *_EsAsChangeMgr) WithShipMobile(shipMobile string) Option {
	return optionFunc(func(o *options) { o.query["ship_mobile"] = shipMobile })
}

// GetByOption 功能选项模式获取
func (obj *_EsAsChangeMgr) GetByOption(opts ...Option) (result models.EsAsChange, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsAsChangeMgr) GetByOptions(opts ...Option) (results []*models.EsAsChange, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsAsChangeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsAsChange, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where(options.query)
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
func (obj *_EsAsChangeMgr) GetFromID(id int) (result models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *_EsAsChangeMgr) GetBatchFromID(ids []int) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromServiceSn 通过service_sn获取内容 售后服务单号
func (obj *_EsAsChangeMgr) GetFromServiceSn(serviceSn string) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`service_sn` = ?", serviceSn).Find(&results).Error

	return
}

// GetBatchFromServiceSn 批量查找 售后服务单号
func (obj *_EsAsChangeMgr) GetBatchFromServiceSn(serviceSns []string) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`service_sn` IN (?)", serviceSns).Find(&results).Error

	return
}

// GetFromShipName 通过ship_name获取内容 收货人姓名
func (obj *_EsAsChangeMgr) GetFromShipName(shipName string) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`ship_name` = ?", shipName).Find(&results).Error

	return
}

// GetBatchFromShipName 批量查找 收货人姓名
func (obj *_EsAsChangeMgr) GetBatchFromShipName(shipNames []string) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`ship_name` IN (?)", shipNames).Find(&results).Error

	return
}

// GetFromProvinceID 通过province_id获取内容 收货地址省份id
func (obj *_EsAsChangeMgr) GetFromProvinceID(provinceID int) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`province_id` = ?", provinceID).Find(&results).Error

	return
}

// GetBatchFromProvinceID 批量查找 收货地址省份id
func (obj *_EsAsChangeMgr) GetBatchFromProvinceID(provinceIDs []int) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`province_id` IN (?)", provinceIDs).Find(&results).Error

	return
}

// GetFromCityID 通过city_id获取内容 收货地址城市id
func (obj *_EsAsChangeMgr) GetFromCityID(cityID int) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`city_id` = ?", cityID).Find(&results).Error

	return
}

// GetBatchFromCityID 批量查找 收货地址城市id
func (obj *_EsAsChangeMgr) GetBatchFromCityID(cityIDs []int) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`city_id` IN (?)", cityIDs).Find(&results).Error

	return
}

// GetFromCountyID 通过county_id获取内容 收货地址区县id
func (obj *_EsAsChangeMgr) GetFromCountyID(countyID int) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`county_id` = ?", countyID).Find(&results).Error

	return
}

// GetBatchFromCountyID 批量查找 收货地址区县id
func (obj *_EsAsChangeMgr) GetBatchFromCountyID(countyIDs []int) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`county_id` IN (?)", countyIDs).Find(&results).Error

	return
}

// GetFromTownID 通过town_id获取内容 收货地址城镇id
func (obj *_EsAsChangeMgr) GetFromTownID(townID int) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`town_id` = ?", townID).Find(&results).Error

	return
}

// GetBatchFromTownID 批量查找 收货地址城镇id
func (obj *_EsAsChangeMgr) GetBatchFromTownID(townIDs []int) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`town_id` IN (?)", townIDs).Find(&results).Error

	return
}

// GetFromProvince 通过province获取内容 收货地址省份名称
func (obj *_EsAsChangeMgr) GetFromProvince(province string) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`province` = ?", province).Find(&results).Error

	return
}

// GetBatchFromProvince 批量查找 收货地址省份名称
func (obj *_EsAsChangeMgr) GetBatchFromProvince(provinces []string) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`province` IN (?)", provinces).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容 收货地址城市名称
func (obj *_EsAsChangeMgr) GetFromCity(city string) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`city` = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量查找 收货地址城市名称
func (obj *_EsAsChangeMgr) GetBatchFromCity(citys []string) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`city` IN (?)", citys).Find(&results).Error

	return
}

// GetFromCounty 通过county获取内容 收货地址县(区)名称
func (obj *_EsAsChangeMgr) GetFromCounty(county string) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`county` = ?", county).Find(&results).Error

	return
}

// GetBatchFromCounty 批量查找 收货地址县(区)名称
func (obj *_EsAsChangeMgr) GetBatchFromCounty(countys []string) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`county` IN (?)", countys).Find(&results).Error

	return
}

// GetFromTown 通过town获取内容 收货地址城镇名称
func (obj *_EsAsChangeMgr) GetFromTown(town string) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`town` = ?", town).Find(&results).Error

	return
}

// GetBatchFromTown 批量查找 收货地址城镇名称
func (obj *_EsAsChangeMgr) GetBatchFromTown(towns []string) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`town` IN (?)", towns).Find(&results).Error

	return
}

// GetFromShipAddr 通过ship_addr获取内容 收货地址
func (obj *_EsAsChangeMgr) GetFromShipAddr(shipAddr string) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`ship_addr` = ?", shipAddr).Find(&results).Error

	return
}

// GetBatchFromShipAddr 批量查找 收货地址
func (obj *_EsAsChangeMgr) GetBatchFromShipAddr(shipAddrs []string) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`ship_addr` IN (?)", shipAddrs).Find(&results).Error

	return
}

// GetFromShipMobile 通过ship_mobile获取内容 收货人手机号
func (obj *_EsAsChangeMgr) GetFromShipMobile(shipMobile string) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`ship_mobile` = ?", shipMobile).Find(&results).Error

	return
}

// GetBatchFromShipMobile 批量查找 收货人手机号
func (obj *_EsAsChangeMgr) GetBatchFromShipMobile(shipMobiles []string) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`ship_mobile` IN (?)", shipMobiles).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsAsChangeMgr) FetchByPrimaryKey(id int) (result models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByEsIndexID  获取多个内容
func (obj *_EsAsChangeMgr) FetchIndexByEsIndexID(id int) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// FetchIndexByEsIndexServiceSn  获取多个内容
func (obj *_EsAsChangeMgr) FetchIndexByEsIndexServiceSn(id int) (results []*models.EsAsChange, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsChange{}).Where("`id` = ?", id).Find(&results).Error

	return
}
