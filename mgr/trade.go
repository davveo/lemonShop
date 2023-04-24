package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type TradeMgr struct {
	*_BaseMgr
}

// NewTradeMgr open func
func NewTradeMgr(db db.Repo) *TradeMgr {
	if db == nil {
		panic(fmt.Errorf("NewTradeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &TradeMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_trade"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *TradeMgr) GetTableName() string {
	return "es_trade"
}

// Reset 重置gorm会话
func (obj *TradeMgr) Reset() *TradeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *TradeMgr) Get() (result models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *TradeMgr) Gets() (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *TradeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTradeID trade_id获取 trade_id
func (obj *TradeMgr) WithTradeID(tradeID int64) Option {
	return optionFunc(func(o *options) { o.query["trade_id"] = tradeID })
}

// WithTradeSn trade_sn获取 交易编号
func (obj *TradeMgr) WithTradeSn(tradeSn string) Option {
	return optionFunc(func(o *options) { o.query["trade_sn"] = tradeSn })
}

// WithMemberID member_id获取 买家id
func (obj *TradeMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMemberName member_name获取 买家用户名
func (obj *TradeMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithPaymentMethodID payment_method_id获取 支付方式id
func (obj *TradeMgr) WithPaymentMethodID(paymentMethodID string) Option {
	return optionFunc(func(o *options) { o.query["payment_method_id"] = paymentMethodID })
}

// WithPaymentPluginID payment_plugin_id获取 支付插件id
func (obj *TradeMgr) WithPaymentPluginID(paymentPluginID string) Option {
	return optionFunc(func(o *options) { o.query["payment_plugin_id"] = paymentPluginID })
}

// WithPaymentMethodName payment_method_name获取 支付方式名称
func (obj *TradeMgr) WithPaymentMethodName(paymentMethodName string) Option {
	return optionFunc(func(o *options) { o.query["payment_method_name"] = paymentMethodName })
}

// WithPaymentType payment_type获取 支付方式类型
func (obj *TradeMgr) WithPaymentType(paymentType string) Option {
	return optionFunc(func(o *options) { o.query["payment_type"] = paymentType })
}

// WithTotalPrice total_price获取 总价格
func (obj *TradeMgr) WithTotalPrice(totalPrice float64) Option {
	return optionFunc(func(o *options) { o.query["total_price"] = totalPrice })
}

// WithGoodsPrice goods_price获取 商品价格
func (obj *TradeMgr) WithGoodsPrice(goodsPrice float64) Option {
	return optionFunc(func(o *options) { o.query["goods_price"] = goodsPrice })
}

// WithFreightPrice freight_price获取 运费
func (obj *TradeMgr) WithFreightPrice(freightPrice float64) Option {
	return optionFunc(func(o *options) { o.query["freight_price"] = freightPrice })
}

// WithDiscountPrice discount_price获取 优惠的金额
func (obj *TradeMgr) WithDiscountPrice(discountPrice float64) Option {
	return optionFunc(func(o *options) { o.query["discount_price"] = discountPrice })
}

// WithConsigneeID consignee_id获取 收货人id
func (obj *TradeMgr) WithConsigneeID(consigneeID int) Option {
	return optionFunc(func(o *options) { o.query["consignee_id"] = consigneeID })
}

// WithConsigneeName consignee_name获取 收货人姓名
func (obj *TradeMgr) WithConsigneeName(consigneeName string) Option {
	return optionFunc(func(o *options) { o.query["consignee_name"] = consigneeName })
}

// WithConsigneeCountry consignee_country获取 收货国家
func (obj *TradeMgr) WithConsigneeCountry(consigneeCountry string) Option {
	return optionFunc(func(o *options) { o.query["consignee_country"] = consigneeCountry })
}

// WithConsigneeCountryID consignee_country_id获取 收货国家id
func (obj *TradeMgr) WithConsigneeCountryID(consigneeCountryID int) Option {
	return optionFunc(func(o *options) { o.query["consignee_country_id"] = consigneeCountryID })
}

// WithConsigneeProvince consignee_province获取 收货省
func (obj *TradeMgr) WithConsigneeProvince(consigneeProvince string) Option {
	return optionFunc(func(o *options) { o.query["consignee_province"] = consigneeProvince })
}

// WithConsigneeProvinceID consignee_province_id获取 收货省id
func (obj *TradeMgr) WithConsigneeProvinceID(consigneeProvinceID int) Option {
	return optionFunc(func(o *options) { o.query["consignee_province_id"] = consigneeProvinceID })
}

// WithConsigneeCity consignee_city获取 收货市
func (obj *TradeMgr) WithConsigneeCity(consigneeCity string) Option {
	return optionFunc(func(o *options) { o.query["consignee_city"] = consigneeCity })
}

// WithConsigneeCityID consignee_city_id获取 收货市id
func (obj *TradeMgr) WithConsigneeCityID(consigneeCityID int) Option {
	return optionFunc(func(o *options) { o.query["consignee_city_id"] = consigneeCityID })
}

// WithConsigneeCounty consignee_county获取 收货区
func (obj *TradeMgr) WithConsigneeCounty(consigneeCounty string) Option {
	return optionFunc(func(o *options) { o.query["consignee_county"] = consigneeCounty })
}

// WithConsigneeCountyID consignee_county_id获取 收货区id
func (obj *TradeMgr) WithConsigneeCountyID(consigneeCountyID int) Option {
	return optionFunc(func(o *options) { o.query["consignee_county_id"] = consigneeCountyID })
}

// WithConsigneeTown consignee_town获取 收货镇
func (obj *TradeMgr) WithConsigneeTown(consigneeTown string) Option {
	return optionFunc(func(o *options) { o.query["consignee_town"] = consigneeTown })
}

// WithConsigneeTownID consignee_town_id获取 收货镇id
func (obj *TradeMgr) WithConsigneeTownID(consigneeTownID int) Option {
	return optionFunc(func(o *options) { o.query["consignee_town_id"] = consigneeTownID })
}

// WithConsigneeAddress consignee_address获取 收货详细地址
func (obj *TradeMgr) WithConsigneeAddress(consigneeAddress string) Option {
	return optionFunc(func(o *options) { o.query["consignee_address"] = consigneeAddress })
}

// WithConsigneeMobile consignee_mobile获取 收货人手机号
func (obj *TradeMgr) WithConsigneeMobile(consigneeMobile string) Option {
	return optionFunc(func(o *options) { o.query["consignee_mobile"] = consigneeMobile })
}

// WithConsigneeTelephone consignee_telephone获取 收货人电话
func (obj *TradeMgr) WithConsigneeTelephone(consigneeTelephone string) Option {
	return optionFunc(func(o *options) { o.query["consignee_telephone"] = consigneeTelephone })
}

// WithCreateTime create_time获取 交易创建时间
func (obj *TradeMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithOrderJSON order_json获取 订单json(预留，7.0可能废弃)
func (obj *TradeMgr) WithOrderJSON(orderJSON string) Option {
	return optionFunc(func(o *options) { o.query["order_json"] = orderJSON })
}

// WithTradeStatus trade_status获取 订单状态
func (obj *TradeMgr) WithTradeStatus(tradeStatus string) Option {
	return optionFunc(func(o *options) { o.query["trade_status"] = tradeStatus })
}

// GetByOption 功能选项模式获取
func (obj *TradeMgr) GetByOption(opts ...Option) (result models.EsTrade, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *TradeMgr) GetByOptions(opts ...Option) (results []*models.EsTrade, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *TradeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsTrade, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where(options.query)
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

// GetFromTradeID 通过trade_id获取内容 trade_id
func (obj *TradeMgr) GetFromTradeID(tradeID int64) (result models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`trade_id` = ?", tradeID).First(&result).Error

	return
}

// GetBatchFromTradeID 批量查找 trade_id
func (obj *TradeMgr) GetBatchFromTradeID(tradeIDs []int64) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`trade_id` IN (?)", tradeIDs).Find(&results).Error

	return
}

// GetFromTradeSn 通过trade_sn获取内容 交易编号
func (obj *TradeMgr) GetFromTradeSn(tradeSn string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`trade_sn` = ?", tradeSn).Find(&results).Error

	return
}

// GetBatchFromTradeSn 批量查找 交易编号
func (obj *TradeMgr) GetBatchFromTradeSn(tradeSns []string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`trade_sn` IN (?)", tradeSns).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 买家id
func (obj *TradeMgr) GetFromMemberID(memberID int) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 买家id
func (obj *TradeMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 买家用户名
func (obj *TradeMgr) GetFromMemberName(memberName string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 买家用户名
func (obj *TradeMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromPaymentMethodID 通过payment_method_id获取内容 支付方式id
func (obj *TradeMgr) GetFromPaymentMethodID(paymentMethodID string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`payment_method_id` = ?", paymentMethodID).Find(&results).Error

	return
}

// GetBatchFromPaymentMethodID 批量查找 支付方式id
func (obj *TradeMgr) GetBatchFromPaymentMethodID(paymentMethodIDs []string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`payment_method_id` IN (?)", paymentMethodIDs).Find(&results).Error

	return
}

// GetFromPaymentPluginID 通过payment_plugin_id获取内容 支付插件id
func (obj *TradeMgr) GetFromPaymentPluginID(paymentPluginID string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`payment_plugin_id` = ?", paymentPluginID).Find(&results).Error

	return
}

// GetBatchFromPaymentPluginID 批量查找 支付插件id
func (obj *TradeMgr) GetBatchFromPaymentPluginID(paymentPluginIDs []string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`payment_plugin_id` IN (?)", paymentPluginIDs).Find(&results).Error

	return
}

// GetFromPaymentMethodName 通过payment_method_name获取内容 支付方式名称
func (obj *TradeMgr) GetFromPaymentMethodName(paymentMethodName string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`payment_method_name` = ?", paymentMethodName).Find(&results).Error

	return
}

// GetBatchFromPaymentMethodName 批量查找 支付方式名称
func (obj *TradeMgr) GetBatchFromPaymentMethodName(paymentMethodNames []string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`payment_method_name` IN (?)", paymentMethodNames).Find(&results).Error

	return
}

// GetFromPaymentType 通过payment_type获取内容 支付方式类型
func (obj *TradeMgr) GetFromPaymentType(paymentType string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`payment_type` = ?", paymentType).Find(&results).Error

	return
}

// GetBatchFromPaymentType 批量查找 支付方式类型
func (obj *TradeMgr) GetBatchFromPaymentType(paymentTypes []string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`payment_type` IN (?)", paymentTypes).Find(&results).Error

	return
}

// GetFromTotalPrice 通过total_price获取内容 总价格
func (obj *TradeMgr) GetFromTotalPrice(totalPrice float64) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`total_price` = ?", totalPrice).Find(&results).Error

	return
}

// GetBatchFromTotalPrice 批量查找 总价格
func (obj *TradeMgr) GetBatchFromTotalPrice(totalPrices []float64) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`total_price` IN (?)", totalPrices).Find(&results).Error

	return
}

// GetFromGoodsPrice 通过goods_price获取内容 商品价格
func (obj *TradeMgr) GetFromGoodsPrice(goodsPrice float64) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`goods_price` = ?", goodsPrice).Find(&results).Error

	return
}

// GetBatchFromGoodsPrice 批量查找 商品价格
func (obj *TradeMgr) GetBatchFromGoodsPrice(goodsPrices []float64) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`goods_price` IN (?)", goodsPrices).Find(&results).Error

	return
}

// GetFromFreightPrice 通过freight_price获取内容 运费
func (obj *TradeMgr) GetFromFreightPrice(freightPrice float64) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`freight_price` = ?", freightPrice).Find(&results).Error

	return
}

// GetBatchFromFreightPrice 批量查找 运费
func (obj *TradeMgr) GetBatchFromFreightPrice(freightPrices []float64) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`freight_price` IN (?)", freightPrices).Find(&results).Error

	return
}

// GetFromDiscountPrice 通过discount_price获取内容 优惠的金额
func (obj *TradeMgr) GetFromDiscountPrice(discountPrice float64) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`discount_price` = ?", discountPrice).Find(&results).Error

	return
}

// GetBatchFromDiscountPrice 批量查找 优惠的金额
func (obj *TradeMgr) GetBatchFromDiscountPrice(discountPrices []float64) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`discount_price` IN (?)", discountPrices).Find(&results).Error

	return
}

// GetFromConsigneeID 通过consignee_id获取内容 收货人id
func (obj *TradeMgr) GetFromConsigneeID(consigneeID int) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_id` = ?", consigneeID).Find(&results).Error

	return
}

// GetBatchFromConsigneeID 批量查找 收货人id
func (obj *TradeMgr) GetBatchFromConsigneeID(consigneeIDs []int) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_id` IN (?)", consigneeIDs).Find(&results).Error

	return
}

// GetFromConsigneeName 通过consignee_name获取内容 收货人姓名
func (obj *TradeMgr) GetFromConsigneeName(consigneeName string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_name` = ?", consigneeName).Find(&results).Error

	return
}

// GetBatchFromConsigneeName 批量查找 收货人姓名
func (obj *TradeMgr) GetBatchFromConsigneeName(consigneeNames []string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_name` IN (?)", consigneeNames).Find(&results).Error

	return
}

// GetFromConsigneeCountry 通过consignee_country获取内容 收货国家
func (obj *TradeMgr) GetFromConsigneeCountry(consigneeCountry string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_country` = ?", consigneeCountry).Find(&results).Error

	return
}

// GetBatchFromConsigneeCountry 批量查找 收货国家
func (obj *TradeMgr) GetBatchFromConsigneeCountry(consigneeCountrys []string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_country` IN (?)", consigneeCountrys).Find(&results).Error

	return
}

// GetFromConsigneeCountryID 通过consignee_country_id获取内容 收货国家id
func (obj *TradeMgr) GetFromConsigneeCountryID(consigneeCountryID int) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_country_id` = ?", consigneeCountryID).Find(&results).Error

	return
}

// GetBatchFromConsigneeCountryID 批量查找 收货国家id
func (obj *TradeMgr) GetBatchFromConsigneeCountryID(consigneeCountryIDs []int) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_country_id` IN (?)", consigneeCountryIDs).Find(&results).Error

	return
}

