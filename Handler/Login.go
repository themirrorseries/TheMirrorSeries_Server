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
	aesDec, err := NetFrame.AesDecrypt(login.data.messages[login.data.bytesStart:login.data.bytesEnd], NetFrame.MyKey)
	if err != nil {
		Global.ErrorLog.Log.Errorln(err, "客户端登录解码出错！")
	}
	any := DTO.UserDTO{}
	proto.Unmarshal(aesDec, &any)

	if Global.UseMongo {
		login.sendLoginMessage(Global.GetUser(Global.UserCollection, any.Uuid))
	} else {
		Global.NextUserIDMu.Lock()
		login.sendLoginMessage(Global.NextUserID)
		Global.NextUserID++
		Global.NextUserIDMu.Unlock()
	}
}

func (login *Login) sendLoginMessage(id int32) {
	any := DTO.UserDTO{}
	any.Id = id
	data, _ := proto.Marshal(&any)
	buffer, err := NetFrame.WriteMessage(int32(DTO.MsgTypes_TYPE_LOGIN), int32(DTO.LoginTypes_LOGIN_SRES), data, any.XXX_Size())
	if err != nil {
		Global.ErrorLog.Log.Errorln(err, "发送客户端登录消息编码出错！")
	}
	login.data.client.Client.Write(buffer.Bytes())
	Global.DetailedLog.Detailed(id, Global.Login_IN)
}
