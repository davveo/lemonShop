package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsShopThemesMgr struct {
	*_BaseMgr
}

// EsShopThemesMgr open func
func EsShopThemesMgr(db *gorm.DB) *_EsShopThemesMgr {
	if db == nil {
		panic(fmt.Errorf("EsShopThemesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsShopThemesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_shop_themes"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsShopThemesMgr) GetTableName() string {
	return "es_shop_themes"
}

// Reset 重置gorm会话
func (obj *_EsShopThemesMgr) Reset() *_EsShopThemesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsShopThemesMgr) Get() (result models.EsShopThemes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsShopThemesMgr) Gets() (results []*models.EsShopThemes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsShopThemesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 模版id
func (obj *_EsShopThemesMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 模版名称
func (obj *_EsShopThemesMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithImagePath image_path获取 模版图片路径
func (obj *_EsShopThemesMgr) WithImagePath(imagePath string) Option {
	return optionFunc(func(o *options) { o.query["image_path"] = imagePath })
}

// WithIsDefault is_default获取 是否为默认
func (obj *_EsShopThemesMgr) WithIsDefault(isDefault int) Option {
	return optionFunc(func(o *options) { o.query["is_default"] = isDefault })
}

// WithType type获取 模版类型
func (obj *_EsShopThemesMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithMark mark获取 模版标识
func (obj *_EsShopThemesMgr) WithMark(mark string) Option {
	return optionFunc(func(o *options) { o.query["mark"] = mark })
}

// GetByOption 功能选项模式获取
func (obj *_EsShopThemesMgr) GetByOption(opts ...Option) (result models.EsShopThemes, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsShopThemesMgr) GetByOptions(opts ...Option) (results []*models.EsShopThemes, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsShopThemesMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsShopThemes, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where(options.query)
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

// GetFromID 通过id获取内容 模版id
func (obj *_EsShopThemesMgr) GetFromID(id int) (result models.EsShopThemes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 模版id
func (obj *_EsShopThemesMgr) GetBatchFromID(ids []int) (results []*models.EsShopThemes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 模版名称
func (obj *_EsShopThemesMgr) GetFromName(name string) (results []*models.EsShopThemes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 模版名称
func (obj *_EsShopThemesMgr) GetBatchFromName(names []string) (results []*models.EsShopThemes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromImagePath 通过image_path获取内容 模版图片路径
func (obj *_EsShopThemesMgr) GetFromImagePath(imagePath string) (results []*models.EsShopThemes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`image_path` = ?", imagePath).Find(&results).Error

	return
}

// GetBatchFromImagePath 批量查找 模版图片路径
func (obj *_EsShopThemesMgr) GetBatchFromImagePath(imagePaths []string) (results []*models.EsShopThemes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`image_path` IN (?)", imagePaths).Find(&results).Error

	return
}

// GetFromIsDefault 通过is_default获取内容 是否为默认
func (obj *_EsShopThemesMgr) GetFromIsDefault(isDefault int) (results []*models.EsShopThemes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`is_default` = ?", isDefault).Find(&results).Error

	return
}

// GetBatchFromIsDefault 批量查找 是否为默认
func (obj *_EsShopThemesMgr) GetBatchFromIsDefault(isDefaults []int) (results []*models.EsShopThemes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`is_default` IN (?)", isDefaults).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 模版类型
func (obj *_EsShopThemesMgr) GetFromType(_type string) (results []*models.EsShopThemes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 模版类型
func (obj *_EsShopThemesMgr) GetBatchFromType(_types []string) (results []*models.EsShopThemes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromMark 通过mark获取内容 模版标识
func (obj *_EsShopThemesMgr) GetFromMark(mark string) (results []*models.EsShopThemes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`mark` = ?", mark).Find(&results).Error

	return
}

// GetBatchFromMark 批量查找 模版标识
func (obj *_EsShopThemesMgr) GetBatchFromMark(marks []string) (results []*models.EsShopThemes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`mark` IN (?)", marks).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsShopThemesMgr) FetchByPrimaryKey(id int) (result models.EsShopThemes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`id` = ?", id).First(&result).Error

	return
}
