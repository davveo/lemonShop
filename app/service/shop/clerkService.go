package shop

import "github.com/davveo/lemonShop/models"

var _ ClerkService = (*clerkService)(nil)

type (
	ClerkService interface {
		i()
		// GetClerkByMemberId 根据会员id查询店员
		GetClerkByMemberId(memberId int64) (*models.EsClerk, error)
	}
	clerkService struct {
	}
)

func (c clerkService) GetClerkByMemberId(memberId int64) (*models.EsClerk, error) {
	//TODO implement me
	panic("implement me")
}

func (c clerkService) i() {}

func NewClerkService() ClerkService {
	return &clerkService{}
}
