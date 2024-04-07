package main

import (
	"fmt"
	"os"
)

func VPrint(msg string) {
	if verbose {
		fmt.Fprintln(os.Stderr, msg)
	}
}
