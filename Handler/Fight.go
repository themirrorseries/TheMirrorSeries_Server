package Handler

import (
	"../Global"
	"../proto/dto"
	"github.com/golang/protobuf/proto"
	//log "github.com/sirupsen/logrus"
	"net"
)

type Fight struct {
	command    int32
	messages   []byte
	bytesStart int32
	bytesEnd   int32
	client     net.Conn
}

func NewFight(c, start, end int32, msg []byte, _client *Global.ClientState) *Fight {
	fight := &Fight{
		command:    c,
		bytesStart: start,
		bytesEnd:   end,
		messages:   msg,
		client:     _client.Client,
	}
	return fight
}

func (fight *Fight) ReveiveMessage() {
	switch fight.command {
	case int32(DTO.FightTypes_MOVE_CREQ):
		//client move
		fight.move()
	}
}

func (fight *Fight) move() {

	move := DTO.ClientMoveDTO{}
	proto.Unmarshal(fight.messages[fight.bytesStart:fight.bytesEnd], &move)

	Global.RoomMng[move.Roomid].InsertMsg(&move)
}
