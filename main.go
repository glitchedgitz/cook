package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var Blue = "\033[96m"
var White = "\033[97m"
var gopath = os.Getenv("GOPATH")

// Files
var extensions_txt = gopath + "\\src\\cook\\ingredients\\extensions.txt"
var files_txt = gopath + "\\src\\cook\\ingredients\\files.txt"
var raft_large_extentions_txt = gopath + "\\src\\cook\\ingredients\\raft-large-extensions.txt"

var dirbust_files = `%n%0
%n%1
%n%1%e%
%n%2
%n%2%e%
%n%3
%n%3%e%
%n%4
%n%4%e%
%n%5
%n%5%e%
%n% (3rd copy)%e%
%n% (4th copy)%e%
%n% (another copy)%e%
%n% (copy)%e%
%n% (third copy)%e%
%n% - Copy (2)%e%
%n% - Copy Copy%e%
%n% - Copy%e%
%n%-
%n%.1
%n%.2
%n%.3
%n%.7z
%n%%e%1
%n%%e%2
%n%%e%3
%n%%e%4
%n%%e%5
%n%%e%.1
%n%%e%.7z
%n%%e%.a
%n%%e%.ar
%n%%e%.bac
%n%%e%.backup
%n%%e%.bak
%n%%e%.bz2
%n%%e%.cbz
%n%%e%.exe
%n%%e%.gz
%n%%e%.inc
%n%%e%.include
%n%%e%.jar
%n%%e%.lzma
%n%%e%.old
%n%%e%.rar
%n%%e%.tar
%n%%e%.tar.7z
%n%%e%.tar.bz2
%n%%e%.tar.gz
%n%%e%.tar.lzma
%n%%e%.tar.xz
%n%%e%.wim
%n%%e%.xz
%n%%e%.zip
%n%%e%_backup
%n%%e%_bak
%n%%e%_inc
%n%%e%_old
%n%%e%backup
%n%%e%bak
%n%%e%inc
%n%%e%old
%n%%e%~
%n%.a
%n%.ar
%n%.bac
%n%.backup
%n%.bak
%n%.bz2
%n%.cache
%n%.cbz
%n%.conf
%n%.cs
%n%.csproj
%n%.dif
%n%.dist
%n%.ear
%n%.err
%n%.exe
%n%.gz
%n%.inc
%n%.include
%n%.ini
%n%.jar
%n%.java
%n%.jspa
%n%.jspf
%n%.jspx
%n%.log
%n%.lst
%n%.lzma
%n%.old
%n%.orig
%n%.part
%n%.rar
%n%.rej
%n%.sass-cache
%n%.sav
%n%.save
%n%.save.1
%n%.sublime-project
%n%.sublime-workspace
%n%.swp
%n%.tar
%n%.tar.7z
%n%.tar.bz2
%n%.tar.gz
%n%.tar.lzma
%n%.tar.xz
%n%.temp
%n%.templ
%n%.tmp
%n%.txt
%n%.un~
%n%.vb
%n%.vbproj
%n%.vi
%n%.war
%n%.wim
%n%.xz
%n%.zip
%n%_
%n%_backup
%n%_backup%e%
%n%_bak
%n%_bak%e%
%n%_inc
%n%_old
%n%_old%e%
%n%backup
%n%bak
%n%inc
%n%old
%n%~
%n%~1
%n%~bk
.%n%.swm
.%n%.swn
.%n%.swo
.%n%.swp
backup_%n%%e%
bak_%n%%e%
Copy (2) of %n%%e%
Copy of %n%%e%
Copy of Copy of %n%%e%
old_%n%%e%`

