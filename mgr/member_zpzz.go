package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type MemberZpzzMgr struct {
	*_BaseMgr
}

// NewMemberZpzzMgr open func
func NewMemberZpzzMgr(db db.Repo) *MemberZpzzMgr {
	if db == nil {
		panic(fmt.Errorf("NewMemberZpzzMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &MemberZpzzMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_member_zpzz"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *MemberZpzzMgr) GetTableName() string {
	return "es_member_zpzz"
}

// Reset 重置gorm会话
func (obj *MemberZpzzMgr) Reset() *MemberZpzzMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *MemberZpzzMgr) Get() (result models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *MemberZpzzMgr) Gets() (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *MemberZpzzMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *MemberZpzzMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员ID
func (obj *MemberZpzzMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithUname uname获取 会员用户名
func (obj *MemberZpzzMgr) WithUname(uname string) Option {
	return optionFunc(func(o *options) { o.query["uname"] = uname })
}

// WithStatus status获取 状态 NEW_APPLY：新申请，AUDIT_PASS：审核通过，AUDIT_REFUSE：审核未通过
func (obj *MemberZpzzMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithCompanyName company_name获取 单位名称
func (obj *MemberZpzzMgr) WithCompanyName(companyName string) Option {
	return optionFunc(func(o *options) { o.query["company_name"] = companyName })
}

// WithTaxpayerCode taxpayer_code获取 纳税人识别码
func (obj *MemberZpzzMgr) WithTaxpayerCode(taxpayerCode string) Option {
	return optionFunc(func(o *options) { o.query["taxpayer_code"] = taxpayerCode })
}

// WithRegisterAddress register_address获取 公司注册地址
func (obj *MemberZpzzMgr) WithRegisterAddress(registerAddress string) Option {
	return optionFunc(func(o *options) { o.query["register_address"] = registerAddress })
}

// WithRegisterTel register_tel获取 公司注册电话
func (obj *MemberZpzzMgr) WithRegisterTel(registerTel string) Option {
	return optionFunc(func(o *options) { o.query["register_tel"] = registerTel })
}

// WithBankName bank_name获取 开户银行
func (obj *MemberZpzzMgr) WithBankName(bankName string) Option {
	return optionFunc(func(o *options) { o.query["bank_name"] = bankName })
}

// WithBankAccount bank_account获取 银行账户
func (obj *MemberZpzzMgr) WithBankAccount(bankAccount string) Option {
	return optionFunc(func(o *options) { o.query["bank_account"] = bankAccount })
}

// WithAuditRemark audit_remark获取 平台审核备注
func (obj *MemberZpzzMgr) WithAuditRemark(auditRemark string) Option {
	return optionFunc(func(o *options) { o.query["audit_remark"] = auditRemark })
}

// WithApplyTime apply_time获取 申请时间
func (obj *MemberZpzzMgr) WithApplyTime(applyTime int64) Option {
	return optionFunc(func(o *options) { o.query["apply_time"] = applyTime })
}

// GetByOption 功能选项模式获取
func (obj *MemberZpzzMgr) GetByOption(opts ...Option) (result models.EsMemberZpzz, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *MemberZpzzMgr) GetByOptions(opts ...Option) (results []*models.EsMemberZpzz, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *MemberZpzzMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsMemberZpzz, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where(options.query)
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
func (obj *MemberZpzzMgr) GetFromID(id int) (result models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *MemberZpzzMgr) GetBatchFromID(ids []int) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员ID
func (obj *MemberZpzzMgr) GetFromMemberID(memberID int) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员ID
func (obj *MemberZpzzMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromUname 通过uname获取内容 会员用户名
func (obj *MemberZpzzMgr) GetFromUname(uname string) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`uname` = ?", uname).Find(&results).Error

	return
}

// GetBatchFromUname 批量查找 会员用户名
func (obj *MemberZpzzMgr) GetBatchFromUname(unames []string) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`uname` IN (?)", unames).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态 NEW_APPLY：新申请，AUDIT_PASS：审核通过，AUDIT_REFUSE：审核未通过
func (obj *MemberZpzzMgr) GetFromStatus(status string) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态 NEW_APPLY：新申请，AUDIT_PASS：审核通过，AUDIT_REFUSE：审核未通过
func (obj *MemberZpzzMgr) GetBatchFromStatus(statuss []string) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromCompanyName 通过company_name获取内容 单位名称
func (obj *MemberZpzzMgr) GetFromCompanyName(companyName string) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`company_name` = ?", companyName).Find(&results).Error

	return
}

// GetBatchFromCompanyName 批量查找 单位名称
func (obj *MemberZpzzMgr) GetBatchFromCompanyName(companyNames []string) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`company_name` IN (?)", companyNames).Find(&results).Error

	return
}

// GetFromTaxpayerCode 通过taxpayer_code获取内容 纳税人识别码
func (obj *MemberZpzzMgr) GetFromTaxpayerCode(taxpayerCode string) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`taxpayer_code` = ?", taxpayerCode).Find(&results).Error

	return
}

// GetBatchFromTaxpayerCode 批量查找 纳税人识别码
func (obj *MemberZpzzMgr) GetBatchFromTaxpayerCode(taxpayerCodes []string) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`taxpayer_code` IN (?)", taxpayerCodes).Find(&results).Error

	return
}

// GetFromRegisterAddress 通过register_address获取内容 公司注册地址
func (obj *MemberZpzzMgr) GetFromRegisterAddress(registerAddress string) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`register_address` = ?", registerAddress).Find(&results).Error

	return
}

// GetBatchFromRegisterAddress 批量查找 公司注册地址
func (obj *MemberZpzzMgr) GetBatchFromRegisterAddress(registerAddresss []string) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`register_address` IN (?)", registerAddresss).Find(&results).Error

	return
}

