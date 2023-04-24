package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type MemberCouponMgr struct {
	*_BaseMgr
}

// NewMemberCouponMgr open func
func NewMemberCouponMgr(db db.Repo) *MemberCouponMgr {
	if db == nil {
		panic(fmt.Errorf("NewMemberCouponMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &MemberCouponMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_member_coupon"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *MemberCouponMgr) GetTableName() string {
	return "es_member_coupon"
}

// Reset 重置gorm会话
func (obj *MemberCouponMgr) Reset() *MemberCouponMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *MemberCouponMgr) Get() (result models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *MemberCouponMgr) Gets() (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *MemberCouponMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithMcID mc_id获取
func (obj *MemberCouponMgr) WithMcID(mcID int) Option {
	return optionFunc(func(o *options) { o.query["mc_id"] = mcID })
}

// WithCouponID coupon_id获取
func (obj *MemberCouponMgr) WithCouponID(couponID int) Option {
	return optionFunc(func(o *options) { o.query["coupon_id"] = couponID })
}

// WithMemberID member_id获取
func (obj *MemberCouponMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithUsedTime used_time获取
func (obj *MemberCouponMgr) WithUsedTime(usedTime int64) Option {
	return optionFunc(func(o *options) { o.query["used_time"] = usedTime })
}

// WithCreateTime create_time获取
func (obj *MemberCouponMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithOrderID order_id获取
func (obj *MemberCouponMgr) WithOrderID(orderID int) Option {
	return optionFunc(func(o *options) { o.query["order_id"] = orderID })
}

// WithOrderSn order_sn获取
func (obj *MemberCouponMgr) WithOrderSn(orderSn string) Option {
	return optionFunc(func(o *options) { o.query["order_sn"] = orderSn })
}

// WithMemberName member_name获取
func (obj *MemberCouponMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithCouponSn coupon_sn获取
func (obj *MemberCouponMgr) WithCouponSn(couponSn string) Option {
	return optionFunc(func(o *options) { o.query["coupon_sn"] = couponSn })
}

// WithTitle title获取
func (obj *MemberCouponMgr) WithTitle(title string) Option {
	return optionFunc(func(o *options) { o.query["title"] = title })
}

// GetByOption 功能选项模式获取
func (obj *MemberCouponMgr) GetByOption(opts ...Option) (result models.EsMemberCoupon, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *MemberCouponMgr) GetByOptions(opts ...Option) (results []*models.EsMemberCoupon, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *MemberCouponMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsMemberCoupon, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where(options.query)
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

// GetFromMcID 通过mc_id获取内容
func (obj *MemberCouponMgr) GetFromMcID(mcID int) (result models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`mc_id` = ?", mcID).First(&result).Error

	return
}

// GetBatchFromMcID 批量查找
func (obj *MemberCouponMgr) GetBatchFromMcID(mcIDs []int) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`mc_id` IN (?)", mcIDs).Find(&results).Error

	return
}

// GetFromCouponID 通过coupon_id获取内容
func (obj *MemberCouponMgr) GetFromCouponID(couponID int) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`coupon_id` = ?", couponID).Find(&results).Error

	return
}

// GetBatchFromCouponID 批量查找
func (obj *MemberCouponMgr) GetBatchFromCouponID(couponIDs []int) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`coupon_id` IN (?)", couponIDs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容
func (obj *MemberCouponMgr) GetFromMemberID(memberID int) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找
func (obj *MemberCouponMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromUsedTime 通过used_time获取内容
func (obj *MemberCouponMgr) GetFromUsedTime(usedTime int64) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`used_time` = ?", usedTime).Find(&results).Error

	return
}

// GetBatchFromUsedTime 批量查找
func (obj *MemberCouponMgr) GetBatchFromUsedTime(usedTimes []int64) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`used_time` IN (?)", usedTimes).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容
func (obj *MemberCouponMgr) GetFromCreateTime(createTime int64) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找
func (obj *MemberCouponMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromOrderID 通过order_id获取内容
func (obj *MemberCouponMgr) GetFromOrderID(orderID int) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`order_id` = ?", orderID).Find(&results).Error

	return
}

// GetBatchFromOrderID 批量查找
func (obj *MemberCouponMgr) GetBatchFromOrderID(orderIDs []int) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`order_id` IN (?)", orderIDs).Find(&results).Error

	return
}

// GetFromOrderSn 通过order_sn获取内容
func (obj *MemberCouponMgr) GetFromOrderSn(orderSn string) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`order_sn` = ?", orderSn).Find(&results).Error

	return
}

// GetBatchFromOrderSn 批量查找
func (obj *MemberCouponMgr) GetBatchFromOrderSn(orderSns []string) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`order_sn` IN (?)", orderSns).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容
func (obj *MemberCouponMgr) GetFromMemberName(memberName string) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找
func (obj *MemberCouponMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromCouponSn 通过coupon_sn获取内容
func (obj *MemberCouponMgr) GetFromCouponSn(couponSn string) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`coupon_sn` = ?", couponSn).Find(&results).Error

	return
}

// GetBatchFromCouponSn 批量查找
func (obj *MemberCouponMgr) GetBatchFromCouponSn(couponSns []string) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`coupon_sn` IN (?)", couponSns).Find(&results).Error

	return
}

// GetFromTitle 通过title获取内容
func (obj *MemberCouponMgr) GetFromTitle(title string) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`title` = ?", title).Find(&results).Error

	return
}

// GetBatchFromTitle 批量查找
func (obj *MemberCouponMgr) GetBatchFromTitle(titles []string) (results []*models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`title` IN (?)", titles).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *MemberCouponMgr) FetchByPrimaryKey(mcID int) (result models.EsMemberCoupon, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCoupon{}).Where("`mc_id` = ?", mcID).First(&result).Error

	return
}
