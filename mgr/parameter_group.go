package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type ParameterGroupMgr struct {
	*_BaseMgr
}

// NewParameterGroupMgr open func
func NewParameterGroupMgr(db db.Repo) *ParameterGroupMgr {
	if db == nil {
		panic(fmt.Errorf("NewParameterGroupMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &ParameterGroupMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_parameter_group"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *ParameterGroupMgr) GetTableName() string {
	return "es_parameter_group"
}

// Reset 重置gorm会话
func (obj *ParameterGroupMgr) Reset() *ParameterGroupMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *ParameterGroupMgr) Get() (result models.EsParameterGroup, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameterGroup{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *ParameterGroupMgr) Gets() (results []*models.EsParameterGroup, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameterGroup{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *ParameterGroupMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsParameterGroup{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithGroupID group_id获取 主键
func (obj *ParameterGroupMgr) WithGroupID(groupID int) Option {
	return optionFunc(func(o *options) { o.query["group_id"] = groupID })
}

// WithGroupName group_name获取 参数组名称
func (obj *ParameterGroupMgr) WithGroupName(groupName string) Option {
	return optionFunc(func(o *options) { o.query["group_name"] = groupName })
}

// WithCategoryID category_id获取 关联分类id
func (obj *ParameterGroupMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithSort sort获取 排序
func (obj *ParameterGroupMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// GetByOption 功能选项模式获取
func (obj *ParameterGroupMgr) GetByOption(opts ...Option) (result models.EsParameterGroup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameterGroup{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *ParameterGroupMgr) GetByOptions(opts ...Option) (results []*models.EsParameterGroup, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameterGroup{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *ParameterGroupMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsParameterGroup, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsParameterGroup{}).Where(options.query)
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

// GetFromGroupID 通过group_id获取内容 主键
func (obj *ParameterGroupMgr) GetFromGroupID(groupID int) (result models.EsParameterGroup, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameterGroup{}).Where("`group_id` = ?", groupID).First(&result).Error

	return
}

// GetBatchFromGroupID 批量查找 主键
func (obj *ParameterGroupMgr) GetBatchFromGroupID(groupIDs []int) (results []*models.EsParameterGroup, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameterGroup{}).Where("`group_id` IN (?)", groupIDs).Find(&results).Error

	return
}

// GetFromGroupName 通过group_name获取内容 参数组名称
func (obj *ParameterGroupMgr) GetFromGroupName(groupName string) (results []*models.EsParameterGroup, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameterGroup{}).Where("`group_name` = ?", groupName).Find(&results).Error

	return
}

// GetBatchFromGroupName 批量查找 参数组名称
func (obj *ParameterGroupMgr) GetBatchFromGroupName(groupNames []string) (results []*models.EsParameterGroup, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameterGroup{}).Where("`group_name` IN (?)", groupNames).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容 关联分类id
func (obj *ParameterGroupMgr) GetFromCategoryID(categoryID int) (results []*models.EsParameterGroup, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameterGroup{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 关联分类id
func (obj *ParameterGroupMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsParameterGroup, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameterGroup{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *ParameterGroupMgr) GetFromSort(sort int) (results []*models.EsParameterGroup, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameterGroup{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *ParameterGroupMgr) GetBatchFromSort(sorts []int) (results []*models.EsParameterGroup, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameterGroup{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *ParameterGroupMgr) FetchByPrimaryKey(groupID int) (result models.EsParameterGroup, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameterGroup{}).Where("`group_id` = ?", groupID).First(&result).Error

	return
}
