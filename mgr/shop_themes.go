package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type ShopThemesMgr struct {
	*_BaseMgr
}

// NewShopThemesMgr open func
func NewShopThemesMgr(db db.Repo) *ShopThemesMgr {
	if db == nil {
		panic(fmt.Errorf("NewShopThemesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &ShopThemesMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_shop_themes"), wdb: db.GetDbW().Table("es_shop_themes"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *ShopThemesMgr) GetTableName() string {
	return "es_shop_themes"
}

// Reset 重置gorm会话
func (obj *ShopThemesMgr) Reset() *ShopThemesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *ShopThemesMgr) Get() (result models.EsShopThemes, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *ShopThemesMgr) Gets() (results []*models.EsShopThemes, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *ShopThemesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 模版id
func (obj *ShopThemesMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 模版名称
func (obj *ShopThemesMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithImagePath image_path获取 模版图片路径
func (obj *ShopThemesMgr) WithImagePath(imagePath string) Option {
	return optionFunc(func(o *options) { o.query["image_path"] = imagePath })
}

// WithIsDefault is_default获取 是否为默认
func (obj *ShopThemesMgr) WithIsDefault(isDefault int) Option {
	return optionFunc(func(o *options) { o.query["is_default"] = isDefault })
}

// WithType type获取 模版类型
func (obj *ShopThemesMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithMark mark获取 模版标识
func (obj *ShopThemesMgr) WithMark(mark string) Option {
	return optionFunc(func(o *options) { o.query["mark"] = mark })
}

// GetByOption 功能选项模式获取
func (obj *ShopThemesMgr) GetByOption(opts ...Option) (result models.EsShopThemes, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *ShopThemesMgr) GetByOptions(opts ...Option) (results []*models.EsShopThemes, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *ShopThemesMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsShopThemes, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where(options.query)
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
func (obj *ShopThemesMgr) GetFromID(id int) (result models.EsShopThemes, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 模版id
func (obj *ShopThemesMgr) GetBatchFromID(ids []int) (results []*models.EsShopThemes, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 模版名称
func (obj *ShopThemesMgr) GetFromName(name string) (results []*models.EsShopThemes, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 模版名称
func (obj *ShopThemesMgr) GetBatchFromName(names []string) (results []*models.EsShopThemes, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromImagePath 通过image_path获取内容 模版图片路径
func (obj *ShopThemesMgr) GetFromImagePath(imagePath string) (results []*models.EsShopThemes, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`image_path` = ?", imagePath).Find(&results).Error

	return
}

// GetBatchFromImagePath 批量查找 模版图片路径
func (obj *ShopThemesMgr) GetBatchFromImagePath(imagePaths []string) (results []*models.EsShopThemes, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`image_path` IN (?)", imagePaths).Find(&results).Error

	return
}

// GetFromIsDefault 通过is_default获取内容 是否为默认
func (obj *ShopThemesMgr) GetFromIsDefault(isDefault int) (results []*models.EsShopThemes, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`is_default` = ?", isDefault).Find(&results).Error

	return
}

// GetBatchFromIsDefault 批量查找 是否为默认
func (obj *ShopThemesMgr) GetBatchFromIsDefault(isDefaults []int) (results []*models.EsShopThemes, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`is_default` IN (?)", isDefaults).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 模版类型
func (obj *ShopThemesMgr) GetFromType(_type string) (results []*models.EsShopThemes, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 模版类型
func (obj *ShopThemesMgr) GetBatchFromType(_types []string) (results []*models.EsShopThemes, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromMark 通过mark获取内容 模版标识
func (obj *ShopThemesMgr) GetFromMark(mark string) (results []*models.EsShopThemes, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`mark` = ?", mark).Find(&results).Error

	return
}

// GetBatchFromMark 批量查找 模版标识
func (obj *ShopThemesMgr) GetBatchFromMark(marks []string) (results []*models.EsShopThemes, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`mark` IN (?)", marks).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *ShopThemesMgr) FetchByPrimaryKey(id int) (result models.EsShopThemes, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopThemes{}).Where("`id` = ?", id).First(&result).Error

	return
}
