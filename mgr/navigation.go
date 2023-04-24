package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type NavigationMgr struct {
	*_BaseMgr
}

// NewNavigationMgr open func
func NewNavigationMgr(db db.Repo) *NavigationMgr {
	if db == nil {
		panic(fmt.Errorf("NewNavigationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &NavigationMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_navigation"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *NavigationMgr) GetTableName() string {
	return "es_navigation"
}

// Reset 重置gorm会话
func (obj *NavigationMgr) Reset() *NavigationMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *NavigationMgr) Get() (result models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *NavigationMgr) Gets() (results []*models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *NavigationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 导航id
func (obj *NavigationMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 名称
func (obj *NavigationMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithDisable disable获取 是否显示
func (obj *NavigationMgr) WithDisable(disable int) Option {
	return optionFunc(func(o *options) { o.query["disable"] = disable })
}

// WithSort sort获取 排序
func (obj *NavigationMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithNavURL nav_url获取 URL
func (obj *NavigationMgr) WithNavURL(navURL string) Option {
	return optionFunc(func(o *options) { o.query["nav_url"] = navURL })
}

// WithTarget target获取 新窗口打开
func (obj *NavigationMgr) WithTarget(target int) Option {
	return optionFunc(func(o *options) { o.query["target"] = target })
}

// WithShopID shop_id获取 店铺id
func (obj *NavigationMgr) WithShopID(shopID int) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithContents contents获取 内容
func (obj *NavigationMgr) WithContents(contents string) Option {
	return optionFunc(func(o *options) { o.query["contents"] = contents })
}

// GetByOption 功能选项模式获取
func (obj *NavigationMgr) GetByOption(opts ...Option) (result models.EsNavigation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *NavigationMgr) GetByOptions(opts ...Option) (results []*models.EsNavigation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *NavigationMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsNavigation, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where(options.query)
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

// GetFromID 通过id获取内容 导航id
func (obj *NavigationMgr) GetFromID(id int) (result models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 导航id
func (obj *NavigationMgr) GetBatchFromID(ids []int) (results []*models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 名称
func (obj *NavigationMgr) GetFromName(name string) (results []*models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 名称
func (obj *NavigationMgr) GetBatchFromName(names []string) (results []*models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromDisable 通过disable获取内容 是否显示
func (obj *NavigationMgr) GetFromDisable(disable int) (results []*models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where("`disable` = ?", disable).Find(&results).Error

	return
}

// GetBatchFromDisable 批量查找 是否显示
func (obj *NavigationMgr) GetBatchFromDisable(disables []int) (results []*models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where("`disable` IN (?)", disables).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *NavigationMgr) GetFromSort(sort int) (results []*models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *NavigationMgr) GetBatchFromSort(sorts []int) (results []*models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromNavURL 通过nav_url获取内容 URL
func (obj *NavigationMgr) GetFromNavURL(navURL string) (results []*models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where("`nav_url` = ?", navURL).Find(&results).Error

	return
}

// GetBatchFromNavURL 批量查找 URL
func (obj *NavigationMgr) GetBatchFromNavURL(navURLs []string) (results []*models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where("`nav_url` IN (?)", navURLs).Find(&results).Error

	return
}

// GetFromTarget 通过target获取内容 新窗口打开
func (obj *NavigationMgr) GetFromTarget(target int) (results []*models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where("`target` = ?", target).Find(&results).Error

	return
}

// GetBatchFromTarget 批量查找 新窗口打开
func (obj *NavigationMgr) GetBatchFromTarget(targets []int) (results []*models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where("`target` IN (?)", targets).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *NavigationMgr) GetFromShopID(shopID int) (results []*models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *NavigationMgr) GetBatchFromShopID(shopIDs []int) (results []*models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromContents 通过contents获取内容 内容
func (obj *NavigationMgr) GetFromContents(contents string) (results []*models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where("`contents` = ?", contents).Find(&results).Error

	return
}

// GetBatchFromContents 批量查找 内容
func (obj *NavigationMgr) GetBatchFromContents(contentss []string) (results []*models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where("`contents` IN (?)", contentss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *NavigationMgr) FetchByPrimaryKey(id int) (result models.EsNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsNavigation{}).Where("`id` = ?", id).First(&result).Error

	return
}
