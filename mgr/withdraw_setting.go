package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type WithdrawSettingMgr struct {
	*_BaseMgr
}

// NewWithdrawSettingMgr open func
func NewWithdrawSettingMgr(db db.Repo) *WithdrawSettingMgr {
	if db == nil {
		panic(fmt.Errorf("NewWithdrawSettingMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &WithdrawSettingMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_withdraw_setting"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *WithdrawSettingMgr) GetTableName() string {
	return "es_withdraw_setting"
}

// Reset 重置gorm会话
func (obj *WithdrawSettingMgr) Reset() *WithdrawSettingMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *WithdrawSettingMgr) Get() (result models.EsWithdrawSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWithdrawSetting{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *WithdrawSettingMgr) Gets() (results []*models.EsWithdrawSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWithdrawSetting{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *WithdrawSettingMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsWithdrawSetting{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *WithdrawSettingMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 用户id
func (obj *WithdrawSettingMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithParam param获取 参数
func (obj *WithdrawSettingMgr) WithParam(param string) Option {
	return optionFunc(func(o *options) { o.query["param"] = param })
}

// GetByOption 功能选项模式获取
func (obj *WithdrawSettingMgr) GetByOption(opts ...Option) (result models.EsWithdrawSetting, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWithdrawSetting{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *WithdrawSettingMgr) GetByOptions(opts ...Option) (results []*models.EsWithdrawSetting, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWithdrawSetting{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *WithdrawSettingMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsWithdrawSetting, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsWithdrawSetting{}).Where(options.query)
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
func (obj *WithdrawSettingMgr) GetFromID(id int) (result models.EsWithdrawSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWithdrawSetting{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *WithdrawSettingMgr) GetBatchFromID(ids []int) (results []*models.EsWithdrawSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWithdrawSetting{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 用户id
func (obj *WithdrawSettingMgr) GetFromMemberID(memberID int) (results []*models.EsWithdrawSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWithdrawSetting{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 用户id
func (obj *WithdrawSettingMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsWithdrawSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWithdrawSetting{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromParam 通过param获取内容 参数
func (obj *WithdrawSettingMgr) GetFromParam(param string) (results []*models.EsWithdrawSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWithdrawSetting{}).Where("`param` = ?", param).Find(&results).Error

	return
}

// GetBatchFromParam 批量查找 参数
func (obj *WithdrawSettingMgr) GetBatchFromParam(params []string) (results []*models.EsWithdrawSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWithdrawSetting{}).Where("`param` IN (?)", params).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *WithdrawSettingMgr) FetchByPrimaryKey(id int) (result models.EsWithdrawSetting, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsWithdrawSetting{}).Where("`id` = ?", id).First(&result).Error

	return
}
