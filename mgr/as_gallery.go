package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"
	"gorm.io/gorm"
)

type AsGalleryMgr struct {
	*_BaseMgr
}

// NewAsGalleryMgr open func
func NewAsGalleryMgr(db db.Repo) *AsGalleryMgr {
	if db == nil {
		panic(fmt.Errorf("NewAsGalleryMgr need init by rdb"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &AsGalleryMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_as_gallery"),
		wdb:       db.GetDbW().Table("es_as_gallery"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *AsGalleryMgr) GetTableName() string {
	return "es_as_gallery"
}

// Reset 重置gorm会话
func (obj *AsGalleryMgr) Reset() *AsGalleryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *AsGalleryMgr) Get() (result models.EsAsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGallery{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *AsGalleryMgr) Gets() (results []*models.EsAsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGallery{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *AsGalleryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsAsGallery{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *AsGalleryMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithServiceSn service_sn获取 售后单号
func (obj *AsGalleryMgr) WithServiceSn(serviceSn string) Option {
	return optionFunc(func(o *options) { o.query["service_sn"] = serviceSn })
}

// WithImg img获取 图片链接
func (obj *AsGalleryMgr) WithImg(img string) Option {
	return optionFunc(func(o *options) { o.query["img"] = img })
}

// GetByOption 功能选项模式获取
func (obj *AsGalleryMgr) GetByOption(opts ...Option) (result models.EsAsGallery, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGallery{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *AsGalleryMgr) GetByOptions(opts ...Option) (results []*models.EsAsGallery, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGallery{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *AsGalleryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsAsGallery, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsAsGallery{}).Where(options.query)
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
func (obj *AsGalleryMgr) GetFromID(id int) (result models.EsAsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGallery{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *AsGalleryMgr) GetBatchFromID(ids []int) (results []*models.EsAsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGallery{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromServiceSn 通过service_sn获取内容 售后单号
func (obj *AsGalleryMgr) GetFromServiceSn(serviceSn string) (results []*models.EsAsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGallery{}).Where("`service_sn` = ?", serviceSn).Find(&results).Error

	return
}

// GetBatchFromServiceSn 批量查找 售后单号
func (obj *AsGalleryMgr) GetBatchFromServiceSn(serviceSns []string) (results []*models.EsAsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGallery{}).Where("`service_sn` IN (?)", serviceSns).Find(&results).Error

	return
}

// GetFromImg 通过img获取内容 图片链接
func (obj *AsGalleryMgr) GetFromImg(img string) (results []*models.EsAsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGallery{}).Where("`img` = ?", img).Find(&results).Error

	return
}

// GetBatchFromImg 批量查找 图片链接
func (obj *AsGalleryMgr) GetBatchFromImg(imgs []string) (results []*models.EsAsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGallery{}).Where("`img` IN (?)", imgs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *AsGalleryMgr) FetchByPrimaryKey(id int) (result models.EsAsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGallery{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByEsIndexID  获取多个内容
func (obj *AsGalleryMgr) FetchIndexByEsIndexID(id int) (results []*models.EsAsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGallery{}).Where("`id` = ?", id).Find(&results).Error

	return
}

// FetchIndexByEsIndexServiceSn  获取多个内容
func (obj *AsGalleryMgr) FetchIndexByEsIndexServiceSn(serviceSn string) (results []*models.EsAsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAsGallery{}).Where("`service_sn` = ?", serviceSn).Find(&results).Error

	return
}
