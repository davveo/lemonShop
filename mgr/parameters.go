package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type ParametersMgr struct {
	*_BaseMgr
}

// NewParametersMgr open func
func NewParametersMgr(db db.Repo) *ParametersMgr {
	if db == nil {
		panic(fmt.Errorf("NewParametersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &ParametersMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_parameters"), wdb: db.GetDbW().Table("es_parameters"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *ParametersMgr) GetTableName() string {
	return "es_parameters"
}

// Reset 重置gorm会话
func (obj *ParametersMgr) Reset() *ParametersMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *ParametersMgr) Get() (result models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *ParametersMgr) Gets() (results []*models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *ParametersMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithParamID param_id获取 主键
func (obj *ParametersMgr) WithParamID(paramID int) Option {
	return optionFunc(func(o *options) { o.query["param_id"] = paramID })
}

// WithParamName param_name获取 参数名称
func (obj *ParametersMgr) WithParamName(paramName string) Option {
	return optionFunc(func(o *options) { o.query["param_name"] = paramName })
}

// WithParamType param_type获取 参数类型1 输入项   2 选择项
func (obj *ParametersMgr) WithParamType(paramType int) Option {
	return optionFunc(func(o *options) { o.query["param_type"] = paramType })
}

// WithOptions options获取 选择值，当参数类型是选择项2时，必填，逗号分隔
func (obj *ParametersMgr) WithOptions(opt string) Option {
	return optionFunc(func(o *options) { o.query["options"] = opt })
}

// WithIsIndex is_index获取 是否可索引，0 不显示 1 显示
func (obj *ParametersMgr) WithIsIndex(isIndex int) Option {
	return optionFunc(func(o *options) { o.query["is_index"] = isIndex })
}

// WithRequired required获取 是否必填是  1    否   0
func (obj *ParametersMgr) WithRequired(required int) Option {
	return optionFunc(func(o *options) { o.query["required"] = required })
}

// WithGroupID group_id获取 参数分组id
func (obj *ParametersMgr) WithGroupID(groupID int) Option {
	return optionFunc(func(o *options) { o.query["group_id"] = groupID })
}

// WithCategoryID category_id获取 分类id
func (obj *ParametersMgr) WithCategoryID(categoryID int) Option {
	return optionFunc(func(o *options) { o.query["category_id"] = categoryID })
}

// WithSort sort获取 排序
func (obj *ParametersMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// GetByOption 功能选项模式获取
func (obj *ParametersMgr) GetByOption(opts ...Option) (result models.EsParameters, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *ParametersMgr) GetByOptions(opts ...Option) (results []*models.EsParameters, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *ParametersMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsParameters, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where(options.query)
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
func (obj *ParametersMgr) GetFromParamID(paramID int) (result models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`param_id` = ?", paramID).First(&result).Error

	return
}

// GetBatchFromParamID 批量查找 主键
func (obj *ParametersMgr) GetBatchFromParamID(paramIDs []int) (results []*models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`param_id` IN (?)", paramIDs).Find(&results).Error

	return
}

// GetFromParamName 通过param_name获取内容 参数名称
func (obj *ParametersMgr) GetFromParamName(paramName string) (results []*models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`param_name` = ?", paramName).Find(&results).Error

	return
}

// GetBatchFromParamName 批量查找 参数名称
func (obj *ParametersMgr) GetBatchFromParamName(paramNames []string) (results []*models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`param_name` IN (?)", paramNames).Find(&results).Error

	return
}

// GetFromParamType 通过param_type获取内容 参数类型1 输入项   2 选择项
func (obj *ParametersMgr) GetFromParamType(paramType int) (results []*models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`param_type` = ?", paramType).Find(&results).Error

	return
}

// GetBatchFromParamType 批量查找 参数类型1 输入项   2 选择项
func (obj *ParametersMgr) GetBatchFromParamType(paramTypes []int) (results []*models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`param_type` IN (?)", paramTypes).Find(&results).Error

	return
}

// GetFromOptions 通过options获取内容 选择值，当参数类型是选择项2时，必填，逗号分隔
func (obj *ParametersMgr) GetFromOptions(options string) (results []*models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`options` = ?", options).Find(&results).Error

	return
}

// GetBatchFromOptions 批量查找 选择值，当参数类型是选择项2时，必填，逗号分隔
func (obj *ParametersMgr) GetBatchFromOptions(optionss []string) (results []*models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`options` IN (?)", optionss).Find(&results).Error

	return
}

// GetFromIsIndex 通过is_index获取内容 是否可索引，0 不显示 1 显示
func (obj *ParametersMgr) GetFromIsIndex(isIndex int) (results []*models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`is_index` = ?", isIndex).Find(&results).Error

	return
}

// GetBatchFromIsIndex 批量查找 是否可索引，0 不显示 1 显示
func (obj *ParametersMgr) GetBatchFromIsIndex(isIndexs []int) (results []*models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`is_index` IN (?)", isIndexs).Find(&results).Error

	return
}

// GetFromRequired 通过required获取内容 是否必填是  1    否   0
func (obj *ParametersMgr) GetFromRequired(required int) (results []*models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`required` = ?", required).Find(&results).Error

	return
}

// GetBatchFromRequired 批量查找 是否必填是  1    否   0
func (obj *ParametersMgr) GetBatchFromRequired(requireds []int) (results []*models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`required` IN (?)", requireds).Find(&results).Error

	return
}

// GetFromGroupID 通过group_id获取内容 参数分组id
func (obj *ParametersMgr) GetFromGroupID(groupID int) (results []*models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`group_id` = ?", groupID).Find(&results).Error

	return
}

// GetBatchFromGroupID 批量查找 参数分组id
func (obj *ParametersMgr) GetBatchFromGroupID(groupIDs []int) (results []*models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`group_id` IN (?)", groupIDs).Find(&results).Error

	return
}

// GetFromCategoryID 通过category_id获取内容 分类id
func (obj *ParametersMgr) GetFromCategoryID(categoryID int) (results []*models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`category_id` = ?", categoryID).Find(&results).Error

	return
}

// GetBatchFromCategoryID 批量查找 分类id
func (obj *ParametersMgr) GetBatchFromCategoryID(categoryIDs []int) (results []*models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`category_id` IN (?)", categoryIDs).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *ParametersMgr) GetFromSort(sort int) (results []*models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *ParametersMgr) GetBatchFromSort(sorts []int) (results []*models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *ParametersMgr) FetchByPrimaryKey(paramID int) (result models.EsParameters, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsParameters{}).Where("`param_id` = ?", paramID).First(&result).Error

	return
}
