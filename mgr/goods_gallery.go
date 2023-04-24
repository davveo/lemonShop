package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type GoodsGalleryMgr struct {
	*_BaseMgr
}

// NewGoodsGalleryMgr open func
func NewGoodsGalleryMgr(db db.Repo) *GoodsGalleryMgr {
	if db == nil {
		panic(fmt.Errorf("NewGoodsGalleryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &GoodsGalleryMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_category_spec"),
		wdb:       db.GetDbW().Table("es_category_spec"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *GoodsGalleryMgr) GetTableName() string {
	return "es_goods_gallery"
}

// Reset 重置gorm会话
func (obj *GoodsGalleryMgr) Reset() *GoodsGalleryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *GoodsGalleryMgr) Get() (result models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *GoodsGalleryMgr) Gets() (results []*models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *GoodsGalleryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithImgID img_id获取 主键
func (obj *GoodsGalleryMgr) WithImgID(imgID int) Option {
	return optionFunc(func(o *options) { o.query["img_id"] = imgID })
}

// WithGoodsID goods_id获取 商品主键
func (obj *GoodsGalleryMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithThumbnail thumbnail获取 缩略图路径
func (obj *GoodsGalleryMgr) WithThumbnail(thumbnail string) Option {
	return optionFunc(func(o *options) { o.query["thumbnail"] = thumbnail })
}

// WithSmall small获取 小图路径
func (obj *GoodsGalleryMgr) WithSmall(small string) Option {
	return optionFunc(func(o *options) { o.query["small"] = small })
}

// WithBig big获取 大图路径
func (obj *GoodsGalleryMgr) WithBig(big string) Option {
	return optionFunc(func(o *options) { o.query["big"] = big })
}

// WithOriginal original获取 原图路径
func (obj *GoodsGalleryMgr) WithOriginal(original string) Option {
	return optionFunc(func(o *options) { o.query["original"] = original })
}

// WithTiny tiny获取 极小图路径
func (obj *GoodsGalleryMgr) WithTiny(tiny string) Option {
	return optionFunc(func(o *options) { o.query["tiny"] = tiny })
}

// WithIsdefault isdefault获取 是否是默认图片1   0没有默认
func (obj *GoodsGalleryMgr) WithIsdefault(isdefault int) Option {
	return optionFunc(func(o *options) { o.query["isdefault"] = isdefault })
}

// WithSort sort获取 排序
func (obj *GoodsGalleryMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// GetByOption 功能选项模式获取
func (obj *GoodsGalleryMgr) GetByOption(opts ...Option) (result models.EsGoodsGallery, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *GoodsGalleryMgr) GetByOptions(opts ...Option) (results []*models.EsGoodsGallery, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *GoodsGalleryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsGoodsGallery, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where(options.query)
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

// GetFromImgID 通过img_id获取内容 主键
func (obj *GoodsGalleryMgr) GetFromImgID(imgID int) (result models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`img_id` = ?", imgID).First(&result).Error

	return
}

// GetBatchFromImgID 批量查找 主键
func (obj *GoodsGalleryMgr) GetBatchFromImgID(imgIDs []int) (results []*models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`img_id` IN (?)", imgIDs).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品主键
func (obj *GoodsGalleryMgr) GetFromGoodsID(goodsID int) (results []*models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品主键
func (obj *GoodsGalleryMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromThumbnail 通过thumbnail获取内容 缩略图路径
func (obj *GoodsGalleryMgr) GetFromThumbnail(thumbnail string) (results []*models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`thumbnail` = ?", thumbnail).Find(&results).Error

	return
}

// GetBatchFromThumbnail 批量查找 缩略图路径
func (obj *GoodsGalleryMgr) GetBatchFromThumbnail(thumbnails []string) (results []*models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`thumbnail` IN (?)", thumbnails).Find(&results).Error

	return
}

// GetFromSmall 通过small获取内容 小图路径
func (obj *GoodsGalleryMgr) GetFromSmall(small string) (results []*models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`small` = ?", small).Find(&results).Error

	return
}

// GetBatchFromSmall 批量查找 小图路径
func (obj *GoodsGalleryMgr) GetBatchFromSmall(smalls []string) (results []*models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`small` IN (?)", smalls).Find(&results).Error

	return
}

// GetFromBig 通过big获取内容 大图路径
func (obj *GoodsGalleryMgr) GetFromBig(big string) (results []*models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`big` = ?", big).Find(&results).Error

	return
}

// GetBatchFromBig 批量查找 大图路径
func (obj *GoodsGalleryMgr) GetBatchFromBig(bigs []string) (results []*models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`big` IN (?)", bigs).Find(&results).Error

	return
}

// GetFromOriginal 通过original获取内容 原图路径
func (obj *GoodsGalleryMgr) GetFromOriginal(original string) (results []*models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`original` = ?", original).Find(&results).Error

	return
}

// GetBatchFromOriginal 批量查找 原图路径
func (obj *GoodsGalleryMgr) GetBatchFromOriginal(originals []string) (results []*models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`original` IN (?)", originals).Find(&results).Error

	return
}

// GetFromTiny 通过tiny获取内容 极小图路径
func (obj *GoodsGalleryMgr) GetFromTiny(tiny string) (results []*models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`tiny` = ?", tiny).Find(&results).Error

	return
}

// GetBatchFromTiny 批量查找 极小图路径
func (obj *GoodsGalleryMgr) GetBatchFromTiny(tinys []string) (results []*models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`tiny` IN (?)", tinys).Find(&results).Error

	return
}

// GetFromIsdefault 通过isdefault获取内容 是否是默认图片1   0没有默认
func (obj *GoodsGalleryMgr) GetFromIsdefault(isdefault int) (results []*models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`isdefault` = ?", isdefault).Find(&results).Error

	return
}

// GetBatchFromIsdefault 批量查找 是否是默认图片1   0没有默认
func (obj *GoodsGalleryMgr) GetBatchFromIsdefault(isdefaults []int) (results []*models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`isdefault` IN (?)", isdefaults).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *GoodsGalleryMgr) GetFromSort(sort int) (results []*models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *GoodsGalleryMgr) GetBatchFromSort(sorts []int) (results []*models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *GoodsGalleryMgr) FetchByPrimaryKey(imgID int) (result models.EsGoodsGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsGoodsGallery{}).Where("`img_id` = ?", imgID).First(&result).Error

	return
}
