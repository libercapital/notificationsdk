package notificationsdk_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/bava-http.git/mocks"
	"gitlab.com/bavatech/architecture/software/libs/go-modules/notificationsdk.git"
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
				notificationsdk.HttpClient = &mocks.HttpClient{}
				notificationsdk.HttpClient.(*mocks.HttpClient).
					On("DoRequest", mock.Anything, http.MethodPost, f.Config.URL+"/send-whatsapp", mock.Anything, mock.Anything).
					Return(&http.Response{StatusCode: http.StatusAccepted, Body: io.NopCloser(bytes.NewReader([]byte{}))}, nil).
					Once()
			},
		},
		{
			name: "error doing request",
			fields: fields{
				Config: notificationsdk.Config{
					URL: "http://localhost:1000",
				},
			},
			wantErr: true,
			mockBehavior: func(f fields) {
				notificationsdk.HttpClient = &mocks.HttpClient{}
				notificationsdk.HttpClient.(*mocks.HttpClient).
					On("DoRequest", mock.Anything, http.MethodPost, f.Config.URL+"/send-whatsapp", mock.Anything, mock.Anything).
					Return(nil, errors.New("test error")).
					Once()
			},
		},
		{
			name: "error status not accepted",
			fields: fields{
				Config: notificationsdk.Config{
					URL: "http://localhost:1000",
				},
			},
			wantErr: true,
			mockBehavior: func(f fields) {
				notificationsdk.HttpClient = &mocks.HttpClient{}
				notificationsdk.HttpClient.(*mocks.HttpClient).
					On("DoRequest", mock.Anything, http.MethodPost, f.Config.URL+"/send-whatsapp", mock.Anything, mock.Anything).
					Return(&http.Response{StatusCode: http.StatusBadRequest, Body: io.NopCloser(bytes.NewReader([]byte{}))}, nil).
					Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := notificationsdk.NewClient(http.Client{}, tt.fields.Config)
			tt.mockBehavior(tt.fields)
			err := c.SendWhatsapp(context.TODO(), "ACCESS_TOKEN", notificationsdk.WhatsappNotifyRequest{})
			if (err != nil) != tt.wantErr {
				t.Errorf("c.SendWhatsapp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
