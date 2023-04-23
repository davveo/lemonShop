package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsOrderOutStatusMgr struct {
	*_BaseMgr
}

// EsOrderOutStatusMgr open func
func EsOrderOutStatusMgr(db *gorm.DB) *_EsOrderOutStatusMgr {
	if db == nil {
		panic(fmt.Errorf("EsOrderOutStatusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsOrderOutStatusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_order_out_status"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsOrderOutStatusMgr) GetTableName() string {
	return "es_order_out_status"
}

// Reset 重置gorm会话
func (obj *_EsOrderOutStatusMgr) Reset() *_EsOrderOutStatusMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsOrderOutStatusMgr) Get() (result models.EsOrderOutStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderOutStatus{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsOrderOutStatusMgr) Gets() (results []*models.EsOrderOutStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderOutStatus{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsOrderOutStatusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsOrderOutStatus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithOutID out_id获取 主键
func (obj *_EsOrderOutStatusMgr) WithOutID(outID int) Option {
	return optionFunc(func(o *options) { o.query["out_id"] = outID })
}

// WithOrderSn order_sn获取 订单编号
func (obj *_EsOrderOutStatusMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithOutType out_type获取 出库类型
func (obj *_EsOrderOutStatusMgr) WithOutType(outType string) Option {
	return optionFunc(func(o *options) { o.query["out_type"] = outType })
}

// WithOutStatus out_status获取 出库状态
func (obj *_EsOrderOutStatusMgr) WithOutStatus(outStatus string) Option {
	return optionFunc(func(o *options) { o.query["out_status"] = outStatus })
}

// GetByOption 功能选项模式获取
func (obj *_EsOrderOutStatusMgr) GetByOption(opts ...Option) (result models.EsOrderOutStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderOutStatus{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsOrderOutStatusMgr) GetByOptions(opts ...Option) (results []*models.EsOrderOutStatus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderOutStatus{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsOrderOutStatusMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsOrderOutStatus, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsOrderOutStatus{}).Where(options.query)
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

// GetFromOutID 通过out_id获取内容 主键
func (obj *_EsOrderOutStatusMgr) GetFromOutID(outID int) (result models.EsOrderOutStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderOutStatus{}).Where("`out_id` = ?", outID).First(&result).Error

	return
}

// GetBatchFromOutID 批量查找 主键
func (obj *_EsOrderOutStatusMgr) GetBatchFromOutID(outIDs []int) (results []*models.EsOrderOutStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderOutStatus{}).Where("`out_id` IN (?)", outIDs).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *_EsOrderOutStatusMgr) GetFromOrderSn(orderSn string) (results []*models.EsOrderOutStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderOutStatus{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *_EsOrderOutStatusMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsOrderOutStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderOutStatus{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromOutType 通过out_type获取内容 出库类型
func (obj *_EsOrderOutStatusMgr) GetFromOutType(outType string) (results []*models.EsOrderOutStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderOutStatus{}).Where("`out_type` = ?", outType).Find(&results).Error

	return
}

// GetBatchFromOutType 批量查找 出库类型
func (obj *_EsOrderOutStatusMgr) GetBatchFromOutType(outTypes []string) (results []*models.EsOrderOutStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderOutStatus{}).Where("`out_type` IN (?)", outTypes).Find(&results).Error

	return
}

// GetFromOutStatus 通过out_status获取内容 出库状态
func (obj *_EsOrderOutStatusMgr) GetFromOutStatus(outStatus string) (results []*models.EsOrderOutStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderOutStatus{}).Where("`out_status` = ?", outStatus).Find(&results).Error

	return
}

// GetBatchFromOutStatus 批量查找 出库状态
func (obj *_EsOrderOutStatusMgr) GetBatchFromOutStatus(outStatuss []string) (results []*models.EsOrderOutStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderOutStatus{}).Where("`out_status` IN (?)", outStatuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsOrderOutStatusMgr) FetchByPrimaryKey(outID int) (result models.EsOrderOutStatus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsOrderOutStatus{}).Where("`out_id` = ?", outID).First(&result).Error

	return
}
