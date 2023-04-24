package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type PintuanMgr struct {
	*_BaseMgr
}

// NewPintuanMgr open func
func NewPintuanMgr(db db.Repo) *PintuanMgr {
	if db == nil {
		panic(fmt.Errorf("NewPintuanMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &PintuanMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_pintuan"), wdb: db.GetDbW().Table("es_pintuan"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *PintuanMgr) GetTableName() string {
	return "es_pintuan"
}

// Reset 重置gorm会话
func (obj *PintuanMgr) Reset() *PintuanMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *PintuanMgr) Get() (result models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *PintuanMgr) Gets() (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *PintuanMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithPromotionID promotion_id获取 id
func (obj *PintuanMgr) WithPromotionID(promotionID int) Option {
	return optionFunc(func(o *options) { o.query["promotion_id"] = promotionID })
}

// WithPromotionName promotion_name获取 活动名称
func (obj *PintuanMgr) WithPromotionName(promotionName string) Option {
	return optionFunc(func(o *options) { o.query["promotion_name"] = promotionName })
}

// WithPromotionTitle promotion_title获取 活动标题
func (obj *PintuanMgr) WithPromotionTitle(promotionTitle string) Option {
	return optionFunc(func(o *options) { o.query["promotion_title"] = promotionTitle })
}

// WithStartTime start_time获取 活动开始时间
func (obj *PintuanMgr) WithStartTime(startTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取 活动结束时间
func (obj *PintuanMgr) WithEndTime(endTime int64) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithRequiredNum required_num获取 成团人数
func (obj *PintuanMgr) WithRequiredNum(requiredNum int) Option {
	return optionFunc(func(o *options) { o.query["required_num"] = requiredNum })
}

// WithLimitNum limit_num获取 限购数量
func (obj *PintuanMgr) WithLimitNum(limitNum int) Option {
	return optionFunc(func(o *options) { o.query["limit_num"] = limitNum })
}

// WithEnableMocker enable_mocker获取 虚拟成团
func (obj *PintuanMgr) WithEnableMocker(enableMocker int16) Option {
	return optionFunc(func(o *options) { o.query["enable_mocker"] = enableMocker })
}

// WithCreateTime create_time获取 创建时间
func (obj *PintuanMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithStatus status获取 活动状态
func (obj *PintuanMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithOptionStatus option_status获取 api请求操作状态
func (obj *PintuanMgr) WithOptionStatus(optionStatus string) Option {
	return optionFunc(func(o *options) { o.query["option_status"] = optionStatus })
}

// WithSellerName seller_name获取 商家名称
func (obj *PintuanMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithSellerID seller_id获取 商家ID
func (obj *PintuanMgr) WithSellerID(sellerID int64) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// GetByOption 功能选项模式获取
func (obj *PintuanMgr) GetByOption(opts ...Option) (result models.EsPintuan, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *PintuanMgr) GetByOptions(opts ...Option) (results []*models.EsPintuan, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *PintuanMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsPintuan, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where(options.query)
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

// GetFromPromotionID 通过promotion_id获取内容 id
func (obj *PintuanMgr) GetFromPromotionID(promotionID int) (result models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`promotion_id` = ?", promotionID).First(&result).Error

	return
}

// GetBatchFromPromotionID 批量查找 id
func (obj *PintuanMgr) GetBatchFromPromotionID(promotionIDs []int) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`promotion_id` IN (?)", promotionIDs).Find(&results).Error

	return
}

// GetFromPromotionName 通过promotion_name获取内容 活动名称
func (obj *PintuanMgr) GetFromPromotionName(promotionName string) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`promotion_name` = ?", promotionName).Find(&results).Error

	return
}

// GetBatchFromPromotionName 批量查找 活动名称
func (obj *PintuanMgr) GetBatchFromPromotionName(promotionNames []string) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`promotion_name` IN (?)", promotionNames).Find(&results).Error

	return
}

// GetFromPromotionTitle 通过promotion_title获取内容 活动标题
func (obj *PintuanMgr) GetFromPromotionTitle(promotionTitle string) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`promotion_title` = ?", promotionTitle).Find(&results).Error

	return
}

// GetBatchFromPromotionTitle 批量查找 活动标题
func (obj *PintuanMgr) GetBatchFromPromotionTitle(promotionTitles []string) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`promotion_title` IN (?)", promotionTitles).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容 活动开始时间
func (obj *PintuanMgr) GetFromStartTime(startTime int64) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找 活动开始时间
func (obj *PintuanMgr) GetBatchFromStartTime(startTimes []int64) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容 活动结束时间
func (obj *PintuanMgr) GetFromEndTime(endTime int64) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找 活动结束时间
func (obj *PintuanMgr) GetBatchFromEndTime(endTimes []int64) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromRequiredNum 通过required_num获取内容 成团人数
func (obj *PintuanMgr) GetFromRequiredNum(requiredNum int) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`required_num` = ?", requiredNum).Find(&results).Error

	return
}

// GetBatchFromRequiredNum 批量查找 成团人数
func (obj *PintuanMgr) GetBatchFromRequiredNum(requiredNums []int) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`required_num` IN (?)", requiredNums).Find(&results).Error

	return
}

// GetFromLimitNum 通过limit_num获取内容 限购数量
func (obj *PintuanMgr) GetFromLimitNum(limitNum int) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`limit_num` = ?", limitNum).Find(&results).Error

	return
}

// GetBatchFromLimitNum 批量查找 限购数量
func (obj *PintuanMgr) GetBatchFromLimitNum(limitNums []int) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`limit_num` IN (?)", limitNums).Find(&results).Error

	return
}

// GetFromEnableMocker 通过enable_mocker获取内容 虚拟成团
func (obj *PintuanMgr) GetFromEnableMocker(enableMocker int16) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`enable_mocker` = ?", enableMocker).Find(&results).Error

	return
}

// GetBatchFromEnableMocker 批量查找 虚拟成团
func (obj *PintuanMgr) GetBatchFromEnableMocker(enableMockers []int16) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`enable_mocker` IN (?)", enableMockers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *PintuanMgr) GetFromCreateTime(createTime int64) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *PintuanMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 活动状态
func (obj *PintuanMgr) GetFromStatus(status string) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 活动状态
func (obj *PintuanMgr) GetBatchFromStatus(statuss []string) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromOptionStatus 通过option_status获取内容 api请求操作状态
func (obj *PintuanMgr) GetFromOptionStatus(optionStatus string) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`option_status` = ?", optionStatus).Find(&results).Error

	return
}

// GetBatchFromOptionStatus 批量查找 api请求操作状态
func (obj *PintuanMgr) GetBatchFromOptionStatus(optionStatuss []string) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`option_status` IN (?)", optionStatuss).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 商家名称
func (obj *PintuanMgr) GetFromSellerName(sellerName string) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 商家名称
func (obj *PintuanMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 商家ID
func (obj *PintuanMgr) GetFromSellerID(sellerID int64) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 商家ID
func (obj *PintuanMgr) GetBatchFromSellerID(sellerIDs []int64) (results []*models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *PintuanMgr) FetchByPrimaryKey(promotionID int) (result models.EsPintuan, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`promotion_id` = ?", promotionID).First(&result).Error

	return
}