// GetFromConsigneeProvince 通过consignee_province获取内容 收货省
func (obj *TradeMgr) GetFromConsigneeProvince(consigneeProvince string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_province` = ?", consigneeProvince).Find(&results).Error

	return
}

// GetBatchFromConsigneeProvince 批量查找 收货省
func (obj *TradeMgr) GetBatchFromConsigneeProvince(consigneeProvinces []string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_province` IN (?)", consigneeProvinces).Find(&results).Error

	return
}

// GetFromConsigneeProvinceID 通过consignee_province_id获取内容 收货省id
func (obj *TradeMgr) GetFromConsigneeProvinceID(consigneeProvinceID int) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_province_id` = ?", consigneeProvinceID).Find(&results).Error

	return
}

// GetBatchFromConsigneeProvinceID 批量查找 收货省id
func (obj *TradeMgr) GetBatchFromConsigneeProvinceID(consigneeProvinceIDs []int) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_province_id` IN (?)", consigneeProvinceIDs).Find(&results).Error

	return
}

// GetFromConsigneeCity 通过consignee_city获取内容 收货市
func (obj *TradeMgr) GetFromConsigneeCity(consigneeCity string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_city` = ?", consigneeCity).Find(&results).Error

	return
}

// GetBatchFromConsigneeCity 批量查找 收货市
func (obj *TradeMgr) GetBatchFromConsigneeCity(consigneeCitys []string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_city` IN (?)", consigneeCitys).Find(&results).Error

	return
}

