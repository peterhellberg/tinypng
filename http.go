package tinypng

import (
	"net/http"
	"os"
)

func uploadPNG(apiKey string, inputFile *os.File) (*http.Response, error) {
	req, err := preparePOSTRequest(apiKey, inputFile)
	check(err)

	return sendHTTPRequest(req)
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

// Send HTTP request
func sendHTTPRequest(req *http.Request) (*http.Response, error) {
	// Create a HTTP client
	client := &http.Client{}

	// Perform the POST request
	res, err := client.Do(req)
	check(err)

	return res, nil
}
