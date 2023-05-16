package api

import (
	"bytes"
	"context"
	"github.com/yshujie/openai/common"
	"net/http"
)

// RequestBuilder 请求构造器
type RequestBuilder interface {
	build(ctx context.Context, method, url string, request any) (*http.Request, error)
}

// HttpRequestBuilder Http 请求构造器
type HttpRequestBuilder struct {
	marshaller common.Marshaller
}

// NewHttpRequestBuilder 创建 Http 请求构造器
func NewHttpRequestBuilder() *HttpRequestBuilder {
	return &HttpRequestBuilder{marshaller: &common.JsonMarshaller{}}
}

// Build 构造 http 请求
func (b *HttpRequestBuilder) Build(ctx context.Context, method, url string, request any) (*http.Request, error) {
	if request == nil {
		return http.NewRequestWithContext(ctx, method, url, nil)
	}

	var reqBytes []byte
	reqBytes, err := b.marshaller.Marshal(request)
	if err != nil {
		return nil, err
	}

	return http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(reqBytes))
}
