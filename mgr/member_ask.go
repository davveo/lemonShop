package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type MemberAskMgr struct {
	*_BaseMgr
}

// NewMemberAskMgr open func
func NewMemberAskMgr(db db.Repo) *MemberAskMgr {
	if db == nil {
		panic(fmt.Errorf("NewMemberAskMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &MemberAskMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_member_ask"), wdb: db.GetDbW().Table("es_member_ask"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *MemberAskMgr) GetTableName() string {
	return "es_member_ask"
}

// Reset 重置gorm会话
func (obj *MemberAskMgr) Reset() *MemberAskMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *MemberAskMgr) Get() (result models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *MemberAskMgr) Gets() (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *MemberAskMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithAskID ask_id获取 主键
func (obj *MemberAskMgr) WithAskID(askID int) Option {
	return optionFunc(func(o *options) { o.query["ask_id"] = askID })
}

// WithGoodsID goods_id获取 商品id
func (obj *MemberAskMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithMemberID member_id获取 会员id
func (obj *MemberAskMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithContent content获取 咨询内容
func (obj *MemberAskMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithCreateTime create_time获取 咨询时间
func (obj *MemberAskMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithSellerID seller_id获取 卖家id
func (obj *MemberAskMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithReply reply获取 卖家回复内容
func (obj *MemberAskMgr) WithReply(reply string) Option {
	return optionFunc(func(o *options) { o.query["reply"] = reply })
}

// WithReplyTime reply_time获取 卖家回复时间
func (obj *MemberAskMgr) WithReplyTime(replyTime int64) Option {
	return optionFunc(func(o *options) { o.query["reply_time"] = replyTime })
}

// WithReplyStatus reply_status获取 商家回复状态 YES:已回复,NO:未回复
func (obj *MemberAskMgr) WithReplyStatus(replyStatus string) Option {
	return optionFunc(func(o *options) { o.query["reply_status"] = replyStatus })
}

// WithStatus status获取 删除状态 DELETED:已删除,NORMAL:正常
func (obj *MemberAskMgr) WithStatus(status string) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithMemberName member_name获取 会员名称
func (obj *MemberAskMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithGoodsName goods_name获取 商品名称
func (obj *MemberAskMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithMemberFace member_face获取 会员头像
func (obj *MemberAskMgr) WithMemberFace(memberFace string) Option {
	return optionFunc(func(o *options) { o.query["member_face"] = memberFace })
}

// WithAuthStatus auth_status获取 审核状态 WAIT_AUDIT:待审核,PASS_AUDIT:审核通过,REFUSE_AUDIT:审核未通过
func (obj *MemberAskMgr) WithAuthStatus(authStatus string) Option {
	return optionFunc(func(o *options) { o.query["auth_status"] = authStatus })
}

// WithGoodsImg goods_img获取 商品图片
func (obj *MemberAskMgr) WithGoodsImg(goodsImg string) Option {
	return optionFunc(func(o *options) { o.query["goods_img"] = goodsImg })
}

// WithAnonymous anonymous获取 是否匿名 YES:是，NO:否
func (obj *MemberAskMgr) WithAnonymous(anonymous string) Option {
	return optionFunc(func(o *options) { o.query["anonymous"] = anonymous })
}

// WithReplyNum reply_num获取 会员问题咨询回复数量
func (obj *MemberAskMgr) WithReplyNum(replyNum int) Option {
	return optionFunc(func(o *options) { o.query["reply_num"] = replyNum })
}

// GetByOption 功能选项模式获取
func (obj *MemberAskMgr) GetByOption(opts ...Option) (result models.EsMemberAsk, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *MemberAskMgr) GetByOptions(opts ...Option) (results []*models.EsMemberAsk, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *MemberAskMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsMemberAsk, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where(options.query)
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

// GetFromAskID 通过ask_id获取内容 主键
func (obj *MemberAskMgr) GetFromAskID(askID int) (result models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`ask_id` = ?", askID).First(&result).Error

	return
}

// GetBatchFromAskID 批量查找 主键
func (obj *MemberAskMgr) GetBatchFromAskID(askIDs []int) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`ask_id` IN (?)", askIDs).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品id
func (obj *MemberAskMgr) GetFromGoodsID(goodsID int) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品id
func (obj *MemberAskMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *MemberAskMgr) GetFromMemberID(memberID int) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *MemberAskMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 咨询内容
func (obj *MemberAskMgr) GetFromContent(content string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 咨询内容
func (obj *MemberAskMgr) GetBatchFromContent(contents []string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 咨询时间
func (obj *MemberAskMgr) GetFromCreateTime(createTime int64) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 咨询时间
func (obj *MemberAskMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 卖家id
func (obj *MemberAskMgr) GetFromSellerID(sellerID int) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 卖家id
func (obj *MemberAskMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromReply 通过reply获取内容 卖家回复内容
func (obj *MemberAskMgr) GetFromReply(reply string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`reply` = ?", reply).Find(&results).Error

	return
}

// GetBatchFromReply 批量查找 卖家回复内容
func (obj *MemberAskMgr) GetBatchFromReply(replys []string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`reply` IN (?)", replys).Find(&results).Error

	return
}

// GetFromReplyTime 通过reply_time获取内容 卖家回复时间
func (obj *MemberAskMgr) GetFromReplyTime(replyTime int64) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`reply_time` = ?", replyTime).Find(&results).Error

	return
}

// GetBatchFromReplyTime 批量查找 卖家回复时间
func (obj *MemberAskMgr) GetBatchFromReplyTime(replyTimes []int64) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`reply_time` IN (?)", replyTimes).Find(&results).Error

	return
}

// GetFromReplyStatus 通过reply_status获取内容 商家回复状态 YES:已回复,NO:未回复
func (obj *MemberAskMgr) GetFromReplyStatus(replyStatus string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`reply_status` = ?", replyStatus).Find(&results).Error

	return
}

// GetBatchFromReplyStatus 批量查找 商家回复状态 YES:已回复,NO:未回复
func (obj *MemberAskMgr) GetBatchFromReplyStatus(replyStatuss []string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`reply_status` IN (?)", replyStatuss).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 删除状态 DELETED:已删除,NORMAL:正常
func (obj *MemberAskMgr) GetFromStatus(status string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 删除状态 DELETED:已删除,NORMAL:正常
func (obj *MemberAskMgr) GetBatchFromStatus(statuss []string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 会员名称
func (obj *MemberAskMgr) GetFromMemberName(memberName string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 会员名称
func (obj *MemberAskMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *MemberAskMgr) GetFromGoodsName(goodsName string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *MemberAskMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromMemberFace 通过member_face获取内容 会员头像
func (obj *MemberAskMgr) GetFromMemberFace(memberFace string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`member_face` = ?", memberFace).Find(&results).Error

	return
}

// GetBatchFromMemberFace 批量查找 会员头像
func (obj *MemberAskMgr) GetBatchFromMemberFace(memberFaces []string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`member_face` IN (?)", memberFaces).Find(&results).Error

	return
}

// GetFromAuthStatus 通过auth_status获取内容 审核状态 WAIT_AUDIT:待审核,PASS_AUDIT:审核通过,REFUSE_AUDIT:审核未通过
func (obj *MemberAskMgr) GetFromAuthStatus(authStatus string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`auth_status` = ?", authStatus).Find(&results).Error

	return
}

// GetBatchFromAuthStatus 批量查找 审核状态 WAIT_AUDIT:待审核,PASS_AUDIT:审核通过,REFUSE_AUDIT:审核未通过
func (obj *MemberAskMgr) GetBatchFromAuthStatus(authStatuss []string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`auth_status` IN (?)", authStatuss).Find(&results).Error

	return
}

// GetFromGoodsImg 通过goods_img获取内容 商品图片
func (obj *MemberAskMgr) GetFromGoodsImg(goodsImg string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`goods_img` = ?", goodsImg).Find(&results).Error

	return
}

// GetBatchFromGoodsImg 批量查找 商品图片
func (obj *MemberAskMgr) GetBatchFromGoodsImg(goodsImgs []string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`goods_img` IN (?)", goodsImgs).Find(&results).Error

	return
}

// GetFromAnonymous 通过anonymous获取内容 是否匿名 YES:是，NO:否
func (obj *MemberAskMgr) GetFromAnonymous(anonymous string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`anonymous` = ?", anonymous).Find(&results).Error

	return
}

// GetBatchFromAnonymous 批量查找 是否匿名 YES:是，NO:否
func (obj *MemberAskMgr) GetBatchFromAnonymous(anonymouss []string) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`anonymous` IN (?)", anonymouss).Find(&results).Error

	return
}

// GetFromReplyNum 通过reply_num获取内容 会员问题咨询回复数量
func (obj *MemberAskMgr) GetFromReplyNum(replyNum int) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`reply_num` = ?", replyNum).Find(&results).Error

	return
}

// GetBatchFromReplyNum 批量查找 会员问题咨询回复数量
func (obj *MemberAskMgr) GetBatchFromReplyNum(replyNums []int) (results []*models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`reply_num` IN (?)", replyNums).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *MemberAskMgr) FetchByPrimaryKey(askID int) (result models.EsMemberAsk, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberAsk{}).Where("`ask_id` = ?", askID).First(&result).Error

	return
}
