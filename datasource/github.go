/**
 * @Author: DollarKillerX
 * @Description: github.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午9:37 2019/12/10
 */
package datasource

import (
	"ddbf/utils"
	"net/url"
	"strconv"
)

type github struct {
}

func githubNew() *github {
	return &github{}
}

func (g *github) ParseDomain(domain string) (Domains, error) {
	dnsurl := g.restDNSURL(domain, 0)
	s, e := get(dnsurl)
	if e != nil {
		return nil, e
	}
	result := Domains{}
	urls := utils.ExtractSubdomains(s, domain)
	for _, v := range urls {
		result[v] = []string{}
	}
	return result, nil
}

func (g *github) restDNSURL(domain string, page int) string {
	pn := strconv.Itoa(page)
	u, _ := url.Parse("https://api.github.com/search/code")

	u.RawQuery = url.Values{
		"s":        {"indexed"},
		"type":     {"Code"},
		"o":        {"desc"},
		"q":        {"\"" + domain + "\""},
		"page":     {pn},
		"per_page": {"100"},
	}.Encode()
	return u.String()
}

// 注入到集中处理
//func init() {
//	DataSourcesMu.Lock()
//	DataSources = append(DataSources, githubNew())
//	DataSourcesMu.Unlock()
//}
