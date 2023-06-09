package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type xchangeMgr struct {
	*_BaseMgr
}

// EsExchangeMgr open func
func EsExchangeMgr(db db.Repo) *xchangeMgr {
	if db == nil {
		panic(fmt.Errorf("EsExchangeMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &xchangeMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_exchange"), wdb: db.GetDbW().Table("es_exchange"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *xchangeMgr) GetTableName() string {
	return "es_exchange"
}

// Reset 重置gorm会话
func (obj *xchangeMgr) Reset() *xchangeMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *xchangeMgr) Get() (result models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *xchangeMgr) Gets() (results []*models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *xchangeMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithExchangeID exchange_id获取 主键
func (obj *xchangeMgr) WithExchangeID(exchangeID int) Option {
	return optionFunc(func(o *options) { o.query["exchange_id"] = exchangeID })
}

// WithGoodsID goods_id获取 商品id
func (obj *xchangeMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithCategoryID category_id获取 商品所属积分分类
func (obj *xchangeMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithEnableExchange enable_exchange获取 是否允许兑换
func (obj *xchangeMgr) WithEnableExchange(enableExchange int) Option {
	return optionFunc(func(o *options) { o.query["enable_exchange"] = enableExchange })
}

// WithExchangeMoney exchange_money获取 兑换所需金额
func (obj *xchangeMgr) WithExchangeMoney(exchangeMoney float64) Option {
	return optionFunc(func(o *options) { o.query["exchange_money"] = exchangeMoney })
}

// WithExchangePoint exchange_point获取 兑换所需积分
func (obj *xchangeMgr) WithExchangePoint(exchangePoint int) Option {
	return optionFunc(func(o *options) { o.query["exchange_point"] = exchangePoint })
}

// WithGoodsName goods_name获取 商品名称
func (obj *xchangeMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithGoodsPrice goods_price获取 商品原价
func (obj *xchangeMgr) WithGoodsPrice(goodsPrice float64) Option {
	return optionFunc(func(o *options) { o.query["goods_price"] = goodsPrice })
}

// WithGoodsImg goods_img获取 商品图片
func (obj *xchangeMgr) WithGoodsImg(goodsImg string) Option {
	return optionFunc(func(o *options) { o.query["goods_img"] = goodsImg })
}

// GetByOption 功能选项模式获取
func (obj *xchangeMgr) GetByOption(opts ...Option) (result models.EsExchange, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *xchangeMgr) GetByOptions(opts ...Option) (results []*models.EsExchange, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *xchangeMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsExchange, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where(options.query)
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

// GetFromExchangeID 通过exchange_id获取内容 主键
func (obj *xchangeMgr) GetFromExchangeID(exchangeID int) (result models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`exchange_id` = ?", exchangeID).First(&result).Error

	return
}

// GetBatchFromExchangeID 批量查找 主键
func (obj *xchangeMgr) GetBatchFromExchangeID(exchangeIDs []int) (results []*models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`exchange_id` IN (?)", exchangeIDs).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品id
func (obj *xchangeMgr) GetFromGoodsID(goodsID int) (results []*models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品id
func (obj *xchangeMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容 商品所属积分分类
func (obj *xchangeMgr) GetFromCategoryID(categoryID int) (results []*models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 商品所属积分分类
func (obj *xchangeMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromEnableExchange 通过enable_exchange获取内容 是否允许兑换
func (obj *xchangeMgr) GetFromEnableExchange(enableExchange int) (results []*models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`enable_exchange` = ?", enableExchange).Find(&results).Error

	return
}

// GetBatchFromEnableExchange 批量查找 是否允许兑换
func (obj *xchangeMgr) GetBatchFromEnableExchange(enableExchanges []int) (results []*models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`enable_exchange` IN (?)", enableExchanges).Find(&results).Error

	return
}

// GetFromExchangeMoney 通过exchange_money获取内容 兑换所需金额
func (obj *xchangeMgr) GetFromExchangeMoney(exchangeMoney float64) (results []*models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`exchange_money` = ?", exchangeMoney).Find(&results).Error

	return
}

// GetBatchFromExchangeMoney 批量查找 兑换所需金额
func (obj *xchangeMgr) GetBatchFromExchangeMoney(exchangeMoneys []float64) (results []*models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`exchange_money` IN (?)", exchangeMoneys).Find(&results).Error

	return
}

// GetFromExchangePoint 通过exchange_point获取内容 兑换所需积分
func (obj *xchangeMgr) GetFromExchangePoint(exchangePoint int) (results []*models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`exchange_point` = ?", exchangePoint).Find(&results).Error

	return
}

// GetBatchFromExchangePoint 批量查找 兑换所需积分
func (obj *xchangeMgr) GetBatchFromExchangePoint(exchangePoints []int) (results []*models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`exchange_point` IN (?)", exchangePoints).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *xchangeMgr) GetFromGoodsName(goodsName string) (results []*models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *xchangeMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromGoodsPrice 通过goods_price获取内容 商品原价
func (obj *xchangeMgr) GetFromGoodsPrice(goodsPrice float64) (results []*models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`goods_price` = ?", goodsPrice).Find(&results).Error

	return
}

// GetBatchFromGoodsPrice 批量查找 商品原价
func (obj *xchangeMgr) GetBatchFromGoodsPrice(goodsPrices []float64) (results []*models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`goods_price` IN (?)", goodsPrices).Find(&results).Error

	return
}

// GetFromGoodsImg 通过goods_img获取内容 商品图片
func (obj *xchangeMgr) GetFromGoodsImg(goodsImg string) (results []*models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`goods_img` = ?", goodsImg).Find(&results).Error

	return
}

// GetBatchFromGoodsImg 批量查找 商品图片
func (obj *xchangeMgr) GetBatchFromGoodsImg(goodsImgs []string) (results []*models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`goods_img` IN (?)", goodsImgs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *xchangeMgr) FetchByPrimaryKey(exchangeID int) (result models.EsExchange, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsExchange{}).Where("`exchange_id` = ?", exchangeID).First(&result).Error

	return
}
