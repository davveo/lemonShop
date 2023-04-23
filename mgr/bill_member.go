package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsBillMemberMgr struct {
	*_BaseMgr
}

// EsBillMemberMgr open func
func EsBillMemberMgr(db *gorm.DB) *_EsBillMemberMgr {
	if db == nil {
		panic(fmt.Errorf("EsBillMemberMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsBillMemberMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_bill_member"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsBillMemberMgr) GetTableName() string {
	return "es_bill_member"
}

// Reset 重置gorm会话
func (obj *_EsBillMemberMgr) Reset() *_EsBillMemberMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsBillMemberMgr) Get() (result models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsBillMemberMgr) Gets() (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsBillMemberMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_EsBillMemberMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTotalID total_id获取 总结算单id
func (obj *_EsBillMemberMgr) WithTotalID(totalID int) Option {
	return optionFunc(func(o *options) { o.query["total_id"] = totalID })
}

// WithMemberID member_id获取 会员id
func (obj *_EsBillMemberMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithStartTime start_time获取 开始时间
func (obj *_EsBillMemberMgr) WithStartTime(startTime int) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取 结束时间
func (obj *_EsBillMemberMgr) WithEndTime(endTime int) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithFinalMoney final_money获取 最终结算金额
func (obj *_EsBillMemberMgr) WithFinalMoney(finalMoney float64) Option {
	return optionFunc(func(o *options) { o.query["final_money"] = finalMoney })
}

// WithPushMoney push_money获取 提成金额
func (obj *_EsBillMemberMgr) WithPushMoney(pushMoney float64) Option {
	return optionFunc(func(o *options) { o.query["push_money"] = pushMoney })
}

// WithOrderMoney order_money获取 订单金额
func (obj *_EsBillMemberMgr) WithOrderMoney(orderMoney float64) Option {
	return optionFunc(func(o *options) { o.query["order_money"] = orderMoney })
}

// WithReturnOrderMoney return_order_money获取 订单返还金额
func (obj *_EsBillMemberMgr) WithReturnOrderMoney(returnOrderMoney float64) Option {
	return optionFunc(func(o *options) { o.query["return_order_money"] = returnOrderMoney })
}

// WithReturnPushMoney return_push_money获取 返还提成金额
func (obj *_EsBillMemberMgr) WithReturnPushMoney(returnPushMoney float64) Option {
	return optionFunc(func(o *options) { o.query["return_push_money"] = returnPushMoney })
}

// WithMemberName member_name获取 会员名称
func (obj *_EsBillMemberMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithSn sn获取 编号
func (obj *_EsBillMemberMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithOrderCount order_count获取 订单数
func (obj *_EsBillMemberMgr) WithOrderCount(orderCount int) Option {
	return optionFunc(func(o *options) { o.query["order_count"] = orderCount })
}

// WithReturnOrderCount return_order_count获取 返还订单数
func (obj *_EsBillMemberMgr) WithReturnOrderCount(returnOrderCount int) Option {
	return optionFunc(func(o *options) { o.query["return_order_count"] = returnOrderCount })
}

// GetByOption 功能选项模式获取
func (obj *_EsBillMemberMgr) GetByOption(opts ...Option) (result models.EsBillMember, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsBillMemberMgr) GetByOptions(opts ...Option) (results []*models.EsBillMember, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsBillMemberMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsBillMember, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where(options.query)
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
func (obj *_EsBillMemberMgr) GetFromID(id int) (result models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_EsBillMemberMgr) GetBatchFromID(ids []int) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTotalID 通过total_id获取内容 总结算单id
func (obj *_EsBillMemberMgr) GetFromTotalID(totalID int) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`total_id` = ?", totalID).Find(&results).Error

	return
}

// GetBatchFromTotalID 批量查找 总结算单id
func (obj *_EsBillMemberMgr) GetBatchFromTotalID(totalIDs []int) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`total_id` IN (?)", totalIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_EsBillMemberMgr) GetFromMemberID(memberID int) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_EsBillMemberMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容 开始时间
func (obj *_EsBillMemberMgr) GetFromStartTime(startTime int) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找 开始时间
func (obj *_EsBillMemberMgr) GetBatchFromStartTime(startTimes []int) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容 结束时间
func (obj *_EsBillMemberMgr) GetFromEndTime(endTime int) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找 结束时间
func (obj *_EsBillMemberMgr) GetBatchFromEndTime(endTimes []int) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromFinalMoney 通过final_money获取内容 最终结算金额
func (obj *_EsBillMemberMgr) GetFromFinalMoney(finalMoney float64) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`final_money` = ?", finalMoney).Find(&results).Error

	return
}

// GetBatchFromFinalMoney 批量查找 最终结算金额
func (obj *_EsBillMemberMgr) GetBatchFromFinalMoney(finalMoneys []float64) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`final_money` IN (?)", finalMoneys).Find(&results).Error

	return
}

// GetFromPushMoney 通过push_money获取内容 提成金额
func (obj *_EsBillMemberMgr) GetFromPushMoney(pushMoney float64) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`push_money` = ?", pushMoney).Find(&results).Error

	return
}

// GetBatchFromPushMoney 批量查找 提成金额
func (obj *_EsBillMemberMgr) GetBatchFromPushMoney(pushMoneys []float64) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`push_money` IN (?)", pushMoneys).Find(&results).Error

	return
}

// GetFromOrderMoney 通过order_money获取内容 订单金额
func (obj *_EsBillMemberMgr) GetFromOrderMoney(orderMoney float64) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`order_money` = ?", orderMoney).Find(&results).Error

	return
}

// GetBatchFromOrderMoney 批量查找 订单金额
func (obj *_EsBillMemberMgr) GetBatchFromOrderMoney(orderMoneys []float64) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`order_money` IN (?)", orderMoneys).Find(&results).Error

	return
}

