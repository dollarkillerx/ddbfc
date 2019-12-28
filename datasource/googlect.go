/**
 * @Author: DollarKillerX
 * @Description: googlect.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:49 2019/12/7
 */
package datasource

import (
	"ddbf/utils/httplib"
	"net/url"
)

type googlect struct {
	baseURL string
}

func googlectNew() *googlect {
	return &googlect{
		baseURL: "https://www.google.com/transparencyreport/api/v3/httpsreport/ct/certsearch",
	}
}

func (g *googlect) ParseDomain(domain string) (Domains, error) {
	return nil, nil
}

func (g *googlect) decode(domain, data string) (Domains, error) {
	return nil, nil
}

func (g *googlect) http(url string) (string, error) {
	lib, e := httplib.Get(url)
	if e != nil {
		return "", e
	}
	return lib.Header("Connection", "close").
		Header("Referer", "https://transparencyreport.google.com/https/certificates").ByteBigString()
}

func (g *googlect) getDNSURL(domain, token string) string {
	var dir string

	if token != "" {
		dir = "/page"
	}
	u, _ := url.Parse(g.baseURL + dir)

	values := url.Values{
		"domain":             {domain},
		"include_expired":    {"true"},
		"include_subdomains": {"true"},
	}

	if token != "" {
		values.Add("p", token)
	}

	u.RawQuery = values.Encode()
	return u.String()
}

func init() {
	DataSourcesMu.Lock()
	DataSources = append(DataSources, googlectNew())
	DataSourcesMu.Unlock()
}
