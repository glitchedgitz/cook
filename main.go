package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

var gopath = os.Getenv("GOPATH")

var m = make(map[interface{}]map[string][]string)

var params = make(map[string][]string)
var pattern = []string{}

var banner = `

                             
  ░            ░ ░      ░ ░  ░  ░   
 ░        ░ ░ ░ ▒  ░ ░ ░ ▒  ░ ░░ ░ 
  ░  ▒     ░ ▒ ▒░   ░ ▒ ▒░ ░ ░▒ ▒░
░ ░▒ ▒  ░░ ▒░▒░▒░ ░ ▒░▒░▒░ ▒ ▒▒ ▓▒
 ▄████▄   ▒█████   ▒█████   ██ ▄█▀           A CUSTOMIZABLE WORDLIST
▒██▀ ▀█  ▒██▒  ██▒▒██▒  ██▒ ██▄█▒            AND PASSWORD GENERATOR
▒▓█    ▄ ▒██░  ██▒▒██░  ██▒▓███▄░ 
▒▓▓▄ ▄██▒▒██   ██░▒██   ██░▓██ █▄            by Gitesh Sharma 
 ▒ ▓███▀ ░░ ████▓▒░░ ████▓▒░▒██▒ █▄ V1       @giteshnxtlvl

`

var help = `

COMPLETE USAGE
	https://github.com/giteshnxtlvl/cook

BASIC USAGE
	cook -start admin,root  -sep _,-  -end secret,critical  start:sep:end

OUTPUT
	admin_secret
	admin_critical
	admin-secret
	admin-critical
	root_secret
	root_critical
	root-secret
	root-critical
`

var config = `

# Character set like crunch
charSet:
    n     : [0123456789]
    A     : [ABCDEFGHIJKLMNOPQRSTUVWXYZ]
    a     : [abcdefghijklmnopqrstuvwxyz]
    aAn   : [abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789]
    An    : [ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789]
    an    : [abcdefghijklmnopqrstuvwxyz0123456789]
    aA    : [abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ]
    s     : ["!#$%&'()*+,-./:;<=>?@[\\]^_` + "`" + `{|}~&\""]
    all   : ["!#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_` + "`" + `abcdefghijklmnopqrstuvwxyz{|}~\""]

# File to access from anywhere
files:
    raft_ext: [E:\\tools\\wordlists\\SecLists\\Discovery\\Web-Content\\raft-large-extensions.txt]
    robot_1000: [E:\\tools\\wordlists\\SecLists\\Discovery\\Web-Content\\RobotsDisallowed-Top1000.txt]

# Create your word's set
words:
	admin_set: [admin, root, su, administration]
	password_set: [123, "@123", "#123"]
	months : [January, February, March, April, May, June, July, August, September, October, November, December]
	mons : [Jan, Feb, Mar, Apr, May, Jun, Jul, Aug, Sep, Oct, Nov, Dec]

# Extension Set, . will added before using this
extensions:
    archive: [7z, a, apk, xapk, ar, bz2, cab, cpio, deb, dmg, egg, gz, iso, jar, lha, mar, pea, rar, rpm, s7z, shar, tar, tbz2, tgz, tlz, war, whl, xpi, zip, zipx, xz, pak]
    config : [conf, config]
    sheet  : [ods, xls, xlsx, csv, ics vcf]
    exec   : [exe, msi, bin, command, sh, bat, crx]
    code   : [c, cc, class, clj, cpp, cs, cxx, el, go, h, java, lua, m, m4, php, php3, php5, php7, pl, po, py, rb, rs, sh, swift, vb, vcxproj, xcodeproj, xml, diff, patch, js, jsx]
    web    : [html, html5, htm, css, js, jsx, less, scss, wasm, php, php3, php5, php7]
    backup : [bak, backup, backup1, backup2]
    slide  : [ppt, odp]
    font   : [eot, otf, ttf, woff, woff2]
    text   : [doc, docx, ebook, log, md, msg, odt, org, pages, pdf, rtf, rst, tex, txt, wpd, wps]
    audio  : [aac, aiff, ape, au, flac, gsm, it, m3u, m4a, mid, mod, mp3, mpa, pls, ra, s3m, sid, wav, wma, xm]
    book   : [mobi, epub, azw1, azw3, azw4, azw6, azw, cbr, cbz]
    video  : [3g2, 3gp, aaf, asf, avchd, avi, drc, flv, m2v, m4p, m4v, mkv, mng, mov, mp2, mp4, mpe, mpeg, mpg, mpv, mxf, nsv, ogg, ogv, ogm, qt, rm, rmvb, roq, srt, svi, vob, webm, wmv, yuv]
    image  : [3dm, 3ds, max, bmp, dds, gif, jpg, jpeg, png, psd, xcf, tga, thm, tif, tiff, yuv, ai, eps, ps, svg, dwg, dxf, gpx, kml, kmz, webp]

`

func parseCommand(list []string, a string) ([]string, bool) {
	for i, b := range list {
		if b == a {
			return append(list[:i], list[i+1:]...), true
		}
	}
	return list, false
}

func parseCommandArg(list []string, a string) ([]string, string) {
	for i, b := range list {
		if b == a {
			return append(list[:i], list[i+2:]...), list[i+1]
		}
	}
	return list, ""
}

