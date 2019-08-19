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
		break
	}
}
func (login *Login) clientLogin() {
	//to do  loginDto
	log.Println("client login")
	//解码dto
	any := DTO.UserDTO{}
	proto.Unmarshal(login.data.messages[login.data.bytesStart:login.data.bytesEnd], &any)

	//数据库环境未搭建选此项
	if !IsExist() {
		Global.NextUserIDMu.Lock()
		login.SendLoginMessage(Global.NextUserID)
		Global.NextUserID++
		Global.NextUserIDMu.Unlock()
	} else {
		login.SendLoginMessage(Global.NextUserID)
	}

	//已经有数据库环境选此项 屏蔽上面
	//login.SendLoginMessage(Global.GetUser(Global.UserCollection, any.Uuid))

}

//检查设备号是否存在
func IsExist() bool {
	return false
}

func (login *Login) SendLoginMessage(id int32) {
	any := DTO.UserDTO{}
	any.Id = id
	//any.XXX_Marshal()
	data, _ := proto.Marshal(&any)
	//any.XXX_Marshal()
	log.Println("encode ok")
	buffer := NetFrame.WriteMessage(int32(DTO.MsgTypes_TYPE_LOGIN), int32(DTO.LoginTypes_LOGIN_SRES), data, any.XXX_Size())
	login.data.client.Client.Write(buffer.Bytes())
	log.Println("send ok")
}
