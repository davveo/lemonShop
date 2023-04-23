package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsShopMgr struct {
	*_BaseMgr
}

// EsShopMgr open func
func EsShopMgr(db *gorm.DB) *_EsShopMgr {
	if db == nil {
		panic(fmt.Errorf("EsShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsShopMgr) GetTableName() string {
	return "es_shop"
}

// Reset 重置gorm会话
func (obj *_EsShopMgr) Reset() *_EsShopMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsShopMgr) Get() (result models.EsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsShopMgr) Gets() (results []*models.EsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsShopMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithShopID shop_id获取 店铺Id
func (obj *_EsShopMgr) WithShopID(shopID int) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithMemberID member_id获取 会员Id
func (obj *_EsShopMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMemberName member_name获取 会员名称
func (obj *_EsShopMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithShopName shop_name获取 店铺名称
func (obj *_EsShopMgr) WithShopName(shopName string) Option {
	return optionFunc(func(o *options) { o.query["shop_name"] = shopName })
}

// WithShopDisable shop_disable获取 店铺状态
func (obj *_EsShopMgr) WithShopDisable(shopDisable string) Option {
	return optionFunc(func(o *options) { o.query["shop_disable"] = shopDisable })
}

// WithShopCreatetime shop_createtime获取 店铺创建时间
func (obj *_EsShopMgr) WithShopCreatetime(shopCreatetime int64) Option {
	return optionFunc(func(o *options) { o.query["shop_createtime"] = shopCreatetime })
}

// WithShopEndtime shop_endtime获取 店铺关闭时间
func (obj *_EsShopMgr) WithShopEndtime(shopEndtime int64) Option {
	return optionFunc(func(o *options) { o.query["shop_endtime"] = shopEndtime })
}

// GetByOption 功能选项模式获取
func (obj *_EsShopMgr) GetByOption(opts ...Option) (result models.EsShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsShopMgr) GetByOptions(opts ...Option) (results []*models.EsShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsShopMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsShop, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Where(options.query)
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
func (obj *_EsShopMgr) GetFromShopID(shopID int) (result models.EsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_id` = ?", shopID).First(&result).Error

	return
}

// GetBatchFromShopID 批量查找 店铺Id
func (obj *_EsShopMgr) GetBatchFromShopID(shopIDs []int) (results []*models.EsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员Id
func (obj *_EsShopMgr) GetFromMemberID(memberID int) (results []*models.EsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员Id
func (obj *_EsShopMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 会员名称
func (obj *_EsShopMgr) GetFromMemberName(memberName string) (results []*models.EsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 会员名称
func (obj *_EsShopMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromShopName 通过shop_name获取内容 店铺名称
func (obj *_EsShopMgr) GetFromShopName(shopName string) (results []*models.EsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_name` = ?", shopName).Find(&results).Error

	return
}

// GetBatchFromShopName 批量查找 店铺名称
func (obj *_EsShopMgr) GetBatchFromShopName(shopNames []string) (results []*models.EsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_name` IN (?)", shopNames).Find(&results).Error

	return
}

// GetFromShopDisable 通过shop_disable获取内容 店铺状态
func (obj *_EsShopMgr) GetFromShopDisable(shopDisable string) (results []*models.EsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_disable` = ?", shopDisable).Find(&results).Error

	return
}

// GetBatchFromShopDisable 批量查找 店铺状态
func (obj *_EsShopMgr) GetBatchFromShopDisable(shopDisables []string) (results []*models.EsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_disable` IN (?)", shopDisables).Find(&results).Error

	return
}

// GetFromShopCreatetime 通过shop_createtime获取内容 店铺创建时间
func (obj *_EsShopMgr) GetFromShopCreatetime(shopCreatetime int64) (results []*models.EsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_createtime` = ?", shopCreatetime).Find(&results).Error

	return
}

// GetBatchFromShopCreatetime 批量查找 店铺创建时间
func (obj *_EsShopMgr) GetBatchFromShopCreatetime(shopCreatetimes []int64) (results []*models.EsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_createtime` IN (?)", shopCreatetimes).Find(&results).Error

	return
}

// GetFromShopEndtime 通过shop_endtime获取内容 店铺关闭时间
func (obj *_EsShopMgr) GetFromShopEndtime(shopEndtime int64) (results []*models.EsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_endtime` = ?", shopEndtime).Find(&results).Error

	return
}

// GetBatchFromShopEndtime 批量查找 店铺关闭时间
func (obj *_EsShopMgr) GetBatchFromShopEndtime(shopEndtimes []int64) (results []*models.EsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_endtime` IN (?)", shopEndtimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsShopMgr) FetchByPrimaryKey(shopID int) (result models.EsShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShop{}).Where("`shop_id` = ?", shopID).First(&result).Error

	return
}
