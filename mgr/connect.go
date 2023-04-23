package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsConnectMgr struct {
	*_BaseMgr
}

// EsConnectMgr open func
func EsConnectMgr(db *gorm.DB) *_EsConnectMgr {
	if db == nil {
		panic(fmt.Errorf("EsConnectMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsConnectMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_connect"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsConnectMgr) GetTableName() string {
	return "es_connect"
}

// Reset 重置gorm会话
func (obj *_EsConnectMgr) Reset() *_EsConnectMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsConnectMgr) Get() (result models.EsConnect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnect{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsConnectMgr) Gets() (results []*models.EsConnect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnect{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsConnectMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsConnect{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *_EsConnectMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员id
func (obj *_EsConnectMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithUnionID union_id获取 第三方唯一标示
func (obj *_EsConnectMgr) WithUnionID(unionID string) Option {
	return optionFunc(func(o *options) { o.query["union_id"] = unionID })
}

// WithUnionType union_type获取 信任登录类型
func (obj *_EsConnectMgr) WithUnionType(unionType string) Option {
	return optionFunc(func(o *options) { o.query["union_type"] = unionType })
}

// WithUnboundTime unbound_time获取 解绑时间
func (obj *_EsConnectMgr) WithUnboundTime(unboundTime int64) Option {
	return optionFunc(func(o *options) { o.query["unbound_time"] = unboundTime })
}

// GetByOption 功能选项模式获取
func (obj *_EsConnectMgr) GetByOption(opts ...Option) (result models.EsConnect, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnect{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsConnectMgr) GetByOptions(opts ...Option) (results []*models.EsConnect, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnect{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsConnectMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsConnect, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsConnect{}).Where(options.query)
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
func (obj *_EsConnectMgr) GetFromID(id int) (result models.EsConnect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *_EsConnectMgr) GetBatchFromID(ids []int) (results []*models.EsConnect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_EsConnectMgr) GetFromMemberID(memberID int) (results []*models.EsConnect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_EsConnectMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsConnect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromUnionID 通过union_id获取内容 第三方唯一标示
func (obj *_EsConnectMgr) GetFromUnionID(unionID string) (results []*models.EsConnect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`union_id` = ?", unionID).Find(&results).Error

	return
}

// GetBatchFromUnionID 批量查找 第三方唯一标示
func (obj *_EsConnectMgr) GetBatchFromUnionID(unionIDs []string) (results []*models.EsConnect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`union_id` IN (?)", unionIDs).Find(&results).Error

	return
}

// GetFromUnionType 通过union_type获取内容 信任登录类型
func (obj *_EsConnectMgr) GetFromUnionType(unionType string) (results []*models.EsConnect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`union_type` = ?", unionType).Find(&results).Error

	return
}

// GetBatchFromUnionType 批量查找 信任登录类型
func (obj *_EsConnectMgr) GetBatchFromUnionType(unionTypes []string) (results []*models.EsConnect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`union_type` IN (?)", unionTypes).Find(&results).Error

	return
}

// GetFromUnboundTime 通过unbound_time获取内容 解绑时间
func (obj *_EsConnectMgr) GetFromUnboundTime(unboundTime int64) (results []*models.EsConnect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`unbound_time` = ?", unboundTime).Find(&results).Error

	return
}

// GetBatchFromUnboundTime 批量查找 解绑时间
func (obj *_EsConnectMgr) GetBatchFromUnboundTime(unboundTimes []int64) (results []*models.EsConnect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`unbound_time` IN (?)", unboundTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsConnectMgr) FetchByPrimaryKey(id int) (result models.EsConnect, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`id` = ?", id).First(&result).Error

	return
}
