package Handler

import ("fmt"
"../Global"
)

type Match struct{
	command int32
	messages []byte
	bytesStart int32
	bytesEnd int32
}

func NewMatch(c, start, end int32, msg []byte) *Match{
	match := &Match{
		command:	c,
		bytesStart:	start,
		bytesEnd:	end,
		messages:	msg,
	}
	return match
}

func (match *Match)ReveiveMessage(){
	switch (match.command){
	case 0:
		//申请进入匹配
		match.matchStart()
		break;
	case 2:
		//申请离开匹配
		match.matchEnd()
		break;
	default:
		fmt.Println("其他错误")
		break;
	}
}
func (match *Match)matchStart(){
	//to do
	fmt.Println("match start")
	//add to match pool
	isMatched :=false
	switch(Global.RoomCache.playernum){
		case 0:
			//playid1
			break
		case 1:
			//playid2
			break 
		case 2:
			//playerid3
			break;
		case 3:
			//playerid4
			//new map(cache) insert(newmap) cache.clear
			//inform
			break;
		default:
			break;
	}
}

func (match *Match)matchEnd(){
	//to do
	fmt.Println("match end")
	//delete from cache room

}

//enter(user token)