/**
 * @Author: DollarKillerX
 * @Description: spyse_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:40 2019/12/5
 */
package datasource

import (
	"io/ioutil"
	"log"
	"regexp"
	"testing"
)

func TestSpyse_ParseDomainReg(t *testing.T) {
	bytes, e := ioutil.ReadFile("spyse_test.json")
	if e != nil {
		panic(e)
	}
	spysereg("baidu.com", string(bytes))

}

func spysereg(domain, data string) {
	reg := `"{1}(\w+.)*` + domain
	compile := regexp.MustCompile(reg)
	submatch := compile.FindAllStringSubmatch(data, -1)

	for _, v := range submatch {
		log.Println(v[0])
	}
}

func TestSpyse_ParseDomain(t *testing.T) {
	i := spyseNew()
	domains, e := i.ParseDomain("dollarkiller.com")
	//domains, e := i.ParseDomain("baidu.com")
	if e != nil {
		panic(e)
	}
	log.Println(len(domains))
	log.Println(domains)
}
