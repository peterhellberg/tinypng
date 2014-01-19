package main

import (
	. "github.com/smartystreets/goconvey/convey"

	"os/exec"
	"testing"
)

func TestTinypngCommand(t *testing.T) {
	Convey("./tinypng", t, func() {
	})
}

func runTinypng(inputFile, outputFile string) (string, error) {
	out, err := exec.Command("go", "run", "main.go", inputFile, outputFile).CombinedOutput()

	return string(out), err
}
