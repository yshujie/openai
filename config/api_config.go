package config

const openaiAPIURLv1 = "https://api.openai.com/v1"

type APIType string

const APITypeOpenAI APIType = "OPEN_AI"

type APIConfig struct {
	BaseUrl string
	APIType APIType
	APIKey  string
}

// DefaultAPIConfig 默认的API配置
func DefaultAPIConfig(APIKey string) APIConfig {
	return APIConfig{
		BaseUrl: openaiAPIURLv1,
		APIType: APITypeOpenAI,
		APIKey:  APIKey,
	}
}

func (APIConfig) string() string {
	return "<OpenAI API Config>"
}
