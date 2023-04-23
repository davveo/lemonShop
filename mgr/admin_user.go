package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type AdminUserMgr struct {
	*_BaseMgr
}

// NewAdminUserMgr open func
func NewAdminUserMgr(db db.Repo) *AdminUserMgr {
	if db == nil {
		panic(fmt.Errorf("EsSpecificationDao need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &AdminUserMgr{_BaseMgr: &_BaseMgr{
		rdb:       db.GetDbR().Table("es_admin_user"),
		wdb:       db.GetDbW().Table("es_admin_user"),
		isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *AdminUserMgr) GetTableName() string {
	return "es_admin_user"
}

func (obj *AdminUserMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Count(count)
}

// WithID id获取 平台管理员id
func (obj *AdminUserMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUsername username获取 管理员名称
func (obj *AdminUserMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithPassword password获取 管理员密码
func (obj *AdminUserMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithDepartment department获取 部门
func (obj *AdminUserMgr) WithDepartment(department string) Option {
	return optionFunc(func(o *options) { o.query["department"] = department })
}

// WithRoleID role_id获取 权限id
func (obj *AdminUserMgr) WithRoleID(roleID int) Option {
	return optionFunc(func(o *options) { o.query["role_id"] = roleID })
}

// WithDateLine date_line获取 创建日期
func (obj *AdminUserMgr) WithDateLine(dateLine int64) Option {
	return optionFunc(func(o *options) { o.query["date_line"] = dateLine })
}

// WithRemark remark获取 备注
func (obj *AdminUserMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithUserState user_state获取 是否删除
func (obj *AdminUserMgr) WithUserState(userState int) Option {
	return optionFunc(func(o *options) { o.query["user_state"] = userState })
}

// WithRealName real_name获取 管理员真实姓名
func (obj *AdminUserMgr) WithRealName(realName string) Option {
	return optionFunc(func(o *options) { o.query["real_name"] = realName })
}

// WithFace face获取 头像
func (obj *AdminUserMgr) WithFace(face string) Option {
	return optionFunc(func(o *options) { o.query["face"] = face })
}

// WithFounder founder获取 是否为超级管理员
func (obj *AdminUserMgr) WithFounder(founder int) Option {
	return optionFunc(func(o *options) { o.query["founder"] = founder })
}

// GetByOption 功能选项模式获取
func (obj *AdminUserMgr) GetByOption(opts ...Option) (result models.EsAdminUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *AdminUserMgr) GetByOptions(opts ...Option) (results []*models.EsAdminUser, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *AdminUserMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsAdminUser, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where(options.query)
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

// GetFromID 通过id获取内容 平台管理员id
func (obj *AdminUserMgr) GetFromID(id int) (result models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 平台管理员id
func (obj *AdminUserMgr) GetBatchFromID(ids []int) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 管理员名称
func (obj *AdminUserMgr) GetFromUsername(username string) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`username` = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量查找 管理员名称
func (obj *AdminUserMgr) GetBatchFromUsername(usernames []string) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 管理员密码
func (obj *AdminUserMgr) GetFromPassword(password string) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 管理员密码
func (obj *AdminUserMgr) GetBatchFromPassword(passwords []string) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromDepartment 通过department获取内容 部门
func (obj *AdminUserMgr) GetFromDepartment(department string) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`department` = ?", department).Find(&results).Error

	return
}

// GetBatchFromDepartment 批量查找 部门
func (obj *AdminUserMgr) GetBatchFromDepartment(departments []string) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`department` IN (?)", departments).Find(&results).Error

	return
}

// GetFromRoleID 通过role_id获取内容 权限id
func (obj *AdminUserMgr) GetFromRoleID(roleID int) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`role_id` = ?", roleID).Find(&results).Error

	return
}

// GetBatchFromRoleID 批量查找 权限id
func (obj *AdminUserMgr) GetBatchFromRoleID(roleIDs []int) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`role_id` IN (?)", roleIDs).Find(&results).Error

	return
}

// GetFromDateLine 通过date_line获取内容 创建日期
func (obj *AdminUserMgr) GetFromDateLine(dateLine int64) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`date_line` = ?", dateLine).Find(&results).Error

	return
}

// GetBatchFromDateLine 批量查找 创建日期
func (obj *AdminUserMgr) GetBatchFromDateLine(dateLines []int64) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`date_line` IN (?)", dateLines).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 备注
func (obj *AdminUserMgr) GetFromRemark(remark string) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 备注
func (obj *AdminUserMgr) GetBatchFromRemark(remarks []string) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromUserState 通过user_state获取内容 是否删除
func (obj *AdminUserMgr) GetFromUserState(userState int) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`user_state` = ?", userState).Find(&results).Error

	return
}

// GetBatchFromUserState 批量查找 是否删除
func (obj *AdminUserMgr) GetBatchFromUserState(userStates []int) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`user_state` IN (?)", userStates).Find(&results).Error

	return
}

// GetFromRealName 通过real_name获取内容 管理员真实姓名
func (obj *AdminUserMgr) GetFromRealName(realName string) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`real_name` = ?", realName).Find(&results).Error

	return
}

// GetBatchFromRealName 批量查找 管理员真实姓名
func (obj *AdminUserMgr) GetBatchFromRealName(realNames []string) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`real_name` IN (?)", realNames).Find(&results).Error

	return
}

// GetFromFace 通过face获取内容 头像
func (obj *AdminUserMgr) GetFromFace(face string) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`face` = ?", face).Find(&results).Error

	return
}

// GetBatchFromFace 批量查找 头像
func (obj *AdminUserMgr) GetBatchFromFace(faces []string) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`face` IN (?)", faces).Find(&results).Error

	return
}

// GetFromFounder 通过founder获取内容 是否为超级管理员
func (obj *AdminUserMgr) GetFromFounder(founder int) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`founder` = ?", founder).Find(&results).Error

	return
}

// GetBatchFromFounder 批量查找 是否为超级管理员
func (obj *AdminUserMgr) GetBatchFromFounder(founders []int) (results []*models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`founder` IN (?)", founders).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *AdminUserMgr) FetchByPrimaryKey(id int) (result models.EsAdminUser, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsAdminUser{}).Where("`id` = ?", id).First(&result).Error

	return
}
