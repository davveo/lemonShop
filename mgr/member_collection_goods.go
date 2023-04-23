package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsMemberCollectionGoodsMgr struct {
	*_BaseMgr
}

// EsMemberCollectionGoodsMgr open func
func EsMemberCollectionGoodsMgr(db *gorm.DB) *_EsMemberCollectionGoodsMgr {
	if db == nil {
		panic(fmt.Errorf("EsMemberCollectionGoodsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsMemberCollectionGoodsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_member_collection_goods"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsMemberCollectionGoodsMgr) GetTableName() string {
	return "es_member_collection_goods"
}

// Reset 重置gorm会话
func (obj *_EsMemberCollectionGoodsMgr) Reset() *_EsMemberCollectionGoodsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsMemberCollectionGoodsMgr) Get() (result models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsMemberCollectionGoodsMgr) Gets() (results []*models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsMemberCollectionGoodsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *_EsMemberCollectionGoodsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员ID
func (obj *_EsMemberCollectionGoodsMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithGoodsID goods_id获取 收藏商品ID
func (obj *_EsMemberCollectionGoodsMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithCreateTime create_time获取 收藏商品时间
func (obj *_EsMemberCollectionGoodsMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithGoodsName goods_name获取 商品名称
func (obj *_EsMemberCollectionGoodsMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithGoodsPrice goods_price获取 商品价格
func (obj *_EsMemberCollectionGoodsMgr) WithGoodsPrice(goodsPrice float64) Option {
	return optionFunc(func(o *options) { o.query["goods_price"] = goodsPrice })
}

// WithGoodsSn goods_sn获取 商品编号
func (obj *_EsMemberCollectionGoodsMgr) WithGoodsSn(goodsSn string) Option {
	return optionFunc(func(o *options) { o.query["goods_sn"] = goodsSn })
}

// WithGoodsImg goods_img获取 商品图片
func (obj *_EsMemberCollectionGoodsMgr) WithGoodsImg(goodsImg string) Option {
	return optionFunc(func(o *options) { o.query["goods_img"] = goodsImg })
}

// WithShopID shop_id获取 店铺id
func (obj *_EsMemberCollectionGoodsMgr) WithShopID(shopID int) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// GetByOption 功能选项模式获取
func (obj *_EsMemberCollectionGoodsMgr) GetByOption(opts ...Option) (result models.EsMemberCollectionGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsMemberCollectionGoodsMgr) GetByOptions(opts ...Option) (results []*models.EsMemberCollectionGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsMemberCollectionGoodsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsMemberCollectionGoods, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where(options.query)
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
func (obj *_EsMemberCollectionGoodsMgr) GetFromID(id int) (result models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *_EsMemberCollectionGoodsMgr) GetBatchFromID(ids []int) (results []*models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员ID
func (obj *_EsMemberCollectionGoodsMgr) GetFromMemberID(memberID int) (results []*models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员ID
func (obj *_EsMemberCollectionGoodsMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 收藏商品ID
func (obj *_EsMemberCollectionGoodsMgr) GetFromGoodsID(goodsID int) (results []*models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 收藏商品ID
func (obj *_EsMemberCollectionGoodsMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 收藏商品时间
func (obj *_EsMemberCollectionGoodsMgr) GetFromCreateTime(createTime int64) (results []*models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 收藏商品时间
func (obj *_EsMemberCollectionGoodsMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *_EsMemberCollectionGoodsMgr) GetFromGoodsName(goodsName string) (results []*models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *_EsMemberCollectionGoodsMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromGoodsPrice 通过goods_price获取内容 商品价格
func (obj *_EsMemberCollectionGoodsMgr) GetFromGoodsPrice(goodsPrice float64) (results []*models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`goods_price` = ?", goodsPrice).Find(&results).Error

	return
}

// GetBatchFromGoodsPrice 批量查找 商品价格
func (obj *_EsMemberCollectionGoodsMgr) GetBatchFromGoodsPrice(goodsPrices []float64) (results []*models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`goods_price` IN (?)", goodsPrices).Find(&results).Error

	return
}

// GetFromGoodsSn 通过goods_sn获取内容 商品编号
func (obj *_EsMemberCollectionGoodsMgr) GetFromGoodsSn(goodsSn string) (results []*models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`goods_sn` = ?", goodsSn).Find(&results).Error

	return
}

// GetBatchFromGoodsSn 批量查找 商品编号
func (obj *_EsMemberCollectionGoodsMgr) GetBatchFromGoodsSn(goodsSns []string) (results []*models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`goods_sn` IN (?)", goodsSns).Find(&results).Error

	return
}

// GetFromGoodsImg 通过goods_img获取内容 商品图片
func (obj *_EsMemberCollectionGoodsMgr) GetFromGoodsImg(goodsImg string) (results []*models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`goods_img` = ?", goodsImg).Find(&results).Error

	return
}

// GetBatchFromGoodsImg 批量查找 商品图片
func (obj *_EsMemberCollectionGoodsMgr) GetBatchFromGoodsImg(goodsImgs []string) (results []*models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`goods_img` IN (?)", goodsImgs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_EsMemberCollectionGoodsMgr) GetFromShopID(shopID int) (results []*models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_EsMemberCollectionGoodsMgr) GetBatchFromShopID(shopIDs []int) (results []*models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsMemberCollectionGoodsMgr) FetchByPrimaryKey(id int) (result models.EsMemberCollectionGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionGoods{}).Where("`id` = ?", id).First(&result).Error

	return
}
