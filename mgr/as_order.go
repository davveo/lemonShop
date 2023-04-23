package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsAsOrderMgr struct {
	*_BaseMgr
}

// EsAsOrderMgr open func
func EsAsOrderMgr(db *gorm.DB) *_EsAsOrderMgr {
	if db == nil {
		panic(fmt.Errorf("EsAsOrderMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsAsOrderMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_as_order"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsAsOrderMgr) GetTableName() string {
	return "es_as_order"
}

// Reset 重置gorm会话
func (obj *_EsAsOrderMgr) Reset() *_EsAsOrderMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsAsOrderMgr) Get() (result models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsAsOrderMgr) Gets() (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsAsOrderMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *_EsAsOrderMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSn sn获取 售后单号
func (obj *_EsAsOrderMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithOrderSn order_sn获取 订单编号
func (obj *_EsAsOrderMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithMemberID member_id获取 用户ID
func (obj *_EsAsOrderMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMemberName member_name获取 用户名
func (obj *_EsAsOrderMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithSellerID seller_id获取 商家ID
func (obj *_EsAsOrderMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithSellerName seller_name获取 商家名称
func (obj *_EsAsOrderMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithCreateTime create_time获取 创建时间
func (obj *_EsAsOrderMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithMobile mobile获取 手机号
func (obj *_EsAsOrderMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithServiceType service_type获取 售后类型 RETURN_GOODS：退货，CHANGE_GOODS：换货，SUPPLY_AGAIN_GOODS：补发货品，ORDER_CANCEL：取消订单（订单已付款且未收货之前）
func (obj *_EsAsOrderMgr) WithServiceType(serviceType string) Option {
	return optionFunc(func(o *options) { o.query["service_type"] = serviceType })
}

// WithServiceStatus service_status获取 售后单状态 APPLY：待审核，PASS：审核通过，REFUSE：审核拒绝，WAIT_FOR_MANUAL：待人工处理，STOCK_IN：退货入库，REFUNDING：退款中，REFUNDFAIL：退款失败，COMPLETE：完成
func (obj *_EsAsOrderMgr) WithServiceStatus(serviceStatus string) Option {
	return optionFunc(func(o *options) { o.query["service_status"] = serviceStatus })
}

// WithReason reason获取 申请原因
func (obj *_EsAsOrderMgr) WithReason(reason string) Option {
	return optionFunc(func(o *options) { o.query["reason"] = reason })
}

// WithApplyVouchers apply_vouchers获取 申请凭证
func (obj *_EsAsOrderMgr) WithApplyVouchers(applyVouchers string) Option {
	return optionFunc(func(o *options) { o.query["apply_vouchers"] = applyVouchers })
}

// WithProblemDesc problem_desc获取 问题描述
func (obj *_EsAsOrderMgr) WithProblemDesc(problemDesc string) Option {
	return optionFunc(func(o *options) { o.query["problem_desc"] = problemDesc })
}

// WithGoodsJSON goods_json获取 售后商品信息json
func (obj *_EsAsOrderMgr) WithGoodsJSON(goodsJSON string) Option {
	return optionFunc(func(o *options) { o.query["goods_json"] = goodsJSON })
}

// WithDisabled disabled获取 删除状态 DELETED：已删除 NORMAL：正常
func (obj *_EsAsOrderMgr) WithDisabled(disabled string) Option {
	return optionFunc(func(o *options) { o.query["disabled"] = disabled })
}

// WithAuditRemark audit_remark获取 商家审核备注
func (obj *_EsAsOrderMgr) WithAuditRemark(auditRemark string) Option {
	return optionFunc(func(o *options) { o.query["audit_remark"] = auditRemark })
}

// WithStockRemark stock_remark获取 商家入库备注
func (obj *_EsAsOrderMgr) WithStockRemark(stockRemark string) Option {
	return optionFunc(func(o *options) { o.query["stock_remark"] = stockRemark })
}

// WithRefundRemark refund_remark获取 平台退款备注
func (obj *_EsAsOrderMgr) WithRefundRemark(refundRemark string) Option {
	return optionFunc(func(o *options) { o.query["refund_remark"] = refundRemark })
}

// WithCloseReason close_reason获取 取消原因
func (obj *_EsAsOrderMgr) WithCloseReason(closeReason string) Option {
	return optionFunc(func(o *options) { o.query["close_reason"] = closeReason })
}

// WithReturnAddr return_addr获取 退货地址信息
func (obj *_EsAsOrderMgr) WithReturnAddr(returnAddr string) Option {
	return optionFunc(func(o *options) { o.query["return_addr"] = returnAddr })
}

// WithNewOrderSn new_order_sn获取 新订单编号
func (obj *_EsAsOrderMgr) WithNewOrderSn(newOrderSn string) Option {
	return optionFunc(func(o *options) { o.query["new_order_sn"] = newOrderSn })
}

// WithCreateChannel create_channel获取 售后服务单创建渠道 NORMAL：正常渠道创建，PINTUAN：拼团失败自动创建
func (obj *_EsAsOrderMgr) WithCreateChannel(createChannel string) Option {
	return optionFunc(func(o *options) { o.query["create_channel"] = createChannel })
}

// GetByOption 功能选项模式获取
func (obj *_EsAsOrderMgr) GetByOption(opts ...Option) (result models.EsAsOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsAsOrderMgr) GetByOptions(opts ...Option) (results []*models.EsAsOrder, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsAsOrderMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsAsOrder, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where(options.query)
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
func (obj *_EsAsOrderMgr) GetFromID(id int) (result models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *_EsAsOrderMgr) GetBatchFromID(ids []int) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 售后单号
func (obj *_EsAsOrderMgr) GetFromSn(sn string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 售后单号
func (obj *_EsAsOrderMgr) GetBatchFromSn(sns []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *_EsAsOrderMgr) GetFromOrderSn(orderSn string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *_EsAsOrderMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 用户ID
func (obj *_EsAsOrderMgr) GetFromMemberID(memberID int) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 用户ID
func (obj *_EsAsOrderMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 用户名
func (obj *_EsAsOrderMgr) GetFromMemberName(memberName string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 用户名
func (obj *_EsAsOrderMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 商家ID
func (obj *_EsAsOrderMgr) GetFromSellerID(sellerID int) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 商家ID
func (obj *_EsAsOrderMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 商家名称
func (obj *_EsAsOrderMgr) GetFromSellerName(sellerName string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 商家名称
func (obj *_EsAsOrderMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_EsAsOrderMgr) GetFromCreateTime(createTime int64) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_EsAsOrderMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容 手机号
func (obj *_EsAsOrderMgr) GetFromMobile(mobile string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`mobile` = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量查找 手机号
func (obj *_EsAsOrderMgr) GetBatchFromMobile(mobiles []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`mobile` IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromServiceType 通过service_type获取内容 售后类型 RETURN_GOODS：退货，CHANGE_GOODS：换货，SUPPLY_AGAIN_GOODS：补发货品，ORDER_CANCEL：取消订单（订单已付款且未收货之前）
func (obj *_EsAsOrderMgr) GetFromServiceType(serviceType string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`service_type` = ?", serviceType).Find(&results).Error

	return
}

// GetBatchFromServiceType 批量查找 售后类型 RETURN_GOODS：退货，CHANGE_GOODS：换货，SUPPLY_AGAIN_GOODS：补发货品，ORDER_CANCEL：取消订单（订单已付款且未收货之前）
func (obj *_EsAsOrderMgr) GetBatchFromServiceType(serviceTypes []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`service_type` IN (?)", serviceTypes).Find(&results).Error

	return
}

// GetFromServiceStatus 通过service_status获取内容 售后单状态 APPLY：待审核，PASS：审核通过，REFUSE：审核拒绝，WAIT_FOR_MANUAL：待人工处理，STOCK_IN：退货入库，REFUNDING：退款中，REFUNDFAIL：退款失败，COMPLETE：完成
func (obj *_EsAsOrderMgr) GetFromServiceStatus(serviceStatus string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`service_status` = ?", serviceStatus).Find(&results).Error

	return
}

// GetBatchFromServiceStatus 批量查找 售后单状态 APPLY：待审核，PASS：审核通过，REFUSE：审核拒绝，WAIT_FOR_MANUAL：待人工处理，STOCK_IN：退货入库，REFUNDING：退款中，REFUNDFAIL：退款失败，COMPLETE：完成
func (obj *_EsAsOrderMgr) GetBatchFromServiceStatus(serviceStatuss []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`service_status` IN (?)", serviceStatuss).Find(&results).Error

	return
}

// GetFromReason 通过reason获取内容 申请原因
func (obj *_EsAsOrderMgr) GetFromReason(reason string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`reason` = ?", reason).Find(&results).Error

	return
}

// GetBatchFromReason 批量查找 申请原因
func (obj *_EsAsOrderMgr) GetBatchFromReason(reasons []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`reason` IN (?)", reasons).Find(&results).Error

	return
}

// GetFromApplyVouchers 通过apply_vouchers获取内容 申请凭证
func (obj *_EsAsOrderMgr) GetFromApplyVouchers(applyVouchers string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`apply_vouchers` = ?", applyVouchers).Find(&results).Error

	return
}

// GetBatchFromApplyVouchers 批量查找 申请凭证
func (obj *_EsAsOrderMgr) GetBatchFromApplyVouchers(applyVoucherss []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`apply_vouchers` IN (?)", applyVoucherss).Find(&results).Error

	return
}

// GetFromProblemDesc 通过problem_desc获取内容 问题描述
func (obj *_EsAsOrderMgr) GetFromProblemDesc(problemDesc string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`problem_desc` = ?", problemDesc).Find(&results).Error

	return
}

// GetBatchFromProblemDesc 批量查找 问题描述
func (obj *_EsAsOrderMgr) GetBatchFromProblemDesc(problemDescs []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`problem_desc` IN (?)", problemDescs).Find(&results).Error

	return
}

// GetFromGoodsJSON 通过goods_json获取内容 售后商品信息json
func (obj *_EsAsOrderMgr) GetFromGoodsJSON(goodsJSON string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`goods_json` = ?", goodsJSON).Find(&results).Error

	return
}

// GetBatchFromGoodsJSON 批量查找 售后商品信息json
func (obj *_EsAsOrderMgr) GetBatchFromGoodsJSON(goodsJSONs []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`goods_json` IN (?)", goodsJSONs).Find(&results).Error

	return
}

// GetFromDisabled 通过disabled获取内容 删除状态 DELETED：已删除 NORMAL：正常
func (obj *_EsAsOrderMgr) GetFromDisabled(disabled string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`disabled` = ?", disabled).Find(&results).Error

	return
}

// GetBatchFromDisabled 批量查找 删除状态 DELETED：已删除 NORMAL：正常
func (obj *_EsAsOrderMgr) GetBatchFromDisabled(disableds []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`disabled` IN (?)", disableds).Find(&results).Error

	return
}

// GetFromAuditRemark 通过audit_remark获取内容 商家审核备注
func (obj *_EsAsOrderMgr) GetFromAuditRemark(auditRemark string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`audit_remark` = ?", auditRemark).Find(&results).Error

	return
}

// GetBatchFromAuditRemark 批量查找 商家审核备注
func (obj *_EsAsOrderMgr) GetBatchFromAuditRemark(auditRemarks []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`audit_remark` IN (?)", auditRemarks).Find(&results).Error

	return
}

// GetFromStockRemark 通过stock_remark获取内容 商家入库备注
func (obj *_EsAsOrderMgr) GetFromStockRemark(stockRemark string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`stock_remark` = ?", stockRemark).Find(&results).Error

	return
}

// GetBatchFromStockRemark 批量查找 商家入库备注
func (obj *_EsAsOrderMgr) GetBatchFromStockRemark(stockRemarks []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`stock_remark` IN (?)", stockRemarks).Find(&results).Error

	return
}

// GetFromRefundRemark 通过refund_remark获取内容 平台退款备注
func (obj *_EsAsOrderMgr) GetFromRefundRemark(refundRemark string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`refund_remark` = ?", refundRemark).Find(&results).Error

	return
}

// GetBatchFromRefundRemark 批量查找 平台退款备注
func (obj *_EsAsOrderMgr) GetBatchFromRefundRemark(refundRemarks []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`refund_remark` IN (?)", refundRemarks).Find(&results).Error

	return
}

// GetFromCloseReason 通过close_reason获取内容 取消原因
func (obj *_EsAsOrderMgr) GetFromCloseReason(closeReason string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`close_reason` = ?", closeReason).Find(&results).Error

	return
}

// GetBatchFromCloseReason 批量查找 取消原因
func (obj *_EsAsOrderMgr) GetBatchFromCloseReason(closeReasons []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`close_reason` IN (?)", closeReasons).Find(&results).Error

	return
}

// GetFromReturnAddr 通过return_addr获取内容 退货地址信息
func (obj *_EsAsOrderMgr) GetFromReturnAddr(returnAddr string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`return_addr` = ?", returnAddr).Find(&results).Error

	return
}

// GetBatchFromReturnAddr 批量查找 退货地址信息
func (obj *_EsAsOrderMgr) GetBatchFromReturnAddr(returnAddrs []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`return_addr` IN (?)", returnAddrs).Find(&results).Error

	return
}

// GetFromNewOrderSn 通过new_order_sn获取内容 新订单编号
func (obj *_EsAsOrderMgr) GetFromNewOrderSn(newOrderSn string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`new_order_sn` = ?", newOrderSn).Find(&results).Error

	return
}

// GetBatchFromNewOrderSn 批量查找 新订单编号
func (obj *_EsAsOrderMgr) GetBatchFromNewOrderSn(newOrderSns []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`new_order_sn` IN (?)", newOrderSns).Find(&results).Error

	return
}

// GetFromCreateChannel 通过create_channel获取内容 售后服务单创建渠道 NORMAL：正常渠道创建，PINTUAN：拼团失败自动创建
func (obj *_EsAsOrderMgr) GetFromCreateChannel(createChannel string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`create_channel` = ?", createChannel).Find(&results).Error

	return
}

// GetBatchFromCreateChannel 批量查找 售后服务单创建渠道 NORMAL：正常渠道创建，PINTUAN：拼团失败自动创建
func (obj *_EsAsOrderMgr) GetBatchFromCreateChannel(createChannels []string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`create_channel` IN (?)", createChannels).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsAsOrderMgr) FetchByPrimaryKey(id int) (result models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByEsIndexID  获取多个内容
func (obj *_EsAsOrderMgr) FetchIndexByEsIndexID(id int) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// FetchIndexByEsIndexSn  获取多个内容
func (obj *_EsAsOrderMgr) FetchIndexByEsIndexSn(sn string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// FetchIndexByEsIndexOrderSn  获取多个内容
func (obj *_EsAsOrderMgr) FetchIndexByEsIndexOrderSn(orderSn string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// FetchIndexByEsIndexMemberID  获取多个内容
func (obj *_EsAsOrderMgr) FetchIndexByEsIndexMemberID(memberID int) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// FetchIndexByEsIndexSellerID  获取多个内容
func (obj *_EsAsOrderMgr) FetchIndexByEsIndexSellerID(sellerID int) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// FetchIndexByEsIndexServiceType  获取多个内容
func (obj *_EsAsOrderMgr) FetchIndexByEsIndexServiceType(serviceType string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`service_type` = ?", serviceType).Find(&results).Error

	return
}

// FetchIndexByEsIndexServiceStatus  获取多个内容
func (obj *_EsAsOrderMgr) FetchIndexByEsIndexServiceStatus(serviceStatus string) (results []*models.EsAsOrder, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAsOrder{}).Where("`service_status` = ?", serviceStatus).Find(&results).Error

	return
}
