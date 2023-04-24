package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type SeckillRangeMgr struct {
	*_BaseMgr
}

// NewSeckillRangeMgr open func
func NewSeckillRangeMgr(db db.Repo) *SeckillRangeMgr {
	if db == nil {
		panic(fmt.Errorf("NewSeckillRangeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &SeckillRangeMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_category_spec"),
		wdb:       db.GetDbW().Table("es_category_spec"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *SeckillRangeMgr) GetTableName() string {
	return "es_seckill_range"
}

// Reset 重置gorm会话
func (obj *SeckillRangeMgr) Reset() *SeckillRangeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *SeckillRangeMgr) Get() (result models.EsSeckillRange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillRange{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *SeckillRangeMgr) Gets() (results []*models.EsSeckillRange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillRange{}).Find(&results).Error

	return
}

func (obj *SeckillRangeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillRange{}).Count(count)
}

// WithRangeID range_id获取 主键
func (obj *SeckillRangeMgr) WithRangeID(rangeID int) Option {
	return optionFunc(func(o *options) { o.query["range_id"] = rangeID })
}

// WithSeckillID seckill_id获取 限时抢购活动id
func (obj *SeckillRangeMgr) WithSeckillID(seckillID int) Option {
	return optionFunc(func(o *options) { o.query["seckill_id"] = seckillID })
}

// WithRangeTime range_time获取 整点时刻
func (obj *SeckillRangeMgr) WithRangeTime(rangeTime int) Option {
	return optionFunc(func(o *options) { o.query["range_time"] = rangeTime })
}

// GetByOption 功能选项模式获取
func (obj *SeckillRangeMgr) GetByOption(opts ...Option) (result models.EsSeckillRange, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillRange{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *SeckillRangeMgr) GetByOptions(opts ...Option) (results []*models.EsSeckillRange, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillRange{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *SeckillRangeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSeckillRange, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillRange{}).Where(options.query)
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

// GetFromRangeID 通过range_id获取内容 主键
func (obj *SeckillRangeMgr) GetFromRangeID(rangeID int) (result models.EsSeckillRange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillRange{}).Where("`range_id` = ?", rangeID).First(&result).Error

	return
}

// GetBatchFromRangeID 批量查找 主键
func (obj *SeckillRangeMgr) GetBatchFromRangeID(rangeIDs []int) (results []*models.EsSeckillRange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillRange{}).Where("`range_id` IN (?)", rangeIDs).Find(&results).Error

	return
}

// GetFromSeckillID 通过seckill_id获取内容 限时抢购活动id
func (obj *SeckillRangeMgr) GetFromSeckillID(seckillID int) (results []*models.EsSeckillRange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillRange{}).Where("`seckill_id` = ?", seckillID).Find(&results).Error

	return
}

// GetBatchFromSeckillID 批量查找 限时抢购活动id
func (obj *SeckillRangeMgr) GetBatchFromSeckillID(seckillIDs []int) (results []*models.EsSeckillRange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillRange{}).Where("`seckill_id` IN (?)", seckillIDs).Find(&results).Error

	return
}

// GetFromRangeTime 通过range_time获取内容 整点时刻
func (obj *SeckillRangeMgr) GetFromRangeTime(rangeTime int) (results []*models.EsSeckillRange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillRange{}).Where("`range_time` = ?", rangeTime).Find(&results).Error

	return
}

// GetBatchFromRangeTime 批量查找 整点时刻
func (obj *SeckillRangeMgr) GetBatchFromRangeTime(rangeTimes []int) (results []*models.EsSeckillRange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillRange{}).Where("`range_time` IN (?)", rangeTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *SeckillRangeMgr) FetchByPrimaryKey(rangeID int) (result models.EsSeckillRange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSeckillRange{}).Where("`range_id` = ?", rangeID).First(&result).Error

	return
}
