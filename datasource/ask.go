/**
 * @Author: DollarKillerX
 * @Description: ask.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:42 2019/12/6
 */
package datasource

import (
	"ddbf/utils"
	"net/url"
	"strconv"
	"sync"
)

type ask struct {
	doMu   sync.Mutex
	domain Domains
	num    int // 并发数量
}

func askNew() *ask {
	return &ask{
		num:    10,
		domain: Domains{},
	}
}

func (a *ask) ParseDomain(domain string) (Domains, error) {
	wg := sync.WaitGroup{}

	for i := 0; i < a.num; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			url := a.urlByPageNum(domain, i)
			s, e := get(url)
			if e != nil {
				return
			}
			a.decode(domain, s)
		}(i, &wg)
	}

	wg.Wait()
	return a.domain, nil
}

func (a *ask) decode(domain string, data string) {
	doms := utils.ExtractSubdomains(data, domain)
	for _, dom := range doms {
		a.doMu.Lock()
		a.domain[dom] = []string{}
		a.doMu.Unlock()
	}
}

func (a *ask) urlByPageNum(domain string, page int) string {
	p := strconv.Itoa(page)
	u, _ := url.Parse("https://www.ask.com/web")

	u.RawQuery = url.Values{
		"q":    {"site:" + domain + " -www." + domain},
		"o":    {"0"},
		"l":    {"dir"},
		"qo":   {"pagination"},
		"page": {p},
	}.Encode()
	return u.String()
}

// 注入到集中处理中
func init() {
	DataSourcesMu.Lock()
	DataSources = append(DataSources, askNew())
	DataSourcesMu.Unlock()
}
