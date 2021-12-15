package notificationsdk_test

import (
	"context"
	"testing"

	"gitlab.com/bavatech/architecture/software/libs/go-modules/notificationsdk.git"
)

func Test_client_SendWebhook(t *testing.T) {
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
			name: "should send webhook message and return nil",
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
							Url:        f.Config.URL + "/webhook/execute",
							StatusCode: 200,
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
			err := c.SendWebhook(context.TODO(), "ACCESS_TOKEN", notificationsdk.WebhookNotifyRequest{})
			if (err != nil) != tt.wantErr {
				t.Errorf("client.createAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
