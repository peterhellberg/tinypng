package main

import (
	. "github.com/smartystreets/goconvey/convey"

	"os/exec"
	"testing"
)

func TestTinypngCommand(t *testing.T) {
	Convey("tinypng", t, func() {
		Convey("outputs usage instructions if no args", func() {
			out, _ := execGo("run", "main.go")

			So(out, ShouldContainSubstring, "tinypng <input.png> [output.png]")
		})

		Convey("outputs error if unknown file", func() {
			out, _ := execGo("run", "main.go", "unknown.png")

			So(out, ShouldContainSubstring, "Input file does not exist.")
		})

		Convey("outputs error it invalid file", func() {
			out, _ := execGo("run", "main.go", "../testdata/invalid.png")

			So(out, ShouldContainSubstring, "Input file is not a valid PNG file.")
		})

		Convey("outputs note about adding TINYPNG_API_KEY to ENV", func() {
			out, _ := execGo("run", "main.go", "../testdata/valid.png")

			So(out, ShouldContainSubstring, "TINYPNG_API_KEY")
		})
	})
}

func execGo(args ...string) (string, error) {
	out, err := exec.Command("go", args...).CombinedOutput()

	return string(out), err
}
