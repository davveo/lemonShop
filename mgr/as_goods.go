package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsAsGoodsMgr struct {
	*_BaseMgr
}

// EsAsGoodsMgr open func
func EsAsGoodsMgr(db *gorm.DB) *_EsAsGoodsMgr {
	if db == nil {
		panic(fmt.Errorf("EsAsGoodsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsAsGoodsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_as_goods"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsAsGoodsMgr) GetTableName() string {
	return "es_as_goods"
}

// Reset 重置gorm会话
func (obj *_EsAsGoodsMgr) Reset() *_EsAsGoodsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsAsGoodsMgr) Get() (result models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsAsGoodsMgr) Gets() (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsAsGoodsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *_EsAsGoodsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithServiceSn service_sn获取 售后服务单号
func (obj *_EsAsGoodsMgr) WithServiceSn(serviceSn string) Option {
	return optionFunc(func(o *options) { o.query["service_sn"] = serviceSn })
}

// WithGoodsID goods_id获取 商品ID
func (obj *_EsAsGoodsMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithSkuID sku_id获取 商品SKUID
func (obj *_EsAsGoodsMgr) WithSkuID(skuID int) Option {
	return optionFunc(func(o *options) { o.query["sku_id"] = skuID })
}

// WithShipNum ship_num获取 发货数量
func (obj *_EsAsGoodsMgr) WithShipNum(shipNum int) Option {
	return optionFunc(func(o *options) { o.query["ship_num"] = shipNum })
}

// WithPrice price获取 商品成交价
func (obj *_EsAsGoodsMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithReturnNum return_num获取 退还数量
func (obj *_EsAsGoodsMgr) WithReturnNum(returnNum int) Option {
	return optionFunc(func(o *options) { o.query["return_num"] = returnNum })
}

// WithStorageNum storage_num获取 入库数量
func (obj *_EsAsGoodsMgr) WithStorageNum(storageNum int) Option {
	return optionFunc(func(o *options) { o.query["storage_num"] = storageNum })
}

// WithGoodsName goods_name获取 商品名称
func (obj *_EsAsGoodsMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithGoodsSn goods_sn获取 商品编号
func (obj *_EsAsGoodsMgr) WithGoodsSn(goodsSn string) Option {
	return optionFunc(func(o *options) { o.query["goods_sn"] = goodsSn })
}

// WithGoodsImage goods_image获取 商品缩略图
func (obj *_EsAsGoodsMgr) WithGoodsImage(goodsImage string) Option {
	return optionFunc(func(o *options) { o.query["goods_image"] = goodsImage })
}

// WithSpecJSON spec_json获取 商品规格信息
func (obj *_EsAsGoodsMgr) WithSpecJSON(specJSON string) Option {
	return optionFunc(func(o *options) { o.query["spec_json"] = specJSON })
}

// GetByOption 功能选项模式获取
func (obj *_EsAsGoodsMgr) GetByOption(opts ...Option) (result models.EsAsGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsAsGoodsMgr) GetByOptions(opts ...Option) (results []*models.EsAsGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsAsGoodsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsAsGoods, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where(options.query)
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
func (obj *_EsAsGoodsMgr) GetFromID(id int) (result models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *_EsAsGoodsMgr) GetBatchFromID(ids []int) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromServiceSn 通过service_sn获取内容 售后服务单号
func (obj *_EsAsGoodsMgr) GetFromServiceSn(serviceSn string) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`service_sn` = ?", serviceSn).Find(&results).Error

	return
}

// GetBatchFromServiceSn 批量查找 售后服务单号
func (obj *_EsAsGoodsMgr) GetBatchFromServiceSn(serviceSns []string) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`service_sn` IN (?)", serviceSns).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品ID
func (obj *_EsAsGoodsMgr) GetFromGoodsID(goodsID int) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品ID
func (obj *_EsAsGoodsMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromSkuID 通过sku_id获取内容 商品SKUID
func (obj *_EsAsGoodsMgr) GetFromSkuID(skuID int) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`sku_id` = ?", skuID).Find(&results).Error

	return
}

// GetBatchFromSkuID 批量查找 商品SKUID
func (obj *_EsAsGoodsMgr) GetBatchFromSkuID(skuIDs []int) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`sku_id` IN (?)", skuIDs).Find(&results).Error

	return
}

// GetFromShipNum 通过ship_num获取内容 发货数量
func (obj *_EsAsGoodsMgr) GetFromShipNum(shipNum int) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`ship_num` = ?", shipNum).Find(&results).Error

	return
}

// GetBatchFromShipNum 批量查找 发货数量
func (obj *_EsAsGoodsMgr) GetBatchFromShipNum(shipNums []int) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`ship_num` IN (?)", shipNums).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 商品成交价
func (obj *_EsAsGoodsMgr) GetFromPrice(price float64) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 商品成交价
func (obj *_EsAsGoodsMgr) GetBatchFromPrice(prices []float64) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromReturnNum 通过return_num获取内容 退还数量
func (obj *_EsAsGoodsMgr) GetFromReturnNum(returnNum int) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`return_num` = ?", returnNum).Find(&results).Error

	return
}

// GetBatchFromReturnNum 批量查找 退还数量
func (obj *_EsAsGoodsMgr) GetBatchFromReturnNum(returnNums []int) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`return_num` IN (?)", returnNums).Find(&results).Error

	return
}

// GetFromStorageNum 通过storage_num获取内容 入库数量
func (obj *_EsAsGoodsMgr) GetFromStorageNum(storageNum int) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`storage_num` = ?", storageNum).Find(&results).Error

	return
}

// GetBatchFromStorageNum 批量查找 入库数量
func (obj *_EsAsGoodsMgr) GetBatchFromStorageNum(storageNums []int) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`storage_num` IN (?)", storageNums).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *_EsAsGoodsMgr) GetFromGoodsName(goodsName string) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *_EsAsGoodsMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromGoodsSn 通过goods_sn获取内容 商品编号
func (obj *_EsAsGoodsMgr) GetFromGoodsSn(goodsSn string) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`goods_sn` = ?", goodsSn).Find(&results).Error

	return
}

// GetBatchFromGoodsSn 批量查找 商品编号
func (obj *_EsAsGoodsMgr) GetBatchFromGoodsSn(goodsSns []string) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`goods_sn` IN (?)", goodsSns).Find(&results).Error

	return
}

// GetFromGoodsImage 通过goods_image获取内容 商品缩略图
func (obj *_EsAsGoodsMgr) GetFromGoodsImage(goodsImage string) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`goods_image` = ?", goodsImage).Find(&results).Error

	return
}

// GetBatchFromGoodsImage 批量查找 商品缩略图
func (obj *_EsAsGoodsMgr) GetBatchFromGoodsImage(goodsImages []string) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`goods_image` IN (?)", goodsImages).Find(&results).Error

	return
}

// GetFromSpecJSON 通过spec_json获取内容 商品规格信息
func (obj *_EsAsGoodsMgr) GetFromSpecJSON(specJSON string) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`spec_json` = ?", specJSON).Find(&results).Error

	return
}

// GetBatchFromSpecJSON 批量查找 商品规格信息
func (obj *_EsAsGoodsMgr) GetBatchFromSpecJSON(specJSONs []string) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`spec_json` IN (?)", specJSONs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsAsGoodsMgr) FetchByPrimaryKey(id int) (result models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByEsIndexID  获取多个内容
func (obj *_EsAsGoodsMgr) FetchIndexByEsIndexID(id int) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// FetchIndexByEsIndexServiceSn  获取多个内容
func (obj *_EsAsGoodsMgr) FetchIndexByEsIndexServiceSn(serviceSn string) (results []*models.EsAsGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsGoods{}).Where("`service_sn` = ?", serviceSn).Find(&results).Error

	return
}
