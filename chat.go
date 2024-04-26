package goai

import ()

func (goai Client) ChatCompletion(messages []Message) (ChatCompletionResponse, error) {
	oaiResponse := ChatCompletionResponse{}
	oaiRequest := ChatCompletionRequest{
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
	return oaiResponse, goai.PostJson(oaiRequest, &oaiResponse, goai.Endpoint+"chat/completions")
}
