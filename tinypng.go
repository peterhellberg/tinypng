package tinypng

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	apiUser = "api"
	apiURL  = "https://api.tinypng.com/"
)

// ShrinkFn allows you to Shrink using a
// filename instead of an open file handle.
func ShrinkFn(apiKey string, inputFilename string) (Response, error) {
	inputFile, err := os.Open(inputFilename)

	if err != nil {
		return Response{}, err
	}

	return Shrink(apiKey, inputFile)
}

// Shrink allows you to shrink a PNG file using an open file handle.
func Shrink(apiKey string, inputFile *os.File) (Response, error) {
	req, err := preparePOSTRequest(apiKey, inputFile)
	check(err)

	res, err := sendPOSTRequest(req)
	check(err)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	check(err)

	var r Response

	err = json.Unmarshal(body, &r)
	check(err)

	// Get the output URL from the Location header
	r.URL = res.Header.Get("Location")

	if res.StatusCode != http.StatusCreated {
		return r, errors.New("unauthorized")
	}

	return r, nil
}

// Prepare POST request
func preparePOSTRequest(apiKey string, inputFile *os.File) (*http.Request, error) {
	req, err := http.NewRequest("POST", apiURL+"shrink", inputFile)

	if err != nil {
		return req, err
	}

	// Authenticate using the API key
	req.SetBasicAuth(apiUser, apiKey)

	return req, nil
}

// Send POST request
func sendPOSTRequest(req *http.Request) (*http.Response, error) {
	// Create a HTTP client
	client := &http.Client{}

	// Perform the POST request
	res, err := client.Do(req)
	check(err)

	return res, nil
}

// Basic error checking
func check(err error) {
	if err != nil {
		fmt.Println("Error:", err)

		os.Exit(1)
	}
}
