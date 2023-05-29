package mgr

import (
	"context"
	"fmt"

	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type ConnectMgr struct {
	*_BaseMgr
}

// NewConnectMgr open func
func NewConnectMgr(db db.Repo) *ConnectMgr {
	if db == nil {
		panic(fmt.Errorf("NewConnectMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &ConnectMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_connect"),
		wdb:       db.GetDbW().Table("es_connect"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *ConnectMgr) GetTableName() string {
	return "es_connect"
}

// Reset 重置gorm会话
func (obj *ConnectMgr) Reset() *ConnectMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *ConnectMgr) Get() (result models.EsConnect, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnect{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *ConnectMgr) Gets() (results []*models.EsConnect, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnect{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *ConnectMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsConnect{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *ConnectMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员id
func (obj *ConnectMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithUnionID union_id获取 第三方唯一标示
func (obj *ConnectMgr) WithUnionID(unionID string) Option {
	return optionFunc(func(o *options) { o.query["union_id"] = unionID })
}

// WithUnionType union_type获取 信任登录类型
func (obj *ConnectMgr) WithUnionType(unionType string) Option {
	return optionFunc(func(o *options) { o.query["union_type"] = unionType })
}

// WithUnboundTime unbound_time获取 解绑时间
func (obj *ConnectMgr) WithUnboundTime(unboundTime int64) Option {
	return optionFunc(func(o *options) { o.query["unbound_time"] = unboundTime })
}

// GetByOption 功能选项模式获取
func (obj *ConnectMgr) GetByOption(opts ...Option) (result models.EsConnect, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnect{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *ConnectMgr) GetByOptions(opts ...Option) (results []*models.EsConnect, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnect{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *ConnectMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsConnect, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsConnect{}).Where(options.query)
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
func (obj *ConnectMgr) GetFromID(id int) (result models.EsConnect, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *ConnectMgr) GetBatchFromID(ids []int) (results []*models.EsConnect, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *ConnectMgr) GetFromMemberID(memberID int) (results []*models.EsConnect, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *ConnectMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsConnect, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromUnionID 通过union_id获取内容 第三方唯一标示
func (obj *ConnectMgr) GetFromUnionID(unionID string) (results []*models.EsConnect, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`union_id` = ?", unionID).Find(&results).Error

	return
}

// GetBatchFromUnionID 批量查找 第三方唯一标示
func (obj *ConnectMgr) GetBatchFromUnionID(unionIDs []string) (results []*models.EsConnect, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`union_id` IN (?)", unionIDs).Find(&results).Error

	return
}

// GetFromUnionType 通过union_type获取内容 信任登录类型
func (obj *ConnectMgr) GetFromUnionType(unionType string) (results []*models.EsConnect, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`union_type` = ?", unionType).Find(&results).Error

	return
}

// GetBatchFromUnionType 批量查找 信任登录类型
func (obj *ConnectMgr) GetBatchFromUnionType(unionTypes []string) (results []*models.EsConnect, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`union_type` IN (?)", unionTypes).Find(&results).Error

	return
}

// GetFromUnboundTime 通过unbound_time获取内容 解绑时间
func (obj *ConnectMgr) GetFromUnboundTime(unboundTime int64) (results []*models.EsConnect, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`unbound_time` = ?", unboundTime).Find(&results).Error

	return
}

// GetBatchFromUnboundTime 批量查找 解绑时间
func (obj *ConnectMgr) GetBatchFromUnboundTime(unboundTimes []int64) (results []*models.EsConnect, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`unbound_time` IN (?)", unboundTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *ConnectMgr) FetchByPrimaryKey(id int) (result models.EsConnect, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsConnect{}).Where("`id` = ?", id).First(&result).Error

	return
}
