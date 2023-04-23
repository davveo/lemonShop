package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsSssMemberRegisterDataMgr struct {
	*_BaseMgr
}

// EsSssMemberRegisterDataMgr open func
func EsSssMemberRegisterDataMgr(db *gorm.DB) *_EsSssMemberRegisterDataMgr {
	if db == nil {
		panic(fmt.Errorf("EsSssMemberRegisterDataMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsSssMemberRegisterDataMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_sss_member_register_data"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsSssMemberRegisterDataMgr) GetTableName() string {
	return "es_sss_member_register_data"
}

// Reset 重置gorm会话
func (obj *_EsSssMemberRegisterDataMgr) Reset() *_EsSssMemberRegisterDataMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsSssMemberRegisterDataMgr) Get() (result models.EsSssMemberRegisterData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsSssMemberRegisterDataMgr) Gets() (results []*models.EsSssMemberRegisterData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsSssMemberRegisterDataMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_EsSssMemberRegisterDataMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员id
func (obj *_EsSssMemberRegisterDataMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMemberName member_name获取 会员名字
func (obj *_EsSssMemberRegisterDataMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithCreateTime create_time获取 注册日期
func (obj *_EsSssMemberRegisterDataMgr) WithCreateTime(createTime int) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *_EsSssMemberRegisterDataMgr) GetByOption(opts ...Option) (result models.EsSssMemberRegisterData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsSssMemberRegisterDataMgr) GetByOptions(opts ...Option) (results []*models.EsSssMemberRegisterData, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsSssMemberRegisterDataMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSssMemberRegisterData, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where(options.query)
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
func (obj *_EsSssMemberRegisterDataMgr) GetFromID(id int) (result models.EsSssMemberRegisterData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_EsSssMemberRegisterDataMgr) GetBatchFromID(ids []int) (results []*models.EsSssMemberRegisterData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_EsSssMemberRegisterDataMgr) GetFromMemberID(memberID int) (results []*models.EsSssMemberRegisterData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_EsSssMemberRegisterDataMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsSssMemberRegisterData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 会员名字
func (obj *_EsSssMemberRegisterDataMgr) GetFromMemberName(memberName string) (results []*models.EsSssMemberRegisterData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 会员名字
func (obj *_EsSssMemberRegisterDataMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsSssMemberRegisterData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 注册日期
func (obj *_EsSssMemberRegisterDataMgr) GetFromCreateTime(createTime int) (results []*models.EsSssMemberRegisterData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 注册日期
func (obj *_EsSssMemberRegisterDataMgr) GetBatchFromCreateTime(createTimes []int) (results []*models.EsSssMemberRegisterData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsSssMemberRegisterDataMgr) FetchByPrimaryKey(id int) (result models.EsSssMemberRegisterData, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSssMemberRegisterData{}).Where("`id` = ?", id).First(&result).Error

	return
}
