package test

import (
	"github.com/ipimClient/imServer"
	"log"
	"testing"
)

func Test_Immsg(t *testing.T) {
	immsg := imServer.ImMsg{
		FromUsernames: "xiaoli",
		ToUsernames:   "xiaowang,xiaohuang",
		Appid:         "xxx",
		Content:       "你好",
		Ext:           "", //可有可无
	}
	message, status := immsg.SendMsg("", "")
	log.Println(message, status)
}
