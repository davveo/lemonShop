package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type GoodsWordsMgr struct {
	*_BaseMgr
}

// NewGoodsWordsMgr open func
func NewGoodsWordsMgr(db db.Repo) *GoodsWordsMgr {
	if db == nil {
		panic(fmt.Errorf("NewGoodsWordsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &GoodsWordsMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_goods_words"),
		wdb:       db.GetDbW().Table("es_goods_words"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *GoodsWordsMgr) GetTableName() string {
	return "es_goods_words"
}

// Reset 重置gorm会话
func (obj *GoodsWordsMgr) Reset() *GoodsWordsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *GoodsWordsMgr) Get() (result models.EsGoodsWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *GoodsWordsMgr) Gets() (results []*models.EsGoodsWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *GoodsWordsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *GoodsWordsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithWords words获取 分词
func (obj *GoodsWordsMgr) WithWords(words string) Option {
	return optionFunc(func(o *options) { o.query["words"] = words })
}

// WithGoodsNum goods_num获取 数量
func (obj *GoodsWordsMgr) WithGoodsNum(goodsNum int64) Option {
	return optionFunc(func(o *options) { o.query["goods_num"] = goodsNum })
}

// WithQuanpin quanpin获取 全拼字母
func (obj *GoodsWordsMgr) WithQuanpin(quanpin string) Option {
	return optionFunc(func(o *options) { o.query["quanpin"] = quanpin })
}

// WithSzm szm获取 首字母
func (obj *GoodsWordsMgr) WithSzm(szm string) Option {
	return optionFunc(func(o *options) { o.query["szm"] = szm })
}

// WithType type获取 类型，系统(SYSTEM)，平台(PLATFORM)
func (obj *GoodsWordsMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithSort sort获取 排序
func (obj *GoodsWordsMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// GetByOption 功能选项模式获取
func (obj *GoodsWordsMgr) GetByOption(opts ...Option) (result models.EsGoodsWords, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *GoodsWordsMgr) GetByOptions(opts ...Option) (results []*models.EsGoodsWords, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *GoodsWordsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsGoodsWords, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where(options.query)
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
func (obj *GoodsWordsMgr) GetFromID(id int) (result models.EsGoodsWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *GoodsWordsMgr) GetBatchFromID(ids []int) (results []*models.EsGoodsWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromWords 通过words获取内容 分词
func (obj *GoodsWordsMgr) GetFromWords(words string) (results []*models.EsGoodsWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`words` = ?", words).Find(&results).Error

	return
}

// GetBatchFromWords 批量查找 分词
func (obj *GoodsWordsMgr) GetBatchFromWords(wordss []string) (results []*models.EsGoodsWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`words` IN (?)", wordss).Find(&results).Error

	return
}

// GetFromGoodsNum 通过goods_num获取内容 数量
func (obj *GoodsWordsMgr) GetFromGoodsNum(goodsNum int64) (results []*models.EsGoodsWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`goods_num` = ?", goodsNum).Find(&results).Error

	return
}

// GetBatchFromGoodsNum 批量查找 数量
func (obj *GoodsWordsMgr) GetBatchFromGoodsNum(goodsNums []int64) (results []*models.EsGoodsWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`goods_num` IN (?)", goodsNums).Find(&results).Error

	return
}

// GetFromQuanpin 通过quanpin获取内容 全拼字母
func (obj *GoodsWordsMgr) GetFromQuanpin(quanpin string) (results []*models.EsGoodsWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`quanpin` = ?", quanpin).Find(&results).Error

	return
}

// GetBatchFromQuanpin 批量查找 全拼字母
func (obj *GoodsWordsMgr) GetBatchFromQuanpin(quanpins []string) (results []*models.EsGoodsWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`quanpin` IN (?)", quanpins).Find(&results).Error

	return
}

// GetFromSzm 通过szm获取内容 首字母
func (obj *GoodsWordsMgr) GetFromSzm(szm string) (results []*models.EsGoodsWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`szm` = ?", szm).Find(&results).Error

	return
}

// GetBatchFromSzm 批量查找 首字母
func (obj *GoodsWordsMgr) GetBatchFromSzm(szms []string) (results []*models.EsGoodsWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`szm` IN (?)", szms).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 类型，系统(SYSTEM)，平台(PLATFORM)
func (obj *GoodsWordsMgr) GetFromType(_type string) (results []*models.EsGoodsWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 类型，系统(SYSTEM)，平台(PLATFORM)
func (obj *GoodsWordsMgr) GetBatchFromType(_types []string) (results []*models.EsGoodsWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *GoodsWordsMgr) GetFromSort(sort int) (results []*models.EsGoodsWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *GoodsWordsMgr) GetBatchFromSort(sorts []int) (results []*models.EsGoodsWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *GoodsWordsMgr) FetchByPrimaryKey(id int) (result models.EsGoodsWords, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsWords{}).Where("`id` = ?", id).First(&result).Error

	return
}
