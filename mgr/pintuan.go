package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsPintuanMgr struct {
	*_BaseMgr
}

// EsPintuanMgr open func
func EsPintuanMgr(db *gorm.DB) *_EsPintuanMgr {
	if db == nil {
		panic(fmt.Errorf("EsPintuanMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsPintuanMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_pintuan"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsPintuanMgr) GetTableName() string {
	return "es_pintuan"
}

// Reset 重置gorm会话
func (obj *_EsPintuanMgr) Reset() *_EsPintuanMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsPintuanMgr) Get() (result models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsPintuanMgr) Gets() (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsPintuanMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithPromotionID promotion_id获取 id
func (obj *_EsPintuanMgr) WithPromotionID(promotionID int) Option {
	return optionFunc(func(o *options) { o.query["promotion_id"] = promotionID })
}

// WithPromotionName promotion_name获取 活动名称
func (obj *_EsPintuanMgr) WithPromotionName(promotionName string) Option {
	return optionFunc(func(o *options) { o.query["promotion_name"] = promotionName })
}

// WithPromotionTitle promotion_title获取 活动标题
func (obj *_EsPintuanMgr) WithPromotionTitle(promotionTitle string) Option {
	return optionFunc(func(o *options) { o.query["promotion_title"] = promotionTitle })
}

// WithStartTime start_time获取 活动开始时间
func (obj *_EsPintuanMgr) WithStartTime(startTime int64) Option {
	return optionFunc(func(o *options) { o.query["start_time"] = startTime })
}

// WithEndTime end_time获取 活动结束时间
func (obj *_EsPintuanMgr) WithEndTime(endTime int64) Option {
	return optionFunc(func(o *options) { o.query["end_time"] = endTime })
}

// WithRequiredNum required_num获取 成团人数
func (obj *_EsPintuanMgr) WithRequiredNum(requiredNum int) Option {
	return optionFunc(func(o *options) { o.query["required_num"] = requiredNum })
}

// WithLimitNum limit_num获取 限购数量
func (obj *_EsPintuanMgr) WithLimitNum(limitNum int) Option {
	return optionFunc(func(o *options) { o.query["limit_num"] = limitNum })
}

// WithEnableMocker enable_mocker获取 虚拟成团
func (obj *_EsPintuanMgr) WithEnableMocker(enableMocker int16) Option {
	return optionFunc(func(o *options) { o.query["enable_mocker"] = enableMocker })
}

// WithCreateTime create_time获取 创建时间
func (obj *_EsPintuanMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithStatus status获取 活动状态
func (obj *_EsPintuanMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithOptionStatus option_status获取 api请求操作状态
func (obj *_EsPintuanMgr) WithOptionStatus(optionStatus string) Option {
	return optionFunc(func(o *options) { o.query["option_status"] = optionStatus })
}

// WithSellerName seller_name获取 商家名称
func (obj *_EsPintuanMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithSellerID seller_id获取 商家ID
func (obj *_EsPintuanMgr) WithSellerID(sellerID int64) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// GetByOption 功能选项模式获取
func (obj *_EsPintuanMgr) GetByOption(opts ...Option) (result models.EsPintuan, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsPintuanMgr) GetByOptions(opts ...Option) (results []*models.EsPintuan, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsPintuanMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsPintuan, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where(options.query)
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
func (obj *_EsPintuanMgr) GetFromPromotionID(promotionID int) (result models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`promotion_id` = ?", promotionID).First(&result).Error

	return
}

// GetBatchFromPromotionID 批量查找 id
func (obj *_EsPintuanMgr) GetBatchFromPromotionID(promotionIDs []int) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`promotion_id` IN (?)", promotionIDs).Find(&results).Error

	return
}

// GetFromPromotionName 通过promotion_name获取内容 活动名称
func (obj *_EsPintuanMgr) GetFromPromotionName(promotionName string) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`promotion_name` = ?", promotionName).Find(&results).Error

	return
}

// GetBatchFromPromotionName 批量查找 活动名称
func (obj *_EsPintuanMgr) GetBatchFromPromotionName(promotionNames []string) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`promotion_name` IN (?)", promotionNames).Find(&results).Error

	return
}

// GetFromPromotionTitle 通过promotion_title获取内容 活动标题
func (obj *_EsPintuanMgr) GetFromPromotionTitle(promotionTitle string) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`promotion_title` = ?", promotionTitle).Find(&results).Error

	return
}

// GetBatchFromPromotionTitle 批量查找 活动标题
func (obj *_EsPintuanMgr) GetBatchFromPromotionTitle(promotionTitles []string) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`promotion_title` IN (?)", promotionTitles).Find(&results).Error

	return
}

// GetFromStartTime 通过start_time获取内容 活动开始时间
func (obj *_EsPintuanMgr) GetFromStartTime(startTime int64) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`start_time` = ?", startTime).Find(&results).Error

	return
}

// GetBatchFromStartTime 批量查找 活动开始时间
func (obj *_EsPintuanMgr) GetBatchFromStartTime(startTimes []int64) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`start_time` IN (?)", startTimes).Find(&results).Error

	return
}

