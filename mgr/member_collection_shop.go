package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type _EsMemberCollectionShopMgr struct {
	*_BaseMgr
}

// EsMemberCollectionShopMgr open func
func EsMemberCollectionShopMgr(db *gorm.DB) *_EsMemberCollectionShopMgr {
	if db == nil {
		panic(fmt.Errorf("EsMemberCollectionShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_EsMemberCollectionShopMgr{_BaseMgr: &_BaseMgr{DB: db.Table("es_member_collection_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_EsMemberCollectionShopMgr) GetTableName() string {
	return "es_member_collection_shop"
}

// Reset 重置gorm会话
func (obj *_EsMemberCollectionShopMgr) Reset() *_EsMemberCollectionShopMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_EsMemberCollectionShopMgr) Get() (result models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_EsMemberCollectionShopMgr) Gets() (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_EsMemberCollectionShopMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 收藏id
func (obj *_EsMemberCollectionShopMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员id
func (obj *_EsMemberCollectionShopMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithShopID shop_id获取 店铺id
func (obj *_EsMemberCollectionShopMgr) WithShopID(shopID int) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithShopName shop_name获取 店铺名称
func (obj *_EsMemberCollectionShopMgr) WithShopName(shopName string) Option {
	return optionFunc(func(o *options) { o.query["shop_name"] = shopName })
}

// WithCreateTime create_time获取 收藏时间
func (obj *_EsMemberCollectionShopMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithLogo logo获取 店铺logo
func (obj *_EsMemberCollectionShopMgr) WithLogo(logo string) Option {
	return optionFunc(func(o *options) { o.query["logo"] = logo })
}

// WithShopProvince shop_province获取 店铺所在省
func (obj *_EsMemberCollectionShopMgr) WithShopProvince(shopProvince string) Option {
	return optionFunc(func(o *options) { o.query["shop_province"] = shopProvince })
}

// WithShopCity shop_city获取 店铺所在市
func (obj *_EsMemberCollectionShopMgr) WithShopCity(shopCity string) Option {
	return optionFunc(func(o *options) { o.query["shop_city"] = shopCity })
}

// WithShopRegion shop_region获取 店铺所在县
func (obj *_EsMemberCollectionShopMgr) WithShopRegion(shopRegion string) Option {
	return optionFunc(func(o *options) { o.query["shop_region"] = shopRegion })
}

// WithShopTown shop_town获取 店铺所在镇
func (obj *_EsMemberCollectionShopMgr) WithShopTown(shopTown string) Option {
	return optionFunc(func(o *options) { o.query["shop_town"] = shopTown })
}

// WithShopPraiseRate shop_praise_rate获取 店铺好评率
func (obj *_EsMemberCollectionShopMgr) WithShopPraiseRate(shopPraiseRate float64) Option {
	return optionFunc(func(o *options) { o.query["shop_praise_rate"] = shopPraiseRate })
}

// WithShopDescriptionCredit shop_description_credit获取 店铺描述相符度
func (obj *_EsMemberCollectionShopMgr) WithShopDescriptionCredit(shopDescriptionCredit float64) Option {
	return optionFunc(func(o *options) { o.query["shop_description_credit"] = shopDescriptionCredit })
}

// WithShopServiceCredit shop_service_credit获取 服务态度分数
func (obj *_EsMemberCollectionShopMgr) WithShopServiceCredit(shopServiceCredit float64) Option {
	return optionFunc(func(o *options) { o.query["shop_service_credit"] = shopServiceCredit })
}

// WithShopDeliveryCredit shop_delivery_credit获取 发货速度分数
func (obj *_EsMemberCollectionShopMgr) WithShopDeliveryCredit(shopDeliveryCredit float64) Option {
	return optionFunc(func(o *options) { o.query["shop_delivery_credit"] = shopDeliveryCredit })
}

// GetByOption 功能选项模式获取
func (obj *_EsMemberCollectionShopMgr) GetByOption(opts ...Option) (result models.EsMemberCollectionShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_EsMemberCollectionShopMgr) GetByOptions(opts ...Option) (results []*models.EsMemberCollectionShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_EsMemberCollectionShopMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsMemberCollectionShop, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where(options.query)
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

// GetFromID 通过id获取内容 收藏id
func (obj *_EsMemberCollectionShopMgr) GetFromID(id int) (result models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 收藏id
func (obj *_EsMemberCollectionShopMgr) GetBatchFromID(ids []int) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *_EsMemberCollectionShopMgr) GetFromMemberID(memberID int) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *_EsMemberCollectionShopMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *_EsMemberCollectionShopMgr) GetFromShopID(shopID int) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *_EsMemberCollectionShopMgr) GetBatchFromShopID(shopIDs []int) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromShopName 通过shop_name获取内容 店铺名称
func (obj *_EsMemberCollectionShopMgr) GetFromShopName(shopName string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_name` = ?", shopName).Find(&results).Error

	return
}

// GetBatchFromShopName 批量查找 店铺名称
func (obj *_EsMemberCollectionShopMgr) GetBatchFromShopName(shopNames []string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_name` IN (?)", shopNames).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 收藏时间
func (obj *_EsMemberCollectionShopMgr) GetFromCreateTime(createTime int64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 收藏时间
func (obj *_EsMemberCollectionShopMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromLogo 通过logo获取内容 店铺logo
func (obj *_EsMemberCollectionShopMgr) GetFromLogo(logo string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`logo` = ?", logo).Find(&results).Error

	return
}

// GetBatchFromLogo 批量查找 店铺logo
func (obj *_EsMemberCollectionShopMgr) GetBatchFromLogo(logos []string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`logo` IN (?)", logos).Find(&results).Error

	return
}

// GetFromShopProvince 通过shop_province获取内容 店铺所在省
func (obj *_EsMemberCollectionShopMgr) GetFromShopProvince(shopProvince string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_province` = ?", shopProvince).Find(&results).Error

	return
}

// GetBatchFromShopProvince 批量查找 店铺所在省
func (obj *_EsMemberCollectionShopMgr) GetBatchFromShopProvince(shopProvinces []string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_province` IN (?)", shopProvinces).Find(&results).Error

	return
}

// GetFromShopCity 通过shop_city获取内容 店铺所在市
func (obj *_EsMemberCollectionShopMgr) GetFromShopCity(shopCity string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_city` = ?", shopCity).Find(&results).Error

	return
}

// GetBatchFromShopCity 批量查找 店铺所在市
func (obj *_EsMemberCollectionShopMgr) GetBatchFromShopCity(shopCitys []string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_city` IN (?)", shopCitys).Find(&results).Error

	return
}

// GetFromShopRegion 通过shop_region获取内容 店铺所在县
func (obj *_EsMemberCollectionShopMgr) GetFromShopRegion(shopRegion string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_region` = ?", shopRegion).Find(&results).Error

	return
}

// GetBatchFromShopRegion 批量查找 店铺所在县
func (obj *_EsMemberCollectionShopMgr) GetBatchFromShopRegion(shopRegions []string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_region` IN (?)", shopRegions).Find(&results).Error

	return
}

// GetFromShopTown 通过shop_town获取内容 店铺所在镇
func (obj *_EsMemberCollectionShopMgr) GetFromShopTown(shopTown string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_town` = ?", shopTown).Find(&results).Error

	return
}

// GetBatchFromShopTown 批量查找 店铺所在镇
func (obj *_EsMemberCollectionShopMgr) GetBatchFromShopTown(shopTowns []string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_town` IN (?)", shopTowns).Find(&results).Error

	return
}

// GetFromShopPraiseRate 通过shop_praise_rate获取内容 店铺好评率
func (obj *_EsMemberCollectionShopMgr) GetFromShopPraiseRate(shopPraiseRate float64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_praise_rate` = ?", shopPraiseRate).Find(&results).Error

	return
}

// GetBatchFromShopPraiseRate 批量查找 店铺好评率
func (obj *_EsMemberCollectionShopMgr) GetBatchFromShopPraiseRate(shopPraiseRates []float64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_praise_rate` IN (?)", shopPraiseRates).Find(&results).Error

	return
}

// GetFromShopDescriptionCredit 通过shop_description_credit获取内容 店铺描述相符度
func (obj *_EsMemberCollectionShopMgr) GetFromShopDescriptionCredit(shopDescriptionCredit float64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_description_credit` = ?", shopDescriptionCredit).Find(&results).Error

	return
}

// GetBatchFromShopDescriptionCredit 批量查找 店铺描述相符度
func (obj *_EsMemberCollectionShopMgr) GetBatchFromShopDescriptionCredit(shopDescriptionCredits []float64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_description_credit` IN (?)", shopDescriptionCredits).Find(&results).Error

	return
}

// GetFromShopServiceCredit 通过shop_service_credit获取内容 服务态度分数
func (obj *_EsMemberCollectionShopMgr) GetFromShopServiceCredit(shopServiceCredit float64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_service_credit` = ?", shopServiceCredit).Find(&results).Error

	return
}

// GetBatchFromShopServiceCredit 批量查找 服务态度分数
func (obj *_EsMemberCollectionShopMgr) GetBatchFromShopServiceCredit(shopServiceCredits []float64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_service_credit` IN (?)", shopServiceCredits).Find(&results).Error

	return
}

// GetFromShopDeliveryCredit 通过shop_delivery_credit获取内容 发货速度分数
func (obj *_EsMemberCollectionShopMgr) GetFromShopDeliveryCredit(shopDeliveryCredit float64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_delivery_credit` = ?", shopDeliveryCredit).Find(&results).Error

	return
}

// GetBatchFromShopDeliveryCredit 批量查找 发货速度分数
func (obj *_EsMemberCollectionShopMgr) GetBatchFromShopDeliveryCredit(shopDeliveryCredits []float64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_delivery_credit` IN (?)", shopDeliveryCredits).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_EsMemberCollectionShopMgr) FetchByPrimaryKey(id int) (result models.EsMemberCollectionShop, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`id` = ?", id).First(&result).Error

	return
}
