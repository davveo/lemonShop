package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type ShopMenuMgr struct {
	*_BaseMgr
}

// NewShopMenuMgr open func
func NewShopMenuMgr(db db.Repo) *ShopMenuMgr {
	if db == nil {
		panic(fmt.Errorf("NewShopMenuMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &ShopMenuMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_shop_menu"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *ShopMenuMgr) GetTableName() string {
	return "es_shop_menu"
}

// Reset 重置gorm会话
func (obj *ShopMenuMgr) Reset() *ShopMenuMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *ShopMenuMgr) Get() (result models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *ShopMenuMgr) Gets() (results []*models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *ShopMenuMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 店铺菜单id
func (obj *ShopMenuMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithParentID parent_id获取 菜单父id
func (obj *ShopMenuMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithTitle title获取 菜单标题
func (obj *ShopMenuMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithIDentifier identifier获取 菜单唯一标识
func (obj *ShopMenuMgr) WithIDentifier(identifier string) Option {
	return optionFunc(func(o *options) { o.query["identifier"] = identifier })
}

// WithAuthRegular auth_regular获取 权限表达式
func (obj *ShopMenuMgr) WithAuthRegular(authRegular string) Option {
	return optionFunc(func(o *options) { o.query["auth_regular"] = authRegular })
}

// WithDeleteFlag delete_flag获取 删除标记
func (obj *ShopMenuMgr) WithDeleteFlag(deleteFlag int) Option {
	return optionFunc(func(o *options) { o.query["delete_flag"] = deleteFlag })
}

// WithPath path获取 菜单级别标识
func (obj *ShopMenuMgr) WithPath(path string) Option {
	return optionFunc(func(o *options) { o.query["path"] = path })
}

// WithGrade grade获取 菜单级别
func (obj *ShopMenuMgr) WithGrade(grade int) Option {
	return optionFunc(func(o *options) { o.query["grade"] = grade })
}

// GetByOption 功能选项模式获取
func (obj *ShopMenuMgr) GetByOption(opts ...Option) (result models.EsShopMenu, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *ShopMenuMgr) GetByOptions(opts ...Option) (results []*models.EsShopMenu, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *ShopMenuMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsShopMenu, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where(options.query)
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

// GetFromID 通过id获取内容 店铺菜单id
func (obj *ShopMenuMgr) GetFromID(id int) (result models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 店铺菜单id
func (obj *ShopMenuMgr) GetBatchFromID(ids []int) (results []*models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 菜单父id
func (obj *ShopMenuMgr) GetFromParentID(parentID int) (results []*models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 菜单父id
func (obj *ShopMenuMgr) GetBatchFromParentID(parentIDs []int) (results []*models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 菜单标题
func (obj *ShopMenuMgr) GetFromTitle(title string) (results []*models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 菜单标题
func (obj *ShopMenuMgr) GetBatchFromTitle(titles []string) (results []*models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromIDentifier 通过identifier获取内容 菜单唯一标识
func (obj *ShopMenuMgr) GetFromIDentifier(identifier string) (results []*models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where("`identifier` = ?", identifier).Find(&results).Error

	return
}

// GetBatchFromIDentifier 批量查找 菜单唯一标识
func (obj *ShopMenuMgr) GetBatchFromIDentifier(identifiers []string) (results []*models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where("`identifier` IN (?)", identifiers).Find(&results).Error

	return
}

// GetFromAuthRegular 通过auth_regular获取内容 权限表达式
func (obj *ShopMenuMgr) GetFromAuthRegular(authRegular string) (results []*models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where("`auth_regular` = ?", authRegular).Find(&results).Error

	return
}

// GetBatchFromAuthRegular 批量查找 权限表达式
func (obj *ShopMenuMgr) GetBatchFromAuthRegular(authRegulars []string) (results []*models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where("`auth_regular` IN (?)", authRegulars).Find(&results).Error

	return
}

// GetFromDeleteFlag 通过delete_flag获取内容 删除标记
func (obj *ShopMenuMgr) GetFromDeleteFlag(deleteFlag int) (results []*models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where("`delete_flag` = ?", deleteFlag).Find(&results).Error

	return
}

// GetBatchFromDeleteFlag 批量查找 删除标记
func (obj *ShopMenuMgr) GetBatchFromDeleteFlag(deleteFlags []int) (results []*models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where("`delete_flag` IN (?)", deleteFlags).Find(&results).Error

	return
}

// GetFromPath 通过path获取内容 菜单级别标识
func (obj *ShopMenuMgr) GetFromPath(path string) (results []*models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where("`path` = ?", path).Find(&results).Error

	return
}

// GetBatchFromPath 批量查找 菜单级别标识
func (obj *ShopMenuMgr) GetBatchFromPath(paths []string) (results []*models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where("`path` IN (?)", paths).Find(&results).Error

	return
}

// GetFromGrade 通过grade获取内容 菜单级别
func (obj *ShopMenuMgr) GetFromGrade(grade int) (results []*models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where("`grade` = ?", grade).Find(&results).Error

	return
}

// GetBatchFromGrade 批量查找 菜单级别
func (obj *ShopMenuMgr) GetBatchFromGrade(grades []int) (results []*models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where("`grade` IN (?)", grades).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *ShopMenuMgr) FetchByPrimaryKey(id int) (result models.EsShopMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopMenu{}).Where("`id` = ?", id).First(&result).Error

	return
}
