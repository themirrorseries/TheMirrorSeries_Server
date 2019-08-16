package Handler

import (
	"../NetFrame"
	"net"
)

func DeSerizalize(msg []byte, client net.Conn) {
	var decode NetFrame.Decode
	decode.Read(msg)

	//HANDLER CENTER
	switch decode.Thetype {
	case 0:
		{
			login := NewLogin(decode.Command, decode.ReadPos, decode.Len+4, msg, client)
			login.ReveiveMessage()
			//return login
		}
		break

	case 1:
		//user
		break
	case 2:
		//match
		match := NewMatch(decode.Command, decode.ReadPos, decode.Len+4, msg, client)
		match.ReveiveMessage()
		break
	case 3:
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
