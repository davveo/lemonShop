package mgr

import (
	"context"
	"fmt"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"

	"gorm.io/gorm"
)

type MemberMgr struct {
	*_BaseMgr
}

// NewMemberMgr open func
func NewMemberMgr(db db.Repo) *MemberMgr {
	if db == nil {
		panic(fmt.Errorf("NewMemberMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &MemberMgr{_BaseMgr: &_BaseMgr{rdb: db.GetDbR().Table("es_member"), wdb: db.GetDbW().Table("es_member"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *MemberMgr) GetTableName() string {
	return "es_member"
}

// Reset 重置gorm会话
func (obj *MemberMgr) Reset() *MemberMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *MemberMgr) Get() (result models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *MemberMgr) Gets() (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *MemberMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithMemberID member_id获取 会员ID
func (obj *MemberMgr) WithMemberID(memberID int) Option {
	return optionFunc(func(o *options) { o.query["member_id"] = memberID })
}

// WithUname uname获取 会员登陆用户名
func (obj *MemberMgr) WithUname(uname string) Option {
	return optionFunc(func(o *options) { o.query["uname"] = uname })
}

// WithEmail email获取 邮箱
func (obj *MemberMgr) WithEmail(email string) Option {
	return optionFunc(func(o *options) { o.query["email"] = email })
}

// WithPassword password获取 会员登陆密码
func (obj *MemberMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithCreateTime create_time获取 会员注册时间
func (obj *MemberMgr) WithCreateTime(createTime int64) Option {
	return optionFunc(func(o *options) { o.query["create_time"] = createTime })
}

// WithSex sex获取 会员性别 1：男，0：女
func (obj *MemberMgr) WithSex(sex int) Option {
	return optionFunc(func(o *options) { o.query["sex"] = sex })
}

// WithBirthday birthday获取 会员生日
func (obj *MemberMgr) WithBirthday(birthday int64) Option {
	return optionFunc(func(o *options) { o.query["birthday"] = birthday })
}

// WithProvinceID province_id获取 所属省份ID
func (obj *MemberMgr) WithProvinceID(provinceID int) Option {
	return optionFunc(func(o *options) { o.query["province_id"] = provinceID })
}

// WithCityID city_id获取 所属城市ID
func (obj *MemberMgr) WithCityID(cityID int) Option {
	return optionFunc(func(o *options) { o.query["city_id"] = cityID })
}

// WithCountyID county_id获取 所属县(区)ID
func (obj *MemberMgr) WithCountyID(countyID int) Option {
	return optionFunc(func(o *options) { o.query["county_id"] = countyID })
}

// WithTownID town_id获取 所属城镇ID
func (obj *MemberMgr) WithTownID(townID int) Option {
	return optionFunc(func(o *options) { o.query["town_id"] = townID })
}

// WithProvince province获取 所属省份名称
func (obj *MemberMgr) WithProvince(province string) Option {
	return optionFunc(func(o *options) { o.query["province"] = province })
}

// WithCity city获取 所属城市名称
func (obj *MemberMgr) WithCity(city string) Option {
	return optionFunc(func(o *options) { o.query["city"] = city })
}

// WithCounty county获取 所属县(区)名称
func (obj *MemberMgr) WithCounty(county string) Option {
	return optionFunc(func(o *options) { o.query["county"] = county })
}

// WithTown town获取 所属城镇名称
func (obj *MemberMgr) WithTown(town string) Option {
	return optionFunc(func(o *options) { o.query["town"] = town })
}

// WithAddress address获取 详细地址
func (obj *MemberMgr) WithAddress(address string) Option {
	return optionFunc(func(o *options) { o.query["address"] = address })
}

// WithMobile mobile获取 手机号码
func (obj *MemberMgr) WithMobile(mobile string) Option {
	return optionFunc(func(o *options) { o.query["mobile"] = mobile })
}

// WithTel tel获取 座机号码
func (obj *MemberMgr) WithTel(tel string) Option {
	return optionFunc(func(o *options) { o.query["tel"] = tel })
}

// WithGradePoint grade_point获取 等级积分
func (obj *MemberMgr) WithGradePoint(gradePoint int64) Option {
	return optionFunc(func(o *options) { o.query["grade_point"] = gradePoint })
}

// WithMsn msn获取 会员MSN
func (obj *MemberMgr) WithMsn(msn string) Option {
	return optionFunc(func(o *options) { o.query["msn"] = msn })
}

// WithRemark remark获取 会员备注
func (obj *MemberMgr) WithRemark(remark string) Option {
	return optionFunc(func(o *options) { o.query["remark"] = remark })
}

// WithLastLogin last_login获取 上次登陆时间
func (obj *MemberMgr) WithLastLogin(lastLogin int64) Option {
	return optionFunc(func(o *options) { o.query["last_login"] = lastLogin })
}

// WithLoginCount login_count获取 登陆次数
func (obj *MemberMgr) WithLoginCount(loginCount int) Option {
	return optionFunc(func(o *options) { o.query["login_count"] = loginCount })
}

// WithIsCheked is_cheked获取 邮件是否已验证
func (obj *MemberMgr) WithIsCheked(isCheked int) Option {
	return optionFunc(func(o *options) { o.query["is_cheked"] = isCheked })
}

// WithRegisterIP register_ip获取 注册IP地址
func (obj *MemberMgr) WithRegisterIP(registerIP string) Option {
	return optionFunc(func(o *options) { o.query["register_ip"] = registerIP })
}

// WithRecommendPointState recommend_point_state获取 是否已经完成了推荐积分
func (obj *MemberMgr) WithRecommendPointState(recommendPointState int) Option {
	return optionFunc(func(o *options) { o.query["recommend_point_state"] = recommendPointState })
}

// WithInfoFull info_full获取 会员信息是否完善
func (obj *MemberMgr) WithInfoFull(infoFull int) Option {
	return optionFunc(func(o *options) { o.query["info_full"] = infoFull })
}

// WithFindCode find_code获取 find_code
func (obj *MemberMgr) WithFindCode(findCode string) Option {
	return optionFunc(func(o *options) { o.query["find_code"] = findCode })
}

// WithFace face获取 会员头像
func (obj *MemberMgr) WithFace(face string) Option {
	return optionFunc(func(o *options) { o.query["face"] = face })
}

// WithMidentity midentity获取 身份证号
func (obj *MemberMgr) WithMidentity(midentity int) Option {
	return optionFunc(func(o *options) { o.query["midentity"] = midentity })
}

// WithDisabled disabled获取 会员状态
func (obj *MemberMgr) WithDisabled(disabled int) Option {
	return optionFunc(func(o *options) { o.query["disabled"] = disabled })
}

// WithShopID shop_id获取 店铺ID
func (obj *MemberMgr) WithShopID(shopID int) Option {
	return optionFunc(func(o *options) { o.query["shop_id"] = shopID })
}

// WithHaveShop have_shop获取 是否开通店铺
func (obj *MemberMgr) WithHaveShop(haveShop int) Option {
	return optionFunc(func(o *options) { o.query["have_shop"] = haveShop })
}

// WithConsumPoint consum_point获取 消费积分
func (obj *MemberMgr) WithConsumPoint(consumPoint int64) Option {
	return optionFunc(func(o *options) { o.query["consum_point"] = consumPoint })
}

// WithNickname nickname获取 昵称
func (obj *MemberMgr) WithNickname(nickname string) Option {
	return optionFunc(func(o *options) { o.query["nickname"] = nickname })
}

// GetByOption 功能选项模式获取
func (obj *MemberMgr) GetByOption(opts ...Option) (result models.EsMember, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *MemberMgr) GetByOptions(opts ...Option) (results []*models.EsMember, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *MemberMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]models.EsMember, 0)
	var count int64 // 统计总的记录数
	query := obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where(options.query)
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

// GetFromMemberID 通过member_id获取内容 会员ID
func (obj *MemberMgr) GetFromMemberID(memberID int) (result models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`member_id` = ?", memberID).First(&result).Error

	return
}

// GetBatchFromMemberID 批量查找 会员ID
func (obj *MemberMgr) GetBatchFromMemberID(memberIDs []int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`member_id` IN (?)", memberIDs).Find(&results).Error

	return
}

// GetFromUname 通过uname获取内容 会员登陆用户名
func (obj *MemberMgr) GetFromUname(uname string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`uname` = ?", uname).Find(&results).Error

	return
}

// GetBatchFromUname 批量查找 会员登陆用户名
func (obj *MemberMgr) GetBatchFromUname(unames []string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`uname` IN (?)", unames).Find(&results).Error

	return
}

// GetFromEmail 通过email获取内容 邮箱
func (obj *MemberMgr) GetFromEmail(email string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`email` = ?", email).Find(&results).Error

	return
}

// GetBatchFromEmail 批量查找 邮箱
func (obj *MemberMgr) GetBatchFromEmail(emails []string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`email` IN (?)", emails).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容 会员登陆密码
func (obj *MemberMgr) GetFromPassword(password string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找 会员登陆密码
func (obj *MemberMgr) GetBatchFromPassword(passwords []string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromCreateTime 通过create_time获取内容 会员注册时间
func (obj *MemberMgr) GetFromCreateTime(createTime int64) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`create_time` = ?", createTime).Find(&results).Error

	return
}

// GetBatchFromCreateTime 批量查找 会员注册时间
func (obj *MemberMgr) GetBatchFromCreateTime(createTimes []int64) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`create_time` IN (?)", createTimes).Find(&results).Error

	return
}

// GetFromSex 通过sex获取内容 会员性别 1：男，0：女
func (obj *MemberMgr) GetFromSex(sex int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`sex` = ?", sex).Find(&results).Error

	return
}

// GetBatchFromSex 批量查找 会员性别 1：男，0：女
func (obj *MemberMgr) GetBatchFromSex(sexs []int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`sex` IN (?)", sexs).Find(&results).Error

	return
}

// GetFromBirthday 通过birthday获取内容 会员生日
func (obj *MemberMgr) GetFromBirthday(birthday int64) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`birthday` = ?", birthday).Find(&results).Error

	return
}

// GetBatchFromBirthday 批量查找 会员生日
func (obj *MemberMgr) GetBatchFromBirthday(birthdays []int64) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`birthday` IN (?)", birthdays).Find(&results).Error

	return
}

// GetFromProvinceID 通过province_id获取内容 所属省份ID
func (obj *MemberMgr) GetFromProvinceID(provinceID int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`province_id` = ?", provinceID).Find(&results).Error

	return
}

// GetBatchFromProvinceID 批量查找 所属省份ID
func (obj *MemberMgr) GetBatchFromProvinceID(provinceIDs []int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`province_id` IN (?)", provinceIDs).Find(&results).Error

	return
}

// GetFromCityID 通过city_id获取内容 所属城市ID
func (obj *MemberMgr) GetFromCityID(cityID int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`city_id` = ?", cityID).Find(&results).Error

	return
}

// GetBatchFromCityID 批量查找 所属城市ID
func (obj *MemberMgr) GetBatchFromCityID(cityIDs []int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`city_id` IN (?)", cityIDs).Find(&results).Error

	return
}

// GetFromCountyID 通过county_id获取内容 所属县(区)ID
func (obj *MemberMgr) GetFromCountyID(countyID int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`county_id` = ?", countyID).Find(&results).Error

	return
}

// GetBatchFromCountyID 批量查找 所属县(区)ID
func (obj *MemberMgr) GetBatchFromCountyID(countyIDs []int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`county_id` IN (?)", countyIDs).Find(&results).Error

	return
}

// GetFromTownID 通过town_id获取内容 所属城镇ID
func (obj *MemberMgr) GetFromTownID(townID int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`town_id` = ?", townID).Find(&results).Error

	return
}

// GetBatchFromTownID 批量查找 所属城镇ID
func (obj *MemberMgr) GetBatchFromTownID(townIDs []int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`town_id` IN (?)", townIDs).Find(&results).Error

	return
}

// GetFromProvince 通过province获取内容 所属省份名称
func (obj *MemberMgr) GetFromProvince(province string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`province` = ?", province).Find(&results).Error

	return
}

// GetBatchFromProvince 批量查找 所属省份名称
func (obj *MemberMgr) GetBatchFromProvince(provinces []string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`province` IN (?)", provinces).Find(&results).Error

	return
}

// GetFromCity 通过city获取内容 所属城市名称
func (obj *MemberMgr) GetFromCity(city string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`city` = ?", city).Find(&results).Error

	return
}

// GetBatchFromCity 批量查找 所属城市名称
func (obj *MemberMgr) GetBatchFromCity(citys []string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`city` IN (?)", citys).Find(&results).Error

	return
}

// GetFromCounty 通过county获取内容 所属县(区)名称
func (obj *MemberMgr) GetFromCounty(county string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`county` = ?", county).Find(&results).Error

	return
}

// GetBatchFromCounty 批量查找 所属县(区)名称
func (obj *MemberMgr) GetBatchFromCounty(countys []string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`county` IN (?)", countys).Find(&results).Error

	return
}

// GetFromTown 通过town获取内容 所属城镇名称
func (obj *MemberMgr) GetFromTown(town string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`town` = ?", town).Find(&results).Error

	return
}

// GetBatchFromTown 批量查找 所属城镇名称
func (obj *MemberMgr) GetBatchFromTown(towns []string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`town` IN (?)", towns).Find(&results).Error

	return
}

// GetFromAddress 通过address获取内容 详细地址
func (obj *MemberMgr) GetFromAddress(address string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`address` = ?", address).Find(&results).Error

	return
}

// GetBatchFromAddress 批量查找 详细地址
func (obj *MemberMgr) GetBatchFromAddress(addresss []string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`address` IN (?)", addresss).Find(&results).Error

	return
}

// GetFromMobile 通过mobile获取内容 手机号码
func (obj *MemberMgr) GetFromMobile(mobile string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`mobile` = ?", mobile).Find(&results).Error

	return
}

// GetBatchFromMobile 批量查找 手机号码
func (obj *MemberMgr) GetBatchFromMobile(mobiles []string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`mobile` IN (?)", mobiles).Find(&results).Error

	return
}

// GetFromTel 通过tel获取内容 座机号码
func (obj *MemberMgr) GetFromTel(tel string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`tel` = ?", tel).Find(&results).Error

	return
}

// GetBatchFromTel 批量查找 座机号码
func (obj *MemberMgr) GetBatchFromTel(tels []string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`tel` IN (?)", tels).Find(&results).Error

	return
}

// GetFromGradePoint 通过grade_point获取内容 等级积分
func (obj *MemberMgr) GetFromGradePoint(gradePoint int64) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`grade_point` = ?", gradePoint).Find(&results).Error

	return
}

// GetBatchFromGradePoint 批量查找 等级积分
func (obj *MemberMgr) GetBatchFromGradePoint(gradePoints []int64) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`grade_point` IN (?)", gradePoints).Find(&results).Error

	return
}

