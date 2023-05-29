package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type BillTotalMgr struct {
	*_BaseMgr
}

// NewBillTotalMgr open func
func NewBillTotalMgr(db db.Repo) *BillTotalMgr {
	if db == nil {
		panic(fmt.Errorf("NewBillTotalMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &BillTotalMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_bill_total"),
		wdb:       db.GetDbW().Table("es_bill_total"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *BillTotalMgr) GetTableName() string {
	return "es_bill_total"
}

// Reset 重置gorm会话
func (obj *BillTotalMgr) Reset() *BillTotalMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *BillTotalMgr) Get() (result models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *BillTotalMgr) Gets() (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *BillTotalMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *BillTotalMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithStartTime start_time获取 开始时间
func (obj *BillTotalMgr) WithStartTime(startTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取 结束时间
func (obj *BillTotalMgr) WithEndTime(endTime int64) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithOrderCount order_count获取 订单数量
func (obj *BillTotalMgr) WithOrderCount(orderCount int64) Option {
	return optionFunc(func(o *options) { o.query["order_count"] = orderCount })
}

// WithReturnOrderCount return_order_count获取 退还订单数量
func (obj *BillTotalMgr) WithReturnOrderCount(returnOrderCount int) Option {
	return optionFunc(func(o *options) { o.query["return_order_count"] = returnOrderCount })
}

// WithFinalMoney final_money获取 结算金额
func (obj *BillTotalMgr) WithFinalMoney(finalMoney float64) Option {
	return optionFunc(func(o *options) { o.query["final_money"] = finalMoney })
}

// WithPushMoney push_money获取 提成金额
func (obj *BillTotalMgr) WithPushMoney(pushMoney float64) Option {
	return optionFunc(func(o *options) { o.query["push_money"] = pushMoney })
}

// WithReturnPushMoney return_push_money获取 退还提成金额
func (obj *BillTotalMgr) WithReturnPushMoney(returnPushMoney float64) Option {
	return optionFunc(func(o *options) { o.query["return_push_money"] = returnPushMoney })
}

// WithOrderMoney order_money获取 订单金额
func (obj *BillTotalMgr) WithOrderMoney(orderMoney float64) Option {
	return optionFunc(func(o *options) { o.query["order_money"] = orderMoney })
}

// WithReturnOrderMoney return_order_money获取 退还订单金额
func (obj *BillTotalMgr) WithReturnOrderMoney(returnOrderMoney float64) Option {
	return optionFunc(func(o *options) { o.query["return_order_money"] = returnOrderMoney })
}

// WithSn sn获取 编号
func (obj *BillTotalMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// GetByOption 功能选项模式获取
func (obj *BillTotalMgr) GetByOption(opts ...Option) (result models.EsBillTotal, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *BillTotalMgr) GetByOptions(opts ...Option) (results []*models.EsBillTotal, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *BillTotalMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsBillTotal, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where(options.query)
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
func (obj *BillTotalMgr) GetFromID(id int64) (result models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *BillTotalMgr) GetBatchFromID(ids []int64) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容 开始时间
func (obj *BillTotalMgr) GetFromStartTime(startTime int64) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找 开始时间
func (obj *BillTotalMgr) GetBatchFromStartTime(startTimes []int64) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容 结束时间
func (obj *BillTotalMgr) GetFromEndTime(endTime int64) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找 结束时间
func (obj *BillTotalMgr) GetBatchFromEndTime(endTimes []int64) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromOrderCount 通过order_count获取内容 订单数量
func (obj *BillTotalMgr) GetFromOrderCount(orderCount int64) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`order_count` = ?", orderCount).Find(&results).Error

	return
}

// GetBatchFromOrderCount 批量查找 订单数量
func (obj *BillTotalMgr) GetBatchFromOrderCount(orderCounts []int64) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`order_count` IN (?)", orderCounts).Find(&results).Error

	return
}

// GetFromReturnOrderCount 通过return_order_count获取内容 退还订单数量
func (obj *BillTotalMgr) GetFromReturnOrderCount(returnOrderCount int) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`return_order_count` = ?", returnOrderCount).Find(&results).Error

	return
}

// GetBatchFromReturnOrderCount 批量查找 退还订单数量
func (obj *BillTotalMgr) GetBatchFromReturnOrderCount(returnOrderCounts []int) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`return_order_count` IN (?)", returnOrderCounts).Find(&results).Error

	return
}

// GetFromFinalMoney 通过final_money获取内容 结算金额
func (obj *BillTotalMgr) GetFromFinalMoney(finalMoney float64) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`final_money` = ?", finalMoney).Find(&results).Error

	return
}

// GetBatchFromFinalMoney 批量查找 结算金额
func (obj *BillTotalMgr) GetBatchFromFinalMoney(finalMoneys []float64) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`final_money` IN (?)", finalMoneys).Find(&results).Error

	return
}

// GetFromPushMoney 通过push_money获取内容 提成金额
func (obj *BillTotalMgr) GetFromPushMoney(pushMoney float64) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`push_money` = ?", pushMoney).Find(&results).Error

	return
}

// GetBatchFromPushMoney 批量查找 提成金额
func (obj *BillTotalMgr) GetBatchFromPushMoney(pushMoneys []float64) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`push_money` IN (?)", pushMoneys).Find(&results).Error

	return
}

// GetFromReturnPushMoney 通过return_push_money获取内容 退还提成金额
func (obj *BillTotalMgr) GetFromReturnPushMoney(returnPushMoney float64) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`return_push_money` = ?", returnPushMoney).Find(&results).Error

	return
}

// GetBatchFromReturnPushMoney 批量查找 退还提成金额
func (obj *BillTotalMgr) GetBatchFromReturnPushMoney(returnPushMoneys []float64) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`return_push_money` IN (?)", returnPushMoneys).Find(&results).Error

	return
}

// GetFromOrderMoney 通过order_money获取内容 订单金额
func (obj *BillTotalMgr) GetFromOrderMoney(orderMoney float64) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`order_money` = ?", orderMoney).Find(&results).Error

	return
}

// GetBatchFromOrderMoney 批量查找 订单金额
func (obj *BillTotalMgr) GetBatchFromOrderMoney(orderMoneys []float64) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`order_money` IN (?)", orderMoneys).Find(&results).Error

	return
}

// GetFromReturnOrderMoney 通过return_order_money获取内容 退还订单金额
func (obj *BillTotalMgr) GetFromReturnOrderMoney(returnOrderMoney float64) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`return_order_money` = ?", returnOrderMoney).Find(&results).Error

	return
}

// GetBatchFromReturnOrderMoney 批量查找 退还订单金额
func (obj *BillTotalMgr) GetBatchFromReturnOrderMoney(returnOrderMoneys []float64) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`return_order_money` IN (?)", returnOrderMoneys).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 编号
func (obj *BillTotalMgr) GetFromSn(sn string) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 编号
func (obj *BillTotalMgr) GetBatchFromSn(sns []string) (results []*models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *BillTotalMgr) FetchByPrimaryKey(id int64) (result models.EsBillTotal, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsBillTotal{}).Where("`id` = ?", id).First(&result).Error

	return
}
