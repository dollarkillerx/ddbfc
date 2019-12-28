/**
 * @Author: DollarKillerX
 * @Description: riddler.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午9:56 2019/12/10
 */
package datasource

import (
	"ddbf/utils"
	"fmt"
)

type riddler struct {
}

func riddlerNew() *riddler {
	return &riddler{}
}

func (r *riddler) ParseDomain(domain string) (Domains, error) {
	url := r.getUrl(domain)
	s, e := get(url)
	if e != nil {
		return nil, e
	}
	return r.decode(domain, s)
}

func (r *riddler) decode(domain, data string) (Domains, error) {
	result := Domains{}
	doms := utils.ExtractSubdomains(data, domain)
	for _, v := range doms {
		result[v] = []string{}
	}
	return result, nil
}

func (r *riddler) getUrl(domain string) string {
	return fmt.Sprintf("https://riddler.io/search?q=pld:%s", domain)
}

// 注入到集中处理
func init() {
	DataSourcesMu.Lock()
	DataSources = append(DataSources, riddlerNew())
	DataSourcesMu.Unlock()
}
