/**
 * @Author: DollarKillerX
 * @Description: task 相关
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午8:17 2019/12/27
 */
package shared

import (
	"ddbf/pb/pb_master"
)

var Over = make(chan []*pb_master.DomainItem, 1)

//// 消息
//var TaskChannel = make(chan string, 10000)
//
//// 处理完毕的消息
//var OverChannel = make(chan *definition.DomainItem, 5000)
//
//// 当前处理完消息数
//var OverNum int64 = 0
