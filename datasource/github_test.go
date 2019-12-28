/**
 * @Author: DollarKillerX
 * @Description: github_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 上午9:37 2019/12/10
 */
package datasource

import (
	"log"
	"testing"
)

func TestGithub_ParseDomainGetUrl(t *testing.T) {
	g := githubNew()
	dnsurl := g.restDNSURL("baidu.com", 0)
	log.Println(dnsurl)
}
