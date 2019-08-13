package Global

import (
	"../NetFrame"
	"../Tools"
	"../proto/dto"
	"bytes"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"net"
	//"time"
	"sync"
	"time"
)

type PlayerList struct {
	PlayerID     int32
	Name         string
	PlayerClient net.Conn
	PlayerRole   int32
}

type Room struct {
	roomid          int32 //唯一id
	isfull          bool  //是否满员
	playernum       int32 //人数
	players         []PlayerList
	CacheMsg        []DTO.ClientMoveDTO //缓存消息，达到RoomPeople数量或者达到最长等待时间后广播一次
	CacheMsgIndex   int32
	CacheMsgIndexMu sync.Mutex
	StartTime       time.Time //用于计算两次广播之间的最长等待时间
}

func NewRoom() *Room {
	room := &Room{
		roomid:        0,
		isfull:        false,
		playernum:     0,
		players:       make([]PlayerList, RoomPeople),
		CacheMsg:      make([]DTO.ClientMoveDTO, RoomPeople),
		CacheMsgIndex: 0,
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
		room.CacheMsg[i].Seat = 0
	}
}

//游戏初始化  单独通知玩家所属座位编号	光线信息
//其余群发通知	每隔t时间收集玩家操作 广播给所有房间内玩家
//通知消息需要的数据	4个位置playerid	4个socket

//多线程执行房间 not use now
func (room *Room) RoomRun() {
	//通知当前房间玩家匹配成功
	room.RoomInform()
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
	match.Count = 20
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
	log.Println("inform ok")

	//room.RoomStartInit()
	room.StartTime = time.Now()
}

func (room *Room) RoomBroad() {
	//如果某个玩家的消息没收到	Seat设置为0
	log.Println("room broad")

	//广播信息结构 第一层 帧号 自定CacheMsg[i].Bagid
	// 第二层 4个客户端信息（Seat, FrameInfo）Seat=CacheMsg[i].Seat	FrameInfo CacheMsg[i].FrameInfo

	send := DTO.ServerMoveDto{}
	send.Bagid = room.CacheMsg[0].Bagid

	//需要自己分配内存
	TmpClientMoveDTO := make([]DTO.ClientDto, RoomPeople)
	TmpFrameInfo := make([]DTO.FrameInfo, FramesPerBag)
	send.ClientInfo = make([]*DTO.ClientDto, RoomPeople)
	for i := int32(0); i < RoomPeople; i++ {
		send.ClientInfo[i] = &TmpClientMoveDTO[i]
		send.ClientInfo[i].Msg = make([]*DTO.FrameInfo, 5)
		for j := int32(0); j < FramesPerBag; j++ {
			send.ClientInfo[i].Msg[j] = &TmpFrameInfo[j]
			send.ClientInfo[i].Msg[j].Move = new(DTO.Dir)
			send.ClientInfo[i].Msg[j].SkillDir = new(DTO.Dir)
		}
		if room.CacheMsg[i].Seat != 0 {
			send.ClientInfo[i].Seat = room.CacheMsg[i].Seat
			send.ClientInfo[i].Msg = room.CacheMsg[i].Msg
		} else {
			break
		}
	}
	data, _ := proto.Marshal(&send)
	encode := NetFrame.NewEncode(int32(8+send.XXX_Size()), 3, 2)
	encode.Write()
	var buffer bytes.Buffer
	buffer.Write(encode.GetBytes())
	buffer.Write(data)
	for i := 0; i < len(room.players); i++ {
		room.players[i].PlayerClient.Write(buffer.Bytes())
	}
	log.Println("broad ok")

	//广播完后重置缓存消息和时间
	room.ClearCacheMsg()
	room.StartTime = time.Now()
}

//客户端发来一个包，当缓存中包的数量为RoomPeople或者距离上一次发送过了9毫秒 就广播一次
func (room *Room) InsertMsg(move *DTO.ClientMoveDTO) {
	log.Println("insert bag")
	room.CacheMsgIndexMu.Lock()
	room.CacheMsg[room.CacheMsgIndex] = *move
	room.CacheMsgIndex++
	if room.CacheMsgIndex == RoomPeople || time.Since(room.StartTime) >= time.Duration(time.Millisecond*9) {
		room.RoomBroad()
	}
	room.CacheMsgIndexMu.Unlock()
}

func (room *Room) ClearCacheMsg() {
	for i := int32(0); i < RoomPeople; i++ {
		room.CacheMsg[i].Seat = 0
	}
	room.CacheMsgIndex = 0
}
