package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsComplainTopicMgr struct {
	*_BaseMgr
}

// EsComplainTopicMgr open func
func EsComplainTopicMgr(db *gorm.DB) *_EsComplainTopicMgr {
	if db == nil {
		panic(fmt.Errorf("EsComplainTopicMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsComplainTopicMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_complain_topic"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsComplainTopicMgr) GetTableName() string {
	return "es_complain_topic"
}

// Reset 重置gorm会话
func (obj *_EsComplainTopicMgr) Reset() *_EsComplainTopicMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsComplainTopicMgr) Get() (result models.EsComplainTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsComplainTopic{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsComplainTopicMgr) Gets() (results []*models.EsComplainTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsComplainTopic{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsComplainTopicMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsComplainTopic{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithTopicID topic_id获取  主键
func (obj *_EsComplainTopicMgr) WithTopicID(topicID int) Option {
	return optionFunc(func(o *options) { o.query["topic_id"] = topicID })
}

// WithTopicName topic_name获取 投诉主题
func (obj *_EsComplainTopicMgr) WithTopicName(topicName string) Option {
	return optionFunc(func(o *options) { o.query["topic_name"] = topicName })
}

// WithCreateTime create_time获取 添加时间
func (obj *_EsComplainTopicMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithTopicRemark topic_remark获取 主题说明
func (obj *_EsComplainTopicMgr) WithTopicRemark(topicRemark string) Option {
	return optionFunc(func(o *options) { o.query["topic_remark"] = topicRemark })
}

// GetByOption 功能选项模式获取
func (obj *_EsComplainTopicMgr) GetByOption(opts ...Option) (result models.EsComplainTopic, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsComplainTopic{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsComplainTopicMgr) GetByOptions(opts ...Option) (results []*models.EsComplainTopic, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsComplainTopic{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsComplainTopicMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsComplainTopic, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsComplainTopic{}).Where(options.query)
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

// GetFromTopicID 通过topic_id获取内容  主键
func (obj *_EsComplainTopicMgr) GetFromTopicID(topicID int) (result models.EsComplainTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsComplainTopic{}).Where("`topic_id` = ?", topicID).First(&result).Error

	return
}

// GetBatchFromTopicID 批量查找  主键
func (obj *_EsComplainTopicMgr) GetBatchFromTopicID(topicIDs []int) (results []*models.EsComplainTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsComplainTopic{}).Where("`topic_id` IN (?)", topicIDs).Find(&results).Error

	return
}

// GetFromTopicName 通过topic_name获取内容 投诉主题
func (obj *_EsComplainTopicMgr) GetFromTopicName(topicName string) (results []*models.EsComplainTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsComplainTopic{}).Where("`topic_name` = ?", topicName).Find(&results).Error

	return
}

// GetBatchFromTopicName 批量查找 投诉主题
func (obj *_EsComplainTopicMgr) GetBatchFromTopicName(topicNames []string) (results []*models.EsComplainTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsComplainTopic{}).Where("`topic_name` IN (?)", topicNames).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 添加时间
func (obj *_EsComplainTopicMgr) GetFromCreateTime(createTime int64) (results []*models.EsComplainTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsComplainTopic{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 添加时间
func (obj *_EsComplainTopicMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsComplainTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsComplainTopic{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromTopicRemark 通过topic_remark获取内容 主题说明
func (obj *_EsComplainTopicMgr) GetFromTopicRemark(topicRemark string) (results []*models.EsComplainTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsComplainTopic{}).Where("`topic_remark` = ?", topicRemark).Find(&results).Error

	return
}

// GetBatchFromTopicRemark 批量查找 主题说明
func (obj *_EsComplainTopicMgr) GetBatchFromTopicRemark(topicRemarks []string) (results []*models.EsComplainTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsComplainTopic{}).Where("`topic_remark` IN (?)", topicRemarks).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsComplainTopicMgr) FetchByPrimaryKey(topicID int) (result models.EsComplainTopic, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsComplainTopic{}).Where("`topic_id` = ?", topicID).First(&result).Error

	return
}
