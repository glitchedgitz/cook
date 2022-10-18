package methods

import (
	"log"
	"net"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/glitchedgitz/cook/v2/pkg/cook"
	"golang.org/x/net/publicsuffix"
)

func FileBase(urls []string, useless string, array *[]string) {
	for _, u := range urls {
		file := filepath.Base(u)
		*array = append(*array, file)
	}
}

func UrlScheme(u *url.URL, array *[]string) {
	*array = append(*array, u.Scheme)
}
func UrlUser(u *url.URL, array *[]string) {
	*array = append(*array, u.User.Username())
}
func UrlPass(u *url.URL, array *[]string) {
	p, _ := u.User.Password()
	*array = append(*array, p)
}
func UrlHost(u *url.URL, array *[]string) {
	host, _, _ := net.SplitHostPort(u.Host)
	if strings.Contains(u.Host, ":") {
		*array = append(*array, host)
	} else {
		*array = append(*array, u.Host)
	}
}
func UrlPort(u *url.URL, array *[]string) {
	_, port, _ := net.SplitHostPort(u.Host)
	*array = append(*array, port)
}
func UrlPath(u *url.URL, array *[]string) {
	*array = append(*array, u.Path)
}
func UrlFrag(u *url.URL, array *[]string) {
	*array = append(*array, u.Fragment)
}
func UrlRawQuery(u *url.URL, array *[]string) {
	*array = append(*array, u.RawQuery)
}
func UrlKey(u *url.URL, array *[]string) {
	for k := range u.Query() {
		*array = append(*array, k)
	}
}
func UrlValue(u *url.URL, array *[]string) {
	for _, vals := range u.Query() {
		*array = append(*array, vals...)
	}
}
func UrlDomain(u *url.URL, array *[]string) {
	*array = append(*array, u.Scheme+"://"+u.Host)
}
func UrlTld(u *url.URL, array *[]string) {
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
func UrlSub(u *url.URL, array *[]string) {
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
func UrlAllSub(u *url.URL, array *[]string) {
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
func UrlAllDir(u *url.URL, array *[]string) {
	*array = append(*array, strings.Split(u.Path, "/")...)
}

func AnalyzeURLs(urls []string, fn func(*url.URL, *[]string), array *[]string) {

	for _, s := range urls {
		u, err := url.Parse(s)
		if err != nil {
			cook.VPrint("Err: AnalyseURLs in url " + s)
			continue
		}

		fn(u, array)
	}
}

func init() {
	log.SetFlags(0)
}
