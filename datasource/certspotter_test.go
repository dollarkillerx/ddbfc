/**
 * @Author: DollarKillerX
 * @Description: certspotter_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:00 2019/12/5
 */
package datasource

import (
	"ddbf/utils"
	"io/ioutil"
	"log"
	"testing"
)

func TestJsonDecode(t *testing.T) {
	bytes, e := ioutil.ReadFile("certspotter_test.json")
	if e != nil {
		panic(e)
	}
	domains, e := jsonD(string(bytes))
	if e != nil {
		panic(e)
	}
	log.Println(domains)

}

func jsonD(data string) (Domains, error) {
	// 如果他的证书是共用的证书  就会有一些其他域名跑出来
	// 算了还是加进去
	domains := Domains{}
	jsonD := utils.JsonDNew()
	err := jsonD.Decode(data)
	if err != nil {
		return nil, err
	}
	ds, b := jsonD.GetSliceMap()
	if !b {
		return nil, utils.JsonDParseError
	}
	for _, v := range ds {
		d, i := v.GetSlice("dns_names")
		if !i {
			continue
		}
		strings, i2 := d.TraversalSliceString()
		if !i2 {
			continue
		}
		for _, v := range strings {
			domains[v] = []string{}
		}
	}
	return domains, nil
}
