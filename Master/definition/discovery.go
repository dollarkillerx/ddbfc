/**
 * @Author: DollarKillerX
 * @Description: discovery 相关的定义
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午1:43 2019/12/27
 */
package definition

import "time"

type Server struct {
	Host string // 服务地址
	Id   string // 服务id

	Load    int64     // 服务负载
	TimeOut time.Time // 服务超时时间
	TryNum  int       // 尝试请求次数
}
