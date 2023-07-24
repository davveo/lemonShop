package vo

type ShopVo struct {
	ShopId                  int     `json:"shopId"`                  // 店铺Id
	MemberId                int     `json:"memberId"`                // 会员Id
	MemberName              string  `json:"memberName"`              // 会员名称
	ShopName                string  `json:"shopName"`                //店铺名称
	ShopDisable             string  `json:"shopDisable"`             //店铺状态
	ShopCreateTime          int64   `json:"shopCreateTime"`          //店铺创建时间
	ShopEndTime             int64   `json:"shopEndTime"`             // 店铺关闭时间
	ShopProvinceId          int     `json:"shopProvinceId"`          /**店铺所在省id*/
	ShopCityId              int     `json:"shopCityId"`              /**店铺所在市id*/
	ShopCountyId            int     `json:"shopCountyId"`            /**店铺所在县id*/
	ShopTownId              int     `json:"shopTownId"`              /**店铺所在镇id*/
	ShopProvince            string  `json:"shopProvince"`            /**店铺所在省*/
	ShopCity                string  `json:"shopCity"`                /**店铺所在市*/
	ShopCounty              string  `json:"shopCounty"`              /**店铺所在县*/
	ShopTown                string  `json:"shopTown"`                /**店铺所在镇*/
	ShopAdd                 string  `json:"shopAdd"`                 /**店铺详细地址*/
	CompanyName             string  `json:"companyName"`             /**公司名称*/
	CompanyAddress          string  `json:"companyAddress"`          /**公司地址*/
	CompanyPhone            string  `json:"companyPhone"`            /**公司电话*/
	CompanyEmail            string  `json:"companyEmail"`            /**电子邮箱*/
	EmployeeNum             int     `json:"employeeNum"`             /**员工总数*/
	RegMoney                float64 `json:"regMoney"`                /**注册资金*/
	LinkName                string  `json:"linkName"`                /**联系人姓名*/
	LinkPhone               string  `json:"linkPhone"`               /**联系人电话*/
	LegalName               string  `json:"legalName"`               /**法人姓名*/
	LegalId                 string  `json:"legalId"`                 /**法人身份证*/
	LegalImg                string  `json:"legalImg"`                /**法人身份证照片*/
	LicenseNum              string  `json:"licenseNum"`              /**营业执照号*/
	LicenseProvinceId       int     `json:"licenseProvinceId"`       /**营业执照所在省id*/
	LicenseCityId           int     `json:"licenseCityId"`           /**营业执照所在市id*/
	LicenseCountyId         int     `json:"licenseCountyId"`         /**营业执照所在县id*/
	LicenseTownId           int     `json:"licenseTownId"`           /**营业执照所在镇id*/
	LicenseProvince         string  `json:"licenseProvince"`         /**营业执照所在省*/
	LicenseCity             string  `json:"licenseCity"`             /**营业执照所在市*/
	LicenseCounty           string  `json:"licenseCounty"`           /**营业执照所在县*/
	LicenseTown             string  `json:"licenseTown"`             /**营业执照所在镇*/
	LicenseAdd              string  `json:"licenseAdd"`              /**营业执照详细地址*/
	EstablishDate           int64   `json:"establishDate"`           /**成立日期*/
	LicenceStart            int64   `json:"licenceStart"`            /**营业执照有效期开始*/
	LicenceEnd              int64   `json:"licenceEnd"`              /**营业执照有效期结束*/
	Scope                   string  `json:"scope"`                   /**法定经营范围*/
	LicenceImg              string  `json:"licenceImg"`              /**营业执照电子版*/
	OrganizationCode        string  `json:"organizationCode"`        /**组织机构代码*/
	CodeImg                 string  `json:"codeImg"`                 /**组织机构电子版*/
	TaxesImg                string  `json:"taxesImg"`                /**一般纳税人证明电子版*/
	BankAccountName         string  `json:"bankAccountName"`         /**银行开户名*/
	BankNumber              string  `json:"bankNumber"`              /**银行开户账号*/
	BankName                string  `json:"bankName"`                /**开户银行支行名称*/
	BankProvinceId          int     `json:"bankProvinceId"`          /**开户银行所在省id*/
	BankCityId              int     `json:"bankCityId"`              /**开户银行所在市id*/
	BankCountyId            int     `json:"bankCountyId"`            /**开户银行所在县id*/
	BankTownId              int     `json:"bankTownId"`              /**开户银行所在镇id*/
	BankProvince            string  `json:"bankProvince"`            /**开户银行所在省*/
	BankCity                string  `json:"bankCity"`                /**开户银行所在市*/
	BankCounty              string  `json:"bankCounty"`              /**开户银行所在县*/
	BankTown                string  `json:"bankTown"`                /**开户银行所在镇*/
	BankImg                 string  `json:"bankImg"`                 /**开户银行许可证电子版*/
	TaxesCertificateNum     string  `json:"taxesCertificateNum"`     /**税务登记证号*/
	TaxesDistinguishNum     string  `json:"taxesDistinguishNum"`     /**纳税人识别号*/
	TaxesCertificateImg     string  `json:"taxesCertificateImg"`     /**税务登记证书*/
	GoodsManagementCategory string  `json:"goodsManagementCategory"` /**店铺经营类目*/
	ShopLevel               int     `json:"shopLevel"`               /**店铺等级*/
	ShopLevelApply          int     `json:"shopLevelApply"`          /**店铺等级申请*/
	StoreSpaceCapacity      float64 `json:"storeSpaceCapacity"`      /**店铺相册已用存储量*/
	ShopLogo                string  `json:"shopLogo"`                /**店铺logo*/
	ShopBanner              string  `json:"shopBanner"`              /**店铺横幅*/
	ShopDesc                string  `json:"shopDesc"`                /**店铺简介*/
	ShopRecommend           int     `json:"shopRecommend"`           /**是否推荐*/
	ShopThemeId             int     `json:"shopThemeId"`             /**店铺主题id*/
	ShopThemePath           string  `json:"shopThemePath"`           /**店铺主题模版路径*/
	WapThemeId              int     `json:"wapThemeId"`              /**店铺主题id*/
	WapThemePath            string  `json:"wapThemePath"`            /**wap店铺主题*/
	ShopCredit              float64 `json:"shopCredit"`              /**店铺信用*/
	ShopPraiseRate          float64 `json:"shopPraiseRate"`          /**店铺好评率*/
	ShopDescriptionCredit   float64 `json:"shopDescriptionCredit"`   /**店铺描述相符度*/
	ShopServiceCredit       float64 `json:"shopServiceCredit"`       /**服务态度分数*/
	ShopDeliveryCredit      float64 `json:"shopDeliveryCredit"`      /**发货速度分数*/
	ShopCollect             int     `json:"shopCollect"`             /**店铺收藏数*/
	GoodsNum                int     `json:"goodsNum"`                /**店铺商品数*/
	ShopQq                  string  `json:"shopQq"`                  /**店铺客服qq*/
	ShopCommission          float64 `json:"shopCommission"`          /**店铺佣金比例*/
	GoodsWarningCount       int     `json:"goodsWarningCount"`       /**货品预警数*/
	SelfOperated            int     `json:"selfOperated"`            /**是否自营*/
	Step                    int     `json:"step"`                    /**申请开店进度*/
	OrdinReceiptStatus      int     `json:"ordinReceiptStatus"`      /**是否允许开具增值税普通发票 0：否，1：是*/
	ElecReceiptStatus       int     `json:"elecReceiptStatus"`       /**是否允许开具电子普通发票 0：否，1：是*/
	TaxReceiptStatus        int     `json:"taxReceiptStatus"`        /**是否允许开具增值税专用发票 0：否，1：是*/
}
