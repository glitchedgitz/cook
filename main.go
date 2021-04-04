package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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
 ▄████▄   ▒█████   ▒█████   ██ ▄█▀
▒██▀ ▀█  ▒██▒  ██▒▒██▒  ██▒ ██▄█▒ 
▒▓█    ▄ ▒██░  ██▒▒██░  ██▒▓███▄░ 
▒▓▓▄ ▄██▒▒██   ██░▒██   ██░▓██ █▄ 
 ▒ ▓███▀ ░░ ████▓▒░░ ████▓▒░▒██▒ █▄ V1
======================================

HIGHLY CUSTOMIZABLE WORDLIST GENERATOR
	by Gitesh Sharma @giteshnxtlvl
`

var help = `
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
` + config

var config = `
# This is COOK's config file

charSet:
    n     : [0123456789]
    A     : [ABCDEFGHIJKLMNOPQRSTUVWXYZ]
    a     : [abcdefghijklmnopqrstuvwxyz]
    aAN   : [abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789]
    AN    : [ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789]
    an    : [abcdefghijklmnopqrstuvwxyz0123456789]
    aA    : [abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ]
    s     : ["!#$%&'()*+,-./:;<=>?@[\\]^_` + "`" + `{|}~&\""]
    all   : ["!#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_` + "`" + `abcdefghijklmnopqrstuvwxyz{|}~\""]

words:
    admin: [admin, root, su]
    files: [masters, files, password]

extensions:
    archive: [.7z, .a, .apk, .xapk, .ar, .bz2, .cab, .cpio, .deb, .dmg, .egg, .gz, .iso, .jar, .lha, .mar, .pea, .rar, .rpm, .s7z, .shar, .tar, .tbz2, .tgz, .tlz, .war, .whl, .xpi, .zip, .zipx, .xz, .pak]
    config : [.conf, .config]
    sheet  : [.ods, .xls, .xlsx, .csv, .ics .vcf]
    exec   : [.exe, .msi, .bin, .command, .sh, .bat, .crx]
    code   : [.c, .cc, .class, .clj, .cpp, .cs, .cxx, .el, .go, .h, .java, .lua, .m, .m4, .php, .php3, .php5, .php7, .pl, .po, .py, .rb, .rs, .sh, .swift, .vb, .vcxproj, .xcodeproj, .xml, .diff, .patch, .js, .jsx]
    web    : [.html, .html5, .htm, .css, .js, .jsx, .less, .scss, .wasm, .php, .php3, .php5, .php7]
    backup : [.bak, .backup, .backup1, .backup2]
    slide  : [.ppt, .odp]
    font   : [.eot, .otf, .ttf, .woff, .woff2]
    text   : [.doc, .docx, .ebook, .log, .md, .msg, .odt, .org, .pages, .pdf, .rtf, .rst, .tex, .txt, .wpd, .wps]
    audio  : [.aac, .aiff, .ape, .au, .flac, .gsm, .it, .m3u, .m4a, .mid, .mod, .mp3, .mpa, .pls, .ra, .s3m, .sid, .wav, .wma, .xm]
    book   : [.mobi, .epub, .azw1, .azw3, .azw4, .azw6, .azw, .cbr, .cbz]
    video  : [.3g2, .3gp, .aaf, .asf, .avchd, .avi, .drc, .flv, .m2v, .m4p, .m4v, .mkv, .mng, .mov, .mp2, .mp4, .mpe, .mpeg, .mpg, .mpv, .mxf, .nsv, .ogg, .ogv, .ogm, .qt, .rm, .rmvb, .roq, .srt, .svi, .vob, .webm, .wmv, .yuv]
    image  : [.3dm, .3ds, .max, .bmp, .dds, .gif, .jpg, .jpeg, .png, .psd, .xcf, .tga, .thm, .tif, .tiff, .yuv, .ai, .eps, .ps, .svg, .dwg, .dxf, .gpx, .kml, .kmz, .webp]
`

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

//Parse Input
func parseInput(commands []string) {

	if len(commands) == 0 {
		fmt.Println(banner)
		os.Exit(0)
	}
	if stringInSlice(commands, "-h") {
		fmt.Println(help)
		os.Exit(0)
	}

	last := len(commands) - 1

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
					fileData := string(content)
					fileData = strings.ReplaceAll(fileData, "\r\n", "\n")
					values = strings.Split(fileData, "\n")
				}
			} else {
				values = strings.Split(value, ",")
			}

			params[cmd] = values
		}
	}
	pattern = strings.Split(commands[last], ":")
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

	for k, v := range m["charSet"] {
		m["charSet"][k] = strings.Split(v[0], "")
	}

}

func main() {

	cookConfig()
	parseInput(os.Args[1:])

	final := []string{}

	for i, param := range pattern {
		tmp1 := []string{}

		for _, p := range strings.Split(param, ",") {

			if _, exists := params[p]; exists {
				tmp1 = append(tmp1, params[p]...)
				continue
			}

			var notFound = true
			for _, char := range m {
				if _, exists := char[p]; exists {
					tmp1 = append(tmp1, char[p]...)
					notFound = false
					break
				}
			}

			if notFound {
				tmp1 = append(tmp1, p)
			}
		}

		if i == 0 {
			final = tmp1
			continue
		}

		tmp2 := []string{}
		for _, t := range final {
			for _, v := range tmp1 {
				tmp2 = append(tmp2, t+v)
			}
		}

		final = tmp2
	}

	for _, t := range final {
		fmt.Println(t)
	}

}
