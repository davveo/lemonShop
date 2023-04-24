package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type TransactionRecordMgr struct {
	*_BaseMgr
}

// NewTransactionRecordMgr open func
func NewTransactionRecordMgr(db db.Repo) *TransactionRecordMgr {
	if db == nil {
		panic(fmt.Errorf("NewTransactionRecordMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &TransactionRecordMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_transaction_record"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *TransactionRecordMgr) GetTableName() string {
	return "es_transaction_record"
}

// Reset 重置gorm会话
func (obj *TransactionRecordMgr) Reset() *TransactionRecordMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *TransactionRecordMgr) Get() (result models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *TransactionRecordMgr) Gets() (results []*models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *TransactionRecordMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithRecordID record_id获取 主键ID
func (obj *TransactionRecordMgr) WithRecordID(recordID int) Option {
	return optionFunc(func(o *options) { o.query["record_id"] = recordID })
}

// WithOrderSn order_sn获取 订单编号
func (obj *TransactionRecordMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithGoodsID goods_id获取 商品ID
func (obj *TransactionRecordMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithGoodsNum goods_num获取 商品数量
func (obj *TransactionRecordMgr) WithGoodsNum(goodsNum int) Option {
	return optionFunc(func(o *options) { o.query["goods_num"] = goodsNum })
}

// WithRogTime rog_time获取 确认收货时间
func (obj *TransactionRecordMgr) WithRogTime(rogTime int64) Option {
	return optionFunc(func(o *options) { o.query["rog_time"] = rogTime })
}

// WithUname uname获取 用户名
func (obj *TransactionRecordMgr) WithUname(uname string) Option {
	return optionFunc(func(o *options) { o.query["uname"] = uname })
}

// WithPrice price获取 交易价格
func (obj *TransactionRecordMgr) WithPrice(price float64) Option {
	return optionFunc(func(o *options) { o.query["price"] = price })
}

// WithMemberID member_id获取 会员ID
func (obj *TransactionRecordMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// GetByOption 功能选项模式获取
func (obj *TransactionRecordMgr) GetByOption(opts ...Option) (result models.EsTransactionRecord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *TransactionRecordMgr) GetByOptions(opts ...Option) (results []*models.EsTransactionRecord, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *TransactionRecordMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsTransactionRecord, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where(options.query)
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

// GetFromRecordID 通过record_id获取内容 主键ID
func (obj *TransactionRecordMgr) GetFromRecordID(recordID int) (result models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where("`record_id` = ?", recordID).First(&result).Error

	return
}

// GetBatchFromRecordID 批量查找 主键ID
func (obj *TransactionRecordMgr) GetBatchFromRecordID(recordIDs []int) (results []*models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where("`record_id` IN (?)", recordIDs).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *TransactionRecordMgr) GetFromOrderSn(orderSn string) (results []*models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *TransactionRecordMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品ID
func (obj *TransactionRecordMgr) GetFromGoodsID(goodsID int) (results []*models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品ID
func (obj *TransactionRecordMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromGoodsNum 通过goods_num获取内容 商品数量
func (obj *TransactionRecordMgr) GetFromGoodsNum(goodsNum int) (results []*models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where("`goods_num` = ?", goodsNum).Find(&results).Error

	return
}

// GetBatchFromGoodsNum 批量查找 商品数量
func (obj *TransactionRecordMgr) GetBatchFromGoodsNum(goodsNums []int) (results []*models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where("`goods_num` IN (?)", goodsNums).Find(&results).Error

	return
}

// GetFromRogTime 通过rog_time获取内容 确认收货时间
func (obj *TransactionRecordMgr) GetFromRogTime(rogTime int64) (results []*models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where("`rog_time` = ?", rogTime).Find(&results).Error

	return
}

// GetBatchFromRogTime 批量查找 确认收货时间
func (obj *TransactionRecordMgr) GetBatchFromRogTime(rogTimes []int64) (results []*models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where("`rog_time` IN (?)", rogTimes).Find(&results).Error

	return
}

// GetFromUname 通过uname获取内容 用户名
func (obj *TransactionRecordMgr) GetFromUname(uname string) (results []*models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where("`uname` = ?", uname).Find(&results).Error

	return
}

// GetBatchFromUname 批量查找 用户名
func (obj *TransactionRecordMgr) GetBatchFromUname(unames []string) (results []*models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where("`uname` IN (?)", unames).Find(&results).Error

	return
}

// GetFromPrice 通过price获取内容 交易价格
func (obj *TransactionRecordMgr) GetFromPrice(price float64) (results []*models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where("`price` = ?", price).Find(&results).Error

	return
}

// GetBatchFromPrice 批量查找 交易价格
func (obj *TransactionRecordMgr) GetBatchFromPrice(prices []float64) (results []*models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where("`price` IN (?)", prices).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员ID
func (obj *TransactionRecordMgr) GetFromMemberID(memberID int) (results []*models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员ID
func (obj *TransactionRecordMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *TransactionRecordMgr) FetchByPrimaryKey(recordID int) (result models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where("`record_id` = ?", recordID).First(&result).Error

	return
}

// FetchIndexByIndexGoodsID  获取多个内容
func (obj *TransactionRecordMgr) FetchIndexByIndexGoodsID(goodsID int) (results []*models.EsTransactionRecord, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTransactionRecord{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}
