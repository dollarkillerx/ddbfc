/**
 * @Author: DollarKillerX
 * @Description: entrust.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午5:59 2019/12/5
 */
package datasource

import (
	"ddbf/utils/httplib"
	"net/url"
	"regexp"
	"strings"
)

type entrust struct {
}

func entrustNew() *entrust {
	return &entrust{}
}

func (e *entrust) ParseDomain(domain string) (Domains, error) {
	url := e.getURL(domain)
	lib, err := httplib.Get(url)
	if err != nil {
		return nil, err
	}
	sc, err := lib.String()
	if err != nil {
		return nil, err
	}
	return e.decode(domain, sc)
}

func (e *entrust) decode(domain, str string) (Domains, error) {
	domains := Domains{}
	reg := `"subjectDN": "(\w+.)*` + domain + `"`
	compile := regexp.MustCompile(reg)
	submatch := compile.FindAllStringSubmatch(str, -1)
	for _, v := range submatch {
		vr := v[0]
		index := strings.Index(vr, "003d")
		if index != -1 {
			domains[vr[index+4:]] = []string{}
		}
	}
	return domains, nil
}

func (e *entrust) getURL(domain string) string {
	u, _ := url.Parse("https://ctsearch.entrust.com/api/v1/certificates")

	u.RawQuery = url.Values{
		"fields":         {"subjectO,issuerDN,subjectDN,signAlg,san,sn,subjectCNReversed,cert"},
		"domain":         {domain},
		"includeExpired": {"true"},
		"exactMatch":     {"false"},
		"limit":          {"5000"},
	}.Encode()
	return u.String()
}

func init() {
	DataSourcesMu.Lock()
	DataSources = append(DataSources, entrustNew())
	DataSourcesMu.Unlock()
}
