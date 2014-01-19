package tinypng

import (
	"errors"
	"fmt"
	"os"
)

// Basic error checking
func check(err error) {
	if err != nil {
		fmt.Println("Error:", err)

		os.Exit(1)
	}
}

func e(message string) error {
	return errors.New(message)
}
