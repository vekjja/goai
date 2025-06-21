package goai

// ChatCompletionRequest represents the request structure for OpenAI's chat completion API.
type ChatCompletionRequest struct {
	Model            string    `json:"model"`
	Messages         []Message `json:"messages"`
	MaxTokens        *int      `json:"max_tokens,omitempty"`
	Temperature      *float64  `json:"temperature,omitempty"`
	TopP             *float64  `json:"top_p,omitempty"`
	N                *int      `json:"n,omitempty"`
	Stream           bool      `json:"stream"`
	PresencePenalty  *float64  `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float64  `json:"frequency_penalty,omitempty"`
	Stop             []string  `json:"stop,omitempty"`
	User             string    `json:"user,omitempty"`
}

// ChatCompletionResponse represents the response structure from OpenAI's chat completion API.
type ChatCompletionResponse struct {
	ID                string `json:"id"`
	Object            string `json:"object"`
	Created           int    `json:"created"`
	Model             string `json:"model"`
	SystemFingerprint string `json:"system_fingerprint"`
	Choices           []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs     interface{} `json:"logprobs"`
		FinishReason string      `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens            int `json:"prompt_tokens"`
		CompletionTokens        int `json:"completion_tokens"`
		TotalTokens             int `json:"total_tokens"`
		CompletionTokensDetails struct {
			ReasoningTokens          int `json:"reasoning_tokens"`
			AcceptedPredictionTokens int `json:"accepted_prediction_tokens"`
			RejectedPredictionTokens int `json:"rejected_prediction_tokens"`
		} `json:"completion_tokens_details"`
	} `json:"usage"`
}

// Message represents a single message in the conversation history.
type Message struct {
	Role    string `json:"role"`    // "system", "assistant", or "user"
	Content string `json:"content"` // The text content of the message
}

func (goai Client) ChatCompletion(messages []Message) (ChatCompletionResponse, error) {
	res := ChatCompletionResponse{}

	req := ChatCompletionRequest{
		N:                IntPtr(1), // Convert to pointer using helper function
		Messages:         messages,
		User:             goai.User,
		TopP:             Float64Ptr(goai.TopP), // Convert to pointer
		Model:            goai.ChatModel,
		MaxTokens:        IntPtr(goai.MaxTokens),            // Convert to pointer
		Temperature:      Float64Ptr(goai.Temperature),      // Convert to pointer
		PresencePenalty:  Float64Ptr(goai.PresencePenalty),  // Convert to pointer
		FrequencyPenalty: Float64Ptr(goai.FrequencyPenalty), // Convert to pointer
	}

	_, err := goai.PostJson(req, &res, goai.Endpoint+"chat/completions")
	return res, err
}
