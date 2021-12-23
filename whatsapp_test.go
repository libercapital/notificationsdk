package notificationsdk_test

import (
	"context"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/notificationsdk.git"
	"net/http"
	"testing"
)

func Test_client_SendWhatsapp(t *testing.T) {
	type fields struct {
		Config notificationsdk.Config
	}
	tests := []struct {
		name         string
		fields       fields
		wantErr      bool
		mockBehavior func(fields)
	}{
		{
			name: "should send whatsapp message and return nil",
			fields: fields{
				Config: notificationsdk.Config{
					URL: "http://localhost:1000",
				},
			},
			wantErr: false,
			mockBehavior: func(f fields) {
				notificationsdk.HttpClient = &requestClientMock{
					Responses: []mockResponse{
						{
							Url:        f.Config.URL + "/send-whatsapp",
							StatusCode: http.StatusAccepted,
							Body:       nil,
						},
					},
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := notificationsdk.NewClient(tt.fields.Config)
			tt.mockBehavior(tt.fields)
			err := c.SendWhatsapp(context.TODO(), "ACCESS_TOKEN", notificationsdk.WhatsappNotifyRequest{})
			if (err != nil) != tt.wantErr {
				t.Errorf("client.SendWhatsapp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
