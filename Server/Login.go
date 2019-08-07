package main

type Login struct{
	command int32
	bytesStart int32
	bytesEnd int32
	messages []byte
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

func (login *Login)reveiveMessage(){
	switch (login.command){
	case 0:
		//login	客户端申请登录
		break;
	case 2:
		//客户端申请注册
		break;
	default:
		break;
	}
}
func (login *Login)login(command int32, messages []byte){
	switch command{
	case 0:
		//login	客户端申请登录
		break;
	case 2:
		//客户端申请注册
		break;
	default:
		break;
	}
}
