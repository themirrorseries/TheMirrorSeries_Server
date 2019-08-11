package Handler

import("fmt"
	"net"
	"../NetFrame"
)


func DeSerizalize(msg []byte, client net.Conn){
	var decode NetFrame.Decode
	decode.Read(msg);
	
	fmt.Println(decode.Len)
	
	//HANDLER CENTER
	switch(decode.Thetype){
	case 0:{
			login := NewLogin(decode.Command, decode.ReadPos, decode.Len+4, msg, client)
			login.ReveiveMessage()
			//return login
	}
		break;
	
	case 1:
		//user
		break;
	case 2:
		//match
		match:=NewMatch(decode.Command, decode.ReadPos, decode.Len+4, msg, client)
		match.ReveiveMessage()
		break;
	case 3:
		//fight
		//roomManager()
		//Global.ChanMap[decode.Command] <- msg
		//fmt.Println("fight")
		fight:=NewFight(decode.Command, decode.ReadPos, decode.Len+4, msg, client)
		fight.ReveiveMessage()
		break;
	default:
		break
	}
}
