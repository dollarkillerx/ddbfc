/**
 * @Author: DollarKillerX
 * @Description: dns_dumpster.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:05 2019/12/6
 */
package datasource

import (
	"ddbf/utils"
	"ddbf/utils/httplib"
	"errors"
	"net/http"
	"regexp"
)

type dnsDumpster struct {
	domain1 string
	csrf    string
}

func dnsDumpsterNew() *dnsDumpster {
	return &dnsDumpster{
		domain1: "https://dnsdumpster.com/",
	}
}

func (d *dnsDumpster) ParseDomain(domain string) (Domains, error) {
	err := d.getCsrfToken()
	if err != nil {
		return nil, err
	}

	s, err := d.postFrom(domain)
	if err != nil {
		return nil, err
	}
	return d.decode(domain, s)
}

func (d *dnsDumpster) decode(domain, data string) (Domains, error) {
	Regex, _ := regexp.Compile("<td class=\"col-md-4\">(.*\\..*\\..*)<br>")
	match := Regex.FindAllStringSubmatch(data, -1)
	// String to hold initial subdomains
	var initialSubs []string

	for _, data := range match {
		initialSubs = append(initialSubs, data[1])
	}

	validSubdomains := utils.Validate(domain, initialSubs)
	result := Domains{}
	for _, v := range validSubdomains {
		result[v] = []string{}
	}
	return result, nil
}

func (d *dnsDumpster) postFrom(domain string) (string, error) {
	lib, e := httplib.Post("https://dnsdumpster.com/")
	if e != nil {
		return "", e
	}

	cookie := &http.Cookie{
		Name:   "csrftoken",
		Domain: "dnsdumpster.com",
		Value:  d.csrf,
	}

	s, e := lib.Params("csrfmiddlewaretoken", d.csrf).
		Params("targetip", domain).
		Header("Content-Type", "application/x-www-form-urlencoded").
		Header("Referer", "https://dnsdumpster.com").
		Header("X-CSRF-Token", d.csrf).
		AddCookie(cookie).
		ByteBigString()
	if e != nil {
		return "", e
	}
	return s, nil
}

func (d *dnsDumpster) getCsrfToken() error {
	s, e := get(d.domain1)
	if e != nil {
		return e
	}
	re := regexp.MustCompile("<input type=\"hidden\" name=\"csrfmiddlewaretoken\" value=\"(.*)\">")
	match := re.FindAllStringSubmatch(s, -1)
	if len(match) >= 1 {
		if len(match[0]) == 2 {
			d.csrf = match[0][1]
			return nil
		}
	}
	return errors.New("not csrf token")
}

func init() {
	DataSourcesMu.Lock()
	DataSources = append(DataSources, dnsDumpsterNew())
	DataSourcesMu.Unlock()
}
