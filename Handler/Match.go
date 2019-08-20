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
		log.Println("start match")
		match.matchStart()
		break
	case int32(DTO.MatchTypes_LEAVE_CREQ):
		//申请离开匹配
		log.Println("live match")
		match.matchEnd()
		break
	default:
		log.Println("其他错误")
		break
	}
}
func (match *Match) matchStart() {
	log.Println("match start")
	match.data.client.IsMatch = true
	any := DTO.MatchDTO{}
	proto.Unmarshal(match.data.messages[match.data.bytesStart:match.data.bytesEnd], &any)
	match.data.client.PlayerID = any.Id
	Global.RoomCache.InsertPlayer(any.Id, any.RoleID, any.Name, match.data.client)
}

func (match *Match) matchEnd() {

	log.Println("match end")
	match.data.client.IsMatch = false
	any := DTO.MatchRtnDTO{}
	proto.Unmarshal(match.data.messages[match.data.bytesStart:match.data.bytesEnd], &any)
	Global.RoomCache.RemovePlayer(any.Id, match.data.client.Client)
}
