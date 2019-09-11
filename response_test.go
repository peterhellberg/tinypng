package tinypng

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPopulateFromHTTPResponse(t *testing.T) {
	t.Run("Successful HTTP response", func(t *testing.T) {
		r := populatedResponse(201, "http://foo",
			`{"input":{"size":207},"output":{"size":63,"ratio":0.307}}`)

		if got, want := r.Input.Size, int32(207); got != want {
			t.Fatalf("r.Input.Size = %d, want %d", got, want)
		}

		if got, want := r.Output.Size, int32(63); got != want {
			t.Fatalf("r.Output.Size = %d, want %d", got, want)
		}

		if got, want := r.Output.Ratio, 0.307; got != want {
			t.Fatalf("r.Output.Ratio = %f, want %f", got, want)
		}

		if got, want := r.URL, "http://foo"; got != want {
			t.Fatalf("r.URL = %q, want %q", got, want)
		}
	})

	t.Run("Unsuccessful HTTP response", func(t *testing.T) {
		r := populatedResponse(404, "",
			`{"error":"BadSignature","message":"Does not appear to be a PNG file"}`)

		if got, want := r.Error, "BadSignature"; got != want {
			t.Fatalf("r.Error = %q, want %q", got, want)
		}

		if got, want := r.Message, "Does not appear to be a PNG file"; got != want {
			t.Fatalf("r.Message = %q, want %q", got, want)
		}
	})
}

func populatedResponse(statusCode int, location, body string) Response {
	var r Response

	r.PopulateFromHTTPResponse(fakeHTTPResponse(statusCode, location, body))

	return r
}

func fakeHTTPResponse(statusCode int, location, body string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Header:     http.Header{"Location": {location}},
		Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
	}
}
