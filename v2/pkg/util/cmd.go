package util

import (
	"fmt"
)

func TerminalColor(text, color string) string {
	return fmt.Sprintf("%s%s%s", color, text, Reset)
}

func TerminalLink(url, text, color string) string {
	// When colors are disabled (Reset is empty), skip hyperlink escape sequences
	if Reset == "" {
		return text
	}
	return fmt.Sprintf("%s\033]8;;%s\033\\%s\033]8;;\033\\%s", color, url, text, Reset)
}
