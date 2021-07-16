package cook

import (
	"log"
	"net"
	"net/url"
	"path"
	"path/filepath"
	"strings"

	"github.com/buger/jsonparser"
	"github.com/giteshnxtlvl/pencode/pkg/pencode"
	"golang.org/x/net/publicsuffix"
)

var leetValues = make(map[string][]string)

func LeetBegin() {
	readInfoYaml(path.Join(ConfigFolder, "leet.yaml"), leetValues)
}

func Leet(values []string, mode int, array *[]string) {
	for _, v := range values {
		var tmp []string
		v2 := v
		for l, ch := range leetValues {
			for _, c := range ch {
				if strings.Contains(v, c) {
					t := strings.ReplaceAll(v, c, l)
					v2 = strings.ReplaceAll(v2, c, l)
					tmp = append(tmp, t)
					if t != v2 {
						tmp = append(tmp, v2)
					}
				}
			}
		}
		if mode == 1 {
			*array = append(*array, tmp...)
		} else if mode == 0 {
			*array = append(*array, tmp[len(tmp)-1])
		}
	}
}

func Cases(values []string, cc []string, array *[]string) {
	var fn func(string) string
	for _, c := range cc {
		c = strings.ToUpper(c)
		if c == "U" {
			fn = strings.ToUpper
		} else if c == "L" {
			fn = strings.ToLower
		} else if c == "T" {
			fn = strings.Title
		} else {
			log.Fatalln("Err: Unknown value for case")
		}

		for _, v := range values {
			*array = append(*array, fn(v))
		}
	}
}

func Regex(values []string, regex string, array *[]string) {
	data := []byte{}
	for _, v := range values {
		data = append(data, []byte(v+"\n")...)
	}
	FindRegex(data, regex, array)
}

func Split(values []string, split string, array *[]string) {
	for _, v := range values {
		*array = append(*array, strings.Split(v, split)...)
	}
}

func SplitIndex(values []string, split string, index int, array *[]string) {
	for _, v := range values {
		vv := strings.Split(v, split)
		if len(vv) >= index+1 {
			*array = append(*array, vv[index])
		}
	}
}

func GetJsonField(lines []string, get []string, array *[]string) {
	for _, line := range lines {
		data := []byte(line)
		value, _, _, _ := jsonparser.Get(data, get...)
		v := string(value)
		*array = append(*array, v)
	}
}

func Encode(lines []string, encodings []string, array *[]string) {
	chain := pencode.NewChain()
	for _, line := range lines {
		err := chain.Initialize(encodings)
		if err != nil {
			log.Fatalln(err)
		}
		output, err := chain.Encode([]byte(line))
		if err != nil {
			log.Fatalln(err)
		}
		*array = append(*array, string(output))
	}
}

func SmartWords(words []string, fn func(string) string, array *[]string) {
	for _, word := range words {
		str := []string{}

		if strings.Contains(word, "_") {
			str = strings.Split(word, "_")

		} else if strings.Contains(word, "-") {
			str = strings.Split(word, "-")

		} else {

			j := 0
			for i, letter := range word {
				if letter > 'A' && letter < 'Z' {
					str = append(str, word[j:i])
					j = i
				}
			}
			str = append(str, word[j:])
		}

		*array = append(*array, str...)
	}
}

func SmartWordsJoin(words []string, joinWith []string, fn func(string) string, array *[]string) {
	for _, word := range words {
		str := []string{}

		if strings.Contains(word, "_") {
			str = strings.Split(word, "_")

		} else if strings.Contains(word, "-") {
			str = strings.Split(word, "-")

		} else {

			j := 0
			for i, letter := range word {
				if letter > 'A' && letter < 'Z' {
					str = append(str, word[j:i])
					j = i
				}
			}
			str = append(str, word[j:])
		}

		for _, join := range joinWith {
			*array = append(*array, strings.Join(str, join))
		}

	}
}

func FileBase(urls []string, array *[]string) {
	for _, u := range urls {
		file := filepath.Base(u)
		*array = append(*array, file)
	}
}

func AnalyzeURLs(urls []string, get string, array *[]string) {
	get = strings.ToLower(get)

	type f func(*url.URL)
	var fn f

	switch get {

	case "s", "scheme":
		fn = func(u *url.URL) { *array = append(*array, u.Scheme) }

	case "u", "user", "username":
		fn = func(u *url.URL) { *array = append(*array, u.User.Username()) }

	case "p", "pass":
		fn = func(u *url.URL) {
			p, _ := u.User.Password()
			*array = append(*array, p)
		}

	case "u:p", "user:pass":
		fn = func(u *url.URL) {
			p, _ := u.User.Password()
			*array = append(*array, u.User.Username()+":"+p)
		}

	case "h", "host":
		fn = func(u *url.URL) {
			host, _, _ := net.SplitHostPort(u.Host)
			if strings.Contains(u.Host, ":") {
				*array = append(*array, host)
			} else {
				*array = append(*array, u.Host)
			}
		}

	case "pr", "port":
		fn = func(u *url.URL) {
			_, port, _ := net.SplitHostPort(u.Host)
			*array = append(*array, port)
		}

	case "h:p", "h:pr", "host:port":
		fn = func(u *url.URL) {
			host, port, _ := net.SplitHostPort(u.Host)
			*array = append(*array, host+":"+port)
		}

	case "ph", "path":
		fn = func(u *url.URL) { *array = append(*array, u.Path) }

	case "f", "fragment":
		fn = func(u *url.URL) { *array = append(*array, u.Fragment) }

	case "q", "query":
		fn = func(u *url.URL) { *array = append(*array, u.RawQuery) }

	case "k", "key", "keys":
		fn = func(u *url.URL) {
			for k := range u.Query() {
				*array = append(*array, k)
			}
		}

	case "v", "values":
		fn = func(u *url.URL) {
			for _, vals := range u.Query() {
				*array = append(*array, vals...)
			}
		}
	case "k:v", "keys:values":
		fn = func(u *url.URL) {
			for k, v := range u.Query() {
				for _, vv := range v {
					*array = append(*array, k+"="+vv)
				}
			}
		}

	case "d", "domain":
		fn = func(u *url.URL) { *array = append(*array, u.Scheme+"://"+u.Host) }

	case "tld":
		fn = func(u *url.URL) {
			host, _, _ := net.SplitHostPort(u.Host)
			var domain string
			if strings.Contains(u.Host, ":") {
				domain = host
			} else {
				domain = u.Host
			}
			eTLD, _ := publicsuffix.PublicSuffix(domain)
			*array = append(*array, eTLD)
		}

	case "sub", "subdomain":
		fn = func(u *url.URL) {

			host, _, _ := net.SplitHostPort(u.Host)

			var domain string
			if strings.Contains(u.Host, ":") {
				domain = host
			} else {
				domain = u.Host
			}
			mainDomain, _ := publicsuffix.EffectiveTLDPlusOne(domain)
			till := len(domain) - len(mainDomain) - 1
			if till < 0 {
				till = 0
			}
			subdomain := domain[:till]
			*array = append(*array, subdomain)
		}

	case "alldir":
		fn = func(u *url.URL) { *array = append(*array, strings.Split(u.Path, "/")...) }

	default:
		return
	}

	for _, s := range urls {

		u, err := url.Parse(s)
		if err != nil {
			VPrint("Err: AnalyseURLs in url " + s)
			continue
		}

		fn(u)
	}
}

func init() {
	log.SetFlags(0)
}
