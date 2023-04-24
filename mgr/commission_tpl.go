package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type CommissionTplMgr struct {
	*_BaseMgr
}

// NewCommissionTplMgr open func
func NewCommissionTplMgr(db db.Repo) *CommissionTplMgr {
	if db == nil {
		panic(fmt.Errorf("NewCommissionTplMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &CommissionTplMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_category_spec"),
		wdb:       db.GetDbW().Table("es_category_spec"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *CommissionTplMgr) GetTableName() string {
	return "es_commission_tpl"
}

// Reset 重置gorm会话
func (obj *CommissionTplMgr) Reset() *CommissionTplMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *CommissionTplMgr) Get() (result models.EsCommissionTpl, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *CommissionTplMgr) Gets() (results []*models.EsCommissionTpl, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *CommissionTplMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *CommissionTplMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithTplName tpl_name获取 名称
func (obj *CommissionTplMgr) WithTplName(tplName string) Option {
	return optionFunc(func(o *options) { o.query["tpl_name"] = tplName })
}

// WithTplDescribe tpl_describe获取 描述
func (obj *CommissionTplMgr) WithTplDescribe(tplDescribe string) Option {
	return optionFunc(func(o *options) { o.query["tpl_describe"] = tplDescribe })
}

// WithCycle cycle获取 周期
func (obj *CommissionTplMgr) WithCycle(cycle int) Option {
	return optionFunc(func(o *options) { o.query["cycle"] = cycle })
}

// WithGrade1 grade1获取 1级返利
func (obj *CommissionTplMgr) WithGrade1(grade1 float64) Option {
	return optionFunc(func(o *options) { o.query["grade1"] = grade1 })
}

// WithGrade2 grade2获取 2级返利
func (obj *CommissionTplMgr) WithGrade2(grade2 float64) Option {
	return optionFunc(func(o *options) { o.query["grade2"] = grade2 })
}

// WithIsDefault is_default获取 是否默认
func (obj *CommissionTplMgr) WithIsDefault(isDefault int16) Option {
	return optionFunc(func(o *options) { o.query["is_default"] = isDefault })
}

// GetByOption 功能选项模式获取
func (obj *CommissionTplMgr) GetByOption(opts ...Option) (result models.EsCommissionTpl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *CommissionTplMgr) GetByOptions(opts ...Option) (results []*models.EsCommissionTpl, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *CommissionTplMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsCommissionTpl, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where(options.query)
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
func (obj *CommissionTplMgr) GetFromID(id int64) (result models.EsCommissionTpl, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *CommissionTplMgr) GetBatchFromID(ids []int64) (results []*models.EsCommissionTpl, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromTplName 通过tpl_name获取内容 名称
func (obj *CommissionTplMgr) GetFromTplName(tplName string) (results []*models.EsCommissionTpl, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`tpl_name` = ?", tplName).Find(&results).Error

	return
}

// GetBatchFromTplName 批量查找 名称
func (obj *CommissionTplMgr) GetBatchFromTplName(tplNames []string) (results []*models.EsCommissionTpl, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`tpl_name` IN (?)", tplNames).Find(&results).Error

	return
}

// GetFromTplDescribe 通过tpl_describe获取内容 描述
func (obj *CommissionTplMgr) GetFromTplDescribe(tplDescribe string) (results []*models.EsCommissionTpl, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`tpl_describe` = ?", tplDescribe).Find(&results).Error

	return
}

// GetBatchFromTplDescribe 批量查找 描述
func (obj *CommissionTplMgr) GetBatchFromTplDescribe(tplDescribes []string) (results []*models.EsCommissionTpl, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`tpl_describe` IN (?)", tplDescribes).Find(&results).Error

	return
}

// GetFromCycle 通过cycle获取内容 周期
func (obj *CommissionTplMgr) GetFromCycle(cycle int) (results []*models.EsCommissionTpl, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`cycle` = ?", cycle).Find(&results).Error

	return
}

// GetBatchFromCycle 批量查找 周期
func (obj *CommissionTplMgr) GetBatchFromCycle(cycles []int) (results []*models.EsCommissionTpl, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`cycle` IN (?)", cycles).Find(&results).Error

	return
}

// GetFromGrade1 通过grade1获取内容 1级返利
func (obj *CommissionTplMgr) GetFromGrade1(grade1 float64) (results []*models.EsCommissionTpl, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`grade1` = ?", grade1).Find(&results).Error

	return
}

// GetBatchFromGrade1 批量查找 1级返利
func (obj *CommissionTplMgr) GetBatchFromGrade1(grade1s []float64) (results []*models.EsCommissionTpl, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`grade1` IN (?)", grade1s).Find(&results).Error

	return
}

// GetFromGrade2 通过grade2获取内容 2级返利
func (obj *CommissionTplMgr) GetFromGrade2(grade2 float64) (results []*models.EsCommissionTpl, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`grade2` = ?", grade2).Find(&results).Error

	return
}

// GetBatchFromGrade2 批量查找 2级返利
func (obj *CommissionTplMgr) GetBatchFromGrade2(grade2s []float64) (results []*models.EsCommissionTpl, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`grade2` IN (?)", grade2s).Find(&results).Error

	return
}

// GetFromIsDefault 通过is_default获取内容 是否默认
func (obj *CommissionTplMgr) GetFromIsDefault(isDefault int16) (results []*models.EsCommissionTpl, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`is_default` = ?", isDefault).Find(&results).Error

	return
}

// GetBatchFromIsDefault 批量查找 是否默认
func (obj *CommissionTplMgr) GetBatchFromIsDefault(isDefaults []int16) (results []*models.EsCommissionTpl, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`is_default` IN (?)", isDefaults).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *CommissionTplMgr) FetchByPrimaryKey(id int64) (result models.EsCommissionTpl, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommissionTpl{}).Where("`id` = ?", id).First(&result).Error

	return
}
