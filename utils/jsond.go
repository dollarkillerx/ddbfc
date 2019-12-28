/**
 * @Author: DollarKillerX
 * @Description: jsond.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午7:15 2019/12/3
 */
package utils

import (
	"encoding/json"
	"errors"
)

var JsonDParseError = errors.New("JsonDParseError")

type JsonD struct {
	jMap   map[string]interface{}
	jSlice []interface{}
	typ    bool // true jMap false JSlice
}

func JsonDNew() *JsonD {
	return &JsonD{
		jMap: map[string]interface{}{},
	}
}

func (j *JsonD) Decode(str string) error {
	byt := []byte(str)
	if string(str[0]) == "[" {
		err := json.Unmarshal(byt, &j.jSlice)
		if err != nil {
			return err
		}
		j.typ = false
		return nil
	}
	j.typ = true
	err := json.Unmarshal(byt, &j.jMap)
	return err
}

func (j *JsonD) GetString(key string) (string, bool) {
	if j.typ {
		str, bool := j.jMap[key].(string)
		return str, bool
	}
	return "", false
}

func (j *JsonD) Get(key string) (interface{}, bool) {
	if j.typ {
		inter, bool := j.jMap[key]
		return inter, bool
	}
	return nil, false
}

func (j *JsonD) GetMap(key string) (*JsonD, bool) {
	if j.typ {
		jIntr, bool := j.jMap[key]
		if !bool {
			return nil, false
		}
		jMap, bool := jIntr.(map[string]interface{})
		if !bool {
			return nil, false
		}
		return &JsonD{jMap: jMap, typ: true}, true
	}
	return nil, false
}

func (j *JsonD) GetSlice(key string) (*JsonD, bool) {
	if j.typ {
		jIntr, bool := j.jMap[key]
		if !bool {
			return nil, false
		}
		jSlice, bool := jIntr.([]interface{})
		if !bool {
			return nil, false
		}
		return &JsonD{
			jSlice: jSlice,
			typ:    false,
		}, true
	}
	return nil, false
}

func (j *JsonD) GetSliceMap(key ...string) ([]*JsonD, bool) {
	le := len(key)
	if j.typ && le == 1 {
		jIntr, bool := j.jMap[key[0]]
		if !bool {
			return nil, false
		}
		jSlice, bool := jIntr.([]interface{})
		if !bool {
			return nil, false
		}
		items := []*JsonD{}
		for _, v := range jSlice {
			i, bool := v.(map[string]interface{})
			if bool {
				item := &JsonD{jMap: i, typ: true}
				items = append(items, item)
			}
		}
		return items, true
	} else if le == 0 {
		items := []*JsonD{}
		for _, v := range j.jSlice {
			i, bool := v.(map[string]interface{})
			if bool {
				item := &JsonD{jMap: i, typ: true}
				items = append(items, item)
			}
		}
		return items, true
	}
	return nil, false
}

func (j *JsonD) TraversalSlice() ([]interface{}, bool) {
	if !j.typ {
		return j.jSlice, true
	}
	return nil, false
}

func (j *JsonD) TraversalSliceString() ([]string, bool) {
	if !j.typ {
		result := []string{}
		for _, v := range j.jSlice {
			s, ok := v.(string)
			if !ok {
				return nil, false
			}
			result = append(result, s)
		}
		return result, true
	}
	return nil, false
}
