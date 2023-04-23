package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsLogisticsCompanyMgr struct {
	*_BaseMgr
}

// EsLogisticsCompanyMgr open func
func EsLogisticsCompanyMgr(db *gorm.DB) *_EsLogisticsCompanyMgr {
	if db == nil {
		panic(fmt.Errorf("EsLogisticsCompanyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsLogisticsCompanyMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_logistics_company"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsLogisticsCompanyMgr) GetTableName() string {
	return "es_logistics_company"
}

// Reset 重置gorm会话
func (obj *_EsLogisticsCompanyMgr) Reset() *_EsLogisticsCompanyMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsLogisticsCompanyMgr) Get() (result models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsLogisticsCompanyMgr) Gets() (results []*models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsLogisticsCompanyMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_EsLogisticsCompanyMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取
func (obj *_EsLogisticsCompanyMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCode code获取
func (obj *_EsLogisticsCompanyMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithIsWaybill is_waybill获取
func (obj *_EsLogisticsCompanyMgr) WithIsWaybill(isWaybill int) Option {
	return optionFunc(func(o *options) { o.query["is_waybill"] = isWaybill })
}

// WithKdcode kdcode获取
func (obj *_EsLogisticsCompanyMgr) WithKdcode(kdcode string) Option {
	return optionFunc(func(o *options) { o.query["kdcode"] = kdcode })
}

// WithFormItems form_items获取
func (obj *_EsLogisticsCompanyMgr) WithFormItems(formItems string) Option {
	return optionFunc(func(o *options) { o.query["form_items"] = formItems })
}

// WithDeleteStatus delete_status获取 是否删除 DELETED：已删除，NORMAL：正常
func (obj *_EsLogisticsCompanyMgr) WithDeleteStatus(deleteStatus string) Option {
	return optionFunc(func(o *options) { o.query["delete_status"] = deleteStatus })
}

// WithDisabled disabled获取 是否禁用 OPEN：开启，CLOSE：禁用
func (obj *_EsLogisticsCompanyMgr) WithDisabled(disabled string) Option {
	return optionFunc(func(o *options) { o.query["disabled"] = disabled })
}

// GetByOption 功能选项模式获取
func (obj *_EsLogisticsCompanyMgr) GetByOption(opts ...Option) (result models.EsLogisticsCompany, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsLogisticsCompanyMgr) GetByOptions(opts ...Option) (results []*models.EsLogisticsCompany, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsLogisticsCompanyMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsLogisticsCompany, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *_EsLogisticsCompanyMgr) GetFromID(id int) (result models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_EsLogisticsCompanyMgr) GetBatchFromID(ids []int) (results []*models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容
func (obj *_EsLogisticsCompanyMgr) GetFromName(name string) (results []*models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找
func (obj *_EsLogisticsCompanyMgr) GetBatchFromName(names []string) (results []*models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容
func (obj *_EsLogisticsCompanyMgr) GetFromCode(code string) (results []*models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找
func (obj *_EsLogisticsCompanyMgr) GetBatchFromCode(codes []string) (results []*models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromIsWaybill 通过is_waybill获取内容
func (obj *_EsLogisticsCompanyMgr) GetFromIsWaybill(isWaybill int) (results []*models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where("`is_waybill` = ?", isWaybill).Find(&results).Error

	return
}

// GetBatchFromIsWaybill 批量查找
func (obj *_EsLogisticsCompanyMgr) GetBatchFromIsWaybill(isWaybills []int) (results []*models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where("`is_waybill` IN (?)", isWaybills).Find(&results).Error

	return
}

// GetFromKdcode 通过kdcode获取内容
func (obj *_EsLogisticsCompanyMgr) GetFromKdcode(kdcode string) (results []*models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where("`kdcode` = ?", kdcode).Find(&results).Error

	return
}

// GetBatchFromKdcode 批量查找
func (obj *_EsLogisticsCompanyMgr) GetBatchFromKdcode(kdcodes []string) (results []*models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where("`kdcode` IN (?)", kdcodes).Find(&results).Error

	return
}

// GetFromFormItems 通过form_items获取内容
func (obj *_EsLogisticsCompanyMgr) GetFromFormItems(formItems string) (results []*models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where("`form_items` = ?", formItems).Find(&results).Error

	return
}

// GetBatchFromFormItems 批量查找
func (obj *_EsLogisticsCompanyMgr) GetBatchFromFormItems(formItemss []string) (results []*models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where("`form_items` IN (?)", formItemss).Find(&results).Error

	return
}

// GetFromDeleteStatus 通过delete_status获取内容 是否删除 DELETED：已删除，NORMAL：正常
func (obj *_EsLogisticsCompanyMgr) GetFromDeleteStatus(deleteStatus string) (results []*models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where("`delete_status` = ?", deleteStatus).Find(&results).Error

	return
}

// GetBatchFromDeleteStatus 批量查找 是否删除 DELETED：已删除，NORMAL：正常
func (obj *_EsLogisticsCompanyMgr) GetBatchFromDeleteStatus(deleteStatuss []string) (results []*models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where("`delete_status` IN (?)", deleteStatuss).Find(&results).Error

	return
}

// GetFromDisabled 通过disabled获取内容 是否禁用 OPEN：开启，CLOSE：禁用
func (obj *_EsLogisticsCompanyMgr) GetFromDisabled(disabled string) (results []*models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where("`disabled` = ?", disabled).Find(&results).Error

	return
}

// GetBatchFromDisabled 批量查找 是否禁用 OPEN：开启，CLOSE：禁用
func (obj *_EsLogisticsCompanyMgr) GetBatchFromDisabled(disableds []string) (results []*models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where("`disabled` IN (?)", disableds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsLogisticsCompanyMgr) FetchByPrimaryKey(id int) (result models.EsLogisticsCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogisticsCompany{}).Where("`id` = ?", id).First(&result).Error

	return
}
