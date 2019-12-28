/**
 * @Author: DollarKillerX
 * @Description: virustotal_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:44 2019/12/3
 */
package datasource

import (
	"ddbf/utils"
	"io/ioutil"
	"log"
	"testing"
)

//func TestVirustotalNew(t *testing.T) {
//	vir := VirustotalNew()
//	lists, e := vir.ParseDomain("dollarkiller.com")
//	if e != nil {
//		panic(e)
//	}
//	log.Println(lists)
//}

func TestVirustotalJson(t *testing.T) {
	bytes, e := ioutil.ReadFile("virustotal_test.json")
	if e != nil {
		log.Fatalln(e)
	}
	//dac := virustotalJson{}

	//i := make(map[string]interface{})
	//e = json.Unmarshal(bytes, &i)
	//if e != nil {
	//	log.Fatalln(e)
	//}
	//
	////log.Println(dac)
	//ac := i["data"].([]interface{})[0].(map[string]interface{})
	//
	//i2 := ac["attributes"].(map[string]interface{})
	//i3 := i2["last_dns_records"].([]interface{})[0].(map[string]interface{})
	//log.Println(ac["id"])
	//log.Println(i3["value"])

	jsonD := utils.JsonDNew()
	e = jsonD.Decode(string(bytes))
	if e != nil {
		panic(e)
	}
	d, b := jsonD.GetSlice("data")
	if !b {
		log.Panic("err get slice")
	}
	ds, b := d.GetSliceMap()
	if !b {
		log.Panic("err get slice map")
	}

	for _, v := range ds {
		get, i4 := v.Get("id")
		if i4 {
			log.Println(get)
		}
		//"last_dns_records"
		gMap, e := v.GetMap("attributes")
		if !e {
			log.Panic("err get gMap2")
		}
		slice, e := gMap.GetSlice("last_dns_records")
		if !e {
			log.Panic("err get slice2")
		}
		jsonDS, i2 := slice.GetSliceMap()
		if !i2 {
			log.Panic("err get slice map2")
		}
		for _, v := range jsonDS {
			i3, i5 := v.Get("value")
			if i5 {
				log.Println(i3)
			}
		}
	}
}

func TestParseDomain(t *testing.T) {
	vir := virustotalNew()
	domains, e := vir.ParseDomain("dollarkiller.com")
	//domains, e := vir.ParseDomain("baidu.com")
	//domains, e := vir.ParseDomain("baidu.com")
	if e != nil {
		panic(e)
	}
	for k, v := range domains {
		log.Println(k, " : ", v)
	}
}
