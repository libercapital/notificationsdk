package notificationsdk

import (
	"context"
	"net/http"
	"time"

	"gitlab.com/bavatech/architecture/software/libs/go-modules/bava-http.git/httpclient"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
)

type client struct {
	Config Config
}

type Client interface {
	SendWebhook(ctx context.Context, accessToken string, payload WebhookNotifyRequest) error
	SendWhatsapp(ctx context.Context, accessToken string, payload WhatsappNotifyRequest) error
	SendEmail(ctx context.Context, accessToken string, payload EmailRequest) error
}

var HttpClient httpclient.HttpClient

func init() {
	client := &http.Client{Timeout: 60 * time.Second}

	opt := []httptrace.RoundTripperOption{
		httptrace.RTWithServiceName("notification-sdk"),
		httptrace.RTWithResourceNamer(func(req *http.Request) string {
			return req.Method + " " + req.URL.String()
		}),
	}

	HttpClient = httpclient.New(httptrace.WrapClient(client, opt...))
}

func NewClient(config Config) Client {
	return client{
		Config: config,
	}
}