// GetFromEndTime 通过end_time获取内容 活动结束时间
func (obj *_EsPintuanMgr) GetFromEndTime(endTime int64) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`end_time` = ?", endTime).Find(&results).Error

	return
}

// GetBatchFromEndTime 批量查找 活动结束时间
func (obj *_EsPintuanMgr) GetBatchFromEndTime(endTimes []int64) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`end_time` IN (?)", endTimes).Find(&results).Error

	return
}

// GetFromRequiredNum 通过required_num获取内容 成团人数
func (obj *_EsPintuanMgr) GetFromRequiredNum(requiredNum int) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`required_num` = ?", requiredNum).Find(&results).Error

	return
}

// GetBatchFromRequiredNum 批量查找 成团人数
func (obj *_EsPintuanMgr) GetBatchFromRequiredNum(requiredNums []int) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`required_num` IN (?)", requiredNums).Find(&results).Error

	return
}

// GetFromLimitNum 通过limit_num获取内容 限购数量
func (obj *_EsPintuanMgr) GetFromLimitNum(limitNum int) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`limit_num` = ?", limitNum).Find(&results).Error

	return
}

// GetBatchFromLimitNum 批量查找 限购数量
func (obj *_EsPintuanMgr) GetBatchFromLimitNum(limitNums []int) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`limit_num` IN (?)", limitNums).Find(&results).Error

	return
}

// GetFromEnableMocker 通过enable_mocker获取内容 虚拟成团
func (obj *_EsPintuanMgr) GetFromEnableMocker(enableMocker int16) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`enable_mocker` = ?", enableMocker).Find(&results).Error

	return
}

// GetBatchFromEnableMocker 批量查找 虚拟成团
func (obj *_EsPintuanMgr) GetBatchFromEnableMocker(enableMockers []int16) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`enable_mocker` IN (?)", enableMockers).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_EsPintuanMgr) GetFromCreateTime(createTime int64) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_EsPintuanMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 活动状态
func (obj *_EsPintuanMgr) GetFromStatus(status string) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 活动状态
func (obj *_EsPintuanMgr) GetBatchFromStatus(statuss []string) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromOptionStatus 通过option_status获取内容 api请求操作状态
func (obj *_EsPintuanMgr) GetFromOptionStatus(optionStatus string) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`option_status` = ?", optionStatus).Find(&results).Error

	return
}

// GetBatchFromOptionStatus 批量查找 api请求操作状态
func (obj *_EsPintuanMgr) GetBatchFromOptionStatus(optionStatuss []string) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`option_status` IN (?)", optionStatuss).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 商家名称
func (obj *_EsPintuanMgr) GetFromSellerName(sellerName string) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 商家名称
func (obj *_EsPintuanMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 商家ID
func (obj *_EsPintuanMgr) GetFromSellerID(sellerID int64) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 商家ID
func (obj *_EsPintuanMgr) GetBatchFromSellerID(sellerIDs []int64) (results []*models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsPintuanMgr) FetchByPrimaryKey(promotionID int) (result models.EsPintuan, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsPintuan{}).Where("`promotion_id` = ?", promotionID).First(&result).Error

	return
}
