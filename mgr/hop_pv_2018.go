package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type SssShopPv2018Mgr struct {
	*_BaseMgr
}

// NewSssShopPv2018Mgr open func
func NewSssShopPv2018Mgr(db db.Repo) *SssShopPv2018Mgr {
	if db == nil {
		panic(fmt.Errorf("NewSssShopPv2018Mgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &SssShopPv2018Mgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_sss_shop_pv_2018"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *SssShopPv2018Mgr) GetTableName() string {
	return "es_sss_shop_pv_2018"
}

// Reset 重置gorm会话
func (obj *SssShopPv2018Mgr) Reset() *SssShopPv2018Mgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *SssShopPv2018Mgr) Get() (result models.EsSssShopPv2018, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopPv2018{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *SssShopPv2018Mgr) Gets() (results []*models.EsSssShopPv2018, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopPv2018{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *SssShopPv2018Mgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopPv2018{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *SssShopPv2018Mgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSellerID seller_id获取
func (obj *SssShopPv2018Mgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithVsYear vs_year获取
func (obj *SssShopPv2018Mgr) WithVsYear(vsYear int) Option {
	return optionFunc(func(o *options) { o.query["vs_year"] = vsYear })
}

// WithVsMonth vs_month获取
func (obj *SssShopPv2018Mgr) WithVsMonth(vsMonth int) Option {
	return optionFunc(func(o *options) { o.query["vs_month"] = vsMonth })
}

// WithVsDay vs_day获取
func (obj *SssShopPv2018Mgr) WithVsDay(vsDay int) Option {
	return optionFunc(func(o *options) { o.query["vs_day"] = vsDay })
}

// WithVsNum vs_num获取
func (obj *SssShopPv2018Mgr) WithVsNum(vsNum int) Option {
	return optionFunc(func(o *options) { o.query["vs_num"] = vsNum })
}

// GetByOption 功能选项模式获取
func (obj *SssShopPv2018Mgr) GetByOption(opts ...Option) (result models.EsSssShopPv2018, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopPv2018{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *SssShopPv2018Mgr) GetByOptions(opts ...Option) (results []*models.EsSssShopPv2018, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopPv2018{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *SssShopPv2018Mgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSssShopPv2018, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopPv2018{}).Where(options.query)
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
func (obj *SssShopPv2018Mgr) GetFromID(id int) (results []*models.EsSssShopPv2018, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopPv2018{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量查找
func (obj *SssShopPv2018Mgr) GetBatchFromID(ids []int) (results []*models.EsSssShopPv2018, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopPv2018{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容
func (obj *SssShopPv2018Mgr) GetFromSellerID(sellerID int) (results []*models.EsSssShopPv2018, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopPv2018{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找
func (obj *SssShopPv2018Mgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsSssShopPv2018, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopPv2018{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromVsYear 通过vs_year获取内容
func (obj *SssShopPv2018Mgr) GetFromVsYear(vsYear int) (results []*models.EsSssShopPv2018, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopPv2018{}).Where("`vs_year` = ?", vsYear).Find(&results).Error

	return
}

// GetBatchFromVsYear 批量查找
func (obj *SssShopPv2018Mgr) GetBatchFromVsYear(vsYears []int) (results []*models.EsSssShopPv2018, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopPv2018{}).Where("`vs_year` IN (?)", vsYears).Find(&results).Error

	return
}

// GetFromVsMonth 通过vs_month获取内容
func (obj *SssShopPv2018Mgr) GetFromVsMonth(vsMonth int) (results []*models.EsSssShopPv2018, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopPv2018{}).Where("`vs_month` = ?", vsMonth).Find(&results).Error

	return
}

// GetBatchFromVsMonth 批量查找
func (obj *SssShopPv2018Mgr) GetBatchFromVsMonth(vsMonths []int) (results []*models.EsSssShopPv2018, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopPv2018{}).Where("`vs_month` IN (?)", vsMonths).Find(&results).Error

	return
}

// GetFromVsDay 通过vs_day获取内容
func (obj *SssShopPv2018Mgr) GetFromVsDay(vsDay int) (results []*models.EsSssShopPv2018, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopPv2018{}).Where("`vs_day` = ?", vsDay).Find(&results).Error

	return
}

// GetBatchFromVsDay 批量查找
func (obj *SssShopPv2018Mgr) GetBatchFromVsDay(vsDays []int) (results []*models.EsSssShopPv2018, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopPv2018{}).Where("`vs_day` IN (?)", vsDays).Find(&results).Error

	return
}

// GetFromVsNum 通过vs_num获取内容
func (obj *SssShopPv2018Mgr) GetFromVsNum(vsNum int) (results []*models.EsSssShopPv2018, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopPv2018{}).Where("`vs_num` = ?", vsNum).Find(&results).Error

	return
}

// GetBatchFromVsNum 批量查找
func (obj *SssShopPv2018Mgr) GetBatchFromVsNum(vsNums []int) (results []*models.EsSssShopPv2018, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopPv2018{}).Where("`vs_num` IN (?)", vsNums).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
