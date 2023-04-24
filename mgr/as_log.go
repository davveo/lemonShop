package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type AsLogMgr struct {
	*_BaseMgr
}

// NewAsLogMgr open func
func NewAsLogMgr(db db.Repo) *AsLogMgr {
	if db == nil {
		panic(fmt.Errorf("NewAsLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &AsLogMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_category_spec"),
		wdb:       db.GetDbW().Table("es_category_spec"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *AsLogMgr) GetTableName() string {
	return "es_as_log"
}

// Reset 重置gorm会话
func (obj *AsLogMgr) Reset() *AsLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *AsLogMgr) Get() (result models.EsAsLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *AsLogMgr) Gets() (results []*models.EsAsLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *AsLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *AsLogMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithSn sn获取 售后/退款编号
func (obj *AsLogMgr) WithSn(sn string) Option {
	return optionFunc(func(o *options) { o.query["sn"] = sn })
}

// WithLogTime log_time获取 创建时间
func (obj *AsLogMgr) WithLogTime(logTime int64) Option {
	return optionFunc(func(o *options) { o.query["log_time"] = logTime })
}

// WithLogDetail log_detail获取 详细信息
func (obj *AsLogMgr) WithLogDetail(logDetail string) Option {
	return optionFunc(func(o *options) { o.query["log_detail"] = logDetail })
}

// WithOperator operator获取 操作者
func (obj *AsLogMgr) WithOperator(operator string) Option {
	return optionFunc(func(o *options) { o.query["operator"] = operator })
}

// GetByOption 功能选项模式获取
func (obj *AsLogMgr) GetByOption(opts ...Option) (result models.EsAsLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *AsLogMgr) GetByOptions(opts ...Option) (results []*models.EsAsLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *AsLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsAsLog, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).Where(options.query)
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
func (obj *AsLogMgr) GetFromID(id int) (result models.EsAsLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *AsLogMgr) GetBatchFromID(ids []int) (results []*models.EsAsLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromSn 通过sn获取内容 售后/退款编号
func (obj *AsLogMgr) GetFromSn(sn string) (results []*models.EsAsLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}

// GetBatchFromSn 批量查找 售后/退款编号
func (obj *AsLogMgr) GetBatchFromSn(sns []string) (results []*models.EsAsLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).Where("`sn` IN (?)", sns).Find(&results).Error

	return
}

// GetFromLogTime 通过log_time获取内容 创建时间
func (obj *AsLogMgr) GetFromLogTime(logTime int64) (results []*models.EsAsLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).Where("`log_time` = ?", logTime).Find(&results).Error

	return
}

// GetBatchFromLogTime 批量查找 创建时间
func (obj *AsLogMgr) GetBatchFromLogTime(logTimes []int64) (results []*models.EsAsLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).Where("`log_time` IN (?)", logTimes).Find(&results).Error

	return
}

// GetFromLogDetail 通过log_detail获取内容 详细信息
func (obj *AsLogMgr) GetFromLogDetail(logDetail string) (results []*models.EsAsLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).Where("`log_detail` = ?", logDetail).Find(&results).Error

	return
}

// GetBatchFromLogDetail 批量查找 详细信息
func (obj *AsLogMgr) GetBatchFromLogDetail(logDetails []string) (results []*models.EsAsLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).Where("`log_detail` IN (?)", logDetails).Find(&results).Error

	return
}

// GetFromOperator 通过operator获取内容 操作者
func (obj *AsLogMgr) GetFromOperator(operator string) (results []*models.EsAsLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).Where("`operator` = ?", operator).Find(&results).Error

	return
}

// GetBatchFromOperator 批量查找 操作者
func (obj *AsLogMgr) GetBatchFromOperator(operators []string) (results []*models.EsAsLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).Where("`operator` IN (?)", operators).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *AsLogMgr) FetchByPrimaryKey(id int) (result models.EsAsLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByEsIndexID  获取多个内容
func (obj *AsLogMgr) FetchIndexByEsIndexID(id int) (results []*models.EsAsLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// FetchIndexByEsIndexSn  获取多个内容
func (obj *AsLogMgr) FetchIndexByEsIndexSn(sn string) (results []*models.EsAsLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsLog{}).Where("`sn` = ?", sn).Find(&results).Error

	return
}
