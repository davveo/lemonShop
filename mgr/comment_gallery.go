package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type CommentGalleryMgr struct {
	*_BaseMgr
}

// NewCommentGalleryMgr open func
func NewCommentGalleryMgr(db db.Repo) *CommentGalleryMgr {
	if db == nil {
		panic(fmt.Errorf("NewCommentGalleryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &CommentGalleryMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_comment_gallery"),
		wdb:       db.GetDbW().Table("es_comment_gallery"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *CommentGalleryMgr) GetTableName() string {
	return "es_comment_gallery"
}

// Reset 重置gorm会话
func (obj *CommentGalleryMgr) Reset() *CommentGalleryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *CommentGalleryMgr) Get() (result models.EsCommentGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentGallery{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *CommentGalleryMgr) Gets() (results []*models.EsCommentGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentGallery{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *CommentGalleryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsCommentGallery{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithImgID img_id获取 主键
func (obj *CommentGalleryMgr) WithImgID(imgID int) Option {
	return optionFunc(func(o *options) { o.query["img_id"] = imgID })
}

// WithCommentID comment_id获取 主键
func (obj *CommentGalleryMgr) WithCommentID(commentID int) Option {
	return optionFunc(func(o *options) { o.query["comment_id"] = commentID })
}

// WithOriginal original获取 图片路径
func (obj *CommentGalleryMgr) WithOriginal(original string) Option {
	return optionFunc(func(o *options) { o.query["original"] = original })
}

// WithSort sort获取 排序
func (obj *CommentGalleryMgr) WithSort(sort int) Option {
	return optionFunc(func(o *options) { o.query["sort"] = sort })
}

// WithImgBelong img_belong获取 图片所属 0：初评，1：追评
func (obj *CommentGalleryMgr) WithImgBelong(imgBelong int) Option {
	return optionFunc(func(o *options) { o.query["img_belong"] = imgBelong })
}

// GetByOption 功能选项模式获取
func (obj *CommentGalleryMgr) GetByOption(opts ...Option) (result models.EsCommentGallery, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentGallery{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *CommentGalleryMgr) GetByOptions(opts ...Option) (results []*models.EsCommentGallery, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentGallery{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *CommentGalleryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsCommentGallery, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsCommentGallery{}).Where(options.query)
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
func (obj *CommentGalleryMgr) GetFromImgID(imgID int) (result models.EsCommentGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentGallery{}).Where("`img_id` = ?", imgID).First(&result).Error

	return
}

// GetBatchFromImgID 批量查找 主键
func (obj *CommentGalleryMgr) GetBatchFromImgID(imgIDs []int) (results []*models.EsCommentGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentGallery{}).Where("`img_id` IN (?)", imgIDs).Find(&results).Error

	return
}

// GetFromCommentID 通过comment_id获取内容 主键
func (obj *CommentGalleryMgr) GetFromCommentID(commentID int) (results []*models.EsCommentGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentGallery{}).Where("`comment_id` = ?", commentID).Find(&results).Error

	return
}

// GetBatchFromCommentID 批量查找 主键
func (obj *CommentGalleryMgr) GetBatchFromCommentID(commentIDs []int) (results []*models.EsCommentGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentGallery{}).Where("`comment_id` IN (?)", commentIDs).Find(&results).Error

	return
}

// GetFromOriginal 通过original获取内容 图片路径
func (obj *CommentGalleryMgr) GetFromOriginal(original string) (results []*models.EsCommentGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentGallery{}).Where("`original` = ?", original).Find(&results).Error

	return
}

// GetBatchFromOriginal 批量查找 图片路径
func (obj *CommentGalleryMgr) GetBatchFromOriginal(originals []string) (results []*models.EsCommentGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentGallery{}).Where("`original` IN (?)", originals).Find(&results).Error

	return
}

// GetFromSort 通过sort获取内容 排序
func (obj *CommentGalleryMgr) GetFromSort(sort int) (results []*models.EsCommentGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentGallery{}).Where("`sort` = ?", sort).Find(&results).Error

	return
}

// GetBatchFromSort 批量查找 排序
func (obj *CommentGalleryMgr) GetBatchFromSort(sorts []int) (results []*models.EsCommentGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentGallery{}).Where("`sort` IN (?)", sorts).Find(&results).Error

	return
}

// GetFromImgBelong 通过img_belong获取内容 图片所属 0：初评，1：追评
func (obj *CommentGalleryMgr) GetFromImgBelong(imgBelong int) (results []*models.EsCommentGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentGallery{}).Where("`img_belong` = ?", imgBelong).Find(&results).Error

	return
}

// GetBatchFromImgBelong 批量查找 图片所属 0：初评，1：追评
func (obj *CommentGalleryMgr) GetBatchFromImgBelong(imgBelongs []int) (results []*models.EsCommentGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentGallery{}).Where("`img_belong` IN (?)", imgBelongs).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *CommentGalleryMgr) FetchByPrimaryKey(imgID int) (result models.EsCommentGallery, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentGallery{}).Where("`img_id` = ?", imgID).First(&result).Error

	return
}
