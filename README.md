# GOAI
Go [OpenAI API](https://platform.openai.com/docs/api-reference) Library



`go get github.com/vekjja/goai`

```golang
// main.go
package main

import (
	"os"

	"github.com/vekjja/goai"
)

func main() {
	oai := goai.NewClient(os.Getenv("OPENAI_API_KEY"), false)
	oaiRes, err := oai.TextCompletion("This is a test.")
	if err != nil {
		panic(err)
	}
	println(oaiRes.Choices[0].Text)
}
```