// GetFromConsigneeCityID 通过consignee_city_id获取内容 收货市id
func (obj *TradeMgr) GetFromConsigneeCityID(consigneeCityID int) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_city_id` = ?", consigneeCityID).Find(&results).Error

	return
}

// GetBatchFromConsigneeCityID 批量查找 收货市id
func (obj *TradeMgr) GetBatchFromConsigneeCityID(consigneeCityIDs []int) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_city_id` IN (?)", consigneeCityIDs).Find(&results).Error

	return
}

// GetFromConsigneeCounty 通过consignee_county获取内容 收货区
func (obj *TradeMgr) GetFromConsigneeCounty(consigneeCounty string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_county` = ?", consigneeCounty).Find(&results).Error

	return
}

// GetBatchFromConsigneeCounty 批量查找 收货区
func (obj *TradeMgr) GetBatchFromConsigneeCounty(consigneeCountys []string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_county` IN (?)", consigneeCountys).Find(&results).Error

	return
}

// GetFromConsigneeCountyID 通过consignee_county_id获取内容 收货区id
func (obj *TradeMgr) GetFromConsigneeCountyID(consigneeCountyID int) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_county_id` = ?", consigneeCountyID).Find(&results).Error

	return
}

// GetBatchFromConsigneeCountyID 批量查找 收货区id
func (obj *TradeMgr) GetBatchFromConsigneeCountyID(consigneeCountyIDs []int) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_county_id` IN (?)", consigneeCountyIDs).Find(&results).Error

	return
}

