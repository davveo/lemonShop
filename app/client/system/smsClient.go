package system

import (
	"encoding/json"
	"time"

	"github.com/davveo/lemonShop/pkg/logger"

	"github.com/davveo/lemonShop/app/service/system"

	"github.com/davveo/lemonShop/conf"

	"github.com/davveo/lemonShop/pkg/cache"

	"github.com/davveo/lemonShop/app/event_handler"

	"github.com/davveo/lemonShop/pkg/event"

	"github.com/davveo/lemonShop/pkg/string_util"

	"github.com/davveo/lemonShop/pkg/random"

	"github.com/davveo/lemonShop/app/consts"
	"github.com/davveo/lemonShop/app/service/base"
	"github.com/davveo/lemonShop/models/vo"
)

var _ SmsClient = (*smsClient)(nil)

type (
	SmsClient interface {
		i()
		// Valid 验证手机验证码
		Valid(scene, mobile, code string) (bool, error)
		// ValidMobile 获取验证通过的手机号
		ValidMobile(scene, mobile string) (string, error)
		// SendSmsMessage 发送(发送手机短信)消息
		SendSmsMessage(byName, mobile string, sceneType consts.SceneType) error
		// Send 真实发送手机短信
		Send(smsSendVO vo.SmsSendVo) error
	}
	smsClient struct {
		smsPlatformService    system.SmsPlatformService
		messageTemplateClient MessageTemplateClient
		settingService        base.SettingService
		eventMgr              event.EventManager
	}
)

func (s smsClient) Valid(scene, mobile, code string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s smsClient) ValidMobile(scene, mobile string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (s smsClient) SendSmsMessage(byName, mobile string, sceneType consts.SceneType) error {
	template, err := s.messageTemplateClient.GetByCode(consts.MOBILECODESEND)
	if err != nil {
		return err
	}
	settingCfg, err := s.settingService.Get(consts.SETTINGSITE)
	if err != nil {
		return err
	}
	var siteSetting vo.SiteSetting
	if err := json.Unmarshal([]byte(settingCfg), &siteSetting); err != nil {
		return err
	}
	dynamicCode := random.GetRandomCode(siteSetting.TestMode)
	content := string_util.StrSubstitutor(template.SmsContent, map[string]interface{}{
		"byName":   byName,
		"code":     dynamicCode,
		"siteName": siteSetting.SiteName,
	})
	smsSendVo := vo.SmsSendVo{
		Content: content,
		Mobile:  mobile,
	}
	//发送短信验证码
	s.eventMgr.Bind(consts.SEND_MESSAGE, &event_handler.SmsEventHandler{}).
		Trigger(consts.SEND_MESSAGE, smsSendVo)
	//缓存中记录验证码
	return cache.Cache.Set(consts.SMS_CODE+string(sceneType)+"_"+mobile,
		dynamicCode, time.Duration(conf.Conf.SmsCodeTimout))
}

func (s smsClient) Send(smsSendVO vo.SmsSendVo) error {
	platformVO, err := s.smsPlatformService.GetOpen()
	if err != nil {
		return err
	}
	if platformVO != nil {
		//todo
	}
	logger.GLogger.Info("已发送短信:短信内容为:" + smsSendVO.Content + "手机号:" + smsSendVO.Mobile)
	return nil
}

func (s smsClient) i() {}

func NewSmsClient() SmsClient {
	return &smsClient{
		settingService:        base.NewSettingService(),
		messageTemplateClient: NewMessageTemplateClient(),
		smsPlatformService:    system.NewSmsPlatformService(),
	}
}
