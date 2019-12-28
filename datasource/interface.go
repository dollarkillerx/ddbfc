/**
 * @Author: DollarKillerX
 * @Description: interface.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午8:57 2019/11/29
 */
package datasource

import (
	"ddbf/utils/httplib"
	"sync"
)

// 集中
var DataSourcesMu sync.Mutex
var DataSources = make([]BaseInterface, 0)

// 定义域名存储单元  一个域名 可能对应多个ip
type Domains map[string][]string

type BaseInterface interface {
	ParseDomain(domain string) (Domains, error)
}

func get(url string) (string, error) {
	lib, e := httplib.Get(url)
	if e != nil {
		return "", e
	}
	return lib.ByteBigString()
}
