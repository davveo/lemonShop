package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsBrandMgr struct {
	*_BaseMgr
}

// EsBrandMgr open func
func EsBrandMgr(db *gorm.DB) *_EsBrandMgr {
	if db == nil {
		panic(fmt.Errorf("EsBrandMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsBrandMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_brand"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsBrandMgr) GetTableName() string {
	return "es_brand"
}

// Reset 重置gorm会话
func (obj *_EsBrandMgr) Reset() *_EsBrandMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsBrandMgr) Get() (result models.EsBrand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBrand{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsBrandMgr) Gets() (results []*models.EsBrand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBrand{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsBrandMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsBrand{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithBrandID brand_id获取 主键
func (obj *_EsBrandMgr) WithBrandID(brandID int) Option {
	return optionFunc(func(o *options) { o.query["brand_id"] = brandID })
}

// WithName name获取 品牌名称
func (obj *_EsBrandMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithLogo logo获取 品牌图标
func (obj *_EsBrandMgr) WithLogo(logo string) Option {
	return optionFunc(func(o *options) { o.query["logo"] = logo })
}

// WithDisabled disabled获取 是否删除，0删除1未删除
func (obj *_EsBrandMgr) WithDisabled(disabled int) Option {
	return optionFunc(func(o *options) { o.query["disabled"] = disabled })
}

// GetByOption 功能选项模式获取
func (obj *_EsBrandMgr) GetByOption(opts ...Option) (result models.EsBrand, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsBrand{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsBrandMgr) GetByOptions(opts ...Option) (results []*models.EsBrand, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsBrand{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsBrandMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsBrand, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsBrand{}).Where(options.query)
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

// GetFromBrandID 通过brand_id获取内容 主键
func (obj *_EsBrandMgr) GetFromBrandID(brandID int) (result models.EsBrand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBrand{}).Where("`brand_id` = ?", brandID).First(&result).Error

	return
}

// GetBatchFromBrandID 批量查找 主键
func (obj *_EsBrandMgr) GetBatchFromBrandID(brandIDs []int) (results []*models.EsBrand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBrand{}).Where("`brand_id` IN (?)", brandIDs).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 品牌名称
func (obj *_EsBrandMgr) GetFromName(name string) (results []*models.EsBrand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBrand{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 品牌名称
func (obj *_EsBrandMgr) GetBatchFromName(names []string) (results []*models.EsBrand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBrand{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromLogo 通过logo获取内容 品牌图标
func (obj *_EsBrandMgr) GetFromLogo(logo string) (results []*models.EsBrand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBrand{}).Where("`logo` = ?", logo).Find(&results).Error

	return
}

// GetBatchFromLogo 批量查找 品牌图标
func (obj *_EsBrandMgr) GetBatchFromLogo(logos []string) (results []*models.EsBrand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBrand{}).Where("`logo` IN (?)", logos).Find(&results).Error

	return
}

// GetFromDisabled 通过disabled获取内容 是否删除，0删除1未删除
func (obj *_EsBrandMgr) GetFromDisabled(disabled int) (results []*models.EsBrand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBrand{}).Where("`disabled` = ?", disabled).Find(&results).Error

	return
}

// GetBatchFromDisabled 批量查找 是否删除，0删除1未删除
func (obj *_EsBrandMgr) GetBatchFromDisabled(disableds []int) (results []*models.EsBrand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBrand{}).Where("`disabled` IN (?)", disableds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsBrandMgr) FetchByPrimaryKey(brandID int) (result models.EsBrand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBrand{}).Where("`brand_id` = ?", brandID).First(&result).Error

	return
}

// FetchIndexByIndBrand  获取多个内容
func (obj *_EsBrandMgr) FetchIndexByIndBrand(disabled int) (results []*models.EsBrand, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsBrand{}).Where("`disabled` = ?", disabled).Find(&results).Error

	return
}
