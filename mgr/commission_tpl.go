package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsCommissionTplMgr struct {
	*_BaseMgr
}

// EsCommissionTplMgr open func
func EsCommissionTplMgr(db *gorm.DB) *_EsCommissionTplMgr {
	if db == nil {
		panic(fmt.Errorf("EsCommissionTplMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsCommissionTplMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_commission_tpl"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsCommissionTplMgr) GetTableName() string {
	return "es_commission_tpl"
}

// Reset 重置gorm会话
func (obj *_EsCommissionTplMgr) Reset() *_EsCommissionTplMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsCommissionTplMgr) Get() (result models.EsCommissionTpl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsCommissionTplMgr) Gets() (results []*models.EsCommissionTpl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsCommissionTplMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_EsCommissionTplMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTplName tpl_name获取 名称
func (obj *_EsCommissionTplMgr) WithTplName(tplName string) Option {
	return optionFunc(func(o *options) { o.query["tpl_name"] = tplName })
}

// WithTplDescribe tpl_describe获取 描述
func (obj *_EsCommissionTplMgr) WithTplDescribe(tplDescribe string) Option {
	return optionFunc(func(o *options) { o.query["tpl_describe"] = tplDescribe })
}

// WithCycle cycle获取 周期
func (obj *_EsCommissionTplMgr) WithCycle(cycle int) Option {
	return optionFunc(func(o *options) { o.query["cycle"] = cycle })
}

// WithGrade1 grade1获取 1级返利
func (obj *_EsCommissionTplMgr) WithGrade1(grade1 float64) Option {
	return optionFunc(func(o *options) { o.query["grade1"] = grade1 })
}

// WithGrade2 grade2获取 2级返利
func (obj *_EsCommissionTplMgr) WithGrade2(grade2 float64) Option {
	return optionFunc(func(o *options) { o.query["grade2"] = grade2 })
}

// WithIsDefault is_default获取 是否默认
func (obj *_EsCommissionTplMgr) WithIsDefault(isDefault int16) Option {
	return optionFunc(func(o *options) { o.query["is_default"] = isDefault })
}

// GetByOption 功能选项模式获取
func (obj *_EsCommissionTplMgr) GetByOption(opts ...Option) (result models.EsCommissionTpl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsCommissionTplMgr) GetByOptions(opts ...Option) (results []*models.EsCommissionTpl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsCommissionTplMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsCommissionTpl, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where(options.query)
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

// GetFromID 通过id获取内容 id
func (obj *_EsCommissionTplMgr) GetFromID(id int64) (result models.EsCommissionTpl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_EsCommissionTplMgr) GetBatchFromID(ids []int64) (results []*models.EsCommissionTpl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTplName 通过tpl_name获取内容 名称
func (obj *_EsCommissionTplMgr) GetFromTplName(tplName string) (results []*models.EsCommissionTpl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`tpl_name` = ?", tplName).Find(&results).Error

	return
}

// GetBatchFromTplName 批量查找 名称
func (obj *_EsCommissionTplMgr) GetBatchFromTplName(tplNames []string) (results []*models.EsCommissionTpl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`tpl_name` IN (?)", tplNames).Find(&results).Error

	return
}

// GetFromTplDescribe 通过tpl_describe获取内容 描述
func (obj *_EsCommissionTplMgr) GetFromTplDescribe(tplDescribe string) (results []*models.EsCommissionTpl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`tpl_describe` = ?", tplDescribe).Find(&results).Error

	return
}

// GetBatchFromTplDescribe 批量查找 描述
func (obj *_EsCommissionTplMgr) GetBatchFromTplDescribe(tplDescribes []string) (results []*models.EsCommissionTpl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`tpl_describe` IN (?)", tplDescribes).Find(&results).Error

	return
}

// GetFromCycle 通过cycle获取内容 周期
func (obj *_EsCommissionTplMgr) GetFromCycle(cycle int) (results []*models.EsCommissionTpl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`cycle` = ?", cycle).Find(&results).Error

	return
}

// GetBatchFromCycle 批量查找 周期
func (obj *_EsCommissionTplMgr) GetBatchFromCycle(cycles []int) (results []*models.EsCommissionTpl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`cycle` IN (?)", cycles).Find(&results).Error

	return
}

// GetFromGrade1 通过grade1获取内容 1级返利
func (obj *_EsCommissionTplMgr) GetFromGrade1(grade1 float64) (results []*models.EsCommissionTpl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`grade1` = ?", grade1).Find(&results).Error

	return
}

// GetBatchFromGrade1 批量查找 1级返利
func (obj *_EsCommissionTplMgr) GetBatchFromGrade1(grade1s []float64) (results []*models.EsCommissionTpl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`grade1` IN (?)", grade1s).Find(&results).Error

	return
}

// GetFromGrade2 通过grade2获取内容 2级返利
func (obj *_EsCommissionTplMgr) GetFromGrade2(grade2 float64) (results []*models.EsCommissionTpl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`grade2` = ?", grade2).Find(&results).Error

	return
}

// GetBatchFromGrade2 批量查找 2级返利
func (obj *_EsCommissionTplMgr) GetBatchFromGrade2(grade2s []float64) (results []*models.EsCommissionTpl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`grade2` IN (?)", grade2s).Find(&results).Error

	return
}

// GetFromIsDefault 通过is_default获取内容 是否默认
func (obj *_EsCommissionTplMgr) GetFromIsDefault(isDefault int16) (results []*models.EsCommissionTpl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`is_default` = ?", isDefault).Find(&results).Error

	return
}

// GetBatchFromIsDefault 批量查找 是否默认
func (obj *_EsCommissionTplMgr) GetBatchFromIsDefault(isDefaults []int16) (results []*models.EsCommissionTpl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`is_default` IN (?)", isDefaults).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsCommissionTplMgr) FetchByPrimaryKey(id int64) (result models.EsCommissionTpl, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`id` = ?", id).First(&result).Error

	return
}
