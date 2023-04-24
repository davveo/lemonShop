package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type AsGoodsMgr struct {
	*_BaseMgr
}

// NewAsGoodsMgr open func
func NewAsGoodsMgr(db db.Repo) *AsGoodsMgr {
	if db == nil {
		panic(fmt.Errorf("NewAsGoodsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &AsGoodsMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_category_spec"),
		wdb:       db.GetDbW().Table("es_category_spec"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *AsGoodsMgr) GetTableName() string {
	return "es_as_goods"
}

// Reset 重置gorm会话
func (obj *AsGoodsMgr) Reset() *AsGoodsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *AsGoodsMgr) Get() (result models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *AsGoodsMgr) Gets() (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *AsGoodsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *AsGoodsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithServiceSn service_sn获取 售后服务单号
func (obj *AsGoodsMgr) WithServiceSn(serviceSn string) Option {
	return optionFunc(func(o *options) { o.query["service_sn"] = serviceSn })
}

// WithGoodsID goods_id获取 商品ID
func (obj *AsGoodsMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithSkuID sku_id获取 商品SKUID
func (obj *AsGoodsMgr) WithSkuID(skuID int) Option {
	return optionFunc(func(o *options) { o.query["sku_id"] = skuID })
}

// WithShipNum ship_num获取 发货数量
func (obj *AsGoodsMgr) WithShipNum(shipNum int) Option {
	return optionFunc(func(o *options) { o.query["ship_num"] = shipNum })
}

// WithPrice price获取 商品成交价
func (obj *AsGoodsMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithReturnNum return_num获取 退还数量
func (obj *AsGoodsMgr) WithReturnNum(returnNum int) Option {
	return optionFunc(func(o *options) { o.query["return_num"] = returnNum })
}

// WithStorageNum storage_num获取 入库数量
func (obj *AsGoodsMgr) WithStorageNum(storageNum int) Option {
	return optionFunc(func(o *options) { o.query["storage_num"] = storageNum })
}

// WithGoodsName goods_name获取 商品名称
func (obj *AsGoodsMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithGoodsSn goods_sn获取 商品编号
func (obj *AsGoodsMgr) WithGoodsSn(goodsSn string) Option {
	return optionFunc(func(o *options) { o.query["goods_sn"] = goodsSn })
}

// WithGoodsImage goods_image获取 商品缩略图
func (obj *AsGoodsMgr) WithGoodsImage(goodsImage string) Option {
	return optionFunc(func(o *options) { o.query["goods_image"] = goodsImage })
}

// WithSpecJSON spec_json获取 商品规格信息
func (obj *AsGoodsMgr) WithSpecJSON(specJSON string) Option {
	return optionFunc(func(o *options) { o.query["spec_json"] = specJSON })
}

// GetByOption 功能选项模式获取
func (obj *AsGoodsMgr) GetByOption(opts ...Option) (result models.EsAsGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *AsGoodsMgr) GetByOptions(opts ...Option) (results []*models.EsAsGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *AsGoodsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsAsGoods, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键ID
func (obj *AsGoodsMgr) GetFromID(id int) (result models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *AsGoodsMgr) GetBatchFromID(ids []int) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromServiceSn 通过service_sn获取内容 售后服务单号
func (obj *AsGoodsMgr) GetFromServiceSn(serviceSn string) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`service_sn` = ?", serviceSn).Find(&results).Error

	return
}

// GetBatchFromServiceSn 批量查找 售后服务单号
func (obj *AsGoodsMgr) GetBatchFromServiceSn(serviceSns []string) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`service_sn` IN (?)", serviceSns).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品ID
func (obj *AsGoodsMgr) GetFromGoodsID(goodsID int) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品ID
func (obj *AsGoodsMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromSkuID 通过sku_id获取内容 商品SKUID
func (obj *AsGoodsMgr) GetFromSkuID(skuID int) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`sku_id` = ?", skuID).Find(&results).Error

	return
}

// GetBatchFromSkuID 批量查找 商品SKUID
func (obj *AsGoodsMgr) GetBatchFromSkuID(skuIDs []int) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`sku_id` IN (?)", skuIDs).Find(&results).Error

	return
}

// GetFromShipNum 通过ship_num获取内容 发货数量
func (obj *AsGoodsMgr) GetFromShipNum(shipNum int) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`ship_num` = ?", shipNum).Find(&results).Error

	return
}

// GetBatchFromShipNum 批量查找 发货数量
func (obj *AsGoodsMgr) GetBatchFromShipNum(shipNums []int) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`ship_num` IN (?)", shipNums).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 商品成交价
func (obj *AsGoodsMgr) GetFromPrice(price float64) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 商品成交价
func (obj *AsGoodsMgr) GetBatchFromPrice(prices []float64) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromReturnNum 通过return_num获取内容 退还数量
func (obj *AsGoodsMgr) GetFromReturnNum(returnNum int) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`return_num` = ?", returnNum).Find(&results).Error

	return
}

// GetBatchFromReturnNum 批量查找 退还数量
func (obj *AsGoodsMgr) GetBatchFromReturnNum(returnNums []int) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`return_num` IN (?)", returnNums).Find(&results).Error

	return
}

// GetFromStorageNum 通过storage_num获取内容 入库数量
func (obj *AsGoodsMgr) GetFromStorageNum(storageNum int) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`storage_num` = ?", storageNum).Find(&results).Error

	return
}

// GetBatchFromStorageNum 批量查找 入库数量
func (obj *AsGoodsMgr) GetBatchFromStorageNum(storageNums []int) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`storage_num` IN (?)", storageNums).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *AsGoodsMgr) GetFromGoodsName(goodsName string) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *AsGoodsMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromGoodsSn 通过goods_sn获取内容 商品编号
func (obj *AsGoodsMgr) GetFromGoodsSn(goodsSn string) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`goods_sn` = ?", goodsSn).Find(&results).Error

	return
}

// GetBatchFromGoodsSn 批量查找 商品编号
func (obj *AsGoodsMgr) GetBatchFromGoodsSn(goodsSns []string) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`goods_sn` IN (?)", goodsSns).Find(&results).Error

	return
}

// GetFromGoodsImage 通过goods_image获取内容 商品缩略图
func (obj *AsGoodsMgr) GetFromGoodsImage(goodsImage string) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`goods_image` = ?", goodsImage).Find(&results).Error

	return
}

// GetBatchFromGoodsImage 批量查找 商品缩略图
func (obj *AsGoodsMgr) GetBatchFromGoodsImage(goodsImages []string) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`goods_image` IN (?)", goodsImages).Find(&results).Error

	return
}

// GetFromSpecJSON 通过spec_json获取内容 商品规格信息
func (obj *AsGoodsMgr) GetFromSpecJSON(specJSON string) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`spec_json` = ?", specJSON).Find(&results).Error

	return
}

// GetBatchFromSpecJSON 批量查找 商品规格信息
func (obj *AsGoodsMgr) GetBatchFromSpecJSON(specJSONs []string) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`spec_json` IN (?)", specJSONs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *AsGoodsMgr) FetchByPrimaryKey(id int) (result models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByEsIndexID  获取多个内容
func (obj *AsGoodsMgr) FetchIndexByEsIndexID(id int) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// FetchIndexByEsIndexServiceSn  获取多个内容
func (obj *AsGoodsMgr) FetchIndexByEsIndexServiceSn(serviceSn string) (results []*models.EsAsGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`service_sn` = ?", serviceSn).Find(&results).Error

	return
}