// GetFromRegisterTel 通过register_tel获取内容 公司注册电话
func (obj *MemberZpzzMgr) GetFromRegisterTel(registerTel string) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`register_tel` = ?", registerTel).Find(&results).Error

	return
}

// GetBatchFromRegisterTel 批量查找 公司注册电话
func (obj *MemberZpzzMgr) GetBatchFromRegisterTel(registerTels []string) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`register_tel` IN (?)", registerTels).Find(&results).Error

	return
}

// GetFromBankName 通过bank_name获取内容 开户银行
func (obj *MemberZpzzMgr) GetFromBankName(bankName string) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`bank_name` = ?", bankName).Find(&results).Error

	return
}

// GetBatchFromBankName 批量查找 开户银行
func (obj *MemberZpzzMgr) GetBatchFromBankName(bankNames []string) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`bank_name` IN (?)", bankNames).Find(&results).Error

	return
}

// GetFromBankAccount 通过bank_account获取内容 银行账户
func (obj *MemberZpzzMgr) GetFromBankAccount(bankAccount string) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`bank_account` = ?", bankAccount).Find(&results).Error

	return
}

// GetBatchFromBankAccount 批量查找 银行账户
func (obj *MemberZpzzMgr) GetBatchFromBankAccount(bankAccounts []string) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`bank_account` IN (?)", bankAccounts).Find(&results).Error

	return
}

// GetFromAuditRemark 通过audit_remark获取内容 平台审核备注
func (obj *MemberZpzzMgr) GetFromAuditRemark(auditRemark string) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`audit_remark` = ?", auditRemark).Find(&results).Error

	return
}

// GetBatchFromAuditRemark 批量查找 平台审核备注
func (obj *MemberZpzzMgr) GetBatchFromAuditRemark(auditRemarks []string) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`audit_remark` IN (?)", auditRemarks).Find(&results).Error

	return
}

// GetFromApplyTime 通过apply_time获取内容 申请时间
func (obj *MemberZpzzMgr) GetFromApplyTime(applyTime int64) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`apply_time` = ?", applyTime).Find(&results).Error

	return
}

// GetBatchFromApplyTime 批量查找 申请时间
func (obj *MemberZpzzMgr) GetBatchFromApplyTime(applyTimes []int64) (results []*models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`apply_time` IN (?)", applyTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *MemberZpzzMgr) FetchByPrimaryKey(id int) (result models.EsMemberZpzz, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberZpzz{}).Where("`id` = ?", id).First(&result).Error

	return
}