// GetFromConsigneeTown 通过consignee_town获取内容 收货镇
func (obj *TradeMgr) GetFromConsigneeTown(consigneeTown string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_town` = ?", consigneeTown).Find(&results).Error

	return
}

// GetBatchFromConsigneeTown 批量查找 收货镇
func (obj *TradeMgr) GetBatchFromConsigneeTown(consigneeTowns []string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_town` IN (?)", consigneeTowns).Find(&results).Error

	return
}

// GetFromConsigneeTownID 通过consignee_town_id获取内容 收货镇id
func (obj *TradeMgr) GetFromConsigneeTownID(consigneeTownID int) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_town_id` = ?", consigneeTownID).Find(&results).Error

	return
}

// GetBatchFromConsigneeTownID 批量查找 收货镇id
func (obj *TradeMgr) GetBatchFromConsigneeTownID(consigneeTownIDs []int) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_town_id` IN (?)", consigneeTownIDs).Find(&results).Error

	return
}

// GetFromConsigneeAddress 通过consignee_address获取内容 收货详细地址
func (obj *TradeMgr) GetFromConsigneeAddress(consigneeAddress string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_address` = ?", consigneeAddress).Find(&results).Error

	return
}

// GetBatchFromConsigneeAddress 批量查找 收货详细地址
func (obj *TradeMgr) GetBatchFromConsigneeAddress(consigneeAddresss []string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_address` IN (?)", consigneeAddresss).Find(&results).Error

	return
}