// GetFromMsn 通过msn获取内容 会员MSN
func (obj *MemberMgr) GetFromMsn(msn string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`msn` = ?", msn).Find(&results).Error

	return
}

// GetBatchFromMsn 批量查找 会员MSN
func (obj *MemberMgr) GetBatchFromMsn(msns []string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`msn` IN (?)", msns).Find(&results).Error

	return
}

// GetFromRemark 通过remark获取内容 会员备注
func (obj *MemberMgr) GetFromRemark(remark string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`remark` = ?", remark).Find(&results).Error

	return
}

// GetBatchFromRemark 批量查找 会员备注
func (obj *MemberMgr) GetBatchFromRemark(remarks []string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`remark` IN (?)", remarks).Find(&results).Error

	return
}

// GetFromLastLogin 通过last_login获取内容 上次登陆时间
func (obj *MemberMgr) GetFromLastLogin(lastLogin int64) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`last_login` = ?", lastLogin).Find(&results).Error

	return
}

// GetBatchFromLastLogin 批量查找 上次登陆时间
func (obj *MemberMgr) GetBatchFromLastLogin(lastLogins []int64) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`last_login` IN (?)", lastLogins).Find(&results).Error

	return
}

// GetFromLoginCount 通过login_count获取内容 登陆次数
func (obj *MemberMgr) GetFromLoginCount(loginCount int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`login_count` = ?", loginCount).Find(&results).Error

	return
}

