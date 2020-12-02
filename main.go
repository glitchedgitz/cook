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

	var merge string
	flag.StringVar(&merge, "m", "", "Merge Different files in single one")

	var subdomainList string
	flag.StringVar(&subdomainList, "f", "", "Enter list of subdomains")

	var types string
	flag.StringVar(&types, "t", "", "Type of extension")

	var extensions string
	flag.StringVar(&extensions, "e", "", "extension or list of extensions")

	flag.Parse()

	// Creating list of everything
	words_list := strings.Split(words, ",")
	prefix_list := strings.Split(prefixs, ",")
	suffix_list := strings.Split(suffixs, ",")
	separators_list := strings.Split(separators, ",")
	extensions_list := strings.Split(extensions, ",")
	types_list := strings.Split(types, ",")

	// Checking for type of extensions if user added
	if types != "" {
		for _, types := range types_list {
			exts := types_dic[types]
			extensions_list = append(extensions_list, exts...)
		}
	}

	if prefixs != "" && suffixs != "" {
		for _, prefix := range prefix_list {
			for _, separator1 := range separators_list {
				for _, word := range words_list {
					for _, separator2 := range separators_list {
						for _, suffix := range suffix_list {
							for _, ext := range extensions_list {
								fmt.Println(prefix + separator1 + word + separator2 + suffix + ext)
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
						fmt.Println(prefix + separator1 + word + ext)
					}
				}
			}
		}
	} else if prefixs == "" && suffixs != "" {
		for _, word := range words_list {
			for _, separator2 := range separators_list {
				for _, suffix := range suffix_list {
					for _, ext := range extensions_list {
						fmt.Println(word + separator2 + suffix + ext)
					}
				}
			}
		}
	} else {
		for _, word := range words_list {
			for _, ext := range extensions_list {
				fmt.Println(word + "." + ext)
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
