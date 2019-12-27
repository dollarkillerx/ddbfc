/**
 * @Author: DollarKillerX
 * @Description: docker.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午5:21 2019/12/26
 */
package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

// docker横向扩容的模式下获取本地机器地址
func DockerGetHostName() string {
	docker := os.Getenv("inDocker")
	// 不是docker容器下运行
	if docker == "" {
		return os.Args[2]
	}
	// 如果在docker容器下运行获取当前容器的hostname
	bytes, e := ioutil.ReadFile("/etc/hostname")
	if e != nil {
		// 如果不存在 就返回上面的
		return os.Args[2]
	}
	return strings.TrimSpace(string(bytes)) + ":8082"
}
