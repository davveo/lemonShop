package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsHistoryMgr struct {
	*_BaseMgr
}

// EsHistoryMgr open func
func EsHistoryMgr(db *gorm.DB) *_EsHistoryMgr {
	if db == nil {
		panic(fmt.Errorf("EsHistoryMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsHistoryMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_history"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsHistoryMgr) GetTableName() string {
	return "es_history"
}

// Reset 重置gorm会话
func (obj *_EsHistoryMgr) Reset() *_EsHistoryMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsHistoryMgr) Get() (result models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsHistoryMgr) Gets() (results []*models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsHistoryMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 主键
func (obj *_EsHistoryMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithGoodsID goods_id获取 商品id
func (obj *_EsHistoryMgr) WithGoodsID(goodsID int) Option {
	return optionFunc(func(o *options) { o.query["goods_id"] = goodsID })
}

// WithGoodsName goods_name获取 商品名称
func (obj *_EsHistoryMgr) WithGoodsName(goodsName string) Option {
	return optionFunc(func(o *options) { o.query["goods_name"] = goodsName })
}

// WithGoodsPrice goods_price获取 商品价格
func (obj *_EsHistoryMgr) WithGoodsPrice(goodsPrice float64) Option {
	return optionFunc(func(o *options) { o.query["goods_price"] = goodsPrice })
}

// WithGoodsImg goods_img获取 商品主图
func (obj *_EsHistoryMgr) WithGoodsImg(goodsImg string) Option {
	return optionFunc(func(o *options) { o.query["goods_img"] = goodsImg })
}

// WithMemberID member_id获取 会员id
func (obj *_EsHistoryMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithMemberName member_name获取 会员名称
func (obj *_EsHistoryMgr) WithMemberName(memberName string) Option {
	return optionFunc(func(o *options) { o.query["member_name"] = memberName })
}

// WithCreateTime create_time获取 创建时间，按天存
func (obj *_EsHistoryMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithUpdateTime update_time获取 更新时间
func (obj *_EsHistoryMgr) WithUpdateTime(updateTime int64) Option {
	return optionFunc(func(o *options) { o.query["update_time"] = updateTime })
}

// GetByOption 功能选项模式获取
func (obj *_EsHistoryMgr) GetByOption(opts ...Option) (result models.EsHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsHistoryMgr) GetByOptions(opts ...Option) (results []*models.EsHistory, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsHistoryMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsHistory, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where(options.query)
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

// GetFromID 通过id获取内容 主键
func (obj *_EsHistoryMgr) GetFromID(id int) (result models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 主键
func (obj *_EsHistoryMgr) GetBatchFromID(ids []int) (results []*models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromGoodsID 通过goods_id获取内容 商品id
func (obj *_EsHistoryMgr) GetFromGoodsID(goodsID int) (results []*models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`goods_id` = ?", goodsID).Find(&results).Error

	return
}

// GetBatchFromGoodsID 批量查找 商品id
func (obj *_EsHistoryMgr) GetBatchFromGoodsID(goodsIDs []int) (results []*models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`goods_id` IN (?)", goodsIDs).Find(&results).Error

	return
}

// GetFromGoodsName 通过goods_name获取内容 商品名称
func (obj *_EsHistoryMgr) GetFromGoodsName(goodsName string) (results []*models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`goods_name` = ?", goodsName).Find(&results).Error

	return
}

// GetBatchFromGoodsName 批量查找 商品名称
func (obj *_EsHistoryMgr) GetBatchFromGoodsName(goodsNames []string) (results []*models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`goods_name` IN (?)", goodsNames).Find(&results).Error

	return
}

// GetFromGoodsPrice 通过goods_price获取内容 商品价格
func (obj *_EsHistoryMgr) GetFromGoodsPrice(goodsPrice float64) (results []*models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`goods_price` = ?", goodsPrice).Find(&results).Error

	return
}

// GetBatchFromGoodsPrice 批量查找 商品价格
func (obj *_EsHistoryMgr) GetBatchFromGoodsPrice(goodsPrices []float64) (results []*models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`goods_price` IN (?)", goodsPrices).Find(&results).Error

	return
}

// GetFromGoodsImg 通过goods_img获取内容 商品主图
func (obj *_EsHistoryMgr) GetFromGoodsImg(goodsImg string) (results []*models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`goods_img` = ?", goodsImg).Find(&results).Error

	return
}

// GetBatchFromGoodsImg 批量查找 商品主图
func (obj *_EsHistoryMgr) GetBatchFromGoodsImg(goodsImgs []string) (results []*models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`goods_img` IN (?)", goodsImgs).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_EsHistoryMgr) GetFromMemberID(memberID int) (results []*models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_EsHistoryMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromMemberName 通过member_name获取内容 会员名称
func (obj *_EsHistoryMgr) GetFromMemberName(memberName string) (results []*models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`member_name` = ?", memberName).Find(&results).Error

	return
}

// GetBatchFromMemberName 批量查找 会员名称
func (obj *_EsHistoryMgr) GetBatchFromMemberName(memberNames []string) (results []*models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`member_name` IN (?)", memberNames).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 创建时间，按天存
func (obj *_EsHistoryMgr) GetFromCreateTime(createTime int64) (results []*models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 创建时间，按天存
func (obj *_EsHistoryMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromUpdateTime 通过update_time获取内容 更新时间
func (obj *_EsHistoryMgr) GetFromUpdateTime(updateTime int64) (results []*models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`update_time` = ?", updateTime).Find(&results).Error

	return
}

// GetBatchFromUpdateTime 批量查找 更新时间
func (obj *_EsHistoryMgr) GetBatchFromUpdateTime(updateTimes []int64) (results []*models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`update_time` IN (?)", updateTimes).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsHistoryMgr) FetchByPrimaryKey(id int) (result models.EsHistory, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsHistory{}).Where("`id` = ?", id).First(&result).Error

	return
}
