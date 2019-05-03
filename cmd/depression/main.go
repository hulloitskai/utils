// Command depression is a tool for reading and writing depression-encoded
// data.
package main

import (
	"fmt"
	"io"
	"os"

	flag "github.com/spf13/pflag"
	"github.com/stevenxie/utils/depression"
)

func main() {
	// Define and parse flags.
	decode := flag.BoolP("decode", "d", false, "Decode from stdin.")
	flag.Parse()

	// Either encode or decode from stdin.
	var translator io.Writer
	if *decode {
		translator = depression.NewDecoder(os.Stdout)
	} else {
		translator = depression.NewEncoder(os.Stdout)
	}

	// Perform translation.
	if _, err := io.Copy(translator, os.Stdin); err != nil {
		die(err)
	}
	if err := os.Stdout.Close(); err != nil {
		die(err)
	}
}

func die(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v", err)
	os.Exit(1)
}
