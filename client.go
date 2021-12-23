package notificationsdk

import (
	"context"
)

type client struct {
	Config Config
}

type Client interface {
	SendWebhook(ctx context.Context, accessToken string, payload WebhookNotifyRequest) error
	SendWhatsapp(ctx context.Context, accessToken string, payload WhatsappNotifyRequest) error
}

func NewClient(config Config) Client {
	return client{
		Config: config,
	}
}
