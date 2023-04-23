package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsGroupbuyActiveMgr struct {
	*_BaseMgr
}

// EsGroupbuyActiveMgr open func
func EsGroupbuyActiveMgr(db *gorm.DB) *_EsGroupbuyActiveMgr {
	if db == nil {
		panic(fmt.Errorf("EsGroupbuyActiveMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsGroupbuyActiveMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_groupbuy_active"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsGroupbuyActiveMgr) GetTableName() string {
	return "es_groupbuy_active"
}

// Reset 重置gorm会话
func (obj *_EsGroupbuyActiveMgr) Reset() *_EsGroupbuyActiveMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsGroupbuyActiveMgr) Get() (result models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsGroupbuyActiveMgr) Gets() (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsGroupbuyActiveMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithActID act_id获取 活动Id
func (obj *_EsGroupbuyActiveMgr) WithActID(actID int) Option {
	return optionFunc(func(o *options) { o.query["act_id"] = actID })
}

// WithActName act_name获取 活动名称
func (obj *_EsGroupbuyActiveMgr) WithActName(actName string) Option {
	return optionFunc(func(o *options) { o.query["act_name"] = actName })
}

// WithStartTime start_time获取 活动开启时间
func (obj *_EsGroupbuyActiveMgr) WithStartTime(startTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取 团购结束时间
func (obj *_EsGroupbuyActiveMgr) WithEndTime(endTime int64) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithJoinEndTime join_end_time获取 团购报名截止时间
func (obj *_EsGroupbuyActiveMgr) WithJoinEndTime(joinEndTime int64) Option {
	return optionFunc(func(o *options) { o.query["join_end_time"] = joinEndTime })
}

// WithAddTime add_time获取 团购添加时间
func (obj *_EsGroupbuyActiveMgr) WithAddTime(addTime int64) Option {
	return optionFunc(func(o *options) { o.query["add_time"] = addTime })
}

// WithActTagID act_tag_id获取 团购活动标签Id
func (obj *_EsGroupbuyActiveMgr) WithActTagID(actTagID int) Option {
	return optionFunc(func(o *options) { o.query["act_tag_id"] = actTagID })
}

// WithGoodsNum goods_num获取 参与团购商品数量
func (obj *_EsGroupbuyActiveMgr) WithGoodsNum(goodsNum int) Option {
	return optionFunc(func(o *options) { o.query["goods_num"] = goodsNum })
}

// WithDeleteReason delete_reason获取 删除原因
func (obj *_EsGroupbuyActiveMgr) WithDeleteReason(deleteReason string) Option {
	return optionFunc(func(o *options) { o.query["delete_reason"] = deleteReason })
}

// WithDeleteTime delete_time获取 删除时间
func (obj *_EsGroupbuyActiveMgr) WithDeleteTime(deleteTime int64) Option {
	return optionFunc(func(o *options) { o.query["delete_time"] = deleteTime })
}

// WithDeleteName delete_name获取 删除人
func (obj *_EsGroupbuyActiveMgr) WithDeleteName(deleteName string) Option {
	return optionFunc(func(o *options) { o.query["delete_name"] = deleteName })
}

// WithDeleteStatus delete_status获取 是否删除 DELETED：已删除，NORMAL：正常
func (obj *_EsGroupbuyActiveMgr) WithDeleteStatus(deleteStatus string) Option {
	return optionFunc(func(o *options) { o.query["delete_status"] = deleteStatus })
}

// GetByOption 功能选项模式获取
func (obj *_EsGroupbuyActiveMgr) GetByOption(opts ...Option) (result models.EsGroupbuyActive, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsGroupbuyActiveMgr) GetByOptions(opts ...Option) (results []*models.EsGroupbuyActive, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsGroupbuyActiveMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsGroupbuyActive, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where(options.query)
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

// GetFromActID 通过act_id获取内容 活动Id
func (obj *_EsGroupbuyActiveMgr) GetFromActID(actID int) (result models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`act_id` = ?", actID).First(&result).Error

	return
}

// GetBatchFromActID 批量查找 活动Id
func (obj *_EsGroupbuyActiveMgr) GetBatchFromActID(actIDs []int) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`act_id` IN (?)", actIDs).Find(&results).Error

	return
}

// GetFromActName 通过act_name获取内容 活动名称
func (obj *_EsGroupbuyActiveMgr) GetFromActName(actName string) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`act_name` = ?", actName).Find(&results).Error

	return
}

// GetBatchFromActName 批量查找 活动名称
func (obj *_EsGroupbuyActiveMgr) GetBatchFromActName(actNames []string) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`act_name` IN (?)", actNames).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容 活动开启时间
func (obj *_EsGroupbuyActiveMgr) GetFromStartTime(startTime int64) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找 活动开启时间
func (obj *_EsGroupbuyActiveMgr) GetBatchFromStartTime(startTimes []int64) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容 团购结束时间
func (obj *_EsGroupbuyActiveMgr) GetFromEndTime(endTime int64) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找 团购结束时间
func (obj *_EsGroupbuyActiveMgr) GetBatchFromEndTime(endTimes []int64) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromJoinEndTime 通过join_end_time获取内容 团购报名截止时间
func (obj *_EsGroupbuyActiveMgr) GetFromJoinEndTime(joinEndTime int64) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`join_end_time` = ?", joinEndTime).Find(&results).Error

	return
}

// GetBatchFromJoinEndTime 批量查找 团购报名截止时间
func (obj *_EsGroupbuyActiveMgr) GetBatchFromJoinEndTime(joinEndTimes []int64) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`join_end_time` IN (?)", joinEndTimes).Find(&results).Error

	return
}

// GetFromAddTime 通过add_time获取内容 团购添加时间
func (obj *_EsGroupbuyActiveMgr) GetFromAddTime(addTime int64) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`add_time` = ?", addTime).Find(&results).Error

	return
}

