package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type MemberPointHistoryMgr struct {
	*_BaseMgr
}

// NewMemberPointHistoryMgr open func
func NewMemberPointHistoryMgr(db db.Repo) *MemberPointHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("NewMemberPointHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &MemberPointHistoryMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_member_point_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *MemberPointHistoryMgr) GetTableName() string {
	return "es_member_point_history"
}

// Reset 重置gorm会话
func (obj *MemberPointHistoryMgr) Reset() *MemberPointHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *MemberPointHistoryMgr) Get() (result models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *MemberPointHistoryMgr) Gets() (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *MemberPointHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *MemberPointHistoryMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员ID
func (obj *MemberPointHistoryMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithGradePoint grade_point获取 等级积分
func (obj *MemberPointHistoryMgr) WithGradePoint(gradePoint int64) Option {
	return optionFunc(func(o *options) { o.query["grade_point"] = gradePoint })
}

// WithTime time获取 操作时间
func (obj *MemberPointHistoryMgr) WithTime(time int64) Option {
	return optionFunc(func(o *options) { o.query["time"] = time })
}

// WithReason reason获取 操作理由
func (obj *MemberPointHistoryMgr) WithReason(reason string) Option {
	return optionFunc(func(o *options) { o.query["reason"] = reason })
}

// WithGradePointType grade_point_type获取 等级积分类型
func (obj *MemberPointHistoryMgr) WithGradePointType(gradePointType int) Option {
	return optionFunc(func(o *options) { o.query["grade_point_type"] = gradePointType })
}

// WithOperator operator获取 操作者
func (obj *MemberPointHistoryMgr) WithOperator(operator string) Option {
	return optionFunc(func(o *options) { o.query["operator"] = operator })
}

// WithConsumPoint consum_point获取 消费积分
func (obj *MemberPointHistoryMgr) WithConsumPoint(consumPoint int64) Option {
	return optionFunc(func(o *options) { o.query["consum_point"] = consumPoint })
}

// WithConsumPointType consum_point_type获取 消费积分类型
func (obj *MemberPointHistoryMgr) WithConsumPointType(consumPointType int) Option {
	return optionFunc(func(o *options) { o.query["consum_point_type"] = consumPointType })
}

// GetByOption 功能选项模式获取
func (obj *MemberPointHistoryMgr) GetByOption(opts ...Option) (result models.EsMemberPointHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *MemberPointHistoryMgr) GetByOptions(opts ...Option) (results []*models.EsMemberPointHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *MemberPointHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsMemberPointHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where(options.query)
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
func (obj *MemberPointHistoryMgr) GetFromID(id int) (result models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *MemberPointHistoryMgr) GetBatchFromID(ids []int) (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员ID
func (obj *MemberPointHistoryMgr) GetFromMemberID(memberID int) (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员ID
func (obj *MemberPointHistoryMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromGradePoint 通过grade_point获取内容 等级积分
func (obj *MemberPointHistoryMgr) GetFromGradePoint(gradePoint int64) (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`grade_point` = ?", gradePoint).Find(&results).Error

	return
}

// GetBatchFromGradePoint 批量查找 等级积分
func (obj *MemberPointHistoryMgr) GetBatchFromGradePoint(gradePoints []int64) (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`grade_point` IN (?)", gradePoints).Find(&results).Error

	return
}

// GetFromTime 通过time获取内容 操作时间
func (obj *MemberPointHistoryMgr) GetFromTime(time int64) (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`time` = ?", time).Find(&results).Error

	return
}

// GetBatchFromTime 批量查找 操作时间
func (obj *MemberPointHistoryMgr) GetBatchFromTime(times []int64) (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`time` IN (?)", times).Find(&results).Error

	return
}

// GetFromReason 通过reason获取内容 操作理由
func (obj *MemberPointHistoryMgr) GetFromReason(reason string) (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`reason` = ?", reason).Find(&results).Error

	return
}

// GetBatchFromReason 批量查找 操作理由
func (obj *MemberPointHistoryMgr) GetBatchFromReason(reasons []string) (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`reason` IN (?)", reasons).Find(&results).Error

	return
}

// GetFromGradePointType 通过grade_point_type获取内容 等级积分类型
func (obj *MemberPointHistoryMgr) GetFromGradePointType(gradePointType int) (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`grade_point_type` = ?", gradePointType).Find(&results).Error

	return
}

// GetBatchFromGradePointType 批量查找 等级积分类型
func (obj *MemberPointHistoryMgr) GetBatchFromGradePointType(gradePointTypes []int) (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`grade_point_type` IN (?)", gradePointTypes).Find(&results).Error

	return
}

// GetFromOperator 通过operator获取内容 操作者
func (obj *MemberPointHistoryMgr) GetFromOperator(operator string) (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`operator` = ?", operator).Find(&results).Error

	return
}

// GetBatchFromOperator 批量查找 操作者
func (obj *MemberPointHistoryMgr) GetBatchFromOperator(operators []string) (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`operator` IN (?)", operators).Find(&results).Error

	return
}

// GetFromConsumPoint 通过consum_point获取内容 消费积分
func (obj *MemberPointHistoryMgr) GetFromConsumPoint(consumPoint int64) (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`consum_point` = ?", consumPoint).Find(&results).Error

	return
}

// GetBatchFromConsumPoint 批量查找 消费积分
func (obj *MemberPointHistoryMgr) GetBatchFromConsumPoint(consumPoints []int64) (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`consum_point` IN (?)", consumPoints).Find(&results).Error

	return
}

// GetFromConsumPointType 通过consum_point_type获取内容 消费积分类型
func (obj *MemberPointHistoryMgr) GetFromConsumPointType(consumPointType int) (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`consum_point_type` = ?", consumPointType).Find(&results).Error

	return
}

// GetBatchFromConsumPointType 批量查找 消费积分类型
func (obj *MemberPointHistoryMgr) GetBatchFromConsumPointType(consumPointTypes []int) (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`consum_point_type` IN (?)", consumPointTypes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *MemberPointHistoryMgr) FetchByPrimaryKey(id int) (result models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`id` = ?", id).First(&result).Error

	return
}

// FetchIndexByIndPonitHistory  获取多个内容
func (obj *MemberPointHistoryMgr) FetchIndexByIndPonitHistory(memberID int, gradePointType int, consumPointType int) (results []*models.EsMemberPointHistory, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberPointHistory{}).Where("`member_id` = ? AND `grade_point_type` = ? AND `consum_point_type` = ?", memberID, gradePointType, consumPointType).Find(&results).Error

	return
}
