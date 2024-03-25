package config

import (
	"fmt"
	"os"
)

func (conf *Config) VPrint(msg string) {
	if conf.Verbose {
		fmt.Fprintln(os.Stderr, msg)
	}
}
