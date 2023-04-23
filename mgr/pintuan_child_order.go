package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsPintuanChildOrderMgr struct {
	*_BaseMgr
}

// EsPintuanChildOrderMgr open func
func EsPintuanChildOrderMgr(db *gorm.DB) *_EsPintuanChildOrderMgr {
	if db == nil {
		panic(fmt.Errorf("EsPintuanChildOrderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsPintuanChildOrderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_pintuan_child_order"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsPintuanChildOrderMgr) GetTableName() string {
	return "es_pintuan_child_order"
}

// Reset 重置gorm会话
func (obj *_EsPintuanChildOrderMgr) Reset() *_EsPintuanChildOrderMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsPintuanChildOrderMgr) Get() (result models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsPintuanChildOrderMgr) Gets() (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsPintuanChildOrderMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithChildOrderID child_order_id获取 子订单id
func (obj *_EsPintuanChildOrderMgr) WithChildOrderID(childOrderID int) Option {
	return optionFunc(func(o *options) { o.query["child_order_id"] = childOrderID })
}

// WithOrderSn order_sn获取 订单编号
func (obj *_EsPintuanChildOrderMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithMemberID member_id获取 会员id
func (obj *_EsPintuanChildOrderMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithSkuID sku_id获取 会员id
func (obj *_EsPintuanChildOrderMgr) WithSkuID(skuID int) Option {
	return optionFunc(func(o *options) { o.query["sku_id"] = skuID })
}

// WithPintuanID pintuan_id获取 拼团活动id
func (obj *_EsPintuanChildOrderMgr) WithPintuanID(pintuanID int) Option {
	return optionFunc(func(o *options) { o.query["pintuan_id"] = pintuanID })
}

// WithOrderStatus order_status获取 订单状态
func (obj *_EsPintuanChildOrderMgr) WithOrderStatus(orderStatus string) Option {
	return optionFunc(func(o *options) { o.query["order_status"] = orderStatus })
}

// WithOrderID order_id获取 主订单id
func (obj *_EsPintuanChildOrderMgr) WithOrderID(orderID int) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithMemberName member_name获取 买家名称
func (obj *_EsPintuanChildOrderMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithOriginPrice origin_price获取 原价
func (obj *_EsPintuanChildOrderMgr) WithOriginPrice(originPrice float64) Option {
	return optionFunc(func(o *options) { o.query["origin_price"] = originPrice })
}

// WithSalesPrice sales_price获取 拼团价
func (obj *_EsPintuanChildOrderMgr) WithSalesPrice(salesPrice float64) Option {
	return optionFunc(func(o *options) { o.query["sales_price"] = salesPrice })
}

// GetByOption 功能选项模式获取
func (obj *_EsPintuanChildOrderMgr) GetByOption(opts ...Option) (result models.EsPintuanChildOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsPintuanChildOrderMgr) GetByOptions(opts ...Option) (results []*models.EsPintuanChildOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsPintuanChildOrderMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsPintuanChildOrder, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where(options.query)
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

// GetFromChildOrderID 通过child_order_id获取内容 子订单id
func (obj *_EsPintuanChildOrderMgr) GetFromChildOrderID(childOrderID int) (result models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`child_order_id` = ?", childOrderID).First(&result).Error

	return
}

// GetBatchFromChildOrderID 批量查找 子订单id
func (obj *_EsPintuanChildOrderMgr) GetBatchFromChildOrderID(childOrderIDs []int) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`child_order_id` IN (?)", childOrderIDs).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *_EsPintuanChildOrderMgr) GetFromOrderSn(orderSn string) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *_EsPintuanChildOrderMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_EsPintuanChildOrderMgr) GetFromMemberID(memberID int) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_EsPintuanChildOrderMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromSkuID 通过sku_id获取内容 会员id
func (obj *_EsPintuanChildOrderMgr) GetFromSkuID(skuID int) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`sku_id` = ?", skuID).Find(&results).Error

	return
}

// GetBatchFromSkuID 批量查找 会员id
func (obj *_EsPintuanChildOrderMgr) GetBatchFromSkuID(skuIDs []int) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`sku_id` IN (?)", skuIDs).Find(&results).Error

	return
}

// GetFromPintuanID 通过pintuan_id获取内容 拼团活动id
func (obj *_EsPintuanChildOrderMgr) GetFromPintuanID(pintuanID int) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`pintuan_id` = ?", pintuanID).Find(&results).Error

	return
}

// GetBatchFromPintuanID 批量查找 拼团活动id
func (obj *_EsPintuanChildOrderMgr) GetBatchFromPintuanID(pintuanIDs []int) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`pintuan_id` IN (?)", pintuanIDs).Find(&results).Error

	return
}

// GetFromOrderStatus 通过order_status获取内容 订单状态
func (obj *_EsPintuanChildOrderMgr) GetFromOrderStatus(orderStatus string) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`order_status` = ?", orderStatus).Find(&results).Error

	return
}

// GetBatchFromOrderStatus 批量查找 订单状态
func (obj *_EsPintuanChildOrderMgr) GetBatchFromOrderStatus(orderStatuss []string) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`order_status` IN (?)", orderStatuss).Find(&results).Error

	return
}

// GetFromOrderID 通过order_id获取内容 主订单id
func (obj *_EsPintuanChildOrderMgr) GetFromOrderID(orderID int) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`order_id` = ?", orderID).Find(&results).Error

	return
}

// GetBatchFromOrderID 批量查找 主订单id
func (obj *_EsPintuanChildOrderMgr) GetBatchFromOrderID(orderIDs []int) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 买家名称
func (obj *_EsPintuanChildOrderMgr) GetFromMemberName(memberName string) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 买家名称
func (obj *_EsPintuanChildOrderMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromOriginPrice 通过origin_price获取内容 原价
func (obj *_EsPintuanChildOrderMgr) GetFromOriginPrice(originPrice float64) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`origin_price` = ?", originPrice).Find(&results).Error

	return
}

// GetBatchFromOriginPrice 批量查找 原价
func (obj *_EsPintuanChildOrderMgr) GetBatchFromOriginPrice(originPrices []float64) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`origin_price` IN (?)", originPrices).Find(&results).Error

	return
}

// GetFromSalesPrice 通过sales_price获取内容 拼团价
func (obj *_EsPintuanChildOrderMgr) GetFromSalesPrice(salesPrice float64) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`sales_price` = ?", salesPrice).Find(&results).Error

	return
}

// GetBatchFromSalesPrice 批量查找 拼团价
func (obj *_EsPintuanChildOrderMgr) GetBatchFromSalesPrice(salesPrices []float64) (results []*models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`sales_price` IN (?)", salesPrices).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsPintuanChildOrderMgr) FetchByPrimaryKey(childOrderID int) (result models.EsPintuanChildOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuanChildOrder{}).Where("`child_order_id` = ?", childOrderID).First(&result).Error

	return
}
