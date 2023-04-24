package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type SssShopDataMgr struct {
	*_BaseMgr
}

// NewSssShopDataMgr open func
func NewSssShopDataMgr(db db.Repo) *SssShopDataMgr {
	if db == nil {
		panic(fmt.Errorf("NewSssShopDataMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &SssShopDataMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_sss_shop_data"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *SssShopDataMgr) GetTableName() string {
	return "es_sss_shop_data"
}

// Reset 重置gorm会话
func (obj *SssShopDataMgr) Reset() *SssShopDataMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *SssShopDataMgr) Get() (result models.EsSssShopData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopData{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *SssShopDataMgr) Gets() (results []*models.EsSssShopData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopData{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *SssShopDataMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopData{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *SssShopDataMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSellerID seller_id获取 店铺id
func (obj *SssShopDataMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithSellerName seller_name获取 店铺名称
func (obj *SssShopDataMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithFavoriteNum favorite_num获取 店铺被收藏数量
func (obj *SssShopDataMgr) WithFavoriteNum(favoriteNum int) Option {
	return optionFunc(func(o *options) { o.query["favorite_num"] = favoriteNum })
}

// WithShopDisable shop_disable获取 店铺是否开启
func (obj *SssShopDataMgr) WithShopDisable(shopDisable string) Option {
	return optionFunc(func(o *options) { o.query["shop_disable"] = shopDisable })
}

// GetByOption 功能选项模式获取
func (obj *SssShopDataMgr) GetByOption(opts ...Option) (result models.EsSssShopData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopData{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *SssShopDataMgr) GetByOptions(opts ...Option) (results []*models.EsSssShopData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopData{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *SssShopDataMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSssShopData, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopData{}).Where(options.query)
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
func (obj *SssShopDataMgr) GetFromID(id int) (result models.EsSssShopData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopData{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *SssShopDataMgr) GetBatchFromID(ids []int) (results []*models.EsSssShopData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopData{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 店铺id
func (obj *SssShopDataMgr) GetFromSellerID(sellerID int) (results []*models.EsSssShopData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopData{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 店铺id
func (obj *SssShopDataMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsSssShopData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopData{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 店铺名称
func (obj *SssShopDataMgr) GetFromSellerName(sellerName string) (results []*models.EsSssShopData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopData{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 店铺名称
func (obj *SssShopDataMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsSssShopData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopData{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromFavoriteNum 通过favorite_num获取内容 店铺被收藏数量
func (obj *SssShopDataMgr) GetFromFavoriteNum(favoriteNum int) (results []*models.EsSssShopData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopData{}).Where("`favorite_num` = ?", favoriteNum).Find(&results).Error

	return
}

// GetBatchFromFavoriteNum 批量查找 店铺被收藏数量
func (obj *SssShopDataMgr) GetBatchFromFavoriteNum(favoriteNums []int) (results []*models.EsSssShopData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopData{}).Where("`favorite_num` IN (?)", favoriteNums).Find(&results).Error

	return
}

// GetFromShopDisable 通过shop_disable获取内容 店铺是否开启
func (obj *SssShopDataMgr) GetFromShopDisable(shopDisable string) (results []*models.EsSssShopData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopData{}).Where("`shop_disable` = ?", shopDisable).Find(&results).Error

	return
}

// GetBatchFromShopDisable 批量查找 店铺是否开启
func (obj *SssShopDataMgr) GetBatchFromShopDisable(shopDisables []string) (results []*models.EsSssShopData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopData{}).Where("`shop_disable` IN (?)", shopDisables).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *SssShopDataMgr) FetchByPrimaryKey(id int) (result models.EsSssShopData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssShopData{}).Where("`id` = ?", id).First(&result).Error

	return
}
