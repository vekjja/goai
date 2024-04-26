package goai

import ()

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
