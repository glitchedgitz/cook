package core

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var Banner = fmt.Sprintf(`                            

  ░    ░  ░   ░ ░      ░ ░  ░  ░
  ░ ░   ░    ░ ░ ░ ▒  ░ ░  ▒  ░
░░▒ ▒░ ░  ░ ▒ ▒░   ░ ▒ ▒░ ░ ░ ▒ ░
░ ░▒ ▒  ░░ ▒░▒░▒░ ░ ▒░▒░▒░ ▒ ▒▒ ▓▒        
 ▄████▄   ▒█████   ▒█████   ██ ▄█▀           
▒██▀ ▀█  ▒██▒  ██▒▒██▒  ██▒ ██▄█▒            
▒▓█    ▄ ▒██░  ██▒▒██░  ██▒▓███▄░                         
 ▒▓███▀ ░░ ████▓▒░░ ████▓▒░▒██▒ █▄

      THE WORDLIST'S FRAMEWORK

            Version %s
    Gitesh Sharma @giteshnxtlvl
`, Version)

func VPrint(msg string) {
	if Verbose {
		fmt.Fprintln(os.Stderr, msg)
	}
}

func HelpMode(h []string) {
	if len(h) <= 0 {
		log.Fatalln("Ask for these... case, encode, file, function, patterns or usage")
	}

	help := strings.ToLower(h[0])

	if help == "case" {
		fmt.Fprintln(os.Stderr, CaseHelp)
	} else if help == "encode" {
		fmt.Fprintln(os.Stderr, EncodeHelp)
	} else if help == "function" {
		fmt.Fprintln(os.Stderr, FuncHelp)
	} else if help == "pattern" {
		fmt.Fprintln(os.Stderr, "Pattern!! under construction, better not learn about them")
	} else if help == "usage" {
		fmt.Fprintln(os.Stderr, UsageHelp)
	} else {
		fmt.Fprintln(os.Stderr, "Ask for these... case, encode, file, function, patterns or usage")
	}
	os.Exit(0)
}

func ShowHelp() {
	fmt.Fprintln(os.Stderr, Banner)
	fmt.Fprintln(os.Stderr, FlagsHelp)
	fmt.Fprintln(os.Stderr, Reset)
	os.Exit(0)
}
