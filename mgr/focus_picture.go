package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsFocusPictureMgr struct {
	*_BaseMgr
}

// EsFocusPictureMgr open func
func EsFocusPictureMgr(db *gorm.DB) *_EsFocusPictureMgr {
	if db == nil {
		panic(fmt.Errorf("EsFocusPictureMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsFocusPictureMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_focus_picture"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsFocusPictureMgr) GetTableName() string {
	return "es_focus_picture"
}

// Reset 重置gorm会话
func (obj *_EsFocusPictureMgr) Reset() *_EsFocusPictureMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsFocusPictureMgr) Get() (result models.EsFocusPicture, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsFocusPictureMgr) Gets() (results []*models.EsFocusPicture, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsFocusPictureMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键id
func (obj *_EsFocusPictureMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithPicURL pic_url获取 图片地址
func (obj *_EsFocusPictureMgr) WithPicURL(picURL string) Option {
	return optionFunc(func(o *options) { o.query["pic_url"] = picURL })
}

// WithOperationType operation_type获取 操作类型
func (obj *_EsFocusPictureMgr) WithOperationType(operationType string) Option {
	return optionFunc(func(o *options) { o.query["operation_type"] = operationType })
}

// WithOperationParam operation_param获取 操作参数
func (obj *_EsFocusPictureMgr) WithOperationParam(operationParam string) Option {
	return optionFunc(func(o *options) { o.query["operation_param"] = operationParam })
}

// WithOperationURL operation_url获取 操作地址
func (obj *_EsFocusPictureMgr) WithOperationURL(operationURL string) Option {
	return optionFunc(func(o *options) { o.query["operation_url"] = operationURL })
}

// WithClientType client_type获取 客户端类型 pc:pc楼层mobile:移动端楼层
func (obj *_EsFocusPictureMgr) WithClientType(clientType string) Option {
	return optionFunc(func(o *options) { o.query["client_type"] = clientType })
}

// GetByOption 功能选项模式获取
func (obj *_EsFocusPictureMgr) GetByOption(opts ...Option) (result models.EsFocusPicture, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsFocusPictureMgr) GetByOptions(opts ...Option) (results []*models.EsFocusPicture, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsFocusPictureMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsFocusPicture, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where(options.query)
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
func (obj *_EsFocusPictureMgr) GetFromID(id int) (result models.EsFocusPicture, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键id
func (obj *_EsFocusPictureMgr) GetBatchFromID(ids []int) (results []*models.EsFocusPicture, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromPicURL 通过pic_url获取内容 图片地址
func (obj *_EsFocusPictureMgr) GetFromPicURL(picURL string) (results []*models.EsFocusPicture, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`pic_url` = ?", picURL).Find(&results).Error

	return
}

// GetBatchFromPicURL 批量查找 图片地址
func (obj *_EsFocusPictureMgr) GetBatchFromPicURL(picURLs []string) (results []*models.EsFocusPicture, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`pic_url` IN (?)", picURLs).Find(&results).Error

	return
}

// GetFromOperationType 通过operation_type获取内容 操作类型
func (obj *_EsFocusPictureMgr) GetFromOperationType(operationType string) (results []*models.EsFocusPicture, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`operation_type` = ?", operationType).Find(&results).Error

	return
}

// GetBatchFromOperationType 批量查找 操作类型
func (obj *_EsFocusPictureMgr) GetBatchFromOperationType(operationTypes []string) (results []*models.EsFocusPicture, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`operation_type` IN (?)", operationTypes).Find(&results).Error

	return
}

// GetFromOperationParam 通过operation_param获取内容 操作参数
func (obj *_EsFocusPictureMgr) GetFromOperationParam(operationParam string) (results []*models.EsFocusPicture, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`operation_param` = ?", operationParam).Find(&results).Error

	return
}

// GetBatchFromOperationParam 批量查找 操作参数
func (obj *_EsFocusPictureMgr) GetBatchFromOperationParam(operationParams []string) (results []*models.EsFocusPicture, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`operation_param` IN (?)", operationParams).Find(&results).Error

	return
}

// GetFromOperationURL 通过operation_url获取内容 操作地址
func (obj *_EsFocusPictureMgr) GetFromOperationURL(operationURL string) (results []*models.EsFocusPicture, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`operation_url` = ?", operationURL).Find(&results).Error

	return
}

// GetBatchFromOperationURL 批量查找 操作地址
func (obj *_EsFocusPictureMgr) GetBatchFromOperationURL(operationURLs []string) (results []*models.EsFocusPicture, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`operation_url` IN (?)", operationURLs).Find(&results).Error

	return
}

// GetFromClientType 通过client_type获取内容 客户端类型 pc:pc楼层mobile:移动端楼层
func (obj *_EsFocusPictureMgr) GetFromClientType(clientType string) (results []*models.EsFocusPicture, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`client_type` = ?", clientType).Find(&results).Error

	return
}

// GetBatchFromClientType 批量查找 客户端类型 pc:pc楼层mobile:移动端楼层
func (obj *_EsFocusPictureMgr) GetBatchFromClientType(clientTypes []string) (results []*models.EsFocusPicture, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`client_type` IN (?)", clientTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsFocusPictureMgr) FetchByPrimaryKey(id int) (result models.EsFocusPicture, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsFocusPicture{}).Where("`id` = ?", id).First(&result).Error

	return
}
