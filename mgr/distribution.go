package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type DistributionMgr struct {
	*_BaseMgr
}

// NewDistributionMgr open func
func NewDistributionMgr(db db.Repo) *DistributionMgr {
	if db == nil {
		panic(fmt.Errorf("NewDistributionMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &DistributionMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_distribution"),
		wdb:       db.GetDbW().Table("es_distribution"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *DistributionMgr) GetTableName() string {
	return "es_distribution"
}

// Reset 重置gorm会话
func (obj *DistributionMgr) Reset() *DistributionMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *DistributionMgr) Get() (result models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *DistributionMgr) Gets() (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *DistributionMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *DistributionMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员id
func (obj *DistributionMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMemberName member_name获取 会员名称
func (obj *DistributionMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithPath path获取 关系路径
func (obj *DistributionMgr) WithPath(path string) Option {
	return optionFunc(func(o *options) { o.query["path"] = path })
}

// WithMemberIDLv2 member_id_lv2获取 上上级
func (obj *DistributionMgr) WithMemberIDLv2(memberIDLv2 int64) Option {
	return optionFunc(func(o *options) { o.query["member_id_lv2"] = memberIDLv2 })
}

// WithMemberIDLv1 member_id_lv1获取 上级
func (obj *DistributionMgr) WithMemberIDLv1(memberIDLv1 int64) Option {
	return optionFunc(func(o *options) { o.query["member_id_lv1"] = memberIDLv1 })
}

// WithDownline downline获取 下线人数
func (obj *DistributionMgr) WithDownline(downline int64) Option {
	return optionFunc(func(o *options) { o.query["downline"] = downline })
}

// WithOrderNum order_num获取 订单数
func (obj *DistributionMgr) WithOrderNum(orderNum int64) Option {
	return optionFunc(func(o *options) { o.query["order_num"] = orderNum })
}

// WithRebateTotal rebate_total获取 返利总额
func (obj *DistributionMgr) WithRebateTotal(rebateTotal float64) Option {
	return optionFunc(func(o *options) { o.query["rebate_total"] = rebateTotal })
}

// WithTurnoverPrice turnover_price获取 营业额总额
func (obj *DistributionMgr) WithTurnoverPrice(turnoverPrice float64) Option {
	return optionFunc(func(o *options) { o.query["turnover_price"] = turnoverPrice })
}

// WithCanRebate can_rebate获取 可提现金额
func (obj *DistributionMgr) WithCanRebate(canRebate float64) Option {
	return optionFunc(func(o *options) { o.query["can_rebate"] = canRebate })
}

// WithCommissionFrozen commission_frozen获取 返利金额冻结
func (obj *DistributionMgr) WithCommissionFrozen(commissionFrozen float64) Option {
	return optionFunc(func(o *options) { o.query["commission_frozen"] = commissionFrozen })
}

// WithCurrentTplName current_tpl_name获取 使用模版名称
func (obj *DistributionMgr) WithCurrentTplName(currentTplName string) Option {
	return optionFunc(func(o *options) { o.query["current_tpl_name"] = currentTplName })
}

// WithCurrentTplID current_tpl_id获取 使用模版
func (obj *DistributionMgr) WithCurrentTplID(currentTplID int) Option {
	return optionFunc(func(o *options) { o.query["current_tpl_id"] = currentTplID })
}

// WithWithdrawFrozenPrice withdraw_frozen_price获取 提现冻结
func (obj *DistributionMgr) WithWithdrawFrozenPrice(withdrawFrozenPrice float64) Option {
	return optionFunc(func(o *options) { o.query["withdraw_frozen_price"] = withdrawFrozenPrice })
}

// WithCreateTime create_time获取 创建时间
func (obj *DistributionMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *DistributionMgr) GetByOption(opts ...Option) (result models.EsDistribution, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *DistributionMgr) GetByOptions(opts ...Option) (results []*models.EsDistribution, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *DistributionMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsDistribution, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where(options.query)
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
func (obj *DistributionMgr) GetFromID(id int64) (result models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *DistributionMgr) GetBatchFromID(ids []int64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *DistributionMgr) GetFromMemberID(memberID int) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *DistributionMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 会员名称
func (obj *DistributionMgr) GetFromMemberName(memberName string) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 会员名称
func (obj *DistributionMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromPath 通过path获取内容 关系路径
func (obj *DistributionMgr) GetFromPath(path string) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`path` = ?", path).Find(&results).Error

	return
}

// GetBatchFromPath 批量查找 关系路径
func (obj *DistributionMgr) GetBatchFromPath(paths []string) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`path` IN (?)", paths).Find(&results).Error

	return
}

// GetFromMemberIDLv2 通过member_id_lv2获取内容 上上级
func (obj *DistributionMgr) GetFromMemberIDLv2(memberIDLv2 int64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`member_id_lv2` = ?", memberIDLv2).Find(&results).Error

	return
}

// GetBatchFromMemberIDLv2 批量查找 上上级
func (obj *DistributionMgr) GetBatchFromMemberIDLv2(memberIDLv2s []int64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`member_id_lv2` IN (?)", memberIDLv2s).Find(&results).Error

	return
}

// GetFromMemberIDLv1 通过member_id_lv1获取内容 上级
func (obj *DistributionMgr) GetFromMemberIDLv1(memberIDLv1 int64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`member_id_lv1` = ?", memberIDLv1).Find(&results).Error

	return
}

// GetBatchFromMemberIDLv1 批量查找 上级
func (obj *DistributionMgr) GetBatchFromMemberIDLv1(memberIDLv1s []int64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`member_id_lv1` IN (?)", memberIDLv1s).Find(&results).Error

	return
}

// GetFromDownline 通过downline获取内容 下线人数
func (obj *DistributionMgr) GetFromDownline(downline int64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`downline` = ?", downline).Find(&results).Error

	return
}

// GetBatchFromDownline 批量查找 下线人数
func (obj *DistributionMgr) GetBatchFromDownline(downlines []int64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`downline` IN (?)", downlines).Find(&results).Error

	return
}

// GetFromOrderNum 通过order_num获取内容 订单数
func (obj *DistributionMgr) GetFromOrderNum(orderNum int64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`order_num` = ?", orderNum).Find(&results).Error

	return
}

// GetBatchFromOrderNum 批量查找 订单数
func (obj *DistributionMgr) GetBatchFromOrderNum(orderNums []int64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`order_num` IN (?)", orderNums).Find(&results).Error

	return
}

// GetFromRebateTotal 通过rebate_total获取内容 返利总额
func (obj *DistributionMgr) GetFromRebateTotal(rebateTotal float64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`rebate_total` = ?", rebateTotal).Find(&results).Error

	return
}

// GetBatchFromRebateTotal 批量查找 返利总额
func (obj *DistributionMgr) GetBatchFromRebateTotal(rebateTotals []float64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`rebate_total` IN (?)", rebateTotals).Find(&results).Error

	return
}

// GetFromTurnoverPrice 通过turnover_price获取内容 营业额总额
func (obj *DistributionMgr) GetFromTurnoverPrice(turnoverPrice float64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`turnover_price` = ?", turnoverPrice).Find(&results).Error

	return
}

// GetBatchFromTurnoverPrice 批量查找 营业额总额
func (obj *DistributionMgr) GetBatchFromTurnoverPrice(turnoverPrices []float64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`turnover_price` IN (?)", turnoverPrices).Find(&results).Error

	return
}

// GetFromCanRebate 通过can_rebate获取内容 可提现金额
func (obj *DistributionMgr) GetFromCanRebate(canRebate float64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`can_rebate` = ?", canRebate).Find(&results).Error

	return
}

// GetBatchFromCanRebate 批量查找 可提现金额
func (obj *DistributionMgr) GetBatchFromCanRebate(canRebates []float64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`can_rebate` IN (?)", canRebates).Find(&results).Error

	return
}

// GetFromCommissionFrozen 通过commission_frozen获取内容 返利金额冻结
func (obj *DistributionMgr) GetFromCommissionFrozen(commissionFrozen float64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`commission_frozen` = ?", commissionFrozen).Find(&results).Error

	return
}

// GetBatchFromCommissionFrozen 批量查找 返利金额冻结
func (obj *DistributionMgr) GetBatchFromCommissionFrozen(commissionFrozens []float64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`commission_frozen` IN (?)", commissionFrozens).Find(&results).Error

	return
}

// GetFromCurrentTplName 通过current_tpl_name获取内容 使用模版名称
func (obj *DistributionMgr) GetFromCurrentTplName(currentTplName string) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`current_tpl_name` = ?", currentTplName).Find(&results).Error

	return
}

// GetBatchFromCurrentTplName 批量查找 使用模版名称
func (obj *DistributionMgr) GetBatchFromCurrentTplName(currentTplNames []string) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`current_tpl_name` IN (?)", currentTplNames).Find(&results).Error

	return
}

// GetFromCurrentTplID 通过current_tpl_id获取内容 使用模版
func (obj *DistributionMgr) GetFromCurrentTplID(currentTplID int) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`current_tpl_id` = ?", currentTplID).Find(&results).Error

	return
}

// GetBatchFromCurrentTplID 批量查找 使用模版
func (obj *DistributionMgr) GetBatchFromCurrentTplID(currentTplIDs []int) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`current_tpl_id` IN (?)", currentTplIDs).Find(&results).Error

	return
}

// GetFromWithdrawFrozenPrice 通过withdraw_frozen_price获取内容 提现冻结
func (obj *DistributionMgr) GetFromWithdrawFrozenPrice(withdrawFrozenPrice float64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`withdraw_frozen_price` = ?", withdrawFrozenPrice).Find(&results).Error

	return
}

// GetBatchFromWithdrawFrozenPrice 批量查找 提现冻结
func (obj *DistributionMgr) GetBatchFromWithdrawFrozenPrice(withdrawFrozenPrices []float64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`withdraw_frozen_price` IN (?)", withdrawFrozenPrices).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *DistributionMgr) GetFromCreateTime(createTime int64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *DistributionMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *DistributionMgr) FetchByPrimaryKey(id int64) (result models.EsDistribution, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistribution{}).Where("`id` = ?", id).First(&result).Error

	return
}
