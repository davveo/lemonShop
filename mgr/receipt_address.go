package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsReceiptAddressMgr struct {
	*_BaseMgr
}

// EsReceiptAddressMgr open func
func EsReceiptAddressMgr(db *gorm.DB) *_EsReceiptAddressMgr {
	if db == nil {
		panic(fmt.Errorf("EsReceiptAddressMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsReceiptAddressMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_receipt_address"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsReceiptAddressMgr) GetTableName() string {
	return "es_receipt_address"
}

// Get 获取
func (obj *_EsReceiptAddressMgr) Get() (result models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsReceiptAddressMgr) Gets() (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsReceiptAddressMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *_EsReceiptAddressMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员ID
func (obj *_EsReceiptAddressMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMemberName member_name获取 收票人姓名
func (obj *_EsReceiptAddressMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithMemberMobile member_mobile获取 收票人手机号
func (obj *_EsReceiptAddressMgr) WithMemberMobile(memberMobile string) Option {
	return optionFunc(func(o *options) { o.query["member_mobile"] = memberMobile })
}

// WithProvinceID province_id获取 收票地址-所属省份ID
func (obj *_EsReceiptAddressMgr) WithProvinceID(provinceID int) Option {
	return optionFunc(func(o *options) { o.query["province_id"] = provinceID })
}

// WithCityID city_id获取 收票地址-所属城市ID
func (obj *_EsReceiptAddressMgr) WithCityID(cityID int) Option {
	return optionFunc(func(o *options) { o.query["city_id"] = cityID })
}

// WithCountyID county_id获取 收票地址-所属区县ID
func (obj *_EsReceiptAddressMgr) WithCountyID(countyID int) Option {
	return optionFunc(func(o *options) { o.query["county_id"] = countyID })
}

// WithTownID town_id获取 收票地址-所属乡镇ID
func (obj *_EsReceiptAddressMgr) WithTownID(townID int) Option {
	return optionFunc(func(o *options) { o.query["town_id"] = townID })
}

// WithProvince province获取 收票地址-所属省份
func (obj *_EsReceiptAddressMgr) WithProvince(province string) Option {
	return optionFunc(func(o *options) { o.query["province"] = province })
}

// WithCity city获取 收票地址-所属城市
func (obj *_EsReceiptAddressMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithCounty county获取 收票地址-所属区县
func (obj *_EsReceiptAddressMgr) WithCounty(county string) Option {
	return optionFunc(func(o *options) { o.query["county"] = county })
}

// WithTown town获取 收票地址-所属乡镇
func (obj *_EsReceiptAddressMgr) WithTown(town string) Option {
	return optionFunc(func(o *options) { o.query["town"] = town })
}

// WithDetailAddr detail_addr获取 收票地址-详细地址
func (obj *_EsReceiptAddressMgr) WithDetailAddr(detailAddr string) Option {
	return optionFunc(func(o *options) { o.query["detail_addr"] = detailAddr })
}

// GetByOption 功能选项模式获取
func (obj *_EsReceiptAddressMgr) GetByOption(opts ...Option) (result models.EsReceiptAddress, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsReceiptAddressMgr) GetByOptions(opts ...Option) (results []*models.EsReceiptAddress, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsReceiptAddressMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsReceiptAddress, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where(options.query)
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
func (obj *_EsReceiptAddressMgr) GetFromID(id int) (result models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *_EsReceiptAddressMgr) GetBatchFromID(ids []int) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员ID
func (obj *_EsReceiptAddressMgr) GetFromMemberID(memberID int) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员ID
func (obj *_EsReceiptAddressMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 收票人姓名
func (obj *_EsReceiptAddressMgr) GetFromMemberName(memberName string) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 收票人姓名
func (obj *_EsReceiptAddressMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromMemberMobile 通过member_mobile获取内容 收票人手机号
func (obj *_EsReceiptAddressMgr) GetFromMemberMobile(memberMobile string) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`member_mobile` = ?", memberMobile).Find(&results).Error

	return
}

// GetBatchFromMemberMobile 批量查找 收票人手机号
func (obj *_EsReceiptAddressMgr) GetBatchFromMemberMobile(memberMobiles []string) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`member_mobile` IN (?)", memberMobiles).Find(&results).Error

	return
}

// GetFromProvinceID 通过province_id获取内容 收票地址-所属省份ID
func (obj *_EsReceiptAddressMgr) GetFromProvinceID(provinceID int) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`province_id` = ?", provinceID).Find(&results).Error

	return
}

// GetBatchFromProvinceID 批量查找 收票地址-所属省份ID
func (obj *_EsReceiptAddressMgr) GetBatchFromProvinceID(provinceIDs []int) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`province_id` IN (?)", provinceIDs).Find(&results).Error

	return
}

// GetFromCityID 通过city_id获取内容 收票地址-所属城市ID
func (obj *_EsReceiptAddressMgr) GetFromCityID(cityID int) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`city_id` = ?", cityID).Find(&results).Error

	return
}

// GetBatchFromCityID 批量查找 收票地址-所属城市ID
func (obj *_EsReceiptAddressMgr) GetBatchFromCityID(cityIDs []int) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`city_id` IN (?)", cityIDs).Find(&results).Error

	return
}

// GetFromCountyID 通过county_id获取内容 收票地址-所属区县ID
func (obj *_EsReceiptAddressMgr) GetFromCountyID(countyID int) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`county_id` = ?", countyID).Find(&results).Error

	return
}

// GetBatchFromCountyID 批量查找 收票地址-所属区县ID
func (obj *_EsReceiptAddressMgr) GetBatchFromCountyID(countyIDs []int) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`county_id` IN (?)", countyIDs).Find(&results).Error

	return
}

// GetFromTownID 通过town_id获取内容 收票地址-所属乡镇ID
func (obj *_EsReceiptAddressMgr) GetFromTownID(townID int) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`town_id` = ?", townID).Find(&results).Error

	return
}

// GetBatchFromTownID 批量查找 收票地址-所属乡镇ID
func (obj *_EsReceiptAddressMgr) GetBatchFromTownID(townIDs []int) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`town_id` IN (?)", townIDs).Find(&results).Error

	return
}

