package Global

import "time"

//import "../Global"
// map<roomid, room>
type Room struct {
	roomid    int32 //唯一id
	isfull    bool  //是否满员
	playernum int32 //人数
	player1id int32 //人员唯一id
	player2id int32
	player3id int32
	player4id int32
}

func (room *Room) Clear() {
	room.roomid = 0
	room.isfull = false
	room.playernum = 0
	room.player1id = 0
	room.player2id = 0
	room.player3id = 0
	room.player4id = 0
}

func CopyRoom(cacheRoom *Room) *Room {
	room := &Room{
		roomid:    cacheRoom.roomid,
		isfull:    cacheRoom.isfull,
		playernum: cacheRoom.playernum,
		player1id: cacheRoom.player1id,
		player2id: cacheRoom.player2id,
		player3id: cacheRoom.player3id,
		player4id: cacheRoom.player4id,
	}
	return room
}

//游戏初始化  单独通知玩家所属座位编号	光线信息
//其余群发通知	每隔t时间收集玩家操作 广播给所有房间内玩家
//通知消息需要的数据	4个位置playerid	4个socket

//多线程执行房间
func RoomRun(room Room) {
	//通知当前房间玩家匹配成功
	RoomInform(&room)

	//创建定时器
	ticker := time.NewTicker(20 * time.Microsecond)

	for{
		if(false){
			break
		}
		<-ticker.C
		//ChanMap[room.roomid]
	}

}

//cache room调用
func (room *Room) InsertPlayer(playerid int32) {
	insertSuccess := true
	//roomFull
	RoomCacheMu.Lock()
	if !room.isfull {
		if room.player1id == 0 {
			room.player1id = playerid
		} else if room.player1id == 0 {
			room.player2id = playerid
		} else if room.player1id == 0 {
			room.player2id = playerid
		} else if room.player1id == 0 {
			room.player2id = playerid
		} else {
			insertSuccess = false
		}
	}
	if insertSuccess {
		room.playernum++
	}
	if room.playernum == 4 {

		RoomMng[NextRoomID] = RoomCache
		ChanMap[NextRoomID] = make(chan []byte)
		NextRoomID++
		RoomCache.Clear()
		go RoomRun(RoomMng[NextRoomID-1])
	}
	RoomCacheMu.Unlock()
}

func (room *Room) findPlayerByID() {

}

//cache room调用
func ReceivePlayer() {

}

func RoomInform(room *Room) {
	//todo
}
