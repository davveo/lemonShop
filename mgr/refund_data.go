package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsSssRefundDataMgr struct {
	*_BaseMgr
}

// EsSssRefundDataMgr open func
func EsSssRefundDataMgr(db *gorm.DB) *_EsSssRefundDataMgr {
	if db == nil {
		panic(fmt.Errorf("EsSssRefundDataMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsSssRefundDataMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_sss_refund_data"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsSssRefundDataMgr) GetTableName() string {
	return "es_sss_refund_data"
}

// Reset 重置gorm会话
func (obj *_EsSssRefundDataMgr) Reset() *_EsSssRefundDataMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsSssRefundDataMgr) Get() (result models.EsSssRefundData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsSssRefundDataMgr) Gets() (results []*models.EsSssRefundData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsSssRefundDataMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_EsSssRefundDataMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSellerID seller_id获取 商家id
func (obj *_EsSssRefundDataMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithOrderSn order_sn获取 订单sn
func (obj *_EsSssRefundDataMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithRefundSn refund_sn获取 售后订单sn
func (obj *_EsSssRefundDataMgr) WithRefundSn(refundSn string) Option {
	return optionFunc(func(o *options) { o.query["refund_sn"] = refundSn })
}

// WithRefundPrice refund_price获取 退还金额
func (obj *_EsSssRefundDataMgr) WithRefundPrice(refundPrice float64) Option {
	return optionFunc(func(o *options) { o.query["refund_price"] = refundPrice })
}

// WithCreateTime create_time获取 创建日期
func (obj *_EsSssRefundDataMgr) WithCreateTime(createTime int) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *_EsSssRefundDataMgr) GetByOption(opts ...Option) (result models.EsSssRefundData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsSssRefundDataMgr) GetByOptions(opts ...Option) (results []*models.EsSssRefundData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsSssRefundDataMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSssRefundData, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).Where(options.query)
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
func (obj *_EsSssRefundDataMgr) GetFromID(id int) (result models.EsSssRefundData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_EsSssRefundDataMgr) GetBatchFromID(ids []int) (results []*models.EsSssRefundData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 商家id
func (obj *_EsSssRefundDataMgr) GetFromSellerID(sellerID int) (results []*models.EsSssRefundData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 商家id
func (obj *_EsSssRefundDataMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsSssRefundData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单sn
func (obj *_EsSssRefundDataMgr) GetFromOrderSn(orderSn string) (results []*models.EsSssRefundData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单sn
func (obj *_EsSssRefundDataMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsSssRefundData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromRefundSn 通过refund_sn获取内容 售后订单sn
func (obj *_EsSssRefundDataMgr) GetFromRefundSn(refundSn string) (results []*models.EsSssRefundData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).Where("`refund_sn` = ?", refundSn).Find(&results).Error

	return
}

// GetBatchFromRefundSn 批量查找 售后订单sn
func (obj *_EsSssRefundDataMgr) GetBatchFromRefundSn(refundSns []string) (results []*models.EsSssRefundData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).Where("`refund_sn` IN (?)", refundSns).Find(&results).Error

	return
}

// GetFromRefundPrice 通过refund_price获取内容 退还金额
func (obj *_EsSssRefundDataMgr) GetFromRefundPrice(refundPrice float64) (results []*models.EsSssRefundData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).Where("`refund_price` = ?", refundPrice).Find(&results).Error

	return
}

// GetBatchFromRefundPrice 批量查找 退还金额
func (obj *_EsSssRefundDataMgr) GetBatchFromRefundPrice(refundPrices []float64) (results []*models.EsSssRefundData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).Where("`refund_price` IN (?)", refundPrices).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建日期
func (obj *_EsSssRefundDataMgr) GetFromCreateTime(createTime int) (results []*models.EsSssRefundData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建日期
func (obj *_EsSssRefundDataMgr) GetBatchFromCreateTime(createTimes []int) (results []*models.EsSssRefundData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsSssRefundDataMgr) FetchByPrimaryKey(id int) (result models.EsSssRefundData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssRefundData{}).Where("`id` = ?", id).First(&result).Error

	return
}
