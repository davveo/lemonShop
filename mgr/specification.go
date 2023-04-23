package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type SpecificationMgr struct {
	*_BaseMgr
}

func NewSpecificationMgr(db db.Repo) *SpecificationMgr {
	if db == nil {
		panic(fmt.Errorf("EsSpecificationDao need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &SpecificationMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_specification"),
		wdb:       db.GetDbW().Table("es_specification"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *SpecificationMgr) GetTableName() string {
	return "es_specification"
}

// Reset 重置gorm会话
func (obj *SpecificationMgr) Reset() *SpecificationMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *SpecificationMgr) Get() (result models.EsSpecification, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *SpecificationMgr) Gets() (results []*models.EsSpecification, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *SpecificationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithSpecID spec_id获取 主键
func (obj *SpecificationMgr) WithSpecID(specID int) Option {
	return optionFunc(func(o *options) { o.query["spec_id"] = specID })
}

// WithSpecName spec_name获取 规格项名称
func (obj *SpecificationMgr) WithSpecName(specName string) Option {
	return optionFunc(func(o *options) { o.query["spec_name"] = specName })
}

// WithDisabled disabled获取 是否被删除0 删除   1  没有删除
func (obj *SpecificationMgr) WithDisabled(disabled int) Option {
	return optionFunc(func(o *options) { o.query["disabled"] = disabled })
}

// WithSpecMemo spec_memo获取 规格描述
func (obj *SpecificationMgr) WithSpecMemo(specMemo string) Option {
	return optionFunc(func(o *options) { o.query["spec_memo"] = specMemo })
}

// WithSellerID seller_id获取 所属卖家
func (obj *SpecificationMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// GetByOption 功能选项模式获取
func (obj *SpecificationMgr) GetByOption(opts ...Option) (result models.EsSpecification, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *SpecificationMgr) GetByOptions(opts ...Option) (results []*models.EsSpecification, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *SpecificationMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSpecification, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Where(options.query)
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

// GetFromSpecID 通过spec_id获取内容 主键
func (obj *SpecificationMgr) GetFromSpecID(specID int) (result models.EsSpecification, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Where("`spec_id` = ?", specID).First(&result).Error

	return
}

// GetBatchFromSpecID 批量查找 主键
func (obj *SpecificationMgr) GetBatchFromSpecID(specIDs []int) (results []*models.EsSpecification, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Where("`spec_id` IN (?)", specIDs).Find(&results).Error

	return
}

// GetFromSpecName 通过spec_name获取内容 规格项名称
func (obj *SpecificationMgr) GetFromSpecName(specName string) (results []*models.EsSpecification, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Where("`spec_name` = ?", specName).Find(&results).Error

	return
}

// GetBatchFromSpecName 批量查找 规格项名称
func (obj *SpecificationMgr) GetBatchFromSpecName(specNames []string) (results []*models.EsSpecification, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Where("`spec_name` IN (?)", specNames).Find(&results).Error

	return
}

// GetFromDisabled 通过disabled获取内容 是否被删除0 删除   1  没有删除
func (obj *SpecificationMgr) GetFromDisabled(disabled int) (results []*models.EsSpecification, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Where("`disabled` = ?", disabled).Find(&results).Error

	return
}

// GetBatchFromDisabled 批量查找 是否被删除0 删除   1  没有删除
func (obj *SpecificationMgr) GetBatchFromDisabled(disableds []int) (results []*models.EsSpecification, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Where("`disabled` IN (?)", disableds).Find(&results).Error

	return
}

// GetFromSpecMemo 通过spec_memo获取内容 规格描述
func (obj *SpecificationMgr) GetFromSpecMemo(specMemo string) (results []*models.EsSpecification, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Where("`spec_memo` = ?", specMemo).Find(&results).Error

	return
}

// GetBatchFromSpecMemo 批量查找 规格描述
func (obj *SpecificationMgr) GetBatchFromSpecMemo(specMemos []string) (results []*models.EsSpecification, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Where("`spec_memo` IN (?)", specMemos).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 所属卖家
func (obj *SpecificationMgr) GetFromSellerID(sellerID int) (results []*models.EsSpecification, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 所属卖家
func (obj *SpecificationMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsSpecification, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *SpecificationMgr) FetchByPrimaryKey(specID int) (result *models.EsSpecification, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Where("`spec_id` = ?", specID).First(result).Error

	return
}

func (obj *SpecificationMgr) Create(specs *models.EsSpecification) (err error) {
	err = obj.wdb.WithContext(obj.ctx).Create(specs).Error

	return
}

func (obj *SpecificationMgr) Update(specID int, updates map[string]interface{}) (err error) {
	err = obj.wdb.WithContext(obj.ctx).Model(models.EsSpecification{}).
		Where("`spec_id` = ?", specID).Updates(updates).Error

	return
}

func (obj *SpecificationMgr) Raw(sql string, values ...interface{}) (err error) {
	//tx := obj.rdb.WithContext(obj.ctx).Raw(sql, values)
	return nil
}
