package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsMinusMgr struct {
	*_BaseMgr
}

// EsMinusMgr open func
func EsMinusMgr(db *gorm.DB) *_EsMinusMgr {
	if db == nil {
		panic(fmt.Errorf("EsMinusMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsMinusMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_minus"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsMinusMgr) GetTableName() string {
	return "es_minus"
}

// Reset 重置gorm会话
func (obj *_EsMinusMgr) Reset() *_EsMinusMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsMinusMgr) Get() (result models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsMinusMgr) Gets() (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsMinusMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithMinusID minus_id获取 单品立减活动id
func (obj *_EsMinusMgr) WithMinusID(minusID int) Option {
	return optionFunc(func(o *options) { o.query["minus_id"] = minusID })
}

// WithSingleReductionValue single_reduction_value获取 单品立减金额
func (obj *_EsMinusMgr) WithSingleReductionValue(singleReductionValue float64) Option {
	return optionFunc(func(o *options) { o.query["single_reduction_value"] = singleReductionValue })
}

// WithStartTime start_time获取 起始时间
func (obj *_EsMinusMgr) WithStartTime(startTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithStartTimeStr start_time_str获取 起始时间字符串
func (obj *_EsMinusMgr) WithStartTimeStr(startTimeStr string) Option {
	return optionFunc(func(o *options) { o.query["start_time_str"] = startTimeStr })
}

// WithEndTime end_time获取 结束时间
func (obj *_EsMinusMgr) WithEndTime(endTime int64) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithEndTimeStr end_time_str获取 结束时间字符串
func (obj *_EsMinusMgr) WithEndTimeStr(endTimeStr string) Option {
	return optionFunc(func(o *options) { o.query["end_time_str"] = endTimeStr })
}

// WithTitle title获取 单品立减活动标题
func (obj *_EsMinusMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithRangeType range_type获取 是否选择全部商品
func (obj *_EsMinusMgr) WithRangeType(rangeType int) Option {
	return optionFunc(func(o *options) { o.query["range_type"] = rangeType })
}

// WithDisabled disabled获取 是否停用
func (obj *_EsMinusMgr) WithDisabled(disabled int) Option {
	return optionFunc(func(o *options) { o.query["disabled"] = disabled })
}

// WithDescription description获取 描述
func (obj *_EsMinusMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithSellerID seller_id获取 商家id
func (obj *_EsMinusMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// GetByOption 功能选项模式获取
func (obj *_EsMinusMgr) GetByOption(opts ...Option) (result models.EsMinus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsMinusMgr) GetByOptions(opts ...Option) (results []*models.EsMinus, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsMinusMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsMinus, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where(options.query)
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

// GetFromMinusID 通过minus_id获取内容 单品立减活动id
func (obj *_EsMinusMgr) GetFromMinusID(minusID int) (result models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`minus_id` = ?", minusID).First(&result).Error

	return
}

// GetBatchFromMinusID 批量查找 单品立减活动id
func (obj *_EsMinusMgr) GetBatchFromMinusID(minusIDs []int) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`minus_id` IN (?)", minusIDs).Find(&results).Error

	return
}

// GetFromSingleReductionValue 通过single_reduction_value获取内容 单品立减金额
func (obj *_EsMinusMgr) GetFromSingleReductionValue(singleReductionValue float64) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`single_reduction_value` = ?", singleReductionValue).Find(&results).Error

	return
}

// GetBatchFromSingleReductionValue 批量查找 单品立减金额
func (obj *_EsMinusMgr) GetBatchFromSingleReductionValue(singleReductionValues []float64) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`single_reduction_value` IN (?)", singleReductionValues).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容 起始时间
func (obj *_EsMinusMgr) GetFromStartTime(startTime int64) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找 起始时间
func (obj *_EsMinusMgr) GetBatchFromStartTime(startTimes []int64) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromStartTimeStr 通过start_time_str获取内容 起始时间字符串
func (obj *_EsMinusMgr) GetFromStartTimeStr(startTimeStr string) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`start_time_str` = ?", startTimeStr).Find(&results).Error

	return
}

// GetBatchFromStartTimeStr 批量查找 起始时间字符串
func (obj *_EsMinusMgr) GetBatchFromStartTimeStr(startTimeStrs []string) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`start_time_str` IN (?)", startTimeStrs).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容 结束时间
func (obj *_EsMinusMgr) GetFromEndTime(endTime int64) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找 结束时间
func (obj *_EsMinusMgr) GetBatchFromEndTime(endTimes []int64) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromEndTimeStr 通过end_time_str获取内容 结束时间字符串
func (obj *_EsMinusMgr) GetFromEndTimeStr(endTimeStr string) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`end_time_str` = ?", endTimeStr).Find(&results).Error

	return
}

// GetBatchFromEndTimeStr 批量查找 结束时间字符串
func (obj *_EsMinusMgr) GetBatchFromEndTimeStr(endTimeStrs []string) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`end_time_str` IN (?)", endTimeStrs).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 单品立减活动标题
func (obj *_EsMinusMgr) GetFromTitle(title string) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 单品立减活动标题
func (obj *_EsMinusMgr) GetBatchFromTitle(titles []string) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromRangeType 通过range_type获取内容 是否选择全部商品
func (obj *_EsMinusMgr) GetFromRangeType(rangeType int) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`range_type` = ?", rangeType).Find(&results).Error

	return
}

// GetBatchFromRangeType 批量查找 是否选择全部商品
func (obj *_EsMinusMgr) GetBatchFromRangeType(rangeTypes []int) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`range_type` IN (?)", rangeTypes).Find(&results).Error

	return
}

// GetFromDisabled 通过disabled获取内容 是否停用
func (obj *_EsMinusMgr) GetFromDisabled(disabled int) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`disabled` = ?", disabled).Find(&results).Error

	return
}

// GetBatchFromDisabled 批量查找 是否停用
func (obj *_EsMinusMgr) GetBatchFromDisabled(disableds []int) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`disabled` IN (?)", disableds).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容 描述
func (obj *_EsMinusMgr) GetFromDescription(description string) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`description` = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量查找 描述
func (obj *_EsMinusMgr) GetBatchFromDescription(descriptions []string) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`description` IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 商家id
func (obj *_EsMinusMgr) GetFromSellerID(sellerID int) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 商家id
func (obj *_EsMinusMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsMinusMgr) FetchByPrimaryKey(minusID int) (result models.EsMinus, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMinus{}).Where("`minus_id` = ?", minusID).First(&result).Error

	return
}
