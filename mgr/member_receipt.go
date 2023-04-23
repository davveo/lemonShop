package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsMemberReceiptMgr struct {
	*_BaseMgr
}

// EsMemberReceiptMgr open func
func EsMemberReceiptMgr(db *gorm.DB) *_EsMemberReceiptMgr {
	if db == nil {
		panic(fmt.Errorf("EsMemberReceiptMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsMemberReceiptMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_member_receipt"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsMemberReceiptMgr) GetTableName() string {
	return "es_member_receipt"
}

// Reset 重置gorm会话
func (obj *_EsMemberReceiptMgr) Reset() *_EsMemberReceiptMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsMemberReceiptMgr) Get() (result models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsMemberReceiptMgr) Gets() (results []*models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsMemberReceiptMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithReceiptID receipt_id获取 主键ID
func (obj *_EsMemberReceiptMgr) WithReceiptID(receiptID int) Option {
	return optionFunc(func(o *options) { o.query["receipt_id"] = receiptID })
}

// WithMemberID member_id获取 会员ID
func (obj *_EsMemberReceiptMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithReceiptType receipt_type获取 发票类型 ELECTRO：电子普通发票，VATORDINARY：增值税普通发票
func (obj *_EsMemberReceiptMgr) WithReceiptType(receiptType string) Option {
	return optionFunc(func(o *options) { o.query["receipt_type"] = receiptType })
}

// WithReceiptTitle receipt_title获取 发票抬头
func (obj *_EsMemberReceiptMgr) WithReceiptTitle(receiptTitle string) Option {
	return optionFunc(func(o *options) { o.query["receipt_title"] = receiptTitle })
}

// WithReceiptContent receipt_content获取 发票内容
func (obj *_EsMemberReceiptMgr) WithReceiptContent(receiptContent string) Option {
	return optionFunc(func(o *options) { o.query["receipt_content"] = receiptContent })
}

// WithTaxNo tax_no获取 纳税人识别号
func (obj *_EsMemberReceiptMgr) WithTaxNo(taxNo string) Option {
	return optionFunc(func(o *options) { o.query["tax_no"] = taxNo })
}

// WithMemberMobile member_mobile获取 收票人手机号
func (obj *_EsMemberReceiptMgr) WithMemberMobile(memberMobile string) Option {
	return optionFunc(func(o *options) { o.query["member_mobile"] = memberMobile })
}

// WithMemberEmail member_email获取 收票人邮箱
func (obj *_EsMemberReceiptMgr) WithMemberEmail(memberEmail string) Option {
	return optionFunc(func(o *options) { o.query["member_email"] = memberEmail })
}

// WithIsDefault is_default获取 是否为默认选项 0：否，1：是
func (obj *_EsMemberReceiptMgr) WithIsDefault(isDefault int16) Option {
	return optionFunc(func(o *options) { o.query["is_default"] = isDefault })
}

// GetByOption 功能选项模式获取
func (obj *_EsMemberReceiptMgr) GetByOption(opts ...Option) (result models.EsMemberReceipt, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsMemberReceiptMgr) GetByOptions(opts ...Option) (results []*models.EsMemberReceipt, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsMemberReceiptMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsMemberReceipt, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where(options.query)
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

// GetFromReceiptID 通过receipt_id获取内容 主键ID
func (obj *_EsMemberReceiptMgr) GetFromReceiptID(receiptID int) (result models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`receipt_id` = ?", receiptID).First(&result).Error

	return
}

// GetBatchFromReceiptID 批量查找 主键ID
func (obj *_EsMemberReceiptMgr) GetBatchFromReceiptID(receiptIDs []int) (results []*models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`receipt_id` IN (?)", receiptIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员ID
func (obj *_EsMemberReceiptMgr) GetFromMemberID(memberID int) (results []*models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员ID
func (obj *_EsMemberReceiptMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromReceiptType 通过receipt_type获取内容 发票类型 ELECTRO：电子普通发票，VATORDINARY：增值税普通发票
func (obj *_EsMemberReceiptMgr) GetFromReceiptType(receiptType string) (results []*models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`receipt_type` = ?", receiptType).Find(&results).Error

	return
}

// GetBatchFromReceiptType 批量查找 发票类型 ELECTRO：电子普通发票，VATORDINARY：增值税普通发票
func (obj *_EsMemberReceiptMgr) GetBatchFromReceiptType(receiptTypes []string) (results []*models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`receipt_type` IN (?)", receiptTypes).Find(&results).Error

	return
}

// GetFromReceiptTitle 通过receipt_title获取内容 发票抬头
func (obj *_EsMemberReceiptMgr) GetFromReceiptTitle(receiptTitle string) (results []*models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`receipt_title` = ?", receiptTitle).Find(&results).Error

	return
}

// GetBatchFromReceiptTitle 批量查找 发票抬头
func (obj *_EsMemberReceiptMgr) GetBatchFromReceiptTitle(receiptTitles []string) (results []*models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`receipt_title` IN (?)", receiptTitles).Find(&results).Error

	return
}

// GetFromReceiptContent 通过receipt_content获取内容 发票内容
func (obj *_EsMemberReceiptMgr) GetFromReceiptContent(receiptContent string) (results []*models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`receipt_content` = ?", receiptContent).Find(&results).Error

	return
}

// GetBatchFromReceiptContent 批量查找 发票内容
func (obj *_EsMemberReceiptMgr) GetBatchFromReceiptContent(receiptContents []string) (results []*models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`receipt_content` IN (?)", receiptContents).Find(&results).Error

	return
}

// GetFromTaxNo 通过tax_no获取内容 纳税人识别号
func (obj *_EsMemberReceiptMgr) GetFromTaxNo(taxNo string) (results []*models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`tax_no` = ?", taxNo).Find(&results).Error

	return
}

// GetBatchFromTaxNo 批量查找 纳税人识别号
func (obj *_EsMemberReceiptMgr) GetBatchFromTaxNo(taxNos []string) (results []*models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`tax_no` IN (?)", taxNos).Find(&results).Error

	return
}

// GetFromMemberMobile 通过member_mobile获取内容 收票人手机号
func (obj *_EsMemberReceiptMgr) GetFromMemberMobile(memberMobile string) (results []*models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`member_mobile` = ?", memberMobile).Find(&results).Error

	return
}

// GetBatchFromMemberMobile 批量查找 收票人手机号
func (obj *_EsMemberReceiptMgr) GetBatchFromMemberMobile(memberMobiles []string) (results []*models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`member_mobile` IN (?)", memberMobiles).Find(&results).Error

	return
}

// GetFromMemberEmail 通过member_email获取内容 收票人邮箱
func (obj *_EsMemberReceiptMgr) GetFromMemberEmail(memberEmail string) (results []*models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`member_email` = ?", memberEmail).Find(&results).Error

	return
}

// GetBatchFromMemberEmail 批量查找 收票人邮箱
func (obj *_EsMemberReceiptMgr) GetBatchFromMemberEmail(memberEmails []string) (results []*models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`member_email` IN (?)", memberEmails).Find(&results).Error

	return
}

// GetFromIsDefault 通过is_default获取内容 是否为默认选项 0：否，1：是
func (obj *_EsMemberReceiptMgr) GetFromIsDefault(isDefault int16) (results []*models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`is_default` = ?", isDefault).Find(&results).Error

	return
}

// GetBatchFromIsDefault 批量查找 是否为默认选项 0：否，1：是
func (obj *_EsMemberReceiptMgr) GetBatchFromIsDefault(isDefaults []int16) (results []*models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`is_default` IN (?)", isDefaults).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsMemberReceiptMgr) FetchByPrimaryKey(receiptID int) (result models.EsMemberReceipt, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberReceipt{}).Where("`receipt_id` = ?", receiptID).First(&result).Error

	return
}
