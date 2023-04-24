package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsSensitiveWordsMgr struct {
	*_BaseMgr
}

// EsSensitiveWordsMgr open func
func EsSensitiveWordsMgr(db *gorm.DB) *_EsSensitiveWordsMgr {
	if db == nil {
		panic(fmt.Errorf("EsSensitiveWordsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsSensitiveWordsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_sensitive_words"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsSensitiveWordsMgr) GetTableName() string {
	return "es_sensitive_words"
}

// Reset 重置gorm会话
func (obj *_EsSensitiveWordsMgr) Reset() *_EsSensitiveWordsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsSensitiveWordsMgr) Get() (result models.EsSensitiveWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSensitiveWords{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsSensitiveWordsMgr) Gets() (results []*models.EsSensitiveWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSensitiveWords{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsSensitiveWordsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsSensitiveWords{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_EsSensitiveWordsMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithWordName word_name获取 敏感词名称
func (obj *_EsSensitiveWordsMgr) WithWordName(wordName string) Option {
	return optionFunc(func(o *options) { o.query["word_name"] = wordName })
}

// WithCreateTime create_time获取 创建时间
func (obj *_EsSensitiveWordsMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithIsDelete is_delete获取 删除状态  1正常 0 删除
func (obj *_EsSensitiveWordsMgr) WithIsDelete(isDelete int16) Option {
	return optionFunc(func(o *options) { o.query["is_delete"] = isDelete })
}

// GetByOption 功能选项模式获取
func (obj *_EsSensitiveWordsMgr) GetByOption(opts ...Option) (result models.EsSensitiveWords, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSensitiveWords{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsSensitiveWordsMgr) GetByOptions(opts ...Option) (results []*models.EsSensitiveWords, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSensitiveWords{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsSensitiveWordsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSensitiveWords, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsSensitiveWords{}).Where(options.query)
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
func (obj *_EsSensitiveWordsMgr) GetFromID(id int) (result models.EsSensitiveWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSensitiveWords{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_EsSensitiveWordsMgr) GetBatchFromID(ids []int) (results []*models.EsSensitiveWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSensitiveWords{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromWordName 通过word_name获取内容 敏感词名称
func (obj *_EsSensitiveWordsMgr) GetFromWordName(wordName string) (results []*models.EsSensitiveWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSensitiveWords{}).Where("`word_name` = ?", wordName).Find(&results).Error

	return
}

// GetBatchFromWordName 批量查找 敏感词名称
func (obj *_EsSensitiveWordsMgr) GetBatchFromWordName(wordNames []string) (results []*models.EsSensitiveWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSensitiveWords{}).Where("`word_name` IN (?)", wordNames).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *_EsSensitiveWordsMgr) GetFromCreateTime(createTime int64) (results []*models.EsSensitiveWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSensitiveWords{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *_EsSensitiveWordsMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsSensitiveWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSensitiveWords{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromIsDelete 通过is_delete获取内容 删除状态  1正常 0 删除
func (obj *_EsSensitiveWordsMgr) GetFromIsDelete(isDelete int16) (results []*models.EsSensitiveWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSensitiveWords{}).Where("`is_delete` = ?", isDelete).Find(&results).Error

	return
}

// GetBatchFromIsDelete 批量查找 删除状态  1正常 0 删除
func (obj *_EsSensitiveWordsMgr) GetBatchFromIsDelete(isDeletes []int16) (results []*models.EsSensitiveWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSensitiveWords{}).Where("`is_delete` IN (?)", isDeletes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsSensitiveWordsMgr) FetchByPrimaryKey(id int) (result models.EsSensitiveWords, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSensitiveWords{}).Where("`id` = ?", id).First(&result).Error

	return
}