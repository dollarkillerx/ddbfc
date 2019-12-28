/**
 * @Author: DollarKillerX
 * @Description: google.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午5:34 2019/12/7
 */
package datasource

import (
	"ddbf/utils"
	"net/url"
	"strconv"
	"sync"
)

type google struct {
	dnMu   sync.Mutex
	domain Domains

	quantity int
	num      int
}

func googleNew() *google {
	return &google{
		domain:   Domains{},
		quantity: 15,
		num:      10,
	}
}

func (g *google) ParseDomain(domain string) (Domains, error) {
	wg := sync.WaitGroup{}

	for i := 0; i < g.num; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			url := g.urlByPageNum(domain, i, 0)
			s, e := get(url)
			if e != nil {
				return
			}
			g.decode(domain, s)
		}(i, &wg)
	}
	wg.Wait()
	return g.domain, nil
}

func (g *google) decode(domain, data string) {
	doms := utils.ExtractSubdomains(data, domain)
	for _, v := range doms {
		g.dnMu.Lock()
		g.domain[v] = []string{}
		g.dnMu.Unlock()
	}
}

func (g *google) urlByPageNum(domain string, page, numwilds int) string {
	start := strconv.Itoa(g.quantity * page)
	u, _ := url.Parse("https://www.google.com/search")

	var wilds string
	for i := 0; i < numwilds; i++ {
		wilds = "*." + wilds
	}

	u.RawQuery = url.Values{
		"q":      {"site:" + wilds + domain + " -www.*"},
		"btnG":   {"Search"},
		"hl":     {"en"},
		"biw":    {""},
		"bih":    {""},
		"gbv":    {"1"},
		"start":  {start},
		"filter": {"0"},
	}.Encode()
	return u.String()
}

// 注入到集中处理中
func init() {
	DataSourcesMu.Lock()
	DataSources = append(DataSources, googlectNew())
	DataSourcesMu.Unlock()
}
