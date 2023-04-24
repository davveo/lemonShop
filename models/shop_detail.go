package models

// EsShopDetail 店铺详细(es_shop_detail)
type EsShopDetail struct {
	ID                      int     `gorm:"primaryKey;column:id" json:"-"`                                     // 店铺详细id
	ShopID                  int     `gorm:"column:shop_id" json:"shop_id"`                                     // 店铺id
	ShopProvinceID          int     `gorm:"column:shop_province_id" json:"shop_province_id"`                   // 店铺所在省id
	ShopCityID              int     `gorm:"column:shop_city_id" json:"shop_city_id"`                           // 店铺所在市id
	ShopCountyID            int     `gorm:"column:shop_county_id" json:"shop_county_id"`                       // 店铺所在县id
	ShopTownID              int     `gorm:"column:shop_town_id" json:"shop_town_id"`                           // 店铺所在镇id
	ShopProvince            string  `gorm:"column:shop_province" json:"shop_province"`                         // 店铺所在省
	ShopCity                string  `gorm:"column:shop_city" json:"shop_city"`                                 // 店铺所在市
	ShopCounty              string  `gorm:"column:shop_county" json:"shop_county"`                             // 店铺所在县
	ShopTown                string  `gorm:"column:shop_town" json:"shop_town"`                                 // 店铺所在镇
	ShopAdd                 string  `gorm:"column:shop_add" json:"shop_add"`                                   // 店铺详细地址
	CompanyName             string  `gorm:"column:company_name" json:"company_name"`                           // 公司名称
	CompanyAddress          string  `gorm:"column:company_address" json:"company_address"`                     // 公司地址
	CompanyPhone            string  `gorm:"column:company_phone" json:"company_phone"`                         // 公司电话
	CompanyEmail            string  `gorm:"column:company_email" json:"company_email"`                         // 电子邮箱
	EmployeeNum             int     `gorm:"column:employee_num" json:"employee_num"`                           // 员工总数
	RegMoney                float64 `gorm:"column:reg_money" json:"reg_money"`                                 // 注册资金
	LinkName                string  `gorm:"column:link_name" json:"link_name"`                                 // 联系人姓名
	LinkPhone               string  `gorm:"column:link_phone" json:"link_phone"`                               // 联系人电话
	LegalName               string  `gorm:"column:legal_name" json:"legal_name"`                               // 法人姓名
	LegalID                 string  `gorm:"column:legal_id" json:"legal_id"`                                   // 法人身份证
	LegalImg                string  `gorm:"column:legal_img" json:"legal_img"`                                 // 法人身份证照片
	LicenseNum              string  `gorm:"column:license_num" json:"license_num"`                             // 营业执照号
	LicenseProvinceID       int     `gorm:"column:license_province_id" json:"license_province_id"`             // 营业执照所在省id
	LicenseCityID           int     `gorm:"column:license_city_id" json:"license_city_id"`                     // 营业执照所在市id
	LicenseCountyID         int     `gorm:"column:license_county_id" json:"license_county_id"`                 // 营业执照所在县id
	LicenseTownID           int     `gorm:"column:license_town_id" json:"license_town_id"`                     // 营业执照所在镇id
	LicenseProvince         string  `gorm:"column:license_province" json:"license_province"`                   // 营业执照所在省
	LicenseCity             string  `gorm:"column:license_city" json:"license_city"`                           // 营业执照所在市
	LicenseCounty           string  `gorm:"column:license_county" json:"license_county"`                       // 营业执照所在县
	LicenseTown             string  `gorm:"column:license_town" json:"license_town"`                           // 营业执照所在镇
	LicenseAdd              string  `gorm:"column:license_add" json:"license_add"`                             // 营业执照详细地址
	EstablishDate           int64   `gorm:"column:establish_date" json:"establish_date"`                       // 成立日期
	LicenceStart            int64   `gorm:"column:licence_start" json:"licence_start"`                         // 营业执照有效期开始
	LicenceEnd              int64   `gorm:"column:licence_end" json:"licence_end"`                             // 营业执照有效期结束
	Scope                   string  `gorm:"column:scope" json:"scope"`                                         // 法定经营范围
	LicenceImg              string  `gorm:"column:licence_img" json:"licence_img"`                             // 营业执照电子版
	OrganizationCode        string  `gorm:"column:organization_code" json:"organization_code"`                 // 组织机构代码
	CodeImg                 string  `gorm:"column:code_img" json:"code_img"`                                   // 组织机构电子版
	TaxesImg                string  `gorm:"column:taxes_img" json:"taxes_img"`                                 // 一般纳税人证明电子版
	BankAccountName         string  `gorm:"column:bank_account_name" json:"bank_account_name"`                 // 银行开户名
	BankNumber              string  `gorm:"column:bank_number" json:"bank_number"`                             // 银行开户账号
	BankName                string  `gorm:"column:bank_name" json:"bank_name"`                                 // 开户银行支行名称
	BankProvinceID          int     `gorm:"column:bank_province_id" json:"bank_province_id"`                   // 开户银行所在省id
	BankCityID              int     `gorm:"column:bank_city_id" json:"bank_city_id"`                           // 开户银行所在市id
	BankCountyID            int     `gorm:"column:bank_county_id" json:"bank_county_id"`                       // 开户银行所在县id
	BankTownID              int     `gorm:"column:bank_town_id" json:"bank_town_id"`                           // 开户银行所在镇id
	BankProvince            string  `gorm:"column:bank_province" json:"bank_province"`                         // 开户银行所在省
	BankCity                string  `gorm:"column:bank_city" json:"bank_city"`                                 // 开户银行所在市
	BankCounty              string  `gorm:"column:bank_county" json:"bank_county"`                             // 开户银行所在县
	BankTown                string  `gorm:"column:bank_town" json:"bank_town"`                                 // 开户银行所在镇
	BankImg                 string  `gorm:"column:bank_img" json:"bank_img"`                                   // 开户银行许可证电子版
	TaxesCertificateNum     string  `gorm:"column:taxes_certificate_num" json:"taxes_certificate_num"`         // 税务登记证号
	TaxesDistinguishNum     string  `gorm:"column:taxes_distinguish_num" json:"taxes_distinguish_num"`         // 纳税人识别号
	TaxesCertificateImg     string  `gorm:"column:taxes_certificate_img" json:"taxes_certificate_img"`         // 税务登记证书
	GoodsManagementCategory string  `gorm:"column:goods_management_category" json:"goods_management_category"` // 店铺经营类目
	ShopLevel               int     `gorm:"column:shop_level" json:"shop_level"`                               // 店铺等级
	ShopLevelApply          int     `gorm:"column:shop_level_apply" json:"shop_level_apply"`                   // 店铺等级申请
	StoreSpaceCapacity      float64 `gorm:"column:store_space_capacity" json:"store_space_capacity"`           // 店铺相册已用存储量
	ShopLogo                string  `gorm:"column:shop_logo" json:"shop_logo"`                                 // 店铺logo
	ShopBanner              string  `gorm:"column:shop_banner" json:"shop_banner"`                             // 店铺横幅
	ShopDesc                string  `gorm:"column:shop_desc" json:"shop_desc"`                                 // 店铺简介
	ShopRecommend           int     `gorm:"column:shop_recommend" json:"shop_recommend"`                       // 是否推荐
	ShopThemeid             int     `gorm:"column:shop_themeid" json:"shop_themeid"`                           // 店铺主题id
	ShopThemePath           string  `gorm:"column:shop_theme_path" json:"shop_theme_path"`                     // 店铺主题模版路径
	WapThemeid              int     `gorm:"column:wap_themeid" json:"wap_themeid"`                             // 店铺主题id
	WapThemePath            string  `gorm:"column:wap_theme_path" json:"wap_theme_path"`                       // wap店铺主题
	ShopCredit              float64 `gorm:"column:shop_credit" json:"shop_credit"`                             // 店铺信用
	ShopPraiseRate          float64 `gorm:"column:shop_praise_rate" json:"shop_praise_rate"`                   // 店铺好评率
	ShopDescriptionCredit   float64 `gorm:"column:shop_description_credit" json:"shop_description_credit"`     // 店铺描述相符度
	ShopServiceCredit       float64 `gorm:"column:shop_service_credit" json:"shop_service_credit"`             // 服务态度分数
	ShopDeliveryCredit      float64 `gorm:"column:shop_delivery_credit" json:"shop_delivery_credit"`           // 发货速度分数
	ShopCollect             int     `gorm:"column:shop_collect" json:"shop_collect"`                           // 店铺收藏数
	GoodsNum                int     `gorm:"column:goods_num" json:"goods_num"`                                 // 店铺商品数
	ShopQq                  string  `gorm:"column:shop_qq" json:"shop_qq"`                                     // 店铺客服qq
	ShopCommission          float64 `gorm:"column:shop_commission" json:"shop_commission"`                     // 店铺佣金比例
	GoodsWarningCount       int     `gorm:"column:goods_warning_count" json:"goods_warning_count"`             // 货品预警数
	SelfOperated            int     `gorm:"column:self_operated" json:"self_operated"`                         // 是否自营 1:是 0:否
	Step                    int     `gorm:"column:step" json:"step"`                                           // 申请开店进度：1,2.3.4
	OrdinReceiptStatus      int     `gorm:"column:ordin_receipt_status" json:"ordin_receipt_status"`           // 是否允许开具增值税普通发票 0：否，1：是
	ElecReceiptStatus       int     `gorm:"column:elec_receipt_status" json:"elec_receipt_status"`             // 是否允许开具电子普通发票 0：否，1：是
	TaxReceiptStatus        int     `gorm:"column:tax_receipt_status" json:"tax_receipt_status"`               // 是否允许开具增值税专用发票 0：否，1：是
}

// TableName get sql table name.获取数据库表名
func (m *EsShopDetail) TableName() string {
	return "es_shop_detail"
}
