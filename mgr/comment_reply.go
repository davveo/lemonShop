package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type CommentReplyMgr struct {
	*_BaseMgr
}

// NewCommentReplyMgr open func
func NewCommentReplyMgr(db db.Repo) *CommentReplyMgr {
	if db == nil {
		panic(fmt.Errorf("NewCommentReplyMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &CommentReplyMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_category_spec"),
		wdb:       db.GetDbW().Table("es_category_spec"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *CommentReplyMgr) GetTableName() string {
	return "es_comment_reply"
}

// Reset 重置gorm会话
func (obj *CommentReplyMgr) Reset() *CommentReplyMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *CommentReplyMgr) Get() (result models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *CommentReplyMgr) Gets() (results []*models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *CommentReplyMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithReplyID reply_id获取 主键
func (obj *CommentReplyMgr) WithReplyID(replyID int) Option {
	return optionFunc(func(o *options) { o.query["reply_id"] = replyID })
}

// WithCommentID comment_id获取 评论id
func (obj *CommentReplyMgr) WithCommentID(commentID int) Option {
	return optionFunc(func(o *options) { o.query["comment_id"] = commentID })
}

// WithParentID parent_id获取 回复父id
func (obj *CommentReplyMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithContent content获取 回复内容
func (obj *CommentReplyMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithCreateTime create_time获取 创建时间
func (obj *CommentReplyMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithRole role获取 商家或者买家
func (obj *CommentReplyMgr) WithRole(role string) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// WithPath path获取 父子路径0|10|
func (obj *CommentReplyMgr) WithPath(path string) Option {
	return optionFunc(func(o *options) { o.query["path"] = path })
}

// WithReplyType reply_type获取 INITIAL
func (obj *CommentReplyMgr) WithReplyType(replyType string) Option {
	return optionFunc(func(o *options) { o.query["reply_type"] = replyType })
}

// GetByOption 功能选项模式获取
func (obj *CommentReplyMgr) GetByOption(opts ...Option) (result models.EsCommentReply, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *CommentReplyMgr) GetByOptions(opts ...Option) (results []*models.EsCommentReply, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *CommentReplyMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsCommentReply, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where(options.query)
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

// GetFromReplyID 通过reply_id获取内容 主键
func (obj *CommentReplyMgr) GetFromReplyID(replyID int) (result models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where("`reply_id` = ?", replyID).First(&result).Error

	return
}

// GetBatchFromReplyID 批量查找 主键
func (obj *CommentReplyMgr) GetBatchFromReplyID(replyIDs []int) (results []*models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where("`reply_id` IN (?)", replyIDs).Find(&results).Error

	return
}

// GetFromCommentID 通过comment_id获取内容 评论id
func (obj *CommentReplyMgr) GetFromCommentID(commentID int) (results []*models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where("`comment_id` = ?", commentID).Find(&results).Error

	return
}

// GetBatchFromCommentID 批量查找 评论id
func (obj *CommentReplyMgr) GetBatchFromCommentID(commentIDs []int) (results []*models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where("`comment_id` IN (?)", commentIDs).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 回复父id
func (obj *CommentReplyMgr) GetFromParentID(parentID int) (results []*models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 回复父id
func (obj *CommentReplyMgr) GetBatchFromParentID(parentIDs []int) (results []*models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 回复内容
func (obj *CommentReplyMgr) GetFromContent(content string) (results []*models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 回复内容
func (obj *CommentReplyMgr) GetBatchFromContent(contents []string) (results []*models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间
func (obj *CommentReplyMgr) GetFromCreateTime(createTime int64) (results []*models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间
func (obj *CommentReplyMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容 商家或者买家
func (obj *CommentReplyMgr) GetFromRole(role string) (results []*models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找 商家或者买家
func (obj *CommentReplyMgr) GetBatchFromRole(roles []string) (results []*models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

// GetFromPath 通过path获取内容 父子路径0|10|
func (obj *CommentReplyMgr) GetFromPath(path string) (results []*models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where("`path` = ?", path).Find(&results).Error

	return
}

// GetBatchFromPath 批量查找 父子路径0|10|
func (obj *CommentReplyMgr) GetBatchFromPath(paths []string) (results []*models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where("`path` IN (?)", paths).Find(&results).Error

	return
}

// GetFromReplyType 通过reply_type获取内容 INITIAL
func (obj *CommentReplyMgr) GetFromReplyType(replyType string) (results []*models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where("`reply_type` = ?", replyType).Find(&results).Error

	return
}

// GetBatchFromReplyType 批量查找 INITIAL
func (obj *CommentReplyMgr) GetBatchFromReplyType(replyTypes []string) (results []*models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where("`reply_type` IN (?)", replyTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *CommentReplyMgr) FetchByPrimaryKey(replyID int) (result models.EsCommentReply, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsCommentReply{}).Where("`reply_id` = ?", replyID).First(&result).Error

	return
}
