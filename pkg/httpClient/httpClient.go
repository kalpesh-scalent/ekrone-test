package httpClient

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/kalpesh-scalent/ekrone-test/pkg/logger"
)

type HttpClient struct {
	apiBaseURL string
	logger     logger.Logger
	client     *http.Client
}

func GetHttpClient(baseURL string, l *logger.Logger) *HttpClient {

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	return &HttpClient{
		apiBaseURL: baseURL,
		client:     client,
		logger:     *l,
	}
}

func (h HttpClient) ReadResponse(ctx context.Context, resp *http.Response, v interface{}) error {

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		h.logger.WithContext(ctx).Error(err)
		return err
	}

	err = json.Unmarshal(bodyBytes, v)
	if err != nil {
		h.logger.WithContext(ctx).Error(err)
		return err
	}

	return nil
}
