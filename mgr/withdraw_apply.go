package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsWithdrawApplyMgr struct {
	*_BaseMgr
}

// EsWithdrawApplyMgr open func
func EsWithdrawApplyMgr(db *gorm.DB) *_EsWithdrawApplyMgr {
	if db == nil {
		panic(fmt.Errorf("EsWithdrawApplyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsWithdrawApplyMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_withdraw_apply"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsWithdrawApplyMgr) GetTableName() string {
	return "es_withdraw_apply"
}

// Reset 重置gorm会话
func (obj *_EsWithdrawApplyMgr) Reset() *_EsWithdrawApplyMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsWithdrawApplyMgr) Get() (result models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsWithdrawApplyMgr) Gets() (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsWithdrawApplyMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_EsWithdrawApplyMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithApplyMoney apply_money获取 提现金额
func (obj *_EsWithdrawApplyMgr) WithApplyMoney(applyMoney float64) Option {
	return optionFunc(func(o *options) { o.query["apply_money"] = applyMoney })
}

// WithMemberID member_id获取 会员id
func (obj *_EsWithdrawApplyMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMemberName member_name获取 会员名称
func (obj *_EsWithdrawApplyMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithApplyRemark apply_remark获取 申请备注
func (obj *_EsWithdrawApplyMgr) WithApplyRemark(applyRemark string) Option {
	return optionFunc(func(o *options) { o.query["apply_remark"] = applyRemark })
}

// WithInspectRemark inspect_remark获取 审核备注
func (obj *_EsWithdrawApplyMgr) WithInspectRemark(inspectRemark string) Option {
	return optionFunc(func(o *options) { o.query["inspect_remark"] = inspectRemark })
}

// WithTransferRemark transfer_remark获取 转账备注
func (obj *_EsWithdrawApplyMgr) WithTransferRemark(transferRemark string) Option {
	return optionFunc(func(o *options) { o.query["transfer_remark"] = transferRemark })
}

// WithApplyTime apply_time获取 申请时间
func (obj *_EsWithdrawApplyMgr) WithApplyTime(applyTime int) Option {
	return optionFunc(func(o *options) { o.query["apply_time"] = applyTime })
}

// WithInspectTime inspect_time获取 审核时间
func (obj *_EsWithdrawApplyMgr) WithInspectTime(inspectTime int) Option {
	return optionFunc(func(o *options) { o.query["inspect_time"] = inspectTime })
}

// WithTransferTime transfer_time获取 转账时间
func (obj *_EsWithdrawApplyMgr) WithTransferTime(transferTime int) Option {
	return optionFunc(func(o *options) { o.query["transfer_time"] = transferTime })
}

// WithStatus status获取 状态
func (obj *_EsWithdrawApplyMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithSn sn获取 编号
func (obj *_EsWithdrawApplyMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithIP ip获取 ip地址
func (obj *_EsWithdrawApplyMgr) WithIP(ip string) Option {
	return optionFunc(func(o *options) { o.query["ip"] = ip })
}

// GetByOption 功能选项模式获取
func (obj *_EsWithdrawApplyMgr) GetByOption(opts ...Option) (result models.EsWithdrawApply, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsWithdrawApplyMgr) GetByOptions(opts ...Option) (results []*models.EsWithdrawApply, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsWithdrawApplyMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsWithdrawApply, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where(options.query)
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
func (obj *_EsWithdrawApplyMgr) GetFromID(id int) (result models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_EsWithdrawApplyMgr) GetBatchFromID(ids []int) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromApplyMoney 通过apply_money获取内容 提现金额
func (obj *_EsWithdrawApplyMgr) GetFromApplyMoney(applyMoney float64) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`apply_money` = ?", applyMoney).Find(&results).Error

	return
}

// GetBatchFromApplyMoney 批量查找 提现金额
func (obj *_EsWithdrawApplyMgr) GetBatchFromApplyMoney(applyMoneys []float64) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`apply_money` IN (?)", applyMoneys).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_EsWithdrawApplyMgr) GetFromMemberID(memberID int) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_EsWithdrawApplyMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 会员名称
func (obj *_EsWithdrawApplyMgr) GetFromMemberName(memberName string) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 会员名称
func (obj *_EsWithdrawApplyMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromApplyRemark 通过apply_remark获取内容 申请备注
func (obj *_EsWithdrawApplyMgr) GetFromApplyRemark(applyRemark string) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`apply_remark` = ?", applyRemark).Find(&results).Error

	return
}

// GetBatchFromApplyRemark 批量查找 申请备注
func (obj *_EsWithdrawApplyMgr) GetBatchFromApplyRemark(applyRemarks []string) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`apply_remark` IN (?)", applyRemarks).Find(&results).Error

	return
}

// GetFromInspectRemark 通过inspect_remark获取内容 审核备注
func (obj *_EsWithdrawApplyMgr) GetFromInspectRemark(inspectRemark string) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`inspect_remark` = ?", inspectRemark).Find(&results).Error

	return
}

// GetBatchFromInspectRemark 批量查找 审核备注
func (obj *_EsWithdrawApplyMgr) GetBatchFromInspectRemark(inspectRemarks []string) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`inspect_remark` IN (?)", inspectRemarks).Find(&results).Error

	return
}

// GetFromTransferRemark 通过transfer_remark获取内容 转账备注
func (obj *_EsWithdrawApplyMgr) GetFromTransferRemark(transferRemark string) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`transfer_remark` = ?", transferRemark).Find(&results).Error

	return
}

// GetBatchFromTransferRemark 批量查找 转账备注
func (obj *_EsWithdrawApplyMgr) GetBatchFromTransferRemark(transferRemarks []string) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`transfer_remark` IN (?)", transferRemarks).Find(&results).Error

	return
}

// GetFromApplyTime 通过apply_time获取内容 申请时间
func (obj *_EsWithdrawApplyMgr) GetFromApplyTime(applyTime int) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`apply_time` = ?", applyTime).Find(&results).Error

	return
}

// GetBatchFromApplyTime 批量查找 申请时间
func (obj *_EsWithdrawApplyMgr) GetBatchFromApplyTime(applyTimes []int) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`apply_time` IN (?)", applyTimes).Find(&results).Error

	return
}

