package Global

import (
	"../NetFrame"
	"../proto/dto"
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
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
		roomid: 0,
		isfull: false,
		playernum:0,
		players:make([]PlayerList, 2),
	}
	return room
}
func (room *Room) Clear() {
	room.roomid = 0
	room.isfull = false
	room.playernum = 0
	for i:=0;i<len(room.players);i++{
		room.players[i].PlayerID=-1
	}
}

func (room *Room)CopyRoom(cacheRoom *Room) {
	room.roomid = cacheRoom.roomid
	room.isfull = cacheRoom.isfull
	room.playernum = cacheRoom.playernum
	for i:=0;i<len(cacheRoom.players);i++{
		room.players[i] = cacheRoom.players[i]
	}
}

//游戏初始化  单独通知玩家所属座位编号	光线信息
//其余群发通知	每隔t时间收集玩家操作 广播给所有房间内玩家
//通知消息需要的数据	4个位置playerid	4个socket

//多线程执行房间
func RoomRun(room Room) {
	//通知当前房间玩家匹配成功
	//RoomInform(&room)

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
	//roomFull
	RoomCacheMu.Lock()
	fmt.Println("matching")
	insertSuccess := false
	if !room.isfull {
		fmt.Println("isfull")
		for i:=0;i<len(room.players);i++{
			if(room.players[i].PlayerID==-1){
				room.players[i].PlayerID = playerid
				room.players[i].PlayerClient = client
				insertSuccess = true
				fmt.Println("match one")
				break
			}
		}
	}
	fmt.Println(insertSuccess)
	if insertSuccess {
		fmt.Println("被执行")
		room.playernum++
	}
	if room.playernum == 2 {
		fmt.Println("match two people")
		RoomCache.roomid = NextRoomID
		RoomMng[NextRoomID] = NewRoom()
		RoomMng[NextRoomID].CopyRoom(&RoomCache)
		fmt.Println("map ok")
		//ChanMap[NextRoomID] = make(chan []byte)
		NextRoomID++
		RoomCache.Clear()
//		RoomInform(RoomMng[NextRoomID-1])
		RoomMng[NextRoomID-1].RoomInform()
	}
	RoomCacheMu.Unlock()
	fmt.Println("match one success")
}

func (room *Room) findPlayerByID() {

}

//cache room调用
func ReceivePlayer() {

}

func (room *Room)RoomInform(){
	//todo
	fmt.Println("inform")
	match :=DTO.MatchSuccessDTO{}
	match.Roomid = room.roomid
	fmt.Println(len(room.players))
	//match.Players = make([]*DTO.Player,2)
	players1 := new(DTO.Player)
	players2 := new(DTO.Player)
	match.Players = make([]*DTO.Player,2)
	match.Players[0] = players1
	match.Players[1] = players2
	fmt.Println("panic make")
	for i:=0;i<len(room.players);i++{
		match.Players[i].Playerid=room.players[i].PlayerID
		fmt.Println("panic [i]")
		match.Players[i].Name=room.players[i].Name
		match.Players[i].Roleid=room.players[i].PlayerRole
		match.Players[i].Seat=int32(i+1)
	}
	fmt.Println("match ok")
	data, _:=proto.Marshal(&match)
	encode :=NetFrame.NewEncode(int32(8+match.XXX_Size()), 2,4)
	encode.Write()
	var buffer bytes.Buffer
	buffer.Write(encode.GetBytes())
	buffer.Write(data)
	fmt.Println("buffer ok")
	for i:=0;i<len(room.players);i++ {
		room.players[i].PlayerClient.Write(buffer.Bytes())
	}
	fmt.Println("room inform ok")
}
