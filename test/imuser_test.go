package test

import (
	"fmt"
	"github.com/ipimClient/imServer"
	"testing"
)

/**
	如何注册app
 */
func Test_Imuser(t *testing.T) {
	imUser := imServer.ImUser{
		Id:         1,
		Username:   "1111",
		Nickname:   "12",
		Password:   "hahahahaha",
		Appid:      "xxx",
		Logo:       "123123",
		LogoSource: "1231111",
	}
	message, status := imUser.RegisterUser("", "")
	fmt.Println(message, status)
}
