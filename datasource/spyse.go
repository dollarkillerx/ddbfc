/**
 * @Author: DollarKillerX
 * @Description: spyse.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:40 2019/12/5
 */
package datasource

import (
	"ddbf/utils/httplib"
	"fmt"
	"regexp"
)

type spyse struct {
}

func spyseNew() *spyse {
	return &spyse{}
}

func (s *spyse) ParseDomain(domain string) (Domains, error) {
	url := s.getUrl(domain)
	lib, e := httplib.Get(url)
	if e != nil {
		return nil, e
	}
	sc, e := lib.String()
	if e != nil {
		return nil, e
	}
	return s.reg(domain, sc)
}

func (s *spyse) reg(domain, data string) (Domains, error) {
	domains := Domains{}
	reg := `"{1}(\w+.)*` + domain
	compile := regexp.MustCompile(reg)
	submatch := compile.FindAllStringSubmatch(data, -1)
	for _, v := range submatch {
		if len(v) == 2 {
			c := v[0]
			if len(c) > 1 {
				domains[c[1:]] = []string{}
			}
		}
	}
	return domains, nil
}

func (s *spyse) getUrl(domain string) string {
	return fmt.Sprintf("https://api.spyse.com/v1/subdomains-aggregate?domain=%s", domain)
}

// 注入到集中处理中
func init() {
	DataSourcesMu.Lock()
	DataSources = append(DataSources, spyseNew())
	DataSourcesMu.Unlock()
}
