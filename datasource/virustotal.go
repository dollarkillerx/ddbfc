/**
 * @Author: DollarKillerX
 * @Description: virustotal.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午8:55 2019/11/29
 */
package datasource

import (
	"ddbf/utils"
	"ddbf/utils/httplib"
	"fmt"
)

type virustotal struct {
}

func virustotalNew() *virustotal {
	return &virustotal{}
}

func (v *virustotal) ParseDomain(domain string) (Domains, error) {
	url := v.getURL(domain)

	lib, e := httplib.Get(url)
	if e != nil {
		return nil, e
	}
	s, e := lib.String()
	if e != nil {
		return nil, e
	}
	return v.jsonD(s)
}

func (v *virustotal) getURL(domain string) string {
	format := "https://www.virustotal.com/ui/domains/%s/subdomains?limit=40"

	return fmt.Sprintf(format, domain)
}

func (v *virustotal) jsonD(jsn string) (Domains, error) {
	domains := Domains{}
	jsonD := utils.JsonDNew()
	e := jsonD.Decode(jsn)
	if e != nil {
		return nil, e
	}
	d, b := jsonD.GetSlice("data")
	if !b {
		return nil, utils.JsonDParseError
	}
	ds, b := d.GetSliceMap()
	if !b {
		return nil, utils.JsonDParseError
	}
	for _, v := range ds {
		items := []string{}
		doamin, b := v.GetString("id")
		if !b {
			continue
		}
		gMap, b := v.GetMap("attributes")
		if !b {
			continue
		}
		slice, b := gMap.GetSlice("last_dns_records")
		if !b {
			continue
		}
		jsonDS, b := slice.GetSliceMap()
		if !b {
			continue
		}
		for _, vc := range jsonDS {
			ip, b := vc.GetString("value")
			if !b {
				continue
			}
			items = append(items, ip)
		}
		domains[doamin] = items
	}
	return domains, nil
}

// 注入到集中处理中
func init() {
	DataSourcesMu.Lock()
	DataSources = append(DataSources, virustotalNew())
	DataSourcesMu.Unlock()
}
