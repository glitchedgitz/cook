package methods

import (
	"log"
	"net"
	"net/url"
	"path/filepath"
	"strings"

	"golang.org/x/net/publicsuffix"
)

func (m *Methods) FileBase(urls []string, useless string, array *[]string) {
	for _, u := range urls {
		file := filepath.Base(u)
		*array = append(*array, file)
	}
}

func (m *Methods) UrlScheme(u *url.URL, array *[]string) {
	*array = append(*array, u.Scheme)
}
func (m *Methods) UrlUser(u *url.URL, array *[]string) {
	*array = append(*array, u.User.Username())
}
func (m *Methods) UrlPass(u *url.URL, array *[]string) {
	p, _ := u.User.Password()
	*array = append(*array, p)
}
func (m *Methods) UrlHost(u *url.URL, array *[]string) {
	host, _, _ := net.SplitHostPort(u.Host)
	if strings.Contains(u.Host, ":") {
		*array = append(*array, host)
	} else {
		*array = append(*array, u.Host)
	}
}
func (m *Methods) UrlPort(u *url.URL, array *[]string) {
	_, port, _ := net.SplitHostPort(u.Host)
	*array = append(*array, port)
}
func (m *Methods) UrlPath(u *url.URL, array *[]string) {
	*array = append(*array, u.Path)
}
func (m *Methods) UrlFrag(u *url.URL, array *[]string) {
	*array = append(*array, u.Fragment)
}
func (m *Methods) UrlRawQuery(u *url.URL, array *[]string) {
	*array = append(*array, u.RawQuery)
}
func (m *Methods) UrlKey(u *url.URL, array *[]string) {
	for k := range u.Query() {
		*array = append(*array, k)
	}
}
func (m *Methods) UrlValue(u *url.URL, array *[]string) {
	for _, vals := range u.Query() {
		*array = append(*array, vals...)
	}
}
func (m *Methods) UrlDomain(u *url.URL, array *[]string) {
	*array = append(*array, u.Scheme+"://"+u.Host)
}
func (m *Methods) UrlTld(u *url.URL, array *[]string) {
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
func (m *Methods) UrlSub(u *url.URL, array *[]string) {
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
func (m *Methods) UrlAllSub(u *url.URL, array *[]string) {
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
	*array = append(*array, strings.Split(subdomain, ".")...)
}
func (m *Methods) UrlAllDir(u *url.URL, array *[]string) {
	*array = append(*array, strings.Split(u.Path, "/")...)
}

func (m *Methods) AnalyzeURLs(urls []string, fn func(*url.URL, *[]string), array *[]string) {

	for _, s := range urls {
		u, err := url.Parse(s)
		if err != nil {
			log.Println("Err: AnalyseURLs in url " + s)
			continue
		}

		fn(u, array)
	}
}

func (m *Methods) init() {
	log.SetFlags(0)
}
