package core

import (
	"net"
	"net/url"
	"path/filepath"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func WordPlay(words []string, joinWith string, fn func(string) string, array *[]string) {

	for _, word := range words {

		str := []string{}
		w := ""

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

		last := len(str) - 1
		if len(str) > 1 {
			for _, s := range str[:last] {
				w += fn(s) + joinWith
			}
		}
		w += fn(str[last])

		*array = append(*array, w)
	}
}

func FilePath(urls []string, array *[]string) {
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
		fn = func(u *url.URL) {
			*array = append(*array, u.Scheme)
		}

	case "u", "user", "username":
		fn = func(u *url.URL) {
			*array = append(*array, u.User.Username())
		}

	case "p", "pass", "password":
		fn = func(u *url.URL) {
			p, _ := u.User.Password()
			*array = append(*array, p)
		}

	case "u:p", "user:pass", "username:password":
		fn = func(u *url.URL) {
			p, _ := u.User.Password()
			*array = append(*array, u.User.Username()+":"+p)
		}

	case "h", "host", "hostname":
		fn = func(u *url.URL) {
			host, _, _ := net.SplitHostPort(u.Host)
			if strings.Contains(u.Host, ":") {
				*array = append(*array, host)
			} else {
				*array = append(*array, u.Host)
			}
		}

	case "port", "pr", "pt":
		fn = func(u *url.URL) {
			_, port, _ := net.SplitHostPort(u.Host)
			*array = append(*array, port)
		}

	case "h:p", "host:port":
		fn = func(u *url.URL) {
			host, port, _ := net.SplitHostPort(u.Host)
			*array = append(*array, host+":"+port)
		}

	case "path":
		fn = func(u *url.URL) {
			*array = append(*array, u.Path)
		}

	case "f", "fragment":
		fn = func(u *url.URL) {
			*array = append(*array, u.Fragment)
		}

	case "q", "query":
		fn = func(u *url.URL) {

			*array = append(*array, u.RawQuery)
		}
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
		fn = func(u *url.URL) {
			*array = append(*array, u.Scheme+"://"+u.Host)
		}

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
		fn = func(u *url.URL) {
			*array = append(*array, strings.Split(u.Path, "/")...)
		}

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