// GetFromReturnOrderMoney 通过return_order_money获取内容 订单返还金额
func (obj *_EsBillMemberMgr) GetFromReturnOrderMoney(returnOrderMoney float64) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`return_order_money` = ?", returnOrderMoney).Find(&results).Error

	return
}

// GetBatchFromReturnOrderMoney 批量查找 订单返还金额
func (obj *_EsBillMemberMgr) GetBatchFromReturnOrderMoney(returnOrderMoneys []float64) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`return_order_money` IN (?)", returnOrderMoneys).Find(&results).Error

	return
}

// GetFromReturnPushMoney 通过return_push_money获取内容 返还提成金额
func (obj *_EsBillMemberMgr) GetFromReturnPushMoney(returnPushMoney float64) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`return_push_money` = ?", returnPushMoney).Find(&results).Error

	return
}

// GetBatchFromReturnPushMoney 批量查找 返还提成金额
func (obj *_EsBillMemberMgr) GetBatchFromReturnPushMoney(returnPushMoneys []float64) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`return_push_money` IN (?)", returnPushMoneys).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 会员名称
func (obj *_EsBillMemberMgr) GetFromMemberName(memberName string) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 会员名称
func (obj *_EsBillMemberMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 编号
func (obj *_EsBillMemberMgr) GetFromSn(sn string) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 编号
func (obj *_EsBillMemberMgr) GetBatchFromSn(sns []string) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromOrderCount 通过order_count获取内容 订单数
func (obj *_EsBillMemberMgr) GetFromOrderCount(orderCount int) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`order_count` = ?", orderCount).Find(&results).Error

	return
}

// GetBatchFromOrderCount 批量查找 订单数
func (obj *_EsBillMemberMgr) GetBatchFromOrderCount(orderCounts []int) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`order_count` IN (?)", orderCounts).Find(&results).Error

	return
}

// GetFromReturnOrderCount 通过return_order_count获取内容 返还订单数
func (obj *_EsBillMemberMgr) GetFromReturnOrderCount(returnOrderCount int) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`return_order_count` = ?", returnOrderCount).Find(&results).Error

	return
}

// GetBatchFromReturnOrderCount 批量查找 返还订单数
func (obj *_EsBillMemberMgr) GetBatchFromReturnOrderCount(returnOrderCounts []int) (results []*models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`return_order_count` IN (?)", returnOrderCounts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsBillMemberMgr) FetchByPrimaryKey(id int) (result models.EsBillMember, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBillMember{}).Where("`id` = ?", id).First(&result).Error

	return
}