var types_dic = map[string][]string{
	"archive": {".7z", ".a", ".apk", ".xapk", ".ar", ".bz2", ".cab", ".cpio", ".deb", ".dmg", ".egg", ".gz", ".iso", ".jar", ".lha", ".mar", ".pea", ".rar", ".rpm", ".s7z", ".shar", ".tar", ".tbz2", ".tgz", ".tlz", ".war", ".whl", ".xpi", ".zip", ".zipx", ".xz", ".pak"},

	"web": {".css", ".html", ".htm", ".html5", ".less", ".scss", ".wasm"},

	"backup": {".bak", ".backup", ".backup1", ".backup2"},

	"config": {"conf", "config"},

	"code": {".c", ".cc", ".class", ".clj", ".cpp", ".cs", ".cxx", ".el", ".go", ".h", ".java", ".lua", ".m", ".m4", ".php", ".pl", ".po", ".py", ".rb", ".rs", ".swift", ".vb", ".vcxproj", ".xcodeproj", ".xml", ".diff", ".patch", ".html", ".html", ".js"},

	"slide": {".ppt", ".odp"},

	"text": {".doc", ".docx", ".ebook", ".log", ".md", ".msg", ".odt", ".org", ".pages", ".pdf", ".rtf", ".rst", ".tex", ".txt", ".wpd", ".wps"},

	"sheet": {".ods", ".xls", ".xlsx", ".csv", ".ics", ".vcf"},

	"exec": {".exe", ".msi", ".bin", ".command", ".sh", ".bat", ".crx", ".font", ".eot", ".otf", ".ttf", ".woff", ".woff2"},

	"audio": {".aac", ".aiff", ".ape", ".au", ".flac", ".gsm", ".it", ".m3u", ".m4a", ".mid", ".mod", ".mp3", ".mpa", ".pls", ".ra", ".s3m", ".sid", ".wav", ".wma", ".xm"},

	"image": {".3dm", ".3ds", ".max", ".bmp", ".dds", ".gif", ".jpg", ".jpeg", ".png", ".psd", ".xcf", ".tga", ".thm", ".tif", ".tiff", ".ai", ".eps", ".ps", ".svg", ".dwg", ".dxf", ".gpx", ".kml", ".kmz", ".webp"},

	"video": {".3g2", ".3gp", ".aaf", ".asf", ".avchd", ".avi", ".drc", ".flv", ".m2v", ".m4p", ".m4v", ".mkv", ".mng", ".mov", ".mp2", ".mp4", ".mpe", ".mpeg", ".mpg", ".mpv", ".mxf", ".nsv", ".ogg", ".ogv", ".ogm", ".qt", ".rm", ".rmvb", ".roq", ".srt", ".svi", ".vob", ".webm", ".wmv", ".yuv"},

	"book": {".mobi", ".epub", ".azw1", ".azw3", ".azw4", ".azw6", ".azw", ".cbr", ".cbz"},
}

