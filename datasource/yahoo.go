/**
 * @Author: DollarKillerX
 * @Description: yahoo.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:04 2019/12/5
 */
package datasource

import (
	"ddbf/utils"
	"ddbf/utils/httplib"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

type yahoo struct {
	doMu   sync.Mutex
	domain Domains
	num    int // 并发数
}

func yahooNew() *yahoo {
	return &yahoo{
		num:    10,
		domain: Domains{},
	}
}

/**
这边系统的实现方式 就是定义一个最大的取值阶段   不是太友好  先就按照他原来的实现方式来实现把
*/

func (y *yahoo) ParseDomain(domain string) (Domains, error) {
	wg := sync.WaitGroup{}

	for i := 0; i < y.num; i++ {
		wg.Add(1)
		go func(i int, wg *sync.WaitGroup) {
			defer wg.Done()
			url := y.getURL(domain, i)
			s, e := y.httpGet(url)
			if e != nil {
				return
			}
			y.decode(domain, s)
		}(i, &wg)
	}
	wg.Wait()
	return y.domain, nil
}

func (y *yahoo) decode(domain, data string) {
	compile := regexp.MustCompile(`%.{2}`)
	src := compile.ReplaceAllLiteralString(data, " ")
	match := utils.ExtractSubdomains(src, domain)

	xy1 := "-domain:www." + domain + " " + "3Awww." + domain
	for _, subdomain := range match {
		y.doMu.Lock()
		if strings.Index(xy1, subdomain) == -1 {
			y.domain[subdomain] = []string{}
		}
		y.doMu.Unlock()
	}
}

func (y *yahoo) httpGet(url string) (string, error) {
	lib, e := httplib.Get(url)
	if e != nil {
		return "", e
	}
	return lib.String()
}

// 拼接请求url
func (y *yahoo) getURL(domain string, page int) string {
	b := strconv.Itoa(10*page + 1)
	pz := strconv.Itoa(10)

	u, _ := url.Parse("https://search.yahoo.com/search")
	u.RawQuery = url.Values{
		"p":     {"site:" + domain + " -domain:www." + domain},
		"b":     {b},
		"pz":    {pz},
		"bct":   {"0"},
		"xargs": {"0"},
	}.Encode()
	return u.String()
}

// 注入到集中处理中
func init() {
	DataSourcesMu.Lock()
	DataSources = append(DataSources, yahooNew())
	DataSourcesMu.Unlock()
}
