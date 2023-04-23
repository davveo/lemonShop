package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsSssOrderGoodsDataMgr struct {
	*_BaseMgr
}

// EsSssOrderGoodsDataMgr open func
func EsSssOrderGoodsDataMgr(db *gorm.DB) *_EsSssOrderGoodsDataMgr {
	if db == nil {
		panic(fmt.Errorf("EsSssOrderGoodsDataMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsSssOrderGoodsDataMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_sss_order_goods_data"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsSssOrderGoodsDataMgr) GetTableName() string {
	return "es_sss_order_goods_data"
}

// Reset 重置gorm会话
func (obj *_EsSssOrderGoodsDataMgr) Reset() *_EsSssOrderGoodsDataMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsSssOrderGoodsDataMgr) Get() (result models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsSssOrderGoodsDataMgr) Gets() (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsSssOrderGoodsDataMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_EsSssOrderGoodsDataMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderSn order_sn获取 订单编号
func (obj *_EsSssOrderGoodsDataMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithGoodsID goods_id获取 商品id
func (obj *_EsSssOrderGoodsDataMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithGoodsName goods_name获取 商品名称
func (obj *_EsSssOrderGoodsDataMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithGoodsNum goods_num获取 数量
func (obj *_EsSssOrderGoodsDataMgr) WithGoodsNum(goodsNum int) Option {
	return optionFunc(func(o *options) { o.query["goods_num"] = goodsNum })
}

// WithPrice price获取 商品单价
func (obj *_EsSssOrderGoodsDataMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithSubTotal sub_total获取 小计
func (obj *_EsSssOrderGoodsDataMgr) WithSubTotal(subTotal float64) Option {
	return optionFunc(func(o *options) { o.query["sub_total"] = subTotal })
}

// WithCategoryPath category_path获取 分类path
func (obj *_EsSssOrderGoodsDataMgr) WithCategoryPath(categoryPath string) Option {
	return optionFunc(func(o *options) { o.query["category_path"] = categoryPath })
}

// WithCategoryID category_id获取 分类id
func (obj *_EsSssOrderGoodsDataMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithCreateTime create_time获取 创建时间
func (obj *_EsSssOrderGoodsDataMgr) WithCreateTime(createTime int) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithIndustryID industry_id获取 行业id
func (obj *_EsSssOrderGoodsDataMgr) WithIndustryID(industryID int64) Option {
	return optionFunc(func(o *options) { o.query["industry_id"] = industryID })
}

// GetByOption 功能选项模式获取
func (obj *_EsSssOrderGoodsDataMgr) GetByOption(opts ...Option) (result models.EsSssOrderGoodsData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsSssOrderGoodsDataMgr) GetByOptions(opts ...Option) (results []*models.EsSssOrderGoodsData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsSssOrderGoodsDataMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSssOrderGoodsData, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where(options.query)
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

// GetFromID 通过id获取内容 id
func (obj *_EsSssOrderGoodsDataMgr) GetFromID(id int) (result models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_EsSssOrderGoodsDataMgr) GetBatchFromID(ids []int) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *_EsSssOrderGoodsDataMgr) GetFromOrderSn(orderSn string) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *_EsSssOrderGoodsDataMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品id
func (obj *_EsSssOrderGoodsDataMgr) GetFromGoodsID(goodsID int) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品id
func (obj *_EsSssOrderGoodsDataMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *_EsSssOrderGoodsDataMgr) GetFromGoodsName(goodsName string) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *_EsSssOrderGoodsDataMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromGoodsNum 通过goods_num获取内容 数量
func (obj *_EsSssOrderGoodsDataMgr) GetFromGoodsNum(goodsNum int) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`goods_num` = ?", goodsNum).Find(&results).Error

	return
}

// GetBatchFromGoodsNum 批量查找 数量
func (obj *_EsSssOrderGoodsDataMgr) GetBatchFromGoodsNum(goodsNums []int) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`goods_num` IN (?)", goodsNums).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 商品单价
func (obj *_EsSssOrderGoodsDataMgr) GetFromPrice(price float64) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 商品单价
func (obj *_EsSssOrderGoodsDataMgr) GetBatchFromPrice(prices []float64) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromSubTotal 通过sub_total获取内容 小计
func (obj *_EsSssOrderGoodsDataMgr) GetFromSubTotal(subTotal float64) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`sub_total` = ?", subTotal).Find(&results).Error

	return
}

// GetBatchFromSubTotal 批量查找 小计
func (obj *_EsSssOrderGoodsDataMgr) GetBatchFromSubTotal(subTotals []float64) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`sub_total` IN (?)", subTotals).Find(&results).Error

	return
}

// GetFromCategoryPath 通过category_path获取内容 分类path
func (obj *_EsSssOrderGoodsDataMgr) GetFromCategoryPath(categoryPath string) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`category_path` = ?", categoryPath).Find(&results).Error

	return
}

// GetBatchFromCategoryPath 批量查找 分类path
func (obj *_EsSssOrderGoodsDataMgr) GetBatchFromCategoryPath(categoryPaths []string) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`category_path` IN (?)", categoryPaths).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容 分类id
func (obj *_EsSssOrderGoodsDataMgr) GetFromCategoryID(categoryID int) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 分类id
func (obj *_EsSssOrderGoodsDataMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_EsSssOrderGoodsDataMgr) GetFromCreateTime(createTime int) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_EsSssOrderGoodsDataMgr) GetBatchFromCreateTime(createTimes []int) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromIndustryID 通过industry_id获取内容 行业id
func (obj *_EsSssOrderGoodsDataMgr) GetFromIndustryID(industryID int64) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`industry_id` = ?", industryID).Find(&results).Error

	return
}

// GetBatchFromIndustryID 批量查找 行业id
func (obj *_EsSssOrderGoodsDataMgr) GetBatchFromIndustryID(industryIDs []int64) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`industry_id` IN (?)", industryIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsSssOrderGoodsDataMgr) FetchByPrimaryKey(id int) (result models.EsSssOrderGoodsData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`id` = ?", id).First(&result).Error

	return
}
