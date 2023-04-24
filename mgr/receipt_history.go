package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type ReceiptHistoryMgr struct {
	*_BaseMgr
}

// NewReceiptHistoryMgr open func
func NewReceiptHistoryMgr(db db.Repo) *ReceiptHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("NewReceiptHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &ReceiptHistoryMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_receipt_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *ReceiptHistoryMgr) GetTableName() string {
	return "es_receipt_history"
}

// Reset 重置gorm会话
func (obj *ReceiptHistoryMgr) Reset() *ReceiptHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *ReceiptHistoryMgr) Get() (result models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *ReceiptHistoryMgr) Gets() (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *ReceiptHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithHistoryID history_id获取 主键ID
func (obj *ReceiptHistoryMgr) WithHistoryID(historyID int) Option {
	return optionFunc(func(o *options) { o.query["history_id"] = historyID })
}

// WithOrderSn order_sn获取 订单编号
func (obj *ReceiptHistoryMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithOrderPrice order_price获取 订单金额
func (obj *ReceiptHistoryMgr) WithOrderPrice(orderPrice float64) Option {
	return optionFunc(func(o *options) { o.query["order_price"] = orderPrice })
}

// WithSellerID seller_id获取 开票商家ID
func (obj *ReceiptHistoryMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithSellerName seller_name获取 开票商家
func (obj *ReceiptHistoryMgr) WithSellerName(sellerName string) Option {
	return optionFunc(func(o *options) { o.query["seller_name"] = sellerName })
}

// WithMemberID member_id获取 会员ID
func (obj *ReceiptHistoryMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithStatus status获取 开票状态 0：未开，1：已开
func (obj *ReceiptHistoryMgr) WithStatus(status int16) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithReceiptMethod receipt_method获取 开票方式(针对增值税专用发票)
func (obj *ReceiptHistoryMgr) WithReceiptMethod(receiptMethod string) Option {
	return optionFunc(func(o *options) { o.query["receipt_method"] = receiptMethod })
}

// WithReceiptType receipt_type获取 发票类型 ELECTRO：电子普通发票，VATORDINARY：增值税普通发票，VATOSPECIAL：增值税专用发票
func (obj *ReceiptHistoryMgr) WithReceiptType(receiptType string) Option {
	return optionFunc(func(o *options) { o.query["receipt_type"] = receiptType })
}

// WithLogiID logi_id获取 物流公司ID
func (obj *ReceiptHistoryMgr) WithLogiID(logiID int) Option {
	return optionFunc(func(o *options) { o.query["logi_id"] = logiID })
}

// WithLogiName logi_name获取 物流公司名称
func (obj *ReceiptHistoryMgr) WithLogiName(logiName string) Option {
	return optionFunc(func(o *options) { o.query["logi_name"] = logiName })
}

// WithLogiCode logi_code获取 快递单号
func (obj *ReceiptHistoryMgr) WithLogiCode(logiCode string) Option {
	return optionFunc(func(o *options) { o.query["logi_code"] = logiCode })
}

// WithReceiptTitle receipt_title获取 发票抬头
func (obj *ReceiptHistoryMgr) WithReceiptTitle(receiptTitle string) Option {
	return optionFunc(func(o *options) { o.query["receipt_title"] = receiptTitle })
}

// WithReceiptContent receipt_content获取 发票内容
func (obj *ReceiptHistoryMgr) WithReceiptContent(receiptContent string) Option {
	return optionFunc(func(o *options) { o.query["receipt_content"] = receiptContent })
}

// WithTaxNo tax_no获取 纳税人识别号
func (obj *ReceiptHistoryMgr) WithTaxNo(taxNo string) Option {
	return optionFunc(func(o *options) { o.query["tax_no"] = taxNo })
}

// WithRegAddr reg_addr获取 注册地址
func (obj *ReceiptHistoryMgr) WithRegAddr(regAddr string) Option {
	return optionFunc(func(o *options) { o.query["reg_addr"] = regAddr })
}

// WithRegTel reg_tel获取 注册电话
func (obj *ReceiptHistoryMgr) WithRegTel(regTel string) Option {
	return optionFunc(func(o *options) { o.query["reg_tel"] = regTel })
}

// WithBankName bank_name获取 开户银行
func (obj *ReceiptHistoryMgr) WithBankName(bankName string) Option {
	return optionFunc(func(o *options) { o.query["bank_name"] = bankName })
}

// WithBankAccount bank_account获取 银行账户
func (obj *ReceiptHistoryMgr) WithBankAccount(bankAccount string) Option {
	return optionFunc(func(o *options) { o.query["bank_account"] = bankAccount })
}

// WithMemberName member_name获取 收票人名称
func (obj *ReceiptHistoryMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithMemberMobile member_mobile获取 收票人手机号
func (obj *ReceiptHistoryMgr) WithMemberMobile(memberMobile string) Option {
	return optionFunc(func(o *options) { o.query["member_mobile"] = memberMobile })
}

// WithMemberEmail member_email获取 收票人邮箱
func (obj *ReceiptHistoryMgr) WithMemberEmail(memberEmail string) Option {
	return optionFunc(func(o *options) { o.query["member_email"] = memberEmail })
}

// WithProvinceID province_id获取 收票地址-所属省份ID
func (obj *ReceiptHistoryMgr) WithProvinceID(provinceID int) Option {
	return optionFunc(func(o *options) { o.query["province_id"] = provinceID })
}

// WithCityID city_id获取 收票地址-所属城市ID
func (obj *ReceiptHistoryMgr) WithCityID(cityID int) Option {
	return optionFunc(func(o *options) { o.query["city_id"] = cityID })
}

// WithCountyID county_id获取 收票地址-所属区县ID
func (obj *ReceiptHistoryMgr) WithCountyID(countyID int) Option {
	return optionFunc(func(o *options) { o.query["county_id"] = countyID })
}

// WithTownID town_id获取 收票地址-所属乡镇ID
func (obj *ReceiptHistoryMgr) WithTownID(townID int) Option {
	return optionFunc(func(o *options) { o.query["town_id"] = townID })
}

// WithProvince province获取 收票地址-所属省份
func (obj *ReceiptHistoryMgr) WithProvince(province string) Option {
	return optionFunc(func(o *options) { o.query["province"] = province })
}

// WithCity city获取 收票地址-所属城市
func (obj *ReceiptHistoryMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithCounty county获取 收票地址-所属区县
func (obj *ReceiptHistoryMgr) WithCounty(county string) Option {
	return optionFunc(func(o *options) { o.query["county"] = county })
}

// WithTown town获取 收票地址-所属乡镇
func (obj *ReceiptHistoryMgr) WithTown(town string) Option {
	return optionFunc(func(o *options) { o.query["town"] = town })
}

// WithDetailAddr detail_addr获取 收票地址-详细地址
func (obj *ReceiptHistoryMgr) WithDetailAddr(detailAddr string) Option {
	return optionFunc(func(o *options) { o.query["detail_addr"] = detailAddr })
}

// WithAddTime add_time获取 申请开票日期
func (obj *ReceiptHistoryMgr) WithAddTime(addTime int64) Option {
	return optionFunc(func(o *options) { o.query["add_time"] = addTime })
}

// WithGoodsJSON goods_json获取 订单商品信息
func (obj *ReceiptHistoryMgr) WithGoodsJSON(goodsJSON string) Option {
	return optionFunc(func(o *options) { o.query["goods_json"] = goodsJSON })
}

// WithOrderStatus order_status获取 订单出库状态，NEW/CONFIRM
func (obj *ReceiptHistoryMgr) WithOrderStatus(orderStatus string) Option {
	return optionFunc(func(o *options) { o.query["order_status"] = orderStatus })
}

// WithUname uname获取 会员名称
func (obj *ReceiptHistoryMgr) WithUname(uname string) Option {
	return optionFunc(func(o *options) { o.query["uname"] = uname })
}

// GetByOption 功能选项模式获取
func (obj *ReceiptHistoryMgr) GetByOption(opts ...Option) (result models.EsReceiptHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *ReceiptHistoryMgr) GetByOptions(opts ...Option) (results []*models.EsReceiptHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *ReceiptHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsReceiptHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where(options.query)
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

// GetFromHistoryID 通过history_id获取内容 主键ID
func (obj *ReceiptHistoryMgr) GetFromHistoryID(historyID int) (result models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`history_id` = ?", historyID).First(&result).Error

	return
}

// GetBatchFromHistoryID 批量查找 主键ID
func (obj *ReceiptHistoryMgr) GetBatchFromHistoryID(historyIDs []int) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`history_id` IN (?)", historyIDs).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *ReceiptHistoryMgr) GetFromOrderSn(orderSn string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *ReceiptHistoryMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromOrderPrice 通过order_price获取内容 订单金额
func (obj *ReceiptHistoryMgr) GetFromOrderPrice(orderPrice float64) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`order_price` = ?", orderPrice).Find(&results).Error

	return
}

// GetBatchFromOrderPrice 批量查找 订单金额
func (obj *ReceiptHistoryMgr) GetBatchFromOrderPrice(orderPrices []float64) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`order_price` IN (?)", orderPrices).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 开票商家ID
func (obj *ReceiptHistoryMgr) GetFromSellerID(sellerID int) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 开票商家ID
func (obj *ReceiptHistoryMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromSellerName 通过seller_name获取内容 开票商家
func (obj *ReceiptHistoryMgr) GetFromSellerName(sellerName string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`seller_name` = ?", sellerName).Find(&results).Error

	return
}

// GetBatchFromSellerName 批量查找 开票商家
func (obj *ReceiptHistoryMgr) GetBatchFromSellerName(sellerNames []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`seller_name` IN (?)", sellerNames).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员ID
func (obj *ReceiptHistoryMgr) GetFromMemberID(memberID int) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员ID
func (obj *ReceiptHistoryMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 开票状态 0：未开，1：已开
func (obj *ReceiptHistoryMgr) GetFromStatus(status int16) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 开票状态 0：未开，1：已开
func (obj *ReceiptHistoryMgr) GetBatchFromStatus(statuss []int16) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromReceiptMethod 通过receipt_method获取内容 开票方式(针对增值税专用发票)
func (obj *ReceiptHistoryMgr) GetFromReceiptMethod(receiptMethod string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`receipt_method` = ?", receiptMethod).Find(&results).Error

	return
}

// GetBatchFromReceiptMethod 批量查找 开票方式(针对增值税专用发票)
func (obj *ReceiptHistoryMgr) GetBatchFromReceiptMethod(receiptMethods []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`receipt_method` IN (?)", receiptMethods).Find(&results).Error

	return
}

// GetFromReceiptType 通过receipt_type获取内容 发票类型 ELECTRO：电子普通发票，VATORDINARY：增值税普通发票，VATOSPECIAL：增值税专用发票
func (obj *ReceiptHistoryMgr) GetFromReceiptType(receiptType string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`receipt_type` = ?", receiptType).Find(&results).Error

	return
}

// GetBatchFromReceiptType 批量查找 发票类型 ELECTRO：电子普通发票，VATORDINARY：增值税普通发票，VATOSPECIAL：增值税专用发票
func (obj *ReceiptHistoryMgr) GetBatchFromReceiptType(receiptTypes []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`receipt_type` IN (?)", receiptTypes).Find(&results).Error

	return
}

// GetFromLogiID 通过logi_id获取内容 物流公司ID
func (obj *ReceiptHistoryMgr) GetFromLogiID(logiID int) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`logi_id` = ?", logiID).Find(&results).Error

	return
}

// GetBatchFromLogiID 批量查找 物流公司ID
func (obj *ReceiptHistoryMgr) GetBatchFromLogiID(logiIDs []int) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`logi_id` IN (?)", logiIDs).Find(&results).Error

	return
}

// GetFromLogiName 通过logi_name获取内容 物流公司名称
func (obj *ReceiptHistoryMgr) GetFromLogiName(logiName string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`logi_name` = ?", logiName).Find(&results).Error

	return
}

// GetBatchFromLogiName 批量查找 物流公司名称
func (obj *ReceiptHistoryMgr) GetBatchFromLogiName(logiNames []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`logi_name` IN (?)", logiNames).Find(&results).Error

	return
}

// GetFromLogiCode 通过logi_code获取内容 快递单号
func (obj *ReceiptHistoryMgr) GetFromLogiCode(logiCode string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`logi_code` = ?", logiCode).Find(&results).Error

	return
}

// GetBatchFromLogiCode 批量查找 快递单号
func (obj *ReceiptHistoryMgr) GetBatchFromLogiCode(logiCodes []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`logi_code` IN (?)", logiCodes).Find(&results).Error

	return
}

// GetFromReceiptTitle 通过receipt_title获取内容 发票抬头
func (obj *ReceiptHistoryMgr) GetFromReceiptTitle(receiptTitle string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`receipt_title` = ?", receiptTitle).Find(&results).Error

	return
}

// GetBatchFromReceiptTitle 批量查找 发票抬头
func (obj *ReceiptHistoryMgr) GetBatchFromReceiptTitle(receiptTitles []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`receipt_title` IN (?)", receiptTitles).Find(&results).Error

	return
}

// GetFromReceiptContent 通过receipt_content获取内容 发票内容
func (obj *ReceiptHistoryMgr) GetFromReceiptContent(receiptContent string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`receipt_content` = ?", receiptContent).Find(&results).Error

	return
}

// GetBatchFromReceiptContent 批量查找 发票内容
func (obj *ReceiptHistoryMgr) GetBatchFromReceiptContent(receiptContents []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`receipt_content` IN (?)", receiptContents).Find(&results).Error

	return
}

// GetFromTaxNo 通过tax_no获取内容 纳税人识别号
func (obj *ReceiptHistoryMgr) GetFromTaxNo(taxNo string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`tax_no` = ?", taxNo).Find(&results).Error

	return
}

// GetBatchFromTaxNo 批量查找 纳税人识别号
func (obj *ReceiptHistoryMgr) GetBatchFromTaxNo(taxNos []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`tax_no` IN (?)", taxNos).Find(&results).Error

	return
}

// GetFromRegAddr 通过reg_addr获取内容 注册地址
func (obj *ReceiptHistoryMgr) GetFromRegAddr(regAddr string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`reg_addr` = ?", regAddr).Find(&results).Error

	return
}

// GetBatchFromRegAddr 批量查找 注册地址
func (obj *ReceiptHistoryMgr) GetBatchFromRegAddr(regAddrs []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`reg_addr` IN (?)", regAddrs).Find(&results).Error

	return
}

// GetFromRegTel 通过reg_tel获取内容 注册电话
func (obj *ReceiptHistoryMgr) GetFromRegTel(regTel string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`reg_tel` = ?", regTel).Find(&results).Error

	return
}

// GetBatchFromRegTel 批量查找 注册电话
func (obj *ReceiptHistoryMgr) GetBatchFromRegTel(regTels []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`reg_tel` IN (?)", regTels).Find(&results).Error

	return
}

// GetFromBankName 通过bank_name获取内容 开户银行
func (obj *ReceiptHistoryMgr) GetFromBankName(bankName string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`bank_name` = ?", bankName).Find(&results).Error

	return
}

// GetBatchFromBankName 批量查找 开户银行
func (obj *ReceiptHistoryMgr) GetBatchFromBankName(bankNames []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`bank_name` IN (?)", bankNames).Find(&results).Error

	return
}

// GetFromBankAccount 通过bank_account获取内容 银行账户
func (obj *ReceiptHistoryMgr) GetFromBankAccount(bankAccount string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`bank_account` = ?", bankAccount).Find(&results).Error

	return
}

// GetBatchFromBankAccount 批量查找 银行账户
func (obj *ReceiptHistoryMgr) GetBatchFromBankAccount(bankAccounts []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`bank_account` IN (?)", bankAccounts).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 收票人名称
func (obj *ReceiptHistoryMgr) GetFromMemberName(memberName string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 收票人名称
func (obj *ReceiptHistoryMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromMemberMobile 通过member_mobile获取内容 收票人手机号
func (obj *ReceiptHistoryMgr) GetFromMemberMobile(memberMobile string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`member_mobile` = ?", memberMobile).Find(&results).Error

	return
}

// GetBatchFromMemberMobile 批量查找 收票人手机号
func (obj *ReceiptHistoryMgr) GetBatchFromMemberMobile(memberMobiles []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`member_mobile` IN (?)", memberMobiles).Find(&results).Error

	return
}

// GetFromMemberEmail 通过member_email获取内容 收票人邮箱
func (obj *ReceiptHistoryMgr) GetFromMemberEmail(memberEmail string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`member_email` = ?", memberEmail).Find(&results).Error

	return
}

// GetBatchFromMemberEmail 批量查找 收票人邮箱
func (obj *ReceiptHistoryMgr) GetBatchFromMemberEmail(memberEmails []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`member_email` IN (?)", memberEmails).Find(&results).Error

	return
}

// GetFromProvinceID 通过province_id获取内容 收票地址-所属省份ID
func (obj *ReceiptHistoryMgr) GetFromProvinceID(provinceID int) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`province_id` = ?", provinceID).Find(&results).Error

	return
}

// GetBatchFromProvinceID 批量查找 收票地址-所属省份ID
func (obj *ReceiptHistoryMgr) GetBatchFromProvinceID(provinceIDs []int) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`province_id` IN (?)", provinceIDs).Find(&results).Error

	return
}

// GetFromCityID 通过city_id获取内容 收票地址-所属城市ID
func (obj *ReceiptHistoryMgr) GetFromCityID(cityID int) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`city_id` = ?", cityID).Find(&results).Error

	return
}

// GetBatchFromCityID 批量查找 收票地址-所属城市ID
func (obj *ReceiptHistoryMgr) GetBatchFromCityID(cityIDs []int) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`city_id` IN (?)", cityIDs).Find(&results).Error

	return
}

// GetFromCountyID 通过county_id获取内容 收票地址-所属区县ID
func (obj *ReceiptHistoryMgr) GetFromCountyID(countyID int) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`county_id` = ?", countyID).Find(&results).Error

	return
}

// GetBatchFromCountyID 批量查找 收票地址-所属区县ID
func (obj *ReceiptHistoryMgr) GetBatchFromCountyID(countyIDs []int) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`county_id` IN (?)", countyIDs).Find(&results).Error

	return
}

// GetFromTownID 通过town_id获取内容 收票地址-所属乡镇ID
func (obj *ReceiptHistoryMgr) GetFromTownID(townID int) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`town_id` = ?", townID).Find(&results).Error

	return
}

// GetBatchFromTownID 批量查找 收票地址-所属乡镇ID
func (obj *ReceiptHistoryMgr) GetBatchFromTownID(townIDs []int) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`town_id` IN (?)", townIDs).Find(&results).Error

	return
}

// GetFromProvince 通过province获取内容 收票地址-所属省份
func (obj *ReceiptHistoryMgr) GetFromProvince(province string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`province` = ?", province).Find(&results).Error

	return
}

// GetBatchFromProvince 批量查找 收票地址-所属省份
func (obj *ReceiptHistoryMgr) GetBatchFromProvince(provinces []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`province` IN (?)", provinces).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容 收票地址-所属城市
func (obj *ReceiptHistoryMgr) GetFromCity(city string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`city` = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量查找 收票地址-所属城市
func (obj *ReceiptHistoryMgr) GetBatchFromCity(citys []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`city` IN (?)", citys).Find(&results).Error

	return
}

// GetFromCounty 通过county获取内容 收票地址-所属区县
func (obj *ReceiptHistoryMgr) GetFromCounty(county string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`county` = ?", county).Find(&results).Error

	return
}

// GetBatchFromCounty 批量查找 收票地址-所属区县
func (obj *ReceiptHistoryMgr) GetBatchFromCounty(countys []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`county` IN (?)", countys).Find(&results).Error

	return
}

// GetFromTown 通过town获取内容 收票地址-所属乡镇
func (obj *ReceiptHistoryMgr) GetFromTown(town string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`town` = ?", town).Find(&results).Error

	return
}

// GetBatchFromTown 批量查找 收票地址-所属乡镇
func (obj *ReceiptHistoryMgr) GetBatchFromTown(towns []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`town` IN (?)", towns).Find(&results).Error

	return
}

// GetFromDetailAddr 通过detail_addr获取内容 收票地址-详细地址
func (obj *ReceiptHistoryMgr) GetFromDetailAddr(detailAddr string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`detail_addr` = ?", detailAddr).Find(&results).Error

	return
}

// GetBatchFromDetailAddr 批量查找 收票地址-详细地址
func (obj *ReceiptHistoryMgr) GetBatchFromDetailAddr(detailAddrs []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`detail_addr` IN (?)", detailAddrs).Find(&results).Error

	return
}

// GetFromAddTime 通过add_time获取内容 申请开票日期
func (obj *ReceiptHistoryMgr) GetFromAddTime(addTime int64) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`add_time` = ?", addTime).Find(&results).Error

	return
}

// GetBatchFromAddTime 批量查找 申请开票日期
func (obj *ReceiptHistoryMgr) GetBatchFromAddTime(addTimes []int64) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`add_time` IN (?)", addTimes).Find(&results).Error

	return
}

// GetFromGoodsJSON 通过goods_json获取内容 订单商品信息
func (obj *ReceiptHistoryMgr) GetFromGoodsJSON(goodsJSON string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`goods_json` = ?", goodsJSON).Find(&results).Error

	return
}

// GetBatchFromGoodsJSON 批量查找 订单商品信息
func (obj *ReceiptHistoryMgr) GetBatchFromGoodsJSON(goodsJSONs []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`goods_json` IN (?)", goodsJSONs).Find(&results).Error

	return
}

// GetFromOrderStatus 通过order_status获取内容 订单出库状态，NEW/CONFIRM
func (obj *ReceiptHistoryMgr) GetFromOrderStatus(orderStatus string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`order_status` = ?", orderStatus).Find(&results).Error

	return
}

// GetBatchFromOrderStatus 批量查找 订单出库状态，NEW/CONFIRM
func (obj *ReceiptHistoryMgr) GetBatchFromOrderStatus(orderStatuss []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`order_status` IN (?)", orderStatuss).Find(&results).Error

	return
}

// GetFromUname 通过uname获取内容 会员名称
func (obj *ReceiptHistoryMgr) GetFromUname(uname string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`uname` = ?", uname).Find(&results).Error

	return
}

// GetBatchFromUname 批量查找 会员名称
func (obj *ReceiptHistoryMgr) GetBatchFromUname(unames []string) (results []*models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`uname` IN (?)", unames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *ReceiptHistoryMgr) FetchByPrimaryKey(historyID int) (result models.EsReceiptHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsReceiptHistory{}).Where("`history_id` = ?", historyID).First(&result).Error

	return
}
