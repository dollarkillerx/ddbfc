/**
 * @Author: DollarKillerX
 * @Description: crtsh.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午11:24 2019/12/6
 */
package datasource

import "ddbf/utils"

type crtsh struct {
}

func crtshNew() *crtsh {
	return &crtsh{}
}

func (c *crtsh) ParseDomain(domain string) (Domains, error) {
	url := c.getURL(domain)
	s, e := get(url)
	if e != nil {
		return nil, e
	}
	return c.decode(s)
}

func (c *crtsh) decode(data string) (Domains, error) {
	result := Domains{}
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
		s, i := v.GetString("name_value")
		if !i {
			continue
		}
		result[s] = []string{}
	}
	return result, nil
}

func (c *crtsh) getURL(domain string) string {
	return "https://crt.sh/?q=%25." + domain + "&output=json"
}

func init() {
	DataSourcesMu.Lock()
	DataSources = append(DataSources, crtshNew())
	DataSourcesMu.Unlock()
}
