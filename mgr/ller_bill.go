package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsSellerBillMgr struct {
	*_BaseMgr
}

// EsSellerBillMgr open func
func EsSellerBillMgr(db *gorm.DB) *_EsSellerBillMgr {
	if db == nil {
		panic(fmt.Errorf("EsSellerBillMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsSellerBillMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_seller_bill"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsSellerBillMgr) GetTableName() string {
	return "es_seller_bill"
}

// Reset 重置gorm会话
func (obj *_EsSellerBillMgr) Reset() *_EsSellerBillMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsSellerBillMgr) Get() (result models.EsSellerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsSellerBillMgr) Gets() (results []*models.EsSellerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsSellerBillMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_EsSellerBillMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSellerID seller_id获取 商家id
func (obj *_EsSellerBillMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithCreateTime create_time获取 创建时间
func (obj *_EsSellerBillMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithOrderSn order_sn获取 订单编号
func (obj *_EsSellerBillMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithExpenditure expenditure获取 商品反现支出
func (obj *_EsSellerBillMgr) WithExpenditure(expenditure float64) Option {
	return optionFunc(func(o *options) { o.query["expenditure"] = expenditure })
}

// WithReturnExpenditure return_expenditure获取 返还商品反现支出
func (obj *_EsSellerBillMgr) WithReturnExpenditure(returnExpenditure float64) Option {
	return optionFunc(func(o *options) { o.query["return_expenditure"] = returnExpenditure })
}

// GetByOption 功能选项模式获取
func (obj *_EsSellerBillMgr) GetByOption(opts ...Option) (result models.EsSellerBill, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsSellerBillMgr) GetByOptions(opts ...Option) (results []*models.EsSellerBill, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsSellerBillMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSellerBill, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where(options.query)
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
func (obj *_EsSellerBillMgr) GetFromID(id int) (result models.EsSellerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_EsSellerBillMgr) GetBatchFromID(ids []int) (results []*models.EsSellerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 商家id
func (obj *_EsSellerBillMgr) GetFromSellerID(sellerID int) (results []*models.EsSellerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 商家id
func (obj *_EsSellerBillMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsSellerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_EsSellerBillMgr) GetFromCreateTime(createTime int64) (results []*models.EsSellerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_EsSellerBillMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsSellerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *_EsSellerBillMgr) GetFromOrderSn(orderSn string) (results []*models.EsSellerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *_EsSellerBillMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsSellerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromExpenditure 通过expenditure获取内容 商品反现支出
func (obj *_EsSellerBillMgr) GetFromExpenditure(expenditure float64) (results []*models.EsSellerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`expenditure` = ?", expenditure).Find(&results).Error

	return
}

// GetBatchFromExpenditure 批量查找 商品反现支出
func (obj *_EsSellerBillMgr) GetBatchFromExpenditure(expenditures []float64) (results []*models.EsSellerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`expenditure` IN (?)", expenditures).Find(&results).Error

	return
}

// GetFromReturnExpenditure 通过return_expenditure获取内容 返还商品反现支出
func (obj *_EsSellerBillMgr) GetFromReturnExpenditure(returnExpenditure float64) (results []*models.EsSellerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`return_expenditure` = ?", returnExpenditure).Find(&results).Error

	return
}

// GetBatchFromReturnExpenditure 批量查找 返还商品反现支出
func (obj *_EsSellerBillMgr) GetBatchFromReturnExpenditure(returnExpenditures []float64) (results []*models.EsSellerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`return_expenditure` IN (?)", returnExpenditures).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsSellerBillMgr) FetchByPrimaryKey(id int) (result models.EsSellerBill, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`id` = ?", id).First(&result).Error

	return
}