func stringInSlice(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func findRegex(file, expresssion string) []string {
	founded := []string{}

	content, _ := ioutil.ReadFile(file)

	r, err := regexp.Compile(expresssion)
	if err != nil {
		println(err)
	}

	data := strings.ReplaceAll(string(content), "\r\n", "\n")
	extensions_list := r.FindAllString(data, -1)

	for _, found := range extensions_list {
		if !stringInSlice(founded, found) {
			founded = append(founded, found)
		}
	}

	return founded
}

var toUpper = false
var toLower = false
var colCases = make(map[int][]string)

//Parse Input
func parseInput(commands []string) {

	if len(commands) == 0 {
		fmt.Println(banner)
		os.Exit(0)
	}

	if stringInSlice(commands, "-h") {
		fmt.Println(banner)
		fmt.Println(help)
		fmt.Println(config)
		os.Exit(0)
	}

	commands, toUpper = parseCommand(commands, "-upper")
	commands, toLower = parseCommand(commands, "-lower")
	commands, caseValue := parseCommandArg(commands, "-case")

	last := len(commands) - 1
	pattern = strings.Split(commands[last], ":")

	if caseValue != "" {
		if !strings.Contains(caseValue, ":") {
			tmp := strings.Split(caseValue, "")
			for i := 0; i < len(pattern); i++ {
				colCases[i] = tmp
			}

		} else {
			for _, val := range strings.Split(caseValue, ",") {
				v := strings.SplitN(val, ":", 2)
				i, err := strconv.Atoi(v[0])
				if err != nil {
					panic(err)
				}
				colCases[i] = strings.Split(v[1], "")
			}
		}
	}
	// cook admin:_,-:test -case 1:UL,2:T,A

	for i, cmd := range commands[:last] {

		if strings.HasPrefix(cmd, "-") {
			cmd = strings.Replace(cmd, "-", "", 1)
			value := commands[i+1]
			values := []string{}

			if strings.Contains(value, ":") {
				t := strings.SplitN(value, ":", 2)
				file := t[0]
				reg := t[1]

				if strings.HasSuffix(file, ".txt") {
					values = findRegex(file, reg)
				} else {
					values = strings.Split(value, ",")
				}
			} else if strings.HasSuffix(value, ".txt") {
				content, err := ioutil.ReadFile(value)

				if err != nil {
					values = strings.Split(value, ",")
				} else {
					fileData := strings.ReplaceAll(string(content), "\r\n", "\n")
					values = strings.Split(fileData, "\n")
				}
			} else {
				values = strings.Split(value, ",")
			}

			params[cmd] = values
		}
	}
}

func cookConfig() {

	configFile := os.Getenv("COOK")

	content := []byte(config)

	if len(configFile) == 0 {

	} else if _, err := os.Stat(configFile); err == nil {
		// If file exists
		var err2 error
		content, err2 = ioutil.ReadFile(configFile)
		if err2 != nil {
			fmt.Printf("Config File Reading Error: %v\n", err2)
		}

		//If file is empty
		if len(content) == 0 {
			ioutil.WriteFile(configFile, []byte(config), 0644)
			content = []byte(config)
		}
	} else if os.IsNotExist(err) {
		err := ioutil.WriteFile(configFile, []byte(config), 0644)
		if err != nil {
			fmt.Printf("Config File Writing Error: %v\n", err)
		}
	}

	err := yaml.Unmarshal([]byte(content), &m)

	if err != nil {
		fmt.Printf("error: %v", err)
	}
}

func main() {

	cookConfig()
	parseInput(os.Args[1:])

	final := []string{}

	for i, param := range pattern {
		tmp1 := []string{}

		if i == 0 {
			final = append(final, "")
		}

		for _, p := range strings.Split(param, ",") {

			if _, exists := params[p]; exists {
				tmp1 = append(tmp1, params[p]...)
				continue
			}
			if _, exists := m["charSet"][p]; exists {
				chars := strings.Split(m["charSet"][p][0], "")
				tmp1 = append(tmp1, chars...)
				continue
			}
			if _, exists := m["files"][p]; exists {

				content, err := ioutil.ReadFile(m["files"][p][0])

				if err != nil {
					fmt.Println("In cook.yaml, " + m["files"][p][0])
					panic(err)
				}

				fileData := strings.ReplaceAll(string(content), "\r\n", "\n")
				tmp1 = append(tmp1, strings.Split(fileData, "\n")...)
				continue
			}
			if _, exists := m["words"][p]; exists {
				tmp1 = append(tmp1, m["words"][p]...)
				continue
			}
			if _, exists := m["extensions"][p]; exists {
				for _, ext := range m["extensions"][p] {
					tmp1 = append(tmp1, "."+ext)
				}
				continue
			}

			tmp1 = append(tmp1, p)

		}

		tmp2 := []string{}

		if _, exists := colCases[i]; exists {
			A := false
			if colCases[i][0] == "A" {
				A = true
			}

			if A || stringInSlice(colCases[i], "U") {
				for _, t := range final {
					for _, v := range tmp1 {
						tmp2 = append(tmp2, t+strings.ToUpper(v))
					}
				}
			}

			if A || stringInSlice(colCases[i], "L") {
				for _, t := range final {
					for _, v := range tmp1 {
						tmp2 = append(tmp2, t+strings.ToLower(v))
					}
				}
			}

			if A || stringInSlice(colCases[i], "T") {
				for _, t := range final {
					for _, v := range tmp1 {
						tmp2 = append(tmp2, t+strings.Title(v))
					}
				}
			}
		} else {
			for _, t := range final {
				for _, v := range tmp1 {
					tmp2 = append(tmp2, t+v)
				}
			}
		}

		final = tmp2
	}

	if toUpper {
		for _, t := range final {
			fmt.Println(strings.ToUpper(t))
		}
	} else if toLower {
		for _, t := range final {
			fmt.Println(strings.ToLower(t))
		}
	} else {
		for _, t := range final {
			fmt.Println(t)
		}
	}

}
