package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type SssMemberRegisterDataMgr struct {
	*_BaseMgr
}

// NewSssMemberRegisterDataMgr open func
func NewSssMemberRegisterDataMgr(db db.Repo) *SssMemberRegisterDataMgr {
	if db == nil {
		panic(fmt.Errorf("NewSssMemberRegisterDataMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &SssMemberRegisterDataMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_sss_member_register_data"), wdb: db.GetDbW().Table("es_sss_member_register_data"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *SssMemberRegisterDataMgr) GetTableName() string {
	return "es_sss_member_register_data"
}

// Reset 重置gorm会话
func (obj *SssMemberRegisterDataMgr) Reset() *SssMemberRegisterDataMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *SssMemberRegisterDataMgr) Get() (result models.EsSssMemberRegisterData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *SssMemberRegisterDataMgr) Gets() (results []*models.EsSssMemberRegisterData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *SssMemberRegisterDataMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *SssMemberRegisterDataMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员id
func (obj *SssMemberRegisterDataMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMemberName member_name获取 会员名字
func (obj *SssMemberRegisterDataMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithCreateTime create_time获取 注册日期
func (obj *SssMemberRegisterDataMgr) WithCreateTime(createTime int) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *SssMemberRegisterDataMgr) GetByOption(opts ...Option) (result models.EsSssMemberRegisterData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *SssMemberRegisterDataMgr) GetByOptions(opts ...Option) (results []*models.EsSssMemberRegisterData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *SssMemberRegisterDataMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSssMemberRegisterData, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where(options.query)
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

// GetFromID 通过id获取内容 id
func (obj *SssMemberRegisterDataMgr) GetFromID(id int) (result models.EsSssMemberRegisterData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *SssMemberRegisterDataMgr) GetBatchFromID(ids []int) (results []*models.EsSssMemberRegisterData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *SssMemberRegisterDataMgr) GetFromMemberID(memberID int) (results []*models.EsSssMemberRegisterData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *SssMemberRegisterDataMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsSssMemberRegisterData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 会员名字
func (obj *SssMemberRegisterDataMgr) GetFromMemberName(memberName string) (results []*models.EsSssMemberRegisterData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 会员名字
func (obj *SssMemberRegisterDataMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsSssMemberRegisterData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 注册日期
func (obj *SssMemberRegisterDataMgr) GetFromCreateTime(createTime int) (results []*models.EsSssMemberRegisterData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 注册日期
func (obj *SssMemberRegisterDataMgr) GetBatchFromCreateTime(createTimes []int) (results []*models.EsSssMemberRegisterData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *SssMemberRegisterDataMgr) FetchByPrimaryKey(id int) (result models.EsSssMemberRegisterData, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where("`id` = ?", id).First(&result).Error

	return
}
