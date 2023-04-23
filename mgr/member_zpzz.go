package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsMemberZpzzMgr struct {
	*_BaseMgr
}

// EsMemberZpzzMgr open func
func EsMemberZpzzMgr(db *gorm.DB) *_EsMemberZpzzMgr {
	if db == nil {
		panic(fmt.Errorf("EsMemberZpzzMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsMemberZpzzMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_member_zpzz"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsMemberZpzzMgr) GetTableName() string {
	return "es_member_zpzz"
}

// Reset 重置gorm会话
func (obj *_EsMemberZpzzMgr) Reset() *_EsMemberZpzzMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsMemberZpzzMgr) Get() (result models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsMemberZpzzMgr) Gets() (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsMemberZpzzMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *_EsMemberZpzzMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员ID
func (obj *_EsMemberZpzzMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithUname uname获取 会员用户名
func (obj *_EsMemberZpzzMgr) WithUname(uname string) Option {
	return optionFunc(func(o *options) { o.query["uname"] = uname })
}

// WithStatus status获取 状态 NEW_APPLY：新申请，AUDIT_PASS：审核通过，AUDIT_REFUSE：审核未通过
func (obj *_EsMemberZpzzMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCompanyName company_name获取 单位名称
func (obj *_EsMemberZpzzMgr) WithCompanyName(companyName string) Option {
	return optionFunc(func(o *options) { o.query["company_name"] = companyName })
}

// WithTaxpayerCode taxpayer_code获取 纳税人识别码
func (obj *_EsMemberZpzzMgr) WithTaxpayerCode(taxpayerCode string) Option {
	return optionFunc(func(o *options) { o.query["taxpayer_code"] = taxpayerCode })
}

// WithRegisterAddress register_address获取 公司注册地址
func (obj *_EsMemberZpzzMgr) WithRegisterAddress(registerAddress string) Option {
	return optionFunc(func(o *options) { o.query["register_address"] = registerAddress })
}

// WithRegisterTel register_tel获取 公司注册电话
func (obj *_EsMemberZpzzMgr) WithRegisterTel(registerTel string) Option {
	return optionFunc(func(o *options) { o.query["register_tel"] = registerTel })
}

// WithBankName bank_name获取 开户银行
func (obj *_EsMemberZpzzMgr) WithBankName(bankName string) Option {
	return optionFunc(func(o *options) { o.query["bank_name"] = bankName })
}

// WithBankAccount bank_account获取 银行账户
func (obj *_EsMemberZpzzMgr) WithBankAccount(bankAccount string) Option {
	return optionFunc(func(o *options) { o.query["bank_account"] = bankAccount })
}

// WithAuditRemark audit_remark获取 平台审核备注
func (obj *_EsMemberZpzzMgr) WithAuditRemark(auditRemark string) Option {
	return optionFunc(func(o *options) { o.query["audit_remark"] = auditRemark })
}

// WithApplyTime apply_time获取 申请时间
func (obj *_EsMemberZpzzMgr) WithApplyTime(applyTime int64) Option {
	return optionFunc(func(o *options) { o.query["apply_time"] = applyTime })
}

// GetByOption 功能选项模式获取
func (obj *_EsMemberZpzzMgr) GetByOption(opts ...Option) (result models.EsMemberZpzz, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsMemberZpzzMgr) GetByOptions(opts ...Option) (results []*models.EsMemberZpzz, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsMemberZpzzMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsMemberZpzz, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键ID
func (obj *_EsMemberZpzzMgr) GetFromID(id int) (result models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *_EsMemberZpzzMgr) GetBatchFromID(ids []int) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员ID
func (obj *_EsMemberZpzzMgr) GetFromMemberID(memberID int) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员ID
func (obj *_EsMemberZpzzMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromUname 通过uname获取内容 会员用户名
func (obj *_EsMemberZpzzMgr) GetFromUname(uname string) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`uname` = ?", uname).Find(&results).Error

	return
}

// GetBatchFromUname 批量查找 会员用户名
func (obj *_EsMemberZpzzMgr) GetBatchFromUname(unames []string) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`uname` IN (?)", unames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态 NEW_APPLY：新申请，AUDIT_PASS：审核通过，AUDIT_REFUSE：审核未通过
func (obj *_EsMemberZpzzMgr) GetFromStatus(status string) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态 NEW_APPLY：新申请，AUDIT_PASS：审核通过，AUDIT_REFUSE：审核未通过
func (obj *_EsMemberZpzzMgr) GetBatchFromStatus(statuss []string) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCompanyName 通过company_name获取内容 单位名称
func (obj *_EsMemberZpzzMgr) GetFromCompanyName(companyName string) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`company_name` = ?", companyName).Find(&results).Error

	return
}

// GetBatchFromCompanyName 批量查找 单位名称
func (obj *_EsMemberZpzzMgr) GetBatchFromCompanyName(companyNames []string) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`company_name` IN (?)", companyNames).Find(&results).Error

	return
}

// GetFromTaxpayerCode 通过taxpayer_code获取内容 纳税人识别码
func (obj *_EsMemberZpzzMgr) GetFromTaxpayerCode(taxpayerCode string) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`taxpayer_code` = ?", taxpayerCode).Find(&results).Error

	return
}

// GetBatchFromTaxpayerCode 批量查找 纳税人识别码
func (obj *_EsMemberZpzzMgr) GetBatchFromTaxpayerCode(taxpayerCodes []string) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`taxpayer_code` IN (?)", taxpayerCodes).Find(&results).Error

	return
}

// GetFromRegisterAddress 通过register_address获取内容 公司注册地址
func (obj *_EsMemberZpzzMgr) GetFromRegisterAddress(registerAddress string) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`register_address` = ?", registerAddress).Find(&results).Error

	return
}

// GetBatchFromRegisterAddress 批量查找 公司注册地址
func (obj *_EsMemberZpzzMgr) GetBatchFromRegisterAddress(registerAddresss []string) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`register_address` IN (?)", registerAddresss).Find(&results).Error

	return
}

