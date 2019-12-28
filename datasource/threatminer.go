/**
 * @Author: DollarKillerX
 * @Description: threatminer.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午10:15 2019/12/9
 */
package datasource

import (
	"ddbf/utils"
	"fmt"
)

type threatminer struct {
}

func threatminerNew() *threatminer {
	return &threatminer{}
}

func (t *threatminer) ParseDomain(domain string) (Domains, error) {
	url := t.getUrl(domain)
	s, e := get(url)
	if e != nil {
		return nil, e
	}
	return t.decode(domain, s)
}

func (t *threatminer) decode(domain, data string) (Domains, error) {
	doms := utils.ExtractSubdomains(data, domain)
	result := Domains{}
	for _, v := range doms {
		result[v] = []string{}
	}
	return result, nil
}

func (t *threatminer) getUrl(domain string) string {
	return fmt.Sprintf("https://www.threatminer.org/getData.php?e=subdomains_container&q=%s&t=0&rt=10&p=1", domain)
}

func init() {
	DataSourcesMu.Lock()
	DataSources = append(DataSources, threatminerNew())
	DataSourcesMu.Unlock()
}
