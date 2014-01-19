package tinypng

import (
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
	res, err := uploadPNG(apiKey, inputFile)
	check(err)

	defer res.Body.Close()

	var r Response

	r.PopulateFromHTTPResponse(res)

	if res.StatusCode == 201 {
		return r, nil
	}

	return r, e("unsuccessful")
}
