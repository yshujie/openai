package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/yshujie/openai/common"
	"github.com/yshujie/openai/config"
	"io"
	"net/http"
)

var ErrNoValidAPIKey = errors.New("APIKey is not found, please get a valid APIKey")

// HttpClient http 客户端
type HttpClient struct {
	apiConfig config.APIConfig
	http      *http.Client
}

// NewHttpClient 新增 http 客户端
func NewHttpClient(apiKey string) *HttpClient {
	return &HttpClient{
		apiConfig: config.DefaultAPIConfig(apiKey),
		http:      &http.Client{},
	}
}

// BaseUrl 获取 baseUrl
func (c *HttpClient) BaseUrl() string {
	return c.apiConfig.BaseUrl
}

// SendRequest 发送请求
func (c *HttpClient) SendRequest(request *http.Request, v any) error {
	err := c.initRequestHeader(request)
	if err != nil {
		return err
	}

	// 发送 http 请求
	response, err := c.http.Do(request)
	if err != nil {
		return err
	}

	//
	defer response.Body.Close()

	if !c.isSuccess(response) {
		return c.handleErrorResponse(response)
	}

	return c.decodeResponse(response.Body, v)
}

// initRequestHeader
func (c *HttpClient) initRequestHeader(request *http.Request) error {
	if len(c.apiConfig.APIKey) <= 0 {
		return ErrNoValidAPIKey
	}

	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiConfig.APIKey))
	request.Header.Set("Accept", "application/json; charset=utf-8")
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	return nil
}

// isSuccess http 请求是否成功
func (c *HttpClient) isSuccess(response *http.Response) bool {
	return http.StatusOK <= response.StatusCode && response.StatusCode < http.StatusBadRequest
}

// handleErrorResponse http 响应异常处理
func (c *HttpClient) handleErrorResponse(response *http.Response) error {
	var errResp common.ErrorResponse
	err := json.NewDecoder(response.Body).Decode(&errResp)
	if err != nil {
		return &common.RequestError{
			HTTPStatusCode: response.StatusCode,
			Err:            err,
		}
	}
	if errResp.Error == nil {
		return &common.RequestError{
			HTTPStatusCode: response.StatusCode,
			Err:            errResp.Error,
		}
	}

	errResp.Error.HTTPStatusCode = response.StatusCode
	return errResp.Error
}

// decodeResponse http 响应解码
func (c *HttpClient) decodeResponse(body io.Reader, v any) error {
	if v == nil {
		return nil
	}

	// 若结果为字符串，则将响应解码为字符串
	if result, ok := v.(*string); ok {
		return c.decodeString(body, result)
	}

	return json.NewDecoder(body).Decode(v)
}

// decodeString 字符串解码
func (c *HttpClient) decodeString(body io.Reader, output *string) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	*output = string(b)
	return nil
}
