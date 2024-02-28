package notificationsdk_test

// import (
// 	"context"
// 	"net/http"
// 	"testing"
//
// 	"github.com/libercapital/notificationsdk/v2"
// )
//
// func Test_client_SendEmail(t *testing.T) {
// 	type fields struct {
// 		Config notificationsdk.Config
// 	}
// 	type args struct {
// 		ctx         context.Context
// 		accessToken string
// 		payload     notificationsdk.EmailRequest
// 	}
// 	tests := []struct {
// 		name         string
// 		fields       fields
// 		args         args
// 		wantErr      bool
// 		mockBehavior func(fields, args)
// 	}{
// 		{
// 			name: "success",
// 			args: args{ctx: context.TODO(), accessToken: "ACCESS_TOKEN", payload: notificationsdk.EmailRequest{}},
// 			fields: fields{
// 				Config: notificationsdk.Config{
// 					URL: "http://localhost:1000",
// 				},
// 			},
// 			mockBehavior: func(f fields, a args) {
// 				// notificationsdk.HttpClient = &mocks.HttpClient{}
// 				// notificationsdk.HttpClient.(*mocks.HttpClient).
// 				// 	On("DoRequest", mock.Anything, http.MethodPost, f.Config.URL+"/send-email", mock.Anything, mock.Anything).
// 				// 	Return(&http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewReader([]byte{}))}, nil).
// 				// 	Once()
// 			},
// 		},
// 		{
// 			name: "error doing request",
// 			args: args{ctx: context.TODO(), accessToken: "ACCESS_TOKEN", payload: notificationsdk.EmailRequest{}},
// 			fields: fields{
// 				Config: notificationsdk.Config{
// 					URL: "http://localhost:1000",
// 				},
// 			},
// 			mockBehavior: func(f fields, a args) {
// 				// notificationsdk.HttpClient = &mocks.HttpClient{}
// 				// notificationsdk.HttpClient.(*mocks.HttpClient).
// 				// 	On("DoRequest", mock.Anything, http.MethodPost, f.Config.URL+"/send-email", mock.Anything, mock.Anything).
// 				// 	Return(nil, errors.New("test error")).
// 				// 	Once()
// 			},
// 			wantErr: true,
// 		},
// 		{
// 			name: "status code not accepted",
// 			args: args{ctx: context.TODO(), accessToken: "ACCESS_TOKEN", payload: notificationsdk.EmailRequest{}},
// 			fields: fields{
// 				Config: notificationsdk.Config{
// 					URL: "http://localhost:1000",
// 				},
// 			},
// 			mockBehavior: func(f fields, a args) {
// 				// notificationsdk.HttpClient = &mocks.HttpClient{}
// 				// notificationsdk.HttpClient.(*mocks.HttpClient).
// 				// 	On("DoRequest", mock.Anything, http.MethodPost, f.Config.URL+"/send-email", mock.Anything, mock.Anything).
// 				// 	Return(&http.Response{StatusCode: http.StatusBadRequest, Body: io.NopCloser(bytes.NewReader([]byte{}))}, nil).
// 				// 	Once()
// 			},
// 			wantErr: true,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := notificationsdk.NewClient(http.Client{}, tt.fields.Config)
// 			tt.mockBehavior(tt.fields, tt.args)
// 			if err := c.SendEmail(tt.args.ctx, tt.args.payload); (err != nil) != tt.wantErr {
// 				t.Errorf("client.SendEmail() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
