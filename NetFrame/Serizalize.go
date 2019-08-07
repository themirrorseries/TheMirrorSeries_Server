package NetFrame

import ("fmt"
)

func DeSerizalize(msg []byte){
	var decode Decode
	decode.Read(msg);
	
	fmt.Println(decode.len)
	
	//HANDLER CENTER
	switch(decode.thetype){
	case 0:

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