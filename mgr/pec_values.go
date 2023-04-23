package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsSpecValuesMgr struct {
	*_BaseMgr
}

// EsSpecValuesMgr open func
func EsSpecValuesMgr(db *gorm.DB) *_EsSpecValuesMgr {
	if db == nil {
		panic(fmt.Errorf("EsSpecValuesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsSpecValuesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_spec_values"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsSpecValuesMgr) GetTableName() string {
	return "es_spec_values"
}

// Reset 重置gorm会话
func (obj *_EsSpecValuesMgr) Reset() *_EsSpecValuesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsSpecValuesMgr) Get() (result models.EsSpecValues, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSpecValues{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsSpecValuesMgr) Gets() (results []*models.EsSpecValues, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSpecValues{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsSpecValuesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsSpecValues{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSpecValueID spec_value_id获取 主键
func (obj *_EsSpecValuesMgr) WithSpecValueID(specValueID int) Option {
	return optionFunc(func(o *options) { o.query["spec_value_id"] = specValueID })
}

// WithSpecID spec_id获取 规格项id
func (obj *_EsSpecValuesMgr) WithSpecID(specID int) Option {
	return optionFunc(func(o *options) { o.query["spec_id"] = specID })
}

// WithSpecValue spec_value获取 规格值名字
func (obj *_EsSpecValuesMgr) WithSpecValue(specValue string) Option {
	return optionFunc(func(o *options) { o.query["spec_value"] = specValue })
}

// WithSellerID seller_id获取 所属卖家
func (obj *_EsSpecValuesMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithSpecName spec_name获取 规格名字
func (obj *_EsSpecValuesMgr) WithSpecName(specName string) Option {
	return optionFunc(func(o *options) { o.query["spec_name"] = specName })
}

// GetByOption 功能选项模式获取
func (obj *_EsSpecValuesMgr) GetByOption(opts ...Option) (result models.EsSpecValues, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSpecValues{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsSpecValuesMgr) GetByOptions(opts ...Option) (results []*models.EsSpecValues, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSpecValues{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsSpecValuesMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSpecValues, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsSpecValues{}).Where(options.query)
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

// GetFromSpecValueID 通过spec_value_id获取内容 主键
func (obj *_EsSpecValuesMgr) GetFromSpecValueID(specValueID int) (result models.EsSpecValues, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSpecValues{}).Where("`spec_value_id` = ?", specValueID).First(&result).Error

	return
}

// GetBatchFromSpecValueID 批量查找 主键
func (obj *_EsSpecValuesMgr) GetBatchFromSpecValueID(specValueIDs []int) (results []*models.EsSpecValues, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSpecValues{}).Where("`spec_value_id` IN (?)", specValueIDs).Find(&results).Error

	return
}

// GetFromSpecID 通过spec_id获取内容 规格项id
func (obj *_EsSpecValuesMgr) GetFromSpecID(specID int) (results []*models.EsSpecValues, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSpecValues{}).Where("`spec_id` = ?", specID).Find(&results).Error

	return
}

// GetBatchFromSpecID 批量查找 规格项id
func (obj *_EsSpecValuesMgr) GetBatchFromSpecID(specIDs []int) (results []*models.EsSpecValues, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSpecValues{}).Where("`spec_id` IN (?)", specIDs).Find(&results).Error

	return
}

// GetFromSpecValue 通过spec_value获取内容 规格值名字
func (obj *_EsSpecValuesMgr) GetFromSpecValue(specValue string) (results []*models.EsSpecValues, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSpecValues{}).Where("`spec_value` = ?", specValue).Find(&results).Error

	return
}

// GetBatchFromSpecValue 批量查找 规格值名字
func (obj *_EsSpecValuesMgr) GetBatchFromSpecValue(specValues []string) (results []*models.EsSpecValues, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSpecValues{}).Where("`spec_value` IN (?)", specValues).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 所属卖家
func (obj *_EsSpecValuesMgr) GetFromSellerID(sellerID int) (results []*models.EsSpecValues, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSpecValues{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 所属卖家
func (obj *_EsSpecValuesMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsSpecValues, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSpecValues{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromSpecName 通过spec_name获取内容 规格名字
func (obj *_EsSpecValuesMgr) GetFromSpecName(specName string) (results []*models.EsSpecValues, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSpecValues{}).Where("`spec_name` = ?", specName).Find(&results).Error

	return
}

// GetBatchFromSpecName 批量查找 规格名字
func (obj *_EsSpecValuesMgr) GetBatchFromSpecName(specNames []string) (results []*models.EsSpecValues, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSpecValues{}).Where("`spec_name` IN (?)", specNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsSpecValuesMgr) FetchByPrimaryKey(specValueID int) (result models.EsSpecValues, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSpecValues{}).Where("`spec_value_id` = ?", specValueID).First(&result).Error

	return
}
