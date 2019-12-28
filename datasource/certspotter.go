/**
 * @Author: DollarKillerX
 * @Description: certspotter 这个的分析原理是ssl证书透明性  只要他的证书是机构签发的 我们就能获得他的域名
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:41 2019/12/5
 */
package datasource

import (
	"ddbf/utils"
	"ddbf/utils/httplib"
	"net/url"
)

type certSpotter struct {
}

func certSpotterNew() *certSpotter {
	return &certSpotter{}
}

func (c *certSpotter) ParseDomain(domain string) (Domains, error) {
	url := c.getURL(domain)
	lib, e := httplib.Get(url)
	if e != nil {
		return nil, e
	}
	s, e := lib.String()
	if e != nil {
		return nil, e
	}
	return c.jsonD(s)
}

func (c *certSpotter) jsonD(data string) (Domains, error) {
	// 如果他的证书是共用的证书  就会有一些其他域名跑出来
	// 算了还是加进去
	domains := Domains{}
	jsonD := utils.JsonDNew()
	err := jsonD.Decode(data)
	if err != nil {
		return nil, err
	}
	ds, b := jsonD.GetSliceMap()
	if !b {
		return nil, utils.JsonDParseError
	}
	for _, v := range ds {
		d, i := v.GetSlice("dns_names")
		if !i {
			continue
		}
		strings, i2 := d.TraversalSliceString()
		if !i2 {
			continue
		}
		for _, v := range strings {
			domains[v] = []string{}
		}
	}
	return domains, nil
}

func (c *certSpotter) getURL(domain string) string {
	u, _ := url.Parse("https://api.certspotter.com/v1/issuances")

	u.RawQuery = url.Values{
		"domain":             {domain},
		"include_subdomains": {"true"},
		"match_wildcards":    {"true"},
		"expand":             {"dns_names"},
	}.Encode()
	return u.String()
}

// 注入到集中处理中
func init() {
	DataSourcesMu.Lock()
	DataSources = append(DataSources, certSpotterNew())
	DataSourcesMu.Unlock()
}