// GetBatchFromLoginCount 批量查找 登陆次数
func (obj *MemberMgr) GetBatchFromLoginCount(loginCounts []int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`login_count` IN (?)", loginCounts).Find(&results).Error

	return
}

// GetFromIsCheked 通过is_cheked获取内容 邮件是否已验证
func (obj *MemberMgr) GetFromIsCheked(isCheked int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`is_cheked` = ?", isCheked).Find(&results).Error

	return
}

// GetBatchFromIsCheked 批量查找 邮件是否已验证
func (obj *MemberMgr) GetBatchFromIsCheked(isChekeds []int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`is_cheked` IN (?)", isChekeds).Find(&results).Error

	return
}

// GetFromRegisterIP 通过register_ip获取内容 注册IP地址
func (obj *MemberMgr) GetFromRegisterIP(registerIP string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`register_ip` = ?", registerIP).Find(&results).Error

	return
}

// GetBatchFromRegisterIP 批量查找 注册IP地址
func (obj *MemberMgr) GetBatchFromRegisterIP(registerIPs []string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`register_ip` IN (?)", registerIPs).Find(&results).Error

	return
}

// GetFromRecommendPointState 通过recommend_point_state获取内容 是否已经完成了推荐积分
func (obj *MemberMgr) GetFromRecommendPointState(recommendPointState int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`recommend_point_state` = ?", recommendPointState).Find(&results).Error

	return
}

