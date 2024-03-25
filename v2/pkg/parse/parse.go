package parse

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type CookParse struct {
	Args        []string
	showError   bool
	Help        string
	userDefined map[string]string
	//  Help = ""
	//  userDefined = make(map[string]string)
	//  p.Args = os.p.Args[1:]
	//  showError = false
	//  Help = ""
	//  userDefined = make(map[string]string)
}

func (p *CookParse) Boolean(flag, flagL string) bool {
	for i, cmd := range p.Args {
		if cmd == flag || cmd == flagL {
			p.Args = append(p.Args[:i], p.Args[i+1:]...)
			return true
		}
	}
	return false
}

func (p *CookParse) String(flag, flagL string) string {
	l := len(p.Args)
	for i, cmd := range p.Args {
		if cmd == flag || cmd == flagL {
			if i+1 == l {
				fmt.Printf("Err: Flag '%s' doesn't have any value", cmd)
				os.Exit(0)
			}
			value := p.Args[i+1]
			p.Args = append(p.Args[:i], p.Args[i+2:]...)
			return value
		}
	}

	return ""
}

func (p *CookParse) Integer(flag, flagL string) int {
	intValue := 0
	l := len(p.Args)

	for i, cmd := range p.Args {
		if cmd == flag || cmd == flagL {
			if i+1 == l || p.Args[i+1] == "" {
				fmt.Printf("Err: Flag '%s' doesn't have any value", cmd)
				os.Exit(0)
			} else {
				var err error
				intValue, err = strconv.Atoi(p.Args[i+1])
				// min -= 1
				if err != nil {
					log.Fatalf("Err: Flag %s needs integer value", flag)
				}
			}
			p.Args = append(p.Args[:i], p.Args[i+2:]...)
			return intValue
		}
	}
	return -4541
}

func (p *CookParse) UserDefinedFlags() map[string]string {
	tmp := []string{}

	tmp = append(tmp, p.Args...)

	for _, cmd := range tmp {

		if len(cmd) > 1 && strings.Count(cmd, "-") == 1 && strings.HasPrefix(cmd, "-") {
			value := p.String(cmd, cmd)
			cmd = strings.Replace(cmd, "-", "", 1)
			p.userDefined[cmd] = value
		}
	}

	return p.userDefined
}

// Read square brackets
func ReadSqBr(cmd string) (string, string) {
	c := strings.SplitN(cmd, "[", 2)
	name := c[0]
	values := c[1][:len(c[1])-1]
	return name, values
}

// Read square brackets and separate values by a string/char
func ReadSqBrSepBy(cmd string, sep string) (string, []string) {
	c := strings.SplitN(cmd, "[", 2)

	name := c[0]
	values := strings.Split(c[1][:len(c[1])-1], sep)
	return name, values
}

// Read circular brackets
func ReadCrBr(cmd string) (string, string) {
	c := strings.SplitN(cmd, "(", 2)
	name := c[0]
	values := c[1][:len(c[1])-1]
	return name, values
}

// Read circular brackets and separate values by a string/char
func ReadCrBrSepBy(cmd string, sep string) (string, []string) {
	c := strings.SplitN(cmd, "(", 2)
	name := c[0]
	values := strings.Split(c[1][:len(c[1])-1], sep)
	return name, values
}

func NewParse(args ...string) *CookParse {
	if len(args) == 0 {
		args = os.Args[1:]
	}
	return &CookParse{
		Args:        args,
		showError:   false,
		Help:        "",
		userDefined: make(map[string]string),
	}
}

func (p *CookParse) Parse() {

	if len(os.Args) < 2 {
		print(p.Help)
	}

	if p.showError && len(p.userDefined) > 0 {
		panic(fmt.Sprintf("Undefined Flags%v", p.userDefined))
	}

}

func init() {
	log.SetFlags(0)
}
