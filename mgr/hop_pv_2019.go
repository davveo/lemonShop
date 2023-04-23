package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsSssShopPv2019Mgr struct {
	*_BaseMgr
}

// EsSssShopPv2019Mgr open func
func EsSssShopPv2019Mgr(db *gorm.DB) *_EsSssShopPv2019Mgr {
	if db == nil {
		panic(fmt.Errorf("EsSssShopPv2019Mgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsSssShopPv2019Mgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_sss_shop_pv_2019"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsSssShopPv2019Mgr) GetTableName() string {
	return "es_sss_shop_pv_2019"
}

// Reset 重置gorm会话
func (obj *_EsSssShopPv2019Mgr) Reset() *_EsSssShopPv2019Mgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsSssShopPv2019Mgr) Get() (result models.EsSssShopPv2019, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsSssShopPv2019Mgr) Gets() (results []*models.EsSssShopPv2019, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsSssShopPv2019Mgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_EsSssShopPv2019Mgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSellerID seller_id获取
func (obj *_EsSssShopPv2019Mgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithVsYear vs_year获取
func (obj *_EsSssShopPv2019Mgr) WithVsYear(vsYear int) Option {
	return optionFunc(func(o *options) { o.query["vs_year"] = vsYear })
}

// WithVsMonth vs_month获取
func (obj *_EsSssShopPv2019Mgr) WithVsMonth(vsMonth int) Option {
	return optionFunc(func(o *options) { o.query["vs_month"] = vsMonth })
}

// WithVsDay vs_day获取
func (obj *_EsSssShopPv2019Mgr) WithVsDay(vsDay int) Option {
	return optionFunc(func(o *options) { o.query["vs_day"] = vsDay })
}

// WithVsNum vs_num获取
func (obj *_EsSssShopPv2019Mgr) WithVsNum(vsNum int) Option {
	return optionFunc(func(o *options) { o.query["vs_num"] = vsNum })
}

// GetByOption 功能选项模式获取
func (obj *_EsSssShopPv2019Mgr) GetByOption(opts ...Option) (result models.EsSssShopPv2019, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsSssShopPv2019Mgr) GetByOptions(opts ...Option) (results []*models.EsSssShopPv2019, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsSssShopPv2019Mgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSssShopPv2019, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *_EsSssShopPv2019Mgr) GetFromID(id int) (result models.EsSssShopPv2019, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_EsSssShopPv2019Mgr) GetBatchFromID(ids []int) (results []*models.EsSssShopPv2019, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容
func (obj *_EsSssShopPv2019Mgr) GetFromSellerID(sellerID int) (results []*models.EsSssShopPv2019, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找
func (obj *_EsSssShopPv2019Mgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsSssShopPv2019, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromVsYear 通过vs_year获取内容
func (obj *_EsSssShopPv2019Mgr) GetFromVsYear(vsYear int) (results []*models.EsSssShopPv2019, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).Where("`vs_year` = ?", vsYear).Find(&results).Error

	return
}

// GetBatchFromVsYear 批量查找
func (obj *_EsSssShopPv2019Mgr) GetBatchFromVsYear(vsYears []int) (results []*models.EsSssShopPv2019, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).Where("`vs_year` IN (?)", vsYears).Find(&results).Error

	return
}

// GetFromVsMonth 通过vs_month获取内容
func (obj *_EsSssShopPv2019Mgr) GetFromVsMonth(vsMonth int) (results []*models.EsSssShopPv2019, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).Where("`vs_month` = ?", vsMonth).Find(&results).Error

	return
}

// GetBatchFromVsMonth 批量查找
func (obj *_EsSssShopPv2019Mgr) GetBatchFromVsMonth(vsMonths []int) (results []*models.EsSssShopPv2019, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).Where("`vs_month` IN (?)", vsMonths).Find(&results).Error

	return
}

// GetFromVsDay 通过vs_day获取内容
func (obj *_EsSssShopPv2019Mgr) GetFromVsDay(vsDay int) (results []*models.EsSssShopPv2019, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).Where("`vs_day` = ?", vsDay).Find(&results).Error

	return
}

// GetBatchFromVsDay 批量查找
func (obj *_EsSssShopPv2019Mgr) GetBatchFromVsDay(vsDays []int) (results []*models.EsSssShopPv2019, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).Where("`vs_day` IN (?)", vsDays).Find(&results).Error

	return
}

// GetFromVsNum 通过vs_num获取内容
func (obj *_EsSssShopPv2019Mgr) GetFromVsNum(vsNum int) (results []*models.EsSssShopPv2019, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).Where("`vs_num` = ?", vsNum).Find(&results).Error

	return
}

// GetBatchFromVsNum 批量查找
func (obj *_EsSssShopPv2019Mgr) GetBatchFromVsNum(vsNums []int) (results []*models.EsSssShopPv2019, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).Where("`vs_num` IN (?)", vsNums).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsSssShopPv2019Mgr) FetchByPrimaryKey(id int) (result models.EsSssShopPv2019, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssShopPv2019{}).Where("`id` = ?", id).First(&result).Error

	return
}
