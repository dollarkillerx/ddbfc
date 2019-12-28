/**
 * @Author: DollarKillerX
 * @Description: 模拟真实用户访问
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:55 2019/12/4
 */
package httplib

import (
	"ddbf/utils"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type httpLib struct {
	url    string            // 目标地址
	typ    bool              // 本次传输类型 typ: true get , false post
	params url.Values        // 参数
	header map[string]string // header
	cookie *http.Cookie      // cookie
}

func Get(url string) (*httpLib, error) {
	ur, e := urlParse(url)
	if e != nil {
		return nil, e
	}
	return &httpLib{
		url:    ur,
		typ:    true,
		header: map[string]string{},
		params: map[string][]string{},
	}, nil
}

func Post(url string) (*httpLib, error) {
	return &httpLib{
		url:    url,
		typ:    false,
		header: map[string]string{},
		params: map[string][]string{},
	}, nil
}

func (h *httpLib) Header(k, v string) *httpLib {
	h.header[k] = v
	return h
}

func (h *httpLib) Params(k string, v string) *httpLib {
	h.params.Set(k, v)
	return h
}

func (h *httpLib) AddCookie(cookie *http.Cookie) *httpLib {
	h.cookie = cookie
	return h
}

func (h *httpLib) Resp() (*http.Response, error) {
	httpClient := &http.Client{}
	var req *http.Request
	var err error
	if h.typ {
		// get
		if len(h.params) == 0 {
			h.params = nil
		}
		req, err = http.NewRequest("GET", h.url, strings.NewReader(h.params.Encode()))
		if err != nil {
			return nil, err
		}
	} else {
		// post
		req, err = http.NewRequest("POST", h.url, strings.NewReader(h.params.Encode()))
		if err != nil {
			return nil, err
		}
	}
	h.setRequest(req)
	h.addCookie(req)
	return httpClient.Do(req)
}

func (h *httpLib) ByteBig() ([]byte, error) {
	resp, e := h.Resp()
	if e != nil {
		return nil, e
	}
	defer resp.Body.Close()
	code := resp.StatusCode
	if code != 200 {
		return nil, errors.New("not 200 Code: " + strconv.Itoa(code))
	}
	return utils.ReadBigFile(resp.Body)
}

func (h *httpLib) ByteBigString() (string, error) {
	bytes, e := h.ByteBig()
	if e != nil {
		return "", e
	}
	return string(bytes), nil
}

func (h *httpLib) Byte() ([]byte, error) {
	response, e := h.Resp()
	if e != nil {
		return nil, e
	}
	defer response.Body.Close()
	code := response.StatusCode
	if code != 200 {
		return nil, errors.New("not 200 Code: " + strconv.Itoa(code))
	}
	return ioutil.ReadAll(response.Body)
}

func (h *httpLib) String() (string, error) {
	bytes, e := h.Byte()
	if e != nil {
		return "", e
	}
	return string(bytes), nil
}

func (h *httpLib) setRequest(req *http.Request) {
	req.Header.Set("User-Agent", getRandUA())
	req.Header.Set("Accept", accept)
	req.Header.Set("Accept-Language", acceptLang)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// 设置用户设置的header
	for k, v := range h.header {
		req.Header.Set(k, v)
	}
}

func (h *httpLib) addCookie(req *http.Request) {
	if h.cookie != nil {
		req.AddCookie(h.cookie)
	}
}

func urlParse(ur string) (string, error) {
	index := strings.Index(ur, "?")

	if index != -1 {
		values, err := url.Parse(ur)
		if err != nil {
			return "", err
		}
		return ur[:index+1] + values.Query().Encode(), nil
	}

	return ur, nil
}
