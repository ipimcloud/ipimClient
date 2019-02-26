package imServer

/**
	请求返回信息结构体
 */
type Rst struct {
	Message string `json:"message"` //返回消息
	Status  int    `json:"status"`  //返回状态
}
