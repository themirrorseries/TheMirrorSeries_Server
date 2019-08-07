package Handler

import "fmt"
type Login struct{
	command int32
	messages []byte
	bytesStart int32
	bytesEnd int32
}

func NewLogin(c, start, end int32, msg []byte) *Login{
	login := &Login{
		command:	c,
		bytesStart:	start,
		bytesEnd:	end,
		messages:	msg,
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
	
}

func (login *Login)clientRegist(){
	//to do	loginDto
	fmt.Println("client regist")
}