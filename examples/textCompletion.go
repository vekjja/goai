package main

import (
	"os"

	"github.com/seemywingz/goai"
)

func main() {
	oai := goai.NewClient(os.Getenv("OPENAI_API_KEY"), false)
	oaiRes, err := oai.TextCompletion("This is a test.")
	if err != nil {
		panic(err)
	}
	for _, choice := range oaiRes.Choices {
		println(choice.Text)
	}
}
