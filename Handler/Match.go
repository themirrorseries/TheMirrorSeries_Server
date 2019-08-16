package Handler

import (
	"../Global"
	"../proto/dto"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"net"
)

type Match struct {
	command    int32
	messages   []byte
	bytesStart int32
	bytesEnd   int32
	client     net.Conn
}

func NewMatch(c, start, end int32, msg []byte, _client net.Conn) *Match {
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
	case 0:
		//申请进入匹配
		log.Println("start match")
		match.matchStart()
		break
	case 2:
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
	//to do
	log.Println("match start")
	//add to match pool
	//匹配也用anyDTO
	any := DTO.MatchDTO{}
	proto.Unmarshal(match.messages[match.bytesStart:match.bytesEnd], &any)
	Global.RoomCache.InsertPlayer(any.Id, any.RoleID, match.client)
}

func (match *Match) matchEnd() {
	//to do
	log.Println("match end")
	//delete from cache room
	any := DTO.MatchRtnDTO{}
	proto.Unmarshal(match.messages[match.bytesStart:match.bytesEnd], &any)
	Global.RoomCache.RemovePlayer(any.Id, match.client)
}