// GetFromProvince 通过province获取内容 收票地址-所属省份
func (obj *_EsReceiptAddressMgr) GetFromProvince(province string) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`province` = ?", province).Find(&results).Error

	return
}

// GetBatchFromProvince 批量查找 收票地址-所属省份
func (obj *_EsReceiptAddressMgr) GetBatchFromProvince(provinces []string) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`province` IN (?)", provinces).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容 收票地址-所属城市
func (obj *_EsReceiptAddressMgr) GetFromCity(city string) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`city` = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量查找 收票地址-所属城市
func (obj *_EsReceiptAddressMgr) GetBatchFromCity(citys []string) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`city` IN (?)", citys).Find(&results).Error

	return
}

// GetFromCounty 通过county获取内容 收票地址-所属区县
func (obj *_EsReceiptAddressMgr) GetFromCounty(county string) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`county` = ?", county).Find(&results).Error

	return
}

// GetBatchFromCounty 批量查找 收票地址-所属区县
func (obj *_EsReceiptAddressMgr) GetBatchFromCounty(countys []string) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`county` IN (?)", countys).Find(&results).Error

	return
}

// GetFromTown 通过town获取内容 收票地址-所属乡镇
func (obj *_EsReceiptAddressMgr) GetFromTown(town string) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`town` = ?", town).Find(&results).Error

	return
}

// GetBatchFromTown 批量查找 收票地址-所属乡镇
func (obj *_EsReceiptAddressMgr) GetBatchFromTown(towns []string) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`town` IN (?)", towns).Find(&results).Error

	return
}

// GetFromDetailAddr 通过detail_addr获取内容 收票地址-详细地址
func (obj *_EsReceiptAddressMgr) GetFromDetailAddr(detailAddr string) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`detail_addr` = ?", detailAddr).Find(&results).Error

	return
}

// GetBatchFromDetailAddr 批量查找 收票地址-详细地址
func (obj *_EsReceiptAddressMgr) GetBatchFromDetailAddr(detailAddrs []string) (results []*models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`detail_addr` IN (?)", detailAddrs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsReceiptAddressMgr) FetchByPrimaryKey(id int) (result models.EsReceiptAddress, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptAddress{}).Where("`id` = ?", id).First(&result).Error

	return
}
