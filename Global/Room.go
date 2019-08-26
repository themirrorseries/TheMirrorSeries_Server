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
	roomid         int32 //唯一id
	playernum      int32 //人数
	players        []PlayerList
	cacheMsg       []DTO.ClientMoveDTO //缓存消息，达到RoomPeople数量或者达到最长等待时间后广播一次
	cacheMsgIndex  int32
	cacheMsgMu     sync.Mutex
	RoomLivePeople int32 //房间存活人数
	timer          *time.Timer
	loadUp         []int32
	loadUpNum      int32
	loadUpMu       sync.Mutex
}

func NewRoom() *Room {
	room := &Room{
		roomid:         0,
		playernum:      0,
		players:        make([]PlayerList, RoomPeople),
		cacheMsg:       make([]DTO.ClientMoveDTO, RoomPeople),
		cacheMsgIndex:  0,
		RoomLivePeople: RoomPeople,
	}
	return room
}
func (room *Room) Clear() {
	room.roomid = 1
	room.playernum = 0
	for i := int32(0); i < RoomPeople; i++ {
		room.players[i].PlayerID = -1
	}
}

func (room *Room) CopyRoom(cacheRoom *Room) {
	room.roomid = cacheRoom.roomid
	room.playernum = cacheRoom.playernum
	for i := int32(0); i < RoomPeople; i++ {
		room.players[i] = cacheRoom.players[i]
		room.cacheMsg[i].Seat = -1
	}
}

//玩家申请匹配时 将玩家加入匹配池
func (room *Room) InsertPlayer(playerID int32, playerRole int32, playername string, client *ClientState) {
	RoomCacheMu.Lock()
	if room.playernum < RoomPeople {
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
	//匹配房间全局唯一一个，如果满了，将匹配房间信息拷贝到新房间并使用map索引
	if room.playernum == RoomPeople {
		RoomCache.roomid = NextRoomID
		RoomMng[NextRoomID] = NewRoom()
		RoomMng[NextRoomID].CopyRoom(&RoomCache)
		NextRoomID++
		RoomCache.Clear()
		RoomMng[NextRoomID-1].roomInform()
	} else { //如果没满，通知在匹配中，玩家可取消匹配
		room.roomMatchInform(int32(DTO.MatchTypes_ENTER_SRES), client.Client)
	}
	RoomCacheMu.Unlock()
}

//玩家申请离开匹配时 将玩家移出匹配池
func (room *Room) RemovePlayer(playerid int32, client net.Conn) {
	RoomCacheMu.Lock()
	for i := int32(0); i < RoomPeople; i++ {
		if room.players[i].PlayerID == playerid {
			room.players[i].PlayerID = -1
			room.playernum--
			break
		}
	}
	RoomCache.roomMatchInform(int32(DTO.MatchTypes_LEAVE_SRES), client)
	RoomCacheMu.Unlock()
}

//如果是匹配确认需要返回roomid 如果是取消匹配确认 不需要返回特殊数据
func (room *Room) roomMatchInform(command int32, client net.Conn) {
	any := DTO.MatchRtnDTO{}
	any.Cacheroomid = 1
	data, _ := proto.Marshal(&any)
	buffer := NetFrame.WriteMessage(int32(DTO.MsgTypes_TYPE_MATCH), command, data, any.XXX_Size())
	client.Write(buffer.Bytes())
}

//初始化房间数据
func (room *Room) roomInform() {
	room.loadUpNum = 0
	room.loadUp = make([]int32, RoomPeople)
	for i := int32(0); i < RoomPeople; i++ {
		room.loadUp[i] = -1
	}
	match := DTO.MatchSuccessDTO{}
	match.Roomid = room.roomid
	// 暂时写死,后期读表
	//match.Speed = 10
	//match.Count = 20
	//match.X = Tools.RandFloat(-1, 1, 2)
	//match.Z = Tools.RandFloat(-1, 1, 2)

	match.Count = 20
	match.Speed = 10
	match.Lights = make([]*DTO.LightDTO, RoomPeople)
	match.Players = make([]*DTO.PlayerDTO, RoomPeople)
	for i := int32(0); i < RoomPeople; i++ {
		match.Players[i] = new(DTO.PlayerDTO)
		//DTO.Light为protobuf生成，有额外几个参数，不能直接初始化
		match.Lights[i] = new(DTO.LightDTO)
		match.Lights[i].X = Tools.RandFloat(-1, 1, 2)
		match.Lights[i].Z = Tools.RandFloat(-1, 1, 2)
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

	//room.RoomChan = make(chan *DTO.ServerMoveDTO, 100000)
	//把新开的房间信息存入数据库
	AddRoom(RoomCollection, room.roomid, &match)
	//go AddFrame(RoomCollection, room.roomid, room.RoomChan)
}

func (room *Room) RoomBroad() {
	send := DTO.ServerMoveDTO{}
	send.Bagid = room.cacheMsg[0].Bagid
	//需要自己分配内存
	TmpClientMoveDTO := make([]DTO.ClientDTO, RoomPeople)
	send.ClientInfo = make([]*DTO.ClientDTO, RoomPeople)
	num := int32(0)
	for i := int32(0); i < RoomPeople; i++ {
		send.ClientInfo[i] = &TmpClientMoveDTO[i]
		send.ClientInfo[i].Msg = make([]*DTO.FrameInfo, FramesPerBag)
		if room.cacheMsg[i].Seat != -1 {
			send.ClientInfo[i].Seat = room.cacheMsg[i].Seat
			send.ClientInfo[i].Msg = room.cacheMsg[i].Msg
			num++
		} else {
			send.ClientInfo[i].Seat = -1
			break
		}
	}

	data, _ := proto.Marshal(&send)
	buffer := NetFrame.WriteMessage(int32(DTO.MsgTypes_TYPE_FIGHT), int32(DTO.FightTypes_INFORM_SRES), data, send.XXX_Size())
	for i := int32(0); i < RoomPeople; i++ {
		if !room.players[i].IsLeave {
			room.players[i].PlayerClient.Client.Write(buffer.Bytes())
		}
	}
	room.ClearCacheMsg()
}

//客户端发来一个包，当缓存中包的数量为RoomPeople或者距离上一次发送过了9毫秒 就广播一次
func (room *Room) InsertMsg(move *DTO.ClientMoveDTO) {
	room.cacheMsgMu.Lock()

	if room.cacheMsgIndex == 0 {
		room.timer = time.NewTimer(WaitMS * time.Millisecond)
		go func() {
			select {
			case <-room.timer.C:
				room.RoomBroad()
			}
		}()
	}
	room.cacheMsg[room.cacheMsgIndex] = *move
	room.cacheMsgIndex++

	if room.cacheMsgIndex == room.RoomLivePeople {
		room.timer.Stop()
		room.RoomBroad()
	}
	room.cacheMsgMu.Unlock()
}

func (room *Room) ClearCacheMsg() {
	for i := int32(0); i < RoomPeople; i++ {
		room.cacheMsg[i].Seat = -1
	}
	room.cacheMsgIndex = 0
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
		//close(room.RoomChan)
		delete(RoomMng, room.roomid) //清除map，实际内存释放通过GC
	}
}

func (room *Room) AddLoadPeople(seat int32) {
	room.loadUpMu.Lock()
	if room.loadUp[seat-1] == -1 {
		room.loadUp[seat-1] = 1
		room.loadUpNum++
	}
	log.Info("load ok")
	if room.loadUpNum == RoomPeople {
		log.Info("inform load ok")
		send := DTO.AnyDTO{}
		data, _ := proto.Marshal(&send)
		buffer := NetFrame.WriteMessage(int32(DTO.MsgTypes_TYPE_FIGHT), int32(DTO.FightTypes_LOAD_UP_SREQ), data, send.XXX_Size())
		for i := int32(0); i < RoomPeople; i++ {
			room.players[i].PlayerClient.Client.Write(buffer.Bytes())
		}
	}
}
