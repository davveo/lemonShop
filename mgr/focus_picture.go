package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type FocusPictureMgr struct {
	*_BaseMgr
}

// NewFocusPictureMgr open func
func NewFocusPictureMgr(db db.Repo) *FocusPictureMgr {
	if db == nil {
		panic(fmt.Errorf("NewFocusPictureMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &FocusPictureMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_focus_picture"),
		wdb:       db.GetDbW().Table("es_focus_picture"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *FocusPictureMgr) GetTableName() string {
	return "es_focus_picture"
}

// Reset 重置gorm会话
func (obj *FocusPictureMgr) Reset() *FocusPictureMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *FocusPictureMgr) Get() (result models.EsFocusPicture, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *FocusPictureMgr) Gets() (results []*models.EsFocusPicture, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *FocusPictureMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *FocusPictureMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPicURL pic_url获取 图片地址
func (obj *FocusPictureMgr) WithPicURL(picURL string) Option {
	return optionFunc(func(o *options) { o.query["pic_url"] = picURL })
}

// WithOperationType operation_type获取 操作类型
func (obj *FocusPictureMgr) WithOperationType(operationType string) Option {
	return optionFunc(func(o *options) { o.query["operation_type"] = operationType })
}

// WithOperationParam operation_param获取 操作参数
func (obj *FocusPictureMgr) WithOperationParam(operationParam string) Option {
	return optionFunc(func(o *options) { o.query["operation_param"] = operationParam })
}

// WithOperationURL operation_url获取 操作地址
func (obj *FocusPictureMgr) WithOperationURL(operationURL string) Option {
	return optionFunc(func(o *options) { o.query["operation_url"] = operationURL })
}

// WithClientType client_type获取 客户端类型 pc:pc楼层mobile:移动端楼层
func (obj *FocusPictureMgr) WithClientType(clientType string) Option {
	return optionFunc(func(o *options) { o.query["client_type"] = clientType })
}

// GetByOption 功能选项模式获取
func (obj *FocusPictureMgr) GetByOption(opts ...Option) (result models.EsFocusPicture, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *FocusPictureMgr) GetByOptions(opts ...Option) (results []*models.EsFocusPicture, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *FocusPictureMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsFocusPicture, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键id
func (obj *FocusPictureMgr) GetFromID(id int) (result models.EsFocusPicture, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *FocusPictureMgr) GetBatchFromID(ids []int) (results []*models.EsFocusPicture, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPicURL 通过pic_url获取内容 图片地址
func (obj *FocusPictureMgr) GetFromPicURL(picURL string) (results []*models.EsFocusPicture, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`pic_url` = ?", picURL).Find(&results).Error

	return
}

// GetBatchFromPicURL 批量查找 图片地址
func (obj *FocusPictureMgr) GetBatchFromPicURL(picURLs []string) (results []*models.EsFocusPicture, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`pic_url` IN (?)", picURLs).Find(&results).Error

	return
}

// GetFromOperationType 通过operation_type获取内容 操作类型
func (obj *FocusPictureMgr) GetFromOperationType(operationType string) (results []*models.EsFocusPicture, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`operation_type` = ?", operationType).Find(&results).Error

	return
}

// GetBatchFromOperationType 批量查找 操作类型
func (obj *FocusPictureMgr) GetBatchFromOperationType(operationTypes []string) (results []*models.EsFocusPicture, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`operation_type` IN (?)", operationTypes).Find(&results).Error

	return
}

// GetFromOperationParam 通过operation_param获取内容 操作参数
func (obj *FocusPictureMgr) GetFromOperationParam(operationParam string) (results []*models.EsFocusPicture, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`operation_param` = ?", operationParam).Find(&results).Error

	return
}

// GetBatchFromOperationParam 批量查找 操作参数
func (obj *FocusPictureMgr) GetBatchFromOperationParam(operationParams []string) (results []*models.EsFocusPicture, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`operation_param` IN (?)", operationParams).Find(&results).Error

	return
}

// GetFromOperationURL 通过operation_url获取内容 操作地址
func (obj *FocusPictureMgr) GetFromOperationURL(operationURL string) (results []*models.EsFocusPicture, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`operation_url` = ?", operationURL).Find(&results).Error

	return
}

// GetBatchFromOperationURL 批量查找 操作地址
func (obj *FocusPictureMgr) GetBatchFromOperationURL(operationURLs []string) (results []*models.EsFocusPicture, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`operation_url` IN (?)", operationURLs).Find(&results).Error

	return
}

// GetFromClientType 通过client_type获取内容 客户端类型 pc:pc楼层mobile:移动端楼层
func (obj *FocusPictureMgr) GetFromClientType(clientType string) (results []*models.EsFocusPicture, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`client_type` = ?", clientType).Find(&results).Error

	return
}

// GetBatchFromClientType 批量查找 客户端类型 pc:pc楼层mobile:移动端楼层
func (obj *FocusPictureMgr) GetBatchFromClientType(clientTypes []string) (results []*models.EsFocusPicture, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`client_type` IN (?)", clientTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *FocusPictureMgr) FetchByPrimaryKey(id int) (result models.EsFocusPicture, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`id` = ?", id).First(&result).Error

	return
}
