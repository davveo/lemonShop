package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsRegionsMgr struct {
	*_BaseMgr
}

// EsRegionsMgr open func
func EsRegionsMgr(db *gorm.DB) *_EsRegionsMgr {
	if db == nil {
		panic(fmt.Errorf("EsRegionsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsRegionsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_regions"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsRegionsMgr) GetTableName() string {
	return "es_regions"
}

// Reset 重置gorm会话
func (obj *_EsRegionsMgr) Reset() *_EsRegionsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsRegionsMgr) Get() (result models.EsRegions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsRegionsMgr) Gets() (results []*models.EsRegions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsRegionsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 地区id
func (obj *_EsRegionsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithParentID parent_id获取 父地区id
func (obj *_EsRegionsMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithRegionPath region_path获取 路径
func (obj *_EsRegionsMgr) WithRegionPath(regionPath string) Option {
	return optionFunc(func(o *options) { o.query["region_path"] = regionPath })
}

// WithRegionGrade region_grade获取 级别
func (obj *_EsRegionsMgr) WithRegionGrade(regionGrade int) Option {
	return optionFunc(func(o *options) { o.query["region_grade"] = regionGrade })
}

// WithLocalName local_name获取 名称
func (obj *_EsRegionsMgr) WithLocalName(localName string) Option {
	return optionFunc(func(o *options) { o.query["local_name"] = localName })
}

// WithZipcode zipcode获取 邮编
func (obj *_EsRegionsMgr) WithZipcode(zipcode string) Option {
	return optionFunc(func(o *options) { o.query["zipcode"] = zipcode })
}

// WithCod cod获取 是否支持货到付款
func (obj *_EsRegionsMgr) WithCod(cod string) Option {
	return optionFunc(func(o *options) { o.query["cod"] = cod })
}

// GetByOption 功能选项模式获取
func (obj *_EsRegionsMgr) GetByOption(opts ...Option) (result models.EsRegions, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsRegionsMgr) GetByOptions(opts ...Option) (results []*models.EsRegions, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsRegionsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsRegions, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Where(options.query)
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
func (obj *_EsRegionsMgr) GetFromID(id int) (result models.EsRegions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 地区id
func (obj *_EsRegionsMgr) GetBatchFromID(ids []int) (results []*models.EsRegions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 父地区id
func (obj *_EsRegionsMgr) GetFromParentID(parentID int) (results []*models.EsRegions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 父地区id
func (obj *_EsRegionsMgr) GetBatchFromParentID(parentIDs []int) (results []*models.EsRegions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromRegionPath 通过region_path获取内容 路径
func (obj *_EsRegionsMgr) GetFromRegionPath(regionPath string) (results []*models.EsRegions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`region_path` = ?", regionPath).Find(&results).Error

	return
}

// GetBatchFromRegionPath 批量查找 路径
func (obj *_EsRegionsMgr) GetBatchFromRegionPath(regionPaths []string) (results []*models.EsRegions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`region_path` IN (?)", regionPaths).Find(&results).Error

	return
}

// GetFromRegionGrade 通过region_grade获取内容 级别
func (obj *_EsRegionsMgr) GetFromRegionGrade(regionGrade int) (results []*models.EsRegions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`region_grade` = ?", regionGrade).Find(&results).Error

	return
}

// GetBatchFromRegionGrade 批量查找 级别
func (obj *_EsRegionsMgr) GetBatchFromRegionGrade(regionGrades []int) (results []*models.EsRegions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`region_grade` IN (?)", regionGrades).Find(&results).Error

	return
}

// GetFromLocalName 通过local_name获取内容 名称
func (obj *_EsRegionsMgr) GetFromLocalName(localName string) (results []*models.EsRegions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`local_name` = ?", localName).Find(&results).Error

	return
}

// GetBatchFromLocalName 批量查找 名称
func (obj *_EsRegionsMgr) GetBatchFromLocalName(localNames []string) (results []*models.EsRegions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`local_name` IN (?)", localNames).Find(&results).Error

	return
}

// GetFromZipcode 通过zipcode获取内容 邮编
func (obj *_EsRegionsMgr) GetFromZipcode(zipcode string) (results []*models.EsRegions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`zipcode` = ?", zipcode).Find(&results).Error

	return
}

// GetBatchFromZipcode 批量查找 邮编
func (obj *_EsRegionsMgr) GetBatchFromZipcode(zipcodes []string) (results []*models.EsRegions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`zipcode` IN (?)", zipcodes).Find(&results).Error

	return
}

// GetFromCod 通过cod获取内容 是否支持货到付款
func (obj *_EsRegionsMgr) GetFromCod(cod string) (results []*models.EsRegions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`cod` = ?", cod).Find(&results).Error

	return
}

// GetBatchFromCod 批量查找 是否支持货到付款
func (obj *_EsRegionsMgr) GetBatchFromCod(cods []string) (results []*models.EsRegions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`cod` IN (?)", cods).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsRegionsMgr) FetchByPrimaryKey(id int) (result models.EsRegions, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsRegions{}).Where("`id` = ?", id).First(&result).Error

	return
}
