package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type DistributionOrderMgr struct {
	*_BaseMgr
}

// NewDistributionOrderMgr open func
func NewDistributionOrderMgr(db db.Repo) *DistributionOrderMgr {
	if db == nil {
		panic(fmt.Errorf("NewDistributionOrderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &DistributionOrderMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_distribution_order"),
		wdb:       db.GetDbW().Table("es_distribution_order"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *DistributionOrderMgr) GetTableName() string {
	return "es_distribution_order"
}

// Reset 重置gorm会话
func (obj *DistributionOrderMgr) Reset() *DistributionOrderMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *DistributionOrderMgr) Get() (result models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *DistributionOrderMgr) Gets() (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *DistributionOrderMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *DistributionOrderMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithOrderID order_id获取 订单id
func (obj *DistributionOrderMgr) WithOrderID(orderID int64) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithOrderSn order_sn获取 订单sn
func (obj *DistributionOrderMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithBuyerMemberID buyer_member_id获取 购买会员id
func (obj *DistributionOrderMgr) WithBuyerMemberID(buyerMemberID int64) Option {
	return optionFunc(func(o *options) { o.query["buyer_member_id"] = buyerMemberID })
}

// WithBuyerMemberName buyer_member_name获取 购买会员名称
func (obj *DistributionOrderMgr) WithBuyerMemberName(buyerMemberName string) Option {
	return optionFunc(func(o *options) { o.query["buyer_member_name"] = buyerMemberName })
}

// WithMemberIDLv1 member_id_lv1获取 一级分销商id
func (obj *DistributionOrderMgr) WithMemberIDLv1(memberIDLv1 int64) Option {
	return optionFunc(func(o *options) { o.query["member_id_lv1"] = memberIDLv1 })
}

// WithMemberIDLv2 member_id_lv2获取 二级分销商id
func (obj *DistributionOrderMgr) WithMemberIDLv2(memberIDLv2 int64) Option {
	return optionFunc(func(o *options) { o.query["member_id_lv2"] = memberIDLv2 })
}

// WithBillID bill_id获取 结算单id
func (obj *DistributionOrderMgr) WithBillID(billID int64) Option {
	return optionFunc(func(o *options) { o.query["bill_id"] = billID })
}

// WithSettleCycle settle_cycle获取 解冻日期
func (obj *DistributionOrderMgr) WithSettleCycle(settleCycle int64) Option {
	return optionFunc(func(o *options) { o.query["settle_cycle"] = settleCycle })
}

// WithCreateTime create_time获取 订单创建时间
func (obj *DistributionOrderMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithOrderPrice order_price获取 订单金额
func (obj *DistributionOrderMgr) WithOrderPrice(orderPrice float64) Option {
	return optionFunc(func(o *options) { o.query["order_price"] = orderPrice })
}

// WithGrade1Rebate grade1_rebate获取 1级提成金额
func (obj *DistributionOrderMgr) WithGrade1Rebate(grade1Rebate float64) Option {
	return optionFunc(func(o *options) { o.query["grade1_rebate"] = grade1Rebate })
}

// WithGrade2Rebate grade2_rebate获取 2级提成金额
func (obj *DistributionOrderMgr) WithGrade2Rebate(grade2Rebate float64) Option {
	return optionFunc(func(o *options) { o.query["grade2_rebate"] = grade2Rebate })
}

// WithGrade1SellbackPrice grade1_sellback_price获取 1级退款金额
func (obj *DistributionOrderMgr) WithGrade1SellbackPrice(grade1SellbackPrice float64) Option {
	return optionFunc(func(o *options) { o.query["grade1_sellback_price"] = grade1SellbackPrice })
}

// WithGrade2SellbackPrice grade2_sellback_price获取 2级退款金额
func (obj *DistributionOrderMgr) WithGrade2SellbackPrice(grade2SellbackPrice float64) Option {
	return optionFunc(func(o *options) { o.query["grade2_sellback_price"] = grade2SellbackPrice })
}

// WithIsReturn is_return获取 是否退货
func (obj *DistributionOrderMgr) WithIsReturn(isReturn int16) Option {
	return optionFunc(func(o *options) { o.query["is_return"] = isReturn })
}

// WithReturnMoney return_money获取 退款金额
func (obj *DistributionOrderMgr) WithReturnMoney(returnMoney float64) Option {
	return optionFunc(func(o *options) { o.query["return_money"] = returnMoney })
}

// WithIsWithdraw is_withdraw获取 是否已经提现
func (obj *DistributionOrderMgr) WithIsWithdraw(isWithdraw int16) Option {
	return optionFunc(func(o *options) { o.query["is_withdraw"] = isWithdraw })
}

// WithSellerID seller_id获取 店铺id
func (obj *DistributionOrderMgr) WithSellerID(sellerID int64) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithSellerName seller_name获取 店铺名称
func (obj *DistributionOrderMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithLv1Point lv1_point获取 一级饭点比例
func (obj *DistributionOrderMgr) WithLv1Point(lv1Point float64) Option {
	return optionFunc(func(o *options) { o.query["lv1_point"] = lv1Point })
}

// WithLv2Point lv2_point获取 二级返点比例
func (obj *DistributionOrderMgr) WithLv2Point(lv2Point float64) Option {
	return optionFunc(func(o *options) { o.query["lv2_point"] = lv2Point })
}

// WithGoodsRebate goods_rebate获取 商品返现描述详细
func (obj *DistributionOrderMgr) WithGoodsRebate(goodsRebate string) Option {
	return optionFunc(func(o *options) { o.query["goods_rebate"] = goodsRebate })
}

// GetByOption 功能选项模式获取
func (obj *DistributionOrderMgr) GetByOption(opts ...Option) (result models.EsDistributionOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *DistributionOrderMgr) GetByOptions(opts ...Option) (results []*models.EsDistributionOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *DistributionOrderMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsDistributionOrder, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where(options.query)
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
func (obj *DistributionOrderMgr) GetFromID(id int64) (result models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *DistributionOrderMgr) GetBatchFromID(ids []int64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromOrderID 通过order_id获取内容 订单id
func (obj *DistributionOrderMgr) GetFromOrderID(orderID int64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`order_id` = ?", orderID).Find(&results).Error

	return
}

// GetBatchFromOrderID 批量查找 订单id
func (obj *DistributionOrderMgr) GetBatchFromOrderID(orderIDs []int64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单sn
func (obj *DistributionOrderMgr) GetFromOrderSn(orderSn string) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单sn
func (obj *DistributionOrderMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromBuyerMemberID 通过buyer_member_id获取内容 购买会员id
func (obj *DistributionOrderMgr) GetFromBuyerMemberID(buyerMemberID int64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`buyer_member_id` = ?", buyerMemberID).Find(&results).Error

	return
}

// GetBatchFromBuyerMemberID 批量查找 购买会员id
func (obj *DistributionOrderMgr) GetBatchFromBuyerMemberID(buyerMemberIDs []int64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`buyer_member_id` IN (?)", buyerMemberIDs).Find(&results).Error

	return
}

// GetFromBuyerMemberName 通过buyer_member_name获取内容 购买会员名称
func (obj *DistributionOrderMgr) GetFromBuyerMemberName(buyerMemberName string) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`buyer_member_name` = ?", buyerMemberName).Find(&results).Error

	return
}

// GetBatchFromBuyerMemberName 批量查找 购买会员名称
func (obj *DistributionOrderMgr) GetBatchFromBuyerMemberName(buyerMemberNames []string) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`buyer_member_name` IN (?)", buyerMemberNames).Find(&results).Error

	return
}

// GetFromMemberIDLv1 通过member_id_lv1获取内容 一级分销商id
func (obj *DistributionOrderMgr) GetFromMemberIDLv1(memberIDLv1 int64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`member_id_lv1` = ?", memberIDLv1).Find(&results).Error

	return
}

// GetBatchFromMemberIDLv1 批量查找 一级分销商id
func (obj *DistributionOrderMgr) GetBatchFromMemberIDLv1(memberIDLv1s []int64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`member_id_lv1` IN (?)", memberIDLv1s).Find(&results).Error

	return
}

// GetFromMemberIDLv2 通过member_id_lv2获取内容 二级分销商id
func (obj *DistributionOrderMgr) GetFromMemberIDLv2(memberIDLv2 int64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`member_id_lv2` = ?", memberIDLv2).Find(&results).Error

	return
}

// GetBatchFromMemberIDLv2 批量查找 二级分销商id
func (obj *DistributionOrderMgr) GetBatchFromMemberIDLv2(memberIDLv2s []int64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`member_id_lv2` IN (?)", memberIDLv2s).Find(&results).Error

	return
}

// GetFromBillID 通过bill_id获取内容 结算单id
func (obj *DistributionOrderMgr) GetFromBillID(billID int64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`bill_id` = ?", billID).Find(&results).Error

	return
}

// GetBatchFromBillID 批量查找 结算单id
func (obj *DistributionOrderMgr) GetBatchFromBillID(billIDs []int64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`bill_id` IN (?)", billIDs).Find(&results).Error

	return
}

// GetFromSettleCycle 通过settle_cycle获取内容 解冻日期
func (obj *DistributionOrderMgr) GetFromSettleCycle(settleCycle int64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`settle_cycle` = ?", settleCycle).Find(&results).Error

	return
}

// GetBatchFromSettleCycle 批量查找 解冻日期
func (obj *DistributionOrderMgr) GetBatchFromSettleCycle(settleCycles []int64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`settle_cycle` IN (?)", settleCycles).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 订单创建时间
func (obj *DistributionOrderMgr) GetFromCreateTime(createTime int64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 订单创建时间
func (obj *DistributionOrderMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromOrderPrice 通过order_price获取内容 订单金额
func (obj *DistributionOrderMgr) GetFromOrderPrice(orderPrice float64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`order_price` = ?", orderPrice).Find(&results).Error

	return
}

// GetBatchFromOrderPrice 批量查找 订单金额
func (obj *DistributionOrderMgr) GetBatchFromOrderPrice(orderPrices []float64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`order_price` IN (?)", orderPrices).Find(&results).Error

	return
}

// GetFromGrade1Rebate 通过grade1_rebate获取内容 1级提成金额
func (obj *DistributionOrderMgr) GetFromGrade1Rebate(grade1Rebate float64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`grade1_rebate` = ?", grade1Rebate).Find(&results).Error

	return
}

// GetBatchFromGrade1Rebate 批量查找 1级提成金额
func (obj *DistributionOrderMgr) GetBatchFromGrade1Rebate(grade1Rebates []float64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`grade1_rebate` IN (?)", grade1Rebates).Find(&results).Error

	return
}

// GetFromGrade2Rebate 通过grade2_rebate获取内容 2级提成金额
func (obj *DistributionOrderMgr) GetFromGrade2Rebate(grade2Rebate float64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`grade2_rebate` = ?", grade2Rebate).Find(&results).Error

	return
}

// GetBatchFromGrade2Rebate 批量查找 2级提成金额
func (obj *DistributionOrderMgr) GetBatchFromGrade2Rebate(grade2Rebates []float64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`grade2_rebate` IN (?)", grade2Rebates).Find(&results).Error

	return
}

// GetFromGrade1SellbackPrice 通过grade1_sellback_price获取内容 1级退款金额
func (obj *DistributionOrderMgr) GetFromGrade1SellbackPrice(grade1SellbackPrice float64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`grade1_sellback_price` = ?", grade1SellbackPrice).Find(&results).Error

	return
}

// GetBatchFromGrade1SellbackPrice 批量查找 1级退款金额
func (obj *DistributionOrderMgr) GetBatchFromGrade1SellbackPrice(grade1SellbackPrices []float64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`grade1_sellback_price` IN (?)", grade1SellbackPrices).Find(&results).Error

	return
}

// GetFromGrade2SellbackPrice 通过grade2_sellback_price获取内容 2级退款金额
func (obj *DistributionOrderMgr) GetFromGrade2SellbackPrice(grade2SellbackPrice float64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`grade2_sellback_price` = ?", grade2SellbackPrice).Find(&results).Error

	return
}

// GetBatchFromGrade2SellbackPrice 批量查找 2级退款金额
func (obj *DistributionOrderMgr) GetBatchFromGrade2SellbackPrice(grade2SellbackPrices []float64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`grade2_sellback_price` IN (?)", grade2SellbackPrices).Find(&results).Error

	return
}

// GetFromIsReturn 通过is_return获取内容 是否退货
func (obj *DistributionOrderMgr) GetFromIsReturn(isReturn int16) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`is_return` = ?", isReturn).Find(&results).Error

	return
}

// GetBatchFromIsReturn 批量查找 是否退货
func (obj *DistributionOrderMgr) GetBatchFromIsReturn(isReturns []int16) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`is_return` IN (?)", isReturns).Find(&results).Error

	return
}

// GetFromReturnMoney 通过return_money获取内容 退款金额
func (obj *DistributionOrderMgr) GetFromReturnMoney(returnMoney float64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`return_money` = ?", returnMoney).Find(&results).Error

	return
}

// GetBatchFromReturnMoney 批量查找 退款金额
func (obj *DistributionOrderMgr) GetBatchFromReturnMoney(returnMoneys []float64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`return_money` IN (?)", returnMoneys).Find(&results).Error

	return
}

// GetFromIsWithdraw 通过is_withdraw获取内容 是否已经提现
func (obj *DistributionOrderMgr) GetFromIsWithdraw(isWithdraw int16) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`is_withdraw` = ?", isWithdraw).Find(&results).Error

	return
}

// GetBatchFromIsWithdraw 批量查找 是否已经提现
func (obj *DistributionOrderMgr) GetBatchFromIsWithdraw(isWithdraws []int16) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`is_withdraw` IN (?)", isWithdraws).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 店铺id
func (obj *DistributionOrderMgr) GetFromSellerID(sellerID int64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 店铺id
func (obj *DistributionOrderMgr) GetBatchFromSellerID(sellerIDs []int64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 店铺名称
func (obj *DistributionOrderMgr) GetFromSellerName(sellerName string) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 店铺名称
func (obj *DistributionOrderMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromLv1Point 通过lv1_point获取内容 一级饭点比例
func (obj *DistributionOrderMgr) GetFromLv1Point(lv1Point float64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`lv1_point` = ?", lv1Point).Find(&results).Error

	return
}

// GetBatchFromLv1Point 批量查找 一级饭点比例
func (obj *DistributionOrderMgr) GetBatchFromLv1Point(lv1Points []float64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`lv1_point` IN (?)", lv1Points).Find(&results).Error

	return
}

// GetFromLv2Point 通过lv2_point获取内容 二级返点比例
func (obj *DistributionOrderMgr) GetFromLv2Point(lv2Point float64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`lv2_point` = ?", lv2Point).Find(&results).Error

	return
}

// GetBatchFromLv2Point 批量查找 二级返点比例
func (obj *DistributionOrderMgr) GetBatchFromLv2Point(lv2Points []float64) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`lv2_point` IN (?)", lv2Points).Find(&results).Error

	return
}

// GetFromGoodsRebate 通过goods_rebate获取内容 商品返现描述详细
func (obj *DistributionOrderMgr) GetFromGoodsRebate(goodsRebate string) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`goods_rebate` = ?", goodsRebate).Find(&results).Error

	return
}

// GetBatchFromGoodsRebate 批量查找 商品返现描述详细
func (obj *DistributionOrderMgr) GetBatchFromGoodsRebate(goodsRebates []string) (results []*models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`goods_rebate` IN (?)", goodsRebates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *DistributionOrderMgr) FetchByPrimaryKey(id int64) (result models.EsDistributionOrder, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionOrder{}).Where("`id` = ?", id).First(&result).Error

	return
}
