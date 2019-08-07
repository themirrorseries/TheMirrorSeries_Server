package Handler

import "../Global"
// map<roomid, room>
type Room struct{
	roomid		int32	//唯一id
	isfull		bool	//是否满员
	playernum	int32	//人数
	player1id	int32	//人员唯一id
	player2id	int32
	player3id	int32
	player4id	int32
	isStart		bool
}

func (room *Room)Clear(){
	room.roomid		=0
	room.isfull		=false
	room.playernum	=0
	room.player1id	=0
	room.player2id	=0
	room.player3id	=0
	room.player4id	=0
	room.isStart    =false
}

func CopyRoom(cacheRoom *Room) *Room{
	room := &Room{
		roomid:			cacheRoom.roomid,
		isfull:			cacheRoom.isfull,
		playernum:		cacheRoom.playernum,
		player1id:		cacheRoom.player1id,
		player2id:		cacheRoom.player2id,
		player3id:		cacheRoom.player3id,
		player4id:		cacheRoom.player4id,
		isStart:		cacheRoom.isStart,
	}
	return room
}
//游戏初始化  单独通知玩家所属座位编号	光线信息
//其余群发通知	每隔t时间收集玩家操作 广播给所有房间内玩家
//通知消息需要的数据	4个位置playerid	4个socket

func  (room *Room)insertPlayer(){
	insertSuccess := true
	roomFull
	Global.RoomCacheMu.lock()
	if(!room.isfull){
		if(room.player1id==0){
			room.player1id=playerid
		}else if(room.player1id==0)	{
			room.player2id=playerid
		}else if(room.player1id==0)	{
			room.player2id=playerid
		}else if(room.player1id==0)	{
				room.player2id=playerid
		}else
		insertSuccess = false
	}
	if(insertSuccess){
		room.playernum++
	}
	if(playernum==4)
	{
		RoomMng[NextRoomID] = cacheRoom
		NextRoomID++
		cacheRoom.clear()
	}
	Global.RoomCache.unLock()
	if(!RoomMng[NextRoomID-1].isStart)
	
}

func (room *Room)findPlayerByID(){

}

func RoomMatch(){

}

func RoomInform(){

}