package system

import (
	"github.com/davveo/lemonShop/mgr"
	"github.com/davveo/lemonShop/models"
	"github.com/davveo/lemonShop/pkg/db"
)

var _ MessageTemplateClient = (*messageTemplateClient)(nil)

type (
	MessageTemplateClient interface {
		i()
		GetByCode(code string) (*models.EsMessageTemplate, error)
	}

	messageTemplateClient struct {
		messageTplMgr *mgr.MessageTemplateMgr
	}
)

func (m messageTemplateClient) GetByCode(code string) (*models.EsMessageTemplate, error) {
	return m.messageTplMgr.GetByOption(m.messageTplMgr.WithTplCode(code))
}

func (m messageTemplateClient) i() {}

func NewMessageTemplateClient() MessageTemplateClient {
	return &messageTemplateClient{
		messageTplMgr: mgr.NewMessageTemplateMgr(db.GRpo),
	}
}
