package gopenai

import (
	"errors"
	"hash/fnv"
	"net/http"
	"strconv"
	"time"
)

type GOpenAI struct {
	API_KEY    string
	Verbose    bool
	User       string
	httpClient *http.Client
}

func NewGOpenAI(apiKey string, verbose bool) (*GOpenAI, error) {

	if apiKey == "" {
		return nil, errors.New("API Key is required")
	}

	// Create a unique user for OpenAI
	h := fnv.New32a()
	h.Write([]byte(apiKey))
	user := "gopenai-" + strconv.Itoa(int(h.Sum32()))

	return &GOpenAI{
		API_KEY: apiKey,
		Verbose: verbose,
		User:    user,
		httpClient: &http.Client{
			Timeout: time.Second * 60,
		},
	}, nil
}
