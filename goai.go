package goai

import (
	"errors"
	"hash/fnv"
	"net/http"
	"strconv"
	"time"
)

type GoAi struct {
	Endpoint         string
	API_KEY          string
	Verbose          bool
	User             string
	httpClient       *http.Client
	TopP             float64
	ChatModel        string
	TextModel        string
	MaxTokens        int
	Temperature      float64
	FrequencyPenalty float64
	PresencePenalty  float64
}

func New(apiKey string, verbose bool) (*GoAi, error) {

	if apiKey == "" {
		return nil, errors.New("API Key is required")
	}

	// Create a unique user for OpenAI
	h := fnv.New32a()
	h.Write([]byte(apiKey))
	user := "gopenai-" + strconv.Itoa(int(h.Sum32()))

	return &GoAi{
		API_KEY:  apiKey,
		Verbose:  verbose,
		Endpoint: "https://api.openai.com/v1/",
		User:     user,
		httpClient: &http.Client{
			Timeout: time.Second * 60,
		},
		TopP:             0.9,
		ChatModel:        "gpt-3.5-turbo",
		TextModel:        "text-davinci-003",
		MaxTokens:        999,
		Temperature:      0.9,
		FrequencyPenalty: 0.03,
		PresencePenalty:  0.6,
	}, nil
}
