package imServer

import (
	"encoding/json"
	"github.com/ipimClient/util"
	"log"
	"strconv"
)

/**
	聊天用户结构体
 */
type ImUser struct {
	Id         int64  //注册用户的Id(数据库用户的表id)  必须
	Username   string //用户名(用户编号)				 必须
	Nickname   string //用户昵称					    非必须
	Password   string //用户密码						 必须
	LogoSource string //头像原图						 必须
	Logo       string //头像缩略图					 必须
	Mobile     string //手机号						非必须
	Appid      string //用户的appid				     必须
}

/**
	注册用户
	@param url 请求ip或者网址
	@param sign 官网申请的秘钥
	@return string  返回的消息
	@return int  	返回的状态 0为正常,-1为失败
 */
func (user *ImUser) RegisterUser(url string, sign string) (string, int) {
	//参数判断
	if user.Id == 0 {
		return "数据库用户表id不能为0", -1
	}
	if user.Username == "" {
		return "用户编号不能为空", -1
	}
	if user.Password == "" {
		return "用户密码不能为空", -1
	}
	if user.LogoSource == "" {
		return "用户头像原图不能为空", -1
	}
	if user.Logo == "" {
		return "用户头像缩略图不能为空", -1
	}
	if user.Appid == "" {
		return "用户所属appid不能为空", -1
	}
	h := util.NewHttpSend("http://" + url + "/v1/imuser/createuser")
	param := make(map[string]string)
	param["id"] = strconv.FormatInt(user.Id, 10)
	param["username"] = user.Username
	param["nickname"] = user.Nickname
	param["password"] = user.Password
	param["logo_source"] = user.LogoSource
	param["logo"] = user.Logo
	param["mobile"] = user.Mobile
	param["appid"] = user.Appid
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
