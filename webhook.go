package notificationsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
)

type requestClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var HttpClient requestClient

func init() {
	HttpClient = httptrace.WrapClient(
		&http.Client{
			Transport: &logTransport{http.DefaultTransport},
			Timeout:   60 * time.Second,
		},
		httptrace.RTWithResourceNamer(func(req *http.Request) string {
			return "http.notificationsdk"
		}),
	)
}

func (c client) SendWebhook(ctx context.Context, accessToken string, payload WebhookNotifyRequest) error {
	bsBody, _ := json.Marshal(payload)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.Config.URL+"/webhook/execute", io.NopCloser(bytes.NewReader(bsBody)))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("http request error (status=%v)", res.Status)
	}

	return nil
}
