package NetFrame

import("fmt"
"../Handler"
	)


func DeSerizalize(msg []byte){
	var decode Decode
	decode.Read(msg);
	
	fmt.Println(decode.len)
	
	//HANDLER CENTER
	switch(decode.thetype){
	case 0:{
			login := Handler.NewLogin(decode.command, decode.readPos, decode.len, msg)
			login.ReveiveMessage()
			//return login
	}
		break;
	
	case 1:
		//user
		break;
	case 2:
		//match
		match:=Handler.NewMatch(decode.command, decode.readPos, decode.len, msg)
		match.ReveiveMessage()
		break;
	case 3:
		//fight
		//roomManager()
		break;
	default:
		break
	}
}