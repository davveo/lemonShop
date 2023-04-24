package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"

	"gorm.io/gorm"
)

type ShopDetailMgr struct {
	*_BaseMgr
}

// NewShopDetailMgr open func
func NewShopDetailMgr(db db.Repo) *ShopDetailMgr {
	if db == nil {
		panic(fmt.Errorf("NewShopDetailMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &ShopDetailMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_shop_detail"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *ShopDetailMgr) GetTableName() string {
	return "es_shop_detail"
}

// Reset 重置gorm会话
func (obj *ShopDetailMgr) Reset() *ShopDetailMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *ShopDetailMgr) Get() (result models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *ShopDetailMgr) Gets() (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *ShopDetailMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取 店铺详细id
func (obj *ShopDetailMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithShopID shop_id获取 店铺id
func (obj *ShopDetailMgr) WithShopID(shopID int) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithShopProvinceID shop_province_id获取 店铺所在省id
func (obj *ShopDetailMgr) WithShopProvinceID(shopProvinceID int) Option {
	return optionFunc(func(o *options) { o.query["shop_province_id"] = shopProvinceID })
}

// WithShopCityID shop_city_id获取 店铺所在市id
func (obj *ShopDetailMgr) WithShopCityID(shopCityID int) Option {
	return optionFunc(func(o *options) { o.query["shop_city_id"] = shopCityID })
}

// WithShopCountyID shop_county_id获取 店铺所在县id
func (obj *ShopDetailMgr) WithShopCountyID(shopCountyID int) Option {
	return optionFunc(func(o *options) { o.query["shop_county_id"] = shopCountyID })
}

// WithShopTownID shop_town_id获取 店铺所在镇id
func (obj *ShopDetailMgr) WithShopTownID(shopTownID int) Option {
	return optionFunc(func(o *options) { o.query["shop_town_id"] = shopTownID })
}

// WithShopProvince shop_province获取 店铺所在省
func (obj *ShopDetailMgr) WithShopProvince(shopProvince string) Option {
	return optionFunc(func(o *options) { o.query["shop_province"] = shopProvince })
}

// WithShopCity shop_city获取 店铺所在市
func (obj *ShopDetailMgr) WithShopCity(shopCity string) Option {
	return optionFunc(func(o *options) { o.query["shop_city"] = shopCity })
}

// WithShopCounty shop_county获取 店铺所在县
func (obj *ShopDetailMgr) WithShopCounty(shopCounty string) Option {
	return optionFunc(func(o *options) { o.query["shop_county"] = shopCounty })
}

// WithShopTown shop_town获取 店铺所在镇
func (obj *ShopDetailMgr) WithShopTown(shopTown string) Option {
	return optionFunc(func(o *options) { o.query["shop_town"] = shopTown })
}

// WithShopAdd shop_add获取 店铺详细地址
func (obj *ShopDetailMgr) WithShopAdd(shopAdd string) Option {
	return optionFunc(func(o *options) { o.query["shop_add"] = shopAdd })
}

// WithCompanyName company_name获取 公司名称
func (obj *ShopDetailMgr) WithCompanyName(companyName string) Option {
	return optionFunc(func(o *options) { o.query["company_name"] = companyName })
}

// WithCompanyAddress company_address获取 公司地址
func (obj *ShopDetailMgr) WithCompanyAddress(companyAddress string) Option {
	return optionFunc(func(o *options) { o.query["company_address"] = companyAddress })
}

// WithCompanyPhone company_phone获取 公司电话
func (obj *ShopDetailMgr) WithCompanyPhone(companyPhone string) Option {
	return optionFunc(func(o *options) { o.query["company_phone"] = companyPhone })
}

// WithCompanyEmail company_email获取 电子邮箱
func (obj *ShopDetailMgr) WithCompanyEmail(companyEmail string) Option {
	return optionFunc(func(o *options) { o.query["company_email"] = companyEmail })
}

// WithEmployeeNum employee_num获取 员工总数
func (obj *ShopDetailMgr) WithEmployeeNum(employeeNum int) Option {
	return optionFunc(func(o *options) { o.query["employee_num"] = employeeNum })
}

// WithRegMoney reg_money获取 注册资金
func (obj *ShopDetailMgr) WithRegMoney(regMoney float64) Option {
	return optionFunc(func(o *options) { o.query["reg_money"] = regMoney })
}

// WithLinkName link_name获取 联系人姓名
func (obj *ShopDetailMgr) WithLinkName(linkName string) Option {
	return optionFunc(func(o *options) { o.query["link_name"] = linkName })
}

// WithLinkPhone link_phone获取 联系人电话
func (obj *ShopDetailMgr) WithLinkPhone(linkPhone string) Option {
	return optionFunc(func(o *options) { o.query["link_phone"] = linkPhone })
}

// WithLegalName legal_name获取 法人姓名
func (obj *ShopDetailMgr) WithLegalName(legalName string) Option {
	return optionFunc(func(o *options) { o.query["legal_name"] = legalName })
}

// WithLegalID legal_id获取 法人身份证
func (obj *ShopDetailMgr) WithLegalID(legalID string) Option {
	return optionFunc(func(o *options) { o.query["legal_id"] = legalID })
}

// WithLegalImg legal_img获取 法人身份证照片
func (obj *ShopDetailMgr) WithLegalImg(legalImg string) Option {
	return optionFunc(func(o *options) { o.query["legal_img"] = legalImg })
}

// WithLicenseNum license_num获取 营业执照号
func (obj *ShopDetailMgr) WithLicenseNum(licenseNum string) Option {
	return optionFunc(func(o *options) { o.query["license_num"] = licenseNum })
}

// WithLicenseProvinceID license_province_id获取 营业执照所在省id
func (obj *ShopDetailMgr) WithLicenseProvinceID(licenseProvinceID int) Option {
	return optionFunc(func(o *options) { o.query["license_province_id"] = licenseProvinceID })
}

// WithLicenseCityID license_city_id获取 营业执照所在市id
func (obj *ShopDetailMgr) WithLicenseCityID(licenseCityID int) Option {
	return optionFunc(func(o *options) { o.query["license_city_id"] = licenseCityID })
}

// WithLicenseCountyID license_county_id获取 营业执照所在县id
func (obj *ShopDetailMgr) WithLicenseCountyID(licenseCountyID int) Option {
	return optionFunc(func(o *options) { o.query["license_county_id"] = licenseCountyID })
}

// WithLicenseTownID license_town_id获取 营业执照所在镇id
func (obj *ShopDetailMgr) WithLicenseTownID(licenseTownID int) Option {
	return optionFunc(func(o *options) { o.query["license_town_id"] = licenseTownID })
}

// WithLicenseProvince license_province获取 营业执照所在省
func (obj *ShopDetailMgr) WithLicenseProvince(licenseProvince string) Option {
	return optionFunc(func(o *options) { o.query["license_province"] = licenseProvince })
}

// WithLicenseCity license_city获取 营业执照所在市
func (obj *ShopDetailMgr) WithLicenseCity(licenseCity string) Option {
	return optionFunc(func(o *options) { o.query["license_city"] = licenseCity })
}

// WithLicenseCounty license_county获取 营业执照所在县
func (obj *ShopDetailMgr) WithLicenseCounty(licenseCounty string) Option {
	return optionFunc(func(o *options) { o.query["license_county"] = licenseCounty })
}

// WithLicenseTown license_town获取 营业执照所在镇
func (obj *ShopDetailMgr) WithLicenseTown(licenseTown string) Option {
	return optionFunc(func(o *options) { o.query["license_town"] = licenseTown })
}

// WithLicenseAdd license_add获取 营业执照详细地址
func (obj *ShopDetailMgr) WithLicenseAdd(licenseAdd string) Option {
	return optionFunc(func(o *options) { o.query["license_add"] = licenseAdd })
}

// WithEstablishDate establish_date获取 成立日期
func (obj *ShopDetailMgr) WithEstablishDate(establishDate int64) Option {
	return optionFunc(func(o *options) { o.query["establish_date"] = establishDate })
}

// WithLicenceStart licence_start获取 营业执照有效期开始
func (obj *ShopDetailMgr) WithLicenceStart(licenceStart int64) Option {
	return optionFunc(func(o *options) { o.query["licence_start"] = licenceStart })
}

// WithLicenceEnd licence_end获取 营业执照有效期结束
func (obj *ShopDetailMgr) WithLicenceEnd(licenceEnd int64) Option {
	return optionFunc(func(o *options) { o.query["licence_end"] = licenceEnd })
}

// WithScope scope获取 法定经营范围
func (obj *ShopDetailMgr) WithScope(scope string) Option {
	return optionFunc(func(o *options) { o.query["scope"] = scope })
}

// WithLicenceImg licence_img获取 营业执照电子版
func (obj *ShopDetailMgr) WithLicenceImg(licenceImg string) Option {
	return optionFunc(func(o *options) { o.query["licence_img"] = licenceImg })
}

// WithOrganizationCode organization_code获取 组织机构代码
func (obj *ShopDetailMgr) WithOrganizationCode(organizationCode string) Option {
	return optionFunc(func(o *options) { o.query["organization_code"] = organizationCode })
}

// WithCodeImg code_img获取 组织机构电子版
func (obj *ShopDetailMgr) WithCodeImg(codeImg string) Option {
	return optionFunc(func(o *options) { o.query["code_img"] = codeImg })
}

// WithTaxesImg taxes_img获取 一般纳税人证明电子版
func (obj *ShopDetailMgr) WithTaxesImg(taxesImg string) Option {
	return optionFunc(func(o *options) { o.query["taxes_img"] = taxesImg })
}

// WithBankAccountName bank_account_name获取 银行开户名
func (obj *ShopDetailMgr) WithBankAccountName(bankAccountName string) Option {
	return optionFunc(func(o *options) { o.query["bank_account_name"] = bankAccountName })
}

// WithBankNumber bank_number获取 银行开户账号
func (obj *ShopDetailMgr) WithBankNumber(bankNumber string) Option {
	return optionFunc(func(o *options) { o.query["bank_number"] = bankNumber })
}

// WithBankName bank_name获取 开户银行支行名称
func (obj *ShopDetailMgr) WithBankName(bankName string) Option {
	return optionFunc(func(o *options) { o.query["bank_name"] = bankName })
}

// WithBankProvinceID bank_province_id获取 开户银行所在省id
func (obj *ShopDetailMgr) WithBankProvinceID(bankProvinceID int) Option {
	return optionFunc(func(o *options) { o.query["bank_province_id"] = bankProvinceID })
}

// WithBankCityID bank_city_id获取 开户银行所在市id
func (obj *ShopDetailMgr) WithBankCityID(bankCityID int) Option {
	return optionFunc(func(o *options) { o.query["bank_city_id"] = bankCityID })
}

// WithBankCountyID bank_county_id获取 开户银行所在县id
func (obj *ShopDetailMgr) WithBankCountyID(bankCountyID int) Option {
	return optionFunc(func(o *options) { o.query["bank_county_id"] = bankCountyID })
}

// WithBankTownID bank_town_id获取 开户银行所在镇id
func (obj *ShopDetailMgr) WithBankTownID(bankTownID int) Option {
	return optionFunc(func(o *options) { o.query["bank_town_id"] = bankTownID })
}

// WithBankProvince bank_province获取 开户银行所在省
func (obj *ShopDetailMgr) WithBankProvince(bankProvince string) Option {
	return optionFunc(func(o *options) { o.query["bank_province"] = bankProvince })
}

// WithBankCity bank_city获取 开户银行所在市
func (obj *ShopDetailMgr) WithBankCity(bankCity string) Option {
	return optionFunc(func(o *options) { o.query["bank_city"] = bankCity })
}

// WithBankCounty bank_county获取 开户银行所在县
func (obj *ShopDetailMgr) WithBankCounty(bankCounty string) Option {
	return optionFunc(func(o *options) { o.query["bank_county"] = bankCounty })
}

// WithBankTown bank_town获取 开户银行所在镇
func (obj *ShopDetailMgr) WithBankTown(bankTown string) Option {
	return optionFunc(func(o *options) { o.query["bank_town"] = bankTown })
}

// WithBankImg bank_img获取 开户银行许可证电子版
func (obj *ShopDetailMgr) WithBankImg(bankImg string) Option {
	return optionFunc(func(o *options) { o.query["bank_img"] = bankImg })
}

// WithTaxesCertificateNum taxes_certificate_num获取 税务登记证号
func (obj *ShopDetailMgr) WithTaxesCertificateNum(taxesCertificateNum string) Option {
	return optionFunc(func(o *options) { o.query["taxes_certificate_num"] = taxesCertificateNum })
}

// WithTaxesDistinguishNum taxes_distinguish_num获取 纳税人识别号
func (obj *ShopDetailMgr) WithTaxesDistinguishNum(taxesDistinguishNum string) Option {
	return optionFunc(func(o *options) { o.query["taxes_distinguish_num"] = taxesDistinguishNum })
}

// WithTaxesCertificateImg taxes_certificate_img获取 税务登记证书
func (obj *ShopDetailMgr) WithTaxesCertificateImg(taxesCertificateImg string) Option {
	return optionFunc(func(o *options) { o.query["taxes_certificate_img"] = taxesCertificateImg })
}

// WithGoodsManagementCategory goods_management_category获取 店铺经营类目
func (obj *ShopDetailMgr) WithGoodsManagementCategory(goodsManagementCategory string) Option {
	return optionFunc(func(o *options) { o.query["goods_management_category"] = goodsManagementCategory })
}

// WithShopLevel shop_level获取 店铺等级
func (obj *ShopDetailMgr) WithShopLevel(shopLevel int) Option {
	return optionFunc(func(o *options) { o.query["shop_level"] = shopLevel })
}

// WithShopLevelApply shop_level_apply获取 店铺等级申请
func (obj *ShopDetailMgr) WithShopLevelApply(shopLevelApply int) Option {
	return optionFunc(func(o *options) { o.query["shop_level_apply"] = shopLevelApply })
}

// WithStoreSpaceCapacity store_space_capacity获取 店铺相册已用存储量
func (obj *ShopDetailMgr) WithStoreSpaceCapacity(storeSpaceCapacity float64) Option {
	return optionFunc(func(o *options) { o.query["store_space_capacity"] = storeSpaceCapacity })
}

// WithShopLogo shop_logo获取 店铺logo
func (obj *ShopDetailMgr) WithShopLogo(shopLogo string) Option {
	return optionFunc(func(o *options) { o.query["shop_logo"] = shopLogo })
}

// WithShopBanner shop_banner获取 店铺横幅
func (obj *ShopDetailMgr) WithShopBanner(shopBanner string) Option {
	return optionFunc(func(o *options) { o.query["shop_banner"] = shopBanner })
}

// WithShopDesc shop_desc获取 店铺简介
func (obj *ShopDetailMgr) WithShopDesc(shopDesc string) Option {
	return optionFunc(func(o *options) { o.query["shop_desc"] = shopDesc })
}

// WithShopRecommend shop_recommend获取 是否推荐
func (obj *ShopDetailMgr) WithShopRecommend(shopRecommend int) Option {
	return optionFunc(func(o *options) { o.query["shop_recommend"] = shopRecommend })
}

// WithShopThemeid shop_themeid获取 店铺主题id
func (obj *ShopDetailMgr) WithShopThemeid(shopThemeid int) Option {
	return optionFunc(func(o *options) { o.query["shop_themeid"] = shopThemeid })
}

// WithShopThemePath shop_theme_path获取 店铺主题模版路径
func (obj *ShopDetailMgr) WithShopThemePath(shopThemePath string) Option {
	return optionFunc(func(o *options) { o.query["shop_theme_path"] = shopThemePath })
}

// WithWapThemeid wap_themeid获取 店铺主题id
func (obj *ShopDetailMgr) WithWapThemeid(wapThemeid int) Option {
	return optionFunc(func(o *options) { o.query["wap_themeid"] = wapThemeid })
}

// WithWapThemePath wap_theme_path获取 wap店铺主题
func (obj *ShopDetailMgr) WithWapThemePath(wapThemePath string) Option {
	return optionFunc(func(o *options) { o.query["wap_theme_path"] = wapThemePath })
}

// WithShopCredit shop_credit获取 店铺信用
func (obj *ShopDetailMgr) WithShopCredit(shopCredit float64) Option {
	return optionFunc(func(o *options) { o.query["shop_credit"] = shopCredit })
}

// WithShopPraiseRate shop_praise_rate获取 店铺好评率
func (obj *ShopDetailMgr) WithShopPraiseRate(shopPraiseRate float64) Option {
	return optionFunc(func(o *options) { o.query["shop_praise_rate"] = shopPraiseRate })
}

// WithShopDescriptionCredit shop_description_credit获取 店铺描述相符度
func (obj *ShopDetailMgr) WithShopDescriptionCredit(shopDescriptionCredit float64) Option {
	return optionFunc(func(o *options) { o.query["shop_description_credit"] = shopDescriptionCredit })
}

// WithShopServiceCredit shop_service_credit获取 服务态度分数
func (obj *ShopDetailMgr) WithShopServiceCredit(shopServiceCredit float64) Option {
	return optionFunc(func(o *options) { o.query["shop_service_credit"] = shopServiceCredit })
}

// WithShopDeliveryCredit shop_delivery_credit获取 发货速度分数
func (obj *ShopDetailMgr) WithShopDeliveryCredit(shopDeliveryCredit float64) Option {
	return optionFunc(func(o *options) { o.query["shop_delivery_credit"] = shopDeliveryCredit })
}

// WithShopCollect shop_collect获取 店铺收藏数
func (obj *ShopDetailMgr) WithShopCollect(shopCollect int) Option {
	return optionFunc(func(o *options) { o.query["shop_collect"] = shopCollect })
}

// WithGoodsNum goods_num获取 店铺商品数
func (obj *ShopDetailMgr) WithGoodsNum(goodsNum int) Option {
	return optionFunc(func(o *options) { o.query["goods_num"] = goodsNum })
}

// WithShopQq shop_qq获取 店铺客服qq
func (obj *ShopDetailMgr) WithShopQq(shopQq string) Option {
	return optionFunc(func(o *options) { o.query["shop_qq"] = shopQq })
}

// WithShopCommission shop_commission获取 店铺佣金比例
func (obj *ShopDetailMgr) WithShopCommission(shopCommission float64) Option {
	return optionFunc(func(o *options) { o.query["shop_commission"] = shopCommission })
}

// WithGoodsWarningCount goods_warning_count获取 货品预警数
func (obj *ShopDetailMgr) WithGoodsWarningCount(goodsWarningCount int) Option {
	return optionFunc(func(o *options) { o.query["goods_warning_count"] = goodsWarningCount })
}

// WithSelfOperated self_operated获取 是否自营 1:是 0:否
func (obj *ShopDetailMgr) WithSelfOperated(selfOperated int) Option {
	return optionFunc(func(o *options) { o.query["self_operated"] = selfOperated })
}

// WithStep step获取 申请开店进度：1,2.3.4
func (obj *ShopDetailMgr) WithStep(step int) Option {
	return optionFunc(func(o *options) { o.query["step"] = step })
}

// WithOrdinReceiptStatus ordin_receipt_status获取 是否允许开具增值税普通发票 0：否，1：是
func (obj *ShopDetailMgr) WithOrdinReceiptStatus(ordinReceiptStatus int) Option {
	return optionFunc(func(o *options) { o.query["ordin_receipt_status"] = ordinReceiptStatus })
}

// WithElecReceiptStatus elec_receipt_status获取 是否允许开具电子普通发票 0：否，1：是
func (obj *ShopDetailMgr) WithElecReceiptStatus(elecReceiptStatus int) Option {
	return optionFunc(func(o *options) { o.query["elec_receipt_status"] = elecReceiptStatus })
}

// WithTaxReceiptStatus tax_receipt_status获取 是否允许开具增值税专用发票 0：否，1：是
func (obj *ShopDetailMgr) WithTaxReceiptStatus(taxReceiptStatus int) Option {
	return optionFunc(func(o *options) { o.query["tax_receipt_status"] = taxReceiptStatus })
}

// GetByOption 功能选项模式获取
func (obj *ShopDetailMgr) GetByOption(opts ...Option) (result models.EsShopDetail, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *ShopDetailMgr) GetByOptions(opts ...Option) (results []*models.EsShopDetail, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *ShopDetailMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsShopDetail, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where(options.query)
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

// GetFromID 通过id获取内容 店铺详细id
func (obj *ShopDetailMgr) GetFromID(id int) (result models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找 店铺详细id
func (obj *ShopDetailMgr) GetBatchFromID(ids []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺id
func (obj *ShopDetailMgr) GetFromShopID(shopID int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺id
func (obj *ShopDetailMgr) GetBatchFromShopID(shopIDs []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromShopProvinceID 通过shop_province_id获取内容 店铺所在省id
func (obj *ShopDetailMgr) GetFromShopProvinceID(shopProvinceID int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_province_id` = ?", shopProvinceID).Find(&results).Error

	return
}

// GetBatchFromShopProvinceID 批量查找 店铺所在省id
func (obj *ShopDetailMgr) GetBatchFromShopProvinceID(shopProvinceIDs []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_province_id` IN (?)", shopProvinceIDs).Find(&results).Error

	return
}

// GetFromShopCityID 通过shop_city_id获取内容 店铺所在市id
func (obj *ShopDetailMgr) GetFromShopCityID(shopCityID int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_city_id` = ?", shopCityID).Find(&results).Error

	return
}

// GetBatchFromShopCityID 批量查找 店铺所在市id
func (obj *ShopDetailMgr) GetBatchFromShopCityID(shopCityIDs []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_city_id` IN (?)", shopCityIDs).Find(&results).Error

	return
}

// GetFromShopCountyID 通过shop_county_id获取内容 店铺所在县id
func (obj *ShopDetailMgr) GetFromShopCountyID(shopCountyID int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_county_id` = ?", shopCountyID).Find(&results).Error

	return
}

// GetBatchFromShopCountyID 批量查找 店铺所在县id
func (obj *ShopDetailMgr) GetBatchFromShopCountyID(shopCountyIDs []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_county_id` IN (?)", shopCountyIDs).Find(&results).Error

	return
}

// GetFromShopTownID 通过shop_town_id获取内容 店铺所在镇id
func (obj *ShopDetailMgr) GetFromShopTownID(shopTownID int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_town_id` = ?", shopTownID).Find(&results).Error

	return
}

// GetBatchFromShopTownID 批量查找 店铺所在镇id
func (obj *ShopDetailMgr) GetBatchFromShopTownID(shopTownIDs []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_town_id` IN (?)", shopTownIDs).Find(&results).Error

	return
}

// GetFromShopProvince 通过shop_province获取内容 店铺所在省
func (obj *ShopDetailMgr) GetFromShopProvince(shopProvince string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_province` = ?", shopProvince).Find(&results).Error

	return
}

// GetBatchFromShopProvince 批量查找 店铺所在省
func (obj *ShopDetailMgr) GetBatchFromShopProvince(shopProvinces []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_province` IN (?)", shopProvinces).Find(&results).Error

	return
}

// GetFromShopCity 通过shop_city获取内容 店铺所在市
func (obj *ShopDetailMgr) GetFromShopCity(shopCity string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_city` = ?", shopCity).Find(&results).Error

	return
}

// GetBatchFromShopCity 批量查找 店铺所在市
func (obj *ShopDetailMgr) GetBatchFromShopCity(shopCitys []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_city` IN (?)", shopCitys).Find(&results).Error

	return
}

// GetFromShopCounty 通过shop_county获取内容 店铺所在县
func (obj *ShopDetailMgr) GetFromShopCounty(shopCounty string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_county` = ?", shopCounty).Find(&results).Error

	return
}

// GetBatchFromShopCounty 批量查找 店铺所在县
func (obj *ShopDetailMgr) GetBatchFromShopCounty(shopCountys []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_county` IN (?)", shopCountys).Find(&results).Error

	return
}

// GetFromShopTown 通过shop_town获取内容 店铺所在镇
func (obj *ShopDetailMgr) GetFromShopTown(shopTown string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_town` = ?", shopTown).Find(&results).Error

	return
}

// GetBatchFromShopTown 批量查找 店铺所在镇
func (obj *ShopDetailMgr) GetBatchFromShopTown(shopTowns []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_town` IN (?)", shopTowns).Find(&results).Error

	return
}

// GetFromShopAdd 通过shop_add获取内容 店铺详细地址
func (obj *ShopDetailMgr) GetFromShopAdd(shopAdd string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_add` = ?", shopAdd).Find(&results).Error

	return
}

// GetBatchFromShopAdd 批量查找 店铺详细地址
func (obj *ShopDetailMgr) GetBatchFromShopAdd(shopAdds []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_add` IN (?)", shopAdds).Find(&results).Error

	return
}

// GetFromCompanyName 通过company_name获取内容 公司名称
func (obj *ShopDetailMgr) GetFromCompanyName(companyName string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`company_name` = ?", companyName).Find(&results).Error

	return
}

// GetBatchFromCompanyName 批量查找 公司名称
func (obj *ShopDetailMgr) GetBatchFromCompanyName(companyNames []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`company_name` IN (?)", companyNames).Find(&results).Error

	return
}

// GetFromCompanyAddress 通过company_address获取内容 公司地址
func (obj *ShopDetailMgr) GetFromCompanyAddress(companyAddress string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`company_address` = ?", companyAddress).Find(&results).Error

	return
}

// GetBatchFromCompanyAddress 批量查找 公司地址
func (obj *ShopDetailMgr) GetBatchFromCompanyAddress(companyAddresss []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`company_address` IN (?)", companyAddresss).Find(&results).Error

	return
}

// GetFromCompanyPhone 通过company_phone获取内容 公司电话
func (obj *ShopDetailMgr) GetFromCompanyPhone(companyPhone string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`company_phone` = ?", companyPhone).Find(&results).Error

	return
}

// GetBatchFromCompanyPhone 批量查找 公司电话
func (obj *ShopDetailMgr) GetBatchFromCompanyPhone(companyPhones []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`company_phone` IN (?)", companyPhones).Find(&results).Error

	return
}

// GetFromCompanyEmail 通过company_email获取内容 电子邮箱
func (obj *ShopDetailMgr) GetFromCompanyEmail(companyEmail string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`company_email` = ?", companyEmail).Find(&results).Error

	return
}

// GetBatchFromCompanyEmail 批量查找 电子邮箱
func (obj *ShopDetailMgr) GetBatchFromCompanyEmail(companyEmails []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`company_email` IN (?)", companyEmails).Find(&results).Error

	return
}

// GetFromEmployeeNum 通过employee_num获取内容 员工总数
func (obj *ShopDetailMgr) GetFromEmployeeNum(employeeNum int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`employee_num` = ?", employeeNum).Find(&results).Error

	return
}

// GetBatchFromEmployeeNum 批量查找 员工总数
func (obj *ShopDetailMgr) GetBatchFromEmployeeNum(employeeNums []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`employee_num` IN (?)", employeeNums).Find(&results).Error

	return
}

// GetFromRegMoney 通过reg_money获取内容 注册资金
func (obj *ShopDetailMgr) GetFromRegMoney(regMoney float64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`reg_money` = ?", regMoney).Find(&results).Error

	return
}

// GetBatchFromRegMoney 批量查找 注册资金
func (obj *ShopDetailMgr) GetBatchFromRegMoney(regMoneys []float64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`reg_money` IN (?)", regMoneys).Find(&results).Error

	return
}

// GetFromLinkName 通过link_name获取内容 联系人姓名
func (obj *ShopDetailMgr) GetFromLinkName(linkName string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`link_name` = ?", linkName).Find(&results).Error

	return
}

// GetBatchFromLinkName 批量查找 联系人姓名
func (obj *ShopDetailMgr) GetBatchFromLinkName(linkNames []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`link_name` IN (?)", linkNames).Find(&results).Error

	return
}

// GetFromLinkPhone 通过link_phone获取内容 联系人电话
func (obj *ShopDetailMgr) GetFromLinkPhone(linkPhone string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`link_phone` = ?", linkPhone).Find(&results).Error

	return
}

// GetBatchFromLinkPhone 批量查找 联系人电话
func (obj *ShopDetailMgr) GetBatchFromLinkPhone(linkPhones []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`link_phone` IN (?)", linkPhones).Find(&results).Error

	return
}

// GetFromLegalName 通过legal_name获取内容 法人姓名
func (obj *ShopDetailMgr) GetFromLegalName(legalName string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`legal_name` = ?", legalName).Find(&results).Error

	return
}

// GetBatchFromLegalName 批量查找 法人姓名
func (obj *ShopDetailMgr) GetBatchFromLegalName(legalNames []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`legal_name` IN (?)", legalNames).Find(&results).Error

	return
}

// GetFromLegalID 通过legal_id获取内容 法人身份证
func (obj *ShopDetailMgr) GetFromLegalID(legalID string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`legal_id` = ?", legalID).Find(&results).Error

	return
}

// GetBatchFromLegalID 批量查找 法人身份证
func (obj *ShopDetailMgr) GetBatchFromLegalID(legalIDs []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`legal_id` IN (?)", legalIDs).Find(&results).Error

	return
}

// GetFromLegalImg 通过legal_img获取内容 法人身份证照片
func (obj *ShopDetailMgr) GetFromLegalImg(legalImg string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`legal_img` = ?", legalImg).Find(&results).Error

	return
}

// GetBatchFromLegalImg 批量查找 法人身份证照片
func (obj *ShopDetailMgr) GetBatchFromLegalImg(legalImgs []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`legal_img` IN (?)", legalImgs).Find(&results).Error

	return
}

// GetFromLicenseNum 通过license_num获取内容 营业执照号
func (obj *ShopDetailMgr) GetFromLicenseNum(licenseNum string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_num` = ?", licenseNum).Find(&results).Error

	return
}

// GetBatchFromLicenseNum 批量查找 营业执照号
func (obj *ShopDetailMgr) GetBatchFromLicenseNum(licenseNums []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_num` IN (?)", licenseNums).Find(&results).Error

	return
}

// GetFromLicenseProvinceID 通过license_province_id获取内容 营业执照所在省id
func (obj *ShopDetailMgr) GetFromLicenseProvinceID(licenseProvinceID int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_province_id` = ?", licenseProvinceID).Find(&results).Error

	return
}

// GetBatchFromLicenseProvinceID 批量查找 营业执照所在省id
func (obj *ShopDetailMgr) GetBatchFromLicenseProvinceID(licenseProvinceIDs []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_province_id` IN (?)", licenseProvinceIDs).Find(&results).Error

	return
}

// GetFromLicenseCityID 通过license_city_id获取内容 营业执照所在市id
func (obj *ShopDetailMgr) GetFromLicenseCityID(licenseCityID int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_city_id` = ?", licenseCityID).Find(&results).Error

	return
}

// GetBatchFromLicenseCityID 批量查找 营业执照所在市id
func (obj *ShopDetailMgr) GetBatchFromLicenseCityID(licenseCityIDs []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_city_id` IN (?)", licenseCityIDs).Find(&results).Error

	return
}

// GetFromLicenseCountyID 通过license_county_id获取内容 营业执照所在县id
func (obj *ShopDetailMgr) GetFromLicenseCountyID(licenseCountyID int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_county_id` = ?", licenseCountyID).Find(&results).Error

	return
}

// GetBatchFromLicenseCountyID 批量查找 营业执照所在县id
func (obj *ShopDetailMgr) GetBatchFromLicenseCountyID(licenseCountyIDs []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_county_id` IN (?)", licenseCountyIDs).Find(&results).Error

	return
}

// GetFromLicenseTownID 通过license_town_id获取内容 营业执照所在镇id
func (obj *ShopDetailMgr) GetFromLicenseTownID(licenseTownID int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_town_id` = ?", licenseTownID).Find(&results).Error

	return
}

// GetBatchFromLicenseTownID 批量查找 营业执照所在镇id
func (obj *ShopDetailMgr) GetBatchFromLicenseTownID(licenseTownIDs []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_town_id` IN (?)", licenseTownIDs).Find(&results).Error

	return
}

// GetFromLicenseProvince 通过license_province获取内容 营业执照所在省
func (obj *ShopDetailMgr) GetFromLicenseProvince(licenseProvince string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_province` = ?", licenseProvince).Find(&results).Error

	return
}

// GetBatchFromLicenseProvince 批量查找 营业执照所在省
func (obj *ShopDetailMgr) GetBatchFromLicenseProvince(licenseProvinces []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_province` IN (?)", licenseProvinces).Find(&results).Error

	return
}

// GetFromLicenseCity 通过license_city获取内容 营业执照所在市
func (obj *ShopDetailMgr) GetFromLicenseCity(licenseCity string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_city` = ?", licenseCity).Find(&results).Error

	return
}

// GetBatchFromLicenseCity 批量查找 营业执照所在市
func (obj *ShopDetailMgr) GetBatchFromLicenseCity(licenseCitys []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_city` IN (?)", licenseCitys).Find(&results).Error

	return
}

// GetFromLicenseCounty 通过license_county获取内容 营业执照所在县
func (obj *ShopDetailMgr) GetFromLicenseCounty(licenseCounty string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_county` = ?", licenseCounty).Find(&results).Error

	return
}

// GetBatchFromLicenseCounty 批量查找 营业执照所在县
func (obj *ShopDetailMgr) GetBatchFromLicenseCounty(licenseCountys []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_county` IN (?)", licenseCountys).Find(&results).Error

	return
}

// GetFromLicenseTown 通过license_town获取内容 营业执照所在镇
func (obj *ShopDetailMgr) GetFromLicenseTown(licenseTown string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_town` = ?", licenseTown).Find(&results).Error

	return
}

// GetBatchFromLicenseTown 批量查找 营业执照所在镇
func (obj *ShopDetailMgr) GetBatchFromLicenseTown(licenseTowns []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_town` IN (?)", licenseTowns).Find(&results).Error

	return
}

// GetFromLicenseAdd 通过license_add获取内容 营业执照详细地址
func (obj *ShopDetailMgr) GetFromLicenseAdd(licenseAdd string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_add` = ?", licenseAdd).Find(&results).Error

	return
}

// GetBatchFromLicenseAdd 批量查找 营业执照详细地址
func (obj *ShopDetailMgr) GetBatchFromLicenseAdd(licenseAdds []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`license_add` IN (?)", licenseAdds).Find(&results).Error

	return
}

// GetFromEstablishDate 通过establish_date获取内容 成立日期
func (obj *ShopDetailMgr) GetFromEstablishDate(establishDate int64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`establish_date` = ?", establishDate).Find(&results).Error

	return
}

// GetBatchFromEstablishDate 批量查找 成立日期
func (obj *ShopDetailMgr) GetBatchFromEstablishDate(establishDates []int64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`establish_date` IN (?)", establishDates).Find(&results).Error

	return
}

// GetFromLicenceStart 通过licence_start获取内容 营业执照有效期开始
func (obj *ShopDetailMgr) GetFromLicenceStart(licenceStart int64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`licence_start` = ?", licenceStart).Find(&results).Error

	return
}

// GetBatchFromLicenceStart 批量查找 营业执照有效期开始
func (obj *ShopDetailMgr) GetBatchFromLicenceStart(licenceStarts []int64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`licence_start` IN (?)", licenceStarts).Find(&results).Error

	return
}

// GetFromLicenceEnd 通过licence_end获取内容 营业执照有效期结束
func (obj *ShopDetailMgr) GetFromLicenceEnd(licenceEnd int64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`licence_end` = ?", licenceEnd).Find(&results).Error

	return
}

// GetBatchFromLicenceEnd 批量查找 营业执照有效期结束
func (obj *ShopDetailMgr) GetBatchFromLicenceEnd(licenceEnds []int64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`licence_end` IN (?)", licenceEnds).Find(&results).Error

	return
}

// GetFromScope 通过scope获取内容 法定经营范围
func (obj *ShopDetailMgr) GetFromScope(scope string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`scope` = ?", scope).Find(&results).Error

	return
}

// GetBatchFromScope 批量查找 法定经营范围
func (obj *ShopDetailMgr) GetBatchFromScope(scopes []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`scope` IN (?)", scopes).Find(&results).Error

	return
}

// GetFromLicenceImg 通过licence_img获取内容 营业执照电子版
func (obj *ShopDetailMgr) GetFromLicenceImg(licenceImg string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`licence_img` = ?", licenceImg).Find(&results).Error

	return
}

// GetBatchFromLicenceImg 批量查找 营业执照电子版
func (obj *ShopDetailMgr) GetBatchFromLicenceImg(licenceImgs []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`licence_img` IN (?)", licenceImgs).Find(&results).Error

	return
}

// GetFromOrganizationCode 通过organization_code获取内容 组织机构代码
func (obj *ShopDetailMgr) GetFromOrganizationCode(organizationCode string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`organization_code` = ?", organizationCode).Find(&results).Error

	return
}

// GetBatchFromOrganizationCode 批量查找 组织机构代码
func (obj *ShopDetailMgr) GetBatchFromOrganizationCode(organizationCodes []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`organization_code` IN (?)", organizationCodes).Find(&results).Error

	return
}

// GetFromCodeImg 通过code_img获取内容 组织机构电子版
func (obj *ShopDetailMgr) GetFromCodeImg(codeImg string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`code_img` = ?", codeImg).Find(&results).Error

	return
}

// GetBatchFromCodeImg 批量查找 组织机构电子版
func (obj *ShopDetailMgr) GetBatchFromCodeImg(codeImgs []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`code_img` IN (?)", codeImgs).Find(&results).Error

	return
}

// GetFromTaxesImg 通过taxes_img获取内容 一般纳税人证明电子版
func (obj *ShopDetailMgr) GetFromTaxesImg(taxesImg string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`taxes_img` = ?", taxesImg).Find(&results).Error

	return
}

// GetBatchFromTaxesImg 批量查找 一般纳税人证明电子版
func (obj *ShopDetailMgr) GetBatchFromTaxesImg(taxesImgs []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`taxes_img` IN (?)", taxesImgs).Find(&results).Error

	return
}

// GetFromBankAccountName 通过bank_account_name获取内容 银行开户名
func (obj *ShopDetailMgr) GetFromBankAccountName(bankAccountName string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_account_name` = ?", bankAccountName).Find(&results).Error

	return
}

// GetBatchFromBankAccountName 批量查找 银行开户名
func (obj *ShopDetailMgr) GetBatchFromBankAccountName(bankAccountNames []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_account_name` IN (?)", bankAccountNames).Find(&results).Error

	return
}

// GetFromBankNumber 通过bank_number获取内容 银行开户账号
func (obj *ShopDetailMgr) GetFromBankNumber(bankNumber string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_number` = ?", bankNumber).Find(&results).Error

	return
}

// GetBatchFromBankNumber 批量查找 银行开户账号
func (obj *ShopDetailMgr) GetBatchFromBankNumber(bankNumbers []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_number` IN (?)", bankNumbers).Find(&results).Error

	return
}

// GetFromBankName 通过bank_name获取内容 开户银行支行名称
func (obj *ShopDetailMgr) GetFromBankName(bankName string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_name` = ?", bankName).Find(&results).Error

	return
}

// GetBatchFromBankName 批量查找 开户银行支行名称
func (obj *ShopDetailMgr) GetBatchFromBankName(bankNames []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_name` IN (?)", bankNames).Find(&results).Error

	return
}

// GetFromBankProvinceID 通过bank_province_id获取内容 开户银行所在省id
func (obj *ShopDetailMgr) GetFromBankProvinceID(bankProvinceID int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_province_id` = ?", bankProvinceID).Find(&results).Error

	return
}

// GetBatchFromBankProvinceID 批量查找 开户银行所在省id
func (obj *ShopDetailMgr) GetBatchFromBankProvinceID(bankProvinceIDs []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_province_id` IN (?)", bankProvinceIDs).Find(&results).Error

	return
}

// GetFromBankCityID 通过bank_city_id获取内容 开户银行所在市id
func (obj *ShopDetailMgr) GetFromBankCityID(bankCityID int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_city_id` = ?", bankCityID).Find(&results).Error

	return
}

// GetBatchFromBankCityID 批量查找 开户银行所在市id
func (obj *ShopDetailMgr) GetBatchFromBankCityID(bankCityIDs []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_city_id` IN (?)", bankCityIDs).Find(&results).Error

	return
}

// GetFromBankCountyID 通过bank_county_id获取内容 开户银行所在县id
func (obj *ShopDetailMgr) GetFromBankCountyID(bankCountyID int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_county_id` = ?", bankCountyID).Find(&results).Error

	return
}

// GetBatchFromBankCountyID 批量查找 开户银行所在县id
func (obj *ShopDetailMgr) GetBatchFromBankCountyID(bankCountyIDs []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_county_id` IN (?)", bankCountyIDs).Find(&results).Error

	return
}

// GetFromBankTownID 通过bank_town_id获取内容 开户银行所在镇id
func (obj *ShopDetailMgr) GetFromBankTownID(bankTownID int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_town_id` = ?", bankTownID).Find(&results).Error

	return
}

// GetBatchFromBankTownID 批量查找 开户银行所在镇id
func (obj *ShopDetailMgr) GetBatchFromBankTownID(bankTownIDs []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_town_id` IN (?)", bankTownIDs).Find(&results).Error

	return
}

// GetFromBankProvince 通过bank_province获取内容 开户银行所在省
func (obj *ShopDetailMgr) GetFromBankProvince(bankProvince string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_province` = ?", bankProvince).Find(&results).Error

	return
}

// GetBatchFromBankProvince 批量查找 开户银行所在省
func (obj *ShopDetailMgr) GetBatchFromBankProvince(bankProvinces []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_province` IN (?)", bankProvinces).Find(&results).Error

	return
}

// GetFromBankCity 通过bank_city获取内容 开户银行所在市
func (obj *ShopDetailMgr) GetFromBankCity(bankCity string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_city` = ?", bankCity).Find(&results).Error

	return
}

// GetBatchFromBankCity 批量查找 开户银行所在市
func (obj *ShopDetailMgr) GetBatchFromBankCity(bankCitys []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_city` IN (?)", bankCitys).Find(&results).Error

	return
}

// GetFromBankCounty 通过bank_county获取内容 开户银行所在县
func (obj *ShopDetailMgr) GetFromBankCounty(bankCounty string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_county` = ?", bankCounty).Find(&results).Error

	return
}

// GetBatchFromBankCounty 批量查找 开户银行所在县
func (obj *ShopDetailMgr) GetBatchFromBankCounty(bankCountys []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_county` IN (?)", bankCountys).Find(&results).Error

	return
}

// GetFromBankTown 通过bank_town获取内容 开户银行所在镇
func (obj *ShopDetailMgr) GetFromBankTown(bankTown string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_town` = ?", bankTown).Find(&results).Error

	return
}

// GetBatchFromBankTown 批量查找 开户银行所在镇
func (obj *ShopDetailMgr) GetBatchFromBankTown(bankTowns []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_town` IN (?)", bankTowns).Find(&results).Error

	return
}

// GetFromBankImg 通过bank_img获取内容 开户银行许可证电子版
func (obj *ShopDetailMgr) GetFromBankImg(bankImg string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_img` = ?", bankImg).Find(&results).Error

	return
}

// GetBatchFromBankImg 批量查找 开户银行许可证电子版
func (obj *ShopDetailMgr) GetBatchFromBankImg(bankImgs []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`bank_img` IN (?)", bankImgs).Find(&results).Error

	return
}

// GetFromTaxesCertificateNum 通过taxes_certificate_num获取内容 税务登记证号
func (obj *ShopDetailMgr) GetFromTaxesCertificateNum(taxesCertificateNum string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`taxes_certificate_num` = ?", taxesCertificateNum).Find(&results).Error

	return
}

// GetBatchFromTaxesCertificateNum 批量查找 税务登记证号
func (obj *ShopDetailMgr) GetBatchFromTaxesCertificateNum(taxesCertificateNums []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`taxes_certificate_num` IN (?)", taxesCertificateNums).Find(&results).Error

	return
}

// GetFromTaxesDistinguishNum 通过taxes_distinguish_num获取内容 纳税人识别号
func (obj *ShopDetailMgr) GetFromTaxesDistinguishNum(taxesDistinguishNum string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`taxes_distinguish_num` = ?", taxesDistinguishNum).Find(&results).Error

	return
}

// GetBatchFromTaxesDistinguishNum 批量查找 纳税人识别号
func (obj *ShopDetailMgr) GetBatchFromTaxesDistinguishNum(taxesDistinguishNums []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`taxes_distinguish_num` IN (?)", taxesDistinguishNums).Find(&results).Error

	return
}

// GetFromTaxesCertificateImg 通过taxes_certificate_img获取内容 税务登记证书
func (obj *ShopDetailMgr) GetFromTaxesCertificateImg(taxesCertificateImg string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`taxes_certificate_img` = ?", taxesCertificateImg).Find(&results).Error

	return
}

// GetBatchFromTaxesCertificateImg 批量查找 税务登记证书
func (obj *ShopDetailMgr) GetBatchFromTaxesCertificateImg(taxesCertificateImgs []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`taxes_certificate_img` IN (?)", taxesCertificateImgs).Find(&results).Error

	return
}

// GetFromGoodsManagementCategory 通过goods_management_category获取内容 店铺经营类目
func (obj *ShopDetailMgr) GetFromGoodsManagementCategory(goodsManagementCategory string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`goods_management_category` = ?", goodsManagementCategory).Find(&results).Error

	return
}

// GetBatchFromGoodsManagementCategory 批量查找 店铺经营类目
func (obj *ShopDetailMgr) GetBatchFromGoodsManagementCategory(goodsManagementCategorys []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`goods_management_category` IN (?)", goodsManagementCategorys).Find(&results).Error

	return
}

// GetFromShopLevel 通过shop_level获取内容 店铺等级
func (obj *ShopDetailMgr) GetFromShopLevel(shopLevel int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_level` = ?", shopLevel).Find(&results).Error

	return
}

// GetBatchFromShopLevel 批量查找 店铺等级
func (obj *ShopDetailMgr) GetBatchFromShopLevel(shopLevels []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_level` IN (?)", shopLevels).Find(&results).Error

	return
}

// GetFromShopLevelApply 通过shop_level_apply获取内容 店铺等级申请
func (obj *ShopDetailMgr) GetFromShopLevelApply(shopLevelApply int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_level_apply` = ?", shopLevelApply).Find(&results).Error

	return
}

// GetBatchFromShopLevelApply 批量查找 店铺等级申请
func (obj *ShopDetailMgr) GetBatchFromShopLevelApply(shopLevelApplys []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_level_apply` IN (?)", shopLevelApplys).Find(&results).Error

	return
}

// GetFromStoreSpaceCapacity 通过store_space_capacity获取内容 店铺相册已用存储量
func (obj *ShopDetailMgr) GetFromStoreSpaceCapacity(storeSpaceCapacity float64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`store_space_capacity` = ?", storeSpaceCapacity).Find(&results).Error

	return
}

// GetBatchFromStoreSpaceCapacity 批量查找 店铺相册已用存储量
func (obj *ShopDetailMgr) GetBatchFromStoreSpaceCapacity(storeSpaceCapacitys []float64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`store_space_capacity` IN (?)", storeSpaceCapacitys).Find(&results).Error

	return
}

// GetFromShopLogo 通过shop_logo获取内容 店铺logo
func (obj *ShopDetailMgr) GetFromShopLogo(shopLogo string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_logo` = ?", shopLogo).Find(&results).Error

	return
}

// GetBatchFromShopLogo 批量查找 店铺logo
func (obj *ShopDetailMgr) GetBatchFromShopLogo(shopLogos []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_logo` IN (?)", shopLogos).Find(&results).Error

	return
}

// GetFromShopBanner 通过shop_banner获取内容 店铺横幅
func (obj *ShopDetailMgr) GetFromShopBanner(shopBanner string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_banner` = ?", shopBanner).Find(&results).Error

	return
}

// GetBatchFromShopBanner 批量查找 店铺横幅
func (obj *ShopDetailMgr) GetBatchFromShopBanner(shopBanners []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_banner` IN (?)", shopBanners).Find(&results).Error

	return
}

// GetFromShopDesc 通过shop_desc获取内容 店铺简介
func (obj *ShopDetailMgr) GetFromShopDesc(shopDesc string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_desc` = ?", shopDesc).Find(&results).Error

	return
}

// GetBatchFromShopDesc 批量查找 店铺简介
func (obj *ShopDetailMgr) GetBatchFromShopDesc(shopDescs []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_desc` IN (?)", shopDescs).Find(&results).Error

	return
}

// GetFromShopRecommend 通过shop_recommend获取内容 是否推荐
func (obj *ShopDetailMgr) GetFromShopRecommend(shopRecommend int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_recommend` = ?", shopRecommend).Find(&results).Error

	return
}

// GetBatchFromShopRecommend 批量查找 是否推荐
func (obj *ShopDetailMgr) GetBatchFromShopRecommend(shopRecommends []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_recommend` IN (?)", shopRecommends).Find(&results).Error

	return
}

// GetFromShopThemeid 通过shop_themeid获取内容 店铺主题id
func (obj *ShopDetailMgr) GetFromShopThemeid(shopThemeid int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_themeid` = ?", shopThemeid).Find(&results).Error

	return
}

// GetBatchFromShopThemeid 批量查找 店铺主题id
func (obj *ShopDetailMgr) GetBatchFromShopThemeid(shopThemeids []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_themeid` IN (?)", shopThemeids).Find(&results).Error

	return
}

// GetFromShopThemePath 通过shop_theme_path获取内容 店铺主题模版路径
func (obj *ShopDetailMgr) GetFromShopThemePath(shopThemePath string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_theme_path` = ?", shopThemePath).Find(&results).Error

	return
}

// GetBatchFromShopThemePath 批量查找 店铺主题模版路径
func (obj *ShopDetailMgr) GetBatchFromShopThemePath(shopThemePaths []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_theme_path` IN (?)", shopThemePaths).Find(&results).Error

	return
}

// GetFromWapThemeid 通过wap_themeid获取内容 店铺主题id
func (obj *ShopDetailMgr) GetFromWapThemeid(wapThemeid int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`wap_themeid` = ?", wapThemeid).Find(&results).Error

	return
}

// GetBatchFromWapThemeid 批量查找 店铺主题id
func (obj *ShopDetailMgr) GetBatchFromWapThemeid(wapThemeids []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`wap_themeid` IN (?)", wapThemeids).Find(&results).Error

	return
}

// GetFromWapThemePath 通过wap_theme_path获取内容 wap店铺主题
func (obj *ShopDetailMgr) GetFromWapThemePath(wapThemePath string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`wap_theme_path` = ?", wapThemePath).Find(&results).Error

	return
}

// GetBatchFromWapThemePath 批量查找 wap店铺主题
func (obj *ShopDetailMgr) GetBatchFromWapThemePath(wapThemePaths []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`wap_theme_path` IN (?)", wapThemePaths).Find(&results).Error

	return
}

// GetFromShopCredit 通过shop_credit获取内容 店铺信用
func (obj *ShopDetailMgr) GetFromShopCredit(shopCredit float64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_credit` = ?", shopCredit).Find(&results).Error

	return
}

// GetBatchFromShopCredit 批量查找 店铺信用
func (obj *ShopDetailMgr) GetBatchFromShopCredit(shopCredits []float64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_credit` IN (?)", shopCredits).Find(&results).Error

	return
}

// GetFromShopPraiseRate 通过shop_praise_rate获取内容 店铺好评率
func (obj *ShopDetailMgr) GetFromShopPraiseRate(shopPraiseRate float64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_praise_rate` = ?", shopPraiseRate).Find(&results).Error

	return
}

// GetBatchFromShopPraiseRate 批量查找 店铺好评率
func (obj *ShopDetailMgr) GetBatchFromShopPraiseRate(shopPraiseRates []float64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_praise_rate` IN (?)", shopPraiseRates).Find(&results).Error

	return
}

// GetFromShopDescriptionCredit 通过shop_description_credit获取内容 店铺描述相符度
func (obj *ShopDetailMgr) GetFromShopDescriptionCredit(shopDescriptionCredit float64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_description_credit` = ?", shopDescriptionCredit).Find(&results).Error

	return
}

// GetBatchFromShopDescriptionCredit 批量查找 店铺描述相符度
func (obj *ShopDetailMgr) GetBatchFromShopDescriptionCredit(shopDescriptionCredits []float64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_description_credit` IN (?)", shopDescriptionCredits).Find(&results).Error

	return
}

// GetFromShopServiceCredit 通过shop_service_credit获取内容 服务态度分数
func (obj *ShopDetailMgr) GetFromShopServiceCredit(shopServiceCredit float64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_service_credit` = ?", shopServiceCredit).Find(&results).Error

	return
}

// GetBatchFromShopServiceCredit 批量查找 服务态度分数
func (obj *ShopDetailMgr) GetBatchFromShopServiceCredit(shopServiceCredits []float64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_service_credit` IN (?)", shopServiceCredits).Find(&results).Error

	return
}

// GetFromShopDeliveryCredit 通过shop_delivery_credit获取内容 发货速度分数
func (obj *ShopDetailMgr) GetFromShopDeliveryCredit(shopDeliveryCredit float64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_delivery_credit` = ?", shopDeliveryCredit).Find(&results).Error

	return
}

// GetBatchFromShopDeliveryCredit 批量查找 发货速度分数
func (obj *ShopDetailMgr) GetBatchFromShopDeliveryCredit(shopDeliveryCredits []float64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_delivery_credit` IN (?)", shopDeliveryCredits).Find(&results).Error

	return
}

// GetFromShopCollect 通过shop_collect获取内容 店铺收藏数
func (obj *ShopDetailMgr) GetFromShopCollect(shopCollect int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_collect` = ?", shopCollect).Find(&results).Error

	return
}

// GetBatchFromShopCollect 批量查找 店铺收藏数
func (obj *ShopDetailMgr) GetBatchFromShopCollect(shopCollects []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_collect` IN (?)", shopCollects).Find(&results).Error

	return
}

// GetFromGoodsNum 通过goods_num获取内容 店铺商品数
func (obj *ShopDetailMgr) GetFromGoodsNum(goodsNum int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`goods_num` = ?", goodsNum).Find(&results).Error

	return
}

// GetBatchFromGoodsNum 批量查找 店铺商品数
func (obj *ShopDetailMgr) GetBatchFromGoodsNum(goodsNums []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`goods_num` IN (?)", goodsNums).Find(&results).Error

	return
}

// GetFromShopQq 通过shop_qq获取内容 店铺客服qq
func (obj *ShopDetailMgr) GetFromShopQq(shopQq string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_qq` = ?", shopQq).Find(&results).Error

	return
}

// GetBatchFromShopQq 批量查找 店铺客服qq
func (obj *ShopDetailMgr) GetBatchFromShopQq(shopQqs []string) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_qq` IN (?)", shopQqs).Find(&results).Error

	return
}

// GetFromShopCommission 通过shop_commission获取内容 店铺佣金比例
func (obj *ShopDetailMgr) GetFromShopCommission(shopCommission float64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_commission` = ?", shopCommission).Find(&results).Error

	return
}

// GetBatchFromShopCommission 批量查找 店铺佣金比例
func (obj *ShopDetailMgr) GetBatchFromShopCommission(shopCommissions []float64) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`shop_commission` IN (?)", shopCommissions).Find(&results).Error

	return
}

// GetFromGoodsWarningCount 通过goods_warning_count获取内容 货品预警数
func (obj *ShopDetailMgr) GetFromGoodsWarningCount(goodsWarningCount int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`goods_warning_count` = ?", goodsWarningCount).Find(&results).Error

	return
}

// GetBatchFromGoodsWarningCount 批量查找 货品预警数
func (obj *ShopDetailMgr) GetBatchFromGoodsWarningCount(goodsWarningCounts []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`goods_warning_count` IN (?)", goodsWarningCounts).Find(&results).Error

	return
}

// GetFromSelfOperated 通过self_operated获取内容 是否自营 1:是 0:否
func (obj *ShopDetailMgr) GetFromSelfOperated(selfOperated int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`self_operated` = ?", selfOperated).Find(&results).Error

	return
}

// GetBatchFromSelfOperated 批量查找 是否自营 1:是 0:否
func (obj *ShopDetailMgr) GetBatchFromSelfOperated(selfOperateds []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`self_operated` IN (?)", selfOperateds).Find(&results).Error

	return
}

// GetFromStep 通过step获取内容 申请开店进度：1,2.3.4
func (obj *ShopDetailMgr) GetFromStep(step int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`step` = ?", step).Find(&results).Error

	return
}

// GetBatchFromStep 批量查找 申请开店进度：1,2.3.4
func (obj *ShopDetailMgr) GetBatchFromStep(steps []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`step` IN (?)", steps).Find(&results).Error

	return
}

// GetFromOrdinReceiptStatus 通过ordin_receipt_status获取内容 是否允许开具增值税普通发票 0：否，1：是
func (obj *ShopDetailMgr) GetFromOrdinReceiptStatus(ordinReceiptStatus int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`ordin_receipt_status` = ?", ordinReceiptStatus).Find(&results).Error

	return
}

// GetBatchFromOrdinReceiptStatus 批量查找 是否允许开具增值税普通发票 0：否，1：是
func (obj *ShopDetailMgr) GetBatchFromOrdinReceiptStatus(ordinReceiptStatuss []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`ordin_receipt_status` IN (?)", ordinReceiptStatuss).Find(&results).Error

	return
}

// GetFromElecReceiptStatus 通过elec_receipt_status获取内容 是否允许开具电子普通发票 0：否，1：是
func (obj *ShopDetailMgr) GetFromElecReceiptStatus(elecReceiptStatus int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`elec_receipt_status` = ?", elecReceiptStatus).Find(&results).Error

	return
}

// GetBatchFromElecReceiptStatus 批量查找 是否允许开具电子普通发票 0：否，1：是
func (obj *ShopDetailMgr) GetBatchFromElecReceiptStatus(elecReceiptStatuss []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`elec_receipt_status` IN (?)", elecReceiptStatuss).Find(&results).Error

	return
}

// GetFromTaxReceiptStatus 通过tax_receipt_status获取内容 是否允许开具增值税专用发票 0：否，1：是
func (obj *ShopDetailMgr) GetFromTaxReceiptStatus(taxReceiptStatus int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`tax_receipt_status` = ?", taxReceiptStatus).Find(&results).Error

	return
}

// GetBatchFromTaxReceiptStatus 批量查找 是否允许开具增值税专用发票 0：否，1：是
func (obj *ShopDetailMgr) GetBatchFromTaxReceiptStatus(taxReceiptStatuss []int) (results []*models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`tax_receipt_status` IN (?)", taxReceiptStatuss).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *ShopDetailMgr) FetchByPrimaryKey(id int) (result models.EsShopDetail, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsShopDetail{}).Where("`id` = ?", id).First(&result).Error

	return
}
