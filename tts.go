package goai

// TTSRequest represents the request structure for OpenAI's text-to-speech API.
type TTSRequest struct {
	Model          string  `json:"model"`
	Input          string  `json:"input"`
	Voice          string  `json:"voice"`
	ResponseFormat string  `json:"response_format"`
	Speed          float64 `json:"speed"`
}

func (goai Client) TTS(input string) ([]byte, error) {
	req := TTSRequest{
		Input:          input,
		Model:          goai.TTSModel,
		Voice:          goai.Voice,
		Speed:          goai.Speed,
		ResponseFormat: goai.ResponseFormat,
	}
	res, err := goai.PostJson(req, nil, goai.Endpoint+"audio/speech")
	if err != nil {
		return nil, err
	}
	return res, nil
}
