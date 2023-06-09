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
	httpClient *api.HttpClient

	requestBuilder    api.RequestBuilder
	createFormBuilder func(writer io.Writer) common.FormBuilder
}

// NewClient 创建 OpenAI 客户端
func NewClient(apiKey string) *Client {
	return &Client{
		httpClient:     api.NewHttpClient(apiKey),
		requestBuilder: api.NewHttpRequestBuilder(),
		createFormBuilder: func(body io.Writer) common.FormBuilder {
			return common.NewDefaultFormBuilder(body)
		},
	}
}

// CreateChatCompletion 创建聊天完成
func (c *Client) CreateChatCompletion(
	ctx context.Context,
	content string,
) (respContent string, err error) {
	request := chat.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []chat.ChatCompletionMessage{
			{
				Role:    chat.ChatMessageRoleSystem,
				Content: "you are a helpful chatbot",
			},
			{
				Role:    chat.ChatMessageRoleUser,
				Content: content,
			},
		},
	}

	router, err := api.NewRouter(c.httpClient.BaseUrl())
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

	var response chat.ChatCompletionResponse
	err = c.httpClient.SendRequest(req, &response)
	respContent = response.Choices[0].Message.Content
	return
}
