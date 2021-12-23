package notificationsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	PaymentLinkInvoiceCreatedTemplate = "PAYMENT_LINK_INVOICE_CREATED"
)

func (c client) SendWhatsapp(ctx context.Context, accessToken string, payload WhatsappNotifyRequest) error {
	bsBody, _ := json.Marshal(payload)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.Config.URL+"/send-whatsapp", io.NopCloser(bytes.NewReader(bsBody)))
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

	if res.StatusCode != http.StatusAccepted {
		return fmt.Errorf("http request error (status=%v)", res.Status)
	}

	return nil
}
