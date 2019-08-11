package Global

import (
	"../NetFrame"
	"../Tools"
	"../proto/dto"
	"bytes"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
)

//import "../Global"
// map<roomid, room>
type PlayerList struct {
	PlayerID     int32
	Name         string
	PlayerClient net.Conn
	PlayerRole   int32
}

type Room struct {
	roomid            int32 //唯一id
	isfull            bool  //是否满员
	playernum         int32 //人数
	players           []PlayerList
	CurClientFrameNum int32
	BufClientFrameNum int32
}

func NewRoom() *Room {
	room := &Room{
		roomid:            0,
		isfull:            false,
		playernum:         0,
		players:           make([]PlayerList, RoomPeople),
		CurClientFrameNum: 0,
		BufClientFrameNum: 0,
	}
	return room
}
func (room *Room) Clear() {
	room.roomid = 0
	room.isfull = false
	room.playernum = 0
	for i := 0; i < len(room.players); i++ {
		room.players[i].PlayerID = -1
	}
}

func (room *Room) CopyRoom(cacheRoom *Room) {
	room.roomid = cacheRoom.roomid
	room.isfull = cacheRoom.isfull
	room.playernum = cacheRoom.playernum
	for i := 0; i < len(cacheRoom.players); i++ {
		room.players[i] = cacheRoom.players[i]
	}
}

//游戏初始化  单独通知玩家所属座位编号	光线信息
//其余群发通知	每隔t时间收集玩家操作 广播给所有房间内玩家
//通知消息需要的数据	4个位置playerid	4个socket

//多线程执行房间 not use now
func (room *Room) RoomRun() {
	//通知当前房间玩家匹配成功
	//room.RoomInform()
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
	if !room.isfull {
		for i := 0; i < len(room.players); i++ {
			if room.players[i].PlayerID == -1 {
				room.players[i].PlayerID = playerid
				room.players[i].PlayerClient = client
				break
			}
		}
	}
	room.playernum++

	if room.playernum == RoomPeople {
		RoomCache.roomid = NextRoomID
		RoomMng[NextRoomID] = NewRoom()
		RoomMng[NextRoomID].CopyRoom(&RoomCache)
		//ChanMap[NextRoomID] = make(chan []byte)
		NextRoomID++
		RoomCache.Clear()
		RoomMng[NextRoomID-1].RoomInform()
		//go RoomMng[NextRoomID-1].RoomRun()
	}
	RoomCacheMu.Unlock()
}

func (room *Room) findPlayerByID() {

}

//cache room调用
func RemovePlayer() {

}

func (room *Room) RoomInform() {

	match := DTO.MatchSuccessDTO{}
	match.Roomid = room.roomid

	// 暂时写死,后期读表
	match.Speed = 10
	match.Count = 99
	match.X = Tools.RandFloat(-1, 1, 2)
	match.Z = Tools.RandFloat(-1, 1, 2)

	match.Players = make([]*DTO.Player, RoomPeople)
	for i := int32(0); i < RoomPeople; i++ {
		match.Players[i] = new(DTO.Player)
	}

	for i := 0; i < len(room.players); i++ {
		match.Players[i].Playerid = room.players[i].PlayerID
		match.Players[i].Name = room.players[i].Name
		match.Players[i].Roleid = room.players[i].PlayerRole
		match.Players[i].Seat = int32(i + 1)
	}
	data, _ := proto.Marshal(&match)
	encode := NetFrame.NewEncode(int32(8+match.XXX_Size()), 2, 4)
	encode.Write()
	var buffer bytes.Buffer
	buffer.Write(encode.GetBytes())
	buffer.Write(data)
	for i := 0; i < len(room.players); i++ {
		room.players[i].PlayerClient.Write(buffer.Bytes())
	}
	fmt.Println("inform ok")
}

func (room *Room) RoomBroad(move *DTO.MoveDTO) {

	fmt.Println("room broad")
	encode := NetFrame.NewEncode(int32(8+move.XXX_Size()), 3, 2)
	encode.Write()
	data, _ := proto.Marshal(move)
	var buffer bytes.Buffer
	buffer.Write(encode.GetBytes())
	buffer.Write(data)
	for i := 0; i < len(room.players); i++ {
		room.players[i].PlayerClient.Write(buffer.Bytes())
	}
	fmt.Println("broad ok")
}
