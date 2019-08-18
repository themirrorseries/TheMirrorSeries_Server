package Handler

import (
	"../Global"
	"../NetFrame"
	"../proto/dto"
	"bytes"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"net"
)

type Login struct {
	command    int32
	messages   []byte
	bytesStart int32
	bytesEnd   int32
	client     net.Conn
}

func NewLogin(c, start, end int32, msg []byte, _client net.Conn) *Login {
	login := &Login{
		command:    c,
		bytesStart: start,
		bytesEnd:   end,
		messages:   msg,
		client:     _client,
	}
	return login
}

func (login *Login) ReveiveMessage() {
	switch login.command {
	case int32(DTO.LoginTypes_LOGIN_CREQ):
		login.clientLogin()
		break
	//case 2:
	//客户端申请注册
	//	login.clientRegist()
	//	break
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
	//any.XXX_Unmarshal(login.messages[login.bytesStart:login.bytesEnd])
	proto.Unmarshal(login.messages[login.bytesStart:login.bytesEnd], &any)
	//log.Println(any.Id)
	//log.Println("dto ok")

	if !IsExist() {
		Global.NextUserIDMu.Lock()
		login.SendLoginMessage(Global.NextUserID)
		Global.NextUserID++
		Global.NextUserIDMu.Unlock()
	} else {
		login.SendLoginMessage(Global.NextUserID)
	}

	//login.SendLoginMessage(Global.GetUser(Global.UserCollection, any.Uuid))

}

func (login *Login) clientRegist() {
	//to do	loginDto
	log.Println("client regist")
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
	encode := NetFrame.NewEncode(int32(8+any.XXX_Size()), int32(DTO.MsgTypes_TYPE_LOGIN), int32(DTO.LoginTypes_LOGIN_SRES))
	encode.Write()
	var buffer bytes.Buffer
	buffer.Write(encode.GetBytes())
	buffer.Write(data)
	login.client.Write(buffer.Bytes())
	log.Println("send ok")
}
