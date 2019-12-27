/**
 * @Author: DollarKillerX
 * @Description: basic_definition 基础的定义
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:04 2019/12/27
 */
package definition

import (
	"github.com/dollarkillerx/easyutils"
	"log"
)

const (
	ZIPFILE = "zipFile"

	UNZIPFILE = "unzipFile"
)

// 初始化目录
func init() {
	err := easyutils.DirPing(ZIPFILE)
	if err != nil {
		log.Fatalln(err, "文件初始化失败")
	}
	err = easyutils.DirPing(UNZIPFILE)
	if err != nil {
		log.Fatalln(err, "文件初始化失败")
	}
}
