package notificationsdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	PaymentLinkInvoiceCreatedTemplate = "PAYMENT_LINK_INVOICE_CREATED"
)

func (c client) SendWhatsapp(ctx context.Context, accessToken string, payload WhatsappNotifyRequest) error {
	bsBody, _ := json.Marshal(payload)

	headers := map[string]string{
		"Authorization": "Bearer " + accessToken,
		"Content-Type":  "application/json; charset=UTF-8",
	}

	res, err := c.doRequest(ctx, http.MethodPost, c.Config.URL+"/send-whatsapp", headers, bytes.NewReader(bsBody))
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusAccepted {
		return fmt.Errorf("http request error (status=%v)", res.Status)
	}

	return nil
}
