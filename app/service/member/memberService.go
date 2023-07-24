package member

import (
	"github.com/davveo/lemonShop/mgr"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"
)

var _ MemberService = (*memberService)(nil)

type (
	MemberService interface {
		i()
		// GetMemberByMobile 根据用户手机号码查询会员
		GetMemberByMobile(mobile string) (*models.Member, error)
		// Get 获取会员
		Get(id int) (*models.Member, error)
	}
	memberService struct {
		memberMgr *mgr.MemberMgr
	}
)

func (m memberService) Get(id int) (*models.Member, error) {
	return m.memberMgr.GetFromMemberID(id)
}

func (m memberService) i() {}

func (m memberService) GetMemberByMobile(mobile string) (*models.Member, error) {
	return m.memberMgr.GetByOption(m.memberMgr.WithMobile(mobile))
}

func NewMemberService() MemberService {
	return &memberService{
		memberMgr: mgr.NewMemberMgr(db.GRpo),
	}
}
