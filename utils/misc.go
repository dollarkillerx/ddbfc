/**
 * @Author: DollarKillerX
 * @Description: 相关信息搜集
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:18 2019/12/5
 */
package utils

import (
	"strings"

	"github.com/subfinder/urlx"
)

// ExtractSubdomains extracts a subdomain from a big blob of text
func ExtractSubdomains(text, domain string) (urls []string) {
	allUrls := urlx.ExtractSubdomains(text, domain)

	return Validate(domain, allUrls)
}

//Validate returns valid subdomains found ending with target domain
func Validate(domain string, strslice []string) (subdomains []string) {
	for _, entry := range strslice {
		if strings.HasSuffix(entry, "."+domain) {
			subdomains = append(subdomains, entry)
		}
	}

	return subdomains
}
