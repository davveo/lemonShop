package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type SeckillMgr struct {
	*_BaseMgr
}

// NewSeckillMgr open func
func NewSeckillMgr(db db.Repo) *SeckillMgr {
	if db == nil {
		panic(fmt.Errorf("NewSeckillMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &SeckillMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_category_spec"),
		wdb:       db.GetDbW().Table("es_category_spec"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *SeckillMgr) GetTableName() string {
	return "es_seckill"
}

// Reset 重置gorm会话
func (obj *SeckillMgr) Reset() *SeckillMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *SeckillMgr) Get() (result models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *SeckillMgr) Gets() (results []*models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *SeckillMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSeckillID seckill_id获取 主键id
func (obj *SeckillMgr) WithSeckillID(seckillID int) Option {
	return optionFunc(func(o *options) { o.query["seckill_id"] = seckillID })
}

// WithSeckillName seckill_name获取 活动名称
func (obj *SeckillMgr) WithSeckillName(seckillName string) Option {
	return optionFunc(func(o *options) { o.query["seckill_name"] = seckillName })
}

// WithStartDay start_day获取 活动日期
func (obj *SeckillMgr) WithStartDay(startDay int64) Option {
	return optionFunc(func(o *options) { o.query["start_day"] = startDay })
}

// WithApplyEndTime apply_end_time获取 报名截至时间
func (obj *SeckillMgr) WithApplyEndTime(applyEndTime int64) Option {
	return optionFunc(func(o *options) { o.query["apply_end_time"] = applyEndTime })
}

// WithSeckillRule seckill_rule获取 申请规则
func (obj *SeckillMgr) WithSeckillRule(seckillRule string) Option {
	return optionFunc(func(o *options) { o.query["seckill_rule"] = seckillRule })
}

// WithSellerIDs seller_ids获取 商家id
func (obj *SeckillMgr) WithSellerIDs(sellerIDs string) Option {
	return optionFunc(func(o *options) { o.query["seller_ids"] = sellerIDs })
}

// WithSeckillStatus seckill_status获取 状态
func (obj *SeckillMgr) WithSeckillStatus(seckillStatus string) Option {
	return optionFunc(func(o *options) { o.query["seckill_status"] = seckillStatus })
}

// WithDeleteStatus delete_status获取 是否删除 DELETED：已删除，NORMAL：正常
func (obj *SeckillMgr) WithDeleteStatus(deleteStatus string) Option {
	return optionFunc(func(o *options) { o.query["delete_status"] = deleteStatus })
}

// GetByOption 功能选项模式获取
func (obj *SeckillMgr) GetByOption(opts ...Option) (result models.EsSeckill, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *SeckillMgr) GetByOptions(opts ...Option) (results []*models.EsSeckill, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *SeckillMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSeckill, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where(options.query)
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

// GetFromSeckillID 通过seckill_id获取内容 主键id
func (obj *SeckillMgr) GetFromSeckillID(seckillID int) (result models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where("`seckill_id` = ?", seckillID).First(&result).Error

	return
}

// GetBatchFromSeckillID 批量查找 主键id
func (obj *SeckillMgr) GetBatchFromSeckillID(seckillIDs []int) (results []*models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where("`seckill_id` IN (?)", seckillIDs).Find(&results).Error

	return
}

// GetFromSeckillName 通过seckill_name获取内容 活动名称
func (obj *SeckillMgr) GetFromSeckillName(seckillName string) (results []*models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where("`seckill_name` = ?", seckillName).Find(&results).Error

	return
}

// GetBatchFromSeckillName 批量查找 活动名称
func (obj *SeckillMgr) GetBatchFromSeckillName(seckillNames []string) (results []*models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where("`seckill_name` IN (?)", seckillNames).Find(&results).Error

	return
}

// GetFromStartDay 通过start_day获取内容 活动日期
func (obj *SeckillMgr) GetFromStartDay(startDay int64) (results []*models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where("`start_day` = ?", startDay).Find(&results).Error

	return
}

// GetBatchFromStartDay 批量查找 活动日期
func (obj *SeckillMgr) GetBatchFromStartDay(startDays []int64) (results []*models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where("`start_day` IN (?)", startDays).Find(&results).Error

	return
}

// GetFromApplyEndTime 通过apply_end_time获取内容 报名截至时间
func (obj *SeckillMgr) GetFromApplyEndTime(applyEndTime int64) (results []*models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where("`apply_end_time` = ?", applyEndTime).Find(&results).Error

	return
}

// GetBatchFromApplyEndTime 批量查找 报名截至时间
func (obj *SeckillMgr) GetBatchFromApplyEndTime(applyEndTimes []int64) (results []*models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where("`apply_end_time` IN (?)", applyEndTimes).Find(&results).Error

	return
}

// GetFromSeckillRule 通过seckill_rule获取内容 申请规则
func (obj *SeckillMgr) GetFromSeckillRule(seckillRule string) (results []*models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where("`seckill_rule` = ?", seckillRule).Find(&results).Error

	return
}

// GetBatchFromSeckillRule 批量查找 申请规则
func (obj *SeckillMgr) GetBatchFromSeckillRule(seckillRules []string) (results []*models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where("`seckill_rule` IN (?)", seckillRules).Find(&results).Error

	return
}

// GetFromSellerIDs 通过seller_ids获取内容 商家id
func (obj *SeckillMgr) GetFromSellerIDs(sellerIDs string) (results []*models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where("`seller_ids` = ?", sellerIDs).Find(&results).Error

	return
}

// GetBatchFromSellerIDs 批量查找 商家id
func (obj *SeckillMgr) GetBatchFromSellerIDs(sellerIDss []string) (results []*models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where("`seller_ids` IN (?)", sellerIDss).Find(&results).Error

	return
}

// GetFromSeckillStatus 通过seckill_status获取内容 状态
func (obj *SeckillMgr) GetFromSeckillStatus(seckillStatus string) (results []*models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where("`seckill_status` = ?", seckillStatus).Find(&results).Error

	return
}

// GetBatchFromSeckillStatus 批量查找 状态
func (obj *SeckillMgr) GetBatchFromSeckillStatus(seckillStatuss []string) (results []*models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where("`seckill_status` IN (?)", seckillStatuss).Find(&results).Error

	return
}

// GetFromDeleteStatus 通过delete_status获取内容 是否删除 DELETED：已删除，NORMAL：正常
func (obj *SeckillMgr) GetFromDeleteStatus(deleteStatus string) (results []*models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where("`delete_status` = ?", deleteStatus).Find(&results).Error

	return
}

// GetBatchFromDeleteStatus 批量查找 是否删除 DELETED：已删除，NORMAL：正常
func (obj *SeckillMgr) GetBatchFromDeleteStatus(deleteStatuss []string) (results []*models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where("`delete_status` IN (?)", deleteStatuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *SeckillMgr) FetchByPrimaryKey(seckillID int) (result models.EsSeckill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckill{}).Where("`seckill_id` = ?", seckillID).First(&result).Error

	return
}
