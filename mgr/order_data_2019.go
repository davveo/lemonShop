package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type SssOrderData2019Mgr struct {
	*_BaseMgr
}

// NewSssOrderData2019Mgr open func
func NewSssOrderData2019Mgr(db db.Repo) *SssOrderData2019Mgr {
	if db == nil {
		panic(fmt.Errorf("NewSssOrderData2019Mgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &SssOrderData2019Mgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_sss_order_data_2019"), wdb: db.GetDbW().Table("es_sss_order_data_2019"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *SssOrderData2019Mgr) GetTableName() string {
	return "es_sss_order_data_2019"
}

// Reset 重置gorm会话
func (obj *SssOrderData2019Mgr) Reset() *SssOrderData2019Mgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *SssOrderData2019Mgr) Get() (result models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *SssOrderData2019Mgr) Gets() (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *SssOrderData2019Mgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *SssOrderData2019Mgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSn sn获取
func (obj *SssOrderData2019Mgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithBuyerID buyer_id获取
func (obj *SssOrderData2019Mgr) WithBuyerID(buyerID int) Option {
	return optionFunc(func(o *options) { o.query["buyer_id"] = buyerID })
}

// WithBuyerName buyer_name获取
func (obj *SssOrderData2019Mgr) WithBuyerName(buyerName string) Option {
	return optionFunc(func(o *options) { o.query["buyer_name"] = buyerName })
}

// WithSellerID seller_id获取
func (obj *SssOrderData2019Mgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithSellerName seller_name获取
func (obj *SssOrderData2019Mgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithOrderStatus order_status获取
func (obj *SssOrderData2019Mgr) WithOrderStatus(orderStatus string) Option {
	return optionFunc(func(o *options) { o.query["order_status"] = orderStatus })
}

// WithPayStatus pay_status获取
func (obj *SssOrderData2019Mgr) WithPayStatus(payStatus string) Option {
	return optionFunc(func(o *options) { o.query["pay_status"] = payStatus })
}

// WithOrderPrice order_price获取
func (obj *SssOrderData2019Mgr) WithOrderPrice(orderPrice float64) Option {
	return optionFunc(func(o *options) { o.query["order_price"] = orderPrice })
}

// WithGoodsNum goods_num获取
func (obj *SssOrderData2019Mgr) WithGoodsNum(goodsNum int) Option {
	return optionFunc(func(o *options) { o.query["goods_num"] = goodsNum })
}

// WithShipProvinceID ship_province_id获取
func (obj *SssOrderData2019Mgr) WithShipProvinceID(shipProvinceID int) Option {
	return optionFunc(func(o *options) { o.query["ship_province_id"] = shipProvinceID })
}

// WithShipCityID ship_city_id获取
func (obj *SssOrderData2019Mgr) WithShipCityID(shipCityID int) Option {
	return optionFunc(func(o *options) { o.query["ship_city_id"] = shipCityID })
}

// WithCreateTime create_time获取
func (obj *SssOrderData2019Mgr) WithCreateTime(createTime int) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *SssOrderData2019Mgr) GetByOption(opts ...Option) (result models.EsSssOrderData2019, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *SssOrderData2019Mgr) GetByOptions(opts ...Option) (results []*models.EsSssOrderData2019, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *SssOrderData2019Mgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSssOrderData2019, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *SssOrderData2019Mgr) GetFromID(id int) (result models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *SssOrderData2019Mgr) GetBatchFromID(ids []int) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容
func (obj *SssOrderData2019Mgr) GetFromSn(sn string) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找
func (obj *SssOrderData2019Mgr) GetBatchFromSn(sns []string) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromBuyerID 通过buyer_id获取内容
func (obj *SssOrderData2019Mgr) GetFromBuyerID(buyerID int) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`buyer_id` = ?", buyerID).Find(&results).Error

	return
}

// GetBatchFromBuyerID 批量查找
func (obj *SssOrderData2019Mgr) GetBatchFromBuyerID(buyerIDs []int) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`buyer_id` IN (?)", buyerIDs).Find(&results).Error

	return
}

// GetFromBuyerName 通过buyer_name获取内容
func (obj *SssOrderData2019Mgr) GetFromBuyerName(buyerName string) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`buyer_name` = ?", buyerName).Find(&results).Error

	return
}

// GetBatchFromBuyerName 批量查找
func (obj *SssOrderData2019Mgr) GetBatchFromBuyerName(buyerNames []string) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`buyer_name` IN (?)", buyerNames).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容
func (obj *SssOrderData2019Mgr) GetFromSellerID(sellerID int) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找
func (obj *SssOrderData2019Mgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容
func (obj *SssOrderData2019Mgr) GetFromSellerName(sellerName string) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找
func (obj *SssOrderData2019Mgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromOrderStatus 通过order_status获取内容
func (obj *SssOrderData2019Mgr) GetFromOrderStatus(orderStatus string) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`order_status` = ?", orderStatus).Find(&results).Error

	return
}

// GetBatchFromOrderStatus 批量查找
func (obj *SssOrderData2019Mgr) GetBatchFromOrderStatus(orderStatuss []string) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`order_status` IN (?)", orderStatuss).Find(&results).Error

	return
}

// GetFromPayStatus 通过pay_status获取内容
func (obj *SssOrderData2019Mgr) GetFromPayStatus(payStatus string) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`pay_status` = ?", payStatus).Find(&results).Error

	return
}

// GetBatchFromPayStatus 批量查找
func (obj *SssOrderData2019Mgr) GetBatchFromPayStatus(payStatuss []string) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`pay_status` IN (?)", payStatuss).Find(&results).Error

	return
}

