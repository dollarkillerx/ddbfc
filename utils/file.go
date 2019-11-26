/**
 * @Author: DollarKillerX
 * @Description: 文件相关帮助函数
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午4:01 2019/11/26
 */
package utils

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"ddbf/utils/set"
	"encoding/hex"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// 如果文件夹不存在就会创建
func DirPing(path string) error {
	b, e := PathExists(path)
	if e != nil {
		return e
	}
	if !b {
		e := os.MkdirAll(path, 00777)
		if e != nil {
			return e
		}
	}
	return nil
}

// 获取文件后缀
func FileGetPostfix(filename string) (string, error) {
	split := strings.Split(filename, ".")
	if len(split) == 0 {
		return "", errors.New("File Get Postfix Error")
	}
	return split[len(split)-1], nil
}

// 获得文件sha1
func FileGetSha1(file *os.File) string {
	hash := sha1.New()
	io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil))
}

// 获取文件MD5
func FileGetMD5(file *os.File) string {
	_md5 := md5.New()
	io.Copy(_md5, file)
	return hex.EncodeToString(_md5.Sum(nil))
}

// 检测目录下文件是否为空
func FileDirEmpty(file string) bool {
	infos, e := ioutil.ReadDir(file)
	if e != nil {
		return true
	}
	for _, file := range infos {
		if file.IsDir() {
			continue
		}
		return false
	}
	return true
}

// 遍历目录下的文件存放到set中
func LoopDir(dir string) (set.Set, error) {
	out := set.New()
	infos, e := ioutil.ReadDir(dir)
	if e != nil {
		return nil, e
	}
	// 遍历目录下的文件
	for _, file := range infos {
		if file.IsDir() {
			continue
		}
		sets, e := ReadRowToSet(file.Name())
		if e != nil {
			continue
		}
		out.InsertMany(sets...)
	}
	return out, nil
}

// 按行读取文件存入到set中
func ReadRowToSet(file string) ([]string, error) {
	result := make([]string, 0)
	open, e := os.Open(file)
	if e != nil {
		return nil, e
	}
	defer open.Close()
	reader := bufio.NewReader(open)
	for {
		s, e := reader.ReadString('\n')
		if e != nil && e != io.EOF {
			break
		}
		result = append(result, s)
		if e == io.EOF {
			break
		}
	}
	return result, nil
}
