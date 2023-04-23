package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsUpgradeLogMgr struct {
	*_BaseMgr
}

// EsUpgradeLogMgr open func
func EsUpgradeLogMgr(db *gorm.DB) *_EsUpgradeLogMgr {
	if db == nil {
		panic(fmt.Errorf("EsUpgradeLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsUpgradeLogMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_upgrade_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsUpgradeLogMgr) GetTableName() string {
	return "es_upgrade_log"
}

// Reset 重置gorm会话
func (obj *_EsUpgradeLogMgr) Reset() *_EsUpgradeLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsUpgradeLogMgr) Get() (result models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsUpgradeLogMgr) Gets() (results []*models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsUpgradeLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_EsUpgradeLogMgr) WithID(id int64) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员id
func (obj *_EsUpgradeLogMgr) WithMemberID(memberID int64) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMemberName member_name获取 会员名称
func (obj *_EsUpgradeLogMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithType type获取 切换类型
func (obj *_EsUpgradeLogMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// WithOldTplID old_tpl_id获取 旧的模板id
func (obj *_EsUpgradeLogMgr) WithOldTplID(oldTplID int64) Option {
	return optionFunc(func(o *options) { o.query["old_tpl_id"] = oldTplID })
}

// WithNewTplID new_tpl_id获取 新的模板id
func (obj *_EsUpgradeLogMgr) WithNewTplID(newTplID int64) Option {
	return optionFunc(func(o *options) { o.query["new_tpl_id"] = newTplID })
}

// WithNewTplName new_tpl_name获取 新的模板名称
func (obj *_EsUpgradeLogMgr) WithNewTplName(newTplName string) Option {
	return optionFunc(func(o *options) { o.query["new_tpl_name"] = newTplName })
}

// WithOldTplName old_tpl_name获取 旧的模板名称
func (obj *_EsUpgradeLogMgr) WithOldTplName(oldTplName string) Option {
	return optionFunc(func(o *options) { o.query["old_tpl_name"] = oldTplName })
}

// WithCreateTime create_time获取 创建日期
func (obj *_EsUpgradeLogMgr) WithCreateTime(createTime int) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// GetByOption 功能选项模式获取
func (obj *_EsUpgradeLogMgr) GetByOption(opts ...Option) (result models.EsUpgradeLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsUpgradeLogMgr) GetByOptions(opts ...Option) (results []*models.EsUpgradeLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsUpgradeLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsUpgradeLog, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where(options.query)
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
func (obj *_EsUpgradeLogMgr) GetFromID(id int64) (result models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_EsUpgradeLogMgr) GetBatchFromID(ids []int64) (results []*models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_EsUpgradeLogMgr) GetFromMemberID(memberID int64) (results []*models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_EsUpgradeLogMgr) GetBatchFromMemberID(memberIDs []int64) (results []*models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 会员名称
func (obj *_EsUpgradeLogMgr) GetFromMemberName(memberName string) (results []*models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 会员名称
func (obj *_EsUpgradeLogMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 切换类型
func (obj *_EsUpgradeLogMgr) GetFromType(_type string) (results []*models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 切换类型
func (obj *_EsUpgradeLogMgr) GetBatchFromType(_types []string) (results []*models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

// GetFromOldTplID 通过old_tpl_id获取内容 旧的模板id
func (obj *_EsUpgradeLogMgr) GetFromOldTplID(oldTplID int64) (results []*models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`old_tpl_id` = ?", oldTplID).Find(&results).Error

	return
}

// GetBatchFromOldTplID 批量查找 旧的模板id
func (obj *_EsUpgradeLogMgr) GetBatchFromOldTplID(oldTplIDs []int64) (results []*models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`old_tpl_id` IN (?)", oldTplIDs).Find(&results).Error

	return
}

// GetFromNewTplID 通过new_tpl_id获取内容 新的模板id
func (obj *_EsUpgradeLogMgr) GetFromNewTplID(newTplID int64) (results []*models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`new_tpl_id` = ?", newTplID).Find(&results).Error

	return
}

// GetBatchFromNewTplID 批量查找 新的模板id
func (obj *_EsUpgradeLogMgr) GetBatchFromNewTplID(newTplIDs []int64) (results []*models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`new_tpl_id` IN (?)", newTplIDs).Find(&results).Error

	return
}

// GetFromNewTplName 通过new_tpl_name获取内容 新的模板名称
func (obj *_EsUpgradeLogMgr) GetFromNewTplName(newTplName string) (results []*models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`new_tpl_name` = ?", newTplName).Find(&results).Error

	return
}

// GetBatchFromNewTplName 批量查找 新的模板名称
func (obj *_EsUpgradeLogMgr) GetBatchFromNewTplName(newTplNames []string) (results []*models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`new_tpl_name` IN (?)", newTplNames).Find(&results).Error

	return
}

// GetFromOldTplName 通过old_tpl_name获取内容 旧的模板名称
func (obj *_EsUpgradeLogMgr) GetFromOldTplName(oldTplName string) (results []*models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`old_tpl_name` = ?", oldTplName).Find(&results).Error

	return
}

// GetBatchFromOldTplName 批量查找 旧的模板名称
func (obj *_EsUpgradeLogMgr) GetBatchFromOldTplName(oldTplNames []string) (results []*models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`old_tpl_name` IN (?)", oldTplNames).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建日期
func (obj *_EsUpgradeLogMgr) GetFromCreateTime(createTime int) (results []*models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建日期
func (obj *_EsUpgradeLogMgr) GetBatchFromCreateTime(createTimes []int) (results []*models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsUpgradeLogMgr) FetchByPrimaryKey(id int64) (result models.EsUpgradeLog, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsUpgradeLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
