/**
 * @Author: DollarKillerX
 * @Description: yandex.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午10:56 2019/12/9
 */
package datasource

import (
	"ddbf/utils"
	"strconv"
	"sync"
)

type yandex struct {
	domMu   sync.Mutex
	domains Domains
	num     int
}

func yandexNew() *yandex {
	return &yandex{
		domains: Domains{},
		num:     10,
	}
}

func (y *yandex) ParseDomain(domain string) (Domains, error) {
	wg := sync.WaitGroup{}

	for i := 0; i < y.num; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			url := y.getUrl(domain, i)
			s, e := get(url)
			if e != nil {
				return
			}
			y.decode(domain, s)
		}(i, &wg)
	}

	wg.Wait()
	return nil, nil
}

func (y *yandex) decode(domain, data string) {
	doms := utils.ExtractSubdomains(data, domain)
	for _, v := range doms {
		y.domMu.Lock()
		y.domains[v] = []string{}
		y.domMu.Unlock()
	}

}

func (y *yandex) getUrl(domain string, page int) string {
	return "https://yandex.com/search/?text=site%3A" + domain + "%20-www.*&lr=10619&p=" + strconv.Itoa(page)
}

// 注入到集中处理中
func init() {
	DataSourcesMu.Lock()
	DataSources = append(DataSources, yandexNew())
	DataSourcesMu.Unlock()
}
