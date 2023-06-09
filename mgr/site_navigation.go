package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type SiteNavigationMgr struct {
	*_BaseMgr
}

// NewSiteNavigationMgr open func
func NewSiteNavigationMgr(db db.Repo) *SiteNavigationMgr {
	if db == nil {
		panic(fmt.Errorf("NewSiteNavigationMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &SiteNavigationMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_site_navigation"), wdb: db.GetDbW().Table("es_site_navigation"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *SiteNavigationMgr) GetTableName() string {
	return "es_site_navigation"
}

// Reset 重置gorm会话
func (obj *SiteNavigationMgr) Reset() *SiteNavigationMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *SiteNavigationMgr) Get() (result models.EsSiteNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *SiteNavigationMgr) Gets() (results []*models.EsSiteNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *SiteNavigationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithNavigationID navigation_id获取 主键
func (obj *SiteNavigationMgr) WithNavigationID(navigationID int) Option {
	return optionFunc(func(o *options) { o.query["navigation_id"] = navigationID })
}

// WithNavigationName navigation_name获取 导航名称
func (obj *SiteNavigationMgr) WithNavigationName(navigationName string) Option {
	return optionFunc(func(o *options) { o.query["navigation_name"] = navigationName })
}

// WithURL url获取 导航地址
func (obj *SiteNavigationMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithClientType client_type获取 客户端类型
func (obj *SiteNavigationMgr) WithClientType(clientType string) Option {
	return optionFunc(func(o *options) { o.query["client_type"] = clientType })
}

// WithImage image获取 图片
func (obj *SiteNavigationMgr) WithImage(image string) Option {
	return optionFunc(func(o *options) { o.query["image"] = image })
}

// WithSort sort获取 排序
func (obj *SiteNavigationMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// GetByOption 功能选项模式获取
func (obj *SiteNavigationMgr) GetByOption(opts ...Option) (result models.EsSiteNavigation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *SiteNavigationMgr) GetByOptions(opts ...Option) (results []*models.EsSiteNavigation, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *SiteNavigationMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSiteNavigation, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).Where(options.query)
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

// GetFromNavigationID 通过navigation_id获取内容 主键
func (obj *SiteNavigationMgr) GetFromNavigationID(navigationID int) (result models.EsSiteNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).Where("`navigation_id` = ?", navigationID).First(&result).Error

	return
}

// GetBatchFromNavigationID 批量查找 主键
func (obj *SiteNavigationMgr) GetBatchFromNavigationID(navigationIDs []int) (results []*models.EsSiteNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).Where("`navigation_id` IN (?)", navigationIDs).Find(&results).Error

	return
}

// GetFromNavigationName 通过navigation_name获取内容 导航名称
func (obj *SiteNavigationMgr) GetFromNavigationName(navigationName string) (results []*models.EsSiteNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).Where("`navigation_name` = ?", navigationName).Find(&results).Error

	return
}

// GetBatchFromNavigationName 批量查找 导航名称
func (obj *SiteNavigationMgr) GetBatchFromNavigationName(navigationNames []string) (results []*models.EsSiteNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).Where("`navigation_name` IN (?)", navigationNames).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容 导航地址
func (obj *SiteNavigationMgr) GetFromURL(url string) (results []*models.EsSiteNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).Where("`url` = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量查找 导航地址
func (obj *SiteNavigationMgr) GetBatchFromURL(urls []string) (results []*models.EsSiteNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).Where("`url` IN (?)", urls).Find(&results).Error

	return
}

// GetFromClientType 通过client_type获取内容 客户端类型
func (obj *SiteNavigationMgr) GetFromClientType(clientType string) (results []*models.EsSiteNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).Where("`client_type` = ?", clientType).Find(&results).Error

	return
}

// GetBatchFromClientType 批量查找 客户端类型
func (obj *SiteNavigationMgr) GetBatchFromClientType(clientTypes []string) (results []*models.EsSiteNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).Where("`client_type` IN (?)", clientTypes).Find(&results).Error

	return
}

// GetFromImage 通过image获取内容 图片
func (obj *SiteNavigationMgr) GetFromImage(image string) (results []*models.EsSiteNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).Where("`image` = ?", image).Find(&results).Error

	return
}

// GetBatchFromImage 批量查找 图片
func (obj *SiteNavigationMgr) GetBatchFromImage(images []string) (results []*models.EsSiteNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).Where("`image` IN (?)", images).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *SiteNavigationMgr) GetFromSort(sort int) (results []*models.EsSiteNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *SiteNavigationMgr) GetBatchFromSort(sorts []int) (results []*models.EsSiteNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *SiteNavigationMgr) FetchByPrimaryKey(navigationID int) (result models.EsSiteNavigation, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSiteNavigation{}).Where("`navigation_id` = ?", navigationID).First(&result).Error

	return
}
