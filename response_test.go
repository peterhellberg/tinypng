package tinypng

import (
	. "github.com/smartystreets/goconvey/convey"

	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestResponse(t *testing.T) {

	Convey("PopulateFromHTTPResponse", t, func() {
		Convey("Successful HTTP response", func() {

			r := populatedResponse(201, "http://foo",
				`{"input":{"size":207},"output":{"size":63,"ratio":0.307}}`)

			Convey("Input", func() {
				So(r.Input.Size, ShouldEqual, 207)
			})

			Convey("Output", func() {
				So(r.Output.Size, ShouldEqual, 63)
				So(r.Output.Ratio, ShouldEqual, 0.307)
			})

			Convey("URL", func() {
				So(r.URL, ShouldEqual, "http://foo")
			})
		})

		Convey("Unsuccessful HTTP response", func() {
			r := populatedResponse(404, "",
				`{"error":"BadSignature","message":"Does not appear to be a PNG file"}`)

			Convey("Error", func() {
				So(r.Error, ShouldEqual, "BadSignature")
			})

			Convey("Message", func() {
				So(r.Message, ShouldEqual, "Does not appear to be a PNG file")
			})
		})
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
		Body:       nopCloser{bytes.NewBufferString(body)},
	}
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }
