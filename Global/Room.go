package Global

import (
	"../NetFrame"
	"../Tools"
	"../proto/dto"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"net"
	"sync"
	"time"
)

type PlayerList struct {
	PlayerID     int32
	Name         string
	PlayerClient *ClientState
	PlayerRole   int32
	IsLeave      bool
	IsDead       bool
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
	RoomLivePeople  int32     //房间存活人数
	timer           *time.Timer
	RoomChan        chan *DTO.ServerMoveDTO
}

func NewRoom() *Room {
	room := &Room{
		roomid:         0,
		isfull:         false,
		playernum:      0,
		players:        make([]PlayerList, RoomPeople),
		CacheMsg:       make([]DTO.ClientMoveDTO, RoomPeople),
		CacheMsgIndex:  0,
		RoomLivePeople: RoomPeople,
	}
	return room
}
func (room *Room) Clear() {
	room.roomid = 1
	room.isfull = false
	room.playernum = 0
	for i := int32(0); i < RoomPeople; i++ {
		room.players[i].PlayerID = -1
	}
}

func (room *Room) CopyRoom(cacheRoom *Room) {
	room.roomid = cacheRoom.roomid
	room.isfull = cacheRoom.isfull
	room.playernum = cacheRoom.playernum
	for i := int32(0); i < RoomPeople; i++ {
		room.players[i] = cacheRoom.players[i]
		room.CacheMsg[i].Seat = -1
	}
}

//cache room调用
func (room *Room) InsertPlayer(playerID int32, playerRole int32, playername string, client *ClientState) {
	//roomFull
	RoomCacheMu.Lock()
	if !room.isfull {
		for i := int32(0); i < RoomPeople; i++ {
			if room.players[i].PlayerID == -1 {
				room.players[i].PlayerID = playerID
				room.players[i].PlayerRole = playerRole
				room.players[i].PlayerClient = client
				room.players[i].Name = playername
				room.players[i].IsLeave = false
				room.players[i].IsDead = false
				room.playernum++
				break
			}
		}
	}
	if room.playernum == RoomPeople {
		RoomCache.roomid = NextRoomID
		RoomMng[NextRoomID] = NewRoom()
		RoomMng[NextRoomID].CopyRoom(&RoomCache)
		NextRoomID++
		RoomCache.Clear()
		RoomMng[NextRoomID-1].RoomInform()
	} else {
		room.RoomMatchInform(int32(DTO.MatchTypes_ENTER_SRES), client.Client)
	}
	RoomCacheMu.Unlock()
}

//cache room调用
func (room *Room) RemovePlayer(playerid int32, client net.Conn) {
	RoomCacheMu.Lock()
	for i := int32(0); i < RoomPeople; i++ {
		if room.players[i].PlayerID == playerid {
			room.players[i].PlayerID = -1
			room.playernum--
			break
		}
	}
	RoomCache.RoomMatchInform(int32(DTO.MatchTypes_LEAVE_SRES), client)
	RoomCacheMu.Unlock()
}

//如果是匹配确认需要返回roomid 如果是取消匹配确认 不需要返回特殊数据
func (room *Room) RoomMatchInform(command int32, client net.Conn) {
	any := DTO.MatchRtnDTO{}
	any.Cacheroomid = 1
	data, _ := proto.Marshal(&any)
	buffer := NetFrame.WriteMessage(int32(DTO.MsgTypes_TYPE_MATCH), command, data, any.XXX_Size())
	client.Write(buffer.Bytes())
}

