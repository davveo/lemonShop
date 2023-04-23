package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsDraftGoodsParamsMgr struct {
	*_BaseMgr
}

// EsDraftGoodsParamsMgr open func
func EsDraftGoodsParamsMgr(db *gorm.DB) *_EsDraftGoodsParamsMgr {
	if db == nil {
		panic(fmt.Errorf("EsDraftGoodsParamsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsDraftGoodsParamsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_draft_goods_params"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsDraftGoodsParamsMgr) GetTableName() string {
	return "es_draft_goods_params"
}

// Reset 重置gorm会话
func (obj *_EsDraftGoodsParamsMgr) Reset() *_EsDraftGoodsParamsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsDraftGoodsParamsMgr) Get() (result models.EsDraftGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsDraftGoodsParamsMgr) Gets() (results []*models.EsDraftGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsDraftGoodsParamsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *_EsDraftGoodsParamsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithDraftGoodsID draft_goods_id获取 草稿ID
func (obj *_EsDraftGoodsParamsMgr) WithDraftGoodsID(draftGoodsID int) Option {
	return optionFunc(func(o *options) { o.query["draft_goods_id"] = draftGoodsID })
}

// WithParamID param_id获取 参数ID
func (obj *_EsDraftGoodsParamsMgr) WithParamID(paramID int) Option {
	return optionFunc(func(o *options) { o.query["param_id"] = paramID })
}

// WithParamName param_name获取 参数名
func (obj *_EsDraftGoodsParamsMgr) WithParamName(paramName string) Option {
	return optionFunc(func(o *options) { o.query["param_name"] = paramName })
}

// WithParamValue param_value获取 参数值
func (obj *_EsDraftGoodsParamsMgr) WithParamValue(paramValue string) Option {
	return optionFunc(func(o *options) { o.query["param_value"] = paramValue })
}

// GetByOption 功能选项模式获取
func (obj *_EsDraftGoodsParamsMgr) GetByOption(opts ...Option) (result models.EsDraftGoodsParams, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsDraftGoodsParamsMgr) GetByOptions(opts ...Option) (results []*models.EsDraftGoodsParams, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsDraftGoodsParamsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsDraftGoodsParams, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where(options.query)
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

// GetFromID 通过id获取内容 ID
func (obj *_EsDraftGoodsParamsMgr) GetFromID(id int) (result models.EsDraftGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *_EsDraftGoodsParamsMgr) GetBatchFromID(ids []int) (results []*models.EsDraftGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromDraftGoodsID 通过draft_goods_id获取内容 草稿ID
func (obj *_EsDraftGoodsParamsMgr) GetFromDraftGoodsID(draftGoodsID int) (results []*models.EsDraftGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`draft_goods_id` = ?", draftGoodsID).Find(&results).Error

	return
}

// GetBatchFromDraftGoodsID 批量查找 草稿ID
func (obj *_EsDraftGoodsParamsMgr) GetBatchFromDraftGoodsID(draftGoodsIDs []int) (results []*models.EsDraftGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`draft_goods_id` IN (?)", draftGoodsIDs).Find(&results).Error

	return
}

// GetFromParamID 通过param_id获取内容 参数ID
func (obj *_EsDraftGoodsParamsMgr) GetFromParamID(paramID int) (results []*models.EsDraftGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`param_id` = ?", paramID).Find(&results).Error

	return
}

// GetBatchFromParamID 批量查找 参数ID
func (obj *_EsDraftGoodsParamsMgr) GetBatchFromParamID(paramIDs []int) (results []*models.EsDraftGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`param_id` IN (?)", paramIDs).Find(&results).Error

	return
}

// GetFromParamName 通过param_name获取内容 参数名
func (obj *_EsDraftGoodsParamsMgr) GetFromParamName(paramName string) (results []*models.EsDraftGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`param_name` = ?", paramName).Find(&results).Error

	return
}

// GetBatchFromParamName 批量查找 参数名
func (obj *_EsDraftGoodsParamsMgr) GetBatchFromParamName(paramNames []string) (results []*models.EsDraftGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`param_name` IN (?)", paramNames).Find(&results).Error

	return
}

// GetFromParamValue 通过param_value获取内容 参数值
func (obj *_EsDraftGoodsParamsMgr) GetFromParamValue(paramValue string) (results []*models.EsDraftGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`param_value` = ?", paramValue).Find(&results).Error

	return
}

// GetBatchFromParamValue 批量查找 参数值
func (obj *_EsDraftGoodsParamsMgr) GetBatchFromParamValue(paramValues []string) (results []*models.EsDraftGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`param_value` IN (?)", paramValues).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsDraftGoodsParamsMgr) FetchByPrimaryKey(id int) (result models.EsDraftGoodsParams, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`id` = ?", id).First(&result).Error

	return
}
