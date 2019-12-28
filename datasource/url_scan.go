/**
 * @Author: DollarKillerX
 * @Description: url_scan.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午9:32 2019/12/10
 */
package datasource

import (
	"fmt"
)

type urlScan struct {
}

func urlScanNew() *urlScan {
	return &urlScan{}
}

func (u *urlScan) ParseDomain(domain string) (Domains, error) {
	return nil, nil
}

func (u *urlScan) searchURL(domain string) string {
	return fmt.Sprintf("https://urlscan.io/api/v1/search/?q=domain:%s", domain)
}

func (u *urlScan) resultURL(id string) string {
	return fmt.Sprintf("https://urlscan.io/api/v1/result/%s/", id)
}

func (u *urlScan) submitBody(domain string) string {
	return fmt.Sprintf("{\"url\": \"%s\", \"public\": \"on\", \"customagent\": \"%s\"}", domain, "xxx")
}

// 注入到集中处理中
//func init() {
//	DataSourcesMu.Lock()
//	DataSources = append(DataSources, urlScanNew())
//	DataSourcesMu.Unlock()
//}
