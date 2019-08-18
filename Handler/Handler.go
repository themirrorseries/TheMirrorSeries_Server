package Handler

import (
	"../NetFrame"
	"../proto/dto"
	"net"
)

func Handler(msg []byte, client net.Conn) {
	var decode NetFrame.Decode
	decode.Read(msg)

	//HANDLER CENTER
	switch decode.Thetype {
	case int32(DTO.MsgTypes_TYPE_LOGIN):
		{
			login := NewLogin(decode.Command, decode.ReadPos, decode.Len+4, msg, client)
			login.ReveiveMessage()
			//return login
		}
		break

	case int32(DTO.MsgTypes_TYPE_USER):
		//user
		break
	case int32(DTO.MsgTypes_TYPE_MATCH):
		//match
		match := NewMatch(decode.Command, decode.ReadPos, decode.Len+4, msg, client)
		match.ReveiveMessage()
		break
	case int32(DTO.MsgTypes_TYPE_FIGHT):
		//fight
		//roomManager()
		//Global.ChanMap[decode.Command] <- msg
		fight := NewFight(decode.Command, decode.ReadPos, decode.Len+4, msg, client)
		fight.ReveiveMessage()
		break
	default:
		break
	}
}
