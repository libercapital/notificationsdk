package notificationsdk

import (
	"context"
	"net/http"

	"gitlab.com/bavatech/architecture/software/libs/go-modules/bava-http.git/httpclient"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
)

type client struct {
	Config Config
}

type Client interface {
	SendWebhook(ctx context.Context, payload WebhookNotifyRequest) error
	SendWhatsapp(ctx context.Context, payload WhatsappNotifyRequest) error
	SendEmail(ctx context.Context, payload EmailRequest) error
}

var HttpClient httpclient.HttpClient

func NewClient(cli http.Client, config Config) Client {
	opt := []httptrace.RoundTripperOption{
		httptrace.RTWithServiceName("notification-sdk"),
		httptrace.RTWithResourceNamer(func(req *http.Request) string {
			return req.Method + " " + req.URL.String()
		}),
	}

	HttpClient = httpclient.New(httptrace.WrapClient(&cli, opt...))

	return client{
		Config: config,
	}
}
