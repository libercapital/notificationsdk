package notificationsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c client) SendEmail(ctx context.Context, accessToken string, payload EmailRequest) error {
	bsBody, _ := json.Marshal(payload)

	headers := map[string]string{
		"Authorization": "Bearer " + accessToken,
		"Content-Type":  "application/json; charset=UTF-8",
	}

	res, err := HttpClient.DoRequest(ctx, http.MethodPost, c.Config.URL+"/send-email", headers, bytes.NewReader(bsBody))

	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("http request error (status=%v)", res.Status)
	}

	return nil
}
