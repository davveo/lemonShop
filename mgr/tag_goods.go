package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsTagGoodsMgr struct {
	*_BaseMgr
}

// EsTagGoodsMgr open func
func EsTagGoodsMgr(db *gorm.DB) *_EsTagGoodsMgr {
	if db == nil {
		panic(fmt.Errorf("EsTagGoodsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsTagGoodsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_tag_goods"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsTagGoodsMgr) GetTableName() string {
	return "es_tag_goods"
}

// Reset 重置gorm会话
func (obj *_EsTagGoodsMgr) Reset() *_EsTagGoodsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsTagGoodsMgr) Get() (result models.EsTagGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsTagGoods{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsTagGoodsMgr) Gets() (results []*models.EsTagGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsTagGoods{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsTagGoodsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsTagGoods{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTagID tag_id获取 标签id
func (obj *_EsTagGoodsMgr) WithTagID(tagID int) Option {
	return optionFunc(func(o *options) { o.query["tag_id"] = tagID })
}

// WithGoodsID goods_id获取 商品id
func (obj *_EsTagGoodsMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// GetByOption 功能选项模式获取
func (obj *_EsTagGoodsMgr) GetByOption(opts ...Option) (result models.EsTagGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsTagGoods{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsTagGoodsMgr) GetByOptions(opts ...Option) (results []*models.EsTagGoods, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsTagGoods{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsTagGoodsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsTagGoods, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsTagGoods{}).Where(options.query)
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

// GetFromTagID 通过tag_id获取内容 标签id
func (obj *_EsTagGoodsMgr) GetFromTagID(tagID int) (results []*models.EsTagGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsTagGoods{}).Where("`tag_id` = ?", tagID).Find(&results).Error

	return
}

// GetBatchFromTagID 批量查找 标签id
func (obj *_EsTagGoodsMgr) GetBatchFromTagID(tagIDs []int) (results []*models.EsTagGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsTagGoods{}).Where("`tag_id` IN (?)", tagIDs).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品id
func (obj *_EsTagGoodsMgr) GetFromGoodsID(goodsID int) (results []*models.EsTagGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsTagGoods{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品id
func (obj *_EsTagGoodsMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsTagGoods, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsTagGoods{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////
