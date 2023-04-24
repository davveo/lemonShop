package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type MenuMgr struct {
	*_BaseMgr
}

// NewMenuMgr open func
func NewMenuMgr(db db.Repo) *MenuMgr {
	if db == nil {
		panic(fmt.Errorf("NewMenuMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &MenuMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_menu"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *MenuMgr) GetTableName() string {
	return "es_menu"
}

// Reset 重置gorm会话
func (obj *MenuMgr) Reset() *MenuMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *MenuMgr) Get() (result models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *MenuMgr) Gets() (results []*models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *MenuMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 菜单id
func (obj *MenuMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithParentID parent_id获取 父id
func (obj *MenuMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithTitle title获取 菜单标题
func (obj *MenuMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// WithURL url获取 菜单url
func (obj *MenuMgr) WithURL(url string) Option {
	return optionFunc(func(o *options) { o.query["url"] = url })
}

// WithIDentifier identifier获取 菜单唯一标识
func (obj *MenuMgr) WithIDentifier(identifier string) Option {
	return optionFunc(func(o *options) { o.query["identifier"] = identifier })
}

// WithAuthRegular auth_regular获取 权限表达式
func (obj *MenuMgr) WithAuthRegular(authRegular string) Option {
	return optionFunc(func(o *options) { o.query["auth_regular"] = authRegular })
}

// WithDeleteFlag delete_flag获取 删除标记
func (obj *MenuMgr) WithDeleteFlag(deleteFlag int) Option {
	return optionFunc(func(o *options) { o.query["delete_flag"] = deleteFlag })
}

// WithPath path获取 菜单级别标识
func (obj *MenuMgr) WithPath(path string) Option {
	return optionFunc(func(o *options) { o.query["path"] = path })
}

// WithGrade grade获取 菜单级别
func (obj *MenuMgr) WithGrade(grade int) Option {
	return optionFunc(func(o *options) { o.query["grade"] = grade })
}

// GetByOption 功能选项模式获取
func (obj *MenuMgr) GetByOption(opts ...Option) (result models.EsMenu, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *MenuMgr) GetByOptions(opts ...Option) (results []*models.EsMenu, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *MenuMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsMenu, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where(options.query)
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

// GetFromID 通过id获取内容 菜单id
func (obj *MenuMgr) GetFromID(id int) (result models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 菜单id
func (obj *MenuMgr) GetBatchFromID(ids []int) (results []*models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 父id
func (obj *MenuMgr) GetFromParentID(parentID int) (results []*models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 父id
func (obj *MenuMgr) GetBatchFromParentID(parentIDs []int) (results []*models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容 菜单标题
func (obj *MenuMgr) GetFromTitle(title string) (results []*models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找 菜单标题
func (obj *MenuMgr) GetBatchFromTitle(titles []string) (results []*models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

// GetFromURL 通过url获取内容 菜单url
func (obj *MenuMgr) GetFromURL(url string) (results []*models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`url` = ?", url).Find(&results).Error

	return
}

// GetBatchFromURL 批量查找 菜单url
func (obj *MenuMgr) GetBatchFromURL(urls []string) (results []*models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`url` IN (?)", urls).Find(&results).Error

	return
}

// GetFromIDentifier 通过identifier获取内容 菜单唯一标识
func (obj *MenuMgr) GetFromIDentifier(identifier string) (results []*models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`identifier` = ?", identifier).Find(&results).Error

	return
}

// GetBatchFromIDentifier 批量查找 菜单唯一标识
func (obj *MenuMgr) GetBatchFromIDentifier(identifiers []string) (results []*models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`identifier` IN (?)", identifiers).Find(&results).Error

	return
}

// GetFromAuthRegular 通过auth_regular获取内容 权限表达式
func (obj *MenuMgr) GetFromAuthRegular(authRegular string) (results []*models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`auth_regular` = ?", authRegular).Find(&results).Error

	return
}

// GetBatchFromAuthRegular 批量查找 权限表达式
func (obj *MenuMgr) GetBatchFromAuthRegular(authRegulars []string) (results []*models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`auth_regular` IN (?)", authRegulars).Find(&results).Error

	return
}

// GetFromDeleteFlag 通过delete_flag获取内容 删除标记
func (obj *MenuMgr) GetFromDeleteFlag(deleteFlag int) (results []*models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`delete_flag` = ?", deleteFlag).Find(&results).Error

	return
}

// GetBatchFromDeleteFlag 批量查找 删除标记
func (obj *MenuMgr) GetBatchFromDeleteFlag(deleteFlags []int) (results []*models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`delete_flag` IN (?)", deleteFlags).Find(&results).Error

	return
}

// GetFromPath 通过path获取内容 菜单级别标识
func (obj *MenuMgr) GetFromPath(path string) (results []*models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`path` = ?", path).Find(&results).Error

	return
}

// GetBatchFromPath 批量查找 菜单级别标识
func (obj *MenuMgr) GetBatchFromPath(paths []string) (results []*models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`path` IN (?)", paths).Find(&results).Error

	return
}

// GetFromGrade 通过grade获取内容 菜单级别
func (obj *MenuMgr) GetFromGrade(grade int) (results []*models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`grade` = ?", grade).Find(&results).Error

	return
}

// GetBatchFromGrade 批量查找 菜单级别
func (obj *MenuMgr) GetBatchFromGrade(grades []int) (results []*models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`grade` IN (?)", grades).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *MenuMgr) FetchByPrimaryKey(id int) (result models.EsMenu, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMenu{}).Where("`id` = ?", id).First(&result).Error

	return
}
