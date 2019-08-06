package NetFrame

import "fmt"
func Encodeaa(thetype int32, command int32){
	switch(thetype){
	case 0:
		
	}
}

func DeSerizalize(msg []byte){
	var decode Decode
	decode.Read(msg);
	fmt.Println(decode.len)
	switch(decode.thetype){
	case 0:
		//login;
		break;
	
	case 1:
		//user
		break;
	case 2:
		//match
		break;
	case 3:
		//fight
		break;
	default:
		break
	}
}