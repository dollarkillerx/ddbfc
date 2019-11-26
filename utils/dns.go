/**
 * @Author: DollarKillerX
 * @Description: dns解析
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:50 2019/11/26
 */
package utils

import (
	"context"
	"net"
	"time"
)

// dns解析测试
// @param 域名
// @param 超时时间
// @param 尝试次数
func DnsParsing(domain string, timeout time.Duration, tryNum int) error {
	var err error
	for i := 0; i < tryNum; i++ {
		wt, _ := context.WithTimeout(context.TODO(), timeout)
		_, err = net.DefaultResolver.LookupHost(wt, domain)
		if err == nil {
			return nil
		}
	}
	return err
}
