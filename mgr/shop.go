package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type ShopMgr struct {
	*_BaseMgr
}

// NewShopMgr open func
func NewShopMgr(db db.Repo) *ShopMgr {
	if db == nil {
		panic(fmt.Errorf("NewShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &ShopMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_shop"), wdb: db.GetDbW().Table("es_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *ShopMgr) GetTableName() string {
	return "es_shop"
}

// Reset 重置gorm会话
func (obj *ShopMgr) Reset() *ShopMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *ShopMgr) Get() (result models.EsShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *ShopMgr) Gets() (results []*models.EsShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *ShopMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithShopID shop_id获取 店铺Id
func (obj *ShopMgr) WithShopID(shopID int) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithMemberID member_id获取 会员Id
func (obj *ShopMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMemberName member_name获取 会员名称
func (obj *ShopMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithShopName shop_name获取 店铺名称
func (obj *ShopMgr) WithShopName(shopName string) Option {
	return optionFunc(func(o *options) { o.query["shop_name"] = shopName })
}

// WithShopDisable shop_disable获取 店铺状态
func (obj *ShopMgr) WithShopDisable(shopDisable string) Option {
	return optionFunc(func(o *options) { o.query["shop_disable"] = shopDisable })
}

// WithShopCreatetime shop_createtime获取 店铺创建时间
func (obj *ShopMgr) WithShopCreatetime(shopCreatetime int64) Option {
	return optionFunc(func(o *options) { o.query["shop_createtime"] = shopCreatetime })
}

// WithShopEndtime shop_endtime获取 店铺关闭时间
func (obj *ShopMgr) WithShopEndtime(shopEndtime int64) Option {
	return optionFunc(func(o *options) { o.query["shop_endtime"] = shopEndtime })
}

// GetByOption 功能选项模式获取
func (obj *ShopMgr) GetByOption(opts ...Option) (result models.EsShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *ShopMgr) GetByOptions(opts ...Option) (results []*models.EsShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *ShopMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsShop, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Where(options.query)
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

// GetFromShopID 通过shop_id获取内容 店铺Id
func (obj *ShopMgr) GetFromShopID(shopID int) (result models.EsShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_id` = ?", shopID).First(&result).Error

	return
}

// GetBatchFromShopID 批量查找 店铺Id
func (obj *ShopMgr) GetBatchFromShopID(shopIDs []int) (results []*models.EsShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员Id
func (obj *ShopMgr) GetFromMemberID(memberID int) (results []*models.EsShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员Id
func (obj *ShopMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 会员名称
func (obj *ShopMgr) GetFromMemberName(memberName string) (results []*models.EsShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 会员名称
func (obj *ShopMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromShopName 通过shop_name获取内容 店铺名称
func (obj *ShopMgr) GetFromShopName(shopName string) (results []*models.EsShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_name` = ?", shopName).Find(&results).Error

	return
}

// GetBatchFromShopName 批量查找 店铺名称
func (obj *ShopMgr) GetBatchFromShopName(shopNames []string) (results []*models.EsShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_name` IN (?)", shopNames).Find(&results).Error

	return
}

// GetFromShopDisable 通过shop_disable获取内容 店铺状态
func (obj *ShopMgr) GetFromShopDisable(shopDisable string) (results []*models.EsShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_disable` = ?", shopDisable).Find(&results).Error

	return
}

// GetBatchFromShopDisable 批量查找 店铺状态
func (obj *ShopMgr) GetBatchFromShopDisable(shopDisables []string) (results []*models.EsShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_disable` IN (?)", shopDisables).Find(&results).Error

	return
}

// GetFromShopCreatetime 通过shop_createtime获取内容 店铺创建时间
func (obj *ShopMgr) GetFromShopCreatetime(shopCreatetime int64) (results []*models.EsShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_createtime` = ?", shopCreatetime).Find(&results).Error

	return
}

// GetBatchFromShopCreatetime 批量查找 店铺创建时间
func (obj *ShopMgr) GetBatchFromShopCreatetime(shopCreatetimes []int64) (results []*models.EsShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_createtime` IN (?)", shopCreatetimes).Find(&results).Error

	return
}

// GetFromShopEndtime 通过shop_endtime获取内容 店铺关闭时间
func (obj *ShopMgr) GetFromShopEndtime(shopEndtime int64) (results []*models.EsShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_endtime` = ?", shopEndtime).Find(&results).Error

	return
}

// GetBatchFromShopEndtime 批量查找 店铺关闭时间
func (obj *ShopMgr) GetBatchFromShopEndtime(shopEndtimes []int64) (results []*models.EsShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_endtime` IN (?)", shopEndtimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *ShopMgr) FetchByPrimaryKey(shopID int) (result models.EsShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_id` = ?", shopID).First(&result).Error

	return
}
