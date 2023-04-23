package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsSssShopPvMgr struct {
	*_BaseMgr
}

// EsSssShopPvMgr open func
func EsSssShopPvMgr(db *gorm.DB) *_EsSssShopPvMgr {
	if db == nil {
		panic(fmt.Errorf("EsSssShopPvMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsSssShopPvMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_sss_shop_pv"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsSssShopPvMgr) GetTableName() string {
	return "es_sss_shop_pv"
}

// Reset 重置gorm会话
func (obj *_EsSssShopPvMgr) Reset() *_EsSssShopPvMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsSssShopPvMgr) Get() (result models.EsSssShopPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsSssShopPvMgr) Gets() (results []*models.EsSssShopPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsSssShopPvMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_EsSssShopPvMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSellerID seller_id获取 店铺id
func (obj *_EsSssShopPvMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithVsYear vs_year获取 年份
func (obj *_EsSssShopPvMgr) WithVsYear(vsYear int) Option {
	return optionFunc(func(o *options) { o.query["vs_year"] = vsYear })
}

// WithVsMonth vs_month获取 月份
func (obj *_EsSssShopPvMgr) WithVsMonth(vsMonth int) Option {
	return optionFunc(func(o *options) { o.query["vs_month"] = vsMonth })
}

// WithVsDay vs_day获取 日期
func (obj *_EsSssShopPvMgr) WithVsDay(vsDay int) Option {
	return optionFunc(func(o *options) { o.query["vs_day"] = vsDay })
}

// WithVsNum vs_num获取 访问量
func (obj *_EsSssShopPvMgr) WithVsNum(vsNum int) Option {
	return optionFunc(func(o *options) { o.query["vs_num"] = vsNum })
}

// GetByOption 功能选项模式获取
func (obj *_EsSssShopPvMgr) GetByOption(opts ...Option) (result models.EsSssShopPv, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsSssShopPvMgr) GetByOptions(opts ...Option) (results []*models.EsSssShopPv, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsSssShopPvMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSssShopPv, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键id
func (obj *_EsSssShopPvMgr) GetFromID(id int) (result models.EsSssShopPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_EsSssShopPvMgr) GetBatchFromID(ids []int) (results []*models.EsSssShopPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 店铺id
func (obj *_EsSssShopPvMgr) GetFromSellerID(sellerID int) (results []*models.EsSssShopPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 店铺id
func (obj *_EsSssShopPvMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsSssShopPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromVsYear 通过vs_year获取内容 年份
func (obj *_EsSssShopPvMgr) GetFromVsYear(vsYear int) (results []*models.EsSssShopPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).Where("`vs_year` = ?", vsYear).Find(&results).Error

	return
}

// GetBatchFromVsYear 批量查找 年份
func (obj *_EsSssShopPvMgr) GetBatchFromVsYear(vsYears []int) (results []*models.EsSssShopPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).Where("`vs_year` IN (?)", vsYears).Find(&results).Error

	return
}

// GetFromVsMonth 通过vs_month获取内容 月份
func (obj *_EsSssShopPvMgr) GetFromVsMonth(vsMonth int) (results []*models.EsSssShopPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).Where("`vs_month` = ?", vsMonth).Find(&results).Error

	return
}

// GetBatchFromVsMonth 批量查找 月份
func (obj *_EsSssShopPvMgr) GetBatchFromVsMonth(vsMonths []int) (results []*models.EsSssShopPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).Where("`vs_month` IN (?)", vsMonths).Find(&results).Error

	return
}

// GetFromVsDay 通过vs_day获取内容 日期
func (obj *_EsSssShopPvMgr) GetFromVsDay(vsDay int) (results []*models.EsSssShopPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).Where("`vs_day` = ?", vsDay).Find(&results).Error

	return
}

// GetBatchFromVsDay 批量查找 日期
func (obj *_EsSssShopPvMgr) GetBatchFromVsDay(vsDays []int) (results []*models.EsSssShopPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).Where("`vs_day` IN (?)", vsDays).Find(&results).Error

	return
}

// GetFromVsNum 通过vs_num获取内容 访问量
func (obj *_EsSssShopPvMgr) GetFromVsNum(vsNum int) (results []*models.EsSssShopPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).Where("`vs_num` = ?", vsNum).Find(&results).Error

	return
}

// GetBatchFromVsNum 批量查找 访问量
func (obj *_EsSssShopPvMgr) GetBatchFromVsNum(vsNums []int) (results []*models.EsSssShopPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).Where("`vs_num` IN (?)", vsNums).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsSssShopPvMgr) FetchByPrimaryKey(id int) (result models.EsSssShopPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv{}).Where("`id` = ?", id).First(&result).Error

	return
}
