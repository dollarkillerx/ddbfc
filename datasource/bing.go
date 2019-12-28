/**
 * @Author: DollarKillerX
 * @Description: bing.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午12:07 2019/12/7
 */
package datasource

import (
	"ddbf/utils"
	"net/url"
	"regexp"
	"strconv"
	"sync"
)

type bing struct {
	doMu     sync.Mutex
	domain   Domains
	num      int
	quantity int
}

func bingNew() *bing {
	return &bing{
		domain:   Domains{},
		num:      10,
		quantity: 20,
	}
}

func (b *bing) ParseDomain(domain string) (Domains, error) {
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

	wg.Wait()
	return b.domain, nil
}

func (b *bing) decode(domain, data string) {
	reSub := regexp.MustCompile(`%.{2}`)
	src := reSub.ReplaceAllLiteralString(data, " ")
	doms := utils.ExtractSubdomains(src, domain)
	for _, v := range doms {
		b.doMu.Lock()
		b.domain[v] = []string{}
		b.doMu.Unlock()
	}
}

func (b *bing) urlByPageNum(domain string, page int) string {
	count := strconv.Itoa(b.quantity)
	first := strconv.Itoa((page * b.quantity) + 1)
	u, _ := url.Parse("http://www.bing.com/search")

	u.RawQuery = url.Values{
		"q":     {"domain:" + domain + " -www." + domain},
		"count": {count},
		"first": {first},
		"FORM":  {"PORE"},
	}.Encode()
	return u.String()
}

// 注入到集中处理中
func init() {
	DataSourcesMu.Lock()
	DataSources = append(DataSources, bingNew())
	DataSourcesMu.Unlock()
}
