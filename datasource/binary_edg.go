/**
 * @Author: DollarKillerX
 * @Description: binary_edg.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午10:06 2019/12/10
 */
package datasource

import (
	"ddbf/utils"
)

type binaryEdg struct {
}

func binaryEdgNew() *binaryEdg {
	return &binaryEdg{}
}

func (b *binaryEdg) ParseDomain(domain string) (Domains, error) {
	url := b.getUrl(domain)
	s, e := get(url)
	if e != nil {
		return nil, e
	}
	return b.decode(domain, s)
}

func (b *binaryEdg) decode(domain, data string) (Domains, error) {
	result := Domains{}
	urls := utils.ExtractSubdomains(data, domain)
	for _, v := range urls {
		result[v] = []string{}
	}
	return result, nil
}

func (b *binaryEdg) getUrl(domain string) string {
	return "https://api.binaryedge.io/v2/query/domains/subdomain/" + domain
}

// 注入到集中处理
//func init() {
//	DataSourcesMu.Lock()
//	DataSources = append(DataSources, binaryEdgNew())
//	DataSourcesMu.Unlock()
//}
