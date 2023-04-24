package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type MemberCollectionShopMgr struct {
	*_BaseMgr
}

// NewMemberCollectionShopMgr open func
func NewMemberCollectionShopMgr(db db.Repo) *MemberCollectionShopMgr {
	if db == nil {
		panic(fmt.Errorf("NewMemberCollectionShopMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &MemberCollectionShopMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_member_collection_shop"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *MemberCollectionShopMgr) GetTableName() string {
	return "es_member_collection_shop"
}

// Reset 重置gorm会话
func (obj *MemberCollectionShopMgr) Reset() *MemberCollectionShopMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *MemberCollectionShopMgr) Get() (result models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *MemberCollectionShopMgr) Gets() (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *MemberCollectionShopMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 收藏id
func (obj *MemberCollectionShopMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithMemberID member_id获取 会员id
func (obj *MemberCollectionShopMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithShopID shop_id获取 店铺id
func (obj *MemberCollectionShopMgr) WithShopID(shopID int) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithShopName shop_name获取 店铺名称
func (obj *MemberCollectionShopMgr) WithShopName(shopName string) Option {
	return optionFunc(func(o *options) { o.query["shop_name"] = shopName })
}

// WithCreateTime create_time获取 收藏时间
func (obj *MemberCollectionShopMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithLogo logo获取 店铺logo
func (obj *MemberCollectionShopMgr) WithLogo(logo string) Option {
	return optionFunc(func(o *options) { o.query["logo"] = logo })
}

// WithShopProvince shop_province获取 店铺所在省
func (obj *MemberCollectionShopMgr) WithShopProvince(shopProvince string) Option {
	return optionFunc(func(o *options) { o.query["shop_province"] = shopProvince })
}

// WithShopCity shop_city获取 店铺所在市
func (obj *MemberCollectionShopMgr) WithShopCity(shopCity string) Option {
	return optionFunc(func(o *options) { o.query["shop_city"] = shopCity })
}

// WithShopRegion shop_region获取 店铺所在县
func (obj *MemberCollectionShopMgr) WithShopRegion(shopRegion string) Option {
	return optionFunc(func(o *options) { o.query["shop_region"] = shopRegion })
}

// WithShopTown shop_town获取 店铺所在镇
func (obj *MemberCollectionShopMgr) WithShopTown(shopTown string) Option {
	return optionFunc(func(o *options) { o.query["shop_town"] = shopTown })
}

// WithShopPraiseRate shop_praise_rate获取 店铺好评率
func (obj *MemberCollectionShopMgr) WithShopPraiseRate(shopPraiseRate float64) Option {
	return optionFunc(func(o *options) { o.query["shop_praise_rate"] = shopPraiseRate })
}

// WithShopDescriptionCredit shop_description_credit获取 店铺描述相符度
func (obj *MemberCollectionShopMgr) WithShopDescriptionCredit(shopDescriptionCredit float64) Option {
	return optionFunc(func(o *options) { o.query["shop_description_credit"] = shopDescriptionCredit })
}

// WithShopServiceCredit shop_service_credit获取 服务态度分数
func (obj *MemberCollectionShopMgr) WithShopServiceCredit(shopServiceCredit float64) Option {
	return optionFunc(func(o *options) { o.query["shop_service_credit"] = shopServiceCredit })
}

// WithShopDeliveryCredit shop_delivery_credit获取 发货速度分数
func (obj *MemberCollectionShopMgr) WithShopDeliveryCredit(shopDeliveryCredit float64) Option {
	return optionFunc(func(o *options) { o.query["shop_delivery_credit"] = shopDeliveryCredit })
}

// GetByOption 功能选项模式获取
func (obj *MemberCollectionShopMgr) GetByOption(opts ...Option) (result models.EsMemberCollectionShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *MemberCollectionShopMgr) GetByOptions(opts ...Option) (results []*models.EsMemberCollectionShop, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *MemberCollectionShopMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsMemberCollectionShop, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where(options.query)
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
func (obj *MemberCollectionShopMgr) GetFromID(id int) (result models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 收藏id
func (obj *MemberCollectionShopMgr) GetBatchFromID(ids []int) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromMemberID 通过member_id获取内容 会员id
func (obj *MemberCollectionShopMgr) GetFromMemberID(memberID int) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`member_id` = ?", memberID).Find(&results).Error

	return
}

// GetBatchFromMemberID 批量查找 会员id
func (obj *MemberCollectionShopMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *MemberCollectionShopMgr) GetFromShopID(shopID int) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *MemberCollectionShopMgr) GetBatchFromShopID(shopIDs []int) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromShopName 通过shop_name获取内容 店铺名称
func (obj *MemberCollectionShopMgr) GetFromShopName(shopName string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_name` = ?", shopName).Find(&results).Error

	return
}

// GetBatchFromShopName 批量查找 店铺名称
func (obj *MemberCollectionShopMgr) GetBatchFromShopName(shopNames []string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_name` IN (?)", shopNames).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 收藏时间
func (obj *MemberCollectionShopMgr) GetFromCreateTime(createTime int64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 收藏时间
func (obj *MemberCollectionShopMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromLogo 通过logo获取内容 店铺logo
func (obj *MemberCollectionShopMgr) GetFromLogo(logo string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`logo` = ?", logo).Find(&results).Error

	return
}

// GetBatchFromLogo 批量查找 店铺logo
func (obj *MemberCollectionShopMgr) GetBatchFromLogo(logos []string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`logo` IN (?)", logos).Find(&results).Error

	return
}

// GetFromShopProvince 通过shop_province获取内容 店铺所在省
func (obj *MemberCollectionShopMgr) GetFromShopProvince(shopProvince string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_province` = ?", shopProvince).Find(&results).Error

	return
}

// GetBatchFromShopProvince 批量查找 店铺所在省
func (obj *MemberCollectionShopMgr) GetBatchFromShopProvince(shopProvinces []string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_province` IN (?)", shopProvinces).Find(&results).Error

	return
}

// GetFromShopCity 通过shop_city获取内容 店铺所在市
func (obj *MemberCollectionShopMgr) GetFromShopCity(shopCity string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_city` = ?", shopCity).Find(&results).Error

	return
}

// GetBatchFromShopCity 批量查找 店铺所在市
func (obj *MemberCollectionShopMgr) GetBatchFromShopCity(shopCitys []string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_city` IN (?)", shopCitys).Find(&results).Error

	return
}

// GetFromShopRegion 通过shop_region获取内容 店铺所在县
func (obj *MemberCollectionShopMgr) GetFromShopRegion(shopRegion string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_region` = ?", shopRegion).Find(&results).Error

	return
}

// GetBatchFromShopRegion 批量查找 店铺所在县
func (obj *MemberCollectionShopMgr) GetBatchFromShopRegion(shopRegions []string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_region` IN (?)", shopRegions).Find(&results).Error

	return
}

// GetFromShopTown 通过shop_town获取内容 店铺所在镇
func (obj *MemberCollectionShopMgr) GetFromShopTown(shopTown string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_town` = ?", shopTown).Find(&results).Error

	return
}

// GetBatchFromShopTown 批量查找 店铺所在镇
func (obj *MemberCollectionShopMgr) GetBatchFromShopTown(shopTowns []string) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_town` IN (?)", shopTowns).Find(&results).Error

	return
}

// GetFromShopPraiseRate 通过shop_praise_rate获取内容 店铺好评率
func (obj *MemberCollectionShopMgr) GetFromShopPraiseRate(shopPraiseRate float64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_praise_rate` = ?", shopPraiseRate).Find(&results).Error

	return
}

// GetBatchFromShopPraiseRate 批量查找 店铺好评率
func (obj *MemberCollectionShopMgr) GetBatchFromShopPraiseRate(shopPraiseRates []float64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_praise_rate` IN (?)", shopPraiseRates).Find(&results).Error

	return
}

// GetFromShopDescriptionCredit 通过shop_description_credit获取内容 店铺描述相符度
func (obj *MemberCollectionShopMgr) GetFromShopDescriptionCredit(shopDescriptionCredit float64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_description_credit` = ?", shopDescriptionCredit).Find(&results).Error

	return
}

// GetBatchFromShopDescriptionCredit 批量查找 店铺描述相符度
func (obj *MemberCollectionShopMgr) GetBatchFromShopDescriptionCredit(shopDescriptionCredits []float64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_description_credit` IN (?)", shopDescriptionCredits).Find(&results).Error

	return
}

// GetFromShopServiceCredit 通过shop_service_credit获取内容 服务态度分数
func (obj *MemberCollectionShopMgr) GetFromShopServiceCredit(shopServiceCredit float64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_service_credit` = ?", shopServiceCredit).Find(&results).Error

	return
}

// GetBatchFromShopServiceCredit 批量查找 服务态度分数
func (obj *MemberCollectionShopMgr) GetBatchFromShopServiceCredit(shopServiceCredits []float64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_service_credit` IN (?)", shopServiceCredits).Find(&results).Error

	return
}

// GetFromShopDeliveryCredit 通过shop_delivery_credit获取内容 发货速度分数
func (obj *MemberCollectionShopMgr) GetFromShopDeliveryCredit(shopDeliveryCredit float64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_delivery_credit` = ?", shopDeliveryCredit).Find(&results).Error

	return
}

// GetBatchFromShopDeliveryCredit 批量查找 发货速度分数
func (obj *MemberCollectionShopMgr) GetBatchFromShopDeliveryCredit(shopDeliveryCredits []float64) (results []*models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`shop_delivery_credit` IN (?)", shopDeliveryCredits).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *MemberCollectionShopMgr) FetchByPrimaryKey(id int) (result models.EsMemberCollectionShop, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMemberCollectionShop{}).Where("`id` = ?", id).First(&result).Error

	return
}
