/**
 * @Author: DollarKillerX
 * @Description: virustotal.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午8:55 2019/11/29
 */
package datasource

import "fmt"

type Virustotal struct {
}

func (v *Virustotal) New() *Virustotal {
	return &Virustotal{}
}

func (v *Virustotal) ParseDomain(domain string) DomainList {
	url := v.getURL(domain)
	url = url
	return nil
}

func (v *Virustotal) getURL(domain string) string {
	format := "https://www.virustotal.com/ui/domains/%s/subdomains?limit=40"

	return fmt.Sprintf(format, domain)
}
