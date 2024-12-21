package goai

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
