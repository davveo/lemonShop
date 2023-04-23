package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsMemberCommentMgr struct {
	*_BaseMgr
}

// EsMemberCommentMgr open func
func EsMemberCommentMgr(db *gorm.DB) *_EsMemberCommentMgr {
	if db == nil {
		panic(fmt.Errorf("EsMemberCommentMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsMemberCommentMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_member_comment"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsMemberCommentMgr) GetTableName() string {
	return "es_member_comment"
}

// Reset 重置gorm会话
func (obj *_EsMemberCommentMgr) Reset() *_EsMemberCommentMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsMemberCommentMgr) Get() (result models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsMemberCommentMgr) Gets() (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsMemberCommentMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithCommentID comment_id获取 评论主键
func (obj *_EsMemberCommentMgr) WithCommentID(commentID int) Option {
	return optionFunc(func(o *options) { o.query["comment_id"] = commentID })
}

// WithGoodsID goods_id获取 商品id
func (obj *_EsMemberCommentMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithSkuID sku_id获取 skuid
func (obj *_EsMemberCommentMgr) WithSkuID(skuID int) Option {
	return optionFunc(func(o *options) { o.query["sku_id"] = skuID })
}

// WithMemberID member_id获取 会员id
func (obj *_EsMemberCommentMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithSellerID seller_id获取 卖家id
func (obj *_EsMemberCommentMgr) WithSellerID(sellerID int) Option {
	return optionFunc(func(o *options) { o.query["seller_id"] = sellerID })
}

// WithMemberName member_name获取 会员名称
func (obj *_EsMemberCommentMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithMemberFace member_face获取 会员头像
func (obj *_EsMemberCommentMgr) WithMemberFace(memberFace string) Option {
	return optionFunc(func(o *options) { o.query["member_face"] = memberFace })
}

// WithGoodsName goods_name获取 商品名称
func (obj *_EsMemberCommentMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithGoodsImg goods_img获取 商品图片
func (obj *_EsMemberCommentMgr) WithGoodsImg(goodsImg string) Option {
	return optionFunc(func(o *options) { o.query["goods_img"] = goodsImg })
}

// WithContent content获取 评论内容
func (obj *_EsMemberCommentMgr) WithContent(content string) Option {
	return optionFunc(func(o *options) { o.query["content"] = content })
}

// WithCreateTime create_time获取 评论时间
func (obj *_EsMemberCommentMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithHaveImage have_image获取 是否有图片 1 有 0 没有
func (obj *_EsMemberCommentMgr) WithHaveImage(haveImage int16) Option {
	return optionFunc(func(o *options) { o.query["have_image"] = haveImage })
}

// WithStatus status获取 状态  1 正常 0 删除
func (obj *_EsMemberCommentMgr) WithStatus(status int16) Option {
	return optionFunc(func(o *options) { o.query["status"] = status })
}

// WithGrade grade获取 好中差评
func (obj *_EsMemberCommentMgr) WithGrade(grade string) Option {
	return optionFunc(func(o *options) { o.query["grade"] = grade })
}

// WithOrderSn order_sn获取 订单编号
func (obj *_EsMemberCommentMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithReplyStatus reply_status获取 是否回复 1 是 0 否
func (obj *_EsMemberCommentMgr) WithReplyStatus(replyStatus int16) Option {
	return optionFunc(func(o *options) { o.query["reply_status"] = replyStatus })
}

// WithAuditStatus audit_status获取 初评审核:待审核(WAIT_AUDIT),审核通过(PASS_AUDIT),审核拒绝(REFUSE_AUDIT)
func (obj *_EsMemberCommentMgr) WithAuditStatus(auditStatus string) Option {
	return optionFunc(func(o *options) { o.query["audit_status"] = auditStatus })
}

// WithCommentsType comments_type获取 评论类型：初评(INITIAL),追评(ADDITIONAL)
func (obj *_EsMemberCommentMgr) WithCommentsType(commentsType string) Option {
	return optionFunc(func(o *options) { o.query["comments_type"] = commentsType })
}

// WithParentID parent_id获取 初评id，默认为0
func (obj *_EsMemberCommentMgr) WithParentID(parentID int) Option {
	return optionFunc(func(o *options) { o.query["parent_id"] = parentID })
}

// WithAdditionalStatus additional_status获取 追加评论状态 0：未追加，1：已追加
func (obj *_EsMemberCommentMgr) WithAdditionalStatus(additionalStatus int) Option {
	return optionFunc(func(o *options) { o.query["additional_status"] = additionalStatus })
}

// WithAdditionalContent additional_content获取 追加的评论内容
func (obj *_EsMemberCommentMgr) WithAdditionalContent(additionalContent string) Option {
	return optionFunc(func(o *options) { o.query["additional_content"] = additionalContent })
}

// WithAdditionalTime additional_time获取 追加评论时间
func (obj *_EsMemberCommentMgr) WithAdditionalTime(additionalTime int64) Option {
	return optionFunc(func(o *options) { o.query["additional_time"] = additionalTime })
}

// WithAdditionalHaveImage additional_have_image获取 追加的评论是否上传了图片 0：未上传，1：已上传
func (obj *_EsMemberCommentMgr) WithAdditionalHaveImage(additionalHaveImage int) Option {
	return optionFunc(func(o *options) { o.query["additional_have_image"] = additionalHaveImage })
}

// GetByOption 功能选项模式获取
func (obj *_EsMemberCommentMgr) GetByOption(opts ...Option) (result models.EsMemberComment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsMemberCommentMgr) GetByOptions(opts ...Option) (results []*models.EsMemberComment, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsMemberCommentMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsMemberComment, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where(options.query)
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

// GetFromCommentID 通过comment_id获取内容 评论主键
func (obj *_EsMemberCommentMgr) GetFromCommentID(commentID int) (result models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`comment_id` = ?", commentID).First(&result).Error

	return
}

// GetBatchFromCommentID 批量查找 评论主键
func (obj *_EsMemberCommentMgr) GetBatchFromCommentID(commentIDs []int) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`comment_id` IN (?)", commentIDs).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品id
func (obj *_EsMemberCommentMgr) GetFromGoodsID(goodsID int) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品id
func (obj *_EsMemberCommentMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromSkuID 通过sku_id获取内容 skuid
func (obj *_EsMemberCommentMgr) GetFromSkuID(skuID int) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`sku_id` = ?", skuID).Find(&results).Error

	return
}

// GetBatchFromSkuID 批量查找 skuid
func (obj *_EsMemberCommentMgr) GetBatchFromSkuID(skuIDs []int) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`sku_id` IN (?)", skuIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_EsMemberCommentMgr) GetFromMemberID(memberID int) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_EsMemberCommentMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromSellerID 通过seller_id获取内容 卖家id
func (obj *_EsMemberCommentMgr) GetFromSellerID(sellerID int) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`seller_id` = ?", sellerID).Find(&results).Error

	return
}

// GetBatchFromSellerID 批量查找 卖家id
func (obj *_EsMemberCommentMgr) GetBatchFromSellerID(sellerIDs []int) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`seller_id` IN (?)", sellerIDs).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 会员名称
func (obj *_EsMemberCommentMgr) GetFromMemberName(memberName string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 会员名称
func (obj *_EsMemberCommentMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromMemberFace 通过member_face获取内容 会员头像
func (obj *_EsMemberCommentMgr) GetFromMemberFace(memberFace string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`member_face` = ?", memberFace).Find(&results).Error

	return
}

// GetBatchFromMemberFace 批量查找 会员头像
func (obj *_EsMemberCommentMgr) GetBatchFromMemberFace(memberFaces []string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`member_face` IN (?)", memberFaces).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *_EsMemberCommentMgr) GetFromGoodsName(goodsName string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *_EsMemberCommentMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromGoodsImg 通过goods_img获取内容 商品图片
func (obj *_EsMemberCommentMgr) GetFromGoodsImg(goodsImg string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`goods_img` = ?", goodsImg).Find(&results).Error

	return
}

// GetBatchFromGoodsImg 批量查找 商品图片
func (obj *_EsMemberCommentMgr) GetBatchFromGoodsImg(goodsImgs []string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`goods_img` IN (?)", goodsImgs).Find(&results).Error

	return
}

// GetFromContent 通过content获取内容 评论内容
func (obj *_EsMemberCommentMgr) GetFromContent(content string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`content` = ?", content).Find(&results).Error

	return
}

// GetBatchFromContent 批量查找 评论内容
func (obj *_EsMemberCommentMgr) GetBatchFromContent(contents []string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`content` IN (?)", contents).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 评论时间
func (obj *_EsMemberCommentMgr) GetFromCreateTime(createTime int64) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 评论时间
func (obj *_EsMemberCommentMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromHaveImage 通过have_image获取内容 是否有图片 1 有 0 没有
func (obj *_EsMemberCommentMgr) GetFromHaveImage(haveImage int16) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`have_image` = ?", haveImage).Find(&results).Error

	return
}

// GetBatchFromHaveImage 批量查找 是否有图片 1 有 0 没有
func (obj *_EsMemberCommentMgr) GetBatchFromHaveImage(haveImages []int16) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`have_image` IN (?)", haveImages).Find(&results).Error

	return
}

// GetFromStatus 通过status获取内容 状态  1 正常 0 删除
func (obj *_EsMemberCommentMgr) GetFromStatus(status int16) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`status` = ?", status).Find(&results).Error

	return
}

// GetBatchFromStatus 批量查找 状态  1 正常 0 删除
func (obj *_EsMemberCommentMgr) GetBatchFromStatus(statuss []int16) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`status` IN (?)", statuss).Find(&results).Error

	return
}

// GetFromGrade 通过grade获取内容 好中差评
func (obj *_EsMemberCommentMgr) GetFromGrade(grade string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`grade` = ?", grade).Find(&results).Error

	return
}

// GetBatchFromGrade 批量查找 好中差评
func (obj *_EsMemberCommentMgr) GetBatchFromGrade(grades []string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`grade` IN (?)", grades).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容 订单编号
func (obj *_EsMemberCommentMgr) GetFromOrderSn(orderSn string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找 订单编号
func (obj *_EsMemberCommentMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromReplyStatus 通过reply_status获取内容 是否回复 1 是 0 否
func (obj *_EsMemberCommentMgr) GetFromReplyStatus(replyStatus int16) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`reply_status` = ?", replyStatus).Find(&results).Error

	return
}

// GetBatchFromReplyStatus 批量查找 是否回复 1 是 0 否
func (obj *_EsMemberCommentMgr) GetBatchFromReplyStatus(replyStatuss []int16) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`reply_status` IN (?)", replyStatuss).Find(&results).Error

	return
}

// GetFromAuditStatus 通过audit_status获取内容 初评审核:待审核(WAIT_AUDIT),审核通过(PASS_AUDIT),审核拒绝(REFUSE_AUDIT)
func (obj *_EsMemberCommentMgr) GetFromAuditStatus(auditStatus string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`audit_status` = ?", auditStatus).Find(&results).Error

	return
}

// GetBatchFromAuditStatus 批量查找 初评审核:待审核(WAIT_AUDIT),审核通过(PASS_AUDIT),审核拒绝(REFUSE_AUDIT)
func (obj *_EsMemberCommentMgr) GetBatchFromAuditStatus(auditStatuss []string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`audit_status` IN (?)", auditStatuss).Find(&results).Error

	return
}

// GetFromCommentsType 通过comments_type获取内容 评论类型：初评(INITIAL),追评(ADDITIONAL)
func (obj *_EsMemberCommentMgr) GetFromCommentsType(commentsType string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`comments_type` = ?", commentsType).Find(&results).Error

	return
}

// GetBatchFromCommentsType 批量查找 评论类型：初评(INITIAL),追评(ADDITIONAL)
func (obj *_EsMemberCommentMgr) GetBatchFromCommentsType(commentsTypes []string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`comments_type` IN (?)", commentsTypes).Find(&results).Error

	return
}

// GetFromParentID 通过parent_id获取内容 初评id，默认为0
func (obj *_EsMemberCommentMgr) GetFromParentID(parentID int) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`parent_id` = ?", parentID).Find(&results).Error

	return
}

// GetBatchFromParentID 批量查找 初评id，默认为0
func (obj *_EsMemberCommentMgr) GetBatchFromParentID(parentIDs []int) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`parent_id` IN (?)", parentIDs).Find(&results).Error

	return
}

// GetFromAdditionalStatus 通过additional_status获取内容 追加评论状态 0：未追加，1：已追加
func (obj *_EsMemberCommentMgr) GetFromAdditionalStatus(additionalStatus int) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`additional_status` = ?", additionalStatus).Find(&results).Error

	return
}

// GetBatchFromAdditionalStatus 批量查找 追加评论状态 0：未追加，1：已追加
func (obj *_EsMemberCommentMgr) GetBatchFromAdditionalStatus(additionalStatuss []int) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`additional_status` IN (?)", additionalStatuss).Find(&results).Error

	return
}

// GetFromAdditionalContent 通过additional_content获取内容 追加的评论内容
func (obj *_EsMemberCommentMgr) GetFromAdditionalContent(additionalContent string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`additional_content` = ?", additionalContent).Find(&results).Error

	return
}

// GetBatchFromAdditionalContent 批量查找 追加的评论内容
func (obj *_EsMemberCommentMgr) GetBatchFromAdditionalContent(additionalContents []string) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`additional_content` IN (?)", additionalContents).Find(&results).Error

	return
}

// GetFromAdditionalTime 通过additional_time获取内容 追加评论时间
func (obj *_EsMemberCommentMgr) GetFromAdditionalTime(additionalTime int64) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`additional_time` = ?", additionalTime).Find(&results).Error

	return
}

// GetBatchFromAdditionalTime 批量查找 追加评论时间
func (obj *_EsMemberCommentMgr) GetBatchFromAdditionalTime(additionalTimes []int64) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`additional_time` IN (?)", additionalTimes).Find(&results).Error

	return
}

// GetFromAdditionalHaveImage 通过additional_have_image获取内容 追加的评论是否上传了图片 0：未上传，1：已上传
func (obj *_EsMemberCommentMgr) GetFromAdditionalHaveImage(additionalHaveImage int) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`additional_have_image` = ?", additionalHaveImage).Find(&results).Error

	return
}

// GetBatchFromAdditionalHaveImage 批量查找 追加的评论是否上传了图片 0：未上传，1：已上传
func (obj *_EsMemberCommentMgr) GetBatchFromAdditionalHaveImage(additionalHaveImages []int) (results []*models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`additional_have_image` IN (?)", additionalHaveImages).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsMemberCommentMgr) FetchByPrimaryKey(commentID int) (result models.EsMemberComment, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberComment{}).Where("`comment_id` = ?", commentID).First(&result).Error

	return
}
