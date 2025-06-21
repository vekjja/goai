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

func (goai Client) ImageGen(prompt, imageModel string, imageFile string, imageSize string, n int) (ImageResponse, error) {
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

// func (goai Client) UploadImage(requestJson, responseJson interface{}, endpoint, filePath string) error {

// 	// Get the absolute path of the file
// 	fullPath, err := filepath.Abs(filePath)
// 	if err != nil {
// 		return err
// 	}

// 	// https://platform.openAI_com/docs/api-reference/images/create-edit#images/create-edit-image
// 	// The image to edit. Must be a valid PNG file, less than 4MB, and square.
// 	// If mask is not provided, image must have transparency, which will be used as the mask.
// 	//
// 	// Open the PNG image file
// 	file, err := os.Open(fullPath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// Create a new multipart writer
// 	body := &bytes.Buffer{}
// 	writer := multipart.NewWriter(body)

// 	// Add the PNG file to the request
// 	part, err := writer.CreateFormFile("image", filePath)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = io.Copy(part, file)
// 	if err != nil {
// 		return err
// 	}

// 	oaiImageEditJson := requestJson.(*ImageEditRequest)

// 	// Add the JSON payload to the request
// 	part, err = writer.CreateFormField("prompt")
// 	if err != nil {
// 		return err
// 	}
// 	part.Write([]byte(oaiImageEditJson.Prompt))

// 	part, err = writer.CreateFormField("n")
// 	if err != nil {
// 		return err
// 	}
// 	part.Write([]byte(strconv.Itoa(oaiImageEditJson.N)))

// 	part, err = writer.CreateFormField("size")
// 	if err != nil {
// 		return err
// 	}
// 	part.Write([]byte(oaiImageEditJson.Size))

// 	part, err = writer.CreateFormField("user")
// 	if err != nil {
// 		return err
// 	}
// 	part.Write([]byte(oaiImageEditJson.User))

// 	// Close the multipart writer
// 	err = writer.Close()
// 	if err != nil {
// 		return err
// 	}

// 	// Create a new HTTP request
// 	req, err := http.NewRequest("POST", endpoint, body)
// 	if err != nil {
// 		return err
// 	}
// 	req.Header.Set("Content-Type", writer.FormDataContentType())
// 	req.Header.Set("Authorization", "Bearer "+goai.API_KEY)

// 	if goai.Verbose > 0 {
// 		// trace()
// 		fmt.Println("Request Body: ", req.Body)
// 		fmt.Println("Request JSON: ", oaiImageEditJson)
// 	}

// 	// Send the request
// 	fmt.Println("⏳ Uploading File: " + fullPath)
// 	resp, err := goai.HTTPClient.Do(req)
// 	if err != nil {
// 		return err
// 	}

// 	// Read the JSON Response Body
// 	// jsonString, err := io.ReadAll(resp.Body)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// Check for API Errors
// 	jsonString, err := httpCatchErr(resp)
// 	if err != nil {
// 		return err
// 	}

// 	// Unmarshal the JSON Response Body
// 	err = json.Unmarshal([]byte(jsonString), &responseJson)
// 	if err != nil {
// 		return err
// 	}
// 	if goai.Verbose > 0 {
// 		// trace()
// 		fmt.Println(string(jsonString))
// 	}

// 	// Close the HTTP Response Body
// 	defer resp.Body.Close()
// 	return nil
// }
