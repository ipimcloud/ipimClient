package imServer

import (
	"encoding/json"
	"github.com/ipimClient/util"
	"log"
)

/**
	发送消息结构体
	发送者或者接受者如果是多个用户,请用","隔开
	如果发送者为多个用户,那么消息发送会平均分配到多个发送者上
	(如果要发送给100个人消息,发送者有4个,则每个发送者会发送25条消息)
 */
type ImMsg struct {
	FromUsernames string //发送消息用户   必须
	ToUsernames   string //接受消息用户	必须
	Content       string //消息内容		必须
	Ext           string //分享内容 (一般用不到这个参数,如果只是单纯发消息,请保持此为空)
	Appid         string //消息所属appid  必须
}

/**
	发送消息
	@param url 请求ip或者网址
	@param sign 官网申请的秘钥
	@return string  返回的消息
	@return int  	返回的状态 0为正常,-1为失败
 */
func (msg *ImMsg) SendMsg(url string, sign string) (string, int) {
	//参数验证
	if msg.Appid == "" {
		return "消息发送appid不能为空", -1
	}
	if msg.FromUsernames == "" {
		return "消息发送者不能为空", -1
	}
	if msg.ToUsernames == "" {
		return "消息接受者不能为空", -1
	}
	if msg.Content == "" {
		return "内容必须", -1
	}
	h := util.NewHttpSend("http://" + url + "/v1/immsg/sendMsgToMult")
	param := make(map[string]string)
	param["from_usernames"] = msg.FromUsernames
	param["to_usernames"] = msg.ToUsernames
	param["content"] = msg.Content
	param["ext"] = msg.Ext
	param["appid"] = msg.Appid
	//签名验证
	now, signature := util.GetSignature(sign)
	param["time"] = now
	param["signature"] = signature
	h.SetBody(param)
	response, err := h.Post()
	if err != nil {
		log.Println("请求失败:", err)
		return "发送请求失败", -1
	}
	rst := &Rst{}
	err = json.Unmarshal(response, &rst)
	if err != nil {
		log.Println("请求错误:", err)
		return "请求失败", -1
	} else {
		return rst.Message, rst.Status
	}
}
