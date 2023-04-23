package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsParametersMgr struct {
	*_BaseMgr
}

// EsParametersMgr open func
func EsParametersMgr(db *gorm.DB) *_EsParametersMgr {
	if db == nil {
		panic(fmt.Errorf("EsParametersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsParametersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_parameters"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsParametersMgr) GetTableName() string {
	return "es_parameters"
}

// Reset 重置gorm会话
func (obj *_EsParametersMgr) Reset() *_EsParametersMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsParametersMgr) Get() (result models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsParametersMgr) Gets() (results []*models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsParametersMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithParamID param_id获取 主键
func (obj *_EsParametersMgr) WithParamID(paramID int) Option {
	return optionFunc(func(o *options) { o.query["param_id"] = paramID })
}

// WithParamName param_name获取 参数名称
func (obj *_EsParametersMgr) WithParamName(paramName string) Option {
	return optionFunc(func(o *options) { o.query["param_name"] = paramName })
}

// WithParamType param_type获取 参数类型1 输入项   2 选择项
func (obj *_EsParametersMgr) WithParamType(paramType int) Option {
	return optionFunc(func(o *options) { o.query["param_type"] = paramType })
}

// WithOptions options获取 选择值，当参数类型是选择项2时，必填，逗号分隔
func (obj *_EsParametersMgr) WithOptions(options string) Option {
	return optionFunc(func(o *options) { o.query["options"] = options })
}

// WithIsIndex is_index获取 是否可索引，0 不显示 1 显示
func (obj *_EsParametersMgr) WithIsIndex(isIndex int) Option {
	return optionFunc(func(o *options) { o.query["is_index"] = isIndex })
}

// WithRequired required获取 是否必填是  1    否   0
func (obj *_EsParametersMgr) WithRequired(required int) Option {
	return optionFunc(func(o *options) { o.query["required"] = required })
}

// WithGroupID group_id获取 参数分组id
func (obj *_EsParametersMgr) WithGroupID(groupID int) Option {
	return optionFunc(func(o *options) { o.query["group_id"] = groupID })
}

// WithCategoryID category_id获取 分类id
func (obj *_EsParametersMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithSort sort获取 排序
func (obj *_EsParametersMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// GetByOption 功能选项模式获取
func (obj *_EsParametersMgr) GetByOption(opts ...Option) (result models.EsParameters, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsParametersMgr) GetByOptions(opts ...Option) (results []*models.EsParameters, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsParametersMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsParameters, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where(options.query)
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

// GetFromParamID 通过param_id获取内容 主键
func (obj *_EsParametersMgr) GetFromParamID(paramID int) (result models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`param_id` = ?", paramID).First(&result).Error

	return
}

// GetBatchFromParamID 批量查找 主键
func (obj *_EsParametersMgr) GetBatchFromParamID(paramIDs []int) (results []*models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`param_id` IN (?)", paramIDs).Find(&results).Error

	return
}

// GetFromParamName 通过param_name获取内容 参数名称
func (obj *_EsParametersMgr) GetFromParamName(paramName string) (results []*models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`param_name` = ?", paramName).Find(&results).Error

	return
}

// GetBatchFromParamName 批量查找 参数名称
func (obj *_EsParametersMgr) GetBatchFromParamName(paramNames []string) (results []*models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`param_name` IN (?)", paramNames).Find(&results).Error

	return
}

// GetFromParamType 通过param_type获取内容 参数类型1 输入项   2 选择项
func (obj *_EsParametersMgr) GetFromParamType(paramType int) (results []*models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`param_type` = ?", paramType).Find(&results).Error

	return
}

// GetBatchFromParamType 批量查找 参数类型1 输入项   2 选择项
func (obj *_EsParametersMgr) GetBatchFromParamType(paramTypes []int) (results []*models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`param_type` IN (?)", paramTypes).Find(&results).Error

	return
}

// GetFromOptions 通过options获取内容 选择值，当参数类型是选择项2时，必填，逗号分隔
func (obj *_EsParametersMgr) GetFromOptions(options string) (results []*models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`options` = ?", options).Find(&results).Error

	return
}

// GetBatchFromOptions 批量查找 选择值，当参数类型是选择项2时，必填，逗号分隔
func (obj *_EsParametersMgr) GetBatchFromOptions(optionss []string) (results []*models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`options` IN (?)", optionss).Find(&results).Error

	return
}

// GetFromIsIndex 通过is_index获取内容 是否可索引，0 不显示 1 显示
func (obj *_EsParametersMgr) GetFromIsIndex(isIndex int) (results []*models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`is_index` = ?", isIndex).Find(&results).Error

	return
}

// GetBatchFromIsIndex 批量查找 是否可索引，0 不显示 1 显示
func (obj *_EsParametersMgr) GetBatchFromIsIndex(isIndexs []int) (results []*models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`is_index` IN (?)", isIndexs).Find(&results).Error

	return
}

// GetFromRequired 通过required获取内容 是否必填是  1    否   0
func (obj *_EsParametersMgr) GetFromRequired(required int) (results []*models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`required` = ?", required).Find(&results).Error

	return
}

// GetBatchFromRequired 批量查找 是否必填是  1    否   0
func (obj *_EsParametersMgr) GetBatchFromRequired(requireds []int) (results []*models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`required` IN (?)", requireds).Find(&results).Error

	return
}

// GetFromGroupID 通过group_id获取内容 参数分组id
func (obj *_EsParametersMgr) GetFromGroupID(groupID int) (results []*models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`group_id` = ?", groupID).Find(&results).Error

	return
}

// GetBatchFromGroupID 批量查找 参数分组id
func (obj *_EsParametersMgr) GetBatchFromGroupID(groupIDs []int) (results []*models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`group_id` IN (?)", groupIDs).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容 分类id
func (obj *_EsParametersMgr) GetFromCategoryID(categoryID int) (results []*models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 分类id
func (obj *_EsParametersMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_EsParametersMgr) GetFromSort(sort int) (results []*models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_EsParametersMgr) GetBatchFromSort(sorts []int) (results []*models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsParametersMgr) FetchByPrimaryKey(paramID int) (result models.EsParameters, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`param_id` = ?", paramID).First(&result).Error

	return
}
