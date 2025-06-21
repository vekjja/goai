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
	Verbose          int
	User             string
	HTTPClient       *http.Client
	ImageSize        string
	TopP             float64
	ChatModel        string
	ImageModel       string
	TTSModel         string
	Voice            string
	Speed            float64
	ResponseFormat   string
	MaxTokens        int
	Temperature      float64
	FrequencyPenalty float64
	PresencePenalty  float64
}

// Create a unique user for OpenAI to track
func HashAPIKey(apiKey string) string {
	h := fnv.New64a()
	h.Write([]byte(apiKey))
	return strconv.FormatUint(h.Sum64(), 10)
}

func DefaultClient(apiKey string, verbose int) *Client {
	return &Client{
		API_KEY:  apiKey,
		Verbose:  verbose,
		Endpoint: "https://api.openai.com/v1/",
		User:     HashAPIKey(apiKey),
		HTTPClient: &http.Client{
			Timeout: time.Second * 60,
		},
		ImageSize:        "1024x1024",
		TopP:             0.9,
		ChatModel:        "gpt-4",
		TTSModel:         "tts-1",
		Voice:            "onyx",
		Speed:            1,
		ResponseFormat:   "mp3",
		MaxTokens:        999,
		Temperature:      0.9,
		FrequencyPenalty: 0.03,
		PresencePenalty:  0.6,
	}
}
