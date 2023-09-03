package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ubhattac/when"
)

func main() {
	flag.Usage = usage
	verbose := flag.Bool("verbose", false, "Enable verbose descriptions of time")
	flag.Parse()

	input := flag.Arg(0)

	if input == "help" {
		usage()
		os.Exit(0)
	}

	if input == "" {
		usage()
		os.Exit(1)
	}

	var readableTime string
	var err error
	if *verbose {
		readableTime, err = when.WhenVerbose(input)
	} else {
		readableTime, err = when.When(input)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(readableTime)
}

// TODO: adjust usage to indicate verbosity
func usage() {
	const userGuide = `
Usage: when <TIME>

        <TIME> can be any standard computer time format.

        Currently supported formats are:
        * RFC3339
        * Unix time

        [Examples]
        $ when 2023-09-01T14:30:00Z
        $ when 1660460492

`
	fmt.Print(userGuide)
}
