package goai

// ImageRequest represents the request structure for OpenAI's image generation API.
type ImageRequest struct {
	Prompt         string `json:"prompt"`
	N              int    `json:"n"`
	Size           string `json:"size"`
	ResponseFormat string `json:"response_format"`
	User           string `json:"user"`
	Model          string `json:"model"`
}

type ImageResponse struct {
	Created int `json:"created"`
	Data    []struct {
		B64JSON string `json:"b64_json"`
		URL     string `json:"url"`
	} `json:"data"`
	Usage struct {
		TotalTokens        int `json:"total_tokens"`
		InputTokens        int `json:"input_tokens"`
		OutputTokens       int `json:"output_tokens"`
		InputTokensDetails struct {
			TextTokens  int `json:"text_tokens"`
			ImageTokens int `json:"image_tokens"`
		} `json:"input_tokens_details"`
	} `json:"usage"`
}

func (goai Client) ImageGen(prompt, imageModel string, imageSize string, n int) (ImageResponse, error) {
	oaiResponse := ImageResponse{}

	oaiRequest := &ImageRequest{
		N:              n,
		ResponseFormat: "url",
		Prompt:         prompt,
		User:           goai.User,
		Size:           imageSize,
		Model:          imageModel,
	}
	_, err := goai.PostJson(oaiRequest, &oaiResponse, goai.Endpoint+"images/generations")
	return oaiResponse, err
}
