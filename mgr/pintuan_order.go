package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type PintuanOrderMgr struct {
	*_BaseMgr
}

// NewPintuanOrderMgr open func
func NewPintuanOrderMgr(db db.Repo) *PintuanOrderMgr {
	if db == nil {
		panic(fmt.Errorf("NewPintuanOrderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &PintuanOrderMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_pintuan_order"), wdb: db.GetDbW().Table("es_pintuan_order"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *PintuanOrderMgr) GetTableName() string {
	return "es_pintuan_order"
}

// Reset 重置gorm会话
func (obj *PintuanOrderMgr) Reset() *PintuanOrderMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *PintuanOrderMgr) Get() (result models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *PintuanOrderMgr) Gets() (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *PintuanOrderMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithOrderID order_id获取 订单id
func (obj *PintuanOrderMgr) WithOrderID(orderID int) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithPintuanID pintuan_id获取 拼团id
func (obj *PintuanOrderMgr) WithPintuanID(pintuanID int) Option {
	return optionFunc(func(o *options) { o.query["pintuan_id"] = pintuanID })
}

// WithEndTime end_time获取 结束时间
func (obj *PintuanOrderMgr) WithEndTime(endTime int64) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithSkuID sku_id获取 sku_id
func (obj *PintuanOrderMgr) WithSkuID(skuID int) Option {
	return optionFunc(func(o *options) { o.query["sku_id"] = skuID })
}

// WithRequiredNum required_num获取 成团人数
func (obj *PintuanOrderMgr) WithRequiredNum(requiredNum int) Option {
	return optionFunc(func(o *options) { o.query["required_num"] = requiredNum })
}

// WithOfferedNum offered_num获取 已参团人数
func (obj *PintuanOrderMgr) WithOfferedNum(offeredNum int) Option {
	return optionFunc(func(o *options) { o.query["offered_num"] = offeredNum })
}

// WithOfferedPersons offered_persons获取 参团人
func (obj *PintuanOrderMgr) WithOfferedPersons(offeredPersons string) Option {
	return optionFunc(func(o *options) { o.query["offered_persons"] = offeredPersons })
}

// WithOrderStatus order_status获取 订单状态
func (obj *PintuanOrderMgr) WithOrderStatus(orderStatus string) Option {
	return optionFunc(func(o *options) { o.query["order_status"] = orderStatus })
}

// WithGoodsID goods_id获取 商品id
func (obj *PintuanOrderMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithThumbnail thumbnail获取 商品缩略图
func (obj *PintuanOrderMgr) WithThumbnail(thumbnail string) Option {
	return optionFunc(func(o *options) { o.query["thumbnail"] = thumbnail })
}

// WithGoodsName goods_name获取 商品名称
func (obj *PintuanOrderMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// GetByOption 功能选项模式获取
func (obj *PintuanOrderMgr) GetByOption(opts ...Option) (result models.EsPintuanOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *PintuanOrderMgr) GetByOptions(opts ...Option) (results []*models.EsPintuanOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *PintuanOrderMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsPintuanOrder, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where(options.query)
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

// GetFromOrderID 通过order_id获取内容 订单id
func (obj *PintuanOrderMgr) GetFromOrderID(orderID int) (result models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`order_id` = ?", orderID).First(&result).Error

	return
}

// GetBatchFromOrderID 批量查找 订单id
func (obj *PintuanOrderMgr) GetBatchFromOrderID(orderIDs []int) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromPintuanID 通过pintuan_id获取内容 拼团id
func (obj *PintuanOrderMgr) GetFromPintuanID(pintuanID int) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`pintuan_id` = ?", pintuanID).Find(&results).Error

	return
}

// GetBatchFromPintuanID 批量查找 拼团id
func (obj *PintuanOrderMgr) GetBatchFromPintuanID(pintuanIDs []int) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`pintuan_id` IN (?)", pintuanIDs).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容 结束时间
func (obj *PintuanOrderMgr) GetFromEndTime(endTime int64) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找 结束时间
func (obj *PintuanOrderMgr) GetBatchFromEndTime(endTimes []int64) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromSkuID 通过sku_id获取内容 sku_id
func (obj *PintuanOrderMgr) GetFromSkuID(skuID int) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`sku_id` = ?", skuID).Find(&results).Error

	return
}

// GetBatchFromSkuID 批量查找 sku_id
func (obj *PintuanOrderMgr) GetBatchFromSkuID(skuIDs []int) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`sku_id` IN (?)", skuIDs).Find(&results).Error

	return
}

// GetFromRequiredNum 通过required_num获取内容 成团人数
func (obj *PintuanOrderMgr) GetFromRequiredNum(requiredNum int) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`required_num` = ?", requiredNum).Find(&results).Error

	return
}

// GetBatchFromRequiredNum 批量查找 成团人数
func (obj *PintuanOrderMgr) GetBatchFromRequiredNum(requiredNums []int) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`required_num` IN (?)", requiredNums).Find(&results).Error

	return
}

// GetFromOfferedNum 通过offered_num获取内容 已参团人数
func (obj *PintuanOrderMgr) GetFromOfferedNum(offeredNum int) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`offered_num` = ?", offeredNum).Find(&results).Error

	return
}

// GetBatchFromOfferedNum 批量查找 已参团人数
func (obj *PintuanOrderMgr) GetBatchFromOfferedNum(offeredNums []int) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`offered_num` IN (?)", offeredNums).Find(&results).Error

	return
}

// GetFromOfferedPersons 通过offered_persons获取内容 参团人
func (obj *PintuanOrderMgr) GetFromOfferedPersons(offeredPersons string) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`offered_persons` = ?", offeredPersons).Find(&results).Error

	return
}

// GetBatchFromOfferedPersons 批量查找 参团人
func (obj *PintuanOrderMgr) GetBatchFromOfferedPersons(offeredPersonss []string) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`offered_persons` IN (?)", offeredPersonss).Find(&results).Error

	return
}

// GetFromOrderStatus 通过order_status获取内容 订单状态
func (obj *PintuanOrderMgr) GetFromOrderStatus(orderStatus string) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`order_status` = ?", orderStatus).Find(&results).Error

	return
}

// GetBatchFromOrderStatus 批量查找 订单状态
func (obj *PintuanOrderMgr) GetBatchFromOrderStatus(orderStatuss []string) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`order_status` IN (?)", orderStatuss).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品id
func (obj *PintuanOrderMgr) GetFromGoodsID(goodsID int) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品id
func (obj *PintuanOrderMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromThumbnail 通过thumbnail获取内容 商品缩略图
func (obj *PintuanOrderMgr) GetFromThumbnail(thumbnail string) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`thumbnail` = ?", thumbnail).Find(&results).Error

	return
}

// GetBatchFromThumbnail 批量查找 商品缩略图
func (obj *PintuanOrderMgr) GetBatchFromThumbnail(thumbnails []string) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`thumbnail` IN (?)", thumbnails).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *PintuanOrderMgr) GetFromGoodsName(goodsName string) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *PintuanOrderMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *PintuanOrderMgr) FetchByPrimaryKey(orderID int) (result models.EsPintuanOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuanOrder{}).Where("`order_id` = ?", orderID).First(&result).Error

	return
}
