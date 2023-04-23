package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsReceiptFileMgr struct {
	*_BaseMgr
}

// EsReceiptFileMgr open func
func EsReceiptFileMgr(db *gorm.DB) *_EsReceiptFileMgr {
	if db == nil {
		panic(fmt.Errorf("EsReceiptFileMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsReceiptFileMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_receipt_file"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsReceiptFileMgr) GetTableName() string {
	return "es_receipt_file"
}

// Reset 重置gorm会话
func (obj *_EsReceiptFileMgr) Reset() *_EsReceiptFileMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsReceiptFileMgr) Get() (result models.EsReceiptFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptFile{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsReceiptFileMgr) Gets() (results []*models.EsReceiptFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptFile{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsReceiptFileMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsReceiptFile{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_EsReceiptFileMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithHistoryID history_id获取
func (obj *_EsReceiptFileMgr) WithHistoryID(historyID int) Option {
	return optionFunc(func(o *options) { o.query["history_id"] = historyID })
}

// WithElecFile elec_file获取
func (obj *_EsReceiptFileMgr) WithElecFile(elecFile string) Option {
	return optionFunc(func(o *options) { o.query["elec_file"] = elecFile })
}

// GetByOption 功能选项模式获取
func (obj *_EsReceiptFileMgr) GetByOption(opts ...Option) (result models.EsReceiptFile, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptFile{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsReceiptFileMgr) GetByOptions(opts ...Option) (results []*models.EsReceiptFile, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptFile{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsReceiptFileMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsReceiptFile, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsReceiptFile{}).Where(options.query)
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

// GetFromID 通过id获取内容
func (obj *_EsReceiptFileMgr) GetFromID(id int) (result models.EsReceiptFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptFile{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_EsReceiptFileMgr) GetBatchFromID(ids []int) (results []*models.EsReceiptFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptFile{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromHistoryID 通过history_id获取内容
func (obj *_EsReceiptFileMgr) GetFromHistoryID(historyID int) (results []*models.EsReceiptFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptFile{}).Where("`history_id` = ?", historyID).Find(&results).Error

	return
}

// GetBatchFromHistoryID 批量查找
func (obj *_EsReceiptFileMgr) GetBatchFromHistoryID(historyIDs []int) (results []*models.EsReceiptFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptFile{}).Where("`history_id` IN (?)", historyIDs).Find(&results).Error

	return
}

// GetFromElecFile 通过elec_file获取内容
func (obj *_EsReceiptFileMgr) GetFromElecFile(elecFile string) (results []*models.EsReceiptFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptFile{}).Where("`elec_file` = ?", elecFile).Find(&results).Error

	return
}

// GetBatchFromElecFile 批量查找
func (obj *_EsReceiptFileMgr) GetBatchFromElecFile(elecFiles []string) (results []*models.EsReceiptFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptFile{}).Where("`elec_file` IN (?)", elecFiles).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsReceiptFileMgr) FetchByPrimaryKey(id int) (result models.EsReceiptFile, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsReceiptFile{}).Where("`id` = ?", id).First(&result).Error

	return
}
