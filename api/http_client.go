package api

import (
	"github.com/yshujie/openai/config"
	"net/http"
)

// HttpClient http 客户端
type HttpClient struct {
	apiConfig config.APIConfig
	http      *http.Client
}

// NewHttpClient 新增 http 客户端
func NewHttpClient(apiKey string) *HttpClient {
	apiConfig := config.DefaultConfig(apiKey)
	return &HttpClient{
		apiConfig: apiConfig,
		http:      &http.Client{},
	}
}

// SendRequest 发送请求
func (c *HttpClient) SendRequest(request *http.Request, v any) error {
	_, err := c.http.Do(request)
	return err
}
