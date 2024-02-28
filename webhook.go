package notificationsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c client) SendWebhook(ctx context.Context, accessToken string, payload WebhookNotifyRequest) error {
	bsBody, _ := json.Marshal(payload)

	headers := map[string]string{
		"Authorization": "Bearer " + accessToken,
		"Content-Type":  "application/json; charset=UTF-8",
	}

	res, err := c.doRequest(ctx, http.MethodPost, c.Config.URL+"/webhooks/execute", headers, bytes.NewReader(bsBody))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("http request error (status=%v)", res.Status)
	}

	return nil
}
