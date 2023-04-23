package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsSssGoodsPvMgr struct {
	*_BaseMgr
}

// EsSssGoodsPvMgr open func
func EsSssGoodsPvMgr(db *gorm.DB) *_EsSssGoodsPvMgr {
	if db == nil {
		panic(fmt.Errorf("EsSssGoodsPvMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsSssGoodsPvMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_sss_goods_pv"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsSssGoodsPvMgr) GetTableName() string {
	return "es_sss_goods_pv"
}

// Reset 重置gorm会话
func (obj *_EsSssGoodsPvMgr) Reset() *_EsSssGoodsPvMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsSssGoodsPvMgr) Get() (result models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsSssGoodsPvMgr) Gets() (results []*models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsSssGoodsPvMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_EsSssGoodsPvMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSellerID seller_id获取 店铺id
func (obj *_EsSssGoodsPvMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithGoodsID goods_id获取 商品id
func (obj *_EsSssGoodsPvMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithGoodsName goods_name获取 商品名称
func (obj *_EsSssGoodsPvMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithVsYear vs_year获取 年份
func (obj *_EsSssGoodsPvMgr) WithVsYear(vsYear int) Option {
	return optionFunc(func(o *options) { o.query["vs_year"] = vsYear })
}

// WithVsNum vs_num获取 访问量
func (obj *_EsSssGoodsPvMgr) WithVsNum(vsNum int) Option {
	return optionFunc(func(o *options) { o.query["vs_num"] = vsNum })
}

// WithVsMonth vs_month获取 月份
func (obj *_EsSssGoodsPvMgr) WithVsMonth(vsMonth int) Option {
	return optionFunc(func(o *options) { o.query["vs_month"] = vsMonth })
}

// WithVsDay vs_day获取 天
func (obj *_EsSssGoodsPvMgr) WithVsDay(vsDay int) Option {
	return optionFunc(func(o *options) { o.query["vs_day"] = vsDay })
}

// GetByOption 功能选项模式获取
func (obj *_EsSssGoodsPvMgr) GetByOption(opts ...Option) (result models.EsSssGoodsPv, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsSssGoodsPvMgr) GetByOptions(opts ...Option) (results []*models.EsSssGoodsPv, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsSssGoodsPvMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSssGoodsPv, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where(options.query)
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
func (obj *_EsSssGoodsPvMgr) GetFromID(id int) (result models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_EsSssGoodsPvMgr) GetBatchFromID(ids []int) (results []*models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 店铺id
func (obj *_EsSssGoodsPvMgr) GetFromSellerID(sellerID int) (results []*models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 店铺id
func (obj *_EsSssGoodsPvMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品id
func (obj *_EsSssGoodsPvMgr) GetFromGoodsID(goodsID int) (results []*models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品id
func (obj *_EsSssGoodsPvMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *_EsSssGoodsPvMgr) GetFromGoodsName(goodsName string) (results []*models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *_EsSssGoodsPvMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromVsYear 通过vs_year获取内容 年份
func (obj *_EsSssGoodsPvMgr) GetFromVsYear(vsYear int) (results []*models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where("`vs_year` = ?", vsYear).Find(&results).Error

	return
}

// GetBatchFromVsYear 批量查找 年份
func (obj *_EsSssGoodsPvMgr) GetBatchFromVsYear(vsYears []int) (results []*models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where("`vs_year` IN (?)", vsYears).Find(&results).Error

	return
}

// GetFromVsNum 通过vs_num获取内容 访问量
func (obj *_EsSssGoodsPvMgr) GetFromVsNum(vsNum int) (results []*models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where("`vs_num` = ?", vsNum).Find(&results).Error

	return
}

// GetBatchFromVsNum 批量查找 访问量
func (obj *_EsSssGoodsPvMgr) GetBatchFromVsNum(vsNums []int) (results []*models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where("`vs_num` IN (?)", vsNums).Find(&results).Error

	return
}

// GetFromVsMonth 通过vs_month获取内容 月份
func (obj *_EsSssGoodsPvMgr) GetFromVsMonth(vsMonth int) (results []*models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where("`vs_month` = ?", vsMonth).Find(&results).Error

	return
}

// GetBatchFromVsMonth 批量查找 月份
func (obj *_EsSssGoodsPvMgr) GetBatchFromVsMonth(vsMonths []int) (results []*models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where("`vs_month` IN (?)", vsMonths).Find(&results).Error

	return
}

// GetFromVsDay 通过vs_day获取内容 天
func (obj *_EsSssGoodsPvMgr) GetFromVsDay(vsDay int) (results []*models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where("`vs_day` = ?", vsDay).Find(&results).Error

	return
}

// GetBatchFromVsDay 批量查找 天
func (obj *_EsSssGoodsPvMgr) GetBatchFromVsDay(vsDays []int) (results []*models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where("`vs_day` IN (?)", vsDays).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsSssGoodsPvMgr) FetchByPrimaryKey(id int) (result models.EsSssGoodsPv, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssGoodsPv{}).Where("`id` = ?", id).First(&result).Error

	return
}