// GetFromInspectTime 通过inspect_time获取内容 审核时间
func (obj *_EsWithdrawApplyMgr) GetFromInspectTime(inspectTime int) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`inspect_time` = ?", inspectTime).Find(&results).Error

	return
}

// GetBatchFromInspectTime 批量查找 审核时间
func (obj *_EsWithdrawApplyMgr) GetBatchFromInspectTime(inspectTimes []int) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`inspect_time` IN (?)", inspectTimes).Find(&results).Error

	return
}

// GetFromTransferTime 通过transfer_time获取内容 转账时间
func (obj *_EsWithdrawApplyMgr) GetFromTransferTime(transferTime int) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`transfer_time` = ?", transferTime).Find(&results).Error

	return
}

// GetBatchFromTransferTime 批量查找 转账时间
func (obj *_EsWithdrawApplyMgr) GetBatchFromTransferTime(transferTimes []int) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`transfer_time` IN (?)", transferTimes).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态
func (obj *_EsWithdrawApplyMgr) GetFromStatus(status string) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态
func (obj *_EsWithdrawApplyMgr) GetBatchFromStatus(statuss []string) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 编号
func (obj *_EsWithdrawApplyMgr) GetFromSn(sn string) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 编号
func (obj *_EsWithdrawApplyMgr) GetBatchFromSn(sns []string) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromIP 通过ip获取内容 ip地址
func (obj *_EsWithdrawApplyMgr) GetFromIP(ip string) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`ip` = ?", ip).Find(&results).Error

	return
}

// GetBatchFromIP 批量查找 ip地址
func (obj *_EsWithdrawApplyMgr) GetBatchFromIP(ips []string) (results []*models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`ip` IN (?)", ips).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsWithdrawApplyMgr) FetchByPrimaryKey(id int) (result models.EsWithdrawApply, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsWithdrawApply{}).Where("`id` = ?", id).First(&result).Error

	return
}
