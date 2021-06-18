package core

import (
	"fmt"
	"os"
	"strings"
)

var Banner = `                            
  ░            ░ ░      ░ ░  ░  ░            
  ░ ░        ░ ░ ░ ▒  ░ ░ ░ ▒  ░             
░░▒ ▒░    ░ ▒ ▒░   ░ ▒ ▒░ ░ ░ ▒ ░            
░ ░▒ ▒  ░░ ▒░▒░▒░ ░ ▒░▒░▒░ ▒ ▒▒ ▓▒           
 ▄████▄   ▒█████   ▒█████   ██ ▄█▀           
▒██▀ ▀█  ▒██▒  ██▒▒██▒  ██▒ ██▄█▒            
▒▓█    ▄ ▒██░  ██▒▒██░  ██▒▓███▄░            
▒▓▓▄ ▄██▒▒██   ██░▒██   ██░▓██ █▄             
 ▒▓███▀ ░░ ████▓▒░░ ████▓▒░▒██▒ █▄ ` + `       Gitesh Sharma @giteshnxtlvl
 
 Usage  : cook [PARAM-1 VALUES] [PARAM-2 VALUES] ... [PARAM-N VALUES]  [PATTERN]
          cook -param1 val1,val2 -param2 file.txt param1:param2
 Help   : cook -h 
 Config : cook -config`

// var verbose = false

func VPrint(msg string) {
	if Verbose {
		fmt.Fprintln(os.Stderr, msg)
	}
}

func helpCommand(title string, description string, command string) {
	fmt.Println(Red + "\n\n" + title + White)
	fmt.Println(Grey + description)
	fmt.Printf("%s    $ cook", White)
	for _, v := range strings.Split(command, " ") {
		if strings.HasPrefix(v, "-") {
			v = Green + v
		} else {
			v = Blue + v
		}
		fmt.Printf(" " + v)
	}
}

func ShowHelp() {
	fmt.Println(Banner)

	fmt.Println(Green + "\nGITHUB" + White)
	fmt.Println(Blue + "    https://github.com/giteshnxtlvl/cook" + Reset)

	fmt.Println(Green + "\nFLAGS" + White)
	help := `    -case        : Define Cases
                 * Use for complete list
                     -case A for ALL 
                     -case U for Uppercase
                     -case L for Lowercase
                     -case T for Titlecase
                     -case C for Camelcase

                 * Use column wise, no camel case for this
                     -case 0:U,2:T
                         Column 0 will be in Uppercase
                         Column 2 will be in Titlecase,
                         Rest columns will be default output
                     Multiple Cases
                         -case 0:UT,2:A 
    -min         : Minimum no of columns to print			  
    -config      : Config Information *cook.yaml*
    -config-path : Specify path for custom yaml file.
    -update-all  : Update all file's cache
    -h           : Help`

	fmt.Println(help)

	fmt.Println(Green + "\nBASIC USAGE" + White)
	fmt.Printf("   $ cook %[1]s-start %[2]sadmin%[3]s,%[2]sroot  %[1]s-sep %[2]s_%[3]s,%[2]s-  %[1]s-end %[2]ssecret%[3]s,%[2]scritical  %[2]s/%[3]s:%[1]sstart%[3]s:%[1]ssep%[3]s:%[1]send\n", Green, Blue, White)
	fmt.Printf("   %[3]s$ cook %[2]s/%[3]s:%[2]sadmin%[3]s,%[2]sroot%[3]s:%[2]s_%[3]s,%[2]s-%[3]s:%[2]ssecret%[3]s,%[2]scritical\n", Green, Blue, White)

	fmt.Println(Green + "\nFILE WITH REGEX" + White)
	fmt.Printf("   $ cook %[1]s-s %[2]scompany %[1]s-ext %[2]sraft-large-extensions%[3]s:%[3]s\\.asp.*  %[2]s/%[3]s:%[1]ss%[3]s:%[1]sexp\n", Green, Blue, White, Grey)

	helpCommand("FUNCTIONS", "    Use functions such as date for different variations of values", "-name elliot -birth date(17,Sep,1994) name:birth")
	helpCommand("RANGES", "    Use ranges like [1-100], [A-Z], [a-z] or [A-z] in pattern of command", "-name elliot -birth date(17,Sep,1994) name:birth")
	helpCommand("USING CASES", "    Uppercase, lowercase, titlecase, camelcase or ALL \n    For use case check FLGAS above", "camel:[1-10]:case -case C")
	helpCommand("RAW STRINGS", "    Print value without any parsing/modification.\n    Use to take \",\", \":\", \"`\" or any pre-defined sets or functions as raw strings.", "-date `date(17,Sep,1994)` raw:date")
	helpCommand("PIPE INPUT", "    Use - as param value for pipe input", "-d - d:/:test")
	helpCommand("USING -min", "    Print value without any parsing/modification", "n:n:n -min 1")

	fmt.Println(Reset)
	os.Exit(0)
}
