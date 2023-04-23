package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsMemberAddressMgr struct {
	*_BaseMgr
}

// EsMemberAddressMgr open func
func EsMemberAddressMgr(db *gorm.DB) *_EsMemberAddressMgr {
	if db == nil {
		panic(fmt.Errorf("EsMemberAddressMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsMemberAddressMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_member_address"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsMemberAddressMgr) GetTableName() string {
	return "es_member_address"
}

// Reset 重置gorm会话
func (obj *_EsMemberAddressMgr) Reset() *_EsMemberAddressMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsMemberAddressMgr) Get() (result models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsMemberAddressMgr) Gets() (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsMemberAddressMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAddrID addr_id获取 主键ID
func (obj *_EsMemberAddressMgr) WithAddrID(addrID int) Option {
	return optionFunc(func(o *options) { o.query["addr_id"] = addrID })
}

// WithMemberID member_id获取 会员ID
func (obj *_EsMemberAddressMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithName name获取 收货人姓名
func (obj *_EsMemberAddressMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCountry country获取 收货人国籍
func (obj *_EsMemberAddressMgr) WithCountry(country string) Option {
	return optionFunc(func(o *options) { o.query["country"] = country })
}

// WithProvinceID province_id获取 所属省份ID
func (obj *_EsMemberAddressMgr) WithProvinceID(provinceID int) Option {
	return optionFunc(func(o *options) { o.query["province_id"] = provinceID })
}

// WithCityID city_id获取 所属城市ID
func (obj *_EsMemberAddressMgr) WithCityID(cityID int) Option {
	return optionFunc(func(o *options) { o.query["city_id"] = cityID })
}

// WithCountyID county_id获取 所属县(区)ID
func (obj *_EsMemberAddressMgr) WithCountyID(countyID int) Option {
	return optionFunc(func(o *options) { o.query["county_id"] = countyID })
}

// WithTownID town_id获取 所属城镇ID
func (obj *_EsMemberAddressMgr) WithTownID(townID int) Option {
	return optionFunc(func(o *options) { o.query["town_id"] = townID })
}

// WithCounty county获取 所属县(区)名称
func (obj *_EsMemberAddressMgr) WithCounty(county string) Option {
	return optionFunc(func(o *options) { o.query["county"] = county })
}

// WithCity city获取 所属城市名称
func (obj *_EsMemberAddressMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithProvince province获取 所属省份名称
func (obj *_EsMemberAddressMgr) WithProvince(province string) Option {
	return optionFunc(func(o *options) { o.query["province"] = province })
}

// WithTown town获取 所属城镇名称
func (obj *_EsMemberAddressMgr) WithTown(town string) Option {
	return optionFunc(func(o *options) { o.query["town"] = town })
}

// WithAddr addr获取 详细地址
func (obj *_EsMemberAddressMgr) WithAddr(addr string) Option {
	return optionFunc(func(o *options) { o.query["addr"] = addr })
}

// WithTel tel获取 联系电话(一般指座机)
func (obj *_EsMemberAddressMgr) WithTel(tel string) Option {
	return optionFunc(func(o *options) { o.query["tel"] = tel })
}

// WithMobile mobile获取 手机号码
func (obj *_EsMemberAddressMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithDefAddr def_addr获取 是否为默认收货地址
func (obj *_EsMemberAddressMgr) WithDefAddr(defAddr int) Option {
	return optionFunc(func(o *options) { o.query["def_addr"] = defAddr })
}

// WithShipAddressName ship_address_name获取 地址别名
func (obj *_EsMemberAddressMgr) WithShipAddressName(shipAddressName string) Option {
	return optionFunc(func(o *options) { o.query["ship_address_name"] = shipAddressName })
}

// GetByOption 功能选项模式获取
func (obj *_EsMemberAddressMgr) GetByOption(opts ...Option) (result models.EsMemberAddress, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsMemberAddressMgr) GetByOptions(opts ...Option) (results []*models.EsMemberAddress, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsMemberAddressMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsMemberAddress, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where(options.query)
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

// GetFromAddrID 通过addr_id获取内容 主键ID
func (obj *_EsMemberAddressMgr) GetFromAddrID(addrID int) (result models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`addr_id` = ?", addrID).First(&result).Error

	return
}

// GetBatchFromAddrID 批量查找 主键ID
func (obj *_EsMemberAddressMgr) GetBatchFromAddrID(addrIDs []int) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`addr_id` IN (?)", addrIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员ID
func (obj *_EsMemberAddressMgr) GetFromMemberID(memberID int) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员ID
func (obj *_EsMemberAddressMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 收货人姓名
func (obj *_EsMemberAddressMgr) GetFromName(name string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 收货人姓名
func (obj *_EsMemberAddressMgr) GetBatchFromName(names []string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromCountry 通过country获取内容 收货人国籍
func (obj *_EsMemberAddressMgr) GetFromCountry(country string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`country` = ?", country).Find(&results).Error

	return
}

// GetBatchFromCountry 批量查找 收货人国籍
func (obj *_EsMemberAddressMgr) GetBatchFromCountry(countrys []string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`country` IN (?)", countrys).Find(&results).Error

	return
}

// GetFromProvinceID 通过province_id获取内容 所属省份ID
func (obj *_EsMemberAddressMgr) GetFromProvinceID(provinceID int) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`province_id` = ?", provinceID).Find(&results).Error

	return
}

// GetBatchFromProvinceID 批量查找 所属省份ID
func (obj *_EsMemberAddressMgr) GetBatchFromProvinceID(provinceIDs []int) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`province_id` IN (?)", provinceIDs).Find(&results).Error

	return
}

// GetFromCityID 通过city_id获取内容 所属城市ID
func (obj *_EsMemberAddressMgr) GetFromCityID(cityID int) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`city_id` = ?", cityID).Find(&results).Error

	return
}

// GetBatchFromCityID 批量查找 所属城市ID
func (obj *_EsMemberAddressMgr) GetBatchFromCityID(cityIDs []int) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`city_id` IN (?)", cityIDs).Find(&results).Error

	return
}

// GetFromCountyID 通过county_id获取内容 所属县(区)ID
func (obj *_EsMemberAddressMgr) GetFromCountyID(countyID int) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`county_id` = ?", countyID).Find(&results).Error

	return
}

// GetBatchFromCountyID 批量查找 所属县(区)ID
func (obj *_EsMemberAddressMgr) GetBatchFromCountyID(countyIDs []int) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`county_id` IN (?)", countyIDs).Find(&results).Error

	return
}

// GetFromTownID 通过town_id获取内容 所属城镇ID
func (obj *_EsMemberAddressMgr) GetFromTownID(townID int) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`town_id` = ?", townID).Find(&results).Error

	return
}

// GetBatchFromTownID 批量查找 所属城镇ID
func (obj *_EsMemberAddressMgr) GetBatchFromTownID(townIDs []int) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`town_id` IN (?)", townIDs).Find(&results).Error

	return
}

// GetFromCounty 通过county获取内容 所属县(区)名称
func (obj *_EsMemberAddressMgr) GetFromCounty(county string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`county` = ?", county).Find(&results).Error

	return
}

// GetBatchFromCounty 批量查找 所属县(区)名称
func (obj *_EsMemberAddressMgr) GetBatchFromCounty(countys []string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`county` IN (?)", countys).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容 所属城市名称
func (obj *_EsMemberAddressMgr) GetFromCity(city string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`city` = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量查找 所属城市名称
func (obj *_EsMemberAddressMgr) GetBatchFromCity(citys []string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`city` IN (?)", citys).Find(&results).Error

	return
}

// GetFromProvince 通过province获取内容 所属省份名称
func (obj *_EsMemberAddressMgr) GetFromProvince(province string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`province` = ?", province).Find(&results).Error

	return
}

// GetBatchFromProvince 批量查找 所属省份名称
func (obj *_EsMemberAddressMgr) GetBatchFromProvince(provinces []string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`province` IN (?)", provinces).Find(&results).Error

	return
}

// GetFromTown 通过town获取内容 所属城镇名称
func (obj *_EsMemberAddressMgr) GetFromTown(town string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`town` = ?", town).Find(&results).Error

	return
}

// GetBatchFromTown 批量查找 所属城镇名称
func (obj *_EsMemberAddressMgr) GetBatchFromTown(towns []string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`town` IN (?)", towns).Find(&results).Error

	return
}

// GetFromAddr 通过addr获取内容 详细地址
func (obj *_EsMemberAddressMgr) GetFromAddr(addr string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`addr` = ?", addr).Find(&results).Error

	return
}

// GetBatchFromAddr 批量查找 详细地址
func (obj *_EsMemberAddressMgr) GetBatchFromAddr(addrs []string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`addr` IN (?)", addrs).Find(&results).Error

	return
}

// GetFromTel 通过tel获取内容 联系电话(一般指座机)
func (obj *_EsMemberAddressMgr) GetFromTel(tel string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`tel` = ?", tel).Find(&results).Error

	return
}

// GetBatchFromTel 批量查找 联系电话(一般指座机)
func (obj *_EsMemberAddressMgr) GetBatchFromTel(tels []string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`tel` IN (?)", tels).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容 手机号码
func (obj *_EsMemberAddressMgr) GetFromMobile(mobile string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`mobile` = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量查找 手机号码
func (obj *_EsMemberAddressMgr) GetBatchFromMobile(mobiles []string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`mobile` IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromDefAddr 通过def_addr获取内容 是否为默认收货地址
func (obj *_EsMemberAddressMgr) GetFromDefAddr(defAddr int) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`def_addr` = ?", defAddr).Find(&results).Error

	return
}

// GetBatchFromDefAddr 批量查找 是否为默认收货地址
func (obj *_EsMemberAddressMgr) GetBatchFromDefAddr(defAddrs []int) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`def_addr` IN (?)", defAddrs).Find(&results).Error

	return
}

// GetFromShipAddressName 通过ship_address_name获取内容 地址别名
func (obj *_EsMemberAddressMgr) GetFromShipAddressName(shipAddressName string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`ship_address_name` = ?", shipAddressName).Find(&results).Error

	return
}

// GetBatchFromShipAddressName 批量查找 地址别名
func (obj *_EsMemberAddressMgr) GetBatchFromShipAddressName(shipAddressNames []string) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`ship_address_name` IN (?)", shipAddressNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsMemberAddressMgr) FetchByPrimaryKey(addrID int) (result models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`addr_id` = ?", addrID).First(&result).Error

	return
}

// FetchIndexByIndMemAddr  获取多个内容
func (obj *_EsMemberAddressMgr) FetchIndexByIndMemAddr(memberID int, defAddr int) (results []*models.EsMemberAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberAddress{}).Where("`member_id` = ? AND `def_addr` = ?", memberID, defAddr).Find(&results).Error

	return
}
