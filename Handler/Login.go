package Handler

import (
	"../Global"
	"../NetFrame"
	"../proto/dto"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

type Login struct {
	data *HandlerData
}

func (login *Login) ReceiveMessage() {
	switch login.data.command {
	case int32(DTO.LoginTypes_LOGIN_CREQ):
		login.clientLogin()
		break
	default:
		log.Println("其他错误")
		Global.DetailedLog.Log.Warn("客户端登录请求出错！")
		break
	}
}

//检查是否存有设备id，根据设备id发送玩家唯一id号
func (login *Login) clientLogin() {
	log.Println("client login")
	any := DTO.UserDTO{}
	proto.Unmarshal(login.data.messages[login.data.bytesStart:login.data.bytesEnd], &any)
	//数据库环境未搭建选此项
	/*
		if !false() {
			Global.NextUserIDMu.Lock()
			login.SendLoginMessage(Global.NextUserID)
			Global.NextUserID++
			Global.NextUserIDMu.Unlock()
		} else {
			login.SendLoginMessage(Global.NextUserID)
		}
	*/

	//已经有数据库环境选此项 屏蔽上面
	login.sendLoginMessage(Global.GetUser(Global.UserCollection, any.Uuid))
}

func (login *Login) sendLoginMessage(id int32) {
	any := DTO.UserDTO{}
	any.Id = id
	data, _ := proto.Marshal(&any)
	buffer := NetFrame.WriteMessage(int32(DTO.MsgTypes_TYPE_LOGIN), int32(DTO.LoginTypes_LOGIN_SRES), data, any.XXX_Size())
	login.data.client.Client.Write(buffer.Bytes())
	Global.DetailedLog.Detailed(id, Global.Login_IN)
}
