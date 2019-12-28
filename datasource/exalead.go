/**
 * @Author: DollarKillerX
 * @Description: exalead.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:27 2019/12/7
 */
package datasource

import (
	"ddbf/utils"
	"fmt"
	"regexp"
)

type exalead struct {
}

func exaleadNew() *exalead {
	return &exalead{}
}

func (e *exalead) ParseDomain(domain string) (Domains, error) {
	url := e.getURL(domain)
	s, err := get(url)
	if err != nil {
		return nil, err
	}
	return e.decode(domain, s)
}

func (e *exalead) decode(domain, data string) (Domains, error) {
	reSub := regexp.MustCompile(`%.{2}`)
	src := reSub.ReplaceAllLiteralString(data, " ")

	match := utils.ExtractSubdomains(src, domain)
	result := Domains{}
	for _, v := range match {
		result[v] = []string{}
	}
	return result, nil
}

func (e *exalead) getURL(domain string) string {
	base := "http://www.exalead.com/search/web/results/"
	format := base + "?q=site:%s+-www?elements_per_page=50"

	return fmt.Sprintf(format, domain)
}

func init() {
	DataSourcesMu.Lock()
	DataSources = append(DataSources, exaleadNew())
	DataSourcesMu.Unlock()
}
