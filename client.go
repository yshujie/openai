package openai

import (
	"context"
	"github.com/yshujie/openai/api"
	"github.com/yshujie/openai/chat"
	"github.com/yshujie/openai/common"
	"io"
	"net/http"
)

// Client OpenAI 客户端
type Client struct {
	httpClient api.HttpClient

	requestBuilder    api.RequestBuilder
	createFormBuilder func(writer io.Writer) common.FormBuilder
}

// CreateChatCompletion 创建聊天完成
func (c *Client) CreateChatCompletion(
	ctx context.Context,
	request chat.ChatCompletionRequest,
) (response chat.ChatCompletionResponse, err error) {
	if request.Stream {
		err = chat.ErrChatCompletionStreamNotSupported
		return
	}

	router, err := api.NewRouter()
	if err != nil {
		return
	}
	url, err := router.Route("chat-completions")
	if err != nil {
		return
	}

	req, err := c.requestBuilder.Build(
		ctx,
		http.MethodPost,
		url,
		request,
	)
	if err != nil {
		return
	}

	err = c.httpClient.SendRequest(req, &response)
	return
}
