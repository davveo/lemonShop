package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models/vo"

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
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1},
	}
}

// GetTableName get sql table name.获取数据库名字
func (obj *SpecificationMgr) GetTableName() string {
	return "es_specification"
}

// Get 获取
func (obj *SpecificationMgr) Get() (result models.EsSpecification, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).First(&result).Error

	return
}

func (obj *SpecificationMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Count(count)
}

// WithSpecID spec_id获取 主键
func (obj *SpecificationMgr) WithSpecID(specID int) Option {
	return optionFunc(func(o *options) { o.query["spec_id"] = specID })
}

// WithInSpecID spec_id获取 主键
func (obj *SpecificationMgr) WithInSpecID(specID []int64) Option {
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
	var specName string
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	if val, ok := options.query["spec_name"]; ok {
		delete(options.query, "spec_name")
		specName = val.(string)
	}
	resultPage = page
	results := make([]models.EsSpecification, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Where(options.query)

	if specName != "" {
		query.Where("spec_name like ?", "%"+specName+"%")
	}

	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

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

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *SpecificationMgr) FetchByPrimaryKey(specID int) (result *models.EsSpecification, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSpecification{}).Where("`spec_id` = ?", specID).First(&result).Error

	return
}

func (obj *SpecificationMgr) Create(specs *models.EsSpecification) (err error) {
	err = obj.wdb.WithContext(obj.ctx).Create(specs).Error

	return
}

func (obj *SpecificationMgr) Update(updates map[string]interface{}, opts ...Option) (err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.wdb.WithContext(obj.ctx).Model(models.EsSpecification{}).
		Where(options.query).Updates(updates).Error
	return
}

func (obj *SpecificationMgr) UpdateByModel(updates *models.EsSpecification, opts ...Option) (err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.wdb.WithContext(obj.ctx).Model(models.EsSpecification{}).
		Where(options.query).Updates(updates).Error
	return
}

func (obj *SpecificationMgr) DeleteBatch(specIDs []int) (err error) {
	return nil
}

func (obj *SpecificationMgr) RawToStruct(where string, values ...interface{}) (res []*models.EsSpecification, err error) {
	err = obj.rdb.WithContext(obj.ctx).Raw(where, values...).Find(&res).Error

	return
}

func (obj *SpecificationMgr) RawToMap(where string, values ...interface{}) (res []map[string]interface{}, err error) {
	//err = obj.rdb.WithContext(obj.ctx).Raw(obj.baseSql+where, values...).Find(&res).Error
	err = obj.rdb.WithContext(obj.ctx).Raw(where, values...).Find(&res).Error

	return
}

func (obj *SpecificationMgr) RawWithSelect(sql string, values ...interface{}) (res []*vo.SelectVo, err error) {
	err = obj.rdb.WithContext(obj.ctx).Raw(sql, values...).Find(&res).Error

	return
}
