package goai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/spf13/viper"
)

func (goai Client) ImageGen(prompt, imageFile string, n int) ImageResponse {
	var oaiRequest interface{}
	oaiResponse := ImageResponse{}

	// Check if we are editing an image or generating a new one
	if imageFile != "" {

		// Create the JSON Request Body
		oaiRequest = &ImageEditRequest{
			N:              n,
			ResponseFormat: "url",
			Prompt:         prompt,
			User:           goai.User,
			Size:           viper.GetString("openAI_image_size"),
		}
		goai.UploadImage(oaiRequest, &oaiResponse, goai.Endpoint+"images/edits", imageFile)

	} else { // Generate a new image

		oaiRequest = &ImageRequest{
			N:              n,
			ResponseFormat: "url",
			Prompt:         prompt,
			User:           goai.User,
			Size:           viper.GetString("openAI_image_size"),
			Model:          viper.GetString("openAI_image_model"),
		}
		goai.PostJson(oaiRequest, &oaiResponse, goai.Endpoint+"images/generations")
	}
	if goai.Verbose {
		fmt.Println(oaiRequest)
	}
	return oaiResponse
}

func (goai Client) UploadImage(requestJson, responseJson interface{}, endpoint, filePath string) error {

	// Get the absolute path of the file
	fullPath, err := filepath.Abs(filePath)
	if err != nil {
		return err
	}

	// https://platform.openAI_com/docs/api-reference/images/create-edit#images/create-edit-image
	// The image to edit. Must be a valid PNG file, less than 4MB, and square.
	// If mask is not provided, image must have transparency, which will be used as the mask.
	//
	// Open the PNG image file
	file, err := os.Open(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the PNG file to the request
	part, err := writer.CreateFormFile("image", filePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	oaiImageEditJson := requestJson.(*ImageEditRequest)

	// Add the JSON payload to the request
	part, err = writer.CreateFormField("prompt")
	if err != nil {
		return err
	}
	part.Write([]byte(oaiImageEditJson.Prompt))

	part, err = writer.CreateFormField("n")
	if err != nil {
		return err
	}
	part.Write([]byte(strconv.Itoa(oaiImageEditJson.N)))

	part, err = writer.CreateFormField("size")
	if err != nil {
		return err
	}
	part.Write([]byte(oaiImageEditJson.Size))

	part, err = writer.CreateFormField("user")
	if err != nil {
		return err
	}
	part.Write([]byte(oaiImageEditJson.User))

	// Close the multipart writer
	err = writer.Close()
	if err != nil {
		return err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", endpoint, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+goai.API_KEY)

	if goai.Verbose {
		// trace()
		fmt.Println("Request Body: ", req.Body)
		fmt.Println("Request JSON: ", oaiImageEditJson)
	}

	// Send the request
	fmt.Println("⏳ Uploading File: " + fullPath)
	resp, err := goai.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	// Read the JSON Response Body
	jsonString, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check for API Errors
	err = httpCatchErr(resp, jsonString)
	if err != nil {
		return err
	}

	// Unmarshal the JSON Response Body
	err = json.Unmarshal([]byte(jsonString), &responseJson)
	if err != nil {
		return err
	}
	if goai.Verbose {
		// trace()
		fmt.Println(string(jsonString))
	}

	// Close the HTTP Response Body
	defer resp.Body.Close()
	return nil
}
