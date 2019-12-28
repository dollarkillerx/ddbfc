/**
 * @Author: DollarKillerX
 * @Description: dogpile.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午1:51 2019/12/7
 */
package datasource

// 这个数据源废弃了
import (
	"net/url"
	"strconv"
	"sync"
)

type dogpile struct {
	doMu     sync.Mutex
	domain   Domains
	num      int
	quantity int
}

func dogpileNew() *dogpile {
	return &dogpile{
		quantity: 15,
		domain:   Domains{},
		num:      15,
	}
}

func (d *dogpile) ParseDomain(domain string) (Domains, error) {
	return nil, nil
}

func (d *dogpile) urlByPageNum(domain string, page int) string {
	qsi := strconv.Itoa(d.quantity * page)
	u, _ := url.Parse("http://www.dogpile.com/search/web")

	u.RawQuery = url.Values{"qsi": {qsi}, "q": {domain}}.Encode()
	return u.String()
}

// 注入到集中处理中
//func init() {
//	DataSourcesMu.Lock()
//	DataSources = append(DataSources, dogpileNew())
//	DataSourcesMu.Unlock()
//}