// GetBatchFromAddTime 批量查找 团购添加时间
func (obj *_EsGroupbuyActiveMgr) GetBatchFromAddTime(addTimes []int64) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`add_time` IN (?)", addTimes).Find(&results).Error

	return
}

// GetFromActTagID 通过act_tag_id获取内容 团购活动标签Id
func (obj *_EsGroupbuyActiveMgr) GetFromActTagID(actTagID int) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`act_tag_id` = ?", actTagID).Find(&results).Error

	return
}

// GetBatchFromActTagID 批量查找 团购活动标签Id
func (obj *_EsGroupbuyActiveMgr) GetBatchFromActTagID(actTagIDs []int) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`act_tag_id` IN (?)", actTagIDs).Find(&results).Error

	return
}

// GetFromGoodsNum 通过goods_num获取内容 参与团购商品数量
func (obj *_EsGroupbuyActiveMgr) GetFromGoodsNum(goodsNum int) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`goods_num` = ?", goodsNum).Find(&results).Error

	return
}

// GetBatchFromGoodsNum 批量查找 参与团购商品数量
func (obj *_EsGroupbuyActiveMgr) GetBatchFromGoodsNum(goodsNums []int) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`goods_num` IN (?)", goodsNums).Find(&results).Error

	return
}

// GetFromDeleteReason 通过delete_reason获取内容 删除原因
func (obj *_EsGroupbuyActiveMgr) GetFromDeleteReason(deleteReason string) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`delete_reason` = ?", deleteReason).Find(&results).Error

	return
}

// GetBatchFromDeleteReason 批量查找 删除原因
func (obj *_EsGroupbuyActiveMgr) GetBatchFromDeleteReason(deleteReasons []string) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`delete_reason` IN (?)", deleteReasons).Find(&results).Error

	return
}

// GetFromDeleteTime 通过delete_time获取内容 删除时间
func (obj *_EsGroupbuyActiveMgr) GetFromDeleteTime(deleteTime int64) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`delete_time` = ?", deleteTime).Find(&results).Error

	return
}

// GetBatchFromDeleteTime 批量查找 删除时间
func (obj *_EsGroupbuyActiveMgr) GetBatchFromDeleteTime(deleteTimes []int64) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`delete_time` IN (?)", deleteTimes).Find(&results).Error

	return
}

// GetFromDeleteName 通过delete_name获取内容 删除人
func (obj *_EsGroupbuyActiveMgr) GetFromDeleteName(deleteName string) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`delete_name` = ?", deleteName).Find(&results).Error

	return
}

// GetBatchFromDeleteName 批量查找 删除人
func (obj *_EsGroupbuyActiveMgr) GetBatchFromDeleteName(deleteNames []string) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`delete_name` IN (?)", deleteNames).Find(&results).Error

	return
}

// GetFromDeleteStatus 通过delete_status获取内容 是否删除 DELETED：已删除，NORMAL：正常
func (obj *_EsGroupbuyActiveMgr) GetFromDeleteStatus(deleteStatus string) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`delete_status` = ?", deleteStatus).Find(&results).Error

	return
}

// GetBatchFromDeleteStatus 批量查找 是否删除 DELETED：已删除，NORMAL：正常
func (obj *_EsGroupbuyActiveMgr) GetBatchFromDeleteStatus(deleteStatuss []string) (results []*models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`delete_status` IN (?)", deleteStatuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsGroupbuyActiveMgr) FetchByPrimaryKey(actID int) (result models.EsGroupbuyActive, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGroupbuyActive{}).Where("`act_id` = ?", actID).First(&result).Error

	return
}
