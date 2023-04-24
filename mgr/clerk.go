package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type ClerkMgr struct {
	*_BaseMgr
}

// NewClerkMgr open func
func NewClerkMgr(db db.Repo) *ClerkMgr {
	if db == nil {
		panic(fmt.Errorf("NewClerkMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &ClerkMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_category_spec"),
		wdb:       db.GetDbW().Table("es_category_spec"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *ClerkMgr) GetTableName() string {
	return "es_clerk"
}

// Reset 重置gorm会话
func (obj *ClerkMgr) Reset() *ClerkMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *ClerkMgr) Get() (result models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *ClerkMgr) Gets() (results []*models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *ClerkMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithClerkID clerk_id获取 店员id
func (obj *ClerkMgr) WithClerkID(clerkID int) Option {
	return optionFunc(func(o *options) { o.query["clerk_id"] = clerkID })
}

// WithMemberID member_id获取 会员id
func (obj *ClerkMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithClerkName clerk_name获取 店员名称
func (obj *ClerkMgr) WithClerkName(clerkName string) Option {
	return optionFunc(func(o *options) { o.query["clerk_name"] = clerkName })
}

// WithFounder founder获取 是否为超级管理员，1为超级管理员 0为其他管理员
func (obj *ClerkMgr) WithFounder(founder int) Option {
	return optionFunc(func(o *options) { o.query["founder"] = founder })
}

// WithRoleID role_id获取 权限id
func (obj *ClerkMgr) WithRoleID(roleID int) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithUserState user_state获取 店员状态，0为禁用，1为正常
func (obj *ClerkMgr) WithUserState(userState int) Option {
	return optionFunc(func(o *options) { o.query["user_state"] = userState })
}

// WithCreateTime create_time获取 创建日期
func (obj *ClerkMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithShopID shop_id获取 店铺id
func (obj *ClerkMgr) WithShopID(shopID int) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// GetByOption 功能选项模式获取
func (obj *ClerkMgr) GetByOption(opts ...Option) (result models.EsClerk, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *ClerkMgr) GetByOptions(opts ...Option) (results []*models.EsClerk, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *ClerkMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsClerk, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where(options.query)
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

// GetFromClerkID 通过clerk_id获取内容 店员id
func (obj *ClerkMgr) GetFromClerkID(clerkID int) (result models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where("`clerk_id` = ?", clerkID).First(&result).Error

	return
}

// GetBatchFromClerkID 批量查找 店员id
func (obj *ClerkMgr) GetBatchFromClerkID(clerkIDs []int) (results []*models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where("`clerk_id` IN (?)", clerkIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *ClerkMgr) GetFromMemberID(memberID int) (results []*models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *ClerkMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromClerkName 通过clerk_name获取内容 店员名称
func (obj *ClerkMgr) GetFromClerkName(clerkName string) (results []*models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where("`clerk_name` = ?", clerkName).Find(&results).Error

	return
}

// GetBatchFromClerkName 批量查找 店员名称
func (obj *ClerkMgr) GetBatchFromClerkName(clerkNames []string) (results []*models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where("`clerk_name` IN (?)", clerkNames).Find(&results).Error

	return
}

// GetFromFounder 通过founder获取内容 是否为超级管理员，1为超级管理员 0为其他管理员
func (obj *ClerkMgr) GetFromFounder(founder int) (results []*models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where("`founder` = ?", founder).Find(&results).Error

	return
}

// GetBatchFromFounder 批量查找 是否为超级管理员，1为超级管理员 0为其他管理员
func (obj *ClerkMgr) GetBatchFromFounder(founders []int) (results []*models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where("`founder` IN (?)", founders).Find(&results).Error

	return
}

// GetFromRoleID 通过role_id获取内容 权限id
func (obj *ClerkMgr) GetFromRoleID(roleID int) (results []*models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where("`role_id` = ?", roleID).Find(&results).Error

	return
}

// GetBatchFromRoleID 批量查找 权限id
func (obj *ClerkMgr) GetBatchFromRoleID(roleIDs []int) (results []*models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where("`role_id` IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromUserState 通过user_state获取内容 店员状态，0为禁用，1为正常
func (obj *ClerkMgr) GetFromUserState(userState int) (results []*models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where("`user_state` = ?", userState).Find(&results).Error

	return
}

// GetBatchFromUserState 批量查找 店员状态，0为禁用，1为正常
func (obj *ClerkMgr) GetBatchFromUserState(userStates []int) (results []*models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where("`user_state` IN (?)", userStates).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建日期
func (obj *ClerkMgr) GetFromCreateTime(createTime int64) (results []*models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建日期
func (obj *ClerkMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *ClerkMgr) GetFromShopID(shopID int) (results []*models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *ClerkMgr) GetBatchFromShopID(shopIDs []int) (results []*models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *ClerkMgr) FetchByPrimaryKey(clerkID int) (result models.EsClerk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsClerk{}).Where("`clerk_id` = ?", clerkID).First(&result).Error

	return
}
