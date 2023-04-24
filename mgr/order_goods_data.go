package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type SssOrderGoodsDataMgr struct {
	*_BaseMgr
}

// NewSssOrderGoodsDataMgr open func
func NewSssOrderGoodsDataMgr(db db.Repo) *SssOrderGoodsDataMgr {
	if db == nil {
		panic(fmt.Errorf("NewSssOrderGoodsDataMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &SssOrderGoodsDataMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_sss_order_goods_data"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *SssOrderGoodsDataMgr) GetTableName() string {
	return "es_sss_order_goods_data"
}

// Reset 重置gorm会话
func (obj *SssOrderGoodsDataMgr) Reset() *SssOrderGoodsDataMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *SssOrderGoodsDataMgr) Get() (result models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *SssOrderGoodsDataMgr) Gets() (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *SssOrderGoodsDataMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *SssOrderGoodsDataMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderSn order_sn获取 订单编号
func (obj *SssOrderGoodsDataMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithGoodsID goods_id获取 商品id
func (obj *SssOrderGoodsDataMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithGoodsName goods_name获取 商品名称
func (obj *SssOrderGoodsDataMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithGoodsNum goods_num获取 数量
func (obj *SssOrderGoodsDataMgr) WithGoodsNum(goodsNum int) Option {
	return optionFunc(func(o *options) { o.query["goods_num"] = goodsNum })
}

// WithPrice price获取 商品单价
func (obj *SssOrderGoodsDataMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithSubTotal sub_total获取 小计
func (obj *SssOrderGoodsDataMgr) WithSubTotal(subTotal float64) Option {
	return optionFunc(func(o *options) { o.query["sub_total"] = subTotal })
}

// WithCategoryPath category_path获取 分类path
func (obj *SssOrderGoodsDataMgr) WithCategoryPath(categoryPath string) Option {
	return optionFunc(func(o *options) { o.query["category_path"] = categoryPath })
}

// WithCategoryID category_id获取 分类id
func (obj *SssOrderGoodsDataMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithCreateTime create_time获取 创建时间
func (obj *SssOrderGoodsDataMgr) WithCreateTime(createTime int) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithIndustryID industry_id获取 行业id
func (obj *SssOrderGoodsDataMgr) WithIndustryID(industryID int64) Option {
	return optionFunc(func(o *options) { o.query["industry_id"] = industryID })
}

// GetByOption 功能选项模式获取
func (obj *SssOrderGoodsDataMgr) GetByOption(opts ...Option) (result models.EsSssOrderGoodsData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *SssOrderGoodsDataMgr) GetByOptions(opts ...Option) (results []*models.EsSssOrderGoodsData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *SssOrderGoodsDataMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSssOrderGoodsData, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where(options.query)
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
func (obj *SssOrderGoodsDataMgr) GetFromID(id int) (result models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *SssOrderGoodsDataMgr) GetBatchFromID(ids []int) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *SssOrderGoodsDataMgr) GetFromOrderSn(orderSn string) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *SssOrderGoodsDataMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品id
func (obj *SssOrderGoodsDataMgr) GetFromGoodsID(goodsID int) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品id
func (obj *SssOrderGoodsDataMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *SssOrderGoodsDataMgr) GetFromGoodsName(goodsName string) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *SssOrderGoodsDataMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromGoodsNum 通过goods_num获取内容 数量
func (obj *SssOrderGoodsDataMgr) GetFromGoodsNum(goodsNum int) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`goods_num` = ?", goodsNum).Find(&results).Error

	return
}

// GetBatchFromGoodsNum 批量查找 数量
func (obj *SssOrderGoodsDataMgr) GetBatchFromGoodsNum(goodsNums []int) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`goods_num` IN (?)", goodsNums).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 商品单价
func (obj *SssOrderGoodsDataMgr) GetFromPrice(price float64) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 商品单价
func (obj *SssOrderGoodsDataMgr) GetBatchFromPrice(prices []float64) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromSubTotal 通过sub_total获取内容 小计
func (obj *SssOrderGoodsDataMgr) GetFromSubTotal(subTotal float64) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`sub_total` = ?", subTotal).Find(&results).Error

	return
}

// GetBatchFromSubTotal 批量查找 小计
func (obj *SssOrderGoodsDataMgr) GetBatchFromSubTotal(subTotals []float64) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`sub_total` IN (?)", subTotals).Find(&results).Error

	return
}

// GetFromCategoryPath 通过category_path获取内容 分类path
func (obj *SssOrderGoodsDataMgr) GetFromCategoryPath(categoryPath string) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`category_path` = ?", categoryPath).Find(&results).Error

	return
}

// GetBatchFromCategoryPath 批量查找 分类path
func (obj *SssOrderGoodsDataMgr) GetBatchFromCategoryPath(categoryPaths []string) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`category_path` IN (?)", categoryPaths).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容 分类id
func (obj *SssOrderGoodsDataMgr) GetFromCategoryID(categoryID int) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 分类id
func (obj *SssOrderGoodsDataMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *SssOrderGoodsDataMgr) GetFromCreateTime(createTime int) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *SssOrderGoodsDataMgr) GetBatchFromCreateTime(createTimes []int) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromIndustryID 通过industry_id获取内容 行业id
func (obj *SssOrderGoodsDataMgr) GetFromIndustryID(industryID int64) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`industry_id` = ?", industryID).Find(&results).Error

	return
}

// GetBatchFromIndustryID 批量查找 行业id
func (obj *SssOrderGoodsDataMgr) GetBatchFromIndustryID(industryIDs []int64) (results []*models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`industry_id` IN (?)", industryIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *SssOrderGoodsDataMgr) FetchByPrimaryKey(id int) (result models.EsSssOrderGoodsData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssOrderGoodsData{}).Where("`id` = ?", id).First(&result).Error

	return
}
