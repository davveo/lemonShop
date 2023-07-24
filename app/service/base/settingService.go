package base

import (
	"errors"
	"fmt"

	"github.com/davveo/lemonShop/app/consts"
	"github.com/davveo/lemonShop/mgr"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/cache"
	"github.com/davveo/lemonShop/pkg/db"
	"github.com/davveo/lemonShop/pkg/string_util"
	"gorm.io/gorm"
)

var _ SettingService = (*settingService)(nil)

type (
	SettingService interface {
		i()
		Save(group, settings string) error
		Get(group string) (string, error)
	}
	settingService struct {
		settingsMgr *mgr.SettingsMgr
	}
)

func (s settingService) Save(group, settings string) error {
	var err error
	setting, err := s.settingsMgr.GetByOption(s.settingsMgr.WithCfgGroup(group))
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	// 记录不存在
	if setting == nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = s.settingsMgr.Create(&models.EsSettings{
			CfgGroup: group,
			CfgValue: settings,
		})
	} else {
		err = s.settingsMgr.Update(map[string]interface{}{
			"cfg_value": settings,
		}, s.settingsMgr.WithID(setting.ID))
	}
	if err != nil {
		return err
	}
	// 清除缓存
	if success := cache.Cache.Del(consts.SETTING + group); !success {
		return fmt.Errorf("清除缓存失败")
	}
	return nil
}

func (s settingService) Get(group string) (string, error) {
	cacheValue, err := cache.Cache.Get(consts.SETTING + group)
	if err != nil {
		return "", err
	}

	if string_util.IsEmpty(cacheValue) {
		obj, err := s.settingsMgr.GetByOption(s.settingsMgr.WithCfgGroup(group))
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return "", err
		}
		if obj == nil || errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil
		}
		if err = cache.Cache.Set(consts.SETTING+group, obj.CfgValue, 0); err != nil {
			return "", err
		}
		return obj.CfgValue, nil
	}
	return cacheValue, nil
}

func (s settingService) i() {}

func NewSettingService() SettingService {
	return &settingService{
		settingsMgr: mgr.NewSettingsMgr(db.GRpo),
	}
}
