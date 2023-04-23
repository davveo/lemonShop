package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsSMTPMgr struct {
	*_BaseMgr
}

// EsSMTPMgr open func
func EsSMTPMgr(db *gorm.DB) *_EsSMTPMgr {
	if db == nil {
		panic(fmt.Errorf("EsSMTPMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsSMTPMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_smtp"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsSMTPMgr) GetTableName() string {
	return "es_smtp"
}

// Reset 重置gorm会话
func (obj *_EsSMTPMgr) Reset() *_EsSMTPMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsSMTPMgr) Get() (result models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsSMTPMgr) Gets() (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsSMTPMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键ID
func (obj *_EsSMTPMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithHost host获取 主机
func (obj *_EsSMTPMgr) WithHost(host string) Option {
	return optionFunc(func(o *options) { o.query["host"] = host })
}

// WithUsername username获取 用户名
func (obj *_EsSMTPMgr) WithUsername(username string) Option {
	return optionFunc(func(o *options) { o.query["username"] = username })
}

// WithPassword password获取 密码
func (obj *_EsSMTPMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithLastSendTime last_send_time获取 最后发信时间
func (obj *_EsSMTPMgr) WithLastSendTime(lastSendTime int64) Option {
	return optionFunc(func(o *options) { o.query["last_send_time"] = lastSendTime })
}

// WithSendCount send_count获取 已发数
func (obj *_EsSMTPMgr) WithSendCount(sendCount int) Option {
	return optionFunc(func(o *options) { o.query["send_count"] = sendCount })
}

// WithMaxCount max_count获取 最大发信数
func (obj *_EsSMTPMgr) WithMaxCount(maxCount int) Option {
	return optionFunc(func(o *options) { o.query["max_count"] = maxCount })
}

// WithMailFrom mail_from获取 发信邮箱
func (obj *_EsSMTPMgr) WithMailFrom(mailFrom string) Option {
	return optionFunc(func(o *options) { o.query["mail_from"] = mailFrom })
}

// WithPort port获取 端口
func (obj *_EsSMTPMgr) WithPort(port int) Option {
	return optionFunc(func(o *options) { o.query["port"] = port })
}

// WithOpenSsl open_ssl获取 ssl是否开启
func (obj *_EsSMTPMgr) WithOpenSsl(openSsl int) Option {
	return optionFunc(func(o *options) { o.query["open_ssl"] = openSsl })
}

// GetByOption 功能选项模式获取
func (obj *_EsSMTPMgr) GetByOption(opts ...Option) (result models.EsSMTP, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsSMTPMgr) GetByOptions(opts ...Option) (results []*models.EsSMTP, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsSMTPMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsSMTP, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where(options.query)
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
func (obj *_EsSMTPMgr) GetFromID(id int) (result models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键ID
func (obj *_EsSMTPMgr) GetBatchFromID(ids []int) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromHost 通过host获取内容 主机
func (obj *_EsSMTPMgr) GetFromHost(host string) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`host` = ?", host).Find(&results).Error

	return
}

// GetBatchFromHost 批量查找 主机
func (obj *_EsSMTPMgr) GetBatchFromHost(hosts []string) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`host` IN (?)", hosts).Find(&results).Error

	return
}

// GetFromUsername 通过username获取内容 用户名
func (obj *_EsSMTPMgr) GetFromUsername(username string) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`username` = ?", username).Find(&results).Error

	return
}

// GetBatchFromUsername 批量查找 用户名
func (obj *_EsSMTPMgr) GetBatchFromUsername(usernames []string) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`username` IN (?)", usernames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 密码
func (obj *_EsSMTPMgr) GetFromPassword(password string) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 密码
func (obj *_EsSMTPMgr) GetBatchFromPassword(passwords []string) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromLastSendTime 通过last_send_time获取内容 最后发信时间
func (obj *_EsSMTPMgr) GetFromLastSendTime(lastSendTime int64) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`last_send_time` = ?", lastSendTime).Find(&results).Error

	return
}

// GetBatchFromLastSendTime 批量查找 最后发信时间
func (obj *_EsSMTPMgr) GetBatchFromLastSendTime(lastSendTimes []int64) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`last_send_time` IN (?)", lastSendTimes).Find(&results).Error

	return
}

// GetFromSendCount 通过send_count获取内容 已发数
func (obj *_EsSMTPMgr) GetFromSendCount(sendCount int) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`send_count` = ?", sendCount).Find(&results).Error

	return
}

// GetBatchFromSendCount 批量查找 已发数
func (obj *_EsSMTPMgr) GetBatchFromSendCount(sendCounts []int) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`send_count` IN (?)", sendCounts).Find(&results).Error

	return
}

// GetFromMaxCount 通过max_count获取内容 最大发信数
func (obj *_EsSMTPMgr) GetFromMaxCount(maxCount int) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`max_count` = ?", maxCount).Find(&results).Error

	return
}

// GetBatchFromMaxCount 批量查找 最大发信数
func (obj *_EsSMTPMgr) GetBatchFromMaxCount(maxCounts []int) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`max_count` IN (?)", maxCounts).Find(&results).Error

	return
}

// GetFromMailFrom 通过mail_from获取内容 发信邮箱
func (obj *_EsSMTPMgr) GetFromMailFrom(mailFrom string) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`mail_from` = ?", mailFrom).Find(&results).Error

	return
}

// GetBatchFromMailFrom 批量查找 发信邮箱
func (obj *_EsSMTPMgr) GetBatchFromMailFrom(mailFroms []string) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`mail_from` IN (?)", mailFroms).Find(&results).Error

	return
}

// GetFromPort 通过port获取内容 端口
func (obj *_EsSMTPMgr) GetFromPort(port int) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`port` = ?", port).Find(&results).Error

	return
}

// GetBatchFromPort 批量查找 端口
func (obj *_EsSMTPMgr) GetBatchFromPort(ports []int) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`port` IN (?)", ports).Find(&results).Error

	return
}

// GetFromOpenSsl 通过open_ssl获取内容 ssl是否开启
func (obj *_EsSMTPMgr) GetFromOpenSsl(openSsl int) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`open_ssl` = ?", openSsl).Find(&results).Error

	return
}

// GetBatchFromOpenSsl 批量查找 ssl是否开启
func (obj *_EsSMTPMgr) GetBatchFromOpenSsl(openSsls []int) (results []*models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`open_ssl` IN (?)", openSsls).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsSMTPMgr) FetchByPrimaryKey(id int) (result models.EsSMTP, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsSMTP{}).Where("`id` = ?", id).First(&result).Error

	return
}
