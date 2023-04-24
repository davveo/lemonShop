package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type GroupbuyAreaMgr struct {
	*_BaseMgr
}

// NewGroupbuyAreaMgr open func
func NewGroupbuyAreaMgr(db db.Repo) *GroupbuyAreaMgr {
	if db == nil {
		panic(fmt.Errorf("NewGroupbuyAreaMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &GroupbuyAreaMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_category_spec"),
		wdb:       db.GetDbW().Table("es_category_spec"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *GroupbuyAreaMgr) GetTableName() string {
	return "es_groupbuy_area"
}

// Reset 重置gorm会话
func (obj *GroupbuyAreaMgr) Reset() *GroupbuyAreaMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *GroupbuyAreaMgr) Get() (result models.EsGroupbuyArea, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyArea{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *GroupbuyAreaMgr) Gets() (results []*models.EsGroupbuyArea, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyArea{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *GroupbuyAreaMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyArea{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAreaID area_id获取 地区ID
func (obj *GroupbuyAreaMgr) WithAreaID(areaID int) Option {
	return optionFunc(func(o *options) { o.query["area_id"] = areaID })
}

// WithAreaName area_name获取 地区名称
func (obj *GroupbuyAreaMgr) WithAreaName(areaName string) Option {
	return optionFunc(func(o *options) { o.query["area_name"] = areaName })
}

// WithAreaOrder area_order获取 地区排序
func (obj *GroupbuyAreaMgr) WithAreaOrder(areaOrder int) Option {
	return optionFunc(func(o *options) { o.query["area_order"] = areaOrder })
}

// WithAreaPath area_path获取 地区路径结构
func (obj *GroupbuyAreaMgr) WithAreaPath(areaPath string) Option {
	return optionFunc(func(o *options) { o.query["area_path"] = areaPath })
}

// WithParentID parent_id获取 地区父节点
func (obj *GroupbuyAreaMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// GetByOption 功能选项模式获取
func (obj *GroupbuyAreaMgr) GetByOption(opts ...Option) (result models.EsGroupbuyArea, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyArea{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *GroupbuyAreaMgr) GetByOptions(opts ...Option) (results []*models.EsGroupbuyArea, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyArea{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *GroupbuyAreaMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsGroupbuyArea, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyArea{}).Where(options.query)
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

// GetFromAreaID 通过area_id获取内容 地区ID
func (obj *GroupbuyAreaMgr) GetFromAreaID(areaID int) (result models.EsGroupbuyArea, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyArea{}).Where("`area_id` = ?", areaID).First(&result).Error

	return
}

// GetBatchFromAreaID 批量查找 地区ID
func (obj *GroupbuyAreaMgr) GetBatchFromAreaID(areaIDs []int) (results []*models.EsGroupbuyArea, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyArea{}).Where("`area_id` IN (?)", areaIDs).Find(&results).Error

	return
}

// GetFromAreaName 通过area_name获取内容 地区名称
func (obj *GroupbuyAreaMgr) GetFromAreaName(areaName string) (results []*models.EsGroupbuyArea, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyArea{}).Where("`area_name` = ?", areaName).Find(&results).Error

	return
}

// GetBatchFromAreaName 批量查找 地区名称
func (obj *GroupbuyAreaMgr) GetBatchFromAreaName(areaNames []string) (results []*models.EsGroupbuyArea, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyArea{}).Where("`area_name` IN (?)", areaNames).Find(&results).Error

	return
}

// GetFromAreaOrder 通过area_order获取内容 地区排序
func (obj *GroupbuyAreaMgr) GetFromAreaOrder(areaOrder int) (results []*models.EsGroupbuyArea, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyArea{}).Where("`area_order` = ?", areaOrder).Find(&results).Error

	return
}

// GetBatchFromAreaOrder 批量查找 地区排序
func (obj *GroupbuyAreaMgr) GetBatchFromAreaOrder(areaOrders []int) (results []*models.EsGroupbuyArea, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyArea{}).Where("`area_order` IN (?)", areaOrders).Find(&results).Error

	return
}

// GetFromAreaPath 通过area_path获取内容 地区路径结构
func (obj *GroupbuyAreaMgr) GetFromAreaPath(areaPath string) (results []*models.EsGroupbuyArea, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyArea{}).Where("`area_path` = ?", areaPath).Find(&results).Error

	return
}

// GetBatchFromAreaPath 批量查找 地区路径结构
func (obj *GroupbuyAreaMgr) GetBatchFromAreaPath(areaPaths []string) (results []*models.EsGroupbuyArea, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyArea{}).Where("`area_path` IN (?)", areaPaths).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 地区父节点
func (obj *GroupbuyAreaMgr) GetFromParentID(parentID int) (results []*models.EsGroupbuyArea, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyArea{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 地区父节点
func (obj *GroupbuyAreaMgr) GetBatchFromParentID(parentIDs []int) (results []*models.EsGroupbuyArea, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyArea{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *GroupbuyAreaMgr) FetchByPrimaryKey(areaID int) (result models.EsGroupbuyArea, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGroupbuyArea{}).Where("`area_id` = ?", areaID).First(&result).Error

	return
}
