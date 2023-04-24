package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type ShopNoticeLogMgr struct {
	*_BaseMgr
}

// NewShopNoticeLogMgr open func
func NewShopNoticeLogMgr(db db.Repo) *ShopNoticeLogMgr {
	if db == nil {
		panic(fmt.Errorf("NewShopNoticeLogMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &ShopNoticeLogMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_shop_notice_log"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *ShopNoticeLogMgr) GetTableName() string {
	return "es_shop_notice_log"
}

// Reset 重置gorm会话
func (obj *ShopNoticeLogMgr) Reset() *ShopNoticeLogMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *ShopNoticeLogMgr) Get() (result models.EsShopNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *ShopNoticeLogMgr) Gets() (results []*models.EsShopNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *ShopNoticeLogMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 id
func (obj *ShopNoticeLogMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithShopID shop_id获取 店铺ID
func (obj *ShopNoticeLogMgr) WithShopID(shopID int) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithNoticeContent notice_content获取 站内信内容
func (obj *ShopNoticeLogMgr) WithNoticeContent(noticeContent string) Option {
	return optionFunc(func(o *options) { o.query["notice_content"] = noticeContent })
}

// WithSendTime send_time获取 发送时间
func (obj *ShopNoticeLogMgr) WithSendTime(sendTime int64) Option {
	return optionFunc(func(o *options) { o.query["send_time"] = sendTime })
}

// WithIsDelete is_delete获取 是否删除 ：1 删除   0  未删除
func (obj *ShopNoticeLogMgr) WithIsDelete(isDelete int) Option {
	return optionFunc(func(o *options) { o.query["is_delete"] = isDelete })
}

// WithIsRead is_read获取 是否已读 ：1已读   0 未读
func (obj *ShopNoticeLogMgr) WithIsRead(isRead int) Option {
	return optionFunc(func(o *options) { o.query["is_read"] = isRead })
}

// WithType type获取 消息类型
func (obj *ShopNoticeLogMgr) WithType(_type string) Option {
	return optionFunc(func(o *options) { o.query["type"] = _type })
}

// GetByOption 功能选项模式获取
func (obj *ShopNoticeLogMgr) GetByOption(opts ...Option) (result models.EsShopNoticeLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *ShopNoticeLogMgr) GetByOptions(opts ...Option) (results []*models.EsShopNoticeLog, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *ShopNoticeLogMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsShopNoticeLog, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Where(options.query)
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
func (obj *ShopNoticeLogMgr) GetFromID(id int) (result models.EsShopNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 id
func (obj *ShopNoticeLogMgr) GetBatchFromID(ids []int) (results []*models.EsShopNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺ID
func (obj *ShopNoticeLogMgr) GetFromShopID(shopID int) (results []*models.EsShopNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺ID
func (obj *ShopNoticeLogMgr) GetBatchFromShopID(shopIDs []int) (results []*models.EsShopNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromNoticeContent 通过notice_content获取内容 站内信内容
func (obj *ShopNoticeLogMgr) GetFromNoticeContent(noticeContent string) (results []*models.EsShopNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Where("`notice_content` = ?", noticeContent).Find(&results).Error

	return
}

// GetBatchFromNoticeContent 批量查找 站内信内容
func (obj *ShopNoticeLogMgr) GetBatchFromNoticeContent(noticeContents []string) (results []*models.EsShopNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Where("`notice_content` IN (?)", noticeContents).Find(&results).Error

	return
}

// GetFromSendTime 通过send_time获取内容 发送时间
func (obj *ShopNoticeLogMgr) GetFromSendTime(sendTime int64) (results []*models.EsShopNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Where("`send_time` = ?", sendTime).Find(&results).Error

	return
}

// GetBatchFromSendTime 批量查找 发送时间
func (obj *ShopNoticeLogMgr) GetBatchFromSendTime(sendTimes []int64) (results []*models.EsShopNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Where("`send_time` IN (?)", sendTimes).Find(&results).Error

	return
}

// GetFromIsDelete 通过is_delete获取内容 是否删除 ：1 删除   0  未删除
func (obj *ShopNoticeLogMgr) GetFromIsDelete(isDelete int) (results []*models.EsShopNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Where("`is_delete` = ?", isDelete).Find(&results).Error

	return
}

// GetBatchFromIsDelete 批量查找 是否删除 ：1 删除   0  未删除
func (obj *ShopNoticeLogMgr) GetBatchFromIsDelete(isDeletes []int) (results []*models.EsShopNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Where("`is_delete` IN (?)", isDeletes).Find(&results).Error

	return
}

// GetFromIsRead 通过is_read获取内容 是否已读 ：1已读   0 未读
func (obj *ShopNoticeLogMgr) GetFromIsRead(isRead int) (results []*models.EsShopNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Where("`is_read` = ?", isRead).Find(&results).Error

	return
}

// GetBatchFromIsRead 批量查找 是否已读 ：1已读   0 未读
func (obj *ShopNoticeLogMgr) GetBatchFromIsRead(isReads []int) (results []*models.EsShopNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Where("`is_read` IN (?)", isReads).Find(&results).Error

	return
}

// GetFromType 通过type获取内容 消息类型
func (obj *ShopNoticeLogMgr) GetFromType(_type string) (results []*models.EsShopNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Where("`type` = ?", _type).Find(&results).Error

	return
}

// GetBatchFromType 批量查找 消息类型
func (obj *ShopNoticeLogMgr) GetBatchFromType(_types []string) (results []*models.EsShopNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Where("`type` IN (?)", _types).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *ShopNoticeLogMgr) FetchByPrimaryKey(id int) (result models.EsShopNoticeLog, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopNoticeLog{}).Where("`id` = ?", id).First(&result).Error

	return
}