// GetFromRegisterTel 通过register_tel获取内容 公司注册电话
func (obj *_EsMemberZpzzMgr) GetFromRegisterTel(registerTel string) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`register_tel` = ?", registerTel).Find(&results).Error

	return
}

// GetBatchFromRegisterTel 批量查找 公司注册电话
func (obj *_EsMemberZpzzMgr) GetBatchFromRegisterTel(registerTels []string) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`register_tel` IN (?)", registerTels).Find(&results).Error

	return
}

// GetFromBankName 通过bank_name获取内容 开户银行
func (obj *_EsMemberZpzzMgr) GetFromBankName(bankName string) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`bank_name` = ?", bankName).Find(&results).Error

	return
}

// GetBatchFromBankName 批量查找 开户银行
func (obj *_EsMemberZpzzMgr) GetBatchFromBankName(bankNames []string) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`bank_name` IN (?)", bankNames).Find(&results).Error

	return
}

// GetFromBankAccount 通过bank_account获取内容 银行账户
func (obj *_EsMemberZpzzMgr) GetFromBankAccount(bankAccount string) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`bank_account` = ?", bankAccount).Find(&results).Error

	return
}

// GetBatchFromBankAccount 批量查找 银行账户
func (obj *_EsMemberZpzzMgr) GetBatchFromBankAccount(bankAccounts []string) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`bank_account` IN (?)", bankAccounts).Find(&results).Error

	return
}

// GetFromAuditRemark 通过audit_remark获取内容 平台审核备注
func (obj *_EsMemberZpzzMgr) GetFromAuditRemark(auditRemark string) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`audit_remark` = ?", auditRemark).Find(&results).Error

	return
}

// GetBatchFromAuditRemark 批量查找 平台审核备注
func (obj *_EsMemberZpzzMgr) GetBatchFromAuditRemark(auditRemarks []string) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`audit_remark` IN (?)", auditRemarks).Find(&results).Error

	return
}

// GetFromApplyTime 通过apply_time获取内容 申请时间
func (obj *_EsMemberZpzzMgr) GetFromApplyTime(applyTime int64) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`apply_time` = ?", applyTime).Find(&results).Error

	return
}

// GetBatchFromApplyTime 批量查找 申请时间
func (obj *_EsMemberZpzzMgr) GetBatchFromApplyTime(applyTimes []int64) (results []*models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`apply_time` IN (?)", applyTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsMemberZpzzMgr) FetchByPrimaryKey(id int) (result models.EsMemberZpzz, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`id` = ?", id).First(&result).Error

	return
}
