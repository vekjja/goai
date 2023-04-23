package gopenai

import (
	"fmt"

	"github.com/spf13/viper"
)

func (goai GOpenAI) ChatCompletion(messages []OPENAI_Message) (string, error) {
	oaiResponse := OPENAI_ChatCompletionResponse{}
	oaiRequest := OPENAI_ChatCompletionRequest{
		N:                1,
		Messages:         messages,
		User:             goai.User,
		TopP:             viper.GetFloat64("openAI_chat_topP"),
		Model:            viper.GetString("openAI_chat_model"),
		MaxTokens:        viper.GetInt("openAI_chat_maxTokens"),
		Temperature:      viper.GetFloat64("openAI_chat_temperature"),
		FrequencyPenalty: viper.GetFloat64("openAI_chat_frequencyPenalty"),
		PresencePenalty:  viper.GetFloat64("openAI_chat_presencePenalty"),
	}
	return oaiResponse.Choices[0].Message.Content, goai.PostJson(oaiRequest, &oaiResponse, viper.GetString("openAI_endpoint")+"chat/completions")
}

func (goai GOpenAI) TextCompletion(prompt string) (OPENAI_ChatResponse, error) {
	oaiResponse := OPENAI_ChatResponse{}
	oaiRequest := &OPENAI_ChatRequest{
		Prompt:           prompt,
		User:             goai.User,
		Model:            viper.GetString("openAI_text_model"),
		MaxTokens:        viper.GetInt("openAI_text_maxTokens"),
		Temperature:      viper.GetFloat64("openAI_text_temperature"),
		TopP:             viper.GetFloat64("openAI_text_topP"),
		FrequencyPenalty: viper.GetFloat64("openAI_text_frequencyPenalty"),
		PresencePenalty:  viper.GetFloat64("openAI_text_presencePenalty"),
	}
	if goai.Verbose {
		fmt.Println(oaiRequest)
	}
	return oaiResponse, goai.PostJson(oaiRequest, &oaiResponse, viper.GetString("openAI_endpoint")+"completions")
}
