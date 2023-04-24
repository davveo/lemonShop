package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type SellerBillMgr struct {
	*_BaseMgr
}

// NewSellerBillMgr open func
func NewSellerBillMgr(db db.Repo) *SellerBillMgr {
	if db == nil {
		panic(fmt.Errorf("NewSellerBillMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &SellerBillMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_seller_bill"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *SellerBillMgr) GetTableName() string {
	return "es_seller_bill"
}

// Reset 重置gorm会话
func (obj *SellerBillMgr) Reset() *SellerBillMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *SellerBillMgr) Get() (result models.EsSellerBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *SellerBillMgr) Gets() (results []*models.EsSellerBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *SellerBillMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *SellerBillMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSellerID seller_id获取 商家id
func (obj *SellerBillMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithCreateTime create_time获取 创建时间
func (obj *SellerBillMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithOrderSn order_sn获取 订单编号
func (obj *SellerBillMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithExpenditure expenditure获取 商品反现支出
func (obj *SellerBillMgr) WithExpenditure(expenditure float64) Option {
	return optionFunc(func(o *options) { o.query["expenditure"] = expenditure })
}

// WithReturnExpenditure return_expenditure获取 返还商品反现支出
func (obj *SellerBillMgr) WithReturnExpenditure(returnExpenditure float64) Option {
	return optionFunc(func(o *options) { o.query["return_expenditure"] = returnExpenditure })
}

// GetByOption 功能选项模式获取
func (obj *SellerBillMgr) GetByOption(opts ...Option) (result models.EsSellerBill, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *SellerBillMgr) GetByOptions(opts ...Option) (results []*models.EsSellerBill, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *SellerBillMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSellerBill, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where(options.query)
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
func (obj *SellerBillMgr) GetFromID(id int) (result models.EsSellerBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *SellerBillMgr) GetBatchFromID(ids []int) (results []*models.EsSellerBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 商家id
func (obj *SellerBillMgr) GetFromSellerID(sellerID int) (results []*models.EsSellerBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 商家id
func (obj *SellerBillMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsSellerBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *SellerBillMgr) GetFromCreateTime(createTime int64) (results []*models.EsSellerBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *SellerBillMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsSellerBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *SellerBillMgr) GetFromOrderSn(orderSn string) (results []*models.EsSellerBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *SellerBillMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsSellerBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromExpenditure 通过expenditure获取内容 商品反现支出
func (obj *SellerBillMgr) GetFromExpenditure(expenditure float64) (results []*models.EsSellerBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`expenditure` = ?", expenditure).Find(&results).Error

	return
}

// GetBatchFromExpenditure 批量查找 商品反现支出
func (obj *SellerBillMgr) GetBatchFromExpenditure(expenditures []float64) (results []*models.EsSellerBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`expenditure` IN (?)", expenditures).Find(&results).Error

	return
}

// GetFromReturnExpenditure 通过return_expenditure获取内容 返还商品反现支出
func (obj *SellerBillMgr) GetFromReturnExpenditure(returnExpenditure float64) (results []*models.EsSellerBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`return_expenditure` = ?", returnExpenditure).Find(&results).Error

	return
}

// GetBatchFromReturnExpenditure 批量查找 返还商品反现支出
func (obj *SellerBillMgr) GetBatchFromReturnExpenditure(returnExpenditures []float64) (results []*models.EsSellerBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`return_expenditure` IN (?)", returnExpenditures).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *SellerBillMgr) FetchByPrimaryKey(id int) (result models.EsSellerBill, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSellerBill{}).Where("`id` = ?", id).First(&result).Error

	return
}
