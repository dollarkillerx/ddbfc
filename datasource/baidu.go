/**
 * @Author: DollarKillerX
 * @Description: baidu.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:54 2019/12/6
 */
package datasource

import (
	"ddbf/utils"
	"net/url"
	"strconv"
	"sync"
)

type baidu struct {
	doMu   sync.Mutex
	domain Domains
	num    int // 并发数量
}

func baiduNew() *baidu {
	return &baidu{
		domain: Domains{},
	}
}

func (b *baidu) ParseDomain(domain string) (Domains, error) {
	wg := sync.WaitGroup{}
	for i := 0; i < b.num; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			url := b.urlByPageNum(domain, i)
			s, e := get(url)
			if e != nil {
				return
			}
			b.decode(domain, s)
		}(i, &wg)
	}
	wg.Add(1)
	go b.decodeApi(domain, &wg)

	wg.Wait()

	return b.domain, nil
}

func (b *baidu) decode(domain, data string) {
	doms := utils.ExtractSubdomains(data, domain)
	for _, v := range doms {
		b.doMu.Lock()
		b.domain[v] = []string{}
		b.doMu.Unlock()
	}
}

func (b *baidu) decodeApi(domain string, wg *sync.WaitGroup) {
	defer wg.Done()
	url := b.urlForRelatedSites(domain)
	s, e := get(url)
	if e != nil {
		return
	}
	jsonD := utils.JsonDNew()
	e = jsonD.Decode(s)
	if e != nil {
		return
	}
	d, i := jsonD.GetSliceMap("data")
	if !i {
		panic(3)
		return
	}

	ca := map[string]bool{}
	for _, v := range d {
		domain, i3 := v.GetString("domain")
		if !i3 {
			continue
		}
		ca[domain] = true
	}

	b.doMu.Lock()
	for c, _ := range ca {
		b.domain[c] = []string{}
	}
	b.doMu.Unlock()
}

func (b *baidu) urlByPageNum(domain string, page int) string {
	pn := strconv.Itoa(page)
	u, _ := url.Parse("https://www.baidu.com/s")
	query := "site:" + domain + " -site:www." + domain

	u.RawQuery = url.Values{
		"pn": {pn},
		"wd": {query},
		"oq": {query},
	}.Encode()
	return u.String()
}

func (b *baidu) urlForRelatedSites(domain string) string {
	u, _ := url.Parse("https://ce.baidu.com/index/getRelatedSites")

	u.RawQuery = url.Values{"site_address": {domain}}.Encode()
	return u.String()
}

// 注入到集中处理中
func init() {
	DataSourcesMu.Lock()
	DataSources = append(DataSources, baiduNew())
	DataSourcesMu.Unlock()
}
