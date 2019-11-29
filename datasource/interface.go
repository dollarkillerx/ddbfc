/**
 * @Author: DollarKillerX
 * @Description: interface.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午8:57 2019/11/29
 */
package datasource

type DomainList []string

type BaseInterface interface {
	ParseDomain(domain string) DomainList
}
