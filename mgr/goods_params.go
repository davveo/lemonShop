package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsGoodsParamsMgr struct {
	*_BaseMgr
}

// EsGoodsParamsMgr open func
func EsGoodsParamsMgr(db *gorm.DB) *_EsGoodsParamsMgr {
	if db == nil {
		panic(fmt.Errorf("EsGoodsParamsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsGoodsParamsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_goods_params"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsGoodsParamsMgr) GetTableName() string {
	return "es_goods_params"
}

// Reset 重置gorm会话
func (obj *_EsGoodsParamsMgr) Reset() *_EsGoodsParamsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsGoodsParamsMgr) Get() (result models.EsGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsParams{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsGoodsParamsMgr) Gets() (results []*models.EsGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsParams{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsGoodsParamsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsGoodsParams{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_EsGoodsParamsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGoodsID goods_id获取 商品id
func (obj *_EsGoodsParamsMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithParamID param_id获取 参数id
func (obj *_EsGoodsParamsMgr) WithParamID(paramID int) Option {
	return optionFunc(func(o *options) { o.query["param_id"] = paramID })
}

// WithParamName param_name获取 参数名字
func (obj *_EsGoodsParamsMgr) WithParamName(paramName string) Option {
	return optionFunc(func(o *options) { o.query["param_name"] = paramName })
}

// WithParamValue param_value获取 参数值
func (obj *_EsGoodsParamsMgr) WithParamValue(paramValue string) Option {
	return optionFunc(func(o *options) { o.query["param_value"] = paramValue })
}

// GetByOption 功能选项模式获取
func (obj *_EsGoodsParamsMgr) GetByOption(opts ...Option) (result models.EsGoodsParams, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsParams{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsGoodsParamsMgr) GetByOptions(opts ...Option) (results []*models.EsGoodsParams, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsParams{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsGoodsParamsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsGoodsParams, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsGoodsParams{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键
func (obj *_EsGoodsParamsMgr) GetFromID(id int) (result models.EsGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsParams{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_EsGoodsParamsMgr) GetBatchFromID(ids []int) (results []*models.EsGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsParams{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品id
func (obj *_EsGoodsParamsMgr) GetFromGoodsID(goodsID int) (results []*models.EsGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsParams{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品id
func (obj *_EsGoodsParamsMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsParams{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromParamID 通过param_id获取内容 参数id
func (obj *_EsGoodsParamsMgr) GetFromParamID(paramID int) (results []*models.EsGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsParams{}).Where("`param_id` = ?", paramID).Find(&results).Error

	return
}

// GetBatchFromParamID 批量查找 参数id
func (obj *_EsGoodsParamsMgr) GetBatchFromParamID(paramIDs []int) (results []*models.EsGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsParams{}).Where("`param_id` IN (?)", paramIDs).Find(&results).Error

	return
}

// GetFromParamName 通过param_name获取内容 参数名字
func (obj *_EsGoodsParamsMgr) GetFromParamName(paramName string) (results []*models.EsGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsParams{}).Where("`param_name` = ?", paramName).Find(&results).Error

	return
}

// GetBatchFromParamName 批量查找 参数名字
func (obj *_EsGoodsParamsMgr) GetBatchFromParamName(paramNames []string) (results []*models.EsGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsParams{}).Where("`param_name` IN (?)", paramNames).Find(&results).Error

	return
}

// GetFromParamValue 通过param_value获取内容 参数值
func (obj *_EsGoodsParamsMgr) GetFromParamValue(paramValue string) (results []*models.EsGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsParams{}).Where("`param_value` = ?", paramValue).Find(&results).Error

	return
}

// GetBatchFromParamValue 批量查找 参数值
func (obj *_EsGoodsParamsMgr) GetBatchFromParamValue(paramValues []string) (results []*models.EsGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsParams{}).Where("`param_value` IN (?)", paramValues).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsGoodsParamsMgr) FetchByPrimaryKey(id int) (result models.EsGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsParams{}).Where("`id` = ?", id).First(&result).Error

	return
}
