package core

import (
	"net"
	"net/url"
	"path"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func AnalyseURLs(urls []string, get string, array *[]string) {
	get = strings.ToLower(get)
	for _, s := range urls {

		u, err := url.Parse(s)
		if err != nil {
			continue
		}

		if get == "scheme" {
			*array = append(*array, u.Scheme)
			continue
		}
		if get == "user" || get == "username" {
			*array = append(*array, u.User.Username())
			continue
		}

		p, _ := u.User.Password()
		if get == "pass" || get == "password" {
			*array = append(*array, p)
			continue
		}
		if get == "u:p" || get == "user:pass" || get == "username:password" {
			*array = append(*array, u.User.Username()+":"+p)
			continue
		}

		host, port, _ := net.SplitHostPort(u.Host)

		if get == "h" || get == "host" {
			if strings.Contains(u.Host, ":") {
				*array = append(*array, host)
			} else {
				*array = append(*array, u.Host)
			}
			continue
		}

		if get == "port" {
			*array = append(*array, port)
			continue
		}
		if get == "h:p" || get == "host:port" {
			*array = append(*array, host+":"+port)
			continue
		}
		if get == "path" {
			*array = append(*array, u.Path)
			continue
		}
		if get == "f" || get == "fragment" {
			*array = append(*array, u.Fragment)
			continue
		}
		if get == "filepath" {
			_, file := path.Split(s)
			*array = append(*array, file)
			continue
		}
		if get == "q" || get == "query" || get == "k" || get == "key" || get == "keys" {
			*array = append(*array, u.RawQuery)
			continue
		}
		if get == "tld" {
			var domain string
			if strings.Contains(u.Host, ":") {
				domain = host
			} else {
				domain = u.Host
			}
			eTLD, _ := publicsuffix.PublicSuffix(domain)
			*array = append(*array, eTLD)
		}
		if get == "sub" || get == "subdomain" {
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
		if get == "domain" || get == "d" {
			*array = append(*array, u.Scheme+"://"+u.Host)
		}
	}
}
