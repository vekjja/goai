package goai

func (goai Client) ChatCompletion(messages []Message) (ChatCompletionResponse, error) {
	res := ChatCompletionResponse{}
	req := ChatCompletionRequest{
		N:                1,
		Messages:         messages,
		User:             goai.User,
		TopP:             goai.TopP,
		Model:            goai.ChatModel,
		MaxTokens:        goai.MaxTokens,
		Temperature:      goai.Temperature,
		PresencePenalty:  goai.PresencePenalty,
		FrequencyPenalty: goai.FrequencyPenalty,
	}
	_, err := goai.PostJson(req, &res, goai.Endpoint+"chat/completions")
	return res, err
}
