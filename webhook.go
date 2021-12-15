package notificationsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

type OAuthResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   uint16 `json:"expires_in"`
}

func (c client) SendWebhook(ctx context.Context, accessToken string, payload WebhookNotifyRequest) error {
	bsBody, _ := json.Marshal(payload)
	reqURL, _ := url.Parse(c.Config.URL + "/webhook/execute")

	response, err := HttpClient.Do(&http.Request{
		URL:    reqURL,
		Method: http.MethodPost,
		Header: http.Header{
			"Authorization": []string{"Bearer " + accessToken},
			"Content-Type":  []string{"application/json; charset=UTF-8"},
		},
		Body: io.NopCloser(bytes.NewReader(bsBody)),
	})

	if err != nil {
		return err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("http request error (status=%v)", response.Status)
	}

	return nil
}
