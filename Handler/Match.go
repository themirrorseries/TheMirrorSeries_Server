package Handler

import (
	"../Global"
	"../proto/dto"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

type Match struct {
	data *HandlerData
}

func (match *Match) ReceiveMessage() {
	switch match.data.command {
	case int32(DTO.MatchTypes_ENTER_CREQ):
		//申请进入匹配
		match.matchStart()
		break
	case int32(DTO.MatchTypes_LEAVE_CREQ):
		//申请离开匹配
		match.matchEnd()
		break
	default:
		log.Println("其他错误")
		Global.DetailedLog.Log.Warn("客户端匹配请求出错！")
		break
	}
}

//玩家申请匹配，将玩家加入匹配池中
func (match *Match) matchStart() {
	log.Println("match start")
	match.data.client.IsMatch = true
	any := DTO.MatchDTO{}
	proto.Unmarshal(match.data.messages[match.data.bytesStart:match.data.bytesEnd], &any)
	match.data.client.PlayerID = any.Id
	Global.DetailedLog.Detailed(any.Id, Global.Match_IN)
	Global.RoomCache.InsertPlayer(any.Id, any.RoleID, any.Name, match.data.client)
}

//玩家申请结束匹配，将玩家从匹配房间中去除
func (match *Match) matchEnd() {

	log.Println("match end")
	match.data.client.IsMatch = false
	any := DTO.MatchRtnDTO{}
	proto.Unmarshal(match.data.messages[match.data.bytesStart:match.data.bytesEnd], &any)
	Global.DetailedLog.Detailed(any.Id, Global.Match_OUT)
	Global.RoomCache.RemovePlayer(any.Id, match.data.client.Client)
}
