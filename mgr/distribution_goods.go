package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type DistributionGoodsMgr struct {
	*_BaseMgr
}

// NewDistributionGoodsMgr open func
func NewDistributionGoodsMgr(db db.Repo) *DistributionGoodsMgr {
	if db == nil {
		panic(fmt.Errorf("NewDistributionGoodsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &DistributionGoodsMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_category_spec"),
		wdb:       db.GetDbW().Table("es_category_spec"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *DistributionGoodsMgr) GetTableName() string {
	return "es_distribution_goods"
}

// Reset 重置gorm会话
func (obj *DistributionGoodsMgr) Reset() *DistributionGoodsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *DistributionGoodsMgr) Get() (result models.EsDistributionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionGoods{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *DistributionGoodsMgr) Gets() (results []*models.EsDistributionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionGoods{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *DistributionGoodsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionGoods{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *DistributionGoodsMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGoodsID goods_id获取 商品id
func (obj *DistributionGoodsMgr) WithGoodsID(goodsID int64) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithGrade1Rebate grade1_rebate获取 一级返现
func (obj *DistributionGoodsMgr) WithGrade1Rebate(grade1Rebate float64) Option {
	return optionFunc(func(o *options) { o.query["grade1_rebate"] = grade1Rebate })
}

// WithGrade2Rebate grade2_rebate获取 二级返现
func (obj *DistributionGoodsMgr) WithGrade2Rebate(grade2Rebate float64) Option {
	return optionFunc(func(o *options) { o.query["grade2_rebate"] = grade2Rebate })
}

// GetByOption 功能选项模式获取
func (obj *DistributionGoodsMgr) GetByOption(opts ...Option) (result models.EsDistributionGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionGoods{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *DistributionGoodsMgr) GetByOptions(opts ...Option) (results []*models.EsDistributionGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionGoods{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *DistributionGoodsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsDistributionGoods, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionGoods{}).Where(options.query)
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
func (obj *DistributionGoodsMgr) GetFromID(id int64) (result models.EsDistributionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionGoods{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *DistributionGoodsMgr) GetBatchFromID(ids []int64) (results []*models.EsDistributionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionGoods{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品id
func (obj *DistributionGoodsMgr) GetFromGoodsID(goodsID int64) (results []*models.EsDistributionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionGoods{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品id
func (obj *DistributionGoodsMgr) GetBatchFromGoodsID(goodsIDs []int64) (results []*models.EsDistributionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionGoods{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromGrade1Rebate 通过grade1_rebate获取内容 一级返现
func (obj *DistributionGoodsMgr) GetFromGrade1Rebate(grade1Rebate float64) (results []*models.EsDistributionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionGoods{}).Where("`grade1_rebate` = ?", grade1Rebate).Find(&results).Error

	return
}

// GetBatchFromGrade1Rebate 批量查找 一级返现
func (obj *DistributionGoodsMgr) GetBatchFromGrade1Rebate(grade1Rebates []float64) (results []*models.EsDistributionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionGoods{}).Where("`grade1_rebate` IN (?)", grade1Rebates).Find(&results).Error

	return
}

// GetFromGrade2Rebate 通过grade2_rebate获取内容 二级返现
func (obj *DistributionGoodsMgr) GetFromGrade2Rebate(grade2Rebate float64) (results []*models.EsDistributionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionGoods{}).Where("`grade2_rebate` = ?", grade2Rebate).Find(&results).Error

	return
}

// GetBatchFromGrade2Rebate 批量查找 二级返现
func (obj *DistributionGoodsMgr) GetBatchFromGrade2Rebate(grade2Rebates []float64) (results []*models.EsDistributionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionGoods{}).Where("`grade2_rebate` IN (?)", grade2Rebates).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *DistributionGoodsMgr) FetchByPrimaryKey(id int64) (result models.EsDistributionGoods, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDistributionGoods{}).Where("`id` = ?", id).First(&result).Error

	return
}