var type_string = `
archive: 7z a apk xapk ar bz2 cab cpio deb dmg egg gz iso jar lha mar pea rar rpm s7z shar tar tbz2 tgz tlz war whl xpi zip zipx xz pak bak
config : conf config
sheet  : ods xls xlsx csv ics vcf
exec   : exe msi bin command sh bat crx
code   : c cc class clj cpp cs cxx el go h java lua m m4 php php3 php5 php7 pl po py rb rs sh swift vb vcxproj xcodeproj xml diff patch js jsx
web    : html html5 htm css js jsx less scss wasm php php3 php5 php7
backup : .bak .backup .backup1 .backup2
slide  : ppt odp
font   : eot otf ttf woff woff2
text   : doc docx ebook log md msg odt org pages pdf rtf rst tex txt wpd wps
audio  : aac aiff ape au flac gsm it m3u m4a mid mod mp3 mpa pls ra s3m sid wav wma xm
book   : mobi epub azw1 azw3 azw4 azw6 azw cbr cbz
video  : 3g2 3gp aaf asf avchd avi drc flv m2v m4p m4v mkv mng mov mp2 mp4 mpe mpeg mpg mpv mxf nsv ogg ogv ogm qt rm rmvb roq srt svi vob webm wmv yuv
image  : 3dm 3ds max bmp dds gif jpg jpeg png psd xcf tga thm tif tiff yuv ai eps ps svg dwg dxf gpx kml kmz webp`

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println(type_string)
	}

	var words string
	flag.StringVar(&words, "w", "", "word or list of words")

	var prefixs string
	flag.StringVar(&prefixs, "p", "", "word or list of words to add as prefix")

	var suffixs string
	flag.StringVar(&suffixs, "s", "", "word or list of words to add as suffix")

	var separators string
	flag.StringVar(&separators, "sp", "", "char/chars or list of char/chars to add as separator")

	var subdomainList string
	flag.StringVar(&subdomainList, "f", "", "Enter list of subdomains")

	var types string
	flag.StringVar(&types, "t", "", "Type of extension")

	var extensions string
	flag.StringVar(&extensions, "e", "", "extension or list of extensions")

	var addslash bool
	flag.BoolVar(&addslash, "addslash", false, "Add '/' splash before every line")

	var dirbust bool
	flag.BoolVar(&dirbust, "dirbust", false, "Using dirbust file")

	var verbose bool
	flag.BoolVar(&verbose, "v", false, "Show verbose")

	flag.Parse()

	// Creating list of everything
	words_list := strings.Split(words, ",")
	prefix_list := strings.Split(prefixs, ",")
	suffix_list := strings.Split(suffixs, ",")
	separators_list := strings.Split(separators, ",")
	extensions_list := strings.Split(extensions, ",")
	types_list := strings.Split(types, ",")

	if verbose {
		fmt.Printf("words_list       : %v\n", words_list)
		fmt.Printf("prefix_list      : %v\n", prefix_list)
		fmt.Printf("suffix_list      : %v\n", suffix_list)
		fmt.Printf("separators_list  : %v\n", separators_list)
		fmt.Printf("extensions_list  : %v\n", extensions_list)
		fmt.Printf("types_list       : %v\n", types_list)
	}
	var slash = ""

	if addslash {
		slash = "/"
	}

	// Checking for type of extensions if user added
	if types != "" {
		for _, types := range types_list {
			exts := types_dic[types]
			extensions_list = append(extensions_list, exts...)
		}
	}

	for _, word := range words_list {
		for _, ext := range extensions_list {
			fmt.Println(slash + word + ext)
		}
	}

	if dirbust {
		for _, word := range words_list {
			for _, ext := range extensions_list {
				if verbose {
					println("Extension ", ext)
				}
				fmt.Println(strings.ReplaceAll(strings.ReplaceAll(dirbust_files, `%n%`, word), `%e%`, ext))
			}
		}
	}

	if prefixs != "" && suffixs != "" {
		for _, prefix := range prefix_list {
			for _, separator1 := range separators_list {
				for _, word := range words_list {
					for _, separator2 := range separators_list {
						for _, suffix := range suffix_list {
							for _, ext := range extensions_list {
								fmt.Println(slash + prefix + separator1 + word + separator2 + suffix + ext)
							}
						}
					}
				}
			}
		}
	} else if prefixs != "" && suffixs == "" {
		for _, prefix := range prefix_list {
			for _, separator1 := range separators_list {
				for _, word := range words_list {
					for _, ext := range extensions_list {
						fmt.Println(slash + prefix + separator1 + word + ext)
					}
				}
			}
		}
	} else if prefixs == "" && suffixs != "" {
		for _, word := range words_list {
			for _, separator2 := range separators_list {
				for _, suffix := range suffix_list {
					for _, ext := range extensions_list {
						fmt.Println(slash + word + separator2 + suffix + ext)
					}
				}
			}
		}
	}

	// founded := []string{}

	// for _, ext := range extensions_list {
	// 	// print(ext)
	// 	data, _ := ioutil.ReadFile(raft_large_extentions_txt)

	// 	r, err := regexp.Compile(ext)
	// 	if err != nil {
	// 		println(err)
	// 	}
	// 	extensions_list := r.FindAllString(string(data), -1)
	// 	// strings.Split(string(data), "\n")
	// 	for _, found := range extensions_list {
	// 		if !stringInSlice(founded, found) {
	// 			fmt.Println(found)
	// 			founded = append(founded, found)
	// 		}
	// 	}
	// }
}
