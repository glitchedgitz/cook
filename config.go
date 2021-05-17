package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"gopkg.in/yaml.v2"
)

func showHelp() {

	fmt.Println(green + "\nGITHUB" + white)
	fmt.Println(blue + "    https://github.com/giteshnxtlvl/cook" + reset)

	fmt.Println(green + "\nFLAGS" + white)
	help := `    -case   : Define Cases
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

    -min    : Minimum no of columns to print. (Default min = no of columns)
              Same as minimum of crunch			  
    -config : Config Information *cook.yaml*
    -h      : Help
	`
	fmt.Println(help)

	fmt.Println(green + "\nBASIC USAGE" + white)
	fmt.Printf("   $ cook %[1]s-start %[2]sadmin%[3]s,%[2]sroot  %[1]s-sep %[2]s_%[3]s,%[2]s-  %[1]s-end %[2]ssecret%[3]s,%[2]scritical  %[2]s/%[3]s:%[1]sstart%[3]s:%[1]ssep%[3]s:%[1]send\n", green, blue, white)
	fmt.Printf("   %[3]s$ cook %[2]s/%[3]s:%[2]sadmin%[3]s,%[2]sroot%[3]s:%[2]s_%[3]s,%[2]s-%[3]s:%[2]ssecret%[3]s,%[2]scritical\n", green, blue, white)

	fmt.Println(green + "\nFILE WITH REGEX" + white)
	fmt.Printf("   $ cook %[1]s-s %[2]scompany %[1]s-ext %[2]sraft-large-extensions%[3]s:%[3]s\\.asp.*  %[2]s/%[3]s:%[1]ss%[3]s:%[1]sexp\n", green, blue, white, purple)

	os.Exit(0)
}

func cookConfig() {
	if len(configPath) > 0 {
		configFile = configPath
	} else if len(os.Getenv("COOK")) > 0 {
		configFile = os.Getenv("COOK")
	}

	vPrint(fmt.Sprintf("Config File  %s", configFile))

	if _, err := os.Stat(configFile); err == nil {

		content, err = ioutil.ReadFile(configFile)
		if err != nil {
			fmt.Printf("Err: Reading Config File %v\n", err)
		}

		if len(content) == 0 {
			config := getConfigFile()
			ioutil.WriteFile(configFile, []byte(config), 0644)
			content = []byte(config)
		}

	} else {

		err := os.MkdirAll(path.Join(home, ".config", "cook"), os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		config := getConfigFile()
		err = ioutil.WriteFile(configFile, []byte(config), 0644)
		if err != nil {
			fmt.Printf("Err: Reading Config File %v\n", err)
		}
	}

	err := yaml.Unmarshal([]byte(content), &m)

	if err != nil {
		fmt.Printf("error: %v", err)
	}
}

func showMap(set string) {
	fmt.Println("\n" + green + strings.ToUpper(set) + reset)

	keys := []string{}
	for k := range m[set] {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf(blue+"  %-12s "+white+"%v\n", k, m[set][k])
	}
}

func showConfig() {

	fmt.Println(green + "\nCOOK.YAML " + reset)
	fmt.Printf(blue+"  %-11s "+white+" %v\n", "Location", configFile)

	showMap("charSet")
	showMap("files")
	showMap("lists")
	showMap("patterns")
	showMap("extensions")

	os.Exit(0)
}
