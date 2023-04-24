package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type RoleMgr struct {
	*_BaseMgr
}

// NewRoleMgr open func
func NewRoleMgr(db db.Repo) *RoleMgr {
	if db == nil {
		panic(fmt.Errorf("NewRoleMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &RoleMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_role"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *RoleMgr) GetTableName() string {
	return "es_role"
}

// Reset 重置gorm会话
func (obj *RoleMgr) Reset() *RoleMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *RoleMgr) Get() (result models.EsRole, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRole{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *RoleMgr) Gets() (results []*models.EsRole, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRole{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *RoleMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsRole{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithRoleID role_id获取 主键ID
func (obj *RoleMgr) WithRoleID(roleID int) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithRoleName role_name获取 角色名称
func (obj *RoleMgr) WithRoleName(roleName string) Option {
	return optionFunc(func(o *options) { o.query["role_name"] = roleName })
}

// WithAuthIDs auth_ids获取 角色介绍
func (obj *RoleMgr) WithAuthIDs(authIDs string) Option {
	return optionFunc(func(o *options) { o.query["auth_ids"] = authIDs })
}

// WithRoleDescribe role_describe获取 角色描述
func (obj *RoleMgr) WithRoleDescribe(roleDescribe string) Option {
	return optionFunc(func(o *options) { o.query["role_describe"] = roleDescribe })
}

// GetByOption 功能选项模式获取
func (obj *RoleMgr) GetByOption(opts ...Option) (result models.EsRole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRole{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *RoleMgr) GetByOptions(opts ...Option) (results []*models.EsRole, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRole{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *RoleMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsRole, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsRole{}).Where(options.query)
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

// GetFromRoleID 通过role_id获取内容 主键ID
func (obj *RoleMgr) GetFromRoleID(roleID int) (result models.EsRole, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRole{}).Where("`role_id` = ?", roleID).First(&result).Error

	return
}

// GetBatchFromRoleID 批量查找 主键ID
func (obj *RoleMgr) GetBatchFromRoleID(roleIDs []int) (results []*models.EsRole, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRole{}).Where("`role_id` IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromRoleName 通过role_name获取内容 角色名称
func (obj *RoleMgr) GetFromRoleName(roleName string) (results []*models.EsRole, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRole{}).Where("`role_name` = ?", roleName).Find(&results).Error

	return
}

// GetBatchFromRoleName 批量查找 角色名称
func (obj *RoleMgr) GetBatchFromRoleName(roleNames []string) (results []*models.EsRole, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRole{}).Where("`role_name` IN (?)", roleNames).Find(&results).Error

	return
}

// GetFromAuthIDs 通过auth_ids获取内容 角色介绍
func (obj *RoleMgr) GetFromAuthIDs(authIDs string) (results []*models.EsRole, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRole{}).Where("`auth_ids` = ?", authIDs).Find(&results).Error

	return
}

// GetBatchFromAuthIDs 批量查找 角色介绍
func (obj *RoleMgr) GetBatchFromAuthIDs(authIDss []string) (results []*models.EsRole, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRole{}).Where("`auth_ids` IN (?)", authIDss).Find(&results).Error

	return
}

// GetFromRoleDescribe 通过role_describe获取内容 角色描述
func (obj *RoleMgr) GetFromRoleDescribe(roleDescribe string) (results []*models.EsRole, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRole{}).Where("`role_describe` = ?", roleDescribe).Find(&results).Error

	return
}

// GetBatchFromRoleDescribe 批量查找 角色描述
func (obj *RoleMgr) GetBatchFromRoleDescribe(roleDescribes []string) (results []*models.EsRole, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRole{}).Where("`role_describe` IN (?)", roleDescribes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *RoleMgr) FetchByPrimaryKey(roleID int) (result models.EsRole, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsRole{}).Where("`role_id` = ?", roleID).First(&result).Error

	return
}
