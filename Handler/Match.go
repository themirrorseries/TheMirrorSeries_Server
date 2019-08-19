package Handler

import (
	"../Global"
	"../proto/dto"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

type Match struct {
	command    int32
	messages   []byte
	bytesStart int32
	bytesEnd   int32
	client     *Global.ClientState
}

func NewMatch(c, start, end int32, msg []byte, _client *Global.ClientState) *Match {
	match := &Match{
		command:    c,
		bytesStart: start,
		bytesEnd:   end,
		messages:   msg,
		client:     _client,
	}
	return match
}

func (match *Match) ReveiveMessage() {
	log.Println(match.command)
	switch match.command {
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
	match.client.IsMatch = true
	any := DTO.MatchDTO{}
	proto.Unmarshal(match.messages[match.bytesStart:match.bytesEnd], &any)
	match.client.PlayerID = any.Id
	Global.RoomCache.InsertPlayer(any.Id, any.RoleID, match.client)
}

func (match *Match) matchEnd() {

	log.Println("match end")
	match.client.IsMatch = false
	any := DTO.MatchRtnDTO{}
	proto.Unmarshal(match.messages[match.bytesStart:match.bytesEnd], &any)
	Global.RoomCache.RemovePlayer(any.Id, match.client.Client)
}
