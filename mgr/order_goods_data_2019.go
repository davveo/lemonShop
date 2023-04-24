package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type SssOrderGoodsData2019Mgr struct {
	*_BaseMgr
}

// NewSssOrderGoodsData2019Mgr open func
func NewSssOrderGoodsData2019Mgr(db db.Repo) *SssOrderGoodsData2019Mgr {
	if db == nil {
		panic(fmt.Errorf("NewSssOrderGoodsData2019Mgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &SssOrderGoodsData2019Mgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_sss_order_goods_data_2019"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *SssOrderGoodsData2019Mgr) GetTableName() string {
	return "es_sss_order_goods_data_2019"
}

// Reset 重置gorm会话
func (obj *SssOrderGoodsData2019Mgr) Reset() *SssOrderGoodsData2019Mgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *SssOrderGoodsData2019Mgr) Get() (result models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *SssOrderGoodsData2019Mgr) Gets() (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *SssOrderGoodsData2019Mgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *SssOrderGoodsData2019Mgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderSn order_sn获取
func (obj *SssOrderGoodsData2019Mgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithGoodsID goods_id获取
func (obj *SssOrderGoodsData2019Mgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithGoodsName goods_name获取
func (obj *SssOrderGoodsData2019Mgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithGoodsNum goods_num获取
func (obj *SssOrderGoodsData2019Mgr) WithGoodsNum(goodsNum int) Option {
	return optionFunc(func(o *options) { o.query["goods_num"] = goodsNum })
}

// WithPrice price获取
func (obj *SssOrderGoodsData2019Mgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithSubTotal sub_total获取
func (obj *SssOrderGoodsData2019Mgr) WithSubTotal(subTotal float64) Option {
	return optionFunc(func(o *options) { o.query["sub_total"] = subTotal })
}

// WithCategoryPath category_path获取
func (obj *SssOrderGoodsData2019Mgr) WithCategoryPath(categoryPath string) Option {
	return optionFunc(func(o *options) { o.query["category_path"] = categoryPath })
}

// WithCategoryID category_id获取
func (obj *SssOrderGoodsData2019Mgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithCreateTime create_time获取
func (obj *SssOrderGoodsData2019Mgr) WithCreateTime(createTime int) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithIndustryID industry_id获取
func (obj *SssOrderGoodsData2019Mgr) WithIndustryID(industryID int) Option {
	return optionFunc(func(o *options) { o.query["industry_id"] = industryID })
}

// GetByOption 功能选项模式获取
func (obj *SssOrderGoodsData2019Mgr) GetByOption(opts ...Option) (result models.EsSssOrderGoodsData2019, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *SssOrderGoodsData2019Mgr) GetByOptions(opts ...Option) (results []*models.EsSssOrderGoodsData2019, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *SssOrderGoodsData2019Mgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSssOrderGoodsData2019, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where(options.query)
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
func (obj *SssOrderGoodsData2019Mgr) GetFromID(id int) (result models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *SssOrderGoodsData2019Mgr) GetBatchFromID(ids []int) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容
func (obj *SssOrderGoodsData2019Mgr) GetFromOrderSn(orderSn string) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找
func (obj *SssOrderGoodsData2019Mgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容
func (obj *SssOrderGoodsData2019Mgr) GetFromGoodsID(goodsID int) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找
func (obj *SssOrderGoodsData2019Mgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容
func (obj *SssOrderGoodsData2019Mgr) GetFromGoodsName(goodsName string) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找
func (obj *SssOrderGoodsData2019Mgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromGoodsNum 通过goods_num获取内容
func (obj *SssOrderGoodsData2019Mgr) GetFromGoodsNum(goodsNum int) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`goods_num` = ?", goodsNum).Find(&results).Error

	return
}

// GetBatchFromGoodsNum 批量查找
func (obj *SssOrderGoodsData2019Mgr) GetBatchFromGoodsNum(goodsNums []int) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`goods_num` IN (?)", goodsNums).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容
func (obj *SssOrderGoodsData2019Mgr) GetFromPrice(price float64) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找
func (obj *SssOrderGoodsData2019Mgr) GetBatchFromPrice(prices []float64) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromSubTotal 通过sub_total获取内容
func (obj *SssOrderGoodsData2019Mgr) GetFromSubTotal(subTotal float64) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`sub_total` = ?", subTotal).Find(&results).Error

	return
}

// GetBatchFromSubTotal 批量查找
func (obj *SssOrderGoodsData2019Mgr) GetBatchFromSubTotal(subTotals []float64) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`sub_total` IN (?)", subTotals).Find(&results).Error

	return
}

// GetFromCategoryPath 通过category_path获取内容
func (obj *SssOrderGoodsData2019Mgr) GetFromCategoryPath(categoryPath string) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`category_path` = ?", categoryPath).Find(&results).Error

	return
}

// GetBatchFromCategoryPath 批量查找
func (obj *SssOrderGoodsData2019Mgr) GetBatchFromCategoryPath(categoryPaths []string) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`category_path` IN (?)", categoryPaths).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容
func (obj *SssOrderGoodsData2019Mgr) GetFromCategoryID(categoryID int) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找
func (obj *SssOrderGoodsData2019Mgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *SssOrderGoodsData2019Mgr) GetFromCreateTime(createTime int) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *SssOrderGoodsData2019Mgr) GetBatchFromCreateTime(createTimes []int) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromIndustryID 通过industry_id获取内容
func (obj *SssOrderGoodsData2019Mgr) GetFromIndustryID(industryID int) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`industry_id` = ?", industryID).Find(&results).Error

	return
}

// GetBatchFromIndustryID 批量查找
func (obj *SssOrderGoodsData2019Mgr) GetBatchFromIndustryID(industryIDs []int) (results []*models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`industry_id` IN (?)", industryIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *SssOrderGoodsData2019Mgr) FetchByPrimaryKey(id int) (result models.EsSssOrderGoodsData2019, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData2019{}).Where("`id` = ?", id).First(&result).Error

	return
}
