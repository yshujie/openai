package api

import (
	"bytes"
	"context"
	"net/http"
)

type RequestBuilder interface {
	build(ctx context.Context, method, url string, request any) (*http.Request, error)
}

type HttpRequestBuilder struct{}

func (b *HttpRequestBuilder) build(ctx context.Context, method, url string, request any) (*http.Request, error) {
	if request == nil {
		return http.NewRequestWithContext(ctx, method, url, nil)
	}

	var reqBytes []byte
	return http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(reqBytes))
}
