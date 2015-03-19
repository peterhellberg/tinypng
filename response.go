package tinypng

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// Response from the TinyPNG API
type Response struct {
	Input   Input
	Output  Output
	Error   string
	Message string
	URL     string
}

// Input size
type Input struct {
	Size int32
}

// Output size, ratio and url
type Output struct {
	Size  int32
	Ratio float64
}

// PopulateFromHTTPResponse populates response based on HTTP response
func (r *Response) PopulateFromHTTPResponse(res *http.Response) {
	body, err := ioutil.ReadAll(res.Body)
	check(err)

	err = json.Unmarshal(body, &r)

	check(err)

	// Get the output URL from the Location header
	r.URL = res.Header.Get("Location")
}

// SaveAs downloads and saves the compressed PNG file
func (r *Response) SaveAs(fn string) {
	resp, err := http.Get(r.URL)
	check(err)

	defer resp.Body.Close()

	out, err := os.Create(fn)
	check(err)

	defer out.Close()

	io.Copy(out, resp.Body)
}

// Print a line of statistics
func (r *Response) Print() {
	fmt.Print("Input size: ", r.Input.Size)
	fmt.Print(" Output size: ", r.Output.Size)
	fmt.Println(" Ratio:", r.Output.Ratio)
	fmt.Println("\n", r.URL)
}
