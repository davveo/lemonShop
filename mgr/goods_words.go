package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsGoodsWordsMgr struct {
	*_BaseMgr
}

// EsGoodsWordsMgr open func
func EsGoodsWordsMgr(db *gorm.DB) *_EsGoodsWordsMgr {
	if db == nil {
		panic(fmt.Errorf("EsGoodsWordsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsGoodsWordsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_goods_words"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsGoodsWordsMgr) GetTableName() string {
	return "es_goods_words"
}

// Reset 重置gorm会话
func (obj *_EsGoodsWordsMgr) Reset() *_EsGoodsWordsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsGoodsWordsMgr) Get() (result models.EsGoodsWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsGoodsWordsMgr) Gets() (results []*models.EsGoodsWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsGoodsWordsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_EsGoodsWordsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithWords words获取 分词
func (obj *_EsGoodsWordsMgr) WithWords(words string) Option {
	return optionFunc(func(o *options) { o.query["words"] = words })
}

// WithGoodsNum goods_num获取 数量
func (obj *_EsGoodsWordsMgr) WithGoodsNum(goodsNum int64) Option {
	return optionFunc(func(o *options) { o.query["goods_num"] = goodsNum })
}

// WithQuanpin quanpin获取 全拼字母
func (obj *_EsGoodsWordsMgr) WithQuanpin(quanpin string) Option {
	return optionFunc(func(o *options) { o.query["quanpin"] = quanpin })
}

// WithSzm szm获取 首字母
func (obj *_EsGoodsWordsMgr) WithSzm(szm string) Option {
	return optionFunc(func(o *options) { o.query["szm"] = szm })
}

// WithType type获取 类型，系统(SYSTEM)，平台(PLATFORM)
func (obj *_EsGoodsWordsMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithSort sort获取 排序
func (obj *_EsGoodsWordsMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// GetByOption 功能选项模式获取
func (obj *_EsGoodsWordsMgr) GetByOption(opts ...Option) (result models.EsGoodsWords, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsGoodsWordsMgr) GetByOptions(opts ...Option) (results []*models.EsGoodsWords, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsGoodsWordsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsGoodsWords, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键
func (obj *_EsGoodsWordsMgr) GetFromID(id int) (result models.EsGoodsWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_EsGoodsWordsMgr) GetBatchFromID(ids []int) (results []*models.EsGoodsWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromWords 通过words获取内容 分词
func (obj *_EsGoodsWordsMgr) GetFromWords(words string) (results []*models.EsGoodsWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`words` = ?", words).Find(&results).Error

	return
}

// GetBatchFromWords 批量查找 分词
func (obj *_EsGoodsWordsMgr) GetBatchFromWords(wordss []string) (results []*models.EsGoodsWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`words` IN (?)", wordss).Find(&results).Error

	return
}

// GetFromGoodsNum 通过goods_num获取内容 数量
func (obj *_EsGoodsWordsMgr) GetFromGoodsNum(goodsNum int64) (results []*models.EsGoodsWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`goods_num` = ?", goodsNum).Find(&results).Error

	return
}

// GetBatchFromGoodsNum 批量查找 数量
func (obj *_EsGoodsWordsMgr) GetBatchFromGoodsNum(goodsNums []int64) (results []*models.EsGoodsWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`goods_num` IN (?)", goodsNums).Find(&results).Error

	return
}

// GetFromQuanpin 通过quanpin获取内容 全拼字母
func (obj *_EsGoodsWordsMgr) GetFromQuanpin(quanpin string) (results []*models.EsGoodsWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`quanpin` = ?", quanpin).Find(&results).Error

	return
}

// GetBatchFromQuanpin 批量查找 全拼字母
func (obj *_EsGoodsWordsMgr) GetBatchFromQuanpin(quanpins []string) (results []*models.EsGoodsWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`quanpin` IN (?)", quanpins).Find(&results).Error

	return
}

// GetFromSzm 通过szm获取内容 首字母
func (obj *_EsGoodsWordsMgr) GetFromSzm(szm string) (results []*models.EsGoodsWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`szm` = ?", szm).Find(&results).Error

	return
}

// GetBatchFromSzm 批量查找 首字母
func (obj *_EsGoodsWordsMgr) GetBatchFromSzm(szms []string) (results []*models.EsGoodsWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`szm` IN (?)", szms).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型，系统(SYSTEM)，平台(PLATFORM)
func (obj *_EsGoodsWordsMgr) GetFromType(_type string) (results []*models.EsGoodsWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 类型，系统(SYSTEM)，平台(PLATFORM)
func (obj *_EsGoodsWordsMgr) GetBatchFromType(_types []string) (results []*models.EsGoodsWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *_EsGoodsWordsMgr) GetFromSort(sort int) (results []*models.EsGoodsWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *_EsGoodsWordsMgr) GetBatchFromSort(sorts []int) (results []*models.EsGoodsWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsGoodsWordsMgr) FetchByPrimaryKey(id int) (result models.EsGoodsWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`id` = ?", id).First(&result).Error

	return
}
