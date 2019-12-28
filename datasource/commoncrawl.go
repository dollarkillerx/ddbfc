/**
 * @Author: DollarKillerX
 * @Description: commoncrawl.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午8:09 2019/12/5
 */
package datasource

import (
	"bufio"
	"ddbf/utils"
	"ddbf/utils/httplib"
	"github.com/dollarkillerx/easyutils/clog"
	"log"
	"net/url"
	"strings"
	"sync"
)

//https://index.commoncrawl.org/collinfo.json
type commoncrawl struct {
	url       string
	indexURLs []string

	domMu   sync.Mutex
	domains Domains
}

func commoncrawlNew() *commoncrawl {
	return &commoncrawl{
		url:       "https://index.commoncrawl.org/collinfo.json",
		domains:   Domains{},
		indexURLs: []string{},
	}
}

func (c *commoncrawl) ParseDomain(domain string) (Domains, error) {
	c.initIndexList()
	//wg := sync.WaitGroup{}

	// 并发要报错503   一个ip同时只能访问一个数据源
	for _, index := range c.indexURLs {
		//wg.Add(1)
		//c.decode(&wg, domain, index)
		c.decode(domain, index)
	}
	//wg.Wait()
	return c.domains, nil
}

func (c *commoncrawl) decode(domain string, index string) {
	//defer wg.Done()

	url, data, err := c.getData(domain, index)
	if err != nil {
		clog.PrintWa(err)
		log.Println(url)
		return
	}

	domains := Domains{}
	//for _, v := range ds {
	//	url, i := v.GetString("url")
	//	if !i {
	//		continue
	//	}
	//	index := strings.Index(url, "://")
	//	end := strings.Index(url, domain)
	//	domains[url[index+3:end+len(domain)]] = []string{}
	//}
	reader := strings.NewReader(data)
	scanner := bufio.NewScanner(reader)
	jsonD := utils.JsonDNew()
	for scanner.Scan() {
		err := jsonD.Decode(scanner.Text())
		if err != nil {
			clog.PrintWa(err)
			continue
		}
		url, b := jsonD.GetString("url")
		if !b {
			continue
		}
		index := strings.Index(url, "://")
		end := strings.Index(url, domain)
		domains[url[index+3:end+len(domain)]] = []string{}
	}

	c.domMu.Lock()

	for k, v := range domains {
		c.domains[k] = v
	}

	c.domMu.Unlock()
}

func (c *commoncrawl) getData(domain, index string) (string, string, error) {
	url := c.getURL(domain, index)
	lib, e := httplib.Get(url)
	if e != nil {
		return url, "", e
	}
	s, e := lib.ByteBigString()
	return url, s, e
}

func (c *commoncrawl) getURL(domain, index string) string {
	u, _ := url.Parse(index)

	u.RawQuery = url.Values{
		"url":      {"*." + domain},
		"output":   {"json"},
		"filter":   {"=status:200"},
		"fl":       {"url,status"},
		"pageSize": {"2000"},
	}.Encode()
	return u.String()
}

func (c *commoncrawl) initIndexList() error {
	lib, e := httplib.Get(c.url)
	if e != nil {
		return e
	}
	s, e := lib.ByteBigString()
	if e != nil {
		return e
	}

	jsonD := utils.JsonDNew()
	e = jsonD.Decode(s)
	if e != nil {
		return e
	}
	ds, b := jsonD.GetSliceMap()
	if !b {
		return utils.JsonDParseError
	}
	for _, v := range ds {
		key, i := v.GetString("cdx-api")
		if !i {
			continue
		}
		c.indexURLs = append(c.indexURLs, key)
	}
	return nil
}

// 注入到集中处理中
func init() {
	//DataSourcesMu.Lock()
	//DataSources = append(DataSources, commoncrawlNew())
	//DataSourcesMu.Unlock()
}
