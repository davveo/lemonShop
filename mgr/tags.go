package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type TagsMgr struct {
	*_BaseMgr
}

// NewTagsMgr open func
func NewTagsMgr(db db.Repo) *TagsMgr {
	if db == nil {
		panic(fmt.Errorf("NewTagsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &TagsMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_tags"), wdb: db.GetDbW().Table("es_tags"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *TagsMgr) GetTableName() string {
	return "es_tags"
}

// Reset 重置gorm会话
func (obj *TagsMgr) Reset() *TagsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *TagsMgr) Get() (result models.EsTags, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTags{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *TagsMgr) Gets() (results []*models.EsTags, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTags{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *TagsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsTags{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTagID tag_id获取 主键
func (obj *TagsMgr) WithTagID(tagID int) Option {
	return optionFunc(func(o *options) { o.query["tag_id"] = tagID })
}

// WithTagName tag_name获取 标签名字
func (obj *TagsMgr) WithTagName(tagName string) Option {
	return optionFunc(func(o *options) { o.query["tag_name"] = tagName })
}

// WithSellerID seller_id获取 所属卖家
func (obj *TagsMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithMark mark获取 关键字
func (obj *TagsMgr) WithMark(mark string) Option {
	return optionFunc(func(o *options) { o.query["mark"] = mark })
}

// GetByOption 功能选项模式获取
func (obj *TagsMgr) GetByOption(opts ...Option) (result models.EsTags, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTags{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *TagsMgr) GetByOptions(opts ...Option) (results []*models.EsTags, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTags{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *TagsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsTags, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsTags{}).Where(options.query)
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

// GetFromTagID 通过tag_id获取内容 主键
func (obj *TagsMgr) GetFromTagID(tagID int) (result models.EsTags, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTags{}).Where("`tag_id` = ?", tagID).First(&result).Error

	return
}

// GetBatchFromTagID 批量查找 主键
func (obj *TagsMgr) GetBatchFromTagID(tagIDs []int) (results []*models.EsTags, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTags{}).Where("`tag_id` IN (?)", tagIDs).Find(&results).Error

	return
}

// GetFromTagName 通过tag_name获取内容 标签名字
func (obj *TagsMgr) GetFromTagName(tagName string) (results []*models.EsTags, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTags{}).Where("`tag_name` = ?", tagName).Find(&results).Error

	return
}

// GetBatchFromTagName 批量查找 标签名字
func (obj *TagsMgr) GetBatchFromTagName(tagNames []string) (results []*models.EsTags, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTags{}).Where("`tag_name` IN (?)", tagNames).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 所属卖家
func (obj *TagsMgr) GetFromSellerID(sellerID int) (results []*models.EsTags, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTags{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 所属卖家
func (obj *TagsMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsTags, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTags{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromMark 通过mark获取内容 关键字
func (obj *TagsMgr) GetFromMark(mark string) (results []*models.EsTags, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTags{}).Where("`mark` = ?", mark).Find(&results).Error

	return
}

// GetBatchFromMark 批量查找 关键字
func (obj *TagsMgr) GetBatchFromMark(marks []string) (results []*models.EsTags, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTags{}).Where("`mark` IN (?)", marks).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *TagsMgr) FetchByPrimaryKey(tagID int) (result models.EsTags, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsTags{}).Where("`tag_id` = ?", tagID).First(&result).Error

	return
}
