package goai

import (
	"fmt"
)

func (goai Client) ChatCompletion(messages []Message) (string, error) {
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
	return oaiResponse.Choices[0].Message.Content, goai.PostJson(oaiRequest, &oaiResponse, goai.Endpoint+"chat/completions")
}

func (goai Client) TextCompletion(prompt string) (ChatResponse, error) {
	oaiResponse := ChatResponse{}
	oaiRequest := &ChatRequest{
		Prompt:           prompt,
		User:             goai.User,
		TopP:             goai.TopP,
		Model:            goai.TextModel,
		MaxTokens:        goai.MaxTokens,
		Temperature:      goai.Temperature,
		PresencePenalty:  goai.PresencePenalty,
		FrequencyPenalty: goai.FrequencyPenalty,
	}
	if goai.Verbose {
		fmt.Println(oaiRequest)
	}
	return oaiResponse, goai.PostJson(oaiRequest, &oaiResponse, goai.Endpoint+"completions")
}
