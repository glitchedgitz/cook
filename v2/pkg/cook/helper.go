package cook

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

// These are messed up, something is something else
var (
	Blue  = color.FgBlue.Render
	Cyan  = color.FgCyan.Render
	Red   = color.FgRed.Render
	Green = color.FgGreen.Render
	White = color.FgWhite.Render
)

// var (
//
//	Blue    = "\u001b[38;5;45m"
//	Grey    = "\u001b[38;5;252m"
//	Red     = "\u001b[38;5;42m"
//	White   = "\u001b[38;5;255m"
//	Reset   = "\u001b[0m"
//	Reverse = "\u001b[7m"
// )

func VPrint(msg string) {
	if Verbose {
		fmt.Fprintln(os.Stderr, msg)
	}
}
