package notificationsdk

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type logTransport struct {
	Transport http.RoundTripper
}

const REDACTED_MASK = "REDACTED"

var redactedKeys = []string{
	"access_token",
	"client_secret",
	"Authorization",
}

func redactBody(data interface{}) interface{} {
	switch value := data.(type) {
	case []interface{}:
		for index, item := range value {
			value[index] = redactBody(item)
		}
	case map[string]interface{}:
		for _, key := range redactedKeys {
			if value[key] != nil {
				value[key] = REDACTED_MASK
			}
		}
	}
	return data
}

func redactHeader(header map[string][]string) map[string][]string {
	copyHeader := make(map[string][]string)
	for key, value := range header {
		copyHeader[key] = append(copyHeader[key], value...)
	}
	for _, key := range redactedKeys {
		if copyHeader[key] != nil {
			maskField := append(make([]string, 0), REDACTED_MASK)
			copyHeader[key] = maskField
		}
	}
	return copyHeader
}

func extractBody(bodyRef *io.ReadCloser) interface{} {
	var requestBody interface{}
	body := *bodyRef
	buf, err := ioutil.ReadAll(body)
	if err != nil {
		log.Error().Msgf("Error reading request body: %v", err.Error())
	}
	defer func() {
		body.Close()
		*bodyRef = ioutil.NopCloser(bytes.NewBuffer(buf))
	}()
	json.Unmarshal(buf, &requestBody)
	return requestBody
}

func logRequest(req *http.Request, reqId string) {
	body, _ := req.GetBody()
	log.Info().
		Interface("id", reqId).
		Interface("body", redactBody(extractBody(&body))).
		Interface("header", redactHeader(req.Header)).
		Msgf("HTTP Request %s %s", req.Method, req.URL)
}

func logResponse(res *http.Response, reqId string) {
	log.Info().
		Interface("id", reqId).
		Interface("body", redactBody(extractBody(&res.Body))).
		Interface("header", redactHeader(res.Header)).
		Msgf("HTTP Response %d %s", res.StatusCode, res.Request.URL)
}

func logFail(err error, reqId string) {
	log.Error().
		Err(err).
		Interface("id", reqId).
		Msg("Failed to execute http request")
}

func (lt logTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	reqId := uuid.NewString()
	logRequest(req, reqId)

	res, err := lt.Transport.RoundTrip(req)
	if err != nil {
		logFail(err, reqId)
	} else {
		logResponse(res, reqId)
	}

	return res, err
}
