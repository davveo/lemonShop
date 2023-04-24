package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type RegionsMgr struct {
	*_BaseMgr
}

// NewRegionsMgr open func
func NewRegionsMgr(db db.Repo) *RegionsMgr {
	if db == nil {
		panic(fmt.Errorf("NewRegionsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &RegionsMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_regions"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *RegionsMgr) GetTableName() string {
	return "es_regions"
}

// Reset 重置gorm会话
func (obj *RegionsMgr) Reset() *RegionsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *RegionsMgr) Get() (result models.EsRegions, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *RegionsMgr) Gets() (results []*models.EsRegions, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *RegionsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 地区id
func (obj *RegionsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithParentID parent_id获取 父地区id
func (obj *RegionsMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithRegionPath region_path获取 路径
func (obj *RegionsMgr) WithRegionPath(regionPath string) Option {
	return optionFunc(func(o *options) { o.query["region_path"] = regionPath })
}

// WithRegionGrade region_grade获取 级别
func (obj *RegionsMgr) WithRegionGrade(regionGrade int) Option {
	return optionFunc(func(o *options) { o.query["region_grade"] = regionGrade })
}

// WithLocalName local_name获取 名称
func (obj *RegionsMgr) WithLocalName(localName string) Option {
	return optionFunc(func(o *options) { o.query["local_name"] = localName })
}

// WithZipcode zipcode获取 邮编
func (obj *RegionsMgr) WithZipcode(zipcode string) Option {
	return optionFunc(func(o *options) { o.query["zipcode"] = zipcode })
}

// WithCod cod获取 是否支持货到付款
func (obj *RegionsMgr) WithCod(cod string) Option {
	return optionFunc(func(o *options) { o.query["cod"] = cod })
}

// GetByOption 功能选项模式获取
func (obj *RegionsMgr) GetByOption(opts ...Option) (result models.EsRegions, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *RegionsMgr) GetByOptions(opts ...Option) (results []*models.EsRegions, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *RegionsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsRegions, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Where(options.query)
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

// GetFromID 通过id获取内容 地区id
func (obj *RegionsMgr) GetFromID(id int) (result models.EsRegions, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 地区id
func (obj *RegionsMgr) GetBatchFromID(ids []int) (results []*models.EsRegions, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 父地区id
func (obj *RegionsMgr) GetFromParentID(parentID int) (results []*models.EsRegions, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 父地区id
func (obj *RegionsMgr) GetBatchFromParentID(parentIDs []int) (results []*models.EsRegions, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromRegionPath 通过region_path获取内容 路径
func (obj *RegionsMgr) GetFromRegionPath(regionPath string) (results []*models.EsRegions, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`region_path` = ?", regionPath).Find(&results).Error

	return
}

// GetBatchFromRegionPath 批量查找 路径
func (obj *RegionsMgr) GetBatchFromRegionPath(regionPaths []string) (results []*models.EsRegions, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`region_path` IN (?)", regionPaths).Find(&results).Error

	return
}

// GetFromRegionGrade 通过region_grade获取内容 级别
func (obj *RegionsMgr) GetFromRegionGrade(regionGrade int) (results []*models.EsRegions, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`region_grade` = ?", regionGrade).Find(&results).Error

	return
}

// GetBatchFromRegionGrade 批量查找 级别
func (obj *RegionsMgr) GetBatchFromRegionGrade(regionGrades []int) (results []*models.EsRegions, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`region_grade` IN (?)", regionGrades).Find(&results).Error

	return
}

// GetFromLocalName 通过local_name获取内容 名称
func (obj *RegionsMgr) GetFromLocalName(localName string) (results []*models.EsRegions, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`local_name` = ?", localName).Find(&results).Error

	return
}

// GetBatchFromLocalName 批量查找 名称
func (obj *RegionsMgr) GetBatchFromLocalName(localNames []string) (results []*models.EsRegions, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`local_name` IN (?)", localNames).Find(&results).Error

	return
}

// GetFromZipcode 通过zipcode获取内容 邮编
func (obj *RegionsMgr) GetFromZipcode(zipcode string) (results []*models.EsRegions, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`zipcode` = ?", zipcode).Find(&results).Error

	return
}

// GetBatchFromZipcode 批量查找 邮编
func (obj *RegionsMgr) GetBatchFromZipcode(zipcodes []string) (results []*models.EsRegions, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`zipcode` IN (?)", zipcodes).Find(&results).Error

	return
}

// GetFromCod 通过cod获取内容 是否支持货到付款
func (obj *RegionsMgr) GetFromCod(cod string) (results []*models.EsRegions, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`cod` = ?", cod).Find(&results).Error

	return
}

// GetBatchFromCod 批量查找 是否支持货到付款
func (obj *RegionsMgr) GetBatchFromCod(cods []string) (results []*models.EsRegions, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`cod` IN (?)", cods).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *RegionsMgr) FetchByPrimaryKey(id int) (result models.EsRegions, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`id` = ?", id).First(&result).Error

	return
}