// GetFromConsigneeMobile 通过consignee_mobile获取内容 收货人手机号
func (obj *TradeMgr) GetFromConsigneeMobile(consigneeMobile string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_mobile` = ?", consigneeMobile).Find(&results).Error

	return
}

// GetBatchFromConsigneeMobile 批量查找 收货人手机号
func (obj *TradeMgr) GetBatchFromConsigneeMobile(consigneeMobiles []string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_mobile` IN (?)", consigneeMobiles).Find(&results).Error

	return
}

// GetFromConsigneeTelephone 通过consignee_telephone获取内容 收货人电话
func (obj *TradeMgr) GetFromConsigneeTelephone(consigneeTelephone string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_telephone` = ?", consigneeTelephone).Find(&results).Error

	return
}

// GetBatchFromConsigneeTelephone 批量查找 收货人电话
func (obj *TradeMgr) GetBatchFromConsigneeTelephone(consigneeTelephones []string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`consignee_telephone` IN (?)", consigneeTelephones).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 交易创建时间
func (obj *TradeMgr) GetFromCreateTime(createTime int64) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 交易创建时间
func (obj *TradeMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromOrderJSON 通过order_json获取内容 订单json(预留，7.0可能废弃)
func (obj *TradeMgr) GetFromOrderJSON(orderJSON string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`order_json` = ?", orderJSON).Find(&results).Error

	return
}

// GetBatchFromOrderJSON 批量查找 订单json(预留，7.0可能废弃)
func (obj *TradeMgr) GetBatchFromOrderJSON(orderJSONs []string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`order_json` IN (?)", orderJSONs).Find(&results).Error

	return
}

// GetFromTradeStatus 通过trade_status获取内容 订单状态
func (obj *TradeMgr) GetFromTradeStatus(tradeStatus string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`trade_status` = ?", tradeStatus).Find(&results).Error

	return
}

// GetBatchFromTradeStatus 批量查找 订单状态
func (obj *TradeMgr) GetBatchFromTradeStatus(tradeStatuss []string) (results []*models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`trade_status` IN (?)", tradeStatuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *TradeMgr) FetchByPrimaryKey(tradeID int64) (result models.EsTrade, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTrade{}).Where("`trade_id` = ?", tradeID).First(&result).Error

	return
}
