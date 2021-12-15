package notificationsdk_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type mockResponse struct {
	Url        string
	StatusCode int
	Body       interface{}
}

type requestClientMock struct {
	Responses []mockResponse
}

func (r requestClientMock) Do(req *http.Request) (*http.Response, error) {
	var mockRes *mockResponse
	for _, res := range r.Responses {
		if strings.Contains(req.URL.String(), res.Url) {
			mockRes = &res
			break
		}
	}

	httpRes := &http.Response{}
	if mockRes == nil {
		return httpRes, fmt.Errorf(fmt.Sprintf("Not found mock response for url %v", req.RequestURI))
	}

	jsonResp, err := json.Marshal(mockRes.Body)
	if err != nil {
		return nil, err
	}

	httpRes.StatusCode = mockRes.StatusCode
	httpRes.Body = io.NopCloser(bytes.NewReader(jsonResp))

	return httpRes, nil
}
