package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type PageMgr struct {
	*_BaseMgr
}

// NewPageMgr open func
func NewPageMgr(db db.Repo) *PageMgr {
	if db == nil {
		panic(fmt.Errorf("NewPageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &PageMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_page"),
		wdb:       db.GetDbW().Table("es_page"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *PageMgr) GetTableName() string {
	return "es_page"
}

// Reset 重置gorm会话
func (obj *PageMgr) Reset() *PageMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *PageMgr) Get() (result models.EsPage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPage{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *PageMgr) Gets() (results []*models.EsPage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPage{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *PageMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsPage{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithPageID page_id获取 主键id
func (obj *PageMgr) WithPageID(pageID int) Option {
	return optionFunc(func(o *options) { o.query["page_id"] = pageID })
}

// WithPageName page_name获取 楼层名称
func (obj *PageMgr) WithPageName(pageName string) Option {
	return optionFunc(func(o *options) { o.query["page_name"] = pageName })
}

// WithPageData page_data获取 楼层数据
func (obj *PageMgr) WithPageData(pageData string) Option {
	return optionFunc(func(o *options) { o.query["page_data"] = pageData })
}

// WithPageType page_type获取 页面类型
func (obj *PageMgr) WithPageType(pageType string) Option {
	return optionFunc(func(o *options) { o.query["page_type"] = pageType })
}

// WithClientType client_type获取 客户端类型
func (obj *PageMgr) WithClientType(clientType string) Option {
	return optionFunc(func(o *options) { o.query["client_type"] = clientType })
}

// GetByOption 功能选项模式获取
func (obj *PageMgr) GetByOption(opts ...Option) (result models.EsPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPage{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *PageMgr) GetByOptions(opts ...Option) (results []*models.EsPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPage{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *PageMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsPage, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsPage{}).Where(options.query)
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

// GetFromPageID 通过page_id获取内容 主键id
func (obj *PageMgr) GetFromPageID(pageID int) (result models.EsPage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPage{}).Where("`page_id` = ?", pageID).First(&result).Error

	return
}

// GetBatchFromPageID 批量查找 主键id
func (obj *PageMgr) GetBatchFromPageID(pageIDs []int) (results []*models.EsPage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPage{}).Where("`page_id` IN (?)", pageIDs).Find(&results).Error

	return
}

// GetFromPageName 通过page_name获取内容 楼层名称
func (obj *PageMgr) GetFromPageName(pageName string) (results []*models.EsPage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPage{}).Where("`page_name` = ?", pageName).Find(&results).Error

	return
}

// GetBatchFromPageName 批量查找 楼层名称
func (obj *PageMgr) GetBatchFromPageName(pageNames []string) (results []*models.EsPage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPage{}).Where("`page_name` IN (?)", pageNames).Find(&results).Error

	return
}

// GetFromPageData 通过page_data获取内容 楼层数据
func (obj *PageMgr) GetFromPageData(pageData string) (results []*models.EsPage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPage{}).Where("`page_data` = ?", pageData).Find(&results).Error

	return
}

// GetBatchFromPageData 批量查找 楼层数据
func (obj *PageMgr) GetBatchFromPageData(pageDatas []string) (results []*models.EsPage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPage{}).Where("`page_data` IN (?)", pageDatas).Find(&results).Error

	return
}

// GetFromPageType 通过page_type获取内容 页面类型
func (obj *PageMgr) GetFromPageType(pageType string) (results []*models.EsPage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPage{}).Where("`page_type` = ?", pageType).Find(&results).Error

	return
}

// GetBatchFromPageType 批量查找 页面类型
func (obj *PageMgr) GetBatchFromPageType(pageTypes []string) (results []*models.EsPage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPage{}).Where("`page_type` IN (?)", pageTypes).Find(&results).Error

	return
}

// GetFromClientType 通过client_type获取内容 客户端类型
func (obj *PageMgr) GetFromClientType(clientType string) (results []*models.EsPage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPage{}).Where("`client_type` = ?", clientType).Find(&results).Error

	return
}

// GetBatchFromClientType 批量查找 客户端类型
func (obj *PageMgr) GetBatchFromClientType(clientTypes []string) (results []*models.EsPage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPage{}).Where("`client_type` IN (?)", clientTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *PageMgr) FetchByPrimaryKey(pageID int) (result models.EsPage, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsPage{}).Where("`page_id` = ?", pageID).First(&result).Error

	return
}
