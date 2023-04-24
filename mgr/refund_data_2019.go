package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type SssRefundData2019Mgr struct {
	*_BaseMgr
}

// NewSssRefundData2019Mgr open func
func NewSssRefundData2019Mgr(db db.Repo) *SssRefundData2019Mgr {
	if db == nil {
		panic(fmt.Errorf("NewSssRefundData2019Mgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &SssRefundData2019Mgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_sss_refund_data_2019"), wdb: db.GetDbW().Table("es_sss_refund_data_2019"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *SssRefundData2019Mgr) GetTableName() string {
	return "es_sss_refund_data_2019"
}

// Reset 重置gorm会话
func (obj *SssRefundData2019Mgr) Reset() *SssRefundData2019Mgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *SssRefundData2019Mgr) Get() (result models.EsSssRefundData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *SssRefundData2019Mgr) Gets() (results []*models.EsSssRefundData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *SssRefundData2019Mgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *SssRefundData2019Mgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSellerID seller_id获取
func (obj *SssRefundData2019Mgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithRefundSn refund_sn获取
func (obj *SssRefundData2019Mgr) WithRefundSn(refundSn string) Option {
	return optionFunc(func(o *options) { o.query["refund_sn"] = refundSn })
}

// WithOrderSn order_sn获取
func (obj *SssRefundData2019Mgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithRefundPrice refund_price获取
func (obj *SssRefundData2019Mgr) WithRefundPrice(refundPrice float64) Option {
	return optionFunc(func(o *options) { o.query["refund_price"] = refundPrice })
}

// WithCreateTime create_time获取
func (obj *SssRefundData2019Mgr) WithCreateTime(createTime int) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *SssRefundData2019Mgr) GetByOption(opts ...Option) (result models.EsSssRefundData2019, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *SssRefundData2019Mgr) GetByOptions(opts ...Option) (results []*models.EsSssRefundData2019, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *SssRefundData2019Mgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSssRefundData2019, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).Where(options.query)
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
func (obj *SssRefundData2019Mgr) GetFromID(id int) (result models.EsSssRefundData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *SssRefundData2019Mgr) GetBatchFromID(ids []int) (results []*models.EsSssRefundData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容
func (obj *SssRefundData2019Mgr) GetFromSellerID(sellerID int) (results []*models.EsSssRefundData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找
func (obj *SssRefundData2019Mgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsSssRefundData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromRefundSn 通过refund_sn获取内容
func (obj *SssRefundData2019Mgr) GetFromRefundSn(refundSn string) (results []*models.EsSssRefundData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).Where("`refund_sn` = ?", refundSn).Find(&results).Error

	return
}

// GetBatchFromRefundSn 批量查找
func (obj *SssRefundData2019Mgr) GetBatchFromRefundSn(refundSns []string) (results []*models.EsSssRefundData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).Where("`refund_sn` IN (?)", refundSns).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容
func (obj *SssRefundData2019Mgr) GetFromOrderSn(orderSn string) (results []*models.EsSssRefundData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找
func (obj *SssRefundData2019Mgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsSssRefundData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromRefundPrice 通过refund_price获取内容
func (obj *SssRefundData2019Mgr) GetFromRefundPrice(refundPrice float64) (results []*models.EsSssRefundData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).Where("`refund_price` = ?", refundPrice).Find(&results).Error

	return
}

// GetBatchFromRefundPrice 批量查找
func (obj *SssRefundData2019Mgr) GetBatchFromRefundPrice(refundPrices []float64) (results []*models.EsSssRefundData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).Where("`refund_price` IN (?)", refundPrices).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *SssRefundData2019Mgr) GetFromCreateTime(createTime int) (results []*models.EsSssRefundData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *SssRefundData2019Mgr) GetBatchFromCreateTime(createTimes []int) (results []*models.EsSssRefundData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *SssRefundData2019Mgr) FetchByPrimaryKey(id int) (result models.EsSssRefundData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssRefundData2019{}).Where("`id` = ?", id).First(&result).Error

	return
}
