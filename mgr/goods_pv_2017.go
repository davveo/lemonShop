package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type SssGoodsPv2017Mgr struct {
	*_BaseMgr
}

// NewSssGoodsPv2017Mgr open func
func NewSssGoodsPv2017Mgr(db db.Repo) *SssGoodsPv2017Mgr {
	if db == nil {
		panic(fmt.Errorf("NewSssGoodsPv2017Mgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &SssGoodsPv2017Mgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_category_spec"),
		wdb:       db.GetDbW().Table("es_category_spec"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *SssGoodsPv2017Mgr) GetTableName() string {
	return "es_sss_goods_pv_2017"
}

// Reset 重置gorm会话
func (obj *SssGoodsPv2017Mgr) Reset() *SssGoodsPv2017Mgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *SssGoodsPv2017Mgr) Get() (result models.EsSssGoodsPv2017, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *SssGoodsPv2017Mgr) Gets() (results []*models.EsSssGoodsPv2017, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *SssGoodsPv2017Mgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *SssGoodsPv2017Mgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSellerID seller_id获取
func (obj *SssGoodsPv2017Mgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithGoodsID goods_id获取
func (obj *SssGoodsPv2017Mgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithGoodsName goods_name获取
func (obj *SssGoodsPv2017Mgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithVsYear vs_year获取
func (obj *SssGoodsPv2017Mgr) WithVsYear(vsYear int) Option {
	return optionFunc(func(o *options) { o.query["vs_year"] = vsYear })
}

// WithVsMonth vs_month获取
func (obj *SssGoodsPv2017Mgr) WithVsMonth(vsMonth int) Option {
	return optionFunc(func(o *options) { o.query["vs_month"] = vsMonth })
}

// WithVsDay vs_day获取
func (obj *SssGoodsPv2017Mgr) WithVsDay(vsDay int) Option {
	return optionFunc(func(o *options) { o.query["vs_day"] = vsDay })
}

// WithVsNum vs_num获取
func (obj *SssGoodsPv2017Mgr) WithVsNum(vsNum int) Option {
	return optionFunc(func(o *options) { o.query["vs_num"] = vsNum })
}

// GetByOption 功能选项模式获取
func (obj *SssGoodsPv2017Mgr) GetByOption(opts ...Option) (result models.EsSssGoodsPv2017, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *SssGoodsPv2017Mgr) GetByOptions(opts ...Option) (results []*models.EsSssGoodsPv2017, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *SssGoodsPv2017Mgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSssGoodsPv2017, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where(options.query)
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
func (obj *SssGoodsPv2017Mgr) GetFromID(id int) (results []*models.EsSssGoodsPv2017, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// GetBatchFromID 批量查找
func (obj *SssGoodsPv2017Mgr) GetBatchFromID(ids []int) (results []*models.EsSssGoodsPv2017, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容
func (obj *SssGoodsPv2017Mgr) GetFromSellerID(sellerID int) (results []*models.EsSssGoodsPv2017, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找
func (obj *SssGoodsPv2017Mgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsSssGoodsPv2017, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容
func (obj *SssGoodsPv2017Mgr) GetFromGoodsID(goodsID int) (results []*models.EsSssGoodsPv2017, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找
func (obj *SssGoodsPv2017Mgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsSssGoodsPv2017, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容
func (obj *SssGoodsPv2017Mgr) GetFromGoodsName(goodsName string) (results []*models.EsSssGoodsPv2017, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找
func (obj *SssGoodsPv2017Mgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsSssGoodsPv2017, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromVsYear 通过vs_year获取内容
func (obj *SssGoodsPv2017Mgr) GetFromVsYear(vsYear int) (results []*models.EsSssGoodsPv2017, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where("`vs_year` = ?", vsYear).Find(&results).Error

	return
}

// GetBatchFromVsYear 批量查找
func (obj *SssGoodsPv2017Mgr) GetBatchFromVsYear(vsYears []int) (results []*models.EsSssGoodsPv2017, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where("`vs_year` IN (?)", vsYears).Find(&results).Error

	return
}

// GetFromVsMonth 通过vs_month获取内容
func (obj *SssGoodsPv2017Mgr) GetFromVsMonth(vsMonth int) (results []*models.EsSssGoodsPv2017, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where("`vs_month` = ?", vsMonth).Find(&results).Error

	return
}

// GetBatchFromVsMonth 批量查找
func (obj *SssGoodsPv2017Mgr) GetBatchFromVsMonth(vsMonths []int) (results []*models.EsSssGoodsPv2017, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where("`vs_month` IN (?)", vsMonths).Find(&results).Error

	return
}

// GetFromVsDay 通过vs_day获取内容
func (obj *SssGoodsPv2017Mgr) GetFromVsDay(vsDay int) (results []*models.EsSssGoodsPv2017, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where("`vs_day` = ?", vsDay).Find(&results).Error

	return
}

// GetBatchFromVsDay 批量查找
func (obj *SssGoodsPv2017Mgr) GetBatchFromVsDay(vsDays []int) (results []*models.EsSssGoodsPv2017, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where("`vs_day` IN (?)", vsDays).Find(&results).Error

	return
}

// GetFromVsNum 通过vs_num获取内容
func (obj *SssGoodsPv2017Mgr) GetFromVsNum(vsNum int) (results []*models.EsSssGoodsPv2017, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where("`vs_num` = ?", vsNum).Find(&results).Error

	return
}

// GetBatchFromVsNum 批量查找
func (obj *SssGoodsPv2017Mgr) GetBatchFromVsNum(vsNums []int) (results []*models.EsSssGoodsPv2017, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssGoodsPv2017{}).Where("`vs_num` IN (?)", vsNums).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
