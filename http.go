package goai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func httpCatchErr(resp *http.Response) ([]byte, error) {
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// Check for HTTP Response Errors
	if resp.StatusCode != 200 {
		errJson := ErrorResponse{}
		err = json.Unmarshal(resBody, &errJson)
		return nil, errors.New("API Error: " + strconv.Itoa(resp.StatusCode) + "\n" + errJson.Error.Message)
	}
	return resBody, nil
}

func (goai Client) MakeRequest(request *http.Request, responseJson interface{}) ([]byte, error) {

	// Make the HTTP Request
	resp, err := goai.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	// Check for HTTP Errors
	jsonString, err := httpCatchErr(resp)
	if err != nil {
		return nil, err
	}

	if goai.Verbose {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("🌐 HTTP Response", b)
	}

	// Close the HTTP Response Body
	defer resp.Body.Close()

	if responseJson == nil {
		return jsonString, nil
	}

	// Unmarshal the JSON Response Body into provided responseJson
	err = json.Unmarshal([]byte(jsonString), &responseJson)
	if err != nil {
		return nil, errors.New("Error Unmarshalling JSON Response: " + err.Error())
	}
	if goai.Verbose {
		// trace()
		fmt.Println("🌐 HTTP Response String", string(jsonString))
		fmt.Println("🌐 HTTP Response JSON", responseJson)
	}
	return jsonString, nil
}

func (goai Client) PostJson(requestJson, responseJson interface{}, endpoint string) ([]byte, error) {
	// Marshal the JSON Request Body
	requestBodyJson, err := json.Marshal(requestJson)
	if err != nil {
		return nil, err
	}
	if goai.Verbose {
		fmt.Println(string(requestBodyJson))
	}
	// Format HTTP Response and Set Headers
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBodyJson))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+goai.API_KEY)
	return goai.MakeRequest(req, responseJson)
}
