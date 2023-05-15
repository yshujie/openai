package config

const openaiAPIURLv1 = "https://api.openai.com/v1"

type APIType string

const APITypeOpenAI APIType = "OPEN_AI"

type ClientConfig struct {
	BaseUrl string
	APIType APIType
	APIKey  string
}

// DefaultConfig 默认的客户端配置
func DefaultConfig(APIKey string) ClientConfig {
	return ClientConfig{
		BaseUrl: openaiAPIURLv1,
		APIType: APITypeOpenAI,
		APIKey:  APIKey,
	}
}

func (ClientConfig) string() string {
	return "<OpenAI API ClientConfig>"
}
