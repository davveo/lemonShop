package passport

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/davveo/lemonShop/pkg/common"

	"github.com/davveo/lemonShop/app/entity"

	memberClient "github.com/davveo/lemonShop/app/client/member"
	systemClient "github.com/davveo/lemonShop/app/client/system"
	"github.com/davveo/lemonShop/app/consts"
	"github.com/davveo/lemonShop/app/service/member"
	"github.com/davveo/lemonShop/app/service/shop"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/cache"
	"github.com/davveo/lemonShop/pkg/util"
	"gorm.io/gorm"
)

var _ PassportService = (*passportService)(nil)

type (
	PassportService interface {
		i()
		// SendRegisterSmsCode 发送注册短信验证码
		SendRegisterSmsCode(mobile string) error
		// SendFindPasswordCode 发送找回密码短信验证码
		SendFindPasswordCode(mobile string) error
		// SendLoginSmsCode 发送登录短信验证码
		SendLoginSmsCode(mobile string) error
		// SendSmsCode 发送短信验证码
		SendSmsCode(mobile string) error
		// CreateToken 生成token
		CreateToken(member models.Member, time int64) (string, error)
		// ExchangeToken 通过refreshToken 获取 accessToken
		ExchangeToken(ctx context.Context, refreshToken string) (string, error)
		// ClearSign 清除标记缓存
		ClearSign(mobile, scene string) error
	}

	passportService struct {
		smsClient       systemClient.SmsClient
		memberService   member.MemberService
		clerkService    shop.ClerkService
		shopClient      memberClient.ShopClient
		shopRoleService shop.ShopRoleService
	}
)

func (p passportService) SendRegisterSmsCode(mobile string) error {
	if util.IsMobile(mobile) {
		return fmt.Errorf("手机号码格式不正确")
	}
	//校验会员是否存在
	m, err := p.memberService.GetMemberByMobile(mobile)
	if err != nil {
		return err
	}
	if m != nil {
		return fmt.Errorf("该手机号已经被占用")
	}
	return p.smsClient.SendSmsMessage("注册", mobile, consts.REGISTER)
}

func (p passportService) SendFindPasswordCode(mobile string) error {
	//校验会员是否存在
	if _, err := p.memberService.GetMemberByMobile(mobile); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("该手机号未注册")
		}
		return err
	}
	return p.smsClient.SendSmsMessage("找回密码", mobile, consts.VALIDATE_MOBILE)
}

func (p passportService) SendLoginSmsCode(mobile string) error {
	if !util.IsMobile(mobile) {
		return fmt.Errorf("手机号码格式不正确")
	}
	//校验会员是否存在
	m, err := p.memberService.GetMemberByMobile(mobile)
	if err != nil {
		return err
	}
	if m != nil {
		return fmt.Errorf("该手机号未注册")
	}
	return p.smsClient.SendSmsMessage("登录", mobile, consts.LOGIN)
}

func (p passportService) SendSmsCode(mobile string) error {
	if !util.IsMobile(mobile) {
		return fmt.Errorf("手机号码格式不正确")
	}
	return p.smsClient.SendSmsMessage("添加店员", mobile, consts.ADD_CLERK)
}

func (p passportService) CreateToken(member models.Member, t int64) (string, error) {
	clerkDb, err := p.clerkService.GetClerkByMemberId(int64(member.MemberID))
	if err != nil {
		return "", err
	}
	// 如果店员不为空，则说明他是店铺管理员，需要赋值商家权限
	if clerkDb != nil {
		clerk := entity.NewClerk()
		shopVO, err := p.shopClient.GetShop(clerkDb.ShopID)
		if err != nil {
			return "", err
		}
		clerk.SellerId = shopVO.ShopId
		clerk.SellerName = shopVO.ShopName
		clerk.UserName = member.Uname
		clerk.Uid = int64(member.MemberID)
		clerk.ClerkName = clerkDb.ClerkName
		clerk.ClerkId = int64(clerkDb.ClerkID)
		clerk.SelfOperated = shopVO.SelfOperated
		//如果是超级店员则赋值超级店员的权限，否则去查询权限赋值
		if clerkDb.Founder == 1 {
			clerk.Role = "SUPER_SELLER"
		} else {
			shopRole, err := p.shopRoleService.GetByRoleId(clerkDb.RoleID)
			if err != nil {
				return "", err
			}
			clerk.Role = shopRole.RoleName
		}
		clerkMap := common.StructToMap(clerk)
		claims := util.CustomClaims{
			Data: clerkMap,
			StandardClaims: jwt.StandardClaims{
				Subject:   string(consts.Clerk),
				NotBefore: time.Now().Unix() - 1000,
				ExpiresAt: time.Now().Unix() + 1000*t,
			},
		}
		token, err := util.NewJwt().CreateToken(claims)
		if err != nil {
			return "", err
		}
		return token, nil
	} else {
		//如果是会员，则赋值买家权限
		buyer := entity.Buyer{}
		buyer.Uid = int64(member.MemberID)
		buyer.UserName = member.Uname
		buyerMap := common.StructToMap(buyer)
		claims := util.CustomClaims{
			Data: buyerMap,
			StandardClaims: jwt.StandardClaims{
				Subject:   string(consts.Buyer),
				NotBefore: time.Now().Unix() - 1000,
				ExpiresAt: time.Now().Unix() + 1000*t,
			},
		}
		token, err := util.NewJwt().CreateToken(claims)
		if err != nil {
			return "", err
		}
		return token, nil
	}
}

func (p passportService) ExchangeToken(ctx context.Context, refreshToken string) (string, error) {
	if refreshToken == "" {
		return "", fmt.Errorf("当前会员不存在")
	}
	claims, err := util.NewJwt().ParseToken(refreshToken)
	if err != nil {
		return "", err
	}
	uid, ok := claims.Data["uid"].(int)
	if !ok {
		return "", fmt.Errorf("parse uid fail")
	}
	uuid, ok := ctx.Value("uuid").(string)
	if !ok {
		return "", fmt.Errorf("parse uuid fail")
	}
	//根据uid获取用户,获得当前会员是buyer还是seller
	m, err := p.memberService.Get(uid)
	if err != nil {
		return "", err
	}
	// 如果会员token刷新时，会员已经失效，则不颁发新的token
	if m.Disabled == -1 {
		return "", fmt.Errorf("当前token已经失效")
	}
	return uuid, nil
}

func (p passportService) ClearSign(mobile, scene string) error {
	key := fmt.Sprintf("%s_%s_%s", consts.MOBILE_VALIDATE, scene, mobile)
	if !cache.Cache.Del(key) {
		return fmt.Errorf("清除标记缓存")
	}
	return nil
}

func (p passportService) i() {}

func NewPassportService() PassportService {
	return &passportService{
		memberService:   member.NewMemberService(),
		clerkService:    shop.NewClerkService(),
		shopRoleService: shop.NewShopRoleService(),
		smsClient:       systemClient.NewSmsClient(),
		shopClient:      memberClient.NewShopClient(),
	}
}
