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

func httpCatchErr(resp *http.Response, jsonString []byte) error {
	// Check for HTTP Response Errors
	if resp.StatusCode != 200 {
		return errors.New("API Error: " + strconv.Itoa(resp.StatusCode) + "\n" + string(jsonString))
	}
	return nil
}

func (goai GoAi) MakeRequest(request *http.Request, responseJson interface{}) error {

	// Make the HTTP Request
	resp, err := goai.httpClient.Do(request)
	if err != nil {
		return err
	}

	// Read the JSON Response Body
	jsonString, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Check for HTTP Errors
	httpCatchErr(resp, jsonString)
	if goai.Verbose {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("🌐 HTTP Response", b)
	}

	// Unmarshal the JSON Response Body into provided responseJson
	err = json.Unmarshal([]byte(jsonString), &responseJson)
	if err != nil {
		return err
	}
	if goai.Verbose {
		// trace()
		fmt.Println("🌐 HTTP Response String", string(jsonString))
		fmt.Println("🌐 HTTP Response JSON", responseJson)
	}
	// Close the HTTP Response Body
	defer resp.Body.Close()
	return nil
}

func (goai GoAi) PostJson(requestJson, responseJson interface{}, endpoint string) error {
	// Marshal the JSON Request Body
	requestBodyJson, err := json.Marshal(requestJson)
	if err != nil {
		return err
	}
	if goai.Verbose {
		// trace()
		fmt.Println(string(requestBodyJson))
	}
	// Format HTTP Response and Set Headers
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(requestBodyJson))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+goai.API_KEY)
	return goai.MakeRequest(req, responseJson)
}
