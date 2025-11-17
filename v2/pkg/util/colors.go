package util

import "os"

// These are messed up, something is something else
var (
	Blue       = "\u001b[38;5;45m"
	Grey       = "\u001b[38;5;252m"
	Red        = "\u001b[38;5;42m"
	White      = "\u001b[38;5;255m"
	Background = "\u001b[48;5;239m"
	Reset      = "\u001b[0m"
	Reverse    = "\u001b[7m"
)

// InitColors disables colors if --no-color flag or NO_COLOR env var is set
func InitColors(noColor bool) {
	if noColor || os.Getenv("NO_COLOR") != "" {
		Blue = ""
		Grey = ""
		Red = ""
		White = ""
		Background = ""
		Reset = ""
		Reverse = ""
	}
}
