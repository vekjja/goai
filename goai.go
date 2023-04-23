package goai

import (
	"hash/fnv"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	Endpoint         string
	API_KEY          string
	Verbose          bool
	User             string
	httpClient       *http.Client
	ImageSize        string
	TopP             float64
	ChatModel        string
	TextModel        string
	MaxTokens        int
	Temperature      float64
	FrequencyPenalty float64
	PresencePenalty  float64
}

func NewClient(apiKey string, verbose bool) *Client {

	// Create a unique user for OpenAI
	h := fnv.New32a()
	h.Write([]byte(apiKey))
	user := "gopenai-" + strconv.Itoa(int(h.Sum32()))

	return &Client{
		API_KEY:  apiKey,
		Verbose:  verbose,
		Endpoint: "https://api.openai.com/v1/",
		User:     user,
		httpClient: &http.Client{
			Timeout: time.Second * 60,
		},
		ImageSize:        "1024x1024",
		TopP:             0.9,
		ChatModel:        "gpt-3.5-turbo",
		TextModel:        "text-davinci-003",
		MaxTokens:        999,
		Temperature:      0.9,
		FrequencyPenalty: 0.03,
		PresencePenalty:  0.6,
	}
}
