package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsAskMessageMgr struct {
	*_BaseMgr
}

// EsAskMessageMgr open func
func EsAskMessageMgr(db *gorm.DB) *_EsAskMessageMgr {
	if db == nil {
		panic(fmt.Errorf("EsAskMessageMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsAskMessageMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_ask_message"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsAskMessageMgr) GetTableName() string {
	return "es_ask_message"
}

// Reset 重置gorm会话
func (obj *_EsAskMessageMgr) Reset() *_EsAskMessageMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsAskMessageMgr) Get() (result models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsAskMessageMgr) Gets() (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsAskMessageMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *_EsAskMessageMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员ID
func (obj *_EsAskMessageMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithGoodsID goods_id获取 商品ID
func (obj *_EsAskMessageMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithGoodsName goods_name获取 商品名称
func (obj *_EsAskMessageMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithGoodsImg goods_img获取 商品图片
func (obj *_EsAskMessageMgr) WithGoodsImg(goodsImg string) Option {
	return optionFunc(func(o *options) { o.query["goods_img"] = goodsImg })
}

// WithAskID ask_id获取 会员咨询id
func (obj *_EsAskMessageMgr) WithAskID(askID int) Option {
	return optionFunc(func(o *options) { o.query["ask_id"] = askID })
}

// WithAsk ask获取 咨询内容
func (obj *_EsAskMessageMgr) WithAsk(ask string) Option {
	return optionFunc(func(o *options) { o.query["ask"] = ask })
}

// WithAskMember ask_member获取 咨询会员
func (obj *_EsAskMessageMgr) WithAskMember(askMember string) Option {
	return optionFunc(func(o *options) { o.query["ask_member"] = askMember })
}

// WithReplyID reply_id获取 会员回复咨询id
func (obj *_EsAskMessageMgr) WithReplyID(replyID int) Option {
	return optionFunc(func(o *options) { o.query["reply_id"] = replyID })
}

// WithReply reply获取 回复内容
func (obj *_EsAskMessageMgr) WithReply(reply string) Option {
	return optionFunc(func(o *options) { o.query["reply"] = reply })
}

// WithReplyMember reply_member获取 回复会员
func (obj *_EsAskMessageMgr) WithReplyMember(replyMember string) Option {
	return optionFunc(func(o *options) { o.query["reply_member"] = replyMember })
}

// WithSendTime send_time获取 消息发送时间
func (obj *_EsAskMessageMgr) WithSendTime(sendTime int64) Option {
	return optionFunc(func(o *options) { o.query["send_time"] = sendTime })
}

// WithIsDel is_del获取 删除状态 DELETED：已删除 NORMAL：正常
func (obj *_EsAskMessageMgr) WithIsDel(isDel string) Option {
	return optionFunc(func(o *options) { o.query["is_del"] = isDel })
}

// WithIsRead is_read获取 是否已读 YES：是 NO：否
func (obj *_EsAskMessageMgr) WithIsRead(isRead string) Option {
	return optionFunc(func(o *options) { o.query["is_read"] = isRead })
}

// WithReceiveTime receive_time获取 消息接收时间
func (obj *_EsAskMessageMgr) WithReceiveTime(receiveTime int64) Option {
	return optionFunc(func(o *options) { o.query["receive_time"] = receiveTime })
}

// WithMsgType msg_type获取 咨询消息类型 ASK：提问消息 REPLY：回复消息
func (obj *_EsAskMessageMgr) WithMsgType(msgType string) Option {
	return optionFunc(func(o *options) { o.query["msg_type"] = msgType })
}

// WithAskAnonymous ask_anonymous获取 咨询人是否匿名 YES:是，NO:否
func (obj *_EsAskMessageMgr) WithAskAnonymous(askAnonymous string) Option {
	return optionFunc(func(o *options) { o.query["ask_anonymous"] = askAnonymous })
}

// WithReplyAnonymous reply_anonymous获取 回复咨询人是否匿名 YES:是，NO:否
func (obj *_EsAskMessageMgr) WithReplyAnonymous(replyAnonymous string) Option {
	return optionFunc(func(o *options) { o.query["reply_anonymous"] = replyAnonymous })
}

// GetByOption 功能选项模式获取
func (obj *_EsAskMessageMgr) GetByOption(opts ...Option) (result models.EsAskMessage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsAskMessageMgr) GetByOptions(opts ...Option) (results []*models.EsAskMessage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsAskMessageMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsAskMessage, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where(options.query)
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
func (obj *_EsAskMessageMgr) GetFromID(id int) (result models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *_EsAskMessageMgr) GetBatchFromID(ids []int) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员ID
func (obj *_EsAskMessageMgr) GetFromMemberID(memberID int) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员ID
func (obj *_EsAskMessageMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品ID
func (obj *_EsAskMessageMgr) GetFromGoodsID(goodsID int) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品ID
func (obj *_EsAskMessageMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *_EsAskMessageMgr) GetFromGoodsName(goodsName string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *_EsAskMessageMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromGoodsImg 通过goods_img获取内容 商品图片
func (obj *_EsAskMessageMgr) GetFromGoodsImg(goodsImg string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`goods_img` = ?", goodsImg).Find(&results).Error

	return
}

// GetBatchFromGoodsImg 批量查找 商品图片
func (obj *_EsAskMessageMgr) GetBatchFromGoodsImg(goodsImgs []string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`goods_img` IN (?)", goodsImgs).Find(&results).Error

	return
}

// GetFromAskID 通过ask_id获取内容 会员咨询id
func (obj *_EsAskMessageMgr) GetFromAskID(askID int) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`ask_id` = ?", askID).Find(&results).Error

	return
}

// GetBatchFromAskID 批量查找 会员咨询id
func (obj *_EsAskMessageMgr) GetBatchFromAskID(askIDs []int) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`ask_id` IN (?)", askIDs).Find(&results).Error

	return
}

// GetFromAsk 通过ask获取内容 咨询内容
func (obj *_EsAskMessageMgr) GetFromAsk(ask string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`ask` = ?", ask).Find(&results).Error

	return
}

// GetBatchFromAsk 批量查找 咨询内容
func (obj *_EsAskMessageMgr) GetBatchFromAsk(asks []string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`ask` IN (?)", asks).Find(&results).Error

	return
}

// GetFromAskMember 通过ask_member获取内容 咨询会员
func (obj *_EsAskMessageMgr) GetFromAskMember(askMember string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`ask_member` = ?", askMember).Find(&results).Error

	return
}

// GetBatchFromAskMember 批量查找 咨询会员
func (obj *_EsAskMessageMgr) GetBatchFromAskMember(askMembers []string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`ask_member` IN (?)", askMembers).Find(&results).Error

	return
}

// GetFromReplyID 通过reply_id获取内容 会员回复咨询id
func (obj *_EsAskMessageMgr) GetFromReplyID(replyID int) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`reply_id` = ?", replyID).Find(&results).Error

	return
}

// GetBatchFromReplyID 批量查找 会员回复咨询id
func (obj *_EsAskMessageMgr) GetBatchFromReplyID(replyIDs []int) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`reply_id` IN (?)", replyIDs).Find(&results).Error

	return
}

// GetFromReply 通过reply获取内容 回复内容
func (obj *_EsAskMessageMgr) GetFromReply(reply string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`reply` = ?", reply).Find(&results).Error

	return
}

// GetBatchFromReply 批量查找 回复内容
func (obj *_EsAskMessageMgr) GetBatchFromReply(replys []string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`reply` IN (?)", replys).Find(&results).Error

	return
}

// GetFromReplyMember 通过reply_member获取内容 回复会员
func (obj *_EsAskMessageMgr) GetFromReplyMember(replyMember string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`reply_member` = ?", replyMember).Find(&results).Error

	return
}

// GetBatchFromReplyMember 批量查找 回复会员
func (obj *_EsAskMessageMgr) GetBatchFromReplyMember(replyMembers []string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`reply_member` IN (?)", replyMembers).Find(&results).Error

	return
}

// GetFromSendTime 通过send_time获取内容 消息发送时间
func (obj *_EsAskMessageMgr) GetFromSendTime(sendTime int64) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`send_time` = ?", sendTime).Find(&results).Error

	return
}

// GetBatchFromSendTime 批量查找 消息发送时间
func (obj *_EsAskMessageMgr) GetBatchFromSendTime(sendTimes []int64) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`send_time` IN (?)", sendTimes).Find(&results).Error

	return
}

// GetFromIsDel 通过is_del获取内容 删除状态 DELETED：已删除 NORMAL：正常
func (obj *_EsAskMessageMgr) GetFromIsDel(isDel string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`is_del` = ?", isDel).Find(&results).Error

	return
}

// GetBatchFromIsDel 批量查找 删除状态 DELETED：已删除 NORMAL：正常
func (obj *_EsAskMessageMgr) GetBatchFromIsDel(isDels []string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`is_del` IN (?)", isDels).Find(&results).Error

	return
}

// GetFromIsRead 通过is_read获取内容 是否已读 YES：是 NO：否
func (obj *_EsAskMessageMgr) GetFromIsRead(isRead string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`is_read` = ?", isRead).Find(&results).Error

	return
}

// GetBatchFromIsRead 批量查找 是否已读 YES：是 NO：否
func (obj *_EsAskMessageMgr) GetBatchFromIsRead(isReads []string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`is_read` IN (?)", isReads).Find(&results).Error

	return
}

// GetFromReceiveTime 通过receive_time获取内容 消息接收时间
func (obj *_EsAskMessageMgr) GetFromReceiveTime(receiveTime int64) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`receive_time` = ?", receiveTime).Find(&results).Error

	return
}

// GetBatchFromReceiveTime 批量查找 消息接收时间
func (obj *_EsAskMessageMgr) GetBatchFromReceiveTime(receiveTimes []int64) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`receive_time` IN (?)", receiveTimes).Find(&results).Error

	return
}

// GetFromMsgType 通过msg_type获取内容 咨询消息类型 ASK：提问消息 REPLY：回复消息
func (obj *_EsAskMessageMgr) GetFromMsgType(msgType string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`msg_type` = ?", msgType).Find(&results).Error

	return
}

// GetBatchFromMsgType 批量查找 咨询消息类型 ASK：提问消息 REPLY：回复消息
func (obj *_EsAskMessageMgr) GetBatchFromMsgType(msgTypes []string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`msg_type` IN (?)", msgTypes).Find(&results).Error

	return
}

// GetFromAskAnonymous 通过ask_anonymous获取内容 咨询人是否匿名 YES:是，NO:否
func (obj *_EsAskMessageMgr) GetFromAskAnonymous(askAnonymous string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`ask_anonymous` = ?", askAnonymous).Find(&results).Error

	return
}

// GetBatchFromAskAnonymous 批量查找 咨询人是否匿名 YES:是，NO:否
func (obj *_EsAskMessageMgr) GetBatchFromAskAnonymous(askAnonymouss []string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`ask_anonymous` IN (?)", askAnonymouss).Find(&results).Error

	return
}

// GetFromReplyAnonymous 通过reply_anonymous获取内容 回复咨询人是否匿名 YES:是，NO:否
func (obj *_EsAskMessageMgr) GetFromReplyAnonymous(replyAnonymous string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`reply_anonymous` = ?", replyAnonymous).Find(&results).Error

	return
}

// GetBatchFromReplyAnonymous 批量查找 回复咨询人是否匿名 YES:是，NO:否
func (obj *_EsAskMessageMgr) GetBatchFromReplyAnonymous(replyAnonymouss []string) (results []*models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`reply_anonymous` IN (?)", replyAnonymouss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsAskMessageMgr) FetchByPrimaryKey(id int) (result models.EsAskMessage, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsAskMessage{}).Where("`id` = ?", id).First(&result).Error

	return
}
