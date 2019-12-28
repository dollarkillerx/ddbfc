/**
 * @Author: DollarKillerX
 * @Description: discovery_test.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午2:15 2019/12/4
 */
package server

import "testing"

func TestDiscoveryNew(t *testing.T) {
	dic := DiscoveryNew()
	//dic.Run("dollarkiller.com")
	dic.Run("baidu.com")

}
