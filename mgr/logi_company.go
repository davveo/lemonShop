package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type LogiCompanyMgr struct {
	*_BaseMgr
}

// NewLogiCompanyMgr open func
func NewLogiCompanyMgr(db db.Repo) *LogiCompanyMgr {
	if db == nil {
		panic(fmt.Errorf("NewLogiCompanyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &LogiCompanyMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_logi_company"), wdb: db.GetDbW().Table("es_logi_company"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *LogiCompanyMgr) GetTableName() string {
	return "es_logi_company"
}

// Reset 重置gorm会话
func (obj *LogiCompanyMgr) Reset() *LogiCompanyMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *LogiCompanyMgr) Get() (result models.EsLogiCompany, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *LogiCompanyMgr) Gets() (results []*models.EsLogiCompany, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *LogiCompanyMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 物流公司id
func (obj *LogiCompanyMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithName name获取 物流公司名称
func (obj *LogiCompanyMgr) WithName(name string) Option {
	return optionFunc(func(o *options) { o.query["name"] = name })
}

// WithCode code获取 物流公司code
func (obj *LogiCompanyMgr) WithCode(code string) Option {
	return optionFunc(func(o *options) { o.query["code"] = code })
}

// WithKdcode kdcode获取 快递鸟物流公司code
func (obj *LogiCompanyMgr) WithKdcode(kdcode string) Option {
	return optionFunc(func(o *options) { o.query["kdcode"] = kdcode })
}

// WithIsWaybill is_waybill获取 是否支持电子面单1：支持 0：不支持
func (obj *LogiCompanyMgr) WithIsWaybill(isWaybill int) Option {
	return optionFunc(func(o *options) { o.query["is_waybill"] = isWaybill })
}

// WithCustomerName customer_name获取 物流公司客户号
func (obj *LogiCompanyMgr) WithCustomerName(customerName string) Option {
	return optionFunc(func(o *options) { o.query["customer_name"] = customerName })
}

// WithCustomerPwd customer_pwd获取 物流公司电子面单密码
func (obj *LogiCompanyMgr) WithCustomerPwd(customerPwd string) Option {
	return optionFunc(func(o *options) { o.query["customer_pwd"] = customerPwd })
}

// GetByOption 功能选项模式获取
func (obj *LogiCompanyMgr) GetByOption(opts ...Option) (result models.EsLogiCompany, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *LogiCompanyMgr) GetByOptions(opts ...Option) (results []*models.EsLogiCompany, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *LogiCompanyMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsLogiCompany, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where(options.query)
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
func (obj *LogiCompanyMgr) GetFromID(id int) (result models.EsLogiCompany, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 物流公司id
func (obj *LogiCompanyMgr) GetBatchFromID(ids []int) (results []*models.EsLogiCompany, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromName 通过name获取内容 物流公司名称
func (obj *LogiCompanyMgr) GetFromName(name string) (results []*models.EsLogiCompany, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`name` = ?", name).Find(&results).Error

	return
}

// GetBatchFromName 批量查找 物流公司名称
func (obj *LogiCompanyMgr) GetBatchFromName(names []string) (results []*models.EsLogiCompany, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`name` IN (?)", names).Find(&results).Error

	return
}

// GetFromCode 通过code获取内容 物流公司code
func (obj *LogiCompanyMgr) GetFromCode(code string) (results []*models.EsLogiCompany, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`code` = ?", code).Find(&results).Error

	return
}

// GetBatchFromCode 批量查找 物流公司code
func (obj *LogiCompanyMgr) GetBatchFromCode(codes []string) (results []*models.EsLogiCompany, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`code` IN (?)", codes).Find(&results).Error

	return
}

// GetFromKdcode 通过kdcode获取内容 快递鸟物流公司code
func (obj *LogiCompanyMgr) GetFromKdcode(kdcode string) (results []*models.EsLogiCompany, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`kdcode` = ?", kdcode).Find(&results).Error

	return
}

// GetBatchFromKdcode 批量查找 快递鸟物流公司code
func (obj *LogiCompanyMgr) GetBatchFromKdcode(kdcodes []string) (results []*models.EsLogiCompany, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`kdcode` IN (?)", kdcodes).Find(&results).Error

	return
}

// GetFromIsWaybill 通过is_waybill获取内容 是否支持电子面单1：支持 0：不支持
func (obj *LogiCompanyMgr) GetFromIsWaybill(isWaybill int) (results []*models.EsLogiCompany, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`is_waybill` = ?", isWaybill).Find(&results).Error

	return
}

// GetBatchFromIsWaybill 批量查找 是否支持电子面单1：支持 0：不支持
func (obj *LogiCompanyMgr) GetBatchFromIsWaybill(isWaybills []int) (results []*models.EsLogiCompany, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`is_waybill` IN (?)", isWaybills).Find(&results).Error

	return
}

// GetFromCustomerName 通过customer_name获取内容 物流公司客户号
func (obj *LogiCompanyMgr) GetFromCustomerName(customerName string) (results []*models.EsLogiCompany, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`customer_name` = ?", customerName).Find(&results).Error

	return
}

// GetBatchFromCustomerName 批量查找 物流公司客户号
func (obj *LogiCompanyMgr) GetBatchFromCustomerName(customerNames []string) (results []*models.EsLogiCompany, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`customer_name` IN (?)", customerNames).Find(&results).Error

	return
}

// GetFromCustomerPwd 通过customer_pwd获取内容 物流公司电子面单密码
func (obj *LogiCompanyMgr) GetFromCustomerPwd(customerPwd string) (results []*models.EsLogiCompany, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`customer_pwd` = ?", customerPwd).Find(&results).Error

	return
}

// GetBatchFromCustomerPwd 批量查找 物流公司电子面单密码
func (obj *LogiCompanyMgr) GetBatchFromCustomerPwd(customerPwds []string) (results []*models.EsLogiCompany, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`customer_pwd` IN (?)", customerPwds).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *LogiCompanyMgr) FetchByPrimaryKey(id int) (result models.EsLogiCompany, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsLogiCompany{}).Where("`id` = ?", id).First(&result).Error

	return
}
