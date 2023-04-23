package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsLogiCompanyMgr struct {
	*_BaseMgr
}

// EsLogiCompanyMgr open func
func EsLogiCompanyMgr(db *gorm.DB) *_EsLogiCompanyMgr {
	if db == nil {
		panic(fmt.Errorf("EsLogiCompanyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsLogiCompanyMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_logi_company"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsLogiCompanyMgr) GetTableName() string {
	return "es_logi_company"
}

// Reset 重置gorm会话
func (obj *_EsLogiCompanyMgr) Reset() *_EsLogiCompanyMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsLogiCompanyMgr) Get() (result models.EsLogiCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsLogiCompanyMgr) Gets() (results []*models.EsLogiCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsLogiCompanyMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 物流公司id
func (obj *_EsLogiCompanyMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 物流公司名称
func (obj *_EsLogiCompanyMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCode code获取 物流公司code
func (obj *_EsLogiCompanyMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithKdcode kdcode获取 快递鸟物流公司code
func (obj *_EsLogiCompanyMgr) WithKdcode(kdcode string) Option {
	return optionFunc(func(o *options) { o.query["kdcode"] = kdcode })
}

// WithIsWaybill is_waybill获取 是否支持电子面单1：支持 0：不支持
func (obj *_EsLogiCompanyMgr) WithIsWaybill(isWaybill int) Option {
	return optionFunc(func(o *options) { o.query["is_waybill"] = isWaybill })
}

// WithCustomerName customer_name获取 物流公司客户号
func (obj *_EsLogiCompanyMgr) WithCustomerName(customerName string) Option {
	return optionFunc(func(o *options) { o.query["customer_name"] = customerName })
}

// WithCustomerPwd customer_pwd获取 物流公司电子面单密码
func (obj *_EsLogiCompanyMgr) WithCustomerPwd(customerPwd string) Option {
	return optionFunc(func(o *options) { o.query["customer_pwd"] = customerPwd })
}

// GetByOption 功能选项模式获取
func (obj *_EsLogiCompanyMgr) GetByOption(opts ...Option) (result models.EsLogiCompany, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsLogiCompanyMgr) GetByOptions(opts ...Option) (results []*models.EsLogiCompany, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsLogiCompanyMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsLogiCompany, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where(options.query)
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

// GetFromID 通过id获取内容 物流公司id
func (obj *_EsLogiCompanyMgr) GetFromID(id int) (result models.EsLogiCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 物流公司id
func (obj *_EsLogiCompanyMgr) GetBatchFromID(ids []int) (results []*models.EsLogiCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 物流公司名称
func (obj *_EsLogiCompanyMgr) GetFromName(name string) (results []*models.EsLogiCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 物流公司名称
func (obj *_EsLogiCompanyMgr) GetBatchFromName(names []string) (results []*models.EsLogiCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 物流公司code
func (obj *_EsLogiCompanyMgr) GetFromCode(code string) (results []*models.EsLogiCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 物流公司code
func (obj *_EsLogiCompanyMgr) GetBatchFromCode(codes []string) (results []*models.EsLogiCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromKdcode 通过kdcode获取内容 快递鸟物流公司code
func (obj *_EsLogiCompanyMgr) GetFromKdcode(kdcode string) (results []*models.EsLogiCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`kdcode` = ?", kdcode).Find(&results).Error

	return
}

// GetBatchFromKdcode 批量查找 快递鸟物流公司code
func (obj *_EsLogiCompanyMgr) GetBatchFromKdcode(kdcodes []string) (results []*models.EsLogiCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`kdcode` IN (?)", kdcodes).Find(&results).Error

	return
}

// GetFromIsWaybill 通过is_waybill获取内容 是否支持电子面单1：支持 0：不支持
func (obj *_EsLogiCompanyMgr) GetFromIsWaybill(isWaybill int) (results []*models.EsLogiCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`is_waybill` = ?", isWaybill).Find(&results).Error

	return
}

// GetBatchFromIsWaybill 批量查找 是否支持电子面单1：支持 0：不支持
func (obj *_EsLogiCompanyMgr) GetBatchFromIsWaybill(isWaybills []int) (results []*models.EsLogiCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`is_waybill` IN (?)", isWaybills).Find(&results).Error

	return
}

// GetFromCustomerName 通过customer_name获取内容 物流公司客户号
func (obj *_EsLogiCompanyMgr) GetFromCustomerName(customerName string) (results []*models.EsLogiCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`customer_name` = ?", customerName).Find(&results).Error

	return
}

// GetBatchFromCustomerName 批量查找 物流公司客户号
func (obj *_EsLogiCompanyMgr) GetBatchFromCustomerName(customerNames []string) (results []*models.EsLogiCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`customer_name` IN (?)", customerNames).Find(&results).Error

	return
}

// GetFromCustomerPwd 通过customer_pwd获取内容 物流公司电子面单密码
func (obj *_EsLogiCompanyMgr) GetFromCustomerPwd(customerPwd string) (results []*models.EsLogiCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`customer_pwd` = ?", customerPwd).Find(&results).Error

	return
}

// GetBatchFromCustomerPwd 批量查找 物流公司电子面单密码
func (obj *_EsLogiCompanyMgr) GetBatchFromCustomerPwd(customerPwds []string) (results []*models.EsLogiCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`customer_pwd` IN (?)", customerPwds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsLogiCompanyMgr) FetchByPrimaryKey(id int) (result models.EsLogiCompany, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`id` = ?", id).First(&result).Error

	return
}
