package goai

type ChatRequest struct {
	Prompt           string   `json:"prompt"`
	MaxTokens        int      `json:"max_tokens"`
	Model            string   `json:"model"`
	Temperature      float64  `json:"temperature"`
	TopP             float64  `json:"top_p"`
	FrequencyPenalty float64  `json:"frequency_penalty"`
	PresencePenalty  float64  `json:"presence_penalty"`
	Stop             []string `json:"stop"`
	User             string   `json:"user"`
}

type ChatResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Text         string      `json:"text"`
		Index        int         `json:"index"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type ImageRequest struct {
	Prompt         string `json:"prompt"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_format"`
	User           string `json:"user"`
}

type ImageEditRequest struct {
	Prompt         string `json:"prompt"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_format"`
	User           string `json:"user"`
	Image          string `json:"image"`
	Mask           string `json:"mask"`
}

type ImageResponse struct {
	Created int64 `json:"created"`
	Data    []struct {
		URL string `json:"url"`
	} `json:"data"`
}

type ErrorResponse struct {
	Error struct {
		Message string `json:"message"`
		Type    string `json:"type"`
		Param   string `json:"param"`
		Code    string `json:"code"`
	} `json:"error"`
}

type ChatCompletionResponse struct {
	ID      string         `json:"id"`
	Object  string         `json:"object"`
	Created int            `json:"created"`
	Choices []Choice       `json:"choices"`
	Usage   map[string]int `json:"usage"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletionRequest struct {
	Model            string    `json:"model"`
	Messages         []Message `json:"messages"`
	MaxTokens        int       `json:"max_tokens"`
	Temperature      float64   `json:"temperature"`
	TopP             float64   `json:"top_p"`
	N                int       `json:"n"`
	Stream           bool      `json:"stream"`
	PresencePenalty  float64   `json:"presence_penalty"`
	FrequencyPenalty float64   `json:"frequency_penalty"`
	Stop             []string  `json:"stop"`
	User             string    `json:"user"`
}
