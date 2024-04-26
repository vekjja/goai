package main

import (
	"fmt"
	"os"

	"github.com/seemywingz/goai"
)

var messages = []goai.Message{{
	Role:    "system",
	Content: "You are a helpful assistant.",
}}

func main() {
	ai := goai.DefaultClient(os.Getenv("OPENAI_API_KEY"), false)
	oaiRes, err := ai.ChatCompletion(messages)
	if err != nil {
		panic(err)
	}
	fmt.Println(oaiRes.Choices[0].Message.Content)
}