//初始化房间数据
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

	for i := int32(0); i < RoomPeople; i++ {
		match.Players[i].Playerid = room.players[i].PlayerID
		match.Players[i].Name = room.players[i].Name
		match.Players[i].Roleid = room.players[i].PlayerRole
		match.Players[i].Seat = i + 1
		room.players[i].PlayerClient.IsMatch = false
		room.players[i].PlayerClient.IsFight = true
		room.players[i].PlayerClient.RoomSeat = i + 1
		room.players[i].PlayerClient.RoomID = room.roomid
	}
	data, _ := proto.Marshal(&match)
	buffer := NetFrame.WriteMessage(int32(DTO.MsgTypes_TYPE_MATCH), int32(DTO.MatchTypes_ENTER_SELECT_BRO), data, match.XXX_Size())
	for i := int32(0); i < RoomPeople; i++ {
		room.players[i].PlayerClient.Client.Write(buffer.Bytes())
		DetailedLog.Detailed(room.players[i].PlayerID, Fight_IN)
	}
	log.Println("inform ok")

	room.RoomChan = make(chan *DTO.ServerMoveDTO, 100000)
	//把新开的房间信息存入数据库
	AddRoom(RoomCollection, room.roomid, &match)
	go AddFrame(RoomCollection, room.roomid, room.RoomChan)
}

func (room *Room) RoomBroad() {
	send := DTO.ServerMoveDTO{}
	send.Bagid = room.CacheMsg[0].Bagid
	//需要自己分配内存
	TmpClientMoveDTO := make([]DTO.ClientDTO, RoomPeople)
	send.ClientInfo = make([]*DTO.ClientDTO, RoomPeople)
	num := int32(0)
	for i := int32(0); i < RoomPeople; i++ {
		send.ClientInfo[i] = &TmpClientMoveDTO[i]
		send.ClientInfo[i].Msg = make([]*DTO.FrameInfo, FramesPerBag)
		if room.CacheMsg[i].Seat != -1 {
			send.ClientInfo[i].Seat = room.CacheMsg[i].Seat
			send.ClientInfo[i].Msg = room.CacheMsg[i].Msg
			num++
		} else {
			send.ClientInfo[i].Seat = -1
			break
		}
	}
	//DetailedLog.Log.Info("本次包的数量：", num)
	//把要广播的内容写成字节流
	data, _ := proto.Marshal(&send)
	buffer := NetFrame.WriteMessage(int32(DTO.MsgTypes_TYPE_FIGHT), int32(DTO.FightTypes_INFORM_SRES), data, send.XXX_Size())

	for i := int32(0); i < RoomPeople; i++ {
		if !room.players[i].IsLeave {
			room.players[i].PlayerClient.Client.Write(buffer.Bytes())
		}
	}

	//把广播的若干帧存入数据库，有可能碰到房间数据库未创建完毕就开始插入数据了
	//AddFrame(RoomCollection, room.roomid, &send)
	room.RoomChan <- &send
	//广播完后重置缓存消息和时间
	room.ClearCacheMsg()
}

//客户端发来一个包，当缓存中包的数量为RoomPeople或者距离上一次发送过了9毫秒 就广播一次
func (room *Room) InsertMsg(move *DTO.ClientMoveDTO) {
	room.CacheMsgIndexMu.Lock()

	if room.CacheMsgIndex == 0 {
		room.StartTime = time.Now()
		room.timer = time.NewTimer(WaitMS * time.Millisecond)
		go func() {
			select {
			case <-room.timer.C:
				room.RoomBroad()
			}
		}()
	}
	room.CacheMsg[room.CacheMsgIndex] = *move
	room.CacheMsgIndex++

	if room.CacheMsgIndex == room.RoomLivePeople {
		room.RoomBroad()
		room.timer.Stop()
	}
	room.CacheMsgIndexMu.Unlock()
}

func (room *Room) ClearCacheMsg() {
	for i := int32(0); i < RoomPeople; i++ {
		room.CacheMsg[i].Seat = -1
	}
	room.CacheMsgIndex = 0
}

func (room *Room) PlayerDeath(seat int32) {
	room.players[seat-1].IsDead = true
	room.RoomLivePeople--
}

func (room *Room) PlayerLeave(seat int32) {
	room.players[seat-1].IsLeave = true
	room.players[seat-1].PlayerClient.IsFight = false
	if !room.players[seat-1].IsDead {
		room.RoomLivePeople--
	}
	log.Println("leave")
	room.playernum--
	if room.playernum == 0 {
		close(room.RoomChan)
		delete(RoomMng, room.roomid) //清除map，实际内存释放通过GC
	}
}
