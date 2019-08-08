package Global

import (
	"net"
	"../proto/dto"
	"../NetFrame"
	"github.com/golang/protobuf/proto"
	"bytes"
)

//import "../Global"
// map<roomid, room>
type PlayerList struct{
	PlayerID int32
	Name string
	PlayerClient net.Conn
	PlayerRole int32
}

type Room struct {
	roomid    	int32 //唯一id
	isfull    	bool  //是否满员
	playernum 	int32 //人数
	players		[]PlayerList
}

func NewRoom()*Room{
	room :=&Room{
		roomid:		0,
		isfull:false,
		playernum:0,
		players:make([]PlayerList, 2),	//测试暂定2个人
	}
	return room
}
func (room *Room) Clear() {
	room.roomid = 0
	room.isfull = false
	room.playernum = 0
	for i:=0;i<len(room.players);i++{
		room.players[i].PlayerID=0
	}
}

func (room *Room)CopyRoom(cacheRoom *Room) {
	room.roomid = cacheRoom.roomid
	room.isfull = cacheRoom.isfull
	room.playernum = cacheRoom.playernum
}

//游戏初始化  单独通知玩家所属座位编号	光线信息
//其余群发通知	每隔t时间收集玩家操作 广播给所有房间内玩家
//通知消息需要的数据	4个位置playerid	4个socket

//多线程执行房间
func RoomRun(room Room) {
	//通知当前房间玩家匹配成功
	RoomInform(&room)

	//创建定时器
	/*ticker := time.NewTicker(5 * time.Microsecond)

	for{
		if(false){
			break
		}
		<-ticker.C
		//ChanMap[room.roomid]
	}*/

}

//cache room调用
func (room *Room) InsertPlayer(playerid int32, client net.Conn) {
	insertSuccess := false
	//roomFull
	RoomCacheMu.Lock()
	if !room.isfull {
		for _,player := range room.players{
			if(player.PlayerID==0){
				player.PlayerID = playerid
				player.PlayerClient = client
				insertSuccess = true
				break
			}
		}
	}
	if insertSuccess {
		room.playernum++
	}
	if room.playernum == 2 {
		RoomCache.roomid = NextRoomID
		RoomMng[NextRoomID] = NewRoom()
		ChanMap[NextRoomID] = make(chan []byte)
		NextRoomID++
		RoomCache.Clear()
		//go RoomRun(*RoomMng[NextRoomID-1])
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
	match :=DTO.MatchSuccessDTO{}
	match.Roomid = room.roomid
	for i:=0;i<len(room.players);i++{
		match.Players[i].Playerid=room.players[i].PlayerID
		match.Players[i].Name=room.players[i].Name
		match.Players[i].Roleid=room.players[i].PlayerRole
		match.Players[i].Seat=int32(i+1)
	}
	data, _:=proto.Marshal(&match)
	encode :=NetFrame.NewEncode(int32(8+match.XXX_Size()), 2,4)
	var buffer bytes.Buffer
	buffer.Write(encode.GetBytes())
	buffer.Write(data)
	for i:=0;i<len(room.players);i++ {
		room.players[i].PlayerClient.Write(buffer.Bytes())
	}
}
