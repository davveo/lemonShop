package system

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/davveo/lemonShop/app/consts"
	"github.com/davveo/lemonShop/app/entity"
	"github.com/davveo/lemonShop/mgr"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/models/vo"
	"github.com/davveo/lemonShop/pkg/cache"
	"github.com/davveo/lemonShop/pkg/db"
	"github.com/davveo/lemonShop/pkg/string_util"
	"gorm.io/gorm"
)

var _ SmsPlatformService = (*smsPlatformService)(nil)

type (
	SmsPlatformService interface {
		i()
		List(page, pageSize int) (*entity.List, error)
		Add(smsPlatform *vo.SmsPlatformVo) (*vo.SmsPlatformVo, error)
		Edit(smsPlatform *vo.SmsPlatformVo) (*vo.SmsPlatformVo, error)
		Get(id int) (*models.EsSmsPlatform, error)
		GetByBean(bean string) (*models.EsSmsPlatform, error)
		OpenPlatform(bean string) error
		GetConfig(bean string) (*vo.SmsPlatformVo, error)
		GetOpen() (*vo.SmsPlatformVo, error)
	}
	smsPlatformService struct {
		smsPlatformMgr *mgr.SmsPlatformMgr
	}
)

func (s smsPlatformService) List(page, pageSize int) (*entity.List, error) {
	//TODO implement me
	panic("implement me")
}

func (s smsPlatformService) Add(smsPlatform *vo.SmsPlatformVo) (*vo.SmsPlatformVo, error) {
	//TODO implement me
	panic("implement me")
}

func (s smsPlatformService) Edit(smsPlatform *vo.SmsPlatformVo) (*vo.SmsPlatformVo, error) {
	//TODO implement me
	panic("implement me")
}

func (s smsPlatformService) Get(id int) (*models.EsSmsPlatform, error) {
	//TODO implement me
	panic("implement me")
}

func (s smsPlatformService) GetByBean(bean string) (*models.EsSmsPlatform, error) {
	//TODO implement me
	panic("implement me")
}

func (s smsPlatformService) OpenPlatform(bean string) error {
	//TODO implement me
	panic("implement me")
}

func (s smsPlatformService) GetConfig(bean string) (*vo.SmsPlatformVo, error) {
	//TODO implement me
	panic("implement me")
}

func (s smsPlatformService) GetOpen() (*vo.SmsPlatformVo, error) {
	var platformVo vo.SmsPlatformVo
	smsPlatformVoStr, err := cache.Cache.Get(consts.SPlATFORM)
	if err != nil {
		return nil, err
	}
	if string_util.IsEmpty(smsPlatformVoStr) {
		smsPlatform, err := s.smsPlatformMgr.GetByOption(s.smsPlatformMgr.WithOpen(1))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

		if smsPlatform == nil {
			return nil, fmt.Errorf("未找到可用的短信网关")
		}
		platformVo = vo.SmsPlatformVo{
			Config: smsPlatform.Config,
			Bean:   smsPlatform.Bean,
		}
		platformVoStr, err := json.Marshal(platformVo)
		if err != nil {
			return nil, err
		}
		cache.Cache.Set(consts.SPlATFORM, string(platformVoStr), time.Duration(0))
	} else {
		if err := json.Unmarshal([]byte(smsPlatformVoStr), &platformVo); err != nil {
			return nil, err
		}
	}
	return &platformVo, nil
}

func (s smsPlatformService) i() {}

func NewSmsPlatformService() SmsPlatformService {
	return &smsPlatformService{
		smsPlatformMgr: mgr.NewSmsPlatformMgr(db.GRpo),
	}
}