// GetBatchFromRecommendPointState 批量查找 是否已经完成了推荐积分
func (obj *MemberMgr) GetBatchFromRecommendPointState(recommendPointStates []int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`recommend_point_state` IN (?)", recommendPointStates).Find(&results).Error

	return
}

// GetFromInfoFull 通过info_full获取内容 会员信息是否完善
func (obj *MemberMgr) GetFromInfoFull(infoFull int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`info_full` = ?", infoFull).Find(&results).Error

	return
}

// GetBatchFromInfoFull 批量查找 会员信息是否完善
func (obj *MemberMgr) GetBatchFromInfoFull(infoFulls []int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`info_full` IN (?)", infoFulls).Find(&results).Error

	return
}

// GetFromFindCode 通过find_code获取内容 find_code
func (obj *MemberMgr) GetFromFindCode(findCode string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`find_code` = ?", findCode).Find(&results).Error

	return
}

// GetBatchFromFindCode 批量查找 find_code
func (obj *MemberMgr) GetBatchFromFindCode(findCodes []string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`find_code` IN (?)", findCodes).Find(&results).Error

	return
}

// GetFromFace 通过face获取内容 会员头像
func (obj *MemberMgr) GetFromFace(face string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`face` = ?", face).Find(&results).Error

	return
}

