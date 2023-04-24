package models

// EsMember 会员表(es_member)
type EsMember struct {
	MemberID            int    `gorm:"primaryKey;column:member_id" json:"-"`                      // 会员ID
	Uname               string `gorm:"column:uname" json:"uname"`                                 // 会员登陆用户名
	Email               string `gorm:"column:email" json:"email"`                                 // 邮箱
	Password            string `gorm:"column:password" json:"password"`                           // 会员登陆密码
	CreateTime          int64  `gorm:"column:create_time" json:"create_time"`                     // 会员注册时间
	Sex                 int    `gorm:"column:sex" json:"sex"`                                     // 会员性别 1：男，0：女
	Birthday            int64  `gorm:"column:birthday" json:"birthday"`                           // 会员生日
	ProvinceID          int    `gorm:"column:province_id" json:"province_id"`                     // 所属省份ID
	CityID              int    `gorm:"column:city_id" json:"city_id"`                             // 所属城市ID
	CountyID            int    `gorm:"column:county_id" json:"county_id"`                         // 所属县(区)ID
	TownID              int    `gorm:"column:town_id" json:"town_id"`                             // 所属城镇ID
	Province            string `gorm:"column:province" json:"province"`                           // 所属省份名称
	City                string `gorm:"column:city" json:"city"`                                   // 所属城市名称
	County              string `gorm:"column:county" json:"county"`                               // 所属县(区)名称
	Town                string `gorm:"column:town" json:"town"`                                   // 所属城镇名称
	Address             string `gorm:"column:address" json:"address"`                             // 详细地址
	Mobile              string `gorm:"column:mobile" json:"mobile"`                               // 手机号码
	Tel                 string `gorm:"column:tel" json:"tel"`                                     // 座机号码
	GradePoint          int64  `gorm:"column:grade_point" json:"grade_point"`                     // 等级积分
	Msn                 string `gorm:"column:msn" json:"msn"`                                     // 会员MSN
	Remark              string `gorm:"column:remark" json:"remark"`                               // 会员备注
	LastLogin           int64  `gorm:"column:last_login" json:"last_login"`                       // 上次登陆时间
	LoginCount          int    `gorm:"column:login_count" json:"login_count"`                     // 登陆次数
	IsCheked            int    `gorm:"column:is_cheked" json:"is_cheked"`                         // 邮件是否已验证
	RegisterIP          string `gorm:"column:register_ip" json:"register_ip"`                     // 注册IP地址
	RecommendPointState int    `gorm:"column:recommend_point_state" json:"recommend_point_state"` // 是否已经完成了推荐积分
	InfoFull            int    `gorm:"column:info_full" json:"info_full"`                         // 会员信息是否完善
	FindCode            string `gorm:"column:find_code" json:"find_code"`                         // find_code
	Face                string `gorm:"column:face" json:"face"`                                   // 会员头像
	Midentity           int    `gorm:"column:midentity" json:"midentity"`                         // 身份证号
	Disabled            int    `gorm:"column:disabled" json:"disabled"`                           // 会员状态
	ShopID              int    `gorm:"column:shop_id" json:"shop_id"`                             // 店铺ID
	HaveShop            int    `gorm:"column:have_shop" json:"have_shop"`                         // 是否开通店铺
	ConsumPoint         int64  `gorm:"column:consum_point" json:"consum_point"`                   // 消费积分
	Nickname            string `gorm:"column:nickname" json:"nickname"`                           // 昵称
}

// TableName get sql table name.获取数据库表名
func (m *EsMember) TableName() string {
	return "es_member"
}
