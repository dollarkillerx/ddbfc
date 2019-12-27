/**
 * @Author: DollarKillerX
 * @Description: task_limit 限流
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:49 2019/12/27
 */
package shared

var LimitChannel = make(chan bool, 1)
