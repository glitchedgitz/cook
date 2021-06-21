package core

import (
	"net"
	"net/url"
	"path/filepath"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func AnalyzeURLs(urls []string, get string, array *[]string) {
	get = strings.ToLower(get)

	for _, s := range urls {

		u, err := url.Parse(s)
		if err != nil {
			VPrint("Err: AnalyseURLs in url " + s)
			continue
		}

		switch get {

		case "s", "scheme":
			*array = append(*array, u.Scheme)

		case "u", "user", "username":
			*array = append(*array, u.User.Username())

		case "p", "pass", "password":
			p, _ := u.User.Password()
			*array = append(*array, p)

		case "u:p", "user:pass", "username:password":
			p, _ := u.User.Password()
			*array = append(*array, u.User.Username()+":"+p)

		case "h", "host", "hostname":
			host, _, _ := net.SplitHostPort(u.Host)
			if strings.Contains(u.Host, ":") {
				*array = append(*array, host)
			} else {
				*array = append(*array, u.Host)
			}

		case "port", "pr", "pt":
			_, port, _ := net.SplitHostPort(u.Host)
			*array = append(*array, port)

		case "h:p", "host:port":
			host, port, _ := net.SplitHostPort(u.Host)
			*array = append(*array, host+":"+port)

		case "path":
			*array = append(*array, u.Path)

		case "f", "fragment":
			*array = append(*array, u.Fragment)

		case "filepath", "fp", "fb", "filebase":
			file := filepath.Base(s)
			*array = append(*array, file)

		case "q", "query":
			*array = append(*array, u.RawQuery)
		case "k", "key", "keys":
			for k := range u.Query() {
				*array = append(*array, k)
			}

		case "v", "values":
			for _, vals := range u.Query() {
				*array = append(*array, vals...)
			}

		case "d", "domain":
			*array = append(*array, u.Scheme+"://"+u.Host)

		case "tld":
			host, _, _ := net.SplitHostPort(u.Host)
			var domain string
			if strings.Contains(u.Host, ":") {
				domain = host
			} else {
				domain = u.Host
			}
			eTLD, _ := publicsuffix.PublicSuffix(domain)
			*array = append(*array, eTLD)

		case "sub", "subdomain":
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

		case "alldir":
			*array = append(*array, strings.Split(u.Path, "/")...)
		}
	}
}
