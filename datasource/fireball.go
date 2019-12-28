/**
 * @Author: DollarKillerX
 * @Description: fireball.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午11:53 2019/12/9
 */
package datasource

import "ddbf/utils"

type fireball struct {
}

func fireballNew() *fireball {
	return &fireball{}
}

func (f *fireball) ParseDomain(domain string) (Domains, error) {
	url := f.getUrl(domain)
	s, e := get(url)
	if e != nil {
		return nil, e
	}
	return f.decode(domain, s)
}

func (f *fireball) decode(domain, data string) (Domains, error) {
	result := Domains{}
	doms := utils.ExtractSubdomains(data, domain)
	for _, v := range doms {
		result[v] = []string{}
	}
	return result, nil
}

func (f *fireball) getUrl(doamin string) string {
	return "https://fireball.com/search?q=site%3A" + doamin + "+-www.*"
}

// 注入到集中处理
func init() {
	DataSourcesMu.Lock()
	DataSources = append(DataSources, fireballNew())
	DataSourcesMu.Unlock()
}
