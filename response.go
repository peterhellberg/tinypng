package tinypng

import (
	"fmt"
	"io"
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
	fmt.Println("\n", r.URL, "\n")
}
