/**
 * @Author: DollarKillerX
 * @Description: 简单的http请求包
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:26 2019/11/26
 */
package utils

import (
	"io/ioutil"
	"net/http"
)

type EasyHttp struct {
}

func EasyHttpNew() *EasyHttp {
	return &EasyHttp{}
}

// 简单的get请求
func (e *EasyHttp) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
