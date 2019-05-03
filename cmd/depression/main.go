// Command depression is a tool for reading and writing depression-encoded
// data.
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/stevenxie/utils/depression"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Encode and decode using depression representation."
	app.UsageText = "depression [options] < (input) > (output)\n" +
		"   depression [options] (file)"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "decode,d",
			Usage: "Decode from stdin.",
		},
		cli.BoolFlag{
			Name:   "help,h",
			Hidden: true,
		},
	}

	// Configure app.
	app.Author = "Steven Xie <hello@stevenxie.me>"
	app.Action = run
	app.HideVersion = true
	app.HideHelp = true
	app.EnableBashCompletion = true

	// Run and exit with appropriate code.
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	if c.Bool("help") {
		cli.ShowAppHelp(c) // show help and exit
		return nil
	}
	args := c.Args()
	if len(args) > 1 {
		fmt.Fprintf(os.Stderr, "Error: expected 0 or 1 arguments, received %d.\n\n",
			len(args))
		cli.ShowAppHelpAndExit(c, 1)
	}

	// Either encode or decode.
	var (
		translator io.Writer
		decode     = c.Bool("decode")
	)
	if decode {
		translator = depression.NewDecoder(os.Stdout)
	} else {
		translator = depression.NewEncoder(os.Stdout)
	}

	// Either use stdin or read a file.
	var src io.Reader
	if len(args) == 0 {
		src = os.Stdin
	} else {
		var err error
		if src, err = os.Open(args[0]); err != nil {
			return err
		}
	}

	// Perform translation.
	if _, err := io.Copy(translator, src); err != nil {
		return err
	}
	if !decode {
		fmt.Println() // append newline for formatting
	}
	return nil
}
