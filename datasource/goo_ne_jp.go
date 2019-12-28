/**
 * @Author: DollarKillerX
 * @Description: goo_ne_jp.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:26 2019/12/9
 */
package datasource

import "ddbf/utils"

type gooNeJp struct {
}

func gooNeJpNew() *gooNeJp {
	return &gooNeJp{}
}

func (g *gooNeJp) ParseDomain(domain string) (Domains, error) {
	url := g.getUrl(domain)
	s, e := get(url)
	if e != nil {
		return nil, e
	}
	return g.decode(domain, s)
}

func (g *gooNeJp) decode(domain, data string) (Domains, error) {
	doms := utils.ExtractSubdomains(data, domain)
	result := Domains{}
	for _, v := range doms {
		result[v] = []string{}
	}
	return result, nil
}

func (g *gooNeJp) getUrl(domain string) string {
	return "https://search.goo.ne.jp/web.jsp?mode=0&PT=TOP&sbd=goo001&OE=UTF-8&MT=site%3A" + domain + "+-www.*&from=pager_web&IE=UTF-8"
	//return "https://search.goo.ne.jp/web.jsp?mode=0&PT=TOP&sbd=goo001&OE=UTF-8&MT=site%3Abaidu.com+-www.*&from=pager_web&IE=UTF-8"
}

// 注入到集中处理中
func init() {
	DataSourcesMu.Lock()
	DataSources = append(DataSources, gooNeJpNew())
	DataSourcesMu.Unlock()
}
