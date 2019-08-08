package Handler

import ("fmt"
"../proto/dto"
"../Global"
	"net"
	"github.com/golang/protobuf/proto"
	"../NetFrame"
	"bytes"
)
type Login struct{
	command int32
	messages []byte
	bytesStart int32
	bytesEnd int32
	client	net.Conn
}

func NewLogin(c, start, end int32, msg []byte, _client net.Conn) *Login{
	login := &Login{
		command:	c,
		bytesStart:	start,
		bytesEnd:	end,
		messages:	msg,
		client:		_client,
	}
	return login
}

func (login *Login)ReveiveMessage(){
	switch (login.command){
	case 0:
		login.clientLogin()
		break;
	case 2:
		//客户端申请注册
		login.clientRegist()
		break;
	default:
		fmt.Println("其他错误")
		break;
	}
}
func (login *Login)clientLogin(){
	//to do  loginDto
	fmt.Println("client login")
	//解码dto
	any :=DTO.UserDTO{}
	//any.XXX_Unmarshal(login.messages[login.bytesStart:login.bytesEnd])
	proto.Unmarshal(login.messages[login.bytesStart:login.bytesEnd], &any)
	fmt.Println("dto ok")
	if(!IsExist()){
		Global.NextUserIDMu.Lock()
		login.SendLoginMessage(Global.NextUserID)
		Global.NextUserID++
		Global.NextUserIDMu.Unlock()
	}else{
		login.SendLoginMessage(Global.NextUserID)
	}

}

func (login *Login)clientRegist(){
	//to do	loginDto
	fmt.Println("client regist")
}

//检查设备号是否存在
func IsExist() bool{
	return false
}

func (login *Login)SendLoginMessage(id int32){
	any :=DTO.UserDTO{}
	any.Id=id
	//any.XXX_Marshal()
	data,_:= proto.Marshal(&any)
	//any.XXX_Marshal()
	fmt.Println("encode ok")
	encode :=NetFrame.NewEncode(int32(8+any.XXX_Size()), 0, 1)
	encode.Write()
	var buffer bytes.Buffer
	buffer.Write(encode.GetBytes())
	buffer.Write(data)
	login.client.Write(buffer.Bytes())
	fmt.Println("send ok")
}