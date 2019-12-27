/**
 * @Author: DollarKillerX
 * @Description: response 通用相应相关的定义
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:37 2019/12/27
 */
package definition

// 标准相应
type ResponseStandard struct {
	HttpCode int    `json:"-"`
	Code     int    `json:"code"`
	Msg      string `json:"msg,omitempty"`
	Data     string `json:"data,omitempty"`
}

// 相关预定义
var (
	RespOk  = &ResponseStandard{HttpCode: 200, Code: 200, Msg: "200 OK!"}
	Resp400 = &ResponseStandard{HttpCode: 400, Code: 400, Msg: "400 Parameter error"}
	Resp404 = &ResponseStandard{HttpCode: 404, Code: 404, Msg: "404 not file"}

	RespBusy = &ResponseStandard{HttpCode: 200, Code: 2503, Msg: "Server is busy"}

	RespFileErr = &ResponseStandard{HttpCode: 500, Code: 5501, Msg: "File related failure"}
)
