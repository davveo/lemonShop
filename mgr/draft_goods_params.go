package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type DraftGoodsParamsMgr struct {
	*_BaseMgr
}

// NewDraftGoodsParamsMgr open func
func NewDraftGoodsParamsMgr(db db.Repo) *DraftGoodsParamsMgr {
	if db == nil {
		panic(fmt.Errorf("NewDraftGoodsParamsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &DraftGoodsParamsMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_draft_goods_params"),
		wdb:       db.GetDbW().Table("es_draft_goods_params"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *DraftGoodsParamsMgr) GetTableName() string {
	return "es_draft_goods_params"
}

// Reset 重置gorm会话
func (obj *DraftGoodsParamsMgr) Reset() *DraftGoodsParamsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *DraftGoodsParamsMgr) Get() (result models.EsDraftGoodsParams, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *DraftGoodsParamsMgr) Gets() (results []*models.EsDraftGoodsParams, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *DraftGoodsParamsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 ID
func (obj *DraftGoodsParamsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithDraftGoodsID draft_goods_id获取 草稿ID
func (obj *DraftGoodsParamsMgr) WithDraftGoodsID(draftGoodsID int) Option {
	return optionFunc(func(o *options) { o.query["draft_goods_id"] = draftGoodsID })
}

// WithParamID param_id获取 参数ID
func (obj *DraftGoodsParamsMgr) WithParamID(paramID int) Option {
	return optionFunc(func(o *options) { o.query["param_id"] = paramID })
}

// WithParamName param_name获取 参数名
func (obj *DraftGoodsParamsMgr) WithParamName(paramName string) Option {
	return optionFunc(func(o *options) { o.query["param_name"] = paramName })
}

// WithParamValue param_value获取 参数值
func (obj *DraftGoodsParamsMgr) WithParamValue(paramValue string) Option {
	return optionFunc(func(o *options) { o.query["param_value"] = paramValue })
}

// GetByOption 功能选项模式获取
func (obj *DraftGoodsParamsMgr) GetByOption(opts ...Option) (result models.EsDraftGoodsParams, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *DraftGoodsParamsMgr) GetByOptions(opts ...Option) (results []*models.EsDraftGoodsParams, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *DraftGoodsParamsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsDraftGoodsParams, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where(options.query)
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
func (obj *DraftGoodsParamsMgr) GetFromID(id int) (result models.EsDraftGoodsParams, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 ID
func (obj *DraftGoodsParamsMgr) GetBatchFromID(ids []int) (results []*models.EsDraftGoodsParams, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromDraftGoodsID 通过draft_goods_id获取内容 草稿ID
func (obj *DraftGoodsParamsMgr) GetFromDraftGoodsID(draftGoodsID int) (results []*models.EsDraftGoodsParams, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`draft_goods_id` = ?", draftGoodsID).Find(&results).Error

	return
}

// GetBatchFromDraftGoodsID 批量查找 草稿ID
func (obj *DraftGoodsParamsMgr) GetBatchFromDraftGoodsID(draftGoodsIDs []int) (results []*models.EsDraftGoodsParams, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`draft_goods_id` IN (?)", draftGoodsIDs).Find(&results).Error

	return
}

// GetFromParamID 通过param_id获取内容 参数ID
func (obj *DraftGoodsParamsMgr) GetFromParamID(paramID int) (results []*models.EsDraftGoodsParams, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`param_id` = ?", paramID).Find(&results).Error

	return
}

// GetBatchFromParamID 批量查找 参数ID
func (obj *DraftGoodsParamsMgr) GetBatchFromParamID(paramIDs []int) (results []*models.EsDraftGoodsParams, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`param_id` IN (?)", paramIDs).Find(&results).Error

	return
}

// GetFromParamName 通过param_name获取内容 参数名
func (obj *DraftGoodsParamsMgr) GetFromParamName(paramName string) (results []*models.EsDraftGoodsParams, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`param_name` = ?", paramName).Find(&results).Error

	return
}

// GetBatchFromParamName 批量查找 参数名
func (obj *DraftGoodsParamsMgr) GetBatchFromParamName(paramNames []string) (results []*models.EsDraftGoodsParams, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`param_name` IN (?)", paramNames).Find(&results).Error

	return
}

// GetFromParamValue 通过param_value获取内容 参数值
func (obj *DraftGoodsParamsMgr) GetFromParamValue(paramValue string) (results []*models.EsDraftGoodsParams, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`param_value` = ?", paramValue).Find(&results).Error

	return
}

// GetBatchFromParamValue 批量查找 参数值
func (obj *DraftGoodsParamsMgr) GetBatchFromParamValue(paramValues []string) (results []*models.EsDraftGoodsParams, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`param_value` IN (?)", paramValues).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *DraftGoodsParamsMgr) FetchByPrimaryKey(id int) (result models.EsDraftGoodsParams, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsDraftGoodsParams{}).Where("`id` = ?", id).First(&result).Error

	return
}
