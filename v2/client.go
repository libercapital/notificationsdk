package notificationsdk

import (
	"context"
	"io"
	"net/http"

	"github.com/kataras/compress"
	liberlogger "github.com/libercapital/liber-logger-go"
	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
)

type client struct {
	Config     Config
	httpClient *http.Client
}

type Client interface {
	SendWebhook(ctx context.Context, payload WebhookNotifyRequest) error
	SendWhatsapp(ctx context.Context, payload WhatsappNotifyRequest) error
	SendEmail(ctx context.Context, payload EmailRequest) error
}

func NewClient(cli http.Client, config Config) Client {
	opt := []httptrace.RoundTripperOption{
		httptrace.RTWithServiceName("notification-sdk"),
		httptrace.RTWithResourceNamer(func(req *http.Request) string {
			return req.Method + " " + req.URL.String()
		}),
	}

	httpClient := httptrace.WrapClient(&cli, opt...)

	return client{
		Config:     config,
		httpClient: httpClient,
	}
}

func (c client) doRequest(ctx context.Context, httpMethod, url string, headers map[string]string, body io.Reader, redactHeaders ...string) (*http.Response, error) {
	if c.httpClient.Transport == nil {
		c.httpClient.Transport = liberlogger.HttpClient{
			Proxied:      http.DefaultTransport,
			RedactedKeys: redactHeaders,
		}
	}

	request, err := http.NewRequestWithContext(ctx, httpMethod, url, body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.Header.Get("Content-Encoding") == "deflate" {
		reader, _ := compress.NewReader(response.Body, compress.DEFLATE)
		defer reader.Close()

		response.Body = reader
	}

	return response, nil
}
