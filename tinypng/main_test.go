package main

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestTinypngCommand(t *testing.T) {
	t.Run("outputs usage instructions if no args", func(t *testing.T) {
		var (
			got  = execGo("run", "main.go")
			want = "tinypng <input.png> [output.png]"
		)

		if !strings.Contains(got, want) {
			t.Fatalf("%q does not contain %q", got, want)
		}
	})

	t.Run("outputs error if unknown file", func(t *testing.T) {
		var (
			got  = execGo("run", "main.go", "unknown.png")
			want = "Input file does not exist."
		)

		if !strings.Contains(got, want) {
			t.Fatalf("%q does not contain %q", got, want)
		}
	})

	t.Run("outputs error it invalid file", func(t *testing.T) {
		var (
			got  = execGo("run", "main.go", "../testdata/invalid.png")
			want = "Input file is not a valid PNG or JPEG file."
		)

		if !strings.Contains(got, want) {
			t.Fatalf("%q does not contain %q", got, want)
		}
	})

	t.Run("outputs note about adding TINYPNG_API_KEY to ENV", func(t *testing.T) {
		os.Setenv("TINYPNG_API_KEY", "")

		var (
			got  = execGo("run", "main.go", "../testdata/valid.png")
			want = "TINYPNG_API_KEY"
		)

		if !strings.Contains(got, want) {
			t.Fatalf("%q does not contain %q", got, want)
		}
	})
}

func execGo(args ...string) string {
	out, _ := exec.Command("go", args...).CombinedOutput()

	return string(out)
}