// GetFromOrderPrice 通过order_price获取内容
func (obj *SssOrderData2019Mgr) GetFromOrderPrice(orderPrice float64) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`order_price` = ?", orderPrice).Find(&results).Error

	return
}

// GetBatchFromOrderPrice 批量查找
func (obj *SssOrderData2019Mgr) GetBatchFromOrderPrice(orderPrices []float64) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`order_price` IN (?)", orderPrices).Find(&results).Error

	return
}

// GetFromGoodsNum 通过goods_num获取内容
func (obj *SssOrderData2019Mgr) GetFromGoodsNum(goodsNum int) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`goods_num` = ?", goodsNum).Find(&results).Error

	return
}

// GetBatchFromGoodsNum 批量查找
func (obj *SssOrderData2019Mgr) GetBatchFromGoodsNum(goodsNums []int) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`goods_num` IN (?)", goodsNums).Find(&results).Error

	return
}

// GetFromShipProvinceID 通过ship_province_id获取内容
func (obj *SssOrderData2019Mgr) GetFromShipProvinceID(shipProvinceID int) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`ship_province_id` = ?", shipProvinceID).Find(&results).Error

	return
}

// GetBatchFromShipProvinceID 批量查找
func (obj *SssOrderData2019Mgr) GetBatchFromShipProvinceID(shipProvinceIDs []int) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`ship_province_id` IN (?)", shipProvinceIDs).Find(&results).Error

	return
}

// GetFromShipCityID 通过ship_city_id获取内容
func (obj *SssOrderData2019Mgr) GetFromShipCityID(shipCityID int) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`ship_city_id` = ?", shipCityID).Find(&results).Error

	return
}

// GetBatchFromShipCityID 批量查找
func (obj *SssOrderData2019Mgr) GetBatchFromShipCityID(shipCityIDs []int) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`ship_city_id` IN (?)", shipCityIDs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *SssOrderData2019Mgr) GetFromCreateTime(createTime int) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *SssOrderData2019Mgr) GetBatchFromCreateTime(createTimes []int) (results []*models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *SssOrderData2019Mgr) FetchByPrimaryKey(id int) (result models.EsSssOrderData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderData2019{}).Where("`id` = ?", id).First(&result).Error

	return
}
