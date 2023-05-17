package api

import "errors"

var ErrInvalidUrlCode = errors.New("can not find url by this code, please check your url code")

type Router struct {
	baseUrl string
	urlMap  map[string]string
}

func NewRouter(baseUrl string) (*Router, error) {
	return &Router{
		baseUrl: baseUrl,
		urlMap:  loadUrlMap(),
	}, nil
}

var OpenAIUrlMaps = map[string]string{
	"chat-completions": "/chat/completions",
}

func loadUrlMap() map[string]string {
	return OpenAIUrlMaps
}

func (r *Router) Route(code string) (url string, err error) {
	url = r.urlMap[code]
	if len(url) <= 0 {
		return "", ErrInvalidUrlCode
	}

	return r.baseUrl + url, nil
}
