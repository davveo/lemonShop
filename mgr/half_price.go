package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsHalfPriceMgr struct {
	*_BaseMgr
}

// EsHalfPriceMgr open func
func EsHalfPriceMgr(db *gorm.DB) *_EsHalfPriceMgr {
	if db == nil {
		panic(fmt.Errorf("EsHalfPriceMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsHalfPriceMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_half_price"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsHalfPriceMgr) GetTableName() string {
	return "es_half_price"
}

// Reset 重置gorm会话
func (obj *_EsHalfPriceMgr) Reset() *_EsHalfPriceMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsHalfPriceMgr) Get() (result models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsHalfPriceMgr) Gets() (results []*models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsHalfPriceMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithHpID hp_id获取 第二件半价活动ID
func (obj *_EsHalfPriceMgr) WithHpID(hpID int) Option {
	return optionFunc(func(o *options) { o.query["hp_id"] = hpID })
}

// WithStartTime start_time获取 活动开始时间
func (obj *_EsHalfPriceMgr) WithStartTime(startTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取 活动结束时间
func (obj *_EsHalfPriceMgr) WithEndTime(endTime int64) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithTitle title获取 活动标题
func (obj *_EsHalfPriceMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithRangeType range_type获取 是否全部商品参与
func (obj *_EsHalfPriceMgr) WithRangeType(rangeType int) Option {
	return optionFunc(func(o *options) { o.query["range_type"] = rangeType })
}

// WithDisabled disabled获取 是否停用
func (obj *_EsHalfPriceMgr) WithDisabled(disabled int) Option {
	return optionFunc(func(o *options) { o.query["disabled"] = disabled })
}

// WithDescription description获取 活动说明
func (obj *_EsHalfPriceMgr) WithDescription(description string) Option {
	return optionFunc(func(o *options) { o.query["description"] = description })
}

// WithSellerID seller_id获取 商家id
func (obj *_EsHalfPriceMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// GetByOption 功能选项模式获取
func (obj *_EsHalfPriceMgr) GetByOption(opts ...Option) (result models.EsHalfPrice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsHalfPriceMgr) GetByOptions(opts ...Option) (results []*models.EsHalfPrice, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsHalfPriceMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsHalfPrice, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where(options.query)
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

// GetFromHpID 通过hp_id获取内容 第二件半价活动ID
func (obj *_EsHalfPriceMgr) GetFromHpID(hpID int) (result models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where("`hp_id` = ?", hpID).First(&result).Error

	return
}

// GetBatchFromHpID 批量查找 第二件半价活动ID
func (obj *_EsHalfPriceMgr) GetBatchFromHpID(hpIDs []int) (results []*models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where("`hp_id` IN (?)", hpIDs).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容 活动开始时间
func (obj *_EsHalfPriceMgr) GetFromStartTime(startTime int64) (results []*models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找 活动开始时间
func (obj *_EsHalfPriceMgr) GetBatchFromStartTime(startTimes []int64) (results []*models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容 活动结束时间
func (obj *_EsHalfPriceMgr) GetFromEndTime(endTime int64) (results []*models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找 活动结束时间
func (obj *_EsHalfPriceMgr) GetBatchFromEndTime(endTimes []int64) (results []*models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 活动标题
func (obj *_EsHalfPriceMgr) GetFromTitle(title string) (results []*models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 活动标题
func (obj *_EsHalfPriceMgr) GetBatchFromTitle(titles []string) (results []*models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromRangeType 通过range_type获取内容 是否全部商品参与
func (obj *_EsHalfPriceMgr) GetFromRangeType(rangeType int) (results []*models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where("`range_type` = ?", rangeType).Find(&results).Error

	return
}

// GetBatchFromRangeType 批量查找 是否全部商品参与
func (obj *_EsHalfPriceMgr) GetBatchFromRangeType(rangeTypes []int) (results []*models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where("`range_type` IN (?)", rangeTypes).Find(&results).Error

	return
}

// GetFromDisabled 通过disabled获取内容 是否停用
func (obj *_EsHalfPriceMgr) GetFromDisabled(disabled int) (results []*models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where("`disabled` = ?", disabled).Find(&results).Error

	return
}

// GetBatchFromDisabled 批量查找 是否停用
func (obj *_EsHalfPriceMgr) GetBatchFromDisabled(disableds []int) (results []*models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where("`disabled` IN (?)", disableds).Find(&results).Error

	return
}

// GetFromDescription 通过description获取内容 活动说明
func (obj *_EsHalfPriceMgr) GetFromDescription(description string) (results []*models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where("`description` = ?", description).Find(&results).Error

	return
}

// GetBatchFromDescription 批量查找 活动说明
func (obj *_EsHalfPriceMgr) GetBatchFromDescription(descriptions []string) (results []*models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where("`description` IN (?)", descriptions).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 商家id
func (obj *_EsHalfPriceMgr) GetFromSellerID(sellerID int) (results []*models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 商家id
func (obj *_EsHalfPriceMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsHalfPriceMgr) FetchByPrimaryKey(hpID int) (result models.EsHalfPrice, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHalfPrice{}).Where("`hp_id` = ?", hpID).First(&result).Error

	return
}