// GetBatchFromFace 批量查找 会员头像
func (obj *MemberMgr) GetBatchFromFace(faces []string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`face` IN (?)", faces).Find(&results).Error

	return
}

// GetFromMidentity 通过midentity获取内容 身份证号
func (obj *MemberMgr) GetFromMidentity(midentity int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`midentity` = ?", midentity).Find(&results).Error

	return
}

// GetBatchFromMidentity 批量查找 身份证号
func (obj *MemberMgr) GetBatchFromMidentity(midentitys []int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`midentity` IN (?)", midentitys).Find(&results).Error

	return
}

// GetFromDisabled 通过disabled获取内容 会员状态
func (obj *MemberMgr) GetFromDisabled(disabled int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`disabled` = ?", disabled).Find(&results).Error

	return
}

// GetBatchFromDisabled 批量查找 会员状态
func (obj *MemberMgr) GetBatchFromDisabled(disableds []int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`disabled` IN (?)", disableds).Find(&results).Error

	return
}

// GetFromShopID 通过shop_id获取内容 店铺ID
func (obj *MemberMgr) GetFromShopID(shopID int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`shop_id` = ?", shopID).Find(&results).Error

	return
}

// GetBatchFromShopID 批量查找 店铺ID
func (obj *MemberMgr) GetBatchFromShopID(shopIDs []int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`shop_id` IN (?)", shopIDs).Find(&results).Error

	return
}

// GetFromHaveShop 通过have_shop获取内容 是否开通店铺
func (obj *MemberMgr) GetFromHaveShop(haveShop int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`have_shop` = ?", haveShop).Find(&results).Error

	return
}

// GetBatchFromHaveShop 批量查找 是否开通店铺
func (obj *MemberMgr) GetBatchFromHaveShop(haveShops []int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`have_shop` IN (?)", haveShops).Find(&results).Error

	return
}

// GetFromConsumPoint 通过consum_point获取内容 消费积分
func (obj *MemberMgr) GetFromConsumPoint(consumPoint int64) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`consum_point` = ?", consumPoint).Find(&results).Error

	return
}

// GetBatchFromConsumPoint 批量查找 消费积分
func (obj *MemberMgr) GetBatchFromConsumPoint(consumPoints []int64) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`consum_point` IN (?)", consumPoints).Find(&results).Error

	return
}

// GetFromNickname 通过nickname获取内容 昵称
func (obj *MemberMgr) GetFromNickname(nickname string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`nickname` = ?", nickname).Find(&results).Error

	return
}

// GetBatchFromNickname 批量查找 昵称
func (obj *MemberMgr) GetBatchFromNickname(nicknames []string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`nickname` IN (?)", nicknames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *MemberMgr) FetchByPrimaryKey(memberID int) (result models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`member_id` = ?", memberID).First(&result).Error

	return
}

// FetchIndexByIndMemberUname  获取多个内容
func (obj *MemberMgr) FetchIndexByIndMemberUname(uname string, email string) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`uname` = ? AND `email` = ?", uname, email).Find(&results).Error

	return
}

// FetchIndexByIndMemberB2b2c  获取多个内容
func (obj *MemberMgr) FetchIndexByIndMemberB2b2c(shopID int, haveShop int) (results []*models.EsMember, err error) {
	err = obj.rdb.WithContext(obj.ctx).Model(models.EsMember{}).Where("`shop_id` = ? AND `have_shop` = ?", shopID, haveShop).Find(&results).Error

	return
}
