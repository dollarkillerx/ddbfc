/**
 * @Author: DollarKillerX
 * @Description: 文件操作的一些test
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:50 2019/11/26
 */
package utils

import (
	"fmt"
	"log"
	"testing"
)

func TestReadRowToSet(t *testing.T) {
	strings, e := ReadRowToSet("dns_test.go")
	if e != nil {
		log.Fatalln(e)
	}

	for _, v := range strings {
		fmt.Println(v)
	}
}